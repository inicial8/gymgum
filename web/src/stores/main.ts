import { defineStore } from 'pinia'

export const useMainStore = defineStore('main', {
  state: () => ({
    title: '',
  }),
  actions:{
    async setTitle(value: string) {
      this.title = value
    },
  }
})
