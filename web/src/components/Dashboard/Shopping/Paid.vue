<script setup lang="ts">
const options: object = {
  chart: {
    type: 'radialBar'
  },
  title: {
    text: 'Paid'
  }
}

const series: number[] = [30]
</script>

<template>
  <apexchart :options="options" :series="series"></apexchart>
</template>

<style scoped>

</style>
