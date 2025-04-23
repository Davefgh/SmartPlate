<script setup lang="ts">
import type { VehicleRegistrationForm } from '@/types/vehicleRegistration'
import { computed, ref, onMounted, watch } from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

const props = defineProps<{
  registrations: VehicleRegistrationForm[]
}>()

const emit = defineEmits<{
  (e: 'inspectionProcess', registration: VehicleRegistrationForm): void
  (e: 'paymentProcess', registration: VehicleRegistrationForm): void
}>()

const sortBy = ref('submissionDate')
const sortOrder = ref('desc')
const isLoading = ref(true)

// Search and filter
const searchQuery = ref('')
const filterOptions = ref({
  inspectionStatus: '',
  paymentStatus: '',
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
    const matchesInspectionStatus =
      filterOptions.value.inspectionStatus === '' ||
      reg.inspectionStatus === filterOptions.value.inspectionStatus

    const matchesPaymentStatus =
      filterOptions.value.paymentStatus === '' ||
      reg.paymentStatus === filterOptions.value.paymentStatus

    return matchesSearch && matchesInspectionStatus && matchesPaymentStatus
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
    inspectionStatus: '',
    paymentStatus: '',
  }
}

// Simulate loading delay
onMounted(() => {
  isLoading.value = true
  setTimeout(() => {
    isLoading.value = false
  }, 1500)
})
</script>

<template>
  <div class="bg-white rounded-lg shadow overflow-hidden mb-8">
    <!-- Loading State -->
    <div v-if="isLoading" class="p-12 flex flex-col items-center justify-center">
      <div class="flex justify-center items-center mb-4">
        <div class="relative w-20 h-20">
          <div class="absolute top-0 left-0 right-0 bottom-0 w-20 h-20">
            <div
              class="border-4 border-blue-200 border-opacity-50 opacity-75 rounded-full w-20 h-20 animate-spin border-t-blue-600"
            ></div>
          </div>
          <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
            <font-awesome-icon icon="cogs" class="text-blue-600 h-10 w-10" />
          </div>
        </div>
      </div>
      <div class="text-center">
        <p class="text-lg font-medium text-gray-700 mb-1">Processing Registrations</p>
        <p class="text-sm text-gray-500">Fetching data from the server...</p>
      </div>
      <div class="mt-6 w-40 h-2 bg-gray-200 rounded-full overflow-hidden">
        <div class="h-full bg-blue-600 rounded-full animate-progress"></div>
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
              <label class="text-sm text-gray-600">Inspection:</label>
              <select
                v-model="filterOptions.inspectionStatus"
                class="block w-40 pl-3 pr-10 py-2 text-sm border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">All Statuses</option>
                <option value="pending">Pending</option>
                <option value="approved">Approved</option>
                <option value="rejected">Rejected</option>
              </select>
            </div>

            <div class="flex items-center space-x-2">
              <label class="text-sm text-gray-600">Payment:</label>
              <select
                v-model="filterOptions.paymentStatus"
                class="block w-40 pl-3 pr-10 py-2 text-sm border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500"
              >
                <option value="">All Statuses</option>
                <option value="pending">Pending</option>
                <option value="approved">Approved</option>
                <option value="rejected">Rejected</option>
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
        <table class="min-w-full divide-y divide-gray-200 border-separate border-spacing-0">
          <thead class="bg-gray-50 shadow-sm">
            <tr class="divide-x divide-gray-200"></tr>
            <tr>
              <th
                v-for="header in [
                  'Registration ID',
                  'Vehicle Details',
                  'Inspection Status',
                  'Payment Status',
                  'Actions',
                ]"
                :key="header"
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                {{ header }}
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
              <td colspan="5" class="px-6 py-16 text-center">
                <div class="flex flex-col items-center max-w-sm mx-auto">
                  <div class="rounded-full bg-blue-50 p-4 mb-5 ring-4 ring-blue-50">
                    <font-awesome-icon icon="cogs" class="h-8 w-8 text-blue-500" />
                  </div>
                  <h3 class="text-lg font-semibold text-gray-900 mb-2">
                    No Processing Registrations
                  </h3>
                  <p class="text-sm text-gray-600 text-center mb-1 max-w-xs">
                    There are no vehicle registrations currently in the processing stage.
                  </p>
                  <p class="text-sm text-gray-500 text-center max-w-xs">
                    Approved registrations will be displayed here for inspection and payment
                    processing.
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
                <span
                  class="px-2 py-1 text-xs font-medium rounded-full"
                  :class="{
                    'bg-yellow-100 text-yellow-800': registration.inspectionStatus === 'pending',
                    'bg-green-100 text-green-800': registration.inspectionStatus === 'approved',
                    'bg-red-100 text-red-800': registration.inspectionStatus === 'rejected',
                  }"
                >
                  {{ registration.inspectionStatus.toUpperCase() }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <span
                  class="px-2 py-1 text-xs font-medium rounded-full"
                  :class="{
                    'bg-yellow-100 text-yellow-800': registration.paymentStatus === 'pending',
                    'bg-green-100 text-green-800': registration.paymentStatus === 'approved',
                    'bg-red-100 text-red-800': registration.paymentStatus === 'rejected',
                  }"
                >
                  {{ registration.paymentStatus.toUpperCase() }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
                <div v-if="registration.inspectionStatus === 'pending'" class="mb-2">
                  <button
                    @click="emit('inspectionProcess', registration)"
                    class="inline-flex items-center px-2 py-1 border border-transparent text-xs font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none"
                  >
                    Process Inspection
                  </button>
                </div>

                <div
                  v-if="
                    registration.paymentStatus === 'pending' &&
                    registration.inspectionStatus === 'approved'
                  "
                >
                  <button
                    @click="emit('paymentProcess', registration)"
                    class="inline-flex items-center px-2 py-1 border border-transparent text-xs font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none"
                  >
                    Process Payment
                  </button>
                </div>
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

<style scoped>
@keyframes progress {
  0% {
    width: 0%;
  }
  50% {
    width: 70%;
  }
  100% {
    width: 100%;
  }
}

.animate-progress {
  animation: progress 2s ease-in-out infinite;
}
</style>
