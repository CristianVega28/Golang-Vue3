<script setup lang="ts">
import { ref, reactive } from "vue";

interface Fomrdata {
  name: string
  job: string
} 

const form_data = reactive<Fomrdata>({
name:"",
job: ""
})

const sentdata = () => {
  fetch("http://localhost:3000", {
    method: "POST",
    body: JSON.stringify(form_data),
    headers: {
      "Content-Type" : "application/json"
    }
  })
  console.log(JSON.stringify(form_data));
  
}
</script>

<template>
  <form @submit.enter.prevent="sentdata">
    <input type="text" v-model="form_data.name" name="name">
    <input type="text" v-model="form_data.job" name="job">
    <input type="submit" value="Enviar">
  </form>
</template>