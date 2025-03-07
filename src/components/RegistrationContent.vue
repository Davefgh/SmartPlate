
<script setup>
import { ref, computed } from 'vue'

// Mock data for registrations
const registrations = ref([
  {
    id: 1,
    vehicleInfo: 'Toyota Corolla 2022',
    plateNumber: 'ABC-123',
    registrationType: 'New Registration',
    submissionDate: '2024-10-05',
    expiryDate: '2025-10-05',
    status: 'Approved',
    documents: ['Proof of Ownership', 'Insurance Certificate', 'Emission Test'],
    fees: {
      registrationFee: 1500,
      plateIssuanceFee: 450,
      computerFee: 170,
      total: 2120
    }
  },
  {
    id: 2,
    vehicleInfo: 'Honda Civic 2021',
    plateNumber: 'XYZ-789',
    registrationType: 'Renewal',
    submissionDate: '2024-08-10',
    expiryDate: '2025-08-10',
    status: 'Approved',
    documents: ['Insurance Certificate', 'Emission Test'],
    fees: {
      registrationFee: 1500,
      computerFee: 170,
      total: 1670
    }
  },
  {
    id: 3,
    vehicleInfo: 'Ford Mustang 2023',
    plateNumber: 'DEF-456',
    registrationType: 'New Registration',
    submissionDate: '2025-01-05',
    expiryDate: '2026-01-05',
    status: 'Pending',
    documents: ['Proof of Ownership', 'Insurance Certificate', 'Emission Test'],
    fees: {
      registrationFee: 1500,
      plateIssuanceFee: 450,
      computerFee: 170,
      total: 2120
    }
  }
])

// Filter states
const statusFilter = ref('All')
const typeFilter = ref('All')

// Filtered registrations
const filteredRegistrations = computed(() => {
  let result = registrations.value

  // Apply status filter
  if (statusFilter.value !== 'All') {
    result = result.filter(reg => reg.status === statusFilter.value)
  }

  // Apply type filter
  if (typeFilter.value !== 'All') {
    result = result.filter(reg => reg.registrationType === typeFilter.value)
  }

  return result
})

// Active registration detail
const activeRegistration = ref(null)
const showRegistrationDetail = ref(false)

const viewRegistrationDetails = (registration) => {
  activeRegistration.value = registration
  showRegistrationDetail.value = true
}

const closeRegistrationDetails = () => {
  showRegistrationDetail.value = false
  setTimeout(() => {
    activeRegistration.value = null
  }, 300)
}

// Get status badge class
const getStatusBadgeClass = (status) => {
  switch (status) {
    case 'Approved':
      return 'bg-green-100 text-green-800'
    case 'Pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'Rejected':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}
</script>

<template>
  <div class="p-6 fade-in">
    <!-- Page Header -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
      <div>
        <h1 class="text-2xl font-bold text-gray-800">Registration Management</h1>
        <p class="text-gray-600 mt-1">Track and manage your vehicle registrations</p>
      </div>
      <div class="mt-4 md:mt-0">
        <button 
          class="bg-dark-blue hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg flex items-center transition-colors duration-200"
        >
          <font-awesome-icon :icon="['fas', 'plus']" class="mr-2" />
          New Registration
        </button>
      </div>
    </div>

    <!-- Filters -->
    <div class="bg-white rounded-lg shadow-sm p-4 mb-6">
      <div class="flex flex-col md:flex-row md:items-center gap-4">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Status</label>
          <select 
            v-model="statusFilter"
            class="w-full md:w-40 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="All">All Status</option>
            <option value="Approved">Approved</option>
            <option value="Pending">Pending</option>
            <option value="Rejected">Rejected</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Type</label>
          <select 
            v-model="typeFilter"
            class="w-full md:w-40 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
          >
            <option value="All">All Types</option>
            <option value="New Registration">New Registration</option>
            <option value="Renewal">Renewal</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Registration Timeline -->
    <div class="bg-white rounded-lg shadow-sm overflow-hidden">
      <div class="p-4 border-b border-gray-100">
        <h2 class="text-lg font-semibold text-gray-800">Registration History</h2>
      </div>
      
      <div class="p-4">
        <div v-if="filteredRegistrations.length === 0" class="py-8 flex flex-col items-center justify-center">
          <div class="bg-gray-100 rounded-full p-4 mb-4">
            <font-awesome-icon :icon="['fas', 'file-contract']" class="text-gray-400 w-8 h-8" />
          </div>
          <h3 class="text-lg font-medium text-gray-900 mb-1">No registrations found</h3>
          <p class="text-gray-500">Try adjusting your filters or create a new registration</p>
        </div>
        
        <div v-else class="relative">
          <!-- Timeline Line -->
          <div class="absolute top-0 bottom-0 left-6 md:left-8 w-0.5 bg-gray-200"></div>
          
          <!-- Timeline Items -->
          <div v-for="registration in filteredRegistrations" :key="registration.id" class="relative pl-14 md:pl-20 pb-8">
            <!-- Timeline Dot -->
            <div class="absolute left-4 md:left-6 w-4 h-4 rounded-full bg-dark-blue border-4 border-blue-100"></div>
            
            <!-- Timeline Content -->
            <div class="bg-gray-50 rounded-lg p-4 hover:shadow-md transition-shadow duration-200">
              <div class="flex flex-col md:flex-row md:items-center md:justify-between mb-3">
                <div>
                  <h3 class="text-md font-semibold text-gray-800">{{ registration.vehicleInfo }}</h3>
                  <p class="text-sm text-gray-600">Plate Number: {{ registration.plateNumber }}</p>
                </div>
                <span 
                  :class="[
                    'px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full mt-2 md:mt-0',
                    getStatusBadgeClass(registration.status)
                  ]"
                >
                  {{ registration.status }}
                </span>
              </div>
              
              <div class="grid grid-cols-1 md:grid-cols-3 gap-2 mb-3">
                <div>
                  <span class="text-xs text-gray-500 block">Registration Type</span>
                  <span class="text-sm font-medium">{{ registration.registrationType }}</span>
                </div>
                <div>
                  <span class="text-xs text-gray-500 block">Submission Date</span>
                  <span class="text-sm font-medium">{{ registration.submissionDate }}</span>
                </div>
                <div>
                  <span class="text-xs text-gray-500 block">Expiry Date</span>
                  <span class="text-sm font-medium">{{ registration.expiryDate }}</span>
                </div>
              </div>
              
              <div class="flex justify-end">
                <button 
                  @click="viewRegistrationDetails(registration)"
                  class="text-blue-600 hover:text-blue-800 text-sm font-medium transition-colors"
                >
                  <font-awesome-icon :icon="['fas', 'eye']" class="mr-1" />
                  View Details
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Registration Detail Modal -->
    <div v-if="showRegistrationDetail" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Background overlay -->
        <div class="fixed inset-0 transition-opacity" @click="closeRegistrationDetails">
          <div class="absolute inset-0 bg-gray-500 opacity-75"></div>
        </div>

        <!-- Modal panel -->
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left w-full">
                <div class="flex justify-between items-center mb-4">
                  <h3 class="text-lg leading-6 font-medium text-gray-900">Registration Details</h3>
                  <button @click="closeRegistrationDetails" class="text-gray-400 hover:text-gray-500">
                    <font-awesome-icon :icon="['fas', 'times']" class="w-5 h-5" />
                  </button>
                </div>
                
                <div v-if="activeRegistration" class="space-y-4">
                  <!-- Vehicle Info -->
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <h4 class="text-sm font-semibold text-gray-700 mb-2">Vehicle Information</h4>
                    <div class="flex justify-between mb-1">
                      <span class="text-sm text-gray-500">Vehicle</span>
                      <span class="text-sm font-medium">{{ activeRegistration.vehicleInfo }}</span>
                    </div>
                    <div class="flex justify-between">
                      <span class="text-sm text-gray-500">Plate Number</span>
                      <span class="text-sm font-medium">{{ activeRegistration.plateNumber }}</span>
                    </div>
                  </div>
                  
                  <!-- Registration Info -->
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <h4 class="text-sm font-semibold text-gray-700 mb-2">Registration Information</h4>
                    <div class="flex justify-between mb-1">
                      <span class="text-sm text-gray-500">Type</span>
                      <span class="text-sm font-medium">{{ activeRegistration.registrationType }}</span>
                    </div>
                    <div class="flex justify-between mb-1">
                      <span class="text-sm text-gray-500">Submission Date</span>
                      <span class="text-sm font-medium">{{ activeRegistration.submissionDate }}</span>
                    </div>
                    <div class="flex justify-between mb-1">
                      <span class="text-sm text-gray-500">Expiry Date</span>
                      <span class="text-sm font-medium">{{ activeRegistration.expiryDate }}</span>
                    </div>
                    <div class="flex justify-between">
                      <span class="text-sm text-gray-500">Status</span>
                      <span 
                        :class="[
                          'text-sm font-medium px-2 py-0.5 rounded-full text-xs',
                          getStatusBadgeClass(activeRegistration.status)
                        ]"
                      >
                        {{ activeRegistration.status }}
                      </span>
                    </div>
                  </div>
                  
                  <!-- Documents -->
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <h4 class="text-sm font-semibold text-gray-700 mb-2">Submitted Documents</h4>
                    <ul class="space-y-1">
                      <li v-for="(doc, index) in activeRegistration.documents" :key="index" class="text-sm">
                        <font-awesome-icon :icon="['fas', 'file-alt']" class="text-gray-400 mr-2" />
                        {{ doc }}
                      </li>
                    </ul>
                  </div>
                  
                  <!-- Fees -->
                  <div class="bg-gray-50 p-3 rounded-lg">
                    <h4 class="text-sm font-semibold text-gray-700 mb-2">Fees</h4>
                    <div v-for="(value, key) in activeRegistration.fees" :key="key" class="flex justify-between mb-1" :class="{ 'font-semibold': key === 'total' }">
                      <span class="text-sm text-gray-500">{{ key.charAt(0).toUpperCase() + key.slice(1) }}</span>
                      <span class="text-sm">₱{{ value.toLocaleString() }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button 
              v-if="activeRegistration && activeRegistration.status === 'Approved'"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-dark-blue text-base font-medium text-white hover:bg-blue-700 focus:outline-none sm:ml-3 sm:w-auto sm:text-sm"
            >
              <font-awesome-icon :icon="['fas', 'print']" class="mr-2" />
              Print Certificate
            </button>
            <button 
              @click="closeRegistrationDetails"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Close
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