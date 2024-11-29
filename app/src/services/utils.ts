import { useRouter, useRoute } from 'vue-router'
const BASE_URL = 'http://ec2-54-90-88-249.compute-1.amazonaws.com:8080'
// const BASE_URL = 'http://localhost:8080'

interface Options {
  method?: string
  data?: any
  headers?: any
}
type ApiCallOptions = Options | undefined

export async function apiCall(
  path: string,
  { method = 'GET', data, headers }: ApiCallOptions = {}
) {
  const sessionToken = sessionStorage.getItem('token')
  const url = BASE_URL + path
  const response = await fetch(url, {
    method,
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${sessionToken}`,
      ...headers
    },
    body: JSON.stringify(data)
  })

  if (!response.ok) {
    if (response.status === 401) {
      const router = useRouter()
      const route = useRoute()

      if (route?.name !== 'login' && route?.name !== 'register') {
        router?.push('/login')
        return
      }
    }
  }
  if (method === 'DELETE') {
    return response.ok
  }

  const datajson = await response.json()
  return datajson
}
