<script setup lang="ts">
import { useVehicleRegistrationStore } from '@/stores/vehicleRegistration'
import { useVehicleRegistrationFormStore } from '@/stores/vehicleRegistrationForm'
import { computed } from 'vue'
import LTODashboardCharts from '@/components/charts/LTODashboardCharts.vue'

const vehicleRegistrationStore = useVehicleRegistrationStore()
const registrationFormStore = useVehicleRegistrationFormStore()

const pendingRegistrations = computed(
  () => registrationFormStore.forms.filter((form) => form.status === 'pending').length,
)

const platesForIssuance = computed(
  () =>
    vehicleRegistrationStore.registrations.filter(
      (reg) => reg.status === 'pending' || reg.status === 'ready',
    ).length,
)

const completedToday = computed(() => {
  const today = new Date()
  return registrationFormStore.forms.filter((form) => {
    const completionDate = new Date(form.submissionDate)
    return form.status === 'approved' && completionDate.toDateString() === today.toDateString()
  }).length
})
</script>

<template>
  <div>
    <h1 class="text-2xl font-semibold text-gray-900 mb-6">LTO Officer Dashboard</h1>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-xl font-semibold mb-4">Pending Registrations</h2>
        <p class="text-3xl font-bold text-blue-900">{{ pendingRegistrations }}</p>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-xl font-semibold mb-4">Plates for Issuance</h2>
        <p class="text-3xl font-bold text-blue-900">{{ platesForIssuance }}</p>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-xl font-semibold mb-4">Completed Today</h2>
        <p class="text-3xl font-bold text-blue-900">{{ completedToday }}</p>
      </div>
    </div>
  </div>
  <LTODashboardCharts />
</template>
