<template>
  <v-navigation-drawer
    v-model="drawer"
    :rail="rail"
    @click="rail = false"
    location="left"
  >
    <v-list-item
      prepend-avatar="@/assets/logo.png"
      :title="`GymGum ${props.version}`"
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
      <v-list-item prepend-icon="mdi-home-city" title="Home" value="home" to="/"></v-list-item>
      <v-list-item prepend-icon="mdi-account" title="My Account" value="account" to="/account"></v-list-item>
      <v-list-item prepend-icon="mdi-account-group-outline" title="Clients" value="clients" to="/clients"></v-list-item>
    </v-list>
  </v-navigation-drawer>
  <v-app-bar
    color="blue-darken-2"
    flat
    height="74"
    density="compact"
  >
    <v-btn
      icon="mdi-menu"
      @click.stop="drawer = !drawer"
      v-show="!drawer"
    />
    <v-app-bar-title>{{ title }}</v-app-bar-title>
    <template v-slot:append>
      <v-tooltip text="Switch theme" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn icon="mdi-theme-light-dark" class="ml-3" @click="toggleTheme" v-bind="props"></v-btn>
        </template>
      </v-tooltip>

      <v-tooltip text="Login" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn icon="mdi-login" class="ml-3" v-bind="props"></v-btn>
        </template>
      </v-tooltip>


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
import {useTheme} from 'vuetify'
import {useMainStore} from '@/stores/main'
import {storeToRefs} from 'pinia'

const theme = useTheme()

export interface Props {
  version?: string
}

let loading: Ref<boolean> = ref(false)
let year: Ref<number> = ref(20)
let search: Ref<string> = ref('')
let drawer = ref(true)
let rail = ref(false)
const {title} = storeToRefs(useMainStore())

const props = withDefaults(defineProps<Props>(), {
  version: 'expert'
})

onMounted(() => {
  year.value = 2023
})

function handleSearch(): void {
  loading.value = !loading.value
}

function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
}
</script>
