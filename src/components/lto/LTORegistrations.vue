<script setup lang="ts">
import type { VehicleRegistrationForm, AdditionalVehicleData } from '@/types/vehicleRegistration'

interface RegistrationForm extends VehicleRegistrationForm {}

import { useVehicleRegistrationFormStore } from '@/stores/vehicleRegistrationForm'
import { computed, ref, defineAsyncComponent } from 'vue'

// Lazy load the modal components
const VehicleInspectionModal = defineAsyncComponent({
  loader: () => import('../modals/VehicleInspectionModal.vue'),
  // Add error handling
  onError(error, retry, fail, attempts) {
    console.error('Error loading VehicleInspectionModal:', error)
    if (attempts <= 3) {
      // Retry a few times
      retry()
    } else {
      // Show error in console but fail gracefully
      console.error('Failed to load VehicleInspectionModal after multiple attempts')
      fail()
    }
  },
})

const PaymentProcessModal = defineAsyncComponent({
  loader: () => import('../modals/PaymentProcessModal.vue'),
  // Add error handling
  onError(error, retry, fail, attempts) {
    console.error('Error loading PaymentProcessModal:', error)
    if (attempts <= 3) {
      retry()
    } else {
      console.error('Failed to load PaymentProcessModal after multiple attempts')
      fail()
    }
  },
})

const registrationFormStore = useVehicleRegistrationFormStore()
const sortBy = ref('submissionDate')
const sortOrder = ref('desc')

// Active registration for modal
const activeRegistration = ref<RegistrationForm | null>(null)
const showInspectionModal = ref(false)
const showPaymentModal = ref(false)

// Get pending registrations that haven't been approved yet
const pendingRegistrations = computed<RegistrationForm[]>(() =>
  registrationFormStore.forms.filter((form) => form.status === 'pending'),
)

// Get registrations that are approved but still in processing
const processingRegistrations = computed<RegistrationForm[]>(() =>
  registrationFormStore.forms.filter(
    (form) =>
      form.status === 'approved' &&
      (form.inspectionStatus !== 'approved' || form.paymentStatus !== 'approved'),
  ),
)

// Get completed registrations
const completedRegistrations = computed<RegistrationForm[]>(() =>
  registrationFormStore.forms.filter(
    (form) =>
      form.status === 'approved' &&
      form.inspectionStatus === 'approved' &&
      form.paymentStatus === 'approved',
  ),
)

// Sort pending registrations
const sortedPendingRegistrations = computed<RegistrationForm[]>(() => {
  return [...pendingRegistrations.value].sort((a, b) => {
    const aValue = (a as any)[sortBy.value] || ''
    const bValue = (b as any)[sortBy.value] || ''
    const order = sortOrder.value === 'asc' ? 1 : -1
    return aValue > bValue ? order : -order
  })
})

// Sort processing registrations
const sortedProcessingRegistrations = computed<RegistrationForm[]>(() => {
  return [...processingRegistrations.value].sort((a, b) => {
    const aValue = (a as any)[sortBy.value] || ''
    const bValue = (b as any)[sortBy.value] || ''
    const order = sortOrder.value === 'asc' ? 1 : -1
    return aValue > bValue ? order : -order
  })
})

// Sort completed registrations
const sortedCompletedRegistrations = computed<RegistrationForm[]>(() => {
  return [...completedRegistrations.value].sort((a, b) => {
    const aValue = (a as any)[sortBy.value] || ''
    const bValue = (b as any)[sortBy.value] || ''
    const order = sortOrder.value === 'asc' ? 1 : -1
    return aValue > bValue ? order : -order
  })
})

const toggleSort = (field: 'id' | 'vehicleType' | 'applicantName' | 'submissionDate') => {
  if (sortBy.value === field) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortBy.value = field
    sortOrder.value = 'desc'
  }
}

// Process a registration - approve, reject, or update status
const processRegistration = async (registrationId: string, action: 'approve' | 'reject') => {
  await registrationFormStore.$patch((state) => {
    const form = state.forms.find((f) => f.id === registrationId)
    if (form) {
      form.status = action === 'approve' ? 'approved' : 'rejected'

      // If approved, set initial processing state
      if (action === 'approve') {
        form.inspectionStatus = 'pending'
        form.paymentStatus = 'pending'
      }
    }
  })
}

// Open inspection modal for a registration
const openInspectionModal = (registration: RegistrationForm) => {
  activeRegistration.value = JSON.parse(JSON.stringify(registration)) // Deep copy
  showInspectionModal.value = true
}

// Open payment modal for a registration
const openPaymentModal = (registration: RegistrationForm) => {
  activeRegistration.value = JSON.parse(JSON.stringify(registration)) // Deep copy
  showPaymentModal.value = true
}

// Close inspection modal
const closeInspectionModal = () => {
  showInspectionModal.value = false
  activeRegistration.value = null
}

// Close payment modal
const closePaymentModal = () => {
  showPaymentModal.value = false
  activeRegistration.value = null
}

// Handle inspection submission
const handleInspectionSubmit = async (data: {
  id: string
  inspectionStatus: 'approved' | 'rejected'
  inspectionNotes: string
  additionalVehicleData: AdditionalVehicleData
}) => {
  await registrationFormStore.$patch((state) => {
    const form = state.forms.find((f) => f.id === data.id)
    if (form) {
      // Update inspection status
      form.inspectionStatus = data.inspectionStatus

      // Add inspection data to the form
      form.inspectionNotes = data.inspectionNotes

      // Update vehicle data with additional inspection data
      form.additionalVehicleData = data.additionalVehicleData

      // If inspection is rejected, also reject the overall status
      if (data.inspectionStatus === 'rejected') {
        form.status = 'rejected'
      }
    }
  })

  // Close modal
  showInspectionModal.value = false
  activeRegistration.value = null
}

// Handle payment submission
const handlePaymentSubmit = async (data: {
  id: string
  paymentStatus: 'approved' | 'rejected'
  paymentNotes: string
  paymentDetails: {
    amountPaid: number
    paymentDate: string
    paymentMethod: string
    receiptNumber: string
    referenceNumber: string
  }
}) => {
  await registrationFormStore.$patch((state) => {
    const form = state.forms.find((f) => f.id === data.id)
    if (form) {
      // Update payment status
      form.paymentStatus = data.paymentStatus

      // Add payment data to the form
      form.paymentNotes = data.paymentNotes
      form.paymentDetails = data.paymentDetails

      // If payment is approved, update status to 'payment_completed' so it shows up in plate issuance
      if (data.paymentStatus === 'approved') {
        form.status = 'payment_completed' as any

        // Generate payment confirmation code if not present
        if (!(form as any).paymentConfirmationCode) {
          const timestamp = new Date().getTime().toString().slice(-6)
          const random = Math.floor(Math.random() * 10000)
            .toString()
            .padStart(4, '0')
          ;(form as any).paymentConfirmationCode = `CONF-${timestamp}-${random}`
        }
      }
      // If payment is rejected, also reject the overall status
      else if (data.paymentStatus === 'rejected') {
        form.status = 'rejected'
      }
    }
  })

  // Close modal
  showPaymentModal.value = false
  activeRegistration.value = null
}
</script>

<template>
  <div class="p-6">
    <!-- Pending Registrations Table -->
    <h1 class="text-2xl font-semibold text-gray-900 mb-6">Pending Vehicle Registrations</h1>
    <div class="bg-white rounded-lg shadow overflow-hidden mb-8">
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
            <tr v-for="registration in sortedPendingRegistrations" :key="registration.id">
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
            <tr v-if="sortedPendingRegistrations.length === 0">
              <td colspan="5" class="px-6 py-4 text-center text-sm text-gray-500">
                No pending registrations found
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Processing Registrations Table -->
    <h1 class="text-2xl font-semibold text-gray-900 mb-6">Processing Registrations</h1>
    <div class="bg-white rounded-lg shadow overflow-hidden mb-8">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Registration ID
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Vehicle Details
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Inspection Status
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Payment Status
              </th>
              <th
                class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="registration in sortedProcessingRegistrations" :key="registration.id">
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
                <!-- Step 3: Inspection Controls -->
                <div v-if="registration.inspectionStatus === 'pending'" class="mb-2">
                  <button
                    @click="openInspectionModal(registration)"
                    class="inline-flex items-center px-2 py-1 border border-transparent text-xs font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none"
                  >
                    Process Inspection
                  </button>
                </div>

                <!-- Step 4: Payment Controls -->
                <div
                  v-if="
                    registration.paymentStatus === 'pending' &&
                    registration.inspectionStatus === 'approved'
                  "
                >
                  <button
                    @click="openPaymentModal(registration)"
                    class="inline-flex items-center px-2 py-1 border border-transparent text-xs font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none"
                  >
                    Process Payment
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="sortedProcessingRegistrations.length === 0">
              <td colspan="5" class="px-6 py-4 text-center text-sm text-gray-500">
                No registrations in processing
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Completed Registrations Table -->
    <h1 class="text-2xl font-semibold text-gray-900 mb-6">Completed Registrations</h1>
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Registration ID
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Vehicle Details
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Owner
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Completion Date
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Status
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="registration in sortedCompletedRegistrations" :key="registration.id">
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
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <span
                  class="px-2 py-1 text-xs font-medium rounded-full bg-green-100 text-green-800"
                >
                  COMPLETED
                </span>
              </td>
            </tr>
            <tr v-if="sortedCompletedRegistrations.length === 0">
              <td colspan="5" class="px-6 py-4 text-center text-sm text-gray-500">
                No completed registrations
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Lazy-loaded modals -->
    <VehicleInspectionModal
      v-if="showInspectionModal"
      :is-open="showInspectionModal"
      :registration="activeRegistration"
      @close="closeInspectionModal"
      @submit="handleInspectionSubmit"
    />

    <PaymentProcessModal
      v-if="showPaymentModal"
      :is-open="showPaymentModal"
      :registration="activeRegistration"
      @close="closePaymentModal"
      @submit="handlePaymentSubmit"
    />
  </div>
</template>
