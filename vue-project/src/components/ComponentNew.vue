<script setup lang="ts">
import { ref } from 'vue'
const frutas: string[] = ["manzana", "pera", "platano", "cereza"];
</script>



<template>
  <div class="data">
    <ul>
      <li  v-for="(fruta, index) in frutas" :key="index">
        {{ fruta }}
      </li>
    </ul>
    <button @click="$emit('enlarge-txt')">
        alarge
    </button>
    <slot />
  </div>
  <div :class="$attrs.class">
    div con la clase active 
  </div>
</template>
