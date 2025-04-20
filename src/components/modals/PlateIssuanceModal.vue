<script setup lang="ts">
import type { VehicleRegistrationForm } from '@/types/vehicleRegistration'
import { ref, defineProps, defineEmits, watch } from 'vue'

const props = defineProps<{
  isOpen: boolean
  registration: VehicleRegistrationForm | null
  suggestedPlateNumber: string
}>()

const emit = defineEmits<{
  (e: 'close'): void
  (
    e: 'submit',
    data: {
      id: string
      plateNumber: string
      plateType: string
      plateIssueDate: string
      plateExpirationDate: string
      issuanceNotes: string
    },
  ): void
}>()

// Plate issuance data
const plateNumber = ref('')
const plateType = ref('')
const issuanceNotes = ref('')

// Calculate issue date (today) and expiration date (3 years from today)
const today = new Date()
const defaultExpirationDate = new Date(today)
defaultExpirationDate.setFullYear(today.getFullYear() + 3)

const plateIssueDate = ref(today.toISOString().split('T')[0])
const plateExpirationDate = ref(defaultExpirationDate.toISOString().split('T')[0])

// Validation errors
const validationErrors = ref<Record<string, string>>({})

// Plate types based on vehicle type
const getPlateTypes = (vehicleType: string): string[] => {
  if (vehicleType === '2-Wheel') {
    return ['Motorcycle', 'Tricycle']
  } else {
    return ['Private', 'For Hire', 'Government', 'Diplomatic']
  }
}

// Initialize form with suggested plate number when registration changes
watch(
  () => [props.registration, props.suggestedPlateNumber],
  () => {
    if (props.registration && props.suggestedPlateNumber) {
      resetForm()
      plateNumber.value = props.suggestedPlateNumber

      // Set default plate type based on vehicle type
      if (props.registration.vehicleType === '2-Wheel') {
        plateType.value = 'Motorcycle'
      } else {
        plateType.value = 'Private'
      }
    }
  },
  { immediate: true },
)

const resetForm = () => {
  plateNumber.value = props.suggestedPlateNumber || ''
  plateType.value = ''
  issuanceNotes.value = ''
  plateIssueDate.value = today.toISOString().split('T')[0]
  plateExpirationDate.value = defaultExpirationDate.toISOString().split('T')[0]
  validationErrors.value = {}
}

// Validate form
const validateForm = (): boolean => {
  validationErrors.value = {}
  let isValid = true

  // Validate plate number
  if (!plateNumber.value) {
    validationErrors.value.plateNumber = 'Plate number is required'
    isValid = false
  } else if (props.registration?.vehicleType === '2-Wheel' && !plateNumber.value.startsWith('MC')) {
    validationErrors.value.plateNumber = 'Motorcycle plate numbers must start with MC'
    isValid = false
  } else if (
    props.registration?.vehicleType !== '2-Wheel' &&
    !/^[A-Z]{3}\s\d{4}$/.test(plateNumber.value)
  ) {
    validationErrors.value.plateNumber = 'Car plate numbers must follow the format ABC 1234'
    isValid = false
  }

  // Validate plate type
  if (!plateType.value) {
    validationErrors.value.plateType = 'Plate type is required'
    isValid = false
  }

  // Validate issue date
  if (!plateIssueDate.value) {
    validationErrors.value.plateIssueDate = 'Issue date is required'
    isValid = false
  }

  // Validate expiration date
  if (!plateExpirationDate.value) {
    validationErrors.value.plateExpirationDate = 'Expiration date is required'
    isValid = false
  } else {
    const issueDate = new Date(plateIssueDate.value)
    const expirationDate = new Date(plateExpirationDate.value)

    if (expirationDate <= issueDate) {
      validationErrors.value.plateExpirationDate = 'Expiration date must be after issue date'
      isValid = false
    }
  }

  return isValid
}

// Submit plate issuance
const submitPlateIssuance = () => {
  if (!props.registration) return

  if (validateForm()) {
    emit('submit', {
      id: props.registration.id,
      plateNumber: plateNumber.value,
      plateType: plateType.value,
      plateIssueDate: plateIssueDate.value,
      plateExpirationDate: plateExpirationDate.value,
      issuanceNotes: issuanceNotes.value,
    })
  }
}

// Cancel plate issuance
const cancelPlateIssuance = () => {
  emit('close')
}
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center z-50"
  >
    <div
      class="relative mx-auto p-5 border w-full max-w-2xl shadow-lg rounded-md bg-white max-h-screen overflow-y-auto"
    >
      <div class="mt-3">
        <div class="flex justify-between items-center pb-3 border-b">
          <h3 class="text-lg font-medium text-gray-900">Plate Issuance - {{ registration?.id }}</h3>
          <button @click="cancelPlateIssuance" class="text-gray-400 hover:text-gray-500">
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              ></path>
            </svg>
          </button>
        </div>

        <div v-if="registration" class="mt-4">
          <!-- Vehicle Information Section -->
          <div class="bg-gray-50 p-4 rounded-lg mb-4">
            <h4 class="text-md font-medium text-gray-700 mb-2">Vehicle Information</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <p class="text-sm font-medium text-gray-500">Vehicle</p>
                <p class="text-sm">
                  {{ registration.make }} {{ registration.model }} ({{ registration.year }})
                </p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Type</p>
                <p class="text-sm">{{ registration.vehicleType }}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Engine Number</p>
                <p class="text-sm">{{ registration.engineNumber }}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Chassis Number</p>
                <p class="text-sm">{{ registration.chassisNumber }}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Owner</p>
                <p class="text-sm">{{ registration.applicantName || 'Not specified' }}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Registration Type</p>
                <p class="text-sm">{{ registration.registrationType }}</p>
              </div>
            </div>
          </div>

          <!-- Plate Issuance Section -->
          <div class="mt-4">
            <h4 class="text-md font-medium text-gray-700 mb-3">Plate Assignment</h4>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Plate Number <span class="text-red-600">*</span>
                </label>
                <div class="relative">
                  <input
                    v-model="plateNumber"
                    type="text"
                    class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm font-mono"
                    :class="{
                      'border-red-500': validationErrors.plateNumber,
                      'border-gray-300': !validationErrors.plateNumber,
                    }"
                  />
                  <div class="absolute inset-y-0 right-0 flex items-center pr-3">
                    <button
                      @click="plateNumber = suggestedPlateNumber"
                      type="button"
                      class="inline-flex items-center px-2 py-1 border border-gray-300 text-xs font-medium rounded shadow-sm text-gray-700 bg-white hover:bg-gray-50"
                    >
                      Regenerate
                    </button>
                  </div>
                </div>
                <p v-if="validationErrors.plateNumber" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.plateNumber }}
                </p>
                <p class="mt-1 text-xs text-gray-500">
                  {{
                    registration.vehicleType === '2-Wheel' ? 'Format: MC1234' : 'Format: ABC 1234'
                  }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Plate Type <span class="text-red-600">*</span>
                </label>
                <select
                  v-model="plateType"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.plateType,
                    'border-gray-300': !validationErrors.plateType,
                  }"
                >
                  <option value="">Select Plate Type</option>
                  <option
                    v-for="type in getPlateTypes(registration.vehicleType)"
                    :key="type"
                    :value="type"
                  >
                    {{ type }}
                  </option>
                </select>
                <p v-if="validationErrors.plateType" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.plateType }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Issue Date <span class="text-red-600">*</span>
                </label>
                <input
                  v-model="plateIssueDate"
                  type="date"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.plateIssueDate,
                    'border-gray-300': !validationErrors.plateIssueDate,
                  }"
                />
                <p v-if="validationErrors.plateIssueDate" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.plateIssueDate }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Expiration Date <span class="text-red-600">*</span>
                </label>
                <input
                  v-model="plateExpirationDate"
                  type="date"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.plateExpirationDate,
                    'border-gray-300': !validationErrors.plateExpirationDate,
                  }"
                />
                <p v-if="validationErrors.plateExpirationDate" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.plateExpirationDate }}
                </p>
              </div>
            </div>

            <!-- Issuance Notes -->
            <div class="mt-4">
              <h4 class="text-md font-medium text-gray-700 mb-2">Issuance Notes</h4>
              <textarea
                v-model="issuanceNotes"
                rows="3"
                class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm border-gray-300"
                placeholder="Enter any notes regarding this plate issuance..."
              ></textarea>
            </div>
          </div>
        </div>

        <div class="flex justify-end space-x-3 mt-6 pt-3 border-t">
          <button
            @click="cancelPlateIssuance"
            class="px-4 py-2 bg-white border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Cancel
          </button>
          <button
            @click="submitPlateIssuance"
            class="px-4 py-2 bg-blue-600 border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Issue Plate
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
