import { login } from '@/services/auth.service'
import { defineStore } from 'pinia'

export const useAuth = defineStore({
  id: 'user',
  actions: {
    async login(email: string, password: string) {
      await login({email, password})
    }
  }
})