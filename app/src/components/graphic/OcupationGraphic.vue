<template>
  <div class="p-4 bg-gray-100 mt-4 mb-8">
    <h2 class="text-lg font-mono font-semibold">Occupation Stats</h2>
    <div class="flex justify-end gap-1">
      <button @click="() => toggleIncome(true)" :class="{
        'bg-sky-500': !moreThan50K,
        'bg-sky-300 opacity-50 cursor-not-allowed': moreThan50K
      }" class="text-xs rounded-md py-1 px-2 items-end" :disabled="moreThan50K">
        {{ `Mas de $50K ` }}
      </button>
      <button @click="() => toggleIncome(false)" :class="{
        'bg-sky-500': moreThan50K,
        'bg-sky-300 opacity-50 cursor-not-allowed': !moreThan50K
      }" class="text-xs rounded-md p-1 items-end" :disabled="!moreThan50K">
        {{ `<= 50K ` }} </button>
    </div>

    <Bar :key="chartKey" id="my-chart-id" :options="chartOptions" :data="chartData" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { Bar } from 'vue-chartjs';
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js';
import { useStatistics } from '@/services/statistics';

const { fetchOccupationStatistics } = useStatistics();

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale);

const chartKey = ref(0);
const moreThan50K = ref(false);

const chartData = ref({
  labels: [] as string[],
  datasets: [
    {
      label: `% of persons with income ${moreThan50K.value ? '>50K' : '<=50K'}`,
      data: [] as number[],
      backgroundColor: ['#42a5f5', '#42d5f5', '#42a8f5'],
    },
  ],
});

const chartOptions = ref({
  responsive: true,
});

const updateChartData = async () => {
  try {
    const incomeParam = moreThan50K.value ? '>50K' : '<=50K';
    const dataJSON = await fetchOccupationStatistics(`income=${incomeParam}`);

    const statistics = dataJSON.statistics.map((row: { option: string; percentage: string }) => ({
      option: row.option,
      percentage: Number(row.percentage),
    }));

    statistics.sort((a: { percentage: number }, b: { percentage: number }) => b.percentage - a.percentage);
    chartData.value.labels = statistics.map((stat: { option: string }) => stat.option);
    chartData.value.datasets[0].data = statistics.map((stat: { percentage: number }) => stat.percentage);

    chartKey.value++;
  } catch (error) {
    console.error("Error fetching statistics:", error);
  }
};

const toggleIncome = (value: boolean) => {
  moreThan50K.value = value;
  updateChartData();
};

onMounted(updateChartData);
</script>
