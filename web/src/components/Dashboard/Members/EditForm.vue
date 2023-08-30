<script setup lang="ts">
import {toRef} from "vue";
import {useField} from "vee-validate";
import {useMembersStore} from "@/stores/members";
import {IMember} from "@/models/member.model";

const emit = defineEmits(['closeDialog'])

const memberStore = useMembersStore()

const {value}: any = useField('email')

let props: any = defineProps(['dialog', 'member'])
let dialog: any = toRef(props, 'dialog')

let member: IMember = {...memberStore.member}

const submit: any = () => {
  console.log(JSON.stringify(memberStore.member))
  fetch(`http://127.0.0.1:8000/member/${memberStore.member.id}`, {
    method: 'POST',
    body: JSON.stringify(memberStore.member)
  }).then(() => memberStore.getMembers())
      .then(() => emit('closeDialog'))
}
</script>

<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" width="1024" persistent>
      <v-form @submit.prevent="submit()">
        <v-card>
          <v-card-title>
            <span class="text-h5">Edit Member</span>
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
                      v-model="memberStore.member.firstname"
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
                      v-model="memberStore.member.middlename"
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
                      v-model="memberStore.member.lastname"
                      variant="underlined"
                  ></v-text-field>
                </v-col>
                <v-col cols="12">
                  <v-text-field
                      label="Email*"
                      required
                      v-model="memberStore.member.email"
                      name="email"
                      prepend-inner-icon="mdi-email"
                      variant="underlined"
                  ></v-text-field>
                </v-col>
                <v-col cols="12">
                  <v-text-field
                      label="Code*"
                      type="password"
                      name="code"
                      v-model="memberStore.member.code"
                      prepend-inner-icon="mdi-key"
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
                      v-model="memberStore.member.age"
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
                      v-model="memberStore.member.interest"
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
