<script lang="ts" setup>
import {ref, onMounted, Ref} from "vue";
import {useTheme} from "vuetify";
import {useMainStore} from "@/stores/main";
import {storeToRefs} from "pinia";

const {user, isAuthenticated}: any = storeToRefs(useMainStore())

const theme: any = useTheme()

let loading: Ref<boolean> = ref(false)
let year: Ref<number> = ref(20)
let search: Ref<string> = ref('')
let drawer: Ref<boolean> = ref(true)
let rail: Ref<boolean> = ref(false)
const {title}: any = storeToRefs(useMainStore())

onMounted(() => {
  year.value = 2023
})

function handleSearch(): void {
  loading.value = !loading.value
}

function toggleTheme(): void {
  localStorage.setItem('theme', theme.global.current.value.dark ? 'light' : 'dark')
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
}

function logout(): void {
  localStorage.removeItem('user')
  localStorage.removeItem('role')
  window.location.reload()
}
</script>

<template>
  <v-navigation-drawer
      v-model="drawer"
      :rail="rail"
      @click="rail = false"
      location="left"
  >
    <v-list-item
        prepend-avatar="@/assets/gg-logo-96.png"
        nav
        class="logo"
    >
      <template v-slot:append>
        <strong class="logo">Hexa|Gym expert</strong>
        <v-btn
            class="logo"
            variant="text"
            icon="mdi-chevron-left"
            @click.stop="rail = !rail"
        ></v-btn>
      </template>
    </v-list-item>

    <v-divider></v-divider>

    <v-list density="compact" nav>
      <v-list-item prepend-icon="mdi-monitor-dashboard" title="Dashboard" value="dashboards" to="/"></v-list-item>
      <v-list-item prepend-icon="mdi-account" title="My Account" value="account" to="/account"
                   v-if="user.role === 'admin' || user.role === 'client'"></v-list-item>
      <v-list-item prepend-icon="mdi-account-group-outline" title="Clients" value="clients" to="/clients"
                   v-if="user.role === 'admin'"></v-list-item>
    </v-list>
  </v-navigation-drawer>
  <v-app-bar
      color="blue-accent-2"
      flat
      height="56"
  >
    <v-btn
        icon="mdi-menu"
        @click.stop="drawer = !drawer"
    />
    <v-app-bar-title><strong>{{ title }}</strong></v-app-bar-title>
    <template v-slot:append>
      <v-tooltip text="Switch theme" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn icon="mdi-theme-light-dark" class="ml-1" @click="toggleTheme" v-bind="props"></v-btn>
        </template>
      </v-tooltip>
      <span>{{ user.name }}</span>

      <v-tooltip text="Login" location="bottom" v-if="!isAuthenticated">
        <template v-slot:activator="{ props }">
          <v-btn icon="mdi-login" class="ml-1" v-bind="props" to="login"></v-btn>
        </template>
      </v-tooltip>

      <v-tooltip text="Logout" location="bottom" v-else>
        <template v-slot:activator="{ props }">
          <v-btn icon="mdi-logout" class="ml-1" v-bind="props" @click="logout"></v-btn>
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
