<script setup lang="ts">
import { reactive } from "vue";

interface Formdata {
  name: string;
  job: string;
}

const form_data = reactive<Formdata>({
  name: "",
  job: "",
});

const sentdata = async () => {
  try {
    let data = await fetch("http://localhost:3000/data/cr");
    let json = await data.json()
    console.log(json);
    
  } catch (error) {
    console.log(error);
     
  }
  
};
sentdata()
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        One
      </div>
      <div class="col">
        Onea
      </div>
      <div class="w-100">
      </div>
      <div class="col">
        One
      </div>
    </div>
  </div>
</template>
