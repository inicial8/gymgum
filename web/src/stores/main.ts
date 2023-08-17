import { defineStore } from 'pinia'

export const useMainStore = defineStore('main', {
  state: () => ({
    title: '',
    actionBar: true,
    showDialog: false,
  }),
  actions:{
    async setTitle(value: string) {
      this.title = value
    },
    async setActionBar(value: boolean) {
      this.actionBar = value
    },
    async setShowDialog(value: boolean) {
      this.showDialog = value
    },
  }
})
