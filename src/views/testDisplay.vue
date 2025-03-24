<template>
  <div class="container mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4">Vehicle List</h1>

    <div v-if="loading" class="text-gray-600">Loading vehicles...</div>
    <div v-if="error" class="text-red-600">{{ error }}</div>

    <table v-if="vehicles.length" class="min-w-full divide-y divide-gray-200">
      <thead class="bg-gray-50">
        <tr>
          <th
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            ID
          </th>
          <th
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Make
          </th>
          <th
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Model
          </th>
          <th
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            PlateNumber
          </th>
          <th
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Created
          </th>
          <th
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Updated
          </th>
        </tr>
      </thead>
      <tbody class="bg-white divide-y divide-gray-200">
        <tr v-for="vehicle in vehicles" :key="vehicle.id">
          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ vehicle.id }}</td>
          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ vehicle.make }}</td>
          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
            {{ vehicle.platenumber }}
          </td>
          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ vehicle.model }}</td>
          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
            {{ formatDate(vehicle.created) }}
          </td>
          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
            {{ formatDate(vehicle.updated) }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const vehicles = ref([])
const loading = ref(true)
const error = ref(null)

const fetchVehicles = async () => {
  try {
    const response = await axios.get('http://localhost:1323/vehicle')
    vehicles.value = response.data.items
    loading.value = false
  } catch (err) {
    error.value = 'Failed to load vehicles: ' + err.message
    loading.value = false
  }
}

const formatDate = (timestamp) => {
  if (!timestamp) return '-'
  const date = new Date(timestamp)
  return date.toLocaleDateString() + ' ' + date.toLocaleTimeString()
}

onMounted(() => {
  fetchVehicles()
})
</script>

<style scoped>
/* Add any custom styles here if needed */
</style>
