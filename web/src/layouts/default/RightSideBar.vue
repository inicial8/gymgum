<template>
  <v-container class="fill-height">
    <v-responsive class="d-flex align-center text-center fill-height">
      <v-navigation-drawer location="right" v-model="drawer">
<!--        <v-btn icon="mdi-close" variant="text" class="float-left" @click.stop="drawer = !drawer"></v-btn>-->
        <v-list>
          <v-list-item><strong>Actions</strong></v-list-item>
          <v-list-item
            v-for="n in 6"
            :key="n"
            :title="`Action ${ n }`"
            link
          >
          </v-list-item>
        </v-list>
      </v-navigation-drawer>
      <div>Action results</div>
    </v-responsive>
  </v-container>
</template>

<script lang="ts" setup>
import {ref} from 'vue'
import {storeToRefs} from "pinia"
import {useMainStore} from "@/stores/main"

const {actionBar} = storeToRefs(useMainStore())
const drawer = ref(true)
</script>
