<script setup lang="ts">
import { reactive } from "vue";

interface Formdata {
  name: string;
  job: string;
}


/**
 * Headers pagination
 *! X-total-Count: 10
 */
const form_data = reactive<Formdata>({
  name: "",
  job: "",
});

const sentdata = async () => {
  try {
    let data = await fetch("http://localhost:3000/data/40");
    let json = await data.json()
    console.log(json);
    
  } catch (error) {
    console.log(error);
     
  }
  
};
sentdata()
</script>

<template>
  <main class="container">
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
          <tr>
            <th scope="row"> cristian</th>
            <td> cristian</td>
            <td> cristian</td>
            <td> cristian</td>
            <td> cristian</td>
            <td> cristian</td>
            <td> cristian</td>
            <td> cristian</td>
            <td> cristian</td>
            <td> cristian</td>
            <td> cristian</td>
          </tr>
        </tbody>
      </table>
    </div>
  </main>
</template>
