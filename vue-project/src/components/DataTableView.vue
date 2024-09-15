<template>
  <div class="container my-5">
    <label for="" class="mb-2">Tipo de tabla</label>
    <select  @change="getDataTypeTable($event, null)" class="form-select" v-model="selectOption" placeholder="Tipo de tabla" aria-label="Default select example">
      <option value="all" selected>Tabla con todo los datos</option>
      <option value="page" >Tabla con paginación</option>
      <option value="id" >Tabla con buscador de ID</option>
    </select>
  </div>
    <div class="table-responsive container">
      <nav class="container">
        <div class="row">
          <div class="col-3" v-if="computedStateTypeTable.pagination">
            <label for="">Cantidad de usuarios</label>
            <select @change="changeAmountPagination($event)" class="form-select-sm ml-2" v-model="statePagination.amount" aria-label="Default select example">
              <option selected value="5">5</option>
              <option value="10">10</option>
              <option value="15">15</option>
              <option value="20">20</option>
            </select>
          </div>
          <div class="container" v-else-if="computedStateTypeTable.ById">
            <div class="row">
              <div class="col-sm-3 col-10">
                <div class="input-group mb-3">
                  <span class="input-group-text" id="basic-addon1">Customer ID</span>
                  <input v-model="idCustomer"  type="text" name="customer-id" class="form-control" placeholder="Customer" aria-label="Username" aria-describedby="basic-addon1">
                </div>
              </div>
              <div class="col-3">
                <button @click="findById()" type="button" class="btn btn-primary">Primary</button>
              </div>
            </div>
          </div>
        </div>
      </nav>
      <table class="table table-dark">
        <thead>
  
          <tr>
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
            <td> {{ item.customer_id }}</td>
            <td> {{ item.full_name }}</td>
            <td> {{ item.birth_date }}</td>
            <td> {{ item.customer_address }}</td>
            <td> {{ item.customer_sector }}</td>
            <td> {{ item.customer_postal_code }}</td>
            <td> {{ item.phone }}</td>
            <td> {{ item.email }}</td>
            <td> {{ item.discharge_date }}</td>
            <td> {{ item.customer_group }}</td>
          </tr>
        </tbody>
      </table>
      <footer>
        <ul class="first-pagination pagination justify-content-center" v-if="computedStateTypeTable.pagination">
          <li class="page-item disabled previous">
            <a @click.prevent class="page-link pe-auto">Previo</a>
          </li> 
          <li class="page-item">
            <a @click.prevent="getPageInPagination($event)" href="" class="page-link pe-auto active">1</a>
          </li>
          <li class="page-item">
            <a @click.prevent="getPageInPagination($event)" href="" class="page-link pe-auto">2</a>
          </li>
          <li class="page-item">
            <a @click.prevent="getPageInPagination($event)" href="" class="page-link pe-auto">3</a>
          </li>
          <li class="page-item next">
            <a @click.prevent href="" class="page-link pe-auto">Siguiente</a>
          </li>
        </ul>
      </footer>
    </div>

</template>

<script setup lang="ts">
import { ObjectDataQuery, type Clientes, type QueriesUrl, type ResultPaginationFetch, type IStatePagination} from './types'
import { onBeforeMount, ref, computed } from 'vue'

const selectOption = ref("all")
const idCustomer = ref<null | string>(null);

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

const statePagination = ref<IStatePagination>({
  page: 1,
  pages_total: 0,
  numbers: [],
  amount: 5,
  slice_page: 1,
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

const getPageInPagination = ($event : MouseEvent) => {
  const currentLink = $event.target as HTMLAnchorElement;
  const parent = document.querySelector('.first-pagination')
  const arrayLink = parent?.querySelector(".active");
  arrayLink?.classList.remove("active")
  currentLink.classList.add("active")

  statePagination.value.page = Number(currentLink.innerHTML)
  updatePaginationRecords()
}

const getDataTypeTable = async (event : Event, type_optional : string | null) => {
  const selectHtml = event.target as HTMLSelectElement
  console.log(selectHtml.value);
  
  let type = type_optional ? type_optional : selectHtml.value
  switch (type) {
    case 'all':
     let query = queryData() 
     let data = await getData<Clientes[]>(query)
    if (data) {
      clientsState.value = data
    }

      break;
    
    case 'page':
      updatePaginationRecords()
      break;
    
    case 'id':
      clientsState.value = []
      break;

    default:
      break;
  }
}

const updatePaginationRecords = async () => {
      clientsState.value = []
      let query_page = queryData(computedStateTypeTable.value, { pagination: statePagination.value} )
      let dataPagination = await getData<ResultPaginationFetch>(query_page)
      dataPagination?.data_pagination.forEach(pagination => clientsState.value.push(pagination))
      if (dataPagination?.total_rows) {
        statePagination.value.pages_total = Math.round(dataPagination.total_rows / statePagination.value.amount)
        calculateRowPagination(statePagination.value)

        
      }

}

/**
 * 
 * @param object 
 * -> array de numeros de paginación
 * [
 * [números menores (4 números)], -> se actualiza cuando llega al ultimo valores del array
 * [icon], -> se desaparece/omite cuando los 4 valores menores + 4 valores mayores y todo esto menos el valor de array[2][3] y da igual a array[0][0]
 * [números máximos (4 números)]
 * ]
 */
const calculateRowPagination = (object : IStatePagination) => {

  const { amount, page, pages_total, slice_page } = object;
  const array_previous = [] as number[]; 
  const array_next = [] as number[];
  const icon = ["..."] as string[];
  const numbers_pagination = [];

  const breakicon : boolean = page - 8 === 1 ? false : true;

  if (breakicon) {
    
  }
  for (let index = (slice_page - 1) * 4; index < ((slice_page - 1) * 4 ) + 4; index++) {
    
  }

  statePagination.value.numbers = [
    array_previous, 
   icon,
   array_next 
  ]

}

const findById = async () => {
  clientsState.value = []
  const Form = new FormData()
  if (idCustomer.value) {
    Form.append("customer_id", idCustomer.value)
    
  }
  let data = await fetch(`${import.meta.env.VITE_BACKEND_URL}\\clients`, {
    method: 'POST',
    body: Form
  })
  
  let cliente: Clientes = await data.json()
  
  clientsState.value.push(cliente)
}

const changeAmountPagination = (event: Event) => {
  getDataTypeTable(event, 'page')
}

const clearClientes = () => {
  console.log("cccc");
  
}

onBeforeMount(async () => {
    let data = await getData<Clientes[]>('')
    if (data) {
      clientsState.value = data
    }
})

</script>