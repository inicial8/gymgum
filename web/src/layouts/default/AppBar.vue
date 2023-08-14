<template>
  <v-navigation-drawer
    absolute
    :rail="rail"
    permanent
  >
    <v-avatar
      color="transparent"
      size="26"
      class="ma-4"
    >
      <v-icon icon="mdi-dumbbell" size="small"></v-icon>
    </v-avatar>

    <v-divider></v-divider>

    <v-card elevation="0">
      <v-layout>
        <v-navigation-drawer
          v-model="drawer"
          :rail="rail"
          permanent
          @click="rail = false"
        >
          <v-list-item
            prepend-avatar="https://randomuser.me/api/portraits/men/85.jpg"
            title="John Leider"
            nav
          >
            <template v-slot:append>
              <v-btn
                variant="text"
                icon="mdi-chevron-left"
                @click.stop="rail = !rail"
              ></v-btn>
            </template>
          </v-list-item>

          <v-divider></v-divider>

          <v-list density="compact" nav>
            <v-list-item prepend-icon="mdi-home-city" title="Home" value="home" link to="/"></v-list-item>
            <v-list-item prepend-icon="mdi-account" title="My Account" value="account" link to="/account"></v-list-item>
            <v-list-item prepend-icon="mdi-account-group-outline" title="Clients" value="clients" to="/clients" link></v-list-item>
          </v-list>
        </v-navigation-drawer>
        <v-main style="height: 250px"></v-main>
      </v-layout>
    </v-card>
  </v-navigation-drawer>
    <v-app-bar
      class="px-4"
      color="blue-darken-2"
      flat
      height="74"
      density="compact"
    >
          <v-app-bar-title><b>GymGum</b> {{ props.version }}</v-app-bar-title>
          <template v-slot:append>
            <v-btn icon="mdi-theme-light-dark" class="ml-3" @click="toggleTheme"></v-btn>
          </template>

      <v-spacer></v-spacer>

      <v-responsive max-width="250">
        <v-text-field
          v-model="search"
          label="Search"
          placeholder="Start typing for search"
          variant="plain"
          @change="handleSearch"
        >
          <template v-slot:append-inner>
            <v-progress-circular
              v-if="loading"
              color="loading"
              indeterminate
              size="22"
            ></v-progress-circular>

            <v-icon v-if="!loading">mdi-magnify</v-icon>
          </template>
        </v-text-field>
      </v-responsive>
    </v-app-bar>
</template>

<script lang="ts" setup>
import {ref, onMounted, Ref} from "vue"
import { useTheme } from 'vuetify'

const theme = useTheme()

export interface Props {
  version?: string
}
let loading: Ref<boolean> = ref(false)
let year: Ref<number> = ref(20)
let search: Ref<string> = ref('')
let drawer = ref(true)
let rail = ref(true)

const props = withDefaults(defineProps<Props>(), {
  version: 'expert'
})

onMounted(() => {
  year.value = 2023
})

function handleSearch(): void {
  loading.value = !loading.value
}

function toggleTheme () {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
}
</script>
