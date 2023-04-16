import * as express from "express"
import * as aws from "./proto/grpc-server/aws"
import * as grpc from "@grpc/grpc-js"
import { UnaryCallback } from "@grpc/grpc-js/build/src/client"
import * as session from "express-session"
import * as Keycloak from "keycloak-connect"
const cors = require("cors")

const port = 7000
const grpcHost = process.env.GRPC_HOST || `localhost:50051`
const appHost = process.env.APP_HOST || `http://localhost:3000`
const memoryStore = new session.MemoryStore()
const keycloak = new Keycloak({ store: memoryStore }, "./keycloak.json")
const config = keycloak.getConfig()
config["realmUrl"] = "http://localhost:8180/realms/test"

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

app.set("trust proxy", true)

app.use(keycloak.middleware({ logout: "/api/logoff" }))

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
const awsClient = new aws.awsClientImpl(rpc, { service: "aws" })

app.get("/api/login", keycloak.protect(), async (req, res, next) => {
  res.redirect(appHost)
})

app.post("/api/createBucket", keycloak.protect(), async (req, res, next) => {
  try {
    const ret = await awsClient.CreateBucket({})
    res.json(ret)
  } catch (err) {
    next(err)
  }
})

app.get("/api/listBuckets", keycloak.protect(), async (req, res, next) => {
  try {
    const ret = await awsClient.ListBuckets({})
    res.json(ret)
  } catch (err) {
    next(err)
  }
})

app.post("/api/putObject/:key", keycloak.protect(), async (req, res, next) => {
  try {
    if (!req.body) {
      return res.json({})
    }
    const { key } = req.params
    const { content } = req.body
    console.log(content)
    const ret = await awsClient.PutObject({ key, content })
    res.json(ret)
  } catch (err) {
    next(err)
  }
})

app.post("/api/deleteObject/:key", keycloak.protect(), async (req, res, next) => {
  try {
    const { key } = req.params
    const ret = await awsClient.DeleteObject({ key })
    res.json(ret)
  } catch (err) {
    next(err)
  }
})

app.get("/api/getObject/:key", keycloak.protect(), async (req, res, next) => {
  try {
    const { key } = req.params
    const ret = await awsClient.GetObject({ key })
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
