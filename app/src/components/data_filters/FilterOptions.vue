<script setup lang="ts">
import { computed } from 'vue'
import { usePersonStore } from '@/stores/persons'
import { useDataFilterStore } from '@/stores/dataFilter'
import { usePersons } from '@/services/persons'
import { filtersData } from '@/static/filters'
import ExportCsv from './ExportCsv.vue'


const personStore = usePersonStore()
const dataFilterStore = useDataFilterStore()

const filterParams = computed(() => buildFiltersParams(dataFilterStore.state.selectedOptions))

function buildFiltersParams(selectedOptions: Record<string, string[]>): string {
    return Object.entries(selectedOptions)
        .map(([key, values]) => {
            if (values.length > 0) {
                return `${key}=${values.join(',')}`
            }
            return null
        })
        .filter(Boolean)
        .join('&')
}

function selectFilter(filterName: string) {
    const value = dataFilterStore.state.selectedFilter === filterName ? null : filterName
    dataFilterStore.setSelectedFilter(value)
}

const { fetchPaginatedPersons } = usePersons()

async function sendFilter() {
    localStorage.setItem('query', filterParams.value)
    personStore.setFilterQuery(filterParams.value)
    personStore.updateQueryString()
    const newData = await fetchPaginatedPersons(`${filterParams.value}`)
    personStore.setPersonsStore(newData)
}

async function clearFilter() {
    dataFilterStore.setSelectedOptions(filtersData.reduce((map, filter) => {
        map[filter.name] = []
        return map
    }, {} as Record<string, string[]>))
    personStore.setFilterQuery('')
    personStore.updateQueryString()
    const newData = await fetchPaginatedPersons('')
    personStore.setPersonsStore(newData)
}
</script>

<template>
    <section
        :class="['flex flex-col w-full sm:flex mt-2 space-y-2', dataFilterStore.state.selectedFilter === null ? 'pb-6 pt-3' : '']">

        <div class="flex flex-col items-center justify-center sm:flex sm:flex-row">
            <span class="hidden sm:block text-xs px-1 text-gray-700 font-roboto rounded-sm">Filter by:</span>
            <div class="flex">
                <div v-for="filter in Object.keys(dataFilterStore.state.filtersMap)" :key="filter" :class="[
                    'bg-complementarySoft rounded-sm text-xs py-1 px-2 m-1 cursor-pointer',
                    { 'font-semibold bg-yellow-600 ': dataFilterStore.state.selectedFilter === filter },
                ]" @click="selectFilter(filter)">
                    <h3 class="capitalize">{{ filter }}</h3>
                </div>
            </div>
        </div>
        <div class="flex justify-center space-x-3">
            <button @click="sendFilter" class="btn-scondary">
                Apply filters
            </button>
            <button @click="clearFilter" class="btn-scondary">
                Clear Filters
            </button>
            <ExportCsv />
        </div>
    </section>
</template>