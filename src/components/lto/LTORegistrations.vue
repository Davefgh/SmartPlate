<script setup lang="ts">
import type { VehicleRegistrationForm } from '@/types/vehicleRegistration'

interface RegistrationForm extends VehicleRegistrationForm {}

import { useVehicleRegistrationFormStore } from '@/stores/vehicleRegistrationForm'
import { computed, ref } from 'vue'

const registrationFormStore = useVehicleRegistrationFormStore()
const sortBy = ref('submissionDate')
const sortOrder = ref('desc')

const pendingRegistrations = computed<RegistrationForm[]>(() =>
  registrationFormStore.forms.filter((form) => form.status === 'pending'),
)

const sortedRegistrations = computed<RegistrationForm[]>(() => {
  return [...pendingRegistrations.value].sort((a, b) => {
    const aValue = (a as any)[sortBy.value] || ''
    const bValue = (b as any)[sortBy.value] || ''
    const order = sortOrder.value === 'asc' ? 1 : -1
    return aValue > bValue ? order : -order
  })
})

const toggleSort = (field: 'id' | 'vehicleType' | 'applicantName' | 'submissionDate') => {
  if (sortBy.value === field) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortBy.value = field
    sortOrder.value = 'desc'
  }
}

const processRegistration = async (registrationId: string, action: 'approve' | 'reject') => {
  await registrationFormStore.$patch((state) => {
    const form = state.forms.find((f) => f.id === registrationId)
    if (form) {
      form.status = action === 'approve' ? 'approved' : 'rejected'
    }
  })
}
</script>

<template>
  <div class="p-6">
    <h1 class="text-2xl font-semibold text-gray-900 mb-6">Pending Vehicle Registrations</h1>
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                v-for="header in ['Registration ID', 'Vehicle Details', 'Owner', 'Submission Date']"
                :key="header"
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100"
                @click="
                  toggleSort(
                    header.toLowerCase().replace(' ', '') as
                      | 'id'
                      | 'vehicleType'
                      | 'applicantName'
                      | 'submissionDate',
                  )
                "
              >
                {{ header }}
                <span v-if="sortBy === header.toLowerCase().replace(' ', '')">
                  {{ sortOrder === 'asc' ? '↑' : '↓' }}
                </span>
              </th>
              <th
                class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="registration in sortedRegistrations" :key="registration.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ registration.id }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ `${registration.make} ${registration.model} (${registration.year})` }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ registration.applicantName || 'Unknown' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ new Date(registration.submissionDate as string).toLocaleDateString() }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
                <button
                  @click="processRegistration(registration.id, 'approve')"
                  class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                >
                  Approve
                </button>
                <button
                  @click="processRegistration(registration.id, 'reject')"
                  class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                >
                  Reject
                </button>
              </td>
            </tr>
            <tr v-if="sortedRegistrations.length === 0">
              <td colspan="5" class="px-6 py-4 text-center text-sm text-gray-500">
                No pending registrations found
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
