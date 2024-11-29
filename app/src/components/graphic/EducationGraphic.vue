<template>
    <div class="p-4 bg-gray-100 mt-4 mb-8">
        <h2 class="text-lg font-mono font-semibold">Education Stats</h2>
        <div class="flex justify-end gap-1 mb-3">
            <button @click="toggleIncome()" :class="{
                'bg-sky-500': !moreThan50K,
                'bg-sky-300 opacity-50 cursor-not-allowed': moreThan50K
            }" class="text-xs rounded-md py-1 px-2 items-end" :disabled="moreThan50K">
                {{ `Mas de $50K ` }}
            </button>
            <button @click="toggleIncome()" :class="{
                'bg-sky-500': moreThan50K,
                'bg-sky-300 opacity-50 cursor-not-allowed': !moreThan50K
            }" class="text-xs rounded-md p-1 items-end" :disabled="!moreThan50K">
                {{ `Menos de $50K` }}</button>
        </div>
        <div class="w-full h-64">
            <Pie :key="chartKey" :data="data" :options="options" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Pie } from 'vue-chartjs';

import { useStatistics } from '@/services/statistics';

const { fetchEducationStatistics } = useStatistics();

ChartJS.register(ArcElement, Tooltip, Legend);

const chartKey = ref(0);
const moreThan50K = ref(false);

const data = ref({
    labels: [] as string[],
    datasets: [
        {
            label: `% of persons with income ${moreThan50K.value ? '>50K' : '<=50K'}`,
            data: [] as number[],
            backgroundColor: [
                '#3f51b5', '#81c784', '#4caf50', '#9c27b0', '#388e3c',
                '#ffeb3b', '#ffb74d', '#ff7043', '#8d6e63', '#c5e1a5',
                '#ffb300', '#f57c00', '#ff5722', '#c2185b', '#8e24aa',
                '#7e57c2', '#00897b',
            ],
        },
    ],
});

const options = ref({
    responsive: true,
    maintainAspectRatio: false,
});

const updateChartData = async () => {
    try {
        const incomeParam = moreThan50K.value ? '>50K' : '<=50K';
        const dataJSON = await fetchEducationStatistics(`income=${incomeParam}`);

        const statistics = dataJSON.statistics.map((row: { option: string; percentage: string }) => ({
            option: row.option,
            percentage: Number(row.percentage),
        }));

        statistics.sort((a: { percentage: number }, b: { percentage: number }) => b.percentage - a.percentage);
        data.value.labels = statistics.map((stat: { option: string }) => stat.option);
        data.value.datasets[0].data = statistics.map((stat: { percentage: number }) => stat.percentage);

        chartKey.value++;
    } catch (error) {
        console.error("Error fetching statistics:", error);
    }
};

const toggleIncome = () => {
    moreThan50K.value = !moreThan50K.value;
    updateChartData();
};

onMounted(updateChartData);
</script>
