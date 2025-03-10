<script setup>
import { ref, computed } from 'vue'
import PlateModal from './modals/PlateModal.vue'

// Mock data for plates
const plates = ref([
  {
    id: 1,
    plateNumber: 'ABC-123',
    vehicleMake: 'Toyota',
    vehicleModel: 'Corolla',
    registrationDate: '2024-10-15',
    expiryDate: '2025-10-15',
    status: 'Active',
    owner: 'Stanleigh Morales',
    type: 'Private',
  },
  {
    id: 2,
    plateNumber: 'XYZ-789',
    vehicleMake: 'Honda',
    vehicleModel: 'Civic',
    registrationDate: '2024-08-20',
    expiryDate: '2025-08-20',
    status: 'Active',
    owner: 'Stanleigh Morales',
    type: 'Private',
  },
  {
    id: 3,
    plateNumber: 'DEF-456',
    vehicleMake: 'Ford',
    vehicleModel: 'Mustang',
    registrationDate: '2025-01-05',
    expiryDate: '2026-01-05',
    status: 'Pending',
    owner: 'Stanleigh Morales',
    type: 'Private',
  },
])

// Search functionality
const searchQuery = ref('')
const filteredPlates = computed(() => {
  if (!searchQuery.value) return plates.value

  const query = searchQuery.value.toLowerCase()
  return plates.value.filter(
    (plate) =>
      plate.plateNumber.toLowerCase().includes(query) ||
      plate.vehicleMake.toLowerCase().includes(query) ||
      plate.vehicleModel.toLowerCase().includes(query),
  )
})

// Calculate days remaining until expiry
const getDaysRemaining = (expiryDateStr) => {
  const today = new Date()
  const expiryDate = new Date(expiryDateStr)
  const diffTime = expiryDate - today
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  return diffDays
}

// Get status color based on days remaining
const getExpiryStatusColor = (expiryDateStr) => {
  const daysRemaining = getDaysRemaining(expiryDateStr)

  if (daysRemaining < 0) return 'bg-red-100 text-red-800' // Expired
  if (daysRemaining < 30) return 'bg-yellow-100 text-yellow-800' // Expiring soon
  return 'bg-green-100 text-green-800' // Valid
}

// Get expiry status text
const getExpiryStatusText = (expiryDateStr) => {
  const daysRemaining = getDaysRemaining(expiryDateStr)

  if (daysRemaining < 0) return 'Expired'
  if (daysRemaining < 30) return `Expires in ${daysRemaining} days`
  return 'Valid'
}

// Plate details modal
const isPlateModalOpen = ref(false)
const selectedPlate = ref(null)

const openPlateModal = (plate) => {
  selectedPlate.value = plate
  isPlateModalOpen.value = true
}

const closePlateModal = () => {
  isPlateModalOpen.value = false
}
</script>

<template>
  <div class="p-6 fade-in">
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
      <div>
        <h1 class="text-2xl font-bold text-gray-800">Plate Information</h1>
        <p class="text-gray-600 mt-1">Manage your vehicle license plates</p>
      </div>
      <div class="mt-4 md:mt-0">
        <button
          class="bg-dark-blue hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg flex items-center transition-colors duration-200"
        >
          <font-awesome-icon :icon="['fas', 'plus']" class="mr-2" />
          Register New Plate
        </button>
      </div>
    </div>

    <!-- Search and Filter Bar -->
    <div class="bg-white rounded-lg shadow-sm p-4 mb-6">
      <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
        <div class="relative flex-grow">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Search plates by number, vehicle make or model..."
            class="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-200 focus:outline-none focus:ring-2 focus:ring-dark-blue/20 transition-all"
          />
          <div class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400">
            <font-awesome-icon :icon="['fas', 'search']" class="w-4 h-4" />
          </div>
        </div>
        <div class="flex items-center space-x-2">
          <button
            class="bg-gray-100 hover:bg-gray-200 text-gray-700 px-3 py-2 rounded-lg flex items-center transition-colors"
          >
            <font-awesome-icon :icon="['fas', 'filter']" class="mr-2 text-gray-500" />
            Filter
          </button>
          <button
            class="bg-gray-100 hover:bg-gray-200 text-gray-700 px-3 py-2 rounded-lg flex items-center transition-colors"
          >
            <font-awesome-icon :icon="['fas', 'sort']" class="mr-2 text-gray-500" />
            Sort
          </button>
        </div>
      </div>
    </div>

    <!-- Plates Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="plate in filteredPlates"
        :key="plate.id"
        class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300"
      >
        <!-- Plate Header -->
        <div
          :class="[
            'p-4 border-b border-gray-100',
            plate.status === 'Active'
              ? 'bg-gradient-to-r from-blue-50 to-green-50'
              : 'bg-gradient-to-r from-amber-50 to-yellow-50',
          ]"
        >
          <div class="flex justify-between items-center">
            <h3 class="text-lg font-bold text-gray-800">{{ plate.plateNumber }}</h3>
            <span
              :class="[
                'px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full shadow-sm',
                plate.status === 'Active'
                  ? 'bg-gradient-to-r from-green-100 to-green-200 text-green-800'
                  : 'bg-gradient-to-r from-yellow-100 to-amber-200 text-yellow-800',
              ]"
            >
              {{ plate.status }}
            </span>
          </div>
          <p class="text-sm text-gray-600 mt-1">{{ plate.vehicleMake }} {{ plate.vehicleModel }}</p>
        </div>

        <!-- Plate Details -->
        <div class="p-4">
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-sm text-gray-500">Registration Date</span>
              <span class="text-sm font-medium">{{ plate.registrationDate }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-sm text-gray-500">Expiry Date</span>
              <span class="text-sm font-medium">{{ plate.expiryDate }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-sm text-gray-500">Type</span>
              <span class="text-sm font-medium">{{ plate.type }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-sm text-gray-500">Status</span>
              <span
                :class="[
                  'text-sm font-medium px-2 py-1 rounded-full text-xs',
                  getExpiryStatusColor(plate.expiryDate),
                ]"
              >
                {{ getExpiryStatusText(plate.expiryDate) }}
              </span>
            </div>
          </div>
        </div>

        <!-- Plate Actions -->
        <div class="bg-gray-50 p-4 border-t border-gray-100">
          <div class="flex justify-between">
            <button
              @click="openPlateModal(plate)"
              class="text-blue-600 hover:text-blue-800 text-sm font-medium transition-colors flex items-center"
            >
              <font-awesome-icon :icon="['fas', 'eye']" class="mr-1" />
              View Details
            </button>
            <button class="text-gray-600 hover:text-gray-800 text-sm font-medium transition-colors">
              <font-awesome-icon :icon="['fas', 'print']" class="mr-1" />
              Print
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div
      v-if="filteredPlates.length === 0"
      class="bg-white rounded-lg shadow-sm p-12 flex flex-col items-center justify-center"
    >
      <div class="bg-gray-100 rounded-full p-4 mb-4">
        <font-awesome-icon :icon="['fas', 'id-card']" class="text-gray-400 w-8 h-8" />
      </div>
      <h3 class="text-lg font-medium text-gray-900 mb-1">No plates found</h3>
      <p class="text-gray-500 mb-4">Try adjusting your search or filter criteria</p>
      <button
        class="bg-dark-blue hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg flex items-center transition-colors duration-200"
      >
        <font-awesome-icon :icon="['fas', 'plus']" class="mr-2" />
        Register New Plate
      </button>
    </div>

    <!-- Plate Details Modal -->
    <PlateModal
      :plate="selectedPlate"
      :isOpen="isPlateModalOpen && selectedPlate"
      @close="closePlateModal"
    />
  </div>
</template>

<style scoped>
.fade-in {
  animation: fadeIn 0.5s ease-in-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
