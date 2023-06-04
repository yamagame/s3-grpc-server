<script setup>
const host = import.meta.env.VITE_HOST || "localhost"
const port = import.meta.env.VITE_PORT || 3000
const protocol = import.meta.env.VITE_PROTOCOL || "http"
const baseurl = `${protocol}://${host}:${port}`
const useKeycloak = false

const listObjectsRef = ref({})
const getBucketsRef = ref({})

async function callApi(url, options) {
  const { data, pending, error, refresh } = await useFetch(url, options)
  if (error.value != null) {
    if (useKeycloak) {
      login()
    }
    return
  }
  return data
}

async function listObjects() {
  const data = await callApi(`${baseurl}/api/listObjects`, {
    redirect: "error",
  })
  if (data) listObjectsRef.value = data.value[0].keys
}
async function getObject(key) {
  const data = await callApi(`${baseurl}/api/getObject/${key}`, {
    redirect: "error",
  })
  if (data) getBucketsRef.value = data.value
}
async function select(itemname) {
  getObject(itemname)
}
</script>

<template>
  <div style="padding: 0px 0px 8px 0px; border-bottom: solid 1px #E0E0E0;">
    <button style="margin: 4px" @click="listObjects">listObjects</button>
    <button style="margin: 4px" @click="listObjects">listObjects</button>
  </div>
  <div style="font-size: 0.8em; margin: 4px; display: flex;">
    <div style="flex: 1; border-right: solid 1px #E0E0E0;">
      <div v-for="value in listObjectsRef" :key="value" @click="select(value)">
        {{ value }}
      </div>
    </div>
    <div style="flex: 4; margin: 4px;">
      <div v-for="(value, key) in getBucketsRef" :key="value.key" @click="select(value)">
        <span style="color: #8080FF">{{ key }}</span> : {{ value }}
      </div>
    </div>
  </div>
</template>
