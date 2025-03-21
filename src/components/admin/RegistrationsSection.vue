<script setup>
import { ref, computed } from 'vue'
import { useVehicleRegistrationStore } from '@/stores/vehicleRegistration'

const vehicleStore = useVehicleRegistrationStore()

// Get all registrations with details
const registrations = computed(() => vehicleStore.registrationsWithDetails)

// Status filters
const statusFilters = ref([
  { value: 'all', label: 'All Registrations', active: true },
  { value: 'Approved', label: 'Approved', active: false },
  { value: 'Pending', label: 'Pending', active: false },
  { value: 'expired', label: 'Expired', active: false },
])

// Type filters
const typeFilters = ref([
  { value: 'all', label: 'All Types', active: true },
  { value: 'New Registration', label: 'New Registration', active: false },
  { value: 'Renewal', label: 'Renewal', active: false },
])

// Search query
const searchQuery = ref('')

// Active filters
const activeStatusFilter = computed(() => {
  const filter = statusFilters.value.find((f) => f.active === true)
  return filter ? filter.value : 'all'
})

const activeTypeFilter = computed(() => {
  const filter = typeFilters.value.find((f) => f.active === true)
  return filter ? filter.value : 'all'
})

// Apply status filter
const setStatusFilter = (value) => {
  statusFilters.value = statusFilters.value.map((filter) => ({
    ...filter,
    active: filter.value === value,
  }))
}

// Apply type filter
const setTypeFilter = (value) => {
  typeFilters.value = typeFilters.value.map((filter) => ({
    ...filter,
    active: filter.value === value,
  }))
}

// Filtered registrations based on active filters and search query
const filteredRegistrations = computed(() => {
  let result = registrations.value

  // Apply status filter
  if (activeStatusFilter.value !== 'all') {
    if (activeStatusFilter.value === 'expired') {
      // Filter for expired registrations
      result = result.filter((reg) => {
        const expiryDate = new Date(reg.expiryDate)
        return expiryDate < new Date()
      })
    } else {
      // Filter by status
      result = result.filter((reg) => reg.status === activeStatusFilter.value)
    }
  }

  // Apply type filter
  if (activeTypeFilter.value !== 'all') {
    result = result.filter((reg) => reg.registrationType === activeTypeFilter.value)
  }

  // Apply search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    result = result.filter(
      (reg) =>
        reg.vehicleInfo.toLowerCase().includes(query) ||
        reg.plateNumber.toLowerCase().includes(query),
    )
  }

  return result
})

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
    </div>

    <!-- Filters and Search -->
    <div class="bg-white rounded-lg shadow-sm p-4 mb-6">
      <div class="space-y-4">
        <!-- Search Bar -->
        <div class="relative">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search by vehicle or plate number..."
            class="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-200 focus:outline-none focus:ring-2 focus:ring-dark-blue/20 transition-all"
          />
          <div class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400">
            <i class="fas fa-search w-4 h-4"></i>
          </div>
        </div>

        <!-- Filter Options -->
        <div class="flex flex-wrap gap-2">
          <div class="mr-4">
            <span class="text-sm font-medium text-gray-700 mr-2">Status:</span>
            <div class="inline-flex rounded-md shadow-sm mt-1">
              <button
                v-for="filter in statusFilters"
                :key="filter.value"
                @click="setStatusFilter(filter.value)"
                :class="[
                  'px-3 py-1 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-dark-blue',
                  filter.active
                    ? 'bg-dark-blue text-white rounded-md'
                    : 'text-gray-700 hover:bg-gray-100 rounded-md',
                ]"
              >
                {{ filter.label }}
              </button>
            </div>
          </div>

          <div>
            <span class="text-sm font-medium text-gray-700 mr-2">Type:</span>
            <div class="inline-flex rounded-md shadow-sm mt-1">
              <button
                v-for="filter in typeFilters"
                :key="filter.value"
                @click="setTypeFilter(filter.value)"
                :class="[
                  'px-3 py-1 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-dark-blue',
                  filter.active
                    ? 'bg-dark-blue text-white rounded-md'
                    : 'text-gray-700 hover:bg-gray-100 rounded-md',
                ]"
              >
                {{ filter.label }}
              </button>
            </div>
          </div>
        </div>
      </div>
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
          <tr v-for="registration in filteredRegistrations" :key="registration.id">
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
              {{ registration.status }}
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
