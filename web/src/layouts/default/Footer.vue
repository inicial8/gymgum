<script lang="ts" setup>
import {ref, onMounted, Ref} from "vue";

let date: Ref<string> = ref('')

onMounted(() => {
  let nowDate = new Date()
  date.value = nowDate.getFullYear().toString()
})
</script>

<template>
  <v-footer style="height: 80px;max-height: 80px;">
    <v-row justify="center" no-gutters>
      <v-col class="text-center" cols="8">
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
      </v-col>
      <v-col class="text-center pt-4" cols="6">
        {{ date }} © <strong>Hexa|Gym</strong> expert
      </v-col>
    </v-row>
  </v-footer>
</template>
