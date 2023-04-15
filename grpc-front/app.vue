<script setup>
const host = import.meta.env.VITE_HOST || "localhost"
const port = import.meta.env.VITE_PORT || 3000
const protocol = import.meta.env.VITE_PROTOCOL || "http"
const baseurl = `${protocol}://${host}:${port}`

const listBucketsRef = ref({})
const createBucketRef = ref({})
const getObjectRef = ref({})
const putObjectRef = ref({})
const objectKeyRef = ref("test-object")
const objectContentRef = ref("hello world")

const { data } = await useFetch(`${baseurl}/api/listBuckets`)
listBucketsRef.value = data.value

async function createBucket() {
  const { data } = await useFetch(`${baseurl}/api/createBucket`, {
    method: 'POST',
  })
  createBucketRef.value = data.value
}
async function putObject() {
  const key = objectKeyRef.value
  const content = objectContentRef.value
  console.log("putObject", key)
  const { data } = await useFetch(`${baseurl}/api/putObject/${key}`, {
    method: 'POST',
    body: { content, },
  })
  putObjectRef.value = data.value
  console.log(data)
}
async function getObject() {
  const key = objectKeyRef.value
  console.log("createObject", key)
  const { data } = await useFetch(`${baseurl}/api/getObject/${key}`)
  getObjectRef.value = data.value
  console.log(data)
}
</script>

<template>
  <div style="font-size: 0.8em;">
    <div>
      {{ listBucketsRef }}
    </div>
    <div>
      <button @click="createBucket">createBucket</button>
      {{ createBucketRef }}
    </div>
    <div>
      key : <input :value="objectKeyRef" @input="event => objectKeyRef = event.target.value" />
    </div>
    <div>
      Content : <input :value="objectContentRef" @input="event => objectContentRef = event.target.value" />
    </div>
    <div>
      <button @click="putObject">putObject</button>
      {{ putObjectRef }}
    </div>
    <div>
      <button @click="getObject">getObject</button>
      {{ getObjectRef }}
    </div>
  </div>
</template>
