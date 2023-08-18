<script lang="ts" setup>
import {storeToRefs} from "pinia";
import {useMainStore} from "@/stores/main";

const {actionBar} = storeToRefs(useMainStore())
</script>

<template>
      <v-navigation-drawer location="right" v-model="actionBar">
        <!--<v-btn icon="mdi-close" variant="text" class="float-left" @click.stop="drawer = !drawer"></v-btn>-->
        <v-list>
          <v-list-item><strong>Actions</strong></v-list-item>
          <v-list-item
            v-for="n in 6"
            :key="n"
            :title="`Action ${ n }`"
            link
          >
          </v-list-item>
        </v-list>
      </v-navigation-drawer>
</template>
