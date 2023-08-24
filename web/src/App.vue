<script lang="ts" setup>
import {useStorage} from "@vueuse/core";
import {watch} from "vue";
import {useTheme} from "vuetify";
import {RemovableRef} from "@vueuse/shared";

const theme: any = useTheme()

const theDefault: object = {
  role: ''
}

const state: RemovableRef<string> = useStorage('role', theDefault.role)

watch(state, () => {
  window.location.assign('/')
})

theme.global.name.value = localStorage.getItem('theme') || 'dark'
</script>

<template>
  <router-view/>
</template>

<style>

</style>
