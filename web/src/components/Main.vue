<script lang="ts" setup>
import {ref} from "vue";
import MembersTable from "@/components/Dashboard/Members/Table.vue";
import MembersDialog from "@/components/Dashboard/Members/DialogForm.vue";
import ShoppingMain from "@/components/Dashboard/Shopping/Main.vue";
import {useMainStore} from "@/stores/main";
import {storeToRefs} from "pinia";

const {setActionBar} = useMainStore()
const {actionBar} = storeToRefs(useMainStore())

const currentItem = ref('tab-Shopping')
const items = ref(['Members', 'Shopping', 'Videos', 'Images'])
const more = ref(['News', 'Maps', 'Books', 'Flights', 'Apps'])
let dialog = ref(false)

function showActionsBar() {
  let bar = !actionBar.value
  setActionBar(bar)
}

function showDialog() {
  dialog.value = true
}

function closeDialog() {
  dialog.value = false
}
</script>

<template>
  <v-card variant="flat" flat class="ma-4">
    <v-toolbar
      color="gray"
    >

      <v-toolbar-title>Your dashboards</v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn icon @click="showDialog()" v-if="currentItem === 'tab-Members'">
        <v-icon>mdi-account-multiple-plus-outline</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-btn icon @click.stop="showActionsBar">
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn>

      <template v-slot:extension>
        <v-tabs
          v-model="currentItem"
          fixed-tabs
        >
          <v-tab
            v-for="item in items"
            :key="item"
            :value="'tab-' + item"
          >
            {{ item }}
          </v-tab>

          <v-menu
            v-if="more.length"
          >
            <template v-slot:activator="{ props }">
              <v-btn
                variant="plain"
                rounded="0"
                class="align-self-center me-4"
                height="100%"
                v-bind="props"
              >
                more
                <v-icon end>
                  mdi-menu-down
                </v-icon>
              </v-btn>
            </template>

            <v-list class="bg-grey-lighten-3">
              <v-list-item
                v-for="item in more"
                :key="item"
              >
                {{ item }}
              </v-list-item>
            </v-list>
          </v-menu>
        </v-tabs>
      </template>
    </v-toolbar>

    <v-window v-model="currentItem">
      <v-window-item
        v-for="item in items.concat(more)"
        :key="item"
        :value="'tab-' + item"
      >
        <v-card flat v-if="currentItem === 'tab-Members'">
          <v-card-text>
            <div class="pa-6">
              <MembersTable/>
              <MembersDialog :dialog="dialog" @close-dialog="closeDialog"/>
            </div>
          </v-card-text>
        </v-card>
        <v-card flat v-if="currentItem === 'tab-Shopping'">
          <ShoppingMain />
        </v-card>
      </v-window-item>
    </v-window>
  </v-card>
</template>
