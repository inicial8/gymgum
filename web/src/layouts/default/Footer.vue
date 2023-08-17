<template>
  <v-footer height="200">
    <v-row justify="center" no-gutters>
      <v-btn
        variant="text"
      >
        Home
      </v-btn>
      <v-btn
        variant="text"
      >
        About Us
      </v-btn>
      <v-btn
        variant="text"
      >
        Team
      </v-btn>
      <v-btn
        variant="text"
      >
        Services
      </v-btn>
      <v-btn
        variant="text"
      >
        Blog
      </v-btn>
      <v-btn
        variant="text"
      >
        Contact Us
      </v-btn>
      <v-col class="text-center mt-2" cols="12">
        {{ date }} © <strong>Hexa|Gym</strong> expert
      </v-col>
    </v-row>
  </v-footer>
</template>

<script lang="ts" setup>
import {ref, onMounted, Ref} from "vue";

let date: Ref<string> = ref('')

onMounted(() => {
  let nowDate = new Date()
  date.value = nowDate.getFullYear().toString()
  })
</script>
