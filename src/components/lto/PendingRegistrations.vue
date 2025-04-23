<script setup lang="ts">
import type { VehicleRegistrationForm } from '@/types/vehicleRegistration'
import { computed, ref, onMounted, watch } from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

const props = defineProps<{
  registrations: VehicleRegistrationForm[]
}>()

const emit = defineEmits<{
  (e: 'process', data: { registrationId: string; action: 'approve' | 'reject' }): void
}>()

const sortBy = ref('submissionDate')
const sortOrder = ref('desc')

const isLoading = ref(true)

// Simulate loading delay
onMounted(() => {
  isLoading.value = true
  setTimeout(() => {
    isLoading.value = false
  }, 1500)
})

// Search and filter
const searchQuery = ref('')
const filterOptions = ref({
  vehicleType: '',
  year: '',
})

// Pagination
const currentPage = ref(1)
const itemsPerPage = 10
const totalPages = computed(() => Math.ceil(filteredRegistrations.value.length / itemsPerPage))

// Filtered registrations based on search query and filters
const filteredRegistrations = computed(() => {
  return props.registrations.filter((reg) => {
    // Search query filter
    const searchLower = searchQuery.value.toLowerCase()
    const matchesSearch =
      searchLower === '' ||
      reg.id.toLowerCase().includes(searchLower) ||
      `${reg.make} ${reg.model}`.toLowerCase().includes(searchLower) ||
      (reg.applicantName && reg.applicantName.toLowerCase().includes(searchLower))

    // Dropdown filters
    const matchesVehicleType =
      filterOptions.value.vehicleType === '' || reg.vehicleType === filterOptions.value.vehicleType

    const matchesYear =
      filterOptions.value.year === '' || reg.year.toString() === filterOptions.value.year

    return matchesSearch && matchesVehicleType && matchesYear
  })
})

// Sorted and paginated registrations
const displayedRegistrations = computed(() => {
  const sorted = [...filteredRegistrations.value].sort((a, b) => {
    const aValue = (a as any)[sortBy.value] || ''
    const bValue = (b as any)[sortBy.value] || ''
    const order = sortOrder.value === 'asc' ? 1 : -1
    return aValue > bValue ? order : -order
  })

  const startIndex = (currentPage.value - 1) * itemsPerPage
  return sorted.slice(startIndex, startIndex + itemsPerPage)
})

// Get unique values for filter dropdowns
const vehicleTypes = computed(() => {
  const types = new Set(props.registrations.map((reg) => reg.vehicleType))
  return Array.from(types)
})

const years = computed(() => {
  const yearsSet = new Set(props.registrations.map((reg) => reg.year.toString()))
  return Array.from(yearsSet).sort((a, b) => b.localeCompare(a)) // Sort descending
})

// Reset pagination when filters change
watch([searchQuery, filterOptions], () => {
  currentPage.value = 1
})

// Methods for pagination
const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
  }
}

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
  }
}

const goToPage = (page: number) => {
  currentPage.value = page
}

// Reset filters
const resetFilters = () => {
  searchQuery.value = ''
  filterOptions.value = {
    vehicleType: '',
    year: '',
  }
}

const toggleSort = (field: 'id' | 'vehicleType' | 'applicantName' | 'submissionDate') => {
  if (sortBy.value === field) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortBy.value = field
    sortOrder.value = 'desc'
  }
}

const processRegistration = (registrationId: string, action: 'approve' | 'reject') => {
  emit('process', { registrationId, action })
}
</script>

<template>
  <div class="bg-white rounded-lg shadow overflow-hidden mb-8">
    <!-- Loading State -->
    <div v-if="isLoading" class="p-12 flex flex-col items-center justify-center">
      <div class="relative w-20 h-20 mb-4">
        <div class="w-20 h-20 rounded-full border-4 border-blue-100 animate-pulse"></div>
        <div
          class="absolute top-0 left-0 w-20 h-20 rounded-full border-t-4 border-blue-600 animate-spin"
        ></div>
      </div>
      <div class="text-center">
        <p class="text-lg font-medium text-gray-700 mb-1">Loading Registrations</p>
        <p class="text-sm text-gray-500">Please wait while we fetch the latest data...</p>
      </div>
      <div class="mt-6 flex space-x-1">
        <div
          class="w-3 h-3 bg-blue-600 rounded-full animate-bounce"
          style="animation-delay: 0ms"
        ></div>
        <div
          class="w-3 h-3 bg-blue-600 rounded-full animate-bounce"
          style="animation-delay: 150ms"
        ></div>
        <div
          class="w-3 h-3 bg-blue-600 rounded-full animate-bounce"
          style="animation-delay: 300ms"
        ></div>
      </div>
    </div>

    <!-- Content -->
    <div v-else>
      <!-- Search and Filter Bar -->
      <div class="p-4 bg-gray-50 border-b border-gray-200">
        <div
          class="flex flex-col md:flex-row md:items-center md:justify-between space-y-3 md:space-y-0 md:space-x-4"
        >
          <!-- Search Box -->
          <div class="flex flex-1 items-center space-x-2">
            <div class="relative flex-1 max-w-md">
              <span class="absolute inset-y-0 left-0 flex items-center pl-3">
                <font-awesome-icon icon="search" class="h-4 w-4 text-gray-400" />
              </span>
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Search by ID, vehicle or owner..."
                class="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md text-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
          </div>

          <!-- Filters -->
          <div class="flex flex-wrap items-center space-x-2">
            <div class="flex items-center space-x-2">
              <label class="text-sm text-gray-600">Vehicle Type:</label>
              <select
                v-model="filterOptions.vehicleType"
                class="block w-40 pl-3 pr-10 py-2 text-sm border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">All Types</option>
                <option v-for="type in vehicleTypes" :key="type" :value="type">
                  {{ type }}
                </option>
              </select>
            </div>

            <div class="flex items-center space-x-2">
              <label class="text-sm text-gray-600">Year:</label>
              <select
                v-model="filterOptions.year"
                class="block w-28 pl-3 pr-10 py-2 text-sm border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">All Years</option>
                <option v-for="year in years" :key="year" :value="year">
                  {{ year }}
                </option>
              </select>
            </div>

            <button
              @click="resetFilters"
              class="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <font-awesome-icon icon="times" class="h-3 w-3 mr-1" />
              Clear
            </button>
          </div>
        </div>
      </div>

      <!-- Table -->
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
            <tr v-if="!filteredRegistrations.length" class="hover:bg-gray-50">
              <td colspan="5" class="px-6 py-12 text-center">
                <div class="flex flex-col items-center">
                  <div class="rounded-full bg-blue-50 p-3 mb-4">
                    <font-awesome-icon icon="search" class="h-6 w-6 text-blue-500" />
                  </div>
                  <h3 class="text-sm font-medium text-gray-900 mb-1">No Matching Registrations</h3>
                  <p class="text-sm text-gray-500">Try adjusting your search or filter criteria.</p>
                  <button
                    @click="resetFilters"
                    class="mt-3 inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                  >
                    Clear Filters
                  </button>
                </div>
              </td>
            </tr>
            <tr v-else-if="!displayedRegistrations.length" class="hover:bg-gray-50">
              <td colspan="5" class="px-6 py-12 text-center">
                <div class="flex flex-col items-center">
                  <div class="rounded-full bg-blue-50 p-3 mb-4">
                    <font-awesome-icon icon="clipboard-list" class="h-6 w-6 text-blue-500" />
                  </div>
                  <h3 class="text-sm font-medium text-gray-900 mb-1">No Pending Registrations</h3>
                  <p class="text-sm text-gray-500">
                    Waiting for new vehicle registration applications.
                  </p>
                </div>
              </td>
            </tr>
            <tr
              v-for="registration in displayedRegistrations"
              :key="registration.id"
              class="hover:bg-gray-50"
            >
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
                  Accept
                </button>
                <button
                  @click="processRegistration(registration.id, 'reject')"
                  class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                >
                  Reject
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div
        v-if="filteredRegistrations.length > 0"
        class="px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6"
      >
        <div class="flex-1 flex justify-between sm:hidden">
          <button
            @click="prevPage"
            :disabled="currentPage === 1"
            :class="[
              'relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md',
              currentPage === 1
                ? 'text-gray-300 bg-gray-50 cursor-not-allowed'
                : 'text-gray-700 bg-white hover:bg-gray-50',
            ]"
          >
            Previous
          </button>
          <button
            @click="nextPage"
            :disabled="currentPage === totalPages"
            :class="[
              'ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md',
              currentPage === totalPages
                ? 'text-gray-300 bg-gray-50 cursor-not-allowed'
                : 'text-gray-700 bg-white hover:bg-gray-50',
            ]"
          >
            Next
          </button>
        </div>
        <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
          <div>
            <p class="text-sm text-gray-700">
              Showing
              <span class="font-medium">
                {{ filteredRegistrations.length ? (currentPage - 1) * itemsPerPage + 1 : 0 }}
              </span>
              to
              <span class="font-medium">
                {{ Math.min(currentPage * itemsPerPage, filteredRegistrations.length) }}
              </span>
              of
              <span class="font-medium">{{ filteredRegistrations.length }}</span>
              results
            </p>
          </div>
          <div>
            <nav
              class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px"
              aria-label="Pagination"
            >
              <button
                @click="prevPage"
                :disabled="currentPage === 1"
                class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                :class="{ 'opacity-50 cursor-not-allowed': currentPage === 1 }"
              >
                <font-awesome-icon icon="chevron-left" class="h-3 w-3" />
              </button>

              <!-- Show first page -->
              <button
                v-if="totalPages > 0"
                @click="goToPage(1)"
                :class="[
                  'relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium',
                  currentPage === 1 ? 'bg-blue-50 text-blue-600' : 'text-gray-500 hover:bg-gray-50',
                ]"
              >
                1
              </button>

              <!-- Separator for "..." -->
              <span
                v-if="totalPages > 3 && currentPage > 2"
                class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700"
              >
                ...
              </span>

              <!-- Current page (if not first or last) -->
              <button
                v-if="totalPages > 1 && currentPage !== 1 && currentPage !== totalPages"
                class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-blue-50 text-sm font-medium text-blue-600"
              >
                {{ currentPage }}
              </button>

              <!-- Separator for "..." -->
              <span
                v-if="totalPages > 3 && currentPage < totalPages - 1"
                class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700"
              >
                ...
              </span>

              <!-- Show last page -->
              <button
                v-if="totalPages > 1"
                @click="goToPage(totalPages)"
                :class="[
                  'relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium',
                  currentPage === totalPages
                    ? 'bg-blue-50 text-blue-600'
                    : 'text-gray-500 hover:bg-gray-50',
                ]"
              >
                {{ totalPages }}
              </button>

              <button
                @click="nextPage"
                :disabled="currentPage === totalPages"
                class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
                :class="{ 'opacity-50 cursor-not-allowed': currentPage === totalPages }"
              >
                <font-awesome-icon icon="chevron-right" class="h-3 w-3" />
              </button>
            </nav>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
