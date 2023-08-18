<template>
  <v-card variant="flat" class="ma-4">
  <v-card-title>
    <v-text-field
      v-model="search"
      append-icon="mdi-magnify"
      label="Search"
      single-line
      hide-details
    ></v-text-field>
  </v-card-title>
  <v-data-table-server
    v-model:page="page"
    :items-per-page="itemsPerPage"
    :headers="headers"
    :items-length="totalItems"
    :items="serverItems"
    :loading="loading"
    :search="search"
    class="elevation-0"
    item-value="name"
    @update:options="loadItems"
    hide-default-footer
  >
    <template v-slot:bottom>
      <div class="text-center pt-2">
        <v-pagination
          v-model="page"
          :length="pageCount"
        ></v-pagination>
      </div>
    </template>
  </v-data-table-server>
  </v-card>
</template>

<script setup lang="ts">
import {ref, computed, reactive, watch, onMounted} from 'vue'
import { VDataTableServer } from 'vuetify/labs/VDataTable'
import {useMainStore} from "@/stores/main"

const { setTitle } = useMainStore()
let page = ref(1)
let itemsPerPage = ref(5)
let totalItems = 15
let search = ref('')
let name = ref('')
let phone = ref('')
let serverItems = reactive([])
let loading = ref(false)
let clients = reactive([
  {
    name: 'test test',
    phone: '555-55-55',
  },
  {
    name: 'Tu Li',
    phone: '234-64-34',
  },
  {
    name: 'Popov',
    phone: '545-77-22',
  },
  {
    name: 'Popov',
    phone: '545-77-22',
  },
  {
    name: 'Popov',
    phone: '545-77-22',
  },
  {
    name: 'Popov',
    phone: '545-77-22',
  }
])

let headers = [
  {
    title: 'Name',
    align: 'center',
    key: 'name',
  },
  { title: 'Phone', align: 'center', key: 'phone' }
]

watch([name, phone], () => {
  search.value = String(Date.now())
})

const FakeAPI = {
  async fetch ({ page, itemsPerPage, sortBy, search }: any) {
    return new Promise(resolve => {
      setTimeout(() => {
        const start = (page - 1) * itemsPerPage
        const end = start + itemsPerPage
        const items = clients.slice().filter(item => {
          if (search.name.value && !item.name.toLowerCase().includes(search.name.value.toLowerCase())) {
            return false
          }

          return !(search.phone.value && !item.phone.includes(search.phone.value.toLowerCase()));


        })

        if (sortBy.length) {
          const sortKey = sortBy[0].key
          const sortOrder = sortBy[0].order
          items.sort((a: any, b: any) => {
            const aValue = a[sortKey]
            const bValue = b[sortKey]
            return sortOrder === 'desc' ? bValue - aValue : aValue - bValue
          })
        }

        const paginated = items.slice(start, end)

        resolve({ items: paginated, total: items.length })
      }, 500)
    })
  },
}
function loadItems ({ page, itemsPerPage, sortBy }: any) {
  loading.value = true
  FakeAPI.fetch({page, itemsPerPage, sortBy, search: { name: name, phone: phone } }).then(({items, total}: any) => {
    serverItems = items
    totalItems = total
    loading.value = false
  })
}
const pageCount = computed(() => {
  return Math.ceil(clients.length / itemsPerPage.value)
})

onMounted(() => {
  setTitle('Clients')
})
</script>
