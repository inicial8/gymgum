/**
 * plugins/index.ts
 *
 * Automatically included in `./src/main.ts`
 */

// Plugins
import vuetify from "./vuetify";
import router from "../router";
import {createPinia} from "pinia";
import VueApexCharts from "vue3-apexcharts";

// Types
import type {App} from "vue";

export function registerPlugins(app: App) {
  app
    .use(vuetify)
    .use(createPinia())
    .use(router)
    .use(VueApexCharts)
}
