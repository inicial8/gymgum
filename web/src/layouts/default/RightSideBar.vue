<template>
  <v-container class="fill-height">
    <v-responsive class="d-flex align-center text-center fill-height">
      <v-navigation-drawer location="right">
        <v-list>
          <v-list-item
            v-for="n in 6"
            :key="n"
            :title="`Item ${ n }`"
            link
          >
          </v-list-item>
        </v-list>
      </v-navigation-drawer>
    </v-responsive>
  </v-container>
</template>

<script lang="ts" setup>
  //
</script>
