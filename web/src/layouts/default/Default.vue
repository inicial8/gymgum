<template>
  <v-app id="inspire">
    <default-bar />

    <default-view />
    <RightSideBar />

    <default-footer />
  </v-app>
</template>

<script lang="ts" setup>
  import DefaultBar from './AppBar.vue'
  import DefaultView from './View.vue'
  import DefaultFooter from './Footer.vue'
  import RightSideBar from "./RightSideBar.vue"
</script>
