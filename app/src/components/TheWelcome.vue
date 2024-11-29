<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore();
const { isLoggedIn } = authStore;

</script>

<template>
  <main class="w-full h-full flex flex-col justify-center items-center">
    <div class="flex flex-col items-center pb-4 px-6">
      <h1 class="text-center font-roboto font-bold text-2xl mb-4"> ¡Welcome to Analysis Center!</h1>
      <h2 class="font-roboto text-center md:w-4/5 lg:w-3/5">In the Analysis Center you can explore and analyze data from
        the Study created by Barry Becker based on a 1994 census that helps understand what type of people earned more
        or less than $50,000 USD. </h2>
    </div>
    <img src="../assets/img/graphics.png" alt="" class="w-48 mt-2 mb-8">
    <button class="btn-primary mb-6">
      <RouterLink :to="isLoggedIn ? '/analysisCenter' : '/login'">Comenzar</RouterLink>
    </button>
  </main>
</template>
