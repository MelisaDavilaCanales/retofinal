<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';

const authStore = useAuthStore();

function logout() {
    authStore.clearSession();
}
</script>

<template>
    <main class=" w-full h-full flex flex-col justify-center items-center">
        <form class=" bg-softprimary rounded-md py-8 px-16 mb-8 m-4 flex flex-col items-center md:px-28 md:mb-16"
            @submit.prevent="logout">
            <img src="../assets/img/logout.png" alt="img_logout" class="w-32 pb-8">
            <p>Are you sure you want to logout?</p>
            <button class="btn-primary mt-6" type="submit">Logout</button>
        </form>
    </main>
</template>