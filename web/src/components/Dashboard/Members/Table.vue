<script setup lang="ts">
import {inject, onMounted, ref, Ref, toRef} from "vue";
import {VDataTable} from "vuetify/labs/VDataTable";
import {useMembersStore} from "@/stores/members";
import {IMember} from "@/models/member.model";
import {storeToRefs} from "pinia";
import MembersEditForm from "@/components/Dashboard/Members/EditForm.vue";

const memberStore = useMembersStore()
const {members}: any = storeToRefs(memberStore)

let selected: Ref<IMember[]> = ref([])
let dialog: Ref<boolean> = ref(false)
let headers: Ref<object[]> = ref([
  {title: 'Firstname', align: 'start', key: 'firstname'},
  {title: 'Middlename', align: 'end', key: 'middlename'},
  {title: 'Lastname', align: 'end', key: 'lastname'},
  {title: 'Email', align: 'end', key: 'email'},
  {title: 'Enter code', align: 'end', key: 'code'},
  {title: 'Age', align: 'end', key: 'age'},
  {title: 'Interest', align: 'end', key: 'interest'},
])

onMounted(() => {
  memberStore.getMembers()
})
const search = inject('search')

function deleteMember(id: number): void {
  memberStore.deleteMember(id)
}

function editMember(id: number): void {
  memberStore.getMember(id)
  showDialog()
}

function showDialog(): void {
  dialog.value = true
}

function closeDialog(): void {
  dialog.value = false
}
</script>

<template>
  <v-data-table
      v-model="selected"
      class="elevation-0"
      :headers="headers"
      :items="members"
      :search="search"
      items-per-page="10"
      item-value="id"
      return-object
      show-select
      select-strategy="page"
  ></v-data-table>
  <v-card class="mt-2 pa-2" elevation="0">
    <v-card-title v-if="selected.length">selected members: ({{ selected.length }})</v-card-title>
    <v-card-text>
      <v-table v-if="selected.length">
        <thead>
        <tr>
          <th class="text-left">

          </th>
          <th class="text-left">
            Name
          </th>
          <th class="text-left">
            interest
          </th>
          <th class="text-center">
            Actions
          </th>
        </tr>
        </thead>
        <tbody>
        <tr
            v-for="item in selected"
            :key="item.id"
        >
          <td>
            <v-avatar
                color="grey"
                size="40"
                rounded="2"
            >
              <v-img
                  alt="Member"
                  src="https://cdn.vuetifyjs.com/images/lists/1.jpg"
              ></v-img>
            </v-avatar>
          </td>
          <td>{{ item.firstname }} {{ item.lastname }}</td>
          <td>{{ item.interest }}</td>
          <td class="text-center">
            <v-btn icon="mdi-account-edit-outline" variant="plain" @click="editMember(item.id)"/>
            <v-btn icon="mdi-delete-circle-outline" variant="plain" @click.stop="deleteMember(item.id)"/>
          </td>
        </tr>
        </tbody>
      </v-table>
    </v-card-text>
  </v-card>
  <MembersEditForm :dialog="dialog" @close-dialog="closeDialog"/>
</template>

<style scoped>

</style>
