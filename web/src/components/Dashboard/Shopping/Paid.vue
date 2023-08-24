<script setup lang="ts">
const options = {
  chart: {
    type: 'radialBar'
  },
  title: {
    text: 'Paid'
  }
}

const series = [30]
</script>

<template>
  <apexchart :options="options" :series="series"></apexchart>
</template>

<style scoped>

</style>
