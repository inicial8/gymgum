<script lang="ts" setup>
import {storeToRefs} from "pinia";
import {useMainStore} from "@/stores/main";

const {actionBar}: any = storeToRefs(useMainStore())

const items = [
  {type: 'subheader', title: 'Today'},
  {
    prependAvatar: 'https://cdn.vuetifyjs.com/images/lists/1.jpg',
    title: 'Brunch this weekend?',
    subtitle: `<span class="text-primary">Ali Connors</span> &mdash; I'll be in your neighborhood doing errands this weekend. Do you want to hang out?`,
  },
  {type: 'divider', inset: true},
  {
    prependAvatar: 'https://cdn.vuetifyjs.com/images/lists/2.jpg',
    title: 'Summer BBQ',
    subtitle: `<span class="text-primary">to Alex, Scott, Jennifer</span> &mdash; Wish I could come, but I'm out of town this weekend.`,
  },
  {type: 'divider', inset: true},
  {
    prependAvatar: 'https://cdn.vuetifyjs.com/images/lists/3.jpg',
    title: 'Oui oui',
    subtitle: '<span class="text-primary">Sandra Adams</span> &mdash; Do you have Paris recommendations? Have you ever been?',
  },
  {type: 'divider', inset: true},
  {
    prependAvatar: 'https://cdn.vuetifyjs.com/images/lists/4.jpg',
    title: 'Birthday gift',
    subtitle: '<span class="text-primary">Trevor Hansen</span> &mdash; Have any ideas about what we should get Heidi for her birthday?',
  },
  {type: 'divider', inset: true},
  {
    prependAvatar: 'https://cdn.vuetifyjs.com/images/lists/5.jpg',
    title: 'Recipe to try',
    subtitle: '<span class="text-primary">Britta Holt</span> &mdash; We should eat this: Grate, Squash, Corn, and tomatillo Tacos.',
  },
  {type: 'divider', inset: true},
  {
    prependAvatar: 'https://cdn.vuetifyjs.com/images/lists/3.jpg',
    title: 'Oui oui',
    subtitle: '<span class="text-primary">Sandra Adams</span> &mdash; Do you have Paris recommendations? Have you ever been?',
  },
  {type: 'divider', inset: true},
  {
    prependAvatar: 'https://cdn.vuetifyjs.com/images/lists/4.jpg',
    title: 'Birthday gift',
    subtitle: '<span class="text-primary">Trevor Hansen</span> &mdash; Have any ideas about what we should get Heidi for her birthday?',
  },
  {type: 'divider', inset: true},
  {
    prependAvatar: 'https://cdn.vuetifyjs.com/images/lists/2.jpg',
    title: 'Summer BBQ',
    subtitle: `<span class="text-primary">to Alex, Scott, Jennifer</span> &mdash; Wish I could come, but I'm out of town this weekend.`,
  },
  {type: 'divider', inset: true},
  {
    prependAvatar: 'https://cdn.vuetifyjs.com/images/lists/3.jpg',
    title: 'Oui oui',
    subtitle: '<span class="text-primary">Sandra Adams</span> &mdash; Do you have Paris recommendations? Have you ever been?',
  },
]

</script>

<template>
  <v-navigation-drawer location="right" v-model="actionBar">
    <v-toolbar>
      <v-toolbar-title>History</v-toolbar-title>
    </v-toolbar>

    <v-list
      :items="items"
      item-props
      lines="three"
    >
      <template v-slot:subtitle="{ subtitle }">
        <div v-html="subtitle"></div>
      </template>
    </v-list>
  </v-navigation-drawer>
</template>
