<script setup lang="ts">
import {toRef} from "vue";
import {useField, useForm} from "vee-validate";
import {useMembersStore} from "@/stores/members";
import {IMember} from "@/models/member.model";

const emit = defineEmits(['closeDialog'])

const memberStore = useMembersStore()

const {value}: any = useField('email')

let props: any = defineProps(['dialog', 'member'])
let dialog: any = toRef(props, 'dialog')

const {handleSubmit} = useForm({
  validationSchema: {
    firstname(value: string | any[]): string | boolean {
      if (value?.length >= 2) return true
      return "First name needs to be at least 2 characters."
    },
    code(value: string | any[]): string | boolean {
      if (value?.length >= 4) return true
      return "Code needs to be at least 4 characters."
    },
    email(value: string): string | boolean {
      if (value?.length >= 3) return true
      return "Email needs to be at least 3 characters."
    },
  },
})

const email: any = useField("email")
const code: any = useField("code")
const firstname: any = useField("firstname")
const middlename: any = useField("middlename")
const lastname: any = useField("lastname")
const age: any = useField("age")
const interest: any = useField("interest")

const submit: any = handleSubmit((values, actions) => {
  const member: IMember = {
    id: 0,
    firstname: firstname.value.value,
    middlename: middlename.value.value,
    lastname: lastname.value.value,
    email: email.value.value,
    code: code.value.value,
    age: age.value.value,
    interest: interest.value.value
  }
  fetch('http://127.0.0.1:8000/member', {
    method: 'POST',
    body: JSON.stringify(member)
  })
    .then(response => response.json())
    .then(() => memberStore.getMembers())
    .then(() => emit('closeDialog'))
  actions.resetForm()
})
</script>

<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" width="1024" persistent>
      <v-form @submit.prevent="submit()">
        <v-card>
          <v-card-title>
            <span class="text-h5">Add Member</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-row>
                <v-col
                  cols="12"
                  sm="6"
                  md="4"
                >
                  <v-text-field
                    label="Legal first name*"
                    required
                    v-model="firstname.value.value"
                    :error-messages="firstname.errorMessage.value"
                    name="firstname"
                    variant="underlined"
                    hide-details
                  ></v-text-field>
                </v-col>
                <v-col
                  cols="12"
                  sm="6"
                  md="4"
                >
                  <v-text-field
                    label="Legal middle name"
                    hint="not necessary"
                    v-model="middlename.value.value"
                    :error-messages="middlename.errorMessage.value"
                    variant="underlined"
                  ></v-text-field>
                </v-col>
                <v-col
                  cols="12"
                  sm="6"
                  md="4"
                >
                  <v-text-field
                    label="Legal last name*"
                    hint="*please do not ignore this field"
                    persistent-hint
                    required
                    v-model="lastname.value.value"
                    :error-messages="lastname.errorMessage.value"
                    variant="underlined"
                  ></v-text-field>
                </v-col>
                <v-col cols="12">
                  <v-text-field
                    label="Email*"
                    required
                    v-model="email.value.value"
                    name="email"
                    prepend-inner-icon="mdi-email"
                    :error-messages="email.errorMessage.value"
                    variant="underlined"
                  ></v-text-field>
                </v-col>
                <v-col cols="12">
                  <v-text-field
                    label="Code*"
                    type="password"
                    name="code"
                    v-model="code.value.value"
                    prepend-inner-icon="mdi-key"
                    :error-messages="code.errorMessage.value"
                    required
                    variant="underlined"
                  ></v-text-field>
                </v-col>
                <v-col
                  cols="12"
                  sm="6"
                >
                  <v-select
                    :items="['0-17', '18-29', '30-54', '54+']"
                    label="Age*"
                    required
                    v-model="age.value.value"
                    :error-messages="age.errorMessage.value"
                    variant="underlined"
                  ></v-select>
                </v-col>
                <v-col
                  cols="12"
                  sm="6"
                >
                  <v-autocomplete
                    :items="['Grappling', 'Box', 'Stretching', 'Basketball', 'Basejump']"
                    label="Interests"
                    v-model="interest.value.value"
                    :error-messages="interest.errorMessage.value"
                    variant="underlined"
                  ></v-autocomplete>
                </v-col>
              </v-row>
            </v-container>
            <small>*indicates required field</small>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="blue-darken-1"
              variant="text"
              @click="$emit('closeDialog')"
            >
              Close
            </v-btn>
            <v-btn
              color="blue-darken-1"
              variant="text"
              type="submit"
            >
              Save
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-form>
    </v-dialog>
  </v-row>
</template>

<style scoped>

</style>
