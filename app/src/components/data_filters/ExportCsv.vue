<script setup lang="ts">
import { usePersons } from '@/services/persons'
import { usePersonStore } from '@/stores/persons'
import type { IPerson } from '@/types'

const personStore = usePersonStore()
const { fetchPersonns } = usePersons()

async function exportDataToCsv() {
    const data = await fetchPersonns(personStore.state.queryString)
    const persons = data.persons

    const csvContent = convertToCSV(persons);
    const blob = new Blob([csvContent], { type: "text/csv;charset=utf-8;" });
    const url = URL.createObjectURL(blob);
    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", "export_data.csv");
    link.click();
}

const convertToCSV = (data: IPerson[]) => {
    const headers = Object.keys(data[0]);
    const rows = data.map((obj) => headers.map((header) => (obj as any)[header]));
    const headerRow = headers.join(",");
    const csvRows = [headerRow, ...rows.map((row) => row.join(","))];
    return csvRows.join("\n");
};
</script>

<template>
    <button class="btn-scondary" @click="exportDataToCsv">Export Filtered Data</button>
</template>