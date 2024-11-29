import { apiCall } from '@/services/utils'
import type { ICreateUser, ILoginUser } from '@/types'

export function useUsers() {
  async function registerUser(user: ICreateUser) {
    await apiCall('/auth/register', {
      method: 'POST',
      data: user
    })
  }

  async function loginUser(user: ILoginUser) {
    const data = await apiCall('/auth/login', {
      method: 'POST',
      data: user
    })
    return data
  }

  return { registerUser, loginUser }
}
