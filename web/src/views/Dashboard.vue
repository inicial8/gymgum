<template>
  <Main />
</template>

<script lang="ts" setup>
  import Main from '@/components/Main.vue'
  import { useMainStore } from '@/stores/main'
  import {onMounted} from "vue"

  const { setTitle } = useMainStore()

  onMounted(() => {
    setTitle('Dashboard')
  })
</script>
