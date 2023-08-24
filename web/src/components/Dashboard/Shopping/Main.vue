<script setup lang="ts">
import ShoppingRevenue from "@/components/Dashboard/Shopping/Revenue.vue";
import ShoppingScheduled from "@/components/Dashboard/Shopping/Scheduled.vue";
import ShoppingPaid from "@/components/Dashboard/Shopping/Paid.vue";
import ShoppingOverdue from "@/components/Dashboard/Shopping/Overdue.vue";
</script>

<template>
  <v-card
      class="mx-auto"
      max-width="1000"
      elevation="0"
  >
    <v-container fluid>
      <v-row
          no-gutters
          dense
      >
        <v-col>
          <v-sheet class="pa-2 ma-2">
            <ShoppingRevenue/>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-2 ma-2">
            <ShoppingOverdue/>
          </v-sheet>
        </v-col>

      </v-row>
    </v-container>
  </v-card>
  <v-card
      class="mx-auto"
      max-width="1000"
      elevation="0"
  >
    <v-container fluid>
      <v-row
          no-gutters
          dense
      >
        <v-col>
          <v-sheet class="pa-2 ma-2">
            <ShoppingPaid/>
          </v-sheet>
        </v-col>
        <v-col>
          <v-sheet class="pa-2 ma-2">
            <ShoppingScheduled/>
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>
