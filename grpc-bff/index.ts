import * as express from "express"
import * as aws from "./proto/grpc-server/aws"
import * as grpc from "@grpc/grpc-js"
import { UnaryCallback } from "@grpc/grpc-js/build/src/client"

const port = 7000
const grpcHost = process.env.GRPC_HOST || `localhost:50051`

const app = express()

app.use(express.json())
app.use(express.urlencoded({ extended: true }))

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

app.use((req, res, next) => {
  console.log(req.method, req.url)
  console.log(req.body)
  next()
})

app.get("/api", (req, res) => {
  console.log(req.method, req.url)
  res.json({ count: 100, name: "test name" })
})

app.post("/api/createBucket", async (req, res) => {
  const ret = await awsClient.CreateBucket({})
  res.json(ret)
})

app.get("/api/listBuckets", async (req, res) => {
  const ret = await awsClient.ListBuckets({})
  res.json(ret)
})

app.post("/api/putObject/:key", async (req, res) => {
  if (!req.body) {
    return res.json({})
  }
  const { key } = req.params
  const { content } = req.body
  console.log(content)
  const ret = await awsClient.PutObject({ key, content })
  res.json(ret)
})

app.post("/api/deleteObject/:key", async (req, res) => {
  const { key } = req.params
  const ret = await awsClient.DeleteObject({ key })
  res.json(ret)
})

app.get("/api/getObject/:key", async (req, res) => {
  const { key } = req.params
  const ret = await awsClient.GetObject({ key })
  res.json(ret)
})

app.get("/", (req, res) => {
  console.log(req.method, req.url)
  res.send("Hello World!")
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})
