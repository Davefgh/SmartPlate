<script setup>
import { ref, computed } from 'vue'
import VehicleModal from './VehicleModal.vue'

// Mock data for vehicles
const vehicles = ref([
  {
    id: 1,
    make: 'Toyota',
    model: 'Corolla',
    year: 2022,
    plateNumber: 'ABC-123',
    color: 'White',
    status: 'Active',
    lastUpdated: '2025-02-15',
    owner: 'Stanleigh Morales',
  },
  {
    id: 2,
    make: 'Honda',
    model: 'Civic',
    year: 2021,
    plateNumber: 'XYZ-789',
    color: 'Black',
    status: 'Active',
    lastUpdated: '2025-01-20',
    owner: 'Stanleigh Morales',
  },
  {
    id: 3,
    make: 'Ford',
    model: 'Mustang',
    year: 2023,
    plateNumber: 'DEF-456',
    color: 'Red',
    status: 'Pending',
    lastUpdated: '2025-03-01',
    owner: 'Stanleigh Morales',
  },
])

// Search functionality
const searchQuery = ref('')
const filteredVehicles = computed(() => {
  if (!searchQuery.value) return vehicles.value

  const query = searchQuery.value.toLowerCase()
  return vehicles.value.filter(
    (vehicle) =>
      vehicle.make.toLowerCase().includes(query) ||
      vehicle.model.toLowerCase().includes(query) ||
      vehicle.plateNumber.toLowerCase().includes(query),
  )
})

// Add new vehicle modal
const isAddVehicleModalOpen = ref(false)
const newVehicle = ref({
  make: '',
  model: '',
  year: new Date().getFullYear(),
  plateNumber: '',
  color: '',
  status: 'Pending',
})

const toggleAddVehicleModal = () => {
  isAddVehicleModalOpen.value = !isAddVehicleModalOpen.value
  if (!isAddVehicleModalOpen.value) {
    newVehicle.value = {
      make: '',
      model: '',
      year: new Date().getFullYear(),
      plateNumber: '',
      color: '',
      status: 'Pending',
    }
  }
}

const addVehicle = () => {
  vehicles.value.push({
    id: vehicles.value.length + 1,
    ...newVehicle.value,
    lastUpdated: new Date().toISOString().split('T')[0],
    owner: 'Stanleigh Morales',
  })
  toggleAddVehicleModal()
}

// Vehicle details modal
const isVehicleModalOpen = ref(false)
const selectedVehicle = ref(null)

const openVehicleModal = (vehicle) => {
  selectedVehicle.value = vehicle
  isVehicleModalOpen.value = true
}

const closeVehicleModal = () => {
  isVehicleModalOpen.value = false
}
</script>

<template>
  <div class="p-6 fade-in">
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
      <div>
        <h1 class="text-2xl font-bold text-gray-800">Vehicle Information</h1>
        <p class="text-gray-600 mt-1">Manage your registered vehicles</p>
      </div>
      <div class="mt-4 md:mt-0">
        <button
          @click="toggleAddVehicleModal"
          class="bg-dark-blue hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg flex items-center transition-colors duration-200"
        >
          <font-awesome-icon :icon="['fas', 'plus']" class="mr-2" />
          Add New Vehicle
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
            placeholder="Search vehicles by make, model, or plate number..."
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

    <!-- Vehicles Grid Layout -->
    <div class="mb-6">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <!-- Vehicle Card -->
        <div
          v-for="vehicle in filteredVehicles"
          :key="vehicle.id"
          class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300"
        >
          <!-- Vehicle Image -->
          <div
            :class="[
              'h-48 relative overflow-hidden',
              vehicle.status === 'Active'
                ? 'bg-gradient-to-r from-blue-50 to-green-50'
                : 'bg-gradient-to-r from-amber-50 to-yellow-50',
            ]"
          >
            <div class="absolute inset-0 flex items-center justify-center">
              <font-awesome-icon
                :icon="['fas', 'car']"
                :class="[
                  'w-24 h-24 opacity-80',
                  vehicle.color.toLowerCase() === 'white'
                    ? 'text-gray-300'
                    : vehicle.color.toLowerCase() === 'black'
                      ? 'text-gray-700'
                      : vehicle.color.toLowerCase() === 'red'
                        ? 'text-red-400'
                        : vehicle.color.toLowerCase() === 'blue'
                          ? 'text-blue-400'
                          : vehicle.color.toLowerCase() === 'green'
                            ? 'text-green-400'
                            : 'text-gray-300',
                ]"
              />
            </div>
            <!-- Car Make Logo Overlay -->
            <div class="absolute bottom-3 right-3 bg-white/90 rounded-full p-2 shadow-sm">
              <font-awesome-icon :icon="['fas', 'car']" class="text-dark-blue w-5 h-5" />
            </div>
            <!-- Status Badge -->
            <div class="absolute top-3 right-3">
              <span
                :class="[
                  'px-3 py-1 inline-flex text-xs leading-5 font-semibold rounded-full shadow-sm',
                  vehicle.status === 'Active'
                    ? 'bg-gradient-to-r from-green-100 to-green-200 text-green-800'
                    : 'bg-gradient-to-r from-yellow-100 to-amber-200 text-amber-800',
                ]"
              >
                {{ vehicle.status }}
              </span>
            </div>
          </div>

          <!-- Vehicle Details -->
          <div class="p-5">
            <div class="flex justify-between items-start mb-3">
              <h3 class="text-lg font-bold text-gray-900">
                {{ vehicle.make }} {{ vehicle.model }}
              </h3>
              <span class="text-sm font-medium text-gray-600">{{ vehicle.year }}</span>
            </div>

            <div class="space-y-2 mb-4">
              <!-- Plate Number -->
              <div class="flex items-center">
                <div class="w-8 flex-shrink-0 text-gray-400">
                  <font-awesome-icon :icon="['fas', 'id-card']" />
                </div>
                <span class="text-sm text-gray-700">{{ vehicle.plateNumber }}</span>
              </div>

              <!-- Color -->
              <div class="flex items-center">
                <div class="w-8 flex-shrink-0 text-gray-400">
                  <font-awesome-icon :icon="['fas', 'palette']" />
                </div>
                <span class="text-sm text-gray-700">{{ vehicle.color }}</span>
              </div>

              <!-- Last Updated -->
              <div class="flex items-center">
                <div class="w-8 flex-shrink-0 text-gray-400">
                  <font-awesome-icon :icon="['fas', 'calendar-alt']" />
                </div>
                <span class="text-sm text-gray-700">Updated: {{ vehicle.lastUpdated }}</span>
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="flex justify-center items-center pt-3 border-t border-gray-100">
              <button 
                @click="openVehicleModal(vehicle)" 
                class="text-sm text-dark-blue hover:text-blue-700 flex items-center py-2 px-4 rounded-lg hover:bg-blue-50 transition-colors duration-200"
              >
                <font-awesome-icon :icon="['fas', 'eye']" class="mr-2" />
                View Details
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div
        v-if="filteredVehicles.length === 0"
        class="py-12 flex flex-col items-center justify-center"
      >
        <div class="bg-gray-100 rounded-full p-4 mb-4">
          <font-awesome-icon :icon="['fas', 'car']" class="text-gray-400 w-8 h-8" />
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-1">No vehicles found</h3>
        <p class="text-gray-500 mb-4">Try adjusting your search or filter criteria</p>
        <button
          @click="toggleAddVehicleModal"
          class="bg-dark-blue hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg flex items-center transition-colors duration-200"
        >
          <font-awesome-icon :icon="['fas', 'plus']" class="mr-2" />
          Add New Vehicle
        </button>
      </div>
    </div>

    <!-- Add Vehicle Modal -->
    <div v-if="isAddVehicleModalOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div
        class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0"
      >
        <!-- Background overlay -->
        <div class="fixed inset-0 transition-opacity" @click="toggleAddVehicleModal">
          <div class="absolute inset-0 bg-gray-500 opacity-75"></div>
        </div>

        <!-- Modal panel -->
        <div
          class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
        >
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left w-full">
                <h3 class="text-lg leading-6 font-medium text-gray-900 mb-4">Add New Vehicle</h3>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">Make</label>
                    <input
                      v-model="newVehicle.make"
                      type="text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">Model</label>
                    <input
                      v-model="newVehicle.model"
                      type="text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">Year</label>
                    <input
                      v-model="newVehicle.year"
                      type="number"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">Color</label>
                    <input
                      v-model="newVehicle.color"
                      type="text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">Plate Number</label>
                    <input
                      v-model="newVehicle.plateNumber"
                      type="text"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    />
                  </div>
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1">Status</label>
                    <select
                      v-model="newVehicle.status"
                      class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                    >
                      <option value="Pending">Pending</option>
                      <option value="Active">Active</option>
                    </select>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button
              @click="addVehicle"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-dark-blue text-base font-medium text-white hover:bg-blue-700 focus:outline-none sm:ml-3 sm:w-auto sm:text-sm"
            >
              Add Vehicle
            </button>
            <button
              @click="toggleAddVehicleModal"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Vehicle Details Modal -->
    <VehicleModal 
      :vehicle="selectedVehicle" 
      :isOpen="isVehicleModalOpen && selectedVehicle" 
      @close="closeVehicleModal" 
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
