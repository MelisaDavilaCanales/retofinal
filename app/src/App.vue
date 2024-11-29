<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useAuthStore } from './stores/auth'
import { onBeforeMount } from 'vue'

const authStore = useAuthStore()

onBeforeMount(authStore.init)

</script>

<template>
  <div class="w-full h-screen flex flex-col">
    <header class="bg-primary p-2">
      <nav class="flex justify-center text-white font-semi-bold">
        <RouterLink class="px-4 cursor-pointer hover:text-green-400" to="/">Home</RouterLink>
        <RouterLink class="px-4 cursor-pointer hover:text-green-400" v-if="authStore.isLoggedIn" to="/analysisCenter">
          Analysis Center</RouterLink>
        <RouterLink class="px-4 cursor-pointer hover:text-green-400" v-if="authStore.isLoggedIn" to="/statistics">
          Statistics</RouterLink>
        <RouterLink class="px-4 cursor-pointer hover:text-green-400" v-if="!authStore.isLoggedIn" to="/login">Login
        </RouterLink>
        <RouterLink class="px-4 cursor-pointer hover:text-green-400" v-if="!authStore.isLoggedIn" to="/contactUs">
          Contact Us</RouterLink>
        <RouterLink class="px-4 cursor-pointer hover:text-green-400" v-if="authStore.isLoggedIn" to="/logout">
          Logout</RouterLink>
      </nav>
    </header>
    <RouterView class="w-full flex-grow" />
  </div>
</template>
