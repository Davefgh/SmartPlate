<script setup>
import { ref, computed, defineAsyncComponent } from 'vue'
const RegistrationModal = defineAsyncComponent(() => import('./modals/RegistrationModal.vue'))

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
      total: 2120,
    },
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
      total: 1670,
    },
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
      total: 2120,
    },
  },
])

// Filter states
const statusFilter = ref('All')
const typeFilter = ref('All')

// Filtered registrations
const filteredRegistrations = computed(() => {
  let result = registrations.value

  if (statusFilter.value !== 'All') {
    result = result.filter((reg) => reg.status === statusFilter.value)
  }

  if (typeFilter.value !== 'All') {
    result = result.filter((reg) => reg.registrationType === typeFilter.value)
  }

  return result
})

// Get status color based on status
const getStatusColor = (status) => {
  switch (status) {
    case 'Approved':
      return 'bg-gradient-to-r from-green-100 to-green-200 text-green-800'
    case 'Pending':
      return 'bg-gradient-to-r from-yellow-100 to-amber-200 text-amber-800'
    case 'Rejected':
      return 'bg-gradient-to-r from-red-100 to-red-200 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

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
        <div
          v-if="filteredRegistrations.length === 0"
          class="py-8 flex flex-col items-center justify-center"
        >
          <div class="bg-gray-100 rounded-full p-4 mb-4">
            <font-awesome-icon :icon="['fas', 'file-contract']" class="text-gray-400 w-8 h-8" />
          </div>
          <h3 class="text-lg font-medium text-gray-900 mb-1">No registrations found</h3>
          <p class="text-gray-500">Try adjusting your filters</p>
        </div>

        <div v-else class="relative">
          <!-- Timeline Line -->
          <div class="absolute top-0 bottom-0 left-6 md:left-8 w-0.5 bg-gray-200"></div>

          <!-- Timeline Items -->
          <div
            v-for="registration in filteredRegistrations"
            :key="registration.id"
            class="relative pl-14 md:pl-20 pb-8"
          >
            <!-- Timeline Dot -->
            <div
              :class="[
                'absolute left-4 md:left-6 w-5 h-5 rounded-full border-4 flex items-center justify-center transform transition-all duration-300 hover:scale-110 timeline-dot',
                {
                  'bg-gradient-to-br from-green-400 to-green-600 border-green-100 pulse-green':
                    registration.status === 'Approved',
                  'bg-gradient-to-br from-yellow-400 to-amber-500 border-amber-100 pulse-amber':
                    registration.status === 'Pending',
                  'bg-gradient-to-br from-red-400 to-red-600 border-red-100 pulse-red':
                    registration.status === 'Rejected',
                },
              ]"
            ></div>

            <!-- Registration Card -->
            <div
              class="bg-white rounded-lg shadow-sm border border-gray-100 overflow-hidden hover:shadow-md transition-shadow duration-300"
            >
              <!-- Card Header -->
              <div :class="['p-4 border-b border-gray-100', getStatusColor(registration.status)]">
                <div class="flex justify-between items-center">
                  <h3 class="text-base font-semibold text-gray-800">
                    {{ registration.vehicleInfo }}
                  </h3>
                  <span
                    :class="[
                      'px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full shadow-sm',
                      getStatusBadgeClass(registration.status),
                    ]"
                  >
                    {{ registration.status }}
                  </span>
                </div>
                <p class="text-sm text-gray-600 mt-1">
                  Plate: {{ registration.plateNumber }} | Type: {{ registration.registrationType }}
                </p>
              </div>

              <!-- Card Content -->
              <div class="p-4">
                <div class="grid grid-cols-2 gap-3">
                  <div>
                    <p class="text-xs text-gray-500">Submission Date</p>
                    <p class="text-sm font-medium">{{ registration.submissionDate }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-gray-500">Expiry Date</p>
                    <p class="text-sm font-medium">{{ registration.expiryDate }}</p>
                  </div>
                </div>

                <div class="mt-4 flex justify-end">
                  <button
                    @click="viewRegistrationDetails(registration)"
                    class="text-dark-blue hover:text-blue-700 text-sm font-medium flex items-center transition-colors"
                  >
                    View Details
                    <font-awesome-icon :icon="['fas', 'chevron-right']" class="ml-1 w-3 h-3" />
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Registration Modal Component -->
    <RegistrationModal
      :registration="activeRegistration"
      :is-open="showRegistrationDetail"
      @close="closeRegistrationDetails"
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
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.timeline-dot::before {
  content: '';
  position: absolute;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  z-index: -1;
}

.pulse-green::before {
  animation: pulse-green 2s infinite;
  background-color: rgba(74, 222, 128, 0.6);
}

.pulse-amber::before {
  animation: pulse-amber 2s infinite;
  background-color: rgba(251, 191, 36, 0.6);
}

.pulse-red::before {
  animation: pulse-red 2s infinite;
  background-color: rgba(248, 113, 113, 0.6);
}

@keyframes pulse-green {
  0% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(74, 222, 128, 0.6);
  }
  70% {
    transform: scale(1);
    box-shadow: 0 0 0 10px rgba(74, 222, 128, 0);
  }
  100% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(74, 222, 128, 0);
  }
}

@keyframes pulse-amber {
  0% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(251, 191, 36, 0.6);
  }
  70% {
    transform: scale(1);
    box-shadow: 0 0 0 10px rgba(251, 191, 36, 0);
  }
  100% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(251, 191, 36, 0);
  }
}

@keyframes pulse-red {
  0% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(248, 113, 113, 0.6);
  }
  70% {
    transform: scale(1);
    box-shadow: 0 0 0 10px rgba(248, 113, 113, 0);
  }
  100% {
    transform: scale(0.95);
    box-shadow: 0 0 0 0 rgba(248, 113, 113, 0);
  }
}
</style>
