<script setup lang="ts">
import type { VehicleRegistrationForm } from '@/types/vehicleRegistration'
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
      paymentStatus: 'approved' | 'rejected'
      paymentNotes: string
      paymentDetails: {
        amountPaid: number
        paymentDate: string
        paymentMethod: string
        receiptNumber: string
        referenceNumber: string
      }
    },
  ): void
}>()

// Payment data
const paymentNotes = ref('')
const paymentStatus = ref<'approved' | 'rejected'>('approved')
const paymentDetails = ref({
  amountPaid: 0,
  paymentDate: new Date().toISOString().split('T')[0],
  paymentMethod: '',
  receiptNumber: '',
  referenceNumber: '',
})

// Generate a unique payment reference number
const generateReferenceNumber = (): string => {
  const timestamp = new Date().getTime().toString().slice(-6)
  const random = Math.floor(Math.random() * 10000)
    .toString()
    .padStart(4, '0')
  return `PAY-${timestamp}-${random}`
}

// Validation errors
const validationErrors = ref<Record<string, string>>({})

// Get registration fee based on vehicle type
const registrationFee = computed(() => {
  if (props.registration?.vehicleType === '4-Wheel') {
    return 1620.0 // Example fee for 4-wheel vehicles
  } else if (props.registration?.vehicleType === '2-Wheel') {
    return 820.0 // Example fee for 2-wheel vehicles
  }
  return 1000.0 // Default fee
})

// Calculate computer fee
const computerFee = computed(() => 169.0) // Example fixed computer fee

// Calculate total fee
const totalFee = computed(() => {
  let total = registrationFee.value + computerFee.value

  // Add plate issuance fee for new registrations
  if (props.registration?.isNewVehicle) {
    total += 450.0
  }

  return total.toFixed(2)
})

// Define resetForm before it's used in the watch
const resetForm = () => {
  paymentNotes.value = ''
  paymentStatus.value = 'approved'

  // Set default payment date to today
  paymentDetails.value = {
    amountPaid: parseFloat(totalFee.value),
    paymentDate: new Date().toISOString().split('T')[0],
    paymentMethod: '',
    receiptNumber: '',
    referenceNumber: generateReferenceNumber(),
  }

  validationErrors.value = {}
}

// Reset form when registration changes
watch(
  () => props.registration,
  () => {
    if (props.registration) {
      resetForm()
    }
  },
  { immediate: true },
)

// Validate form
const validateForm = (): boolean => {
  validationErrors.value = {}
  let isValid = true

  // Check if amount paid matches total fee
  if (paymentDetails.value.amountPaid <= 0) {
    validationErrors.value.amountPaid = 'Payment amount is required'
    isValid = false
  } else if (paymentDetails.value.amountPaid < parseFloat(totalFee.value)) {
    validationErrors.value.amountPaid = `Amount must be at least ₱${totalFee.value}`
    isValid = false
  }

  // Check payment date
  if (!paymentDetails.value.paymentDate) {
    validationErrors.value.paymentDate = 'Payment date is required'
    isValid = false
  }

  // Check payment method
  if (!paymentDetails.value.paymentMethod) {
    validationErrors.value.paymentMethod = 'Payment method is required'
    isValid = false
  }

  // Check receipt number
  if (!paymentDetails.value.receiptNumber) {
    validationErrors.value.receiptNumber = 'Receipt number is required'
    isValid = false
  }

  // If rejecting, require notes
  if (paymentStatus.value === 'rejected' && !paymentNotes.value.trim()) {
    validationErrors.value.paymentNotes = 'Notes are required when rejecting a payment'
    isValid = false
  }

  return isValid
}

// Submit payment
const submitPayment = () => {
  if (!props.registration) return

  if (validateForm()) {
    // Ensure reference number is generated if not already present
    if (!paymentDetails.value.referenceNumber) {
      paymentDetails.value.referenceNumber = generateReferenceNumber()
    }

    emit('submit', {
      id: props.registration.id,
      paymentStatus: paymentStatus.value,
      paymentNotes: paymentNotes.value,
      paymentDetails: { ...paymentDetails.value },
    })
  }
}

// Cancel payment
const cancelPayment = () => {
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
          <h3 class="text-lg font-medium text-gray-900">
            Payment Processing - {{ registration?.id }}
          </h3>
          <button @click="cancelPayment" class="text-gray-400 hover:text-gray-500">
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
          <!-- Payment Verification Code Banner -->
          <div class="bg-blue-50 p-4 rounded-lg mb-4 border border-blue-200">
            <div class="flex items-center justify-between">
              <div>
                <h4 class="text-md font-medium text-blue-700 mb-1">Payment Verification Code</h4>
                <p class="text-sm text-blue-600">
                  {{ registration.paymentCode || 'No payment code available' }}
                </p>
              </div>
              <div class="text-blue-600 bg-blue-100 px-3 py-1 rounded-lg text-xs font-medium">
                VERIFY BEFORE PROCESSING
              </div>
            </div>
          </div>

          <!-- Payment Information Section -->
          <div class="bg-gray-50 p-4 rounded-lg mb-4">
            <h4 class="text-md font-medium text-gray-700 mb-2">Registration Details</h4>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <p class="text-sm font-medium text-gray-500">Vehicle</p>
                <p class="text-sm">
                  {{ registration.make }} {{ registration.model }} ({{ registration.year }})
                </p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Registration Type</p>
                <p class="text-sm">{{ registration.registrationType }}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500">Owner</p>
                <p class="text-sm">{{ registration.applicantName || 'Not specified' }}</p>
              </div>
            </div>

            <!-- Fee Breakdown -->
            <div class="mt-4 border-t pt-3">
              <h5 class="text-sm font-medium text-gray-700 mb-2">Fee Breakdown</h5>
              <div class="space-y-1">
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">Registration Fee</span>
                  <span class="text-sm text-gray-600">₱{{ registrationFee.toFixed(2) }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-sm text-gray-600">Computer Fee</span>
                  <span class="text-sm text-gray-600">₱{{ computerFee.toFixed(2) }}</span>
                </div>
                <div v-if="registration.isNewVehicle" class="flex justify-between">
                  <span class="text-sm text-gray-600">Plate Issuance Fee</span>
                  <span class="text-sm text-gray-600">₱450.00</span>
                </div>
                <div class="flex justify-between font-medium border-t pt-1 mt-1">
                  <span class="text-sm text-gray-700">Total</span>
                  <span class="text-sm text-gray-700">₱{{ totalFee }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Payment Processing Section -->
          <div class="mt-4">
            <h4 class="text-md font-medium text-gray-700 mb-3">Payment Processing</h4>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Amount Paid (₱) <span class="text-red-600">*</span>
                </label>
                <input
                  v-model.number="paymentDetails.amountPaid"
                  type="number"
                  min="0"
                  step="0.01"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.amountPaid,
                    'border-gray-300': !validationErrors.amountPaid,
                  }"
                />
                <p v-if="validationErrors.amountPaid" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.amountPaid }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Payment Reference Number
                </label>
                <div class="flex">
                  <input
                    v-model="paymentDetails.referenceNumber"
                    type="text"
                    readonly
                    class="w-full px-3 py-2 border border-gray-300 bg-gray-50 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm font-mono"
                  />
                  <button
                    @click="paymentDetails.referenceNumber = generateReferenceNumber()"
                    type="button"
                    class="ml-2 inline-flex items-center px-2 py-1 border border-gray-300 text-xs font-medium rounded shadow-sm text-gray-700 bg-white hover:bg-gray-50"
                  >
                    Regenerate
                  </button>
                </div>
                <p class="mt-1 text-xs text-gray-500">Automatically generated payment reference</p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Payment Date <span class="text-red-600">*</span>
                </label>
                <input
                  v-model="paymentDetails.paymentDate"
                  type="date"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.paymentDate,
                    'border-gray-300': !validationErrors.paymentDate,
                  }"
                />
                <p v-if="validationErrors.paymentDate" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.paymentDate }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Payment Method <span class="text-red-600">*</span>
                </label>
                <select
                  v-model="paymentDetails.paymentMethod"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.paymentMethod,
                    'border-gray-300': !validationErrors.paymentMethod,
                  }"
                >
                  <option value="">Select Payment Method</option>
                  <option value="Cash">Cash</option>
                  <option value="Credit Card">Credit Card</option>
                  <option value="Debit Card">Debit Card</option>
                  <option value="Bank Transfer">Bank Transfer</option>
                  <option value="E-Wallet">E-Wallet</option>
                </select>
                <p v-if="validationErrors.paymentMethod" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.paymentMethod }}
                </p>
              </div>

              <div class="mb-3">
                <label class="block text-sm font-medium text-gray-700 mb-1">
                  Receipt Number <span class="text-red-600">*</span>
                </label>
                <input
                  v-model="paymentDetails.receiptNumber"
                  type="text"
                  class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  :class="{
                    'border-red-500': validationErrors.receiptNumber,
                    'border-gray-300': !validationErrors.receiptNumber,
                  }"
                />
                <p v-if="validationErrors.receiptNumber" class="mt-1 text-sm text-red-600">
                  {{ validationErrors.receiptNumber }}
                </p>
              </div>
            </div>

            <!-- Payment Notes -->
            <div class="mt-4">
              <h4 class="text-md font-medium text-gray-700 mb-2">
                Payment Notes
                <span v-if="paymentStatus === 'rejected'" class="text-red-600">*</span>
              </h4>
              <textarea
                v-model="paymentNotes"
                rows="3"
                class="w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                :class="{
                  'border-red-500': validationErrors.paymentNotes,
                  'border-gray-300': !validationErrors.paymentNotes,
                }"
                placeholder="Enter payment notes or reasons for rejection..."
              ></textarea>
              <p v-if="validationErrors.paymentNotes" class="mt-1 text-sm text-red-600">
                {{ validationErrors.paymentNotes }}
              </p>
            </div>

            <!-- Payment Status -->
            <div class="mt-4">
              <h4 class="text-md font-medium text-gray-700 mb-2">Payment Status</h4>
              <div class="flex space-x-4">
                <label class="inline-flex items-center">
                  <input
                    v-model="paymentStatus"
                    type="radio"
                    class="form-radio h-4 w-4 text-blue-600"
                    name="paymentStatus"
                    value="approved"
                  />
                  <span class="ml-2 text-sm text-gray-700">Approved</span>
                </label>
                <label class="inline-flex items-center">
                  <input
                    v-model="paymentStatus"
                    type="radio"
                    class="form-radio h-4 w-4 text-red-600"
                    name="paymentStatus"
                    value="rejected"
                  />
                  <span class="ml-2 text-sm text-gray-700">Rejected</span>
                </label>
              </div>
            </div>
          </div>
        </div>

        <div class="flex justify-end space-x-3 mt-6 pt-3 border-t">
          <button
            @click="cancelPayment"
            class="px-4 py-2 bg-white border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Cancel
          </button>
          <button
            @click="submitPayment"
            class="px-4 py-2 bg-blue-600 border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Process Payment
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
