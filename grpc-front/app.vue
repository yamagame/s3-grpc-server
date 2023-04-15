<script setup>
const host = import.meta.env.VITE_HOST || "localhost"
const port = import.meta.env.VITE_PORT || 3000
const protocol = import.meta.env.VITE_PROTOCOL || "http"
const baseurl = `${protocol}://${host}:${port}`
const { data } = await useFetch(`${baseurl}/api/listBuckets`)
</script>

<template>
  {{ data }}
</template>
