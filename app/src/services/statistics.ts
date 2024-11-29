import { apiCall } from '@/services/utils'

export function useStatistics() {
  async function fetchOccupationStatistics(querystring: string) {
    const data = await apiCall(`/statistics/ocupation?${querystring}`, {
      method: 'GET'
    })
    return data
  }
  async function fetchEducationStatistics(querystring: string) {
    const data = await apiCall(`/statistics/education?${querystring}`, {
      method: 'GET'
    })
    return data
  }
  async function fetchCountryStatistics(querystring: string) {
    const data = await apiCall(`/statistics/country?${querystring}`, {
      method: 'GET'
    })
    return data
  }
  async function fetchGenderStatistics(querystring: string) {
    const data = await apiCall(`/statistics/gender?${querystring}`, {
      method: 'GET'
    })
    return data
  }

  return {
    fetchOccupationStatistics,
    fetchEducationStatistics,
    fetchCountryStatistics,
    fetchGenderStatistics
  }
}
