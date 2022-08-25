<script setup lang="ts">
import { ref } from 'vue';
import { useAuth } from '@/stores/auth'
const auth = useAuth()

let email = ref("")
let password = ref("")

const login = async () => {
    try {
        await auth.login(email.value, password.value)
    } catch(e) {
        console.error(e)
    }
}
</script>

<template>
  <form>
    <div>
      <label for="email">Email</label>
      <input type="email" name="email" id="email" v-model="email" />
    </div>
    <div>
      <label for="password">Password</label>
      <input type="password" name="password" id="password" v-model="password" />
    </div>
    <button type="button" @click="login">Accedi</button>
  </form>
</template>
