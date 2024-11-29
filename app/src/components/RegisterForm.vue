<script setup lang="ts">
import { reactive } from 'vue'
import { useUsers } from '@/services/auth'
import type { ICreateUser } from '@/types'

const { registerUser } = useUsers()

const newUser = reactive<ICreateUser>({
    username: '',
    firstname: '',
    lastname: '',
    email: '',
    password: ''
})

function register(newUser: ICreateUser) {
    registerUser(newUser)
    clearInputs()
}

function clearInputs() {
    newUser.username = ''
    newUser.firstname = ''
    newUser.lastname = ''
    newUser.email = ''
    newUser.password = ''
}

</script>

<template>
    <main class="w-full h-full flex flex-col justify-center items-center">
        <section class="flex w-full justify-center">
            <aside class="hidden sm:block sm:bg-primary w-4/12 md:w-4/12 lg:w-4/12 xl:w-3/12">
            </aside>
            <form @submit.prevent="register(newUser)"
                class="bg-softprimary px-8 my-8 py-20 flex flex-col items-center justify-center sm:w-6/12 sm:py-6 md:w-6/12 lg:w-4/12 xl:4/12">
                <header class="text-center font-roboto pb-4">
                    <h1 id="login-title" class="font-bold text-2xl">Register</h1>
                </header>
                <div class="[&>div]:input-group-register [&>div]:flex [&>div]:items-center xl:w-4/5">
                    <div>
                        <label for="username">Username</label>
                        <input v-model="newUser.username" type="text" class="form-control" id="username"
                            placeholder="Enter your username" autocomplete="username" required />
                    </div>
                    <div>
                        <label for="email">Email</label>
                        <input v-model="newUser.email" type="email" class="form-control" id="email"
                            placeholder="Enter your email" autocomplete="email" required />
                    </div>
                    <div>
                        <label for="firstName">First Name</label>
                        <input v-model="newUser.firstname" type="text" class="form-control" id="firstName"
                            placeholder="Enter your first name" autocomplete="given-name" required />
                    </div>
                    <div>
                        <label for="lastName">Last Name</label>
                        <input v-model="newUser.lastname" type="text" class="form-control" id="lastName"
                            placeholder="Enter your last name" autocomplete="family-name" required />
                    </div>
                    <div>
                        <label for="password">Password</label>
                        <input v-model="newUser.password" type="password" class="form-control" id="password"
                            placeholder="Enter your password" autocomplete="new-password" required />
                    </div>
                    <div>
                        <button type="submit" class="btn btn-primary mt-4 mx-auto">Register</button>
                    </div>
                    <p class="pt-4 text-center text-sm">
                        Already have an account?
                        <router-link to="/login" class="text-complementary hover:underline">Login</router-link>
                    </p>
                </div>

            </form>
        </section>
    </main>
</template>