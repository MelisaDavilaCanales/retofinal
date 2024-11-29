import { defineStore } from 'pinia'
import { reactive } from 'vue'
import { filtersData } from '@/static/filters'

export const useDataFilterStore = defineStore('dataFilter', () => {
  const state = reactive({
    filtersMap: filtersData.reduce(
      (map, filter) => {
        map[filter.name] = filter.options
        return map
      },
      {} as Record<string, string[]>
    ),
    selectedOptions: filtersData.reduce(
      (map, filter) => {
        map[filter.name] = []
        return map
      },
      {} as Record<string, string[]>
    ),
    selectedFilter: null as string | null
  })

  function setFiltersMap(filtersMap: Record<string, string[]>) {
    state.filtersMap = filtersMap
  }

  function setSelectedOptions(selectedOptions: Record<string, string[]>) {
    state.selectedOptions = selectedOptions
  }

  function setSelectedFilter(selectedFilter: string | null) {
    state.selectedFilter = selectedFilter
  }

  return {
    state,
    setFiltersMap,
    setSelectedOptions,
    setSelectedFilter
  }
})
