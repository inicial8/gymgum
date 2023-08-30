<script lang="ts" setup>
import {storeToRefs} from "pinia";
import {useMainStore} from "@/stores/main";

const {actionBar}: any = storeToRefs(useMainStore())

</script>

<template>
  <v-navigation-drawer location="right" v-model="actionBar">
    <v-list>
      <v-list-item><strong>Actions</strong></v-list-item>
      <v-list-item
        title="Statistics"
        link
      >
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>
