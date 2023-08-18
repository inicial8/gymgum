<script lang="ts" setup>
import DefaultBar from "./AppBar.vue";
import DefaultView from "./View.vue";
import DefaultFooter from "./Footer.vue";
import DefaultRightBar from "./RightBar.vue";
</script>

<template>
  <v-app id="inspire">
    <default-bar/>

    <default-view/>
    <default-right-bar/>

    <default-footer/>
  </v-app>
</template>
