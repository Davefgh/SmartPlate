
<script setup>
import { ref, computed } from 'vue'

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
    owner: 'Stanleigh Morales'
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
    owner: 'Stanleigh Morales'
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
    owner: 'Stanleigh Morales'
  }
])

// Search functionality
const searchQuery = ref('')
const filteredVehicles = computed(() => {
  if (!searchQuery.value) return vehicles.value

  const query = searchQuery.value.toLowerCase()
  return vehicles.value.filter(vehicle =>
    vehicle.make.toLowerCase().includes(query) ||
    vehicle.model.toLowerCase().includes(query) ||
    vehicle.plateNumber.toLowerCase().includes(query)
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
  status: 'Pending'
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
      status: 'Pending'
    }
  }
}

const addVehicle = () => {
  vehicles.value.push({
    id: vehicles.value.length + 1,
    ...newVehicle.value,
    lastUpdated: new Date().toISOString().split('T')[0],
    owner: 'Stanleigh Morales'
  })
  toggleAddVehicleModal()
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
          <button class="bg-gray-100 hover:bg-gray-200 text-gray-700 px-3 py-2 rounded-lg flex items-center transition-colors">
            <font-awesome-icon :icon="['fas', 'filter']" class="mr-2 text-gray-500" />
            Filter
          </button>
          <button class="bg-gray-100 hover:bg-gray-200 text-gray-700 px-3 py-2 rounded-lg flex items-center transition-colors">
            <font-awesome-icon :icon="['fas', 'sort']" class="mr-2 text-gray-500" />
            Sort
          </button>
        </div>
      </div>
    </div>

    <!-- Vehicles List -->
    <div class="bg-white rounded-lg shadow-sm overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Vehicle
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Plate Number
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Status
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Last Updated
              </th>
              <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="vehicle in filteredVehicles" :key="vehicle.id" class="hover:bg-gray-50 transition-colors">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="flex items-center">
                  <div class="flex-shrink-0 h-10 w-10 bg-gray-100 rounded-full flex items-center justify-center">
                    <font-awesome-icon :icon="['fas', 'car']" class="text-gray-500" />
                  </div>
                  <div class="ml-4">
                    <div class="text-sm font-medium text-gray-900">{{ vehicle.make }} {{ vehicle.model }}</div>
                    <div class="text-sm text-gray-500">{{ vehicle.year }} • {{ vehicle.color }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm text-gray-900">{{ vehicle.plateNumber }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span
                  :class="[
                    'px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full',
                    vehicle.status === 'Active' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'
                  ]"
                >
                  {{ vehicle.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ vehicle.lastUpdated }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 mr-3">
                  <font-awesome-icon :icon="['fas', 'edit']" />
                </button>
                <button class="text-red hover:text-red-700">
                  <font-awesome-icon :icon="['fas', 'trash-alt']" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Empty State -->
      <div v-if="filteredVehicles.length === 0" class="py-12 flex flex-col items-center justify-center">
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
      <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Background overlay -->
        <div class="fixed inset-0 transition-opacity" @click="toggleAddVehicleModal">
          <div class="absolute inset-0 bg-gray-500 opacity-75"></div>
        </div>

        <!-- Modal panel -->
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
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
