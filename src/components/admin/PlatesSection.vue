<script setup lang="ts">
import { ref, computed, defineAsyncComponent } from 'vue'
import { useVehicleRegistrationStore } from '@/stores/vehicleRegistration'
import type { Plate, Vehicle } from '@/types/vehicle'

interface ExtendedVehicle extends Vehicle {
  id: number
}

interface PlateWithVehicle extends Plate {
  registrationDate: string
  expiryDate: string
  plateNumber: string
  type: string
  plateType: string
  vehicle: string
  status: string
  id: number
  vehicleId: number
  plate_issue_date: string
  plate_expiration_date: string
  plate_number: string
  plate_type: string
  [key: string]: string | number
}

interface Filter {
  value: string
  label: string
  active: boolean
}

interface TableHeader {
  text: string
  value: string
  sortable: boolean
}

const PlateDetailsModal = defineAsyncComponent(
  () => import('@/components/modals/PlateDetailsModal.vue'),
)

const vehicleStore = useVehicleRegistrationStore()

// Search and filter state
const searchQuery = ref<string>('')
const sortBy = ref<string>('plateNumber')
const sortDesc = ref<boolean>(false)
const currentPage = ref<number>(1)
const itemsPerPage = ref<number>(10)

// Status filters
const statusFilters = ref<Filter[]>([
  { value: 'all', label: 'All Status', active: true },
  { value: 'Active', label: 'Active', active: false },
  { value: 'Expired', label: 'Expired', active: false },
  { value: 'Pending', label: 'Pending', active: false },
])

// Type filters
const typeFilters = ref<Filter[]>([
  { value: 'all', label: 'All Types', active: true },
  { value: 'Regular', label: 'Regular', active: false },
  { value: 'Special', label: 'Special', active: false },
  { value: 'Vanity', label: 'Vanity', active: false },
])

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
const setStatusFilter = (value: string) => {
  statusFilters.value = statusFilters.value.map((filter) => ({
    ...filter,
    active: filter.value === value,
  }))
}

// Apply type filter
const setTypeFilter = (value: string) => {
  typeFilters.value = typeFilters.value.map((filter) => ({
    ...filter,
    active: filter.value === value,
  }))
}

// Get all plates with vehicle information
const plates = computed((): PlateWithVehicle[] => {
  return vehicleStore.platesWithVehicleInfo.map((plate) => {
    const vehicle = vehicleStore.vehicles.find(
      (v): v is ExtendedVehicle => v.id === plate.vehicleId,
    )

    const vehicleInfo = vehicle
      ? `${vehicle.vehicleMake} ${vehicle.vehicleSeries} ${vehicle.yearModel}`
      : 'N/A'

    return {
      ...plate,
      id: plate.plateId,
      registrationDate: plate.plate_issue_date,
      expiryDate: plate.plate_expiration_date,
      plateNumber: plate.plate_number,
      type: 'Private',
      plateType: plate.plate_type,
      vehicle: vehicleInfo,
      status: plate.status || 'Pending',
    }
  })
})

// Filtered and sorted plates
const filteredPlates = computed((): PlateWithVehicle[] => {
  let result = plates.value

  // Apply status filter
  if (activeStatusFilter.value !== 'all') {
    result = result.filter((plate) => plate.status === activeStatusFilter.value)
  }

  // Apply type filter
  if (activeTypeFilter.value !== 'all') {
    result = result.filter((plate) => plate.type === activeTypeFilter.value)
  }

  // Apply search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    result = result.filter(
      (plate) =>
        plate.plateNumber.toLowerCase().includes(query) ||
        plate.vehicle.toLowerCase().includes(query),
    )
  }

  // Apply sorting
  return result.sort((a, b) => {
    const modifier = sortDesc.value ? -1 : 1
    const sortField = sortBy.value as keyof PlateWithVehicle
    const aValue = a[sortField]
    const bValue = b[sortField]

    // Handle undefined or null values
    if (aValue === undefined || aValue === null) return 1 * modifier
    if (bValue === undefined || bValue === null) return -1 * modifier

    // Handle date fields specially
    if (sortField === 'registrationDate' || sortField === 'expiryDate') {
      const aDate = new Date(String(aValue)).getTime()
      const bDate = new Date(String(bValue)).getTime()
      return isNaN(aDate) || isNaN(bDate) ? 0 : (aDate - bDate) * modifier
    }

    // Handle numeric fields
    const aNum = Number(aValue)
    const bNum = Number(bValue)
    if (!isNaN(aNum) && !isNaN(bNum)) {
      return (aNum - bNum) * modifier
    }

    // Default string comparison
    const aStr = String(aValue).toLowerCase()
    const bStr = String(bValue).toLowerCase()
    return aStr.localeCompare(bStr) * modifier
  })
})

// Pagination
const totalPages = computed(() => Math.ceil(filteredPlates.value.length / itemsPerPage.value))
const paginatedPlates = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredPlates.value.slice(start, end)
})

// Status badge color
const getStatusColor = (status: string) => {
  switch (status.toLowerCase()) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'expired':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// Sort handler
const sort = (header: TableHeader) => {
  if (!header.sortable) return
  if (sortBy.value === header.value) {
    sortDesc.value = !sortDesc.value
  } else {
    sortBy.value = header.value
    sortDesc.value = false
  }
}

// Table headers
// Modal state
const selectedPlate = ref<PlateWithVehicle | null>(null)
const showPlateModal = ref<boolean>(false)

// Modal handlers
const openPlateModal = (plate: PlateWithVehicle & { id: number }) => {
  selectedPlate.value = plate
  showPlateModal.value = true
}

const closePlateModal = () => {
  selectedPlate.value = null
  showPlateModal.value = false
}

const headers: TableHeader[] = [
  { text: 'Plate Number', value: 'plateNumber', sortable: true },
  { text: 'Vehicle', value: 'vehicle', sortable: true },
  { text: 'Registration Date', value: 'registrationDate', sortable: true },
  { text: 'Expiry Date', value: 'expiryDate', sortable: true },
  { text: 'Status', value: 'status', sortable: true },
  { text: 'Type', value: 'type', sortable: true },
  { text: 'Actions', value: 'actions', sortable: false },
]
</script>

<template>
  <div class="p-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-semibold text-gray-800">Plates Management</h2>
    </div>

    <!-- Filters and Search -->
    <div class="bg-white rounded-lg shadow-sm p-4 mb-6">
      <div class="space-y-4">
        <!-- Search Bar -->
        <div class="relative">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search by plate number or vehicle..."
            class="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-200 focus:outline-none focus:ring-2 focus:ring-dark-blue/20 transition-all"
          />
          <div class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400">
            <i class="fas fa-search w-4 h-4"></i>
          </div>
        </div>

        <!-- Status and Type Filters -->
        <div class="flex flex-wrap gap-2">
          <!-- Status Filters -->
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

          <!-- Type Filters -->
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

    <!-- Plates Table -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                v-for="header in headers"
                :key="header.value"
                @click="sort(header)"
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100 transition-colors"
                :class="{ 'cursor-default': !header.sortable }"
              >
                <div class="flex items-center gap-2">
                  {{ header.text }}
                  <span v-if="header.sortable" class="text-gray-400">
                    <font-awesome-icon
                      v-if="sortBy === header.value"
                      :icon="['fas', sortDesc ? 'sort-down' : 'sort-up']"
                    />
                    <font-awesome-icon v-else :icon="['fas', 'sort']" />
                  </span>
                </div>
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr
              v-for="plate in paginatedPlates"
              :key="plate.id"
              class="hover:bg-gray-50 transition-colors"
            >
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ plate.plateNumber }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ plate.vehicle }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ plate.registrationDate }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ plate.expiryDate }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <span
                  :class="[
                    'px-2 py-1 rounded-full text-xs font-medium',
                    getStatusColor(plate.status),
                  ]"
                >
                  {{ plate.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ plate.type }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                <div class="flex items-center gap-3">
                  <button
                    class="text-blue-600 hover:text-blue-900 flex items-center gap-1"
                    @click="openPlateModal(plate)"
                  >
                    <font-awesome-icon :icon="['fas', 'eye']" />
                    View
                  </button>
                  <button
                    class="text-indigo-600 hover:text-indigo-900 flex items-center gap-1"
                    @click="() => {}"
                  >
                    <font-awesome-icon :icon="['fas', 'edit']" />
                    Edit
                  </button>
                  <button
                    class="text-red-600 hover:text-red-900 flex items-center gap-1"
                    @click="() => {}"
                  >
                    <font-awesome-icon :icon="['fas', 'trash']" />
                    Delete
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="px-6 py-4 bg-gray-50 border-t border-gray-200">
        <div class="flex items-center justify-between">
          <div class="text-sm text-gray-700">
            Showing {{ (currentPage - 1) * itemsPerPage + 1 }} to
            {{ Math.min(currentPage * itemsPerPage, filteredPlates.length) }} of
            {{ filteredPlates.length }} entries
          </div>
          <div class="flex items-center gap-2">
            <button
              @click="currentPage--"
              :disabled="currentPage === 1"
              class="px-3 py-1 rounded border border-gray-300 disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-100"
            >
              <font-awesome-icon :icon="['fas', 'chevron-left']" />
            </button>
            <span class="text-sm text-gray-700">Page {{ currentPage }} of {{ totalPages }}</span>
            <button
              @click="currentPage++"
              :disabled="currentPage === totalPages"
              class="px-3 py-1 rounded border border-gray-300 disabled:opacity-50 disabled:cursor-not-allowed hover:bg-gray-100"
            >
              <font-awesome-icon :icon="['fas', 'chevron-right']" />
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Plate Details Modal -->
    <PlateDetailsModal
      v-if="selectedPlate"
      :show="showPlateModal"
      :plate="selectedPlate"
      @close="closePlateModal"
    />
  </div>
</template>
