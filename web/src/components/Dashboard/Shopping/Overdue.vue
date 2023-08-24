<script setup lang="ts">

const options: object = {
  chart: {
    type: 'radialBar'
  },
  plotOptions: {
    radialBar: {
      size: undefined,
      inverseOrder: true,
      hollow: {
        margin: 5,
        size: '48%',
        background: 'transparent',

      },
      track: {
        show: false,
      },
      startAngle: -180,
      endAngle: 180
    },
  },
  stroke: {
    lineCap: 'round'
  },
  labels: ['June', 'May', 'April'],
  legend: {
    show: true,
    floating: true,
    position: 'right',
    offsetX: 70,
    offsetY: 240
  },
  title: {
    text: 'Overdue'
  }
}

const series: number[] = [71, 63, 77]
</script>

<template>
  <apexchart :options="options" :series="series"></apexchart>
</template>

<style scoped>

</style>
