import {defineStore} from "pinia";
import router from "@/router";

interface IUser {
  name: string,
  role: string
}

export const useMainStore = defineStore('main', {
  state: () => ({
    title: '',
    actionBar: false,
    showDialog: false,
    isAuthenticated: (localStorage.getItem('user') !== null),
    user: {name: localStorage.getItem('user'), role: localStorage.getItem('role')} as IUser
  }),
  actions: {
    async setTitle(value: string) {
      this.title = value
    },
    async setActionBar(value: boolean) {
      this.actionBar = value
    },
    async setShowDialog(value: boolean) {
      this.showDialog = value
    },
    async login(username: string, password: string) {
      fetch('http://127.0.0.1:8000/auth', {
        method: 'POST',
        body: JSON.stringify({username: username, password: password})
      })
        .then(response => response.json())
        .then(response => {
          this.user.name = response.name
          this.user.role = response.role
          localStorage.setItem('user', response.name)
          localStorage.setItem('role', response.role)
          this.isAuthenticated = true
        })
        .then(() => router.push('/'))
    },
  }
})
