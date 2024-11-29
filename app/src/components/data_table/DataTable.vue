<script setup lang="ts">
import { ref, computed, onBeforeMount } from 'vue'
import { usePersonStore } from '@/stores/persons'
import { usePersons } from '@/services/persons'
import type { IPerson } from '@/types'

const isLoading = ref(true)
const sortOrder = ref<'asc' | 'desc'>('asc')
const sortKey = ref<keyof IPerson>('id')

const personStore = usePersonStore()
const { fetchPaginatedPersons } = usePersons()

const sortedData = computed(() => {
    if (personStore.state.personsData && personStore.state.personsData.length !== 0) {
        const data = [...personStore.state.personsData]
        if (!sortKey.value) return data
        return data.sort((a, b) => {
            const key = sortKey.value as keyof IPerson
            const valA = a[key] || ''
            const valB = b[key] || ''
            if (sortOrder.value === 'asc') return valA > valB ? 1 : -1
            return valA < valB ? 1 : -1
        })
    } else {
        return []
    }
})

onBeforeMount(async () => {
    const query = localStorage.getItem('query') || personStore.state.queryString || '';

    const data = await fetchPaginatedPersons(query);
    personStore.setQueryString(query);
    personStore.setPersonsStore(data);
    isLoading.value = false;
});

const sortTable = (key: keyof IPerson) => {
    if (sortKey.value === key) {
        sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
    } else {
        sortKey.value = key
        sortOrder.value = 'asc'
    }
}

</script>

<template>
    <main class="table-container rounded-md">
        <div v-if="isLoading" class="loading-container text-center">
            <p class="text-lg font-semibold mb-4 text-sky-200">Cargando datos...</p>
        </div>

        <div v-show="!isLoading" class="table-wrapper">
            <table ref="dataTable"
                class="min-w-full table-auto border-collapse bg-white rounded-lg overflow-hidden text-sm">
                <thead class="bg-gray-100 text-gray-700">
                    <tr>
                        <th class="px-4 py-2 text-left">#</th>
                        <th @click="sortTable('id')" class="px-4 py-2 text-left cursor-pointer">Id ↕</th>
                        <th @click="sortTable('age')" class="px-4 py-2 text-left cursor-pointer">Age ↕</th>
                        <th @click="sortTable('workclass')" class="px-4 py-2 text-left cursor-pointer">Workclass ↕</th>
                        <th @click="sortTable('education')" class="px-4 py-2 text-left cursor-pointer">Education ↕</th>
                        <th @click="sortTable('education_num')" class="px-4 py-2 text-left cursor-pointer">Edu Num ↕
                        </th>
                        <th @click="sortTable('marital_status')" class="px-4 py-2 text-left cursor-pointer">Marital
                            Status ↕</th>
                        <th @click="sortTable('occupation')" class="px-4 py-2 text-left cursor-pointer">Occupation ↕
                        </th>
                        <th @click="sortTable('relationship')" class="px-4 py-2 text-left cursor-pointer">Relationship ↕
                        </th>
                        <th @click="sortTable('race')" class="px-4 py-2 text-left cursor-pointer">Race ↕</th>
                        <th @click="sortTable('sex')" class="px-4 py-2 text-left cursor-pointer">Sex ↕</th>
                        <th @click="sortTable('native_country')" class="px-4 py-2 text-left cursor-pointer">Native
                            Country ↕</th>
                        <th @click="sortTable('hours_per_week')" class="px-4 py-2 text-left cursor-pointer">Hours
                            per-week ↕</th>
                        <th @click="sortTable('capital_gain')" class="px-4 py-2 text-left cursor-pointer">Capital gain ↕
                        </th>
                        <th @click="sortTable('capital_loss')" class="px-4 py-2 text-left cursor-pointer">Capital loss ↕
                        </th>
                        <th @click="sortTable('income')" class="px-4 pt-1 text-left cursor-pointer">Income ↕</th>
                    </tr>
                </thead>
                <tbody class="text-gray-600">
                    <tr v-for="(person, index) in sortedData" :key="person.id" class="border-t hover:bg-gray-50">
                        <td class="px-4 py-2">{{ index + 1 }}</td>
                        <td class="px-4 py-2">{{ person.id }}</td>
                        <td class="px-4 py-2">{{ person.age }}</td>
                        <td class="px-4 py-2">{{ person.workclass }}</td>
                        <td class="px-4 py-2">{{ person.education }}</td>
                        <td class="px-4 py-2">{{ person.education_num }}</td>
                        <td class="px-4 py-2">{{ person.marital_status }}</td>
                        <td class="px-4 py-2">{{ person.occupation }}</td>
                        <td class="px-4 py-2">{{ person.relationship }}</td>
                        <td class="px-4 py-2">{{ person.race }}</td>
                        <td class="px-4 py-2">{{ person.sex }}</td>
                        <td class="px-4 py-2">{{ person.native_country }}</td>
                        <td class="px-4 py-2">{{ person.hours_per_week }}</td>
                        <td class="px-4 py-2">{{ person.capital_gain }}</td>
                        <td class="px-4 py-2">{{ person.capital_loss }}</td>
                        <td class="px-4 pt-1">{{ person.income }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </main>
</template>

<style scoped>
.table-container {
    max-width: 100%;
    overflow-x: auto;
    white-space: nowrap;
}

.table-wrapper {
    display: inline-block;
    min-width: 100%;
}
</style>
