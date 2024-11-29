<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useUsers } from '@/services/auth'

const authStore = useAuthStore()
const { loginUser } = useUsers()

const isInvalidCredentials = ref(false)
const loginInputs = reactive({
    username: '',
    password: ''
})

async function signIn() {
    const data = await loginUser(loginInputs)
    clearInputs()
    if (!data || data.token === null) {
        isInvalidCredentials.value = true
        return
    } else if (data.token) {
        authStore.setSession(data.token)
    }
}

function clearInputs() {
    loginInputs.username = ''
    loginInputs.password = ''
}


</script>

<template>
    <main class="w-full h-full flex flex-col justify-center items-center">
        <section class="flex w-full justify-center">
            <aside class="hidden sm:block sm:bg-primary w-4/12 md:w-4/12 xl:w-3/12">
            </aside>
            <form @submit.prevent="signIn"
                class="bg-softprimary px-8 my-8 py-20 flex flex-col items-center justify-center sm:w-5/12 sm:py-6  md:w-4/12 lg:w-3/12 xl:4/12">
                <header class="text-center font-roboto pb-4">
                    <h1 id="login-title" class="font-bold text-2xl">Login</h1>
                </header>
                <div class="[&>div]:input-group-login xl:w-4/5 ">
                    <div>
                        <label for="username">Username</label>
                        <input v-model="loginInputs.username" id="username" type="text"
                            placeholder="Enter your username" autocomplete="username" required />
                    </div>
                    <div>
                        <label for="password">Password</label>
                        <input v-model="loginInputs.password" id="password" type="password"
                            placeholder="Enter your password" autocomplete="current-password" required />
                    </div>
                    <div v-if="isInvalidCredentials" class="text-red-500 text-sm text-center flex justify-center">
                        Invalid Credentials
                    </div>
                    <div><button type="submit" class="btn btn-primary mt-2 mx-auto ">Login</button></div>
                    <p class="pt-4 text-center text-sm">Don't have an account?
                        <router-link to="/register" class="text-complementary hover:underline">Register</router-link>
                    </p>
                </div>
            </form>
        </section>
    </main>
</template>