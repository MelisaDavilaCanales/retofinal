import { apiCall } from '@/services/utils'

export function usePersons() {
  async function fetchPaginatedPersons(querystring: string) {
    const data = await apiCall(`/persons/paginated?${querystring}`, {
      method: 'GET'
    })
    return data
  }

  async function fetchPersonns(querystring: string) {
    const data = await apiCall(`/persons?${querystring}`, {
      method: 'GET'
    })
    return data
  }

  return { fetchPersonns, fetchPaginatedPersons }
}