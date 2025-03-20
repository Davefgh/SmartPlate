<script setup>
import { computed } from 'vue'
import { useVehicleRegistrationStore } from '@/stores/vehicleRegistration'

const vehicleStore = useVehicleRegistrationStore()

// Get all registrations with details
const registrations = computed(() => vehicleStore.registrationsWithDetails)

// Table headers
const headers = [
  { text: 'Vehicle', value: 'vehicleInfo' },
  { text: 'Plate Number', value: 'plateNumber' },
  { text: 'Registration Type', value: 'registrationType' },
  { text: 'Submission Date', value: 'submissionDate' },
  { text: 'Expiry Date', value: 'expiryDate' },
  { text: 'Status', value: 'status' },
  { text: 'Actions', value: 'actions' },
]
</script>

<template>
  <div class="p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-semibold text-gray-800">Registrations Management</h2>
      <button
        class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors"
      >
        Add New Registration
      </button>
    </div>

    <!-- Registrations Table -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th
              v-for="header in headers"
              :key="header.value"
              class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              {{ header.text }}
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="registration in registrations" :key="registration.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ registration.vehicleInfo }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ registration.plateNumber }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ registration.registrationType }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ registration.submissionDate }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ registration.expiryDate }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              <span
                :class="{
                  'px-2 py-1 rounded-full text-xs font-medium': true,
                  'bg-green-100 text-green-800': registration.status === 'Approved',
                  'bg-yellow-100 text-yellow-800': registration.status === 'Pending',
                  'bg-red-100 text-red-800': registration.status === 'Rejected',
                }"
              >
                {{ registration.status }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              <button class="text-indigo-600 hover:text-indigo-900 mr-3" @click="() => {}">
                Edit
              </button>
              <button class="text-red-600 hover:text-red-900" @click="() => {}">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
