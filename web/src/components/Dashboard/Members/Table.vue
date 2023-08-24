<script setup lang="ts">
import {onMounted, ref, Ref} from "vue";
import {VDataTable} from "vuetify/labs/VDataTable";

let selected: Ref<object[]> = ref([])
let members: Ref<object[]> = ref([])
let headers: Ref<object[]> = ref([
  {
    title: 'Firstname',
    align: 'start',
    key: 'firstname',
  },
  {title: 'Middlename', align: 'end', key: 'middlename'},
  {title: 'Lastname', align: 'end', key: 'lastname'},
  {title: 'Email', align: 'end', key: 'email'},
  {title: 'Enter code', align: 'end', key: 'password'},
  {title: 'Age', align: 'end', key: 'age'},
  {title: 'Interest', align: 'end', key: 'interest'},
])

async function getMembers (): Promise<void> {
  const response = await fetch("http://127.0.0.1:8000/members")
  members.value = await response.json()
}

onMounted(() => {
  getMembers()
})
</script>

<template>
  <v-data-table
    v-model="selected"
    class="elevation-0"
    :headers="headers"
    :items="members"
    items-per-page="5"
    item-value="name"
    return-object
    show-select
  ></v-data-table>
  <v-card class="mt-2 pa-2" elevation="0">

    <v-card-title>selected members: ({{ selected.length }})</v-card-title>
    <v-card-text>
      <v-list lines="one">
        <v-list-item
          v-for="item in selected"
          :key="item.id"
          :title="item.firstname"
          :subtitle="item.lastname"
          prepend-avatar="https://cdn.vuetifyjs.com/images/lists/1.jpg"
        ></v-list-item>
      </v-list>
    </v-card-text>
  </v-card>
</template>

<style scoped>

</style>
