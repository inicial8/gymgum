import {defineStore} from "pinia";

interface IMember {
  name: string,
  middlename: string,
  lastname: string,
  email: string,
  enter_code: string,
  age: number,
  interest: string,
}

export const useMembersStore = defineStore('members', {
  state: () => ({
    member: {} as IMember
  }),
  actions: {
    async getMember(): Promise<void> {
      fetch(`http://localhost:8000/member/1`, {
        method: 'GET'
      })
        .then(response => response.json())
        .then(response => {
          this.member = {...response}
        })
    },
  }
})
