<script setup>
import { ref } from 'vue'

// Mock data for pending registrations
const pendingRegistrations = ref([
  {
    id: 1,
    submissionDate: '2023-07-15',
    applicantName: 'John Smith',
    applicantEmail: 'john.smith@example.com',
    vehicleType: 'Sedan',
    plateNumber: 'ABC-123',
    status: 'pending',
  },
  {
    id: 2,
    submissionDate: '2023-07-14',
    applicantName: 'Maria Garcia',
    applicantEmail: 'maria.garcia@example.com',
    vehicleType: 'SUV',
    plateNumber: 'XYZ-789',
    status: 'pending',
  },
  {
    id: 3,
    submissionDate: '2023-07-13',
    applicantName: 'David Lee',
    applicantEmail: 'david.lee@example.com',
    vehicleType: 'Motorcycle',
    plateNumber: 'DEF-456',
    status: 'pending',
  },
])

// Function to approve registration
const approveRegistration = (id) => {
  const registration = pendingRegistrations.value.find((reg) => reg.id === id)
  if (registration) {
    registration.status = 'approved'
  }
}

// Function to reject registration
const rejectRegistration = (id) => {
  const registration = pendingRegistrations.value.find((reg) => reg.id === id)
  if (registration) {
    registration.status = 'rejected'
  }
}
</script>

<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <h2 class="text-2xl font-bold text-gray-900">Pending Registrations</h2>
    </div>

    <!-- Pending Registrations Table -->
    <div class="bg-white shadow-md rounded-lg overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th
              scope="col"
              class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Submission Date
            </th>
            <th
              scope="col"
              class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Applicant
            </th>
            <th
              scope="col"
              class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Vehicle Type
            </th>
            <th
              scope="col"
              class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Plate Number
            </th>
            <th
              scope="col"
              class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Status
            </th>
            <th
              scope="col"
              class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="registration in pendingRegistrations" :key="registration.id">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ registration.submissionDate }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm font-medium text-gray-900">{{ registration.applicantName }}</div>
              <div class="text-sm text-gray-500">{{ registration.applicantEmail }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ registration.vehicleType }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ registration.plateNumber }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span
                :class="[
                  'px-2 inline-flex text-xs leading-5 font-semibold rounded-full',
                  registration.status === 'pending'
                    ? 'bg-yellow-100 text-yellow-800'
                    : registration.status === 'approved'
                      ? 'bg-green-100 text-green-800'
                      : 'bg-red-100 text-red-800',
                ]"
              >
                {{ registration.status }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
              <button
                @click="approveRegistration(registration.id)"
                class="text-green-600 hover:text-green-900"
                :disabled="registration.status !== 'pending'"
              >
                Approve
              </button>
              <button
                @click="rejectRegistration(registration.id)"
                class="text-red-600 hover:text-red-900"
                :disabled="registration.status !== 'pending'"
              >
                Reject
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
