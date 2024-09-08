<template>
  <div class="container m-5">
    <label for="" class="mb-2">Tipo de tabla</label>
    <select  @change="getDataTypeTable($event, null)" class="form-select" v-model="selectOption" placeholder="Tipo de tabla" aria-label="Default select example">
      <option value="all" selected>Tabla con todo los datos</option>
      <option value="page" >Tabla con paginación</option>
      <option value="id" >Tabla con buscador de ID</option>
    </select>
  </div>
    <div class="table-responsive">
      <nav>
        <div class="container">
          <div class="col-3" v-if="computedStateTypeTable.pagination">
            <label for="">Cantidad de usuarios</label>
            <select @change="changeAmountPagination($event)" class="form-select-sm ml-2" v-model="statePagination.amount" aria-label="Default select example">
              <option selected value="5">5</option>
              <option value="10">10</option>
              <option value="15">15</option>
              <option value="20">20</option>
            </select>
          </div>
        </div>
      </nav>
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
            <th scope="col">Discharge Date</th>
            <th scope="col">Group Customer</th>
          </tr>
        </thead>
        <tbody>
            
          <tr v-for="(item, key) in clientsState">
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
      <footer>
        <ul class="pagination justify-content-center">
          <li class="page-item disabled">
            <a @click.prevent href="" class="page-link pe-auto">Previo</a>
          </li>
          <template v-for="number in statePagination.numbers">
            <li class="page-item">
              <a @click.prevent href="" class="page-link pe-auto active">{{ number }}</a>
            </li>
          </template>
          <li class="page-item">
            <a @click.prevent href="" class="page-link pe-auto">Siguiente</a>
          </li>
        </ul>
      </footer>
    </div>

</template>

<script setup lang="ts">
import { ObjectDataQuery, type Clientes, type QueriesUrl, type ResultPaginationFetch} from './types'
import { onBeforeMount, ref, computed } from 'vue'

const selectOption = ref("all")

const computedStateTypeTable = computed<QueriesUrl>(() => {
  const defaultdata = {
  pagination: false,
  allData: false,
  ById: false
}
  switch (selectOption.value) {
    case "all":
      return {
        ...defaultdata,
        allData: true, 
      }
      break;
    
    case "page":
      return {
        ...defaultdata,
        pagination: true
      }
    
    case "id":
      return {
        ...defaultdata,
        ById: true
      }
    default:
      return defaultdata
      break;
  }
})

/**
 * ! Pagination 
 * page (default) = 10
 * amount (defaultl) = 7
 */

const statePagination = ref({
  page: 1,
  pages_total: 0,
  numbers: [] as number[],
  amount: 5,
})
let clientsState = ref<Clientes[]>([])
//* Traer datos mientras el componentes se esta montando
const getData = async <T>(query: string) => {
    /**
     * @query -> page (default) = 1
     * @query -> amount (default) = 7
     */
    try {
       let data = await fetch(`${import.meta.env.VITE_BACKEND_URL}\\clients${query}`)
       let json: T = await data.json()
       return json;
    } catch (error) {
        
    }
}

const queryData = (object_query ?: QueriesUrl , object_data?: ObjectDataQuery ) => {
  if (object_query && object_data) {
    if (object_query.pagination && object_data.pagination) {
      let { amount, page } = object_data.pagination
      return `?page=${page}&&amount=${amount}`
    } else if (object_query.ById && object_data.pagination) {
      let id = object_data.byid
      return `?id=${id}`
    }     
  } 

  return ''
}

const getDataTypeTable = async (event: Event, type_optional : string | null) => {
  const selectHtml = event.target as HTMLSelectElement
  let type = type_optional ? type_optional : selectHtml.value
  switch (type) {
    case 'all':
     let query = queryData() 
     getData(query)
      break;
    
    case 'page':
      clientsState.value = []
      statePagination.value.numbers = []
      let query_page = queryData(computedStateTypeTable.value, { pagination: statePagination.value} )
      let dataPagination = await getData<ResultPaginationFetch>(query_page)
      dataPagination?.data_pagination.forEach(pagination => clientsState.value.push(pagination))
      if (dataPagination?.total_rows) {
        statePagination.value.pages_total = Math.round(dataPagination.total_rows / statePagination.value.amount)

        for (let index = 1; index < statePagination.value.pages_total; index++) {
          console.log(index);
          
          statePagination.value.numbers.push(index) 
        }
        
      }
      console.log(statePagination.value.numbers);
      
      break;

    default:
      break;
  }
}

const changeAmountPagination = (event: Event) => {
  getDataTypeTable(event, 'page')
}

onBeforeMount(async () => {
    let data = await getData<Clientes[]>('')
    data?.forEach(cliente => clientsState.value.push(cliente))
})



</script>