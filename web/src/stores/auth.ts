import { login } from '@/services/auth.service'
import { defineStore } from 'pinia'

export const useAuth = defineStore({
  id: 'user',
  state: () => ({
    token: ""
  }),
  actions: {
    async login(email: string, password: string) {
      const { token } = await login({email, password})
      this.token = token
      console.log(token)
    }
  }
})