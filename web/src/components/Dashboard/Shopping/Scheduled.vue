<script setup lang="ts">
const options = {
  chart: {
    type: 'radialBar'
  },
  title: {
    text: 'Scheduled'
  },
}

const series = [30]
</script>

<template>
  <apexchart type="radialBar" :options="options" :series="series"></apexchart>
</template>

<style scoped>

</style>
