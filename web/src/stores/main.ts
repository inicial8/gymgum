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
    async setTitle(value: string): Promise<void> {
      this.title = value
    },
    async setActionBar(value: boolean): Promise<void> {
      this.actionBar = value
    },
    async setShowDialog(value: boolean): Promise<void> {
      this.showDialog = value
    },
    async login(username: string, password: string): Promise<void> {
      fetch('http://127.0.0.1:8000/auth', {
        method: 'POST',
        body: JSON.stringify({username: username, password: password})
      })
        .then(response => response.json())
        .then(response => {
          this.user.name = response.username
          this.user.role = response.role
          localStorage.setItem('user', response.username)
          localStorage.setItem('role', response.role)
          this.isAuthenticated = true
        })
        .then(() => router.push('/'))
    },
  }
})
