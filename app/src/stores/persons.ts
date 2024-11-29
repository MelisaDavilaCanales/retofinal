import { defineStore } from 'pinia'
import { reactive } from 'vue'
import type { IPerson } from '@/types'

export const usePersonStore = defineStore('person', () => {
  const state = reactive({
    personsData: [] as IPerson[],
    totalPages: 0,
    pageSize: 0,
    currentPage: 0,
    queryString: '',
    queryStringBack: '',
    filterQuery: '',
    paginationQuery: ''
  })

  function setPersonsStore(data: {
    persons: IPerson[]
    totalPages: number
    pageSize: number
    page: number
  }) {
    state.personsData = data.persons
    state.totalPages = data.totalPages
    state.pageSize = data.pageSize
    state.currentPage = data.page
  }

  function setQueryString(query: string) {
    state.queryStringBack = query
  }

  function setPersonsData(persons: IPerson[]) {
    state.personsData = persons
  }

  function setFilterQuery(query: string) {
    state.filterQuery = `&${query}`
  }

  function setPaginationQuery(query: string) {
    state.paginationQuery = `&${query}`
  }

  function updateQueryString() {
    state.queryString = state.filterQuery + state.paginationQuery
  }

  function prevPage() {
    if (state.currentPage > 1) {
      state.currentPage--
    }
  }

  function nextPage() {
    if (state.currentPage < state.totalPages) {
      state.currentPage++
    }
  }

  return {
    state,
    setPersonsStore,
    setQueryString,
    setFilterQuery,
    setPaginationQuery,
    updateQueryString,
    setPersonsData,
    prevPage,
    nextPage
  }
})
