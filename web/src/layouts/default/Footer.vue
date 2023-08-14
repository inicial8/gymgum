<template>
  <v-footer
    app
    height="72"
  >
    {{ date }}@copy
  </v-footer>
</template>

<script lang="ts" setup>
import {ref, onMounted, Ref} from "vue";

let date: Ref<string> = ref('')

onMounted(() => {
  let nowDate = new Date()
  date.value = nowDate.getFullYear().toString()
  })
</script>
