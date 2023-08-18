<script setup>
import {ref} from "vue";
import {useMainStore} from "@/stores/main";

const {login} = useMainStore()

const username = ref('')
const password = ref('')

function auth() {
  login(username.value, password.value)
}
</script>

<template>
  <div class="d-flex align-center justify-center" style="height: 110vh">
    <v-sheet width="400" class="mx-auto">
      <div class="d-flex align-center justify-center">
        <img src="@/assets/gg-logo-96.png" class="logo" alt="Hexa|Gym"/>
      </div>
    </v-sheet>
    <v-sheet width="600" class="mx-auto pa-12" style="border-left: 1px solid #efefef">
      <span class="d-flex align-center justify-center">Welcome to GymGum expert!</span>
      <v-form fast-fail @submit.prevent="auth">
        <v-text-field variant="underlined" v-model="username" label="Username"></v-text-field>

        <v-text-field variant="underlined" v-model="password" label="Password" type="password"></v-text-field>
        <a href="#" class="font-weight-regular">Forgot Password?</a>

        <v-btn type="submit" variant="plain" color="primary" block class="mt-2">Sign in</v-btn>

      </v-form>
    </v-sheet>
  </div>
</template>

<style scoped>
.v-theme--dark .logo {
  filter: invert(100%);
  color: #fff;
}

.v-theme--light .logo {
  filter: invert(10%);
  color: #000;
}
</style>
a
