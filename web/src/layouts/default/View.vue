<template>
  <v-main>
    <router-view v-slot="{ Component, route }">
      <transition name="slide-in-down" mode="out-in">
        <component :is="Component" :key="route.path" />
      </transition>
    </router-view>
  </v-main>
</template>

<script lang="ts" setup>
  //
</script>

<style>
.slide-in-down-enter-active,
.slide-in-down-leave-active {
  transition: transform .3s cubic-bezier(0.25, 0.8, 0.5, 1), opacity .4s;
  opacity: 1;
}

.slide-in-down-enter-from,
.slide-in-down-leave-to {
  opacity: 0;
}

.slide-in-down-enter-from {
  transform: translateX(100px);
}

.slide-in-down-leave-to {
  transform: translateX(-100px);
}
</style>
