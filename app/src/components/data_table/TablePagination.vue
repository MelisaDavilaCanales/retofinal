<script setup lang="ts">
import { usePersonStore } from '@/stores/persons'
import { usePersons } from '@/services/persons'

const personStore = usePersonStore()
const { fetchPaginatedPersons } = usePersons()

async function handlePagination(value: number) {
    if (value === -1) {
        personStore.prevPage()
    } else {
        personStore.nextPage()
    }
    personStore.setPaginationQuery('page=' + personStore.state.currentPage + '&pageSize=10')
    personStore.updateQueryString()
    const data = await fetchPaginatedPersons(personStore.state.queryString)
    personStore.setPersonsData(data.persons)
}
</script>

<template>
    <div id="pagination" class="flex items-center justify-between border-t border-gray-200 bg-white px-4 pt-2 sm:px-6">
        <div class="flex justify-between w-screen  sm:flex sm:flex-1 sm:items-center sm:justify-between">
            <div>
                <p class="text-sm text-gray-700 sm:flex space-x-1">
                    <span>Page</span>
                    <span class="font-extrabold text-primary">{{ personStore.state.currentPage }}</span>
                    <span>of</span>
                    <span class="font-extrabold text-primary">{{ personStore.state.totalPages }}</span>
                    <span class="hidden sm:block"> | Page size:</span>
                    <span class="hidden sm:block font-medium">{{ personStore.state.pageSize }}</span>
                </p>
            </div>
            <div>
                <nav class="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">
                    <a href="#" :disabled="personStore.state.currentPage === 1" @click="() => handlePagination(-1)"
                        :class="{
                            'text-gray-300 cursor-not-allowed': personStore.state.currentPage === 1,
                            'text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0': personStore.state.currentPage !== 1
                        }" class="relative inline-flex items-center rounded-l-md px-2 py-1">
                        <span class="sr-only">Previous</span>
                        <svg class="size-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" data-slot="icon">
                            <path fill-rule="evenodd"
                                d="M11.78 5.22a.75.75 0 0 1 0 1.06L8.06 10l3.72 3.72a.75.75 0 1 1-1.06 1.06l-4.25-4.25a.75.75 0 0 1 0-1.06l4.25-4.25a.75.75 0 0 1 1.06 0Z"
                                clip-rule="evenodd" />
                        </svg>
                    </a>

                    <a href="#" aria-current="page"
                        class="relative z-10 inline-flex items-center justify-center bg-primary px-4 py-1 text-sm font-semibold text-white focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primaryVariant text-center self-center w-14">
                        {{ personStore.state.currentPage }}
                    </a>

                    <a href="#" @click="() => handlePagination(+1)"
                        :disabled="personStore.state.currentPage === personStore.state.totalPages" :class="{
                            'text-gray-300 cursor-not-allowed': personStore.state.currentPage === personStore.state.totalPages,
                            'text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0': personStore.state.currentPage !== personStore.state.totalPages
                        }" class="relative inline-flex items-center rounded-r-md px-2 py-1">
                        <span class="sr-only">Next</span>
                        <svg class="size-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" data-slot="icon">
                            <path fill-rule="evenodd"
                                d="M8.22 5.22a.75.75 0 0 1 1.06 0l4.25 4.25a.75.75 0 0 1 0 1.06l-4.25 4.25a.75.75 0 0 1-1.06-1.06L11.94 10 8.22 6.28a.75.75 0 0 1 0-1.06Z"
                                clip-rule="evenodd" />
                        </svg>
                    </a>
                </nav>
            </div>

        </div>
    </div>
</template>
