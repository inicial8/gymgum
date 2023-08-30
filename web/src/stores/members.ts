import {defineStore} from "pinia";
import {IMember} from "@/models/member.model";

export const useMembersStore = defineStore('members', {
    state: () => ({
        member: {} as IMember,
        members: <IMember[]>[]
    }),
    actions: {
        async getMember(id: number): Promise<void> {
            fetch(`http://localhost:8000/member/${id}`, {
                method: 'GET'
            })
                .then(response => response.json())
                .then(response => {
                    this.member = {...response}
                })
        },
        async getMembers(): Promise<void> {
            const response = await fetch('http://localhost:8000/members')
            this.members = await response.json()
        },
        async deleteMember(id: number): Promise<void> {
            fetch(`http://localhost:8000/member/delete/${id}`, {
                method: 'GET'
            }).then(() => {
                this.getMembers()
            })
        },
    }
})
