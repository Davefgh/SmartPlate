<script setup lang="ts">
import type { VehicleRegistrationForm, AdditionalVehicleData } from '@/types/vehicleRegistration'
import { ref, defineProps, defineEmits, computed, watch } from 'vue'

const props = defineProps<{
  isOpen: boolean
  registration: VehicleRegistrationForm | null
}>()

const emit = defineEmits<{
  (e: 'close'): void
  (
    e: 'submit',
    data: {
      id: string
      inspectionStatus: 'approved' | 'rejected'
      inspectionNotes: string
      additionalVehicleData: AdditionalVehicleData
    },
  ): void
}>()

// Inspection data
const inspectionNotes = ref('')
const inspectionStatus = ref<'approved' | 'rejected'>('approved')

// Additional vehicle data for inspection
const additionalData = ref<AdditionalVehicleData>({
  mvFileNumber: '',
  conductionSticker: '',
  vehicleSeries: '',
  bodyType: '',
  pistonDisplacement: 0,
  numberOfCylinders: 0,
  fuelType: '',
  gvw: 0,
  netWeight: 0,
  shippingWeight: 0,
  usageClassification: '',
  firstRegistrationDate: '',
  ltoOfficeCode: '',
  classification: '',
  denomination: '',
  orNumber: '',
  orDate: '',
})

// Validation errors
const validationErrors = ref<Record<string, string>>({})

// Define resetForm before using it in watch
const resetForm = () => {
  additionalData.value = {
    mvFileNumber: '',
    conductionSticker: '',
    vehicleSeries: '',
    bodyType: '',
    pistonDisplacement: 0,
    numberOfCylinders: 0,
    fuelType: '',
    gvw: 0,
    netWeight: 0,
    shippingWeight: 0,
    usageClassification: '',
    firstRegistrationDate: '',
    ltoOfficeCode: '',
    classification: '',
    denomination: '',
    orNumber: '',
    orDate: '',
  }
  inspectionNotes.value = ''
  inspectionStatus.value = 'approved'
  validationErrors.value = {}
}

// Reset data when registration changes
watch(
  () => props.registration,
  () => {
    // Initialize with existing data if available
    if (props.registration?.additionalVehicleData) {
      additionalData.value = { ...props.registration.additionalVehicleData }
      inspectionNotes.value = props.registration.inspectionNotes || ''
    } else {
      // Reset to defaults otherwise
      resetForm()
    }
  },
  { immediate: true },
)

// Determine required fields based on vehicle type
const requiredFields = computed(() => {
  const baseFields = ['mvFileNumber', 'bodyType', 'fuelType', 'usageClassification']

  // Add fields specific to vehicle type
  if (props.registration?.vehicleType === '4-Wheel') {
    return [...baseFields, 'pistonDisplacement', 'numberOfCylinders', 'gvw', 'netWeight']
  } else if (props.registration?.vehicleType === '2-Wheel') {
    return [...baseFields, 'pistonDisplacement']
  }

  return baseFields
})

// Validate the form
const validateForm = (): boolean => {
  validationErrors.value = {}
  let isValid = true

  // Validate required fields
  requiredFields.value.forEach((field) => {
    const value = additionalData.value[field as keyof AdditionalVehicleData]
    if (!value || (typeof value === 'string' && value.trim() === '')) {
      validationErrors.value[field] = 'This field is required'
      isValid = false
    }
  })

  // Additional validations
  if (additionalData.value.pistonDisplacement < 0) {
    validationErrors.value.pistonDisplacement = 'Must be a positive number'
    isValid = false
  }

  if (additionalData.value.numberOfCylinders < 0) {
    validationErrors.value.numberOfCylinders = 'Must be a positive number'
    isValid = false
  }

  if (additionalData.value.gvw < 0) {
    validationErrors.value.gvw = 'Must be a positive number'
    isValid = false
  }

  if (additionalData.value.netWeight < 0) {
    validationErrors.value.netWeight = 'Must be a positive number'
    isValid = false
  }

  // Require notes if inspection is rejected
  if (inspectionStatus.value === 'rejected' && !inspectionNotes.value.trim()) {
    validationErrors.value.inspectionNotes = 'Notes are required when rejecting an inspection'
    isValid = false
  }

  return isValid
}

// Submit inspection details
const submitInspection = () => {
  if (!props.registration) return

  if (validateForm()) {
    emit('submit', {
      id: props.registration.id,
      inspectionStatus: inspectionStatus.value,
      inspectionNotes: inspectionNotes.value,
      additionalVehicleData: { ...additionalData.value },
    })
  }
}

// Cancel inspection
const cancelInspection = () => {
  emit('close')
}
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center z-50"
  >
    <div
      class="relative mx-auto p-5 border w-full max-w-4xl shadow-lg rounded-md bg-white max-h-screen overflow-y-auto"
    >
      <div class="mt-3">
        <div class="flex justify-between items-center pb-3 border-b">
          <h3 class="text-lg font-medium text-gray-900">
            Vehicle Inspection Process - {{ registration?.id }}
          </h3>
          <button @click="cancelInspection" class="text-gray-400 hover:text-gray-500">
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

        <div v-if="registration" class="mt-4 grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-4">
          <!-- Verification Code Banner -->
          <div class="col-span-2 bg-blue-50 p-4 rounded-lg mb-4 border border-blue-200">
            <div class="flex items-center justify-between">
              <div>
                <h4 class="text-md font-medium text-blue-700 mb-1">Inspection Verification Code</h4>
                <p class="text-sm text-blue-600">
                  {{ registration.inspectionCode || 'No inspection code available' }}
                </p>
              </div>
              <div class="text-blue-600 bg-blue-100 px-3 py-1 rounded-lg text-xs font-medium">
                VERIFY BEFORE INSPECTION
              </div>
            </div>
          </div>

          <!-- Current Vehicle Information Section -->
          <div class="col-span-2 bg-gray-50 p-4 rounded-lg mb-4">
            <h4 class="text-md font-medium text-gray-700 mb-2">Current Vehicle Information</h4>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div>
                <p class="text-sm font-medium text-gray-500">Make</p>
                <p class="text-sm">{{ registration.make }}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Model</p>
                <p class="text-sm">{{ registration.model }}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Year</p>
                <p class="text-sm">{{ registration.year }}</p>
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
                <p class="text-sm font-medium text-gray-500">Color</p>
                <p class="text-sm">{{ registration.color }}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Vehicle Type</p>
                <p class="text-sm">{{ registration.vehicleType }}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Registration Type</p>
                <p class="text-sm">{{ registration.registrationType }}</p>
              </div>
            </div>
          </div>

          <!-- Additional Vehicle Information Section -->
          <div class="col-span-2">
            <h4 class="text-md font-medium text-gray-700 mb-3">Complete Vehicle Information</h4>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  MV File Number
                  <span v-if="requiredFields.includes('mvFileNumber')" class="text-red-600">*</span>
                </label>
                <input
                  v-model="additionalData.mvFileNumber"
                  type="text"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.mvFileNumber,
                    'border-gray-300': !validationErrors.mvFileNumber,
                  }"
                />
                <p v-if="validationErrors.mvFileNumber" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.mvFileNumber }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Conduction Sticker
                </label>
                <input
                  v-model="additionalData.conductionSticker"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1"> Vehicle Series </label>
                <input
                  v-model="additionalData.vehicleSeries"
                  type="text"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Body Type
                  <span v-if="requiredFields.includes('bodyType')" class="text-red-600">*</span>
                </label>
                <input
                  v-model="additionalData.bodyType"
                  type="text"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.bodyType,
                    'border-gray-300': !validationErrors.bodyType,
                  }"
                />
                <p v-if="validationErrors.bodyType" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.bodyType }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Piston Displacement (cc)
                  <span v-if="requiredFields.includes('pistonDisplacement')" class="text-red-600"
                    >*</span
                  >
                </label>
                <input
                  v-model.number="additionalData.pistonDisplacement"
                  type="number"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.pistonDisplacement,
                    'border-gray-300': !validationErrors.pistonDisplacement,
                  }"
                />
                <p v-if="validationErrors.pistonDisplacement" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.pistonDisplacement }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Number of Cylinders
                  <span v-if="requiredFields.includes('numberOfCylinders')" class="text-red-600"
                    >*</span
                  >
                </label>
                <input
                  v-model.number="additionalData.numberOfCylinders"
                  type="number"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.numberOfCylinders,
                    'border-gray-300': !validationErrors.numberOfCylinders,
                  }"
                />
                <p v-if="validationErrors.numberOfCylinders" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.numberOfCylinders }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Fuel Type
                  <span v-if="requiredFields.includes('fuelType')" class="text-red-600">*</span>
                </label>
                <select
                  v-model="additionalData.fuelType"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.fuelType,
                    'border-gray-300': !validationErrors.fuelType,
                  }"
                >
                  <option value="">Select Fuel Type</option>
                  <option value="Gasoline">Gasoline</option>
                  <option value="Diesel">Diesel</option>
                  <option value="Electric">Electric</option>
                  <option value="Hybrid">Hybrid</option>
                  <option value="LPG">LPG</option>
                </select>
                <p v-if="validationErrors.fuelType" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.fuelType }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Gross Vehicle Weight (kg)
                  <span v-if="requiredFields.includes('gvw')" class="text-red-600">*</span>
                </label>
                <input
                  v-model.number="additionalData.gvw"
                  type="number"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.gvw,
                    'border-gray-300': !validationErrors.gvw,
                  }"
                />
                <p v-if="validationErrors.gvw" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.gvw }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Net Weight (kg)
                  <span v-if="requiredFields.includes('netWeight')" class="text-red-600">*</span>
                </label>
                <input
                  v-model.number="additionalData.netWeight"
                  type="number"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.netWeight,
                    'border-gray-300': !validationErrors.netWeight,
                  }"
                />
                <p v-if="validationErrors.netWeight" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.netWeight }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Shipping Weight (kg)
                </label>
                <input
                  v-model.number="additionalData.shippingWeight"
                  type="number"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Usage Classification
                  <span v-if="requiredFields.includes('usageClassification')" class="text-red-600"
                    >*</span
                  >
                </label>
                <select
                  v-model="additionalData.usageClassification"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.usageClassification,
                    'border-gray-300': !validationErrors.usageClassification,
                  }"
                >
                  <option value="">Select Classification</option>
                  <option value="Private">Private</option>
                  <option value="For Hire">For Hire</option>
                  <option value="Government">Government</option>
                  <option value="Diplomatic">Diplomatic</option>
                </select>
                <p v-if="validationErrors.usageClassification" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.usageClassification }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  First Registration Date
                </label>
                <input
                  v-model="additionalData.firstRegistrationDate"
                  type="date"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>
            </div>
          </div>

          <!-- Inspection Findings -->
          <div class="col-span-2 mt-4">
            <h4 class="text-md font-medium text-gray-700 mb-2">
              Inspection Findings & Notes
              <span v-if="inspectionStatus === 'rejected'" class="text-red-600">*</span>
            </h4>
            <textarea
              v-model="inspectionNotes"
              rows="4"
              class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              :class="{
                'border-red-500': validationErrors.inspectionNotes,
                'border-gray-300': !validationErrors.inspectionNotes,
              }"
              placeholder="Enter inspection notes, findings, and recommendations..."
            ></textarea>
            <p v-if="validationErrors.inspectionNotes" class="mt-1 text-sm text-red-600">
              {{ validationErrors.inspectionNotes }}
            </p>
          </div>

          <!-- Inspection Status -->
          <div class="col-span-2 mt-4">
            <h4 class="text-md font-medium text-gray-700 mb-2">Inspection Result</h4>
            <div class="flex space-x-4">
              <label class="inline-flex items-center">
                <input
                  v-model="inspectionStatus"
                  type="radio"
                  class="form-radio h-4 w-4 text-blue-600"
                  name="inspectionStatus"
                  value="approved"
                />
                <span class="ml-2 text-sm text-gray-700">Approved</span>
              </label>
              <label class="inline-flex items-center">
                <input
                  v-model="inspectionStatus"
                  type="radio"
                  class="form-radio h-4 w-4 text-red-600"
                  name="inspectionStatus"
                  value="rejected"
                />
                <span class="ml-2 text-sm text-gray-700">Rejected</span>
              </label>
            </div>
          </div>
        </div>

        <div class="flex justify-end space-x-3 mt-6 pt-3 border-t">
          <button
            @click="cancelInspection"
            class="px-4 py-2 bg-white border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Cancel
          </button>
          <button
            @click="submitInspection"
            class="px-4 py-2 bg-blue-600 border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Submit Inspection
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
