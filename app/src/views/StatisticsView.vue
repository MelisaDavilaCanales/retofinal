<script setup lang="ts">
import OcupationGraphic from '@/components/graphic/OcupationGraphic.vue'
import EducationGraphic from '@/components/graphic/EducationGraphic.vue'
import GenderGraphic from '@/components/graphic/GenderGraphic.vue'
</script>

<template>
    <div class="flex flex-col lg:flex-row justify-center items-center lg:justify-around">
        <div class="w-11/12 lg:w-5/12">
            <OcupationGraphic />
            <EducationGraphic />
        </div>
        <div class="w-11/12 lg:w-5/12">
            <GenderGraphic />
        </div>
    </div>
</template>