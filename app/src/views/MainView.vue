<script setup lang="ts">
import DataTable from '@/components/data_table/DataTable.vue'
import DataFilterContainer from '@/components/data_filters/DataFilterContainer.vue'
import TablePagination from '@/components/data_table/TablePagination.vue'
</script>

<template>
    <div class="w-full flex flex-col justify-between">
        <DataFilterContainer />
        <div class="flex-1 overflow-auto w-full max-w-full px-4 md:px-8">
            <DataTable />
            <TablePagination />
        </div>
    </div>
</template>
