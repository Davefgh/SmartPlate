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
  (e: 'regenerate', vehicleType: string, plateType: string, region: string): string
}>()

// Plate issuance data
const plateNumber = ref('')
const plateType = ref('')
const issuanceNotes = ref('')
const selectedRegion = ref('NCR') // Default region
const lastRegenerationRegion = ref('') // Track the last region used for regeneration
const regenerationError = ref('') // Error message for regeneration

// List of regions for Philippines
const regions = [
  { code: 'NCR', name: 'National Capital Region' },
  { code: 'CALABARZON', name: 'CALABARZON (Region 4A)' },
  { code: 'CENTRAL_LUZON', name: 'Central Luzon (Region 3)' },
  { code: 'WESTERN_VISAYAS', name: 'Western Visayas (Region 6)' },
  { code: 'CENTRAL_VISAYAS', name: 'Central Visayas (Region 7)' },
  { code: 'EASTERN_VISAYAS', name: 'Eastern Visayas (Region 8)' },
  { code: 'NORTHERN_MINDANAO', name: 'Northern Mindanao (Region 10)' },
  { code: 'SOUTHERN_MINDANAO', name: 'Davao Region (Region 11)' },
  { code: 'CAR', name: 'Cordillera Administrative Region' },
  { code: 'CARAGA', name: 'CARAGA (Region 13)' },
  { code: 'BICOL', name: 'Bicol Region (Region 5)' },
  { code: 'ILOCOS', name: 'Ilocos Region (Region 1)' },
  { code: 'MIMAROPA', name: 'MIMAROPA (Region 4B)' },
  { code: 'SOCCSKSARGEN', name: 'SOCCSKSARGEN (Region 12)' },
  { code: 'ZAMBOANGA', name: 'Zamboanga Peninsula (Region 9)' },
  { code: 'BARMM', name: 'Bangsamoro Autonomous Region in Muslim Mindanao' },
]

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
  const commonTypes = ['Private', 'For Hire', 'Government', 'Diplomatic']

  if (vehicleType === '2-Wheel') {
    return ['Motorcycle', 'Tricycle']
  } else if (vehicleType === 'Electric') {
    return [...commonTypes, 'Electric']
  } else if (vehicleType === 'Hybrid') {
    return [...commonTypes, 'Hybrid']
  } else if (vehicleType === 'Trailer') {
    return [...commonTypes, 'Trailer']
  } else if (
    vehicleType.toLowerCase().includes('vintage') ||
    parseInt(props.registration?.year || '0') < 1980
  ) {
    return [...commonTypes, 'Vintage']
  } else {
    return [...commonTypes, 'Electric', 'Hybrid', 'Trailer', 'Vintage']
  }
}

// Generate a plate number based on selected parameters
const regeneratePlateNumber = () => {
  if (!props.registration) return

  // Clear any previous error
  regenerationError.value = ''

  // Only allow regeneration if region has changed since last regeneration
  if (lastRegenerationRegion.value === selectedRegion.value) {
    regenerationError.value = 'Please select a different region to regenerate a new plate number'
    return
  }

  // Call the parent component's regenerate function
  const newPlateNumber = emit(
    'regenerate',
    props.registration.vehicleType,
    plateType.value,
    selectedRegion.value,
  )

  // Record the region used for this regeneration
  lastRegenerationRegion.value = selectedRegion.value

  // In case the parent doesn't return a value, use the suggested plate
  if (!newPlateNumber) {
    plateNumber.value = props.suggestedPlateNumber
  }
}

// Watch for region changes to clear error message
watch(
  () => selectedRegion.value,
  () => {
    regenerationError.value = ''
  },
)

// Initialize form with suggested plate number when registration changes
watch(
  () => [props.registration, props.suggestedPlateNumber],
  () => {
    if (props.registration && props.suggestedPlateNumber) {
      // Keep the current selectedRegion and plateType if they exist
      const currentRegion = selectedRegion.value || 'NCR'
      const currentPlateType = plateType.value || ''

      resetForm()

      // Restore the region selection
      selectedRegion.value = currentRegion
      plateNumber.value = props.suggestedPlateNumber

      // Only set default plate type if it wasn't already set
      if (currentPlateType) {
        // Ensure the plate type is valid for this vehicle
        const validTypes = getPlateTypes(props.registration.vehicleType)
        if (validTypes.includes(currentPlateType)) {
          plateType.value = currentPlateType
        } else {
          // If previous type isn't valid for this vehicle, set default
          setDefaultPlateType(props.registration.vehicleType)
        }
      } else {
        setDefaultPlateType(props.registration.vehicleType)
      }

      // Set the initial regeneration region to track changes
      lastRegenerationRegion.value = currentRegion
    }
  },
  { immediate: true },
)

// Helper function to set default plate type based on vehicle type
const setDefaultPlateType = (vehicleType: string) => {
  if (vehicleType === '2-Wheel') {
    plateType.value = 'Motorcycle'
  } else {
    plateType.value = 'Private'
  }
}

// Reset form to default values
const resetForm = () => {
  plateNumber.value = props.suggestedPlateNumber || ''
  issuanceNotes.value = ''
  plateIssueDate.value = today.toISOString().split('T')[0]
  plateExpirationDate.value = defaultExpirationDate.toISOString().split('T')[0]
  validationErrors.value = {}
}

// Validate form
const validateForm = (): boolean => {
  validationErrors.value = {}
  let isValid = true

  // Validate plate number - the format depends on the selected plate type
  if (!plateNumber.value) {
    validationErrors.value.plateNumber = 'Plate number is required'
    isValid = false
  } else if (props.registration?.vehicleType === '2-Wheel') {
    // Motorcycle format validation
    // Format can be X-NNN or XX-NNNNN
    if (!/^[A-Z]-\d{3}$|^[A-Z]{2}-\d{5}$/.test(plateNumber.value)) {
      validationErrors.value.plateNumber =
        'Motorcycle plate numbers must follow the format: X-NNN or XX-NNNNN'
      isValid = false
    }
  } else if (plateType.value === 'Diplomatic') {
    // Diplomatic format: DDD-NNNN
    if (!/^[A-Z]{3}-\d{4}$/.test(plateNumber.value)) {
      validationErrors.value.plateNumber = 'Diplomatic plates must follow the format DDD-NNNN'
      isValid = false
    }
  } else {
    // Standard format: LLL NNNN (3 letters followed by 4 digits)
    if (!/^[A-Z]{3}\s\d{4}$/.test(plateNumber.value)) {
      validationErrors.value.plateNumber = 'Plate numbers must follow the format LLL NNNN'
      isValid = false
    }
  }

  // Validate plate type
  if (!plateType.value) {
    validationErrors.value.plateType = 'Plate type is required'
    isValid = false
  }

  // Validate selected region
  if (!selectedRegion.value) {
    validationErrors.value.region = 'Region is required'
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
                  Region <span class="text-red-600">*</span>
                </label>
                <select
                  v-model="selectedRegion"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.region,
                    'border-gray-300': !validationErrors.region,
                  }"
                >
                  <option v-for="region in regions" :key="region.code" :value="region.code">
                    {{ region.name }}
                  </option>
                </select>
                <p v-if="validationErrors.region" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.region }}
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

              <div class="mb-3 col-span-2">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Plate Number <span class="text-red-600">*</span>
                </label>
                <div class="flex space-x-2">
                  <div class="relative flex-grow">
                    <input
                      v-model="plateNumber"
                      type="text"
                      class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm font-mono"
                      :class="{
                        'border-red-500': validationErrors.plateNumber,
                        'border-gray-300': !validationErrors.plateNumber,
                      }"
                    />
                  </div>
                  <button
                    @click="plateNumber = suggestedPlateNumber"
                    type="button"
                    class="inline-flex items-center px-3 py-2 border border-gray-300 text-sm font-medium rounded shadow-sm text-gray-700 bg-white hover:bg-gray-50"
                  >
                    Reset
                  </button>
                  <button
                    @click="regeneratePlateNumber"
                    type="button"
                    class="inline-flex items-center px-3 py-2 border border-gray-300 text-sm font-medium rounded shadow-sm"
                    :class="{
                      'text-gray-700 bg-white hover:bg-gray-50':
                        lastRegenerationRegion !== selectedRegion,
                      'opacity-50 cursor-not-allowed text-gray-500 bg-gray-100':
                        lastRegenerationRegion === selectedRegion,
                    }"
                    :title="
                      lastRegenerationRegion === selectedRegion
                        ? 'Change region to generate a new plate'
                        : 'Generate new plate with selected region'
                    "
                  >
                    Regenerate
                  </button>
                </div>
                <p v-if="validationErrors.plateNumber" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.plateNumber }}
                </p>
                <p v-if="regenerationError" class="mt-1 text-sm text-amber-600">
                  {{ regenerationError }}
                </p>
                <p class="mt-1 text-xs text-gray-500">
                  <template v-if="registration.vehicleType === '2-Wheel'">
                    Format: X-NNN or XX-NNNNN
                  </template>
                  <template v-else-if="plateType === 'Diplomatic'"> Format: DDD-NNNN </template>
                  <template v-else>
                    Format: LLL NNNN (Region + Vehicle Type + Sequential)
                  </template>
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
