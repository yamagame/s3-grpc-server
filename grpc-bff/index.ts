import * as express from "express"
import * as storage from "./proto/grpc-server/storage"
import * as grpc from "@grpc/grpc-js"
import { UnaryCallback } from "@grpc/grpc-js/build/src/client"
import * as session from "express-session"
import * as Keycloak from "keycloak-connect"
const cors = require("cors")

const port = 7000
const useKeycloak = false
const grpcHost = process.env.GRPC_HOST || `localhost:50051`
const appHost = process.env.APP_HOST || `http://localhost:3000`
const memoryStore = new session.MemoryStore()

const keycloak = (() => {
  if (useKeycloak) {
    const keycloak = new Keycloak({ store: memoryStore }, "./keycloak.json")
    const config = keycloak.getConfig()
    config["realmUrl"] = "http://localhost:8180/realms/test"
    return keycloak
  }
  return null
})()

const app = express()

app.use(
  session({
    secret: "my-bff-secret",
    resave: false,
    saveUninitialized: true,
    store: memoryStore,
  })
)

app.use(cors())
app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.use((req, res, next) => {
  console.log(req.method, req.url, req.hostname)
  console.log(req.headers)
  console.log(req.body)
  req.headers.host = `${req.headers["x-forwarded-host"]}`
  next()
})

if (useKeycloak) {
  app.set("trust proxy", true)
  app.use(keycloak.middleware({ logout: "/api/logoff" }))
}

function errorHandler(err, req, res, next) {
  res.status(500)
  res.render("error", { error: err })
}

app.use(errorHandler)

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

const conn = new grpc.Client(grpcHost, grpc.ChannelCredentials.createInsecure())

type RpcImpl = (service: string, method: string, data: Uint8Array) => Promise<Uint8Array>

const sendRequest: RpcImpl = (service, method, data) => {
  // Conventionally in gRPC, the request path looks like
  //   "package.names.ServiceName/MethodName",
  // we therefore construct such a string
  const path = `/${service}/${method}`

  return new Promise((resolve, reject) => {
    // makeUnaryRequest transmits the result (and error) with a callback
    // transform this into a promise!
    const resultCallback: UnaryCallback<any> = (err, res) => {
      if (err) {
        return reject(err)
      }
      resolve(res)
    }

    function passThrough(argument: any) {
      return argument
    }

    // Using passThrough as the serialize and deserialize functions
    conn.makeUnaryRequest(path, passThrough, passThrough, data, resultCallback)
  })
}

const rpc: Rpc = { request: sendRequest }
const storageClient = new storage.storageClientImpl(rpc, { service: "storage" })

const protect = () => {
  if (useKeycloak) {
    return keycloak.protect
  }
  return (req, res, next) => {
    next()
  }
}

app.get("/api/login", protect(), async (req, res, next) => {
  res.redirect(appHost)
})

app.post("/api/createBucket", protect(), async (req, res, next) => {
  try {
    const ret = await storageClient.CreateBucket({})
    res.json(ret)
  } catch (err) {
    next(err)
  }
})

app.get("/api/listBuckets", protect(), async (req, res, next) => {
  try {
    const ret = await storageClient.ListBuckets({})
    res.json(ret)
  } catch (err) {
    next(err)
  }
})

app.post("/api/putObject/*", protect(), async (req, res, next) => {
  try {
    if (!req.body) {
      return res.json({})
    }
    const key = req.params[0]
    const { content } = req.body
    console.log(content)
    const ret = await storageClient.PutObject({ key, content })
    res.json(ret)
  } catch (err) {
    next(err)
  }
})

app.post("/api/deleteObject/*", protect(), async (req, res, next) => {
  try {
    const key = req.params[0]
    const ret = await storageClient.DeleteObject({ key })
    res.json(ret)
  } catch (err) {
    next(err)
  }
})

app.get("/api/getObject/*", protect(), async (req, res, next) => {
  try {
    const key = req.params[0]
    console.log("getObject", key)
    const ret = await storageClient.GetObject({ key })
    res.json(ret)
  } catch (err) {
    next(err)
  }
})

app.get("/", (req, res) => {
  res.send("Hello World!")
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})
