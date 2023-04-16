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

async function callApi(url, options) {
  const { data, pending, error, refresh } = await useFetch(url, options)
  if (error.value != null) {
    login()
    return
  }
  return data
}

async function listBuckets() {
  const data = await callApi(`${baseurl}/api/listBuckets`, {
    redirect: "error",
  })
  if (data) listBucketsRef.value = data.value
}
async function createBucket() {
  const data = await callApi(`${baseurl}/api/createBucket`, {
    method: 'POST',
    redirect: "error",
  })
  if (data) createBucketRef.value = data.value
}
async function putObject() {
  const key = objectKeyRef.value
  const content = objectContentRef.value
  const data = await callApi(`${baseurl}/api/putObject/${key}`, {
    method: 'POST',
    body: { content, },
    redirect: "error",
  })
  if (data) putObjectRef.value = data.value
}
async function getObject() {
  const key = objectKeyRef.value
  const data = await callApi(`${baseurl}/api/getObject/${key}`, {
    redirect: "error",
  })
  if (data) getObjectRef.value = data.value
}
async function logout() {
  const { data } = await useFetch(`${baseurl}/api/logoff`, { method: "POST" })
  window.location = `${baseurl}/api/login`
}
async function login() {
  window.location = `${baseurl}/api/login`
}
</script>

<template>
  <div style="font-size: 0.8em;">
    <div>
      <button @click="listBuckets">listBuckets</button>
      {{ listBucketsRef }}
    </div>
    <div>
      <button @click="createBucket">createBucket</button>
      {{ createBucketRef }}
    </div>
    <p>
    <div>
      key : <input :value="objectKeyRef" @input="event => objectKeyRef = event.target.value" />
    </div>
    <div>
      Content : <input :value="objectContentRef" @input="event => objectContentRef = event.target.value" />
    </div>
    </p>
    <div>
      <button @click="putObject">putObject</button>
      {{ putObjectRef }}
    </div>
    <div>
      <button @click="getObject">getObject</button>
      {{ getObjectRef }}
    </div>
    <p>
    <div>
      <button @click="login">Login</button>
    </div>
    <div>
      <button @click="logout">Logout</button>
    </div>
    </p>
  </div>
</template>
