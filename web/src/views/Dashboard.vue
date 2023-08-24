<script lang="ts" setup>
import Main from "@/components/Dashboard/Members/Main.vue";
import {useMainStore} from "@/stores/main";
import {onMounted} from "vue";

const {setTitle}: any = useMainStore()

onMounted(() => {
  setTitle('Dashboard')
})
</script>

<template>
  <Main/>
</template>
