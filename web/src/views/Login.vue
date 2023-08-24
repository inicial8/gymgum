<script lang="ts" setup>
import {ref, Ref} from "vue";
import {useMainStore} from "@/stores/main";
import {useField, useForm} from "vee-validate";
import router from "@/router";

const {errorMessage, meta, value}: any = useField('username')

const {login}: any = useMainStore()

const user: any = localStorage.getItem('user')

let visible: Ref<boolean> = ref(false)

const {handleSubmit} = useForm({
  validationSchema: {
    password(value: string | any[]): string | boolean {
      if (value?.length >= 3) return true
      return "Password needs to be at least 3 characters."
    },
    username(value: string): string | boolean {
      if (value?.length >= 3) return true
      return "Username needs to be at least 3 characters."
    },
  },
})

const username: any = useField("username")
const password: any = useField("password")

const submit: any = handleSubmit(() => {
  if (user) {
    router.push('/');
  } else {
    login(username.value.value)
    localStorage.setItem('user', username.value.value);
    router.push({path: '/'});
  }
})
</script>

<template>
  <div class="d-flex align-center justify-center" style="height: 110vh">
    <v-sheet width="400" class="mx-auto" style="background-color: transparent">
      <div class="d-flex align-center justify-center">
        <img src="@/assets/gg-logo-96.png" class="logo" color="transparent" alt="Hexa|Gym"/>
      </div>
    </v-sheet>
    <v-sheet width="600" class="mx-auto pa-12" style="border-left: 1px solid #efefef">
      <h2 class="d-flex align-center justify-center pb-4">Welcome to Hexa|Gym expert!</h2>
      <v-form fast-fail @submit="submit">
        <v-text-field
            class="pa-1"
            variant="outlined"
            v-model="username.value.value"
            label="Username"
            name="username"
            prepend-inner-icon="mdi-account-outline"
            :error-messages="username.errorMessage.value"
        ></v-text-field>

        <v-text-field
            class="pa-1"
            variant="outlined"
            v-model="password.value.value"
            label="Password"
            :type="visible ? 'text' : 'password'"
            prepend-inner-icon="mdi-lock-outline"
            :append-inner-icon="visible ? 'mdi-eye-off' : 'mdi-eye'"
            @click:append-inner="visible = !visible"
            :error-messages="password.errorMessage.value"
            name="password"
        ></v-text-field>
        <a href="#" class="font-weight-regular">Forgot Password?</a>

        <span v-if="errorMessage && meta.touched">
          {{ errorMessage }} {{ value }}
        </span>

        <v-btn type="submit" variant="plain" color="primary" block class="mt-2"
               :disabled="password === '' && username === ''">Sign in
        </v-btn>

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
