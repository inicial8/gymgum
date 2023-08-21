<script lang="ts" setup>
import {useStorage} from "@vueuse/core";
import {watch} from "vue";

const theDefault = {
  role: ''
}

const state = useStorage('role', theDefault.role)

watch(state, () => {
  window.location.assign('/')
})
</script>

<template>
  <router-view/>
</template>

<style>

</style>
