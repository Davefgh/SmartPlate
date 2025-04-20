<script setup lang="ts">
import { ref, computed, defineAsyncComponent } from 'vue'
import type { VehicleRegistrationForm } from '@/types/vehicleRegistration'
import type { Plate } from '@/types/vehicle'
import { useVehicleRegistrationFormStore } from '@/stores/vehicleRegistrationForm'

// Lazy load the PlateIssuanceModal component
const PlateIssuanceModal = defineAsyncComponent(
  () => import('@/components/modals/PlateIssuanceModal.vue'),
)

const vehicleRegistrationFormStore = useVehicleRegistrationFormStore()

// Reactive variables
const sortBy = ref('registrationdate')
const sortOrder = ref('desc')
const activeRegistration = ref<VehicleRegistrationForm | null>(null)
const showPlateModal = ref(false)
const suggestedPlateNumber = ref('')

// Toggle sort order
const toggleSort = (field: string) => {
  if (sortBy.value === field) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortBy.value = field
    sortOrder.value = 'asc'
  }
}

// Get pending plate issuance registrations (status = 'payment_completed')
const pendingPlateIssuance = computed(() => {
  return vehicleRegistrationFormStore.forms.filter(
    (registration) => registration.status === ('payment_completed' as any),
  )
})

// Get issued plates (registrations with status = 'completed')
const issuedPlates = computed(() => {
  return vehicleRegistrationFormStore.forms.filter(
    (registration) => registration.status === ('completed' as any),
  )
})

// Sort registrations based on current sort settings
const sortedPendingPlateIssuance = computed(() => {
  return [...pendingPlateIssuance.value].sort((a, b) => {
    let valA: string | number, valB: string | number

    switch (sortBy.value) {
      case 'id':
        valA = a.id
        valB = b.id
        break
      case 'vehicledetails':
        valA = `${a.make} ${a.model}`
        valB = `${b.make} ${b.model}`
        break
      case 'owner':
        valA = a.applicantName || ''
        valB = b.applicantName || ''
        break
      case 'registrationdate':
      default:
        valA = new Date(a.submissionDate).getTime()
        valB = new Date(b.submissionDate).getTime()
        break
    }

    // Handle string comparison
    if (typeof valA === 'string' && typeof valB === 'string') {
      return sortOrder.value === 'asc' ? valA.localeCompare(valB) : valB.localeCompare(valA)
    }

    // Handle numeric comparison (ensure they are numbers)
    const numA = typeof valA === 'number' ? valA : 0
    const numB = typeof valB === 'number' ? valB : 0
    return sortOrder.value === 'asc' ? numA - numB : numB - numA
  })
})

// Sort issued plates
const sortedIssuedPlates = computed(() => {
  return [...issuedPlates.value].sort((a, b) => {
    let valA: string | number, valB: string | number

    switch (sortBy.value) {
      case 'id':
        valA = a.id
        valB = b.id
        break
      case 'vehicledetails':
        valA = `${a.make} ${a.model}`
        valB = `${b.make} ${b.model}`
        break
      case 'owner':
        valA = a.applicantName || ''
        valB = b.applicantName || ''
        break
      case 'platenumber':
        valA = a.plateNumber || ''
        valB = b.plateNumber || ''
        break
      case 'registrationdate':
      default:
        valA = new Date(a.submissionDate).getTime()
        valB = new Date(b.submissionDate).getTime()
        break
    }

    // Handle string comparison
    if (typeof valA === 'string' && typeof valB === 'string') {
      return sortOrder.value === 'asc' ? valA.localeCompare(valB) : valB.localeCompare(valA)
    }

    // Handle numeric comparison (ensure they are numbers)
    const numA = typeof valA === 'number' ? valA : 0
    const numB = typeof valB === 'number' ? valB : 0
    return sortOrder.value === 'asc' ? numA - numB : numB - numA
  })
})

// Generate a random plate number
const generatePlateNumber = (vehicleType: string): string => {
  if (vehicleType === '2-Wheel') {
    // Motorcycle format: MC####
    const number = Math.floor(1000 + Math.random() * 9000) // 1000-9999
    return `MC${number}`
  } else {
    // Car format: ABC 1234
    const letters = 'ABCDEFGHJKLMNPQRSTUVWXYZ' // Excluding I and O to avoid confusion
    let plateLetters = ''
    for (let i = 0; i < 3; i++) {
      plateLetters += letters.charAt(Math.floor(Math.random() * letters.length))
    }
    const number = Math.floor(1000 + Math.random() * 9000) // 1000-9999
    return `${plateLetters} ${number}`
  }
}

// Open the plate issuance modal
const openPlateModal = (registration: VehicleRegistrationForm) => {
  activeRegistration.value = registration
  suggestedPlateNumber.value = generatePlateNumber(registration.vehicleType)
  showPlateModal.value = true
}

// Close the plate issuance modal
const closePlateModal = () => {
  showPlateModal.value = false
  activeRegistration.value = null
}

// Handle plate issuance submission
const handlePlateIssuance = (data: {
  id: string
  plateNumber: string
  plateType: string
  plateIssueDate: string
  plateExpirationDate: string
  issuanceNotes: string
}) => {
  if (!activeRegistration.value) return

  // Update the registration with plate information
  const updatedRegistration = {
    ...activeRegistration.value,
    plateNumber: data.plateNumber,
    plateType: data.plateType,
    plateIssueDate: data.plateIssueDate,
    plateExpirationDate: data.plateExpirationDate,
    issuanceNotes: data.issuanceNotes,
    status: 'completed' as any, // Mark the registration as completed
  }

  // Update the registration in the store
  vehicleRegistrationFormStore.forms = vehicleRegistrationFormStore.forms.map((reg) =>
    reg.id === updatedRegistration.id ? (updatedRegistration as VehicleRegistrationForm) : reg,
  )

  // Create a new plate record
  const newPlate: Partial<Plate> = {
    plateId: parseInt(activeRegistration.value.id.replace(/\D/g, '')), // Remove non-digits and convert to number
    vehicleId: activeRegistration.value.vehicleId
      ? parseInt(activeRegistration.value.vehicleId)
      : 0,
    plate_number: data.plateNumber,
    plate_type: data.plateType,
    plate_issue_date: data.plateIssueDate,
    plate_expiration_date: data.plateExpirationDate,
    status: 'Active',
  }

  // In a real application, you would save the plate to a database
  console.log('New plate created:', newPlate)

  // Close the modal
  closePlateModal()
}
</script>

<template>
  <div class="p-6">
    <!-- Pending Plate Issuance Table -->
    <h1 class="text-2xl font-semibold text-gray-900 mb-6">Pending Plate Issuance</h1>
    <div class="bg-white rounded-lg shadow overflow-hidden mb-8">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                v-for="header in [
                  'Registration ID',
                  'Vehicle Details',
                  'Owner',
                  'Registration Date',
                ]"
                :key="header"
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100"
                @click="toggleSort(header.toLowerCase().replace(' ', ''))"
              >
                {{ header }}
                <span v-if="sortBy === header.toLowerCase().replace(' ', '')">{{
                  sortOrder === 'asc' ? '↑' : '↓'
                }}</span>
              </th>
              <th
                class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="registration in sortedPendingPlateIssuance" :key="registration.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ registration.id }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ `${registration.make} ${registration.model} (${registration.year})` }}
                <span class="ml-2 px-2 py-0.5 text-xs rounded-full bg-blue-100 text-blue-700">{{
                  registration.vehicleType
                }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ registration.applicantName || 'Unknown' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ new Date(registration.submissionDate).toLocaleDateString() }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2">
                <button
                  @click="openPlateModal(registration)"
                  class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                >
                  Issue Plate
                </button>
              </td>
            </tr>
            <tr v-if="sortedPendingPlateIssuance.length === 0">
              <td colspan="5" class="px-6 py-4 text-center text-sm text-gray-500">
                No vehicles pending plate issuance
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Issued Plates Table -->
    <h2 class="text-2xl font-semibold text-gray-900 mb-6">Issued Plates</h2>
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th
                v-for="header in [
                  'Registration ID',
                  'Vehicle Details',
                  'Owner',
                  'Plate Number',
                  'Issue Date',
                  'Expiration Date',
                ]"
                :key="header"
                class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100"
                @click="toggleSort(header.toLowerCase().replace(' ', ''))"
              >
                {{ header }}
                <span v-if="sortBy === header.toLowerCase().replace(' ', '')">{{
                  sortOrder === 'asc' ? '↑' : '↓'
                }}</span>
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="registration in sortedIssuedPlates" :key="registration.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ registration.id }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ `${registration.make} ${registration.model} (${registration.year})` }}
                <span class="ml-2 px-2 py-0.5 text-xs rounded-full bg-blue-100 text-blue-700">{{
                  registration.vehicleType
                }}</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ registration.applicantName || 'Unknown' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ registration.plateNumber }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{
                  registration.plateIssueDate
                    ? new Date(registration.plateIssueDate).toLocaleDateString()
                    : 'N/A'
                }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{
                  registration.plateExpirationDate
                    ? new Date(registration.plateExpirationDate).toLocaleDateString()
                    : 'N/A'
                }}
              </td>
            </tr>
            <tr v-if="sortedIssuedPlates.length === 0">
              <td colspan="6" class="px-6 py-4 text-center text-sm text-gray-500">
                No plates have been issued
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Plate Issuance Modal -->
    <PlateIssuanceModal
      :is-open="showPlateModal"
      :registration="activeRegistration"
      :suggested-plate-number="suggestedPlateNumber"
      @close="closePlateModal"
      @submit="handlePlateIssuance"
    />
  </div>
</template>
