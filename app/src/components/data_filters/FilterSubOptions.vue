<script setup lang="ts">
import { useDataFilterStore } from '@/stores/dataFilter'

const dataFilterStore = useDataFilterStore()
</script>

<template>
    <section v-if="dataFilterStore.state.selectedFilter" class="w-full text-xs mx-3 mt-2">
        <hr class="w-[95%] mx-auto sm:-[90%] md:w-[85%] border-t border-gray-200 mb-2">
        <div class="sm:flex sm:flex-col sm:items-center sm:justify-center">
            <div class="grid grid-cols-2 gap-1 gap-x-8 sm:grid-cols-5 md:grid-cols-6 lg:grid-cols-8 xl:grid-cols-8">
                <div v-for="option in dataFilterStore.state.filtersMap[dataFilterStore.state.selectedFilter]"
                    :key="option">
                    <label class="inline-flex items-center space-x-2">
                        <input type="checkbox" :value="option"
                            v-model="dataFilterStore.state.selectedOptions[dataFilterStore.state.selectedFilter]" />
                        <span></span>
                        {{ option }}
                    </label>
                </div>
            </div>
        </div>
        <hr class="w-[95%] sm:-[90%] md:w-[85%] mx-auto border-t border-gray-200 mt-2">
    </section>
</template>