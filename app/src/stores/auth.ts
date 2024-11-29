import type { JWTPayload } from '@/types'
import { usePersonStore } from '@/stores/persons'
import { jwtDecode } from 'jwt-decode'
import { defineStore } from 'pinia'
import { computed, ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

export const useAuthStore = defineStore('auth', () => {
  const router = useRouter()
  const session = ref('')

  const state = reactive({
    authUserId: ''
  })

  const isLoggedIn = computed(() => !!session.value)

  async function init() {
    const tokenStr = sessionStorage.getItem('token')
    if (tokenStr) {
      setSession(tokenStr)
    }
  }

  function setSession(tokenStr: string) {
    const payload = jwtDecode(tokenStr) as JWTPayload
    const now = new Date()
    const diff = payload.MapClaims.eat * 1000 - now.getTime()

    state.authUserId = payload.MapClaims.id
    const personStore = usePersonStore()
    personStore.setQueryString(payload.searchFilters)

    if (diff <= 0) {
      clearSession()
      return
    }

    sessionStorage.setItem('token', tokenStr)
    session.value = payload.session
    router.push('/analysisCenter')

    setTimeout(() => {
      clearSession()
    }, diff)
  }

  function clearSession() {
    sessionStorage.removeItem('token')
    session.value = ''
    router.push('/')
  }

  return { state, isLoggedIn, init, clearSession, setSession }
})
