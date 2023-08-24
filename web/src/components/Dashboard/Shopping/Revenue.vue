<script setup lang="ts">
import {reactive, computed} from "vue";

const options = computed(() => {
  return {
    chart: {
      type: 'bar',
      toolbar: {
        show: false
      },
    },
    xaxis: {
      categories: [1, 2, 3, 4, 5, 6, 7, 8]
    },
    title: {
      text: 'Revenue'
    },
    tooltip: {
      custom: function ({series, seriesIndex, dataPointIndex, w}: any) {
        return '<div>' +
            '<span>' + series[seriesIndex][dataPointIndex] + '$</span>' +
            '</div>'
      }
    }
  }
})

const series = reactive([{
  name: 'Revenue',
  data: [3000, 4000, 4500, 5000, 4900, 6000, 7000, 9100]
}])
</script>

<template>
  <apexchart :options="options" :series="series"></apexchart>
</template>

<style scoped>

</style>
