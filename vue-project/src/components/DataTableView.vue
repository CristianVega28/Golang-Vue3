<template>
    <div class="table-responsive">

      <table class="table table-dark">
        <thead>
  
          <tr>
            <th scope="col">#</th>
            <th scope="col">Customer ID</th>
            <th scope="col">Fullname</th>
            <th scope="col">Birthdate</th>
            <th scope="col">Address Customer</th>
            <th scope="col">Sector Customer</th>
            <th scope="col">Postal Code Customer</th>
            <th scope="col">Phone</th>
            <th scope="col">Email</th>
            <th scope="col">Dischargedate</th>
            <th scope="col">Group Customer</th>
          </tr>
        </thead>
        <tbody>
            
          <tr v-for="(item, key) in clientsState.slice(0,3)">
            <th scope="row"> {{ key }}</th>
            <td> {{ item.customer_id }}</td>
            <td> {{ item.full_name }}</td>
            <td> {{ item.birth_date }}</td>
            <td> {{ item.customer_address }}</td>
            <td> {{ item.customer_sector }}</td>
            <td> {{ item.customer_postal_code }}</td>
            <td> {{ item.phone }}</td>
            <td> {{ item.email }}</td>
            <td> {{ item.email }}</td>
            <td> {{ item.discharge_date }}</td>
            <td> {{ item.customer_group }}</td>
          </tr>
        </tbody>
      </table>
    </div>

</template>

<script setup lang="ts">
interface Clientes {
    customer_id : string,
    full_name : string,
    birth_date : string,
    customer_address : string,
    customer_sector: string,
    customer_postal_code: string,
    phone: number,
    email: string,
    discharge_date: string,
    customer_group: string
}

import { onBeforeMount, reactive } from 'vue'

const clientsState = reactive<Clientes[]>([])
//* Traer datos mientras el componentes se esta montando
const getData = async () => {
    try {
       let data = await fetch(`${import.meta.env.VITE_BACKEND_URL}\\clients`)
       let json: Clientes[] = await data.json()
       json.forEach(cliente => clientsState.push(cliente))
    } catch (error) {
        
    }
}

onBeforeMount(() => {
    getData()
    console.log(import.meta.env.VITE_BACKEND_URL);
    
})



</script>