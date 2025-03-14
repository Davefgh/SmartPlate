<script setup>
import { ref, reactive, computed } from 'vue'

// Current step tracking
const currentStep = ref(1)
const totalSteps = 4

// Form data
const formData = reactive({
  // Step 1: Vehicle Information
  vehicleType: '',
  make: '',
  model: '',
  year: null,
  engineNumber: '',
  chassisNumber: '',
  color: '',
  
  // Step 2: Ownership Documents
  isNewVehicle: true,
  documents: {
    csr: null, // Certificate of Stock Report
    salesInvoice: null,
    insurance: null,
    orCr: null, // Official Receipt/Certificate of Registration
    deedOfSale: null,
    pnpHpgClearance: null
  },
  
  // Step 3: Inspection & Verification
  appointmentDate: null,
  appointmentTime: null,
  referenceCode: '',
  inspectionStatus: 'pending',
  
  // Step 4: Payment & Processing
  referenceSlip: null,
  paymentStatus: 'pending',
  verificationStatus: 'pending'
})

// Required documents based on vehicle type
const requiredDocuments = computed(() => {
  if (formData.isNewVehicle) {
    return ['csr', 'salesInvoice', 'insurance']
  } else {
    return ['orCr', 'deedOfSale', 'pnpHpgClearance']
  }
})

// Navigation functions
const nextStep = () => {
  if (currentStep.value < totalSteps) {
    currentStep.value++
  }
}

const prevStep = () => {
  if (currentStep.value > 1) {
    currentStep.value--
  }
}

// Generate appointment details
const generateAppointment = () => {
  // Generate tomorrow's date
  const tomorrow = new Date()
  tomorrow.setDate(tomorrow.getDate() + 1)
  formData.appointmentDate = tomorrow.toISOString().split('T')[0]
  
  // Generate random time between 9am and 4pm
  const hour = Math.floor(Math.random() * 8) + 9
  formData.appointmentTime = `${hour}:00 ${hour >= 12 ? 'PM' : 'AM'}`
  
  // Generate reference code
  formData.referenceCode = 'LTO-' + Math.random().toString(36).substring(2, 10).toUpperCase()
}

// Handle file uploads
const handleFileUpload = (event, documentType) => {
  const file = event.target.files[0]
  if (file) {
    formData.documents[documentType] = file
  }
}

// Submit the registration
const submitRegistration = () => {
  // In a real app, this would send the data to the backend
  // For now, we'll just show success and reset
  alert('Registration submitted successfully!')
  currentStep.value = 1
  // Reset form data
  Object.keys(formData).forEach(key => {
    if (typeof formData[key] === 'object' && formData[key] !== null) {
      Object.keys(formData[key]).forEach(subKey => {
        formData[key][subKey] = null
      })
    } else {
      formData[key] = typeof formData[key] === 'boolean' ? formData[key] : ''
    }
  })
}
</script>

<template>
  <div class="p-6 max-w-4xl mx-auto">
    <div class="bg-white rounded-xl shadow-md overflow-hidden">
      <!-- Header -->
      <div class="bg-gradient-to-r from-dark-blue to-light-blue p-6 text-white">
        <h1 class="text-2xl font-bold">Step-by-Step Vehicle Registration</h1>
        <p class="mt-2 text-blue-100">Complete all steps to register your vehicle with LTO</p>
        
        <!-- Progress bar -->
        <div class="mt-6 bg-blue-200 bg-opacity-30 h-2 rounded-full">
          <div 
            class="bg-white h-full rounded-full transition-all duration-500"
            :style="`width: ${(currentStep / totalSteps) * 100}%`"
          ></div>
        </div>
        
        <!-- Steps indicator -->
        <div class="flex justify-between mt-2">
          <div 
            v-for="step in totalSteps" 
            :key="step"
            class="flex flex-col items-center"
          >
            <div 
              :class="[
                'w-8 h-8 rounded-full flex items-center justify-center mb-1 transition-all',
                currentStep >= step 
                  ? 'bg-white text-dark-blue' 
                  : 'bg-blue-200 bg-opacity-30 text-white'
              ]"
            >
              {{ step }}
            </div>
            <span class="text-xs text-blue-100">Step {{ step }}</span>
          </div>
        </div>
      </div>
      
      <!-- Form content -->
      <div class="p-6">
        <!-- Step 1: Vehicle Information -->
        <div v-if="currentStep === 1" class="fade-in">
          <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center">
            <span class="bg-dark-blue text-white w-7 h-7 rounded-full flex items-center justify-center mr-2">1</span>
            Vehicle Information
          </h2>
          
          <div class="space-y-6">
            <div class="mb-4">
              <label class="block text-gray-700 font-medium mb-2">Vehicle Type</label>
              <div class="flex space-x-4">
                <label class="flex items-center cursor-pointer">
                  <input 
                    type="radio" 
                    v-model="formData.vehicleType" 
                    value="2-wheel"
                    class="form-radio h-5 w-5 text-dark-blue"
                  >
                  <span class="ml-2">2-Wheel</span>
                </label>
                <label class="flex items-center cursor-pointer">
                  <input 
                    type="radio" 
                    v-model="formData.vehicleType" 
                    value="4-wheel"
                    class="form-radio h-5 w-5 text-dark-blue"
                  >
                  <span class="ml-2">4-Wheel</span>
                </label>
              </div>
            </div>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-gray-700 font-medium mb-2">Make</label>
                <input 
                  type="text" 
                  v-model="formData.make"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-dark-blue focus:border-transparent"
                  placeholder="e.g. Toyota"
                >
              </div>
              
              <div>
                <label class="block text-gray-700 font-medium mb-2">Model</label>
                <input 
                  type="text" 
                  v-model="formData.model"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-dark-blue focus:border-transparent"
                  placeholder="e.g. Corolla"
                >
              </div>
              
              <div>
                <label class="block text-gray-700 font-medium mb-2">Year</label>
                <input 
                  type="number" 
                  v-model="formData.year"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-dark-blue focus:border-transparent"
                  placeholder="e.g. 2023"
                >
              </div>
              
              <div>
                <label class="block text-gray-700 font-medium mb-2">Color</label>
                <input 
                  type="text" 
                  v-model="formData.color"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-dark-blue focus:border-transparent"
                  placeholder="e.g. White"
                >
              </div>
              
              <div>
                <label class="block text-gray-700 font-medium mb-2">Engine Number</label>
                <input 
                  type="text" 
                  v-model="formData.engineNumber"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-dark-blue focus:border-transparent"
                  placeholder="Enter engine number"
                >
              </div>
              
              <div>
                <label class="block text-gray-700 font-medium mb-2">Chassis Number</label>
                <input 
                  type="text" 
                  v-model="formData.chassisNumber"
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-dark-blue focus:border-transparent"
                  placeholder="Enter chassis number"
                >
              </div>
            </div>
          </div>
        </div>
        
        <!-- Step 2: Ownership Documents -->
        <div v-if="currentStep === 2" class="fade-in">
          <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center">
            <span class="bg-dark-blue text-white w-7 h-7 rounded-full flex items-center justify-center mr-2">2</span>
            Ownership Documents
          </h2>
          
          <div class="mb-6">
            <label class="block text-gray-700 font-medium mb-2">Vehicle Status</label>
            <div class="flex space-x-4">
              <label class="flex items-center cursor-pointer">
                <input 
                  type="radio" 
                  v-model="formData.isNewVehicle" 
                  :value="true"
                  class="form-radio h-5 w-5 text-dark-blue"
                >
                <span class="ml-2">Brand New</span>
              </label>
              <label class="flex items-center cursor-pointer">
                <input 
                  type="radio" 
                  v-model="formData.isNewVehicle" 
                  :value="false"
                  class="form-radio h-5 w-5 text-dark-blue"
                >
                <span class="ml-2">Second-Hand</span>
              </label>
            </div>
          </div>
          
          <div class="bg-blue-50 p-4 rounded-lg mb-6">
            <h3 class="font-medium text-dark-blue mb-2">Required Documents</h3>
            <ul class="list-disc pl-5 text-sm text-gray-700">
              <template v-if="formData.isNewVehicle">
                <li v-for="doc in requiredDocuments" :key="doc">
                  {{ doc === 'csr' ? 'Certificate of Stock Report (CSR)' : 
                     doc === 'salesInvoice' ? 'Sales Invoice' : 
                     doc === 'insurance' ? 'Insurance Certificate' : doc }}
                </li>
              </template>
              <template v-else>
                <li v-for="doc in requiredDocuments" :key="doc">
                  {{ doc === 'orCr' ? 'Original Receipt/Certificate of Registration (OR/CR)' : 
                     doc === 'deedOfSale' ? 'Deed of Sale' : 
                     doc === 'pnpHpgClearance' ? 'PNP-HPG Clearance' : doc }}
                </li>
              </template>
            </ul>
          </div>
          
          <div class="space-y-6">
            <template v-if="formData.isNewVehicle">
              <div>
                <label class="block text-gray-700 font-medium mb-2">Certificate of Stock Report (CSR)</label>
                <div class="flex items-center">
                  <input 
                    type="file" 
                    @change="(e) => handleFileUpload(e, 'csr')"
                    class="hidden" 
                    id="csr-upload"
                    accept=".pdf,.jpg,.jpeg,.png"
                  >
                  <label 
                    for="csr-upload"
                    class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg cursor-pointer hover:bg-gray-200 transition-colors"
                  >
                    <font-awesome-icon :icon="['fas', 'upload']" class="mr-2" />
                    Choose File
                  </label>
                  <span class="ml-3 text-sm text-gray-600">
                    {{ formData.documents.csr ? formData.documents.csr.name : 'No file chosen' }}
                  </span>
                </div>
              </div>
              
              <div>
                <label class="block text-gray-700 font-medium mb-2">Sales Invoice</label>
                <div class="flex items-center">
                  <input 
                    type="file" 
                    @change="(e) => handleFileUpload(e, 'salesInvoice')"
                    class="hidden" 
                    id="sales-invoice-upload"
                    accept=".pdf,.jpg,.jpeg,.png"
                  >
                  <label 
                    for="sales-invoice-upload"
                    class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg cursor-pointer hover:bg-gray-200 transition-colors"
                  >
                    <font-awesome-icon :icon="['fas', 'upload']" class="mr-2" />
                    Choose File
                  </label>
                  <span class="ml-3 text-sm text-gray-600">
                    {{ formData.documents.salesInvoice ? formData.documents.salesInvoice.name : 'No file chosen' }}
                  </span>
                </div>
              </div>
              
              <div>
                <label class="block text-gray-700 font-medium mb-2">Insurance Certificate</label>
                <div class="flex items-center">
                  <input 
                    type="file" 
                    @change="(e) => handleFileUpload(e, 'insurance')"
                    class="hidden" 
                    id="insurance-upload"
                    accept=".pdf,.jpg,.jpeg,.png"
                  >
                  <label 
                    for="insurance-upload"
                    class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg cursor-pointer hover:bg-gray-200 transition-colors"
                  >
                    <font-awesome-icon :icon="['fas', 'upload']" class="mr-2" />
                    Choose File
                  </label>
                  <span class="ml-3 text-sm text-gray-600">
                    {{ formData.documents.insurance ? formData.documents.insurance.name : 'No file chosen' }}
                  </span>
                </div>
              </div>
            </template>
            
            <template v-else>
              <div>
                <label class="block text-gray-700 font-medium mb-2">Original Receipt/Certificate of Registration (OR/CR)</label>
                <div class="flex items-center">
                  <input 
                    type="file" 
                    @change="(e) => handleFileUpload(e, 'orCr')"
                    class="hidden" 
                    id="orcr-upload"
                    accept=".pdf,.jpg,.jpeg,.png"
                  >
                  <label 
                    for="orcr-upload"
                    class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg cursor-pointer hover:bg-gray-200 transition-colors"
                  >
                    <font-awesome-icon :icon="['fas', 'upload']" class="mr-2" />
                    Choose File
                  </label>
                  <span class="ml-3 text-sm text-gray-600">
                    {{ formData.documents.orCr ? formData.documents.orCr.name : 'No file chosen' }}
                  </span>
                </div>
              </div>
              
              <div>
                <label class="block text-gray-700 font-medium mb-2">Deed of Sale</label>
                <div class="flex items-center">
                  <input 
                    type="file" 
                    @change="(e) => handleFileUpload(e, 'deedOfSale')"
                    class="hidden" 
                    id="deed-upload"
                    accept=".pdf,.jpg,.jpeg,.png"
                  >
                  <label 
                    for="deed-upload"
                    class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg cursor-pointer hover:bg-gray-200 transition-colors"
                  >
                    <font-awesome-icon :icon="['fas', 'upload']" class="mr-2" />
                    Choose File
                  </label>
                  <span class="ml-3 text-sm text-gray-600">
                    {{ formData.documents.deedOfSale ? formData.documents.deedOfSale.name : 'No file chosen' }}
                  </span>
                </div>
              </div>
              
              <div>
                <label class="block text-gray-700 font-medium mb-2">PNP-HPG Clearance</label>
                <div class="flex items-center">
                  <input 
                    type="file" 
                    @change="(e) => handleFileUpload(e, 'pnpHpgClearance')"
                    class="hidden" 
                    id="clearance-upload"
                    accept=".pdf,.jpg,.jpeg,.png"
                  >
                  <label 
                    for="clearance-upload"
                    class="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg cursor-pointer hover:bg-gray-200 transition-colors"
                  >
                    <font-awesome-icon :icon="['fas', 'upload']" class="mr-2" />
                    Choose File
                  </label>
                  <span class="ml-3 text-sm text-gray-600">
                    {{ formData.documents.pnpHpgClearance ? formData.documents.pnpHpgClearance.name : 'No file chosen' }}
                  </span>
                </div>
              </div>
            </template>
          </div>
        </div>
        
        <!-- Step 3: Inspection & Verification -->
        <div v-if="currentStep === 3" class="fade-in">
          <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center">
            <span class="bg-dark-blue text-white w-7 h-7 rounded-full flex items-center justify-center mr-2">3</span>
            Inspection & Verification
          </h2>
          
          <div class="bg-blue-50 p-6 rounded-lg mb-6">
            <div class="text-center">
              <font-awesome-icon :icon="['fas', 'calendar-check']" class="text-dark-blue text-4xl mb-4" />
              <h3 class="text-lg font-bold text-gray-800 mb-2">Inspection Appointment</h3>
              <p class="text-gray-600 mb-6">Your vehicle needs to be inspected at an LTO-accredited MVIS facility</p>
              
              <button 
                @click="generateAppointment"
                class="bg-dark-blue text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors"
                :disabled="formData.appointmentDate !== null"
              >
                <font-awesome-icon :icon="['fas', 'calendar-plus']" class="mr-2" />
                {{ formData.appointmentDate ? 'Appointment Scheduled' : 'Schedule Appointment' }}
              </button>
            </div>
            
            <div v-if="formData.appointmentDate" class="mt-8 border-t border-blue-200 pt-6">
              <div class="bg-white p-6 rounded-lg shadow-sm">
                <div class="flex items-center justify-between mb-4">
                  <h4 class="font-bold text-gray-800">Appointment Details</h4>
                  <span class="bg-green-100 text-green-800 px-3 py-1 rounded-full text-xs font-medium">
                    Confirmed
                  </span>
                </div>
                
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <p class="text-sm text-gray-500 mb-1">Date</p>
                    <p class="font-medium">{{ formData.appointmentDate }}</p>
                  </div>
                  
                  <div>
                    <p class="text-sm text-gray-500 mb-1">Time</p>
                    <p class="font-medium">{{ formData.appointmentTime }}</p>
                  </div>
                  
                  <div class="md:col-span-2">
                    <p class="text-sm text-gray-500 mb-1">Reference Code</p>
                    <p class="font-medium bg-gray-100 p-2 rounded text-center">{{ formData.referenceCode }}</p>
                  </div>
                </div>
                
                <div class="mt-6 bg-yellow-50 border-l-4 border-yellow-400 p-4">
                  <div class="flex">
                    <div class="flex-shrink-0">
                      <font-awesome-icon :icon="['fas', 'exclamation-triangle']" class="text-yellow-400" />
                    </div>
                    <div class="ml-3">
                      <p class="text-sm text-yellow-700">
                        Please bring your vehicle and all original documents to the inspection center at the scheduled time.
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          
          <div class="bg-gray-50 p-4 rounded-lg">
            <h3 class="font-medium text-gray-800 mb-3">Inspection Process</h3>
            <ol class="list-decimal pl-5 space-y-2 text-gray-600">
              <li>Bring your vehicle to the inspection center at the scheduled time</li>
              <li>Present your reference code to the inspector</li>
              <li>Vehicle will undergo safety and emissions testing</li>
              <li>Results will be uploaded to the LTO system</li>
              <li>You will receive a notification once inspection is complete</li>
            </ol>
          </div>
        </div>
        
        <!-- Step 4: Payment & Processing -->
        <div v-if="currentStep === 4" class="fade-in">
          <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center">
            <span class="bg-dark-blue text-white w-7 h-7 rounded-full flex items-center justify-center mr-2">4</span>
            Payment & Processing
          </h2>
          
          <div class="bg-blue-50 p-6 rounded-lg mb-6">
            <div class="text-center mb-6">
              <font-awesome-icon :icon="['fas', 'receipt']" class="text-dark-blue text-4xl mb-4" />
              <h3 class="text-lg font-bold text-gray-800 mb-2">Digital Reference Slip</h3>
              <p class="text-gray-600">Your digital reference slip for LTO office visit</p>
            </div>
            
            <div class="bg-white p-6 rounded-lg shadow-sm">
              <div class="border-b pb-4 mb-4">
                <h4 class="font-bold text-gray-800 mb-2">Registration Details</h4>
                <div class="grid grid-cols-2 gap-4 text-sm">
                  <div>
                    <p class="text-gray-500">Vehicle</p>
                    <p class="font-medium">{{ formData.make }} {{ formData.model }}</p>
                  </div>
                  <div>
                    <p class="text-gray-500">Year</p>
                    <p class="font-medium">{{ formData.year }}</p>
                  </div>
                  <div>
                    <p class="text-gray-500">Engine No.</p>
                    <p class="font-medium">{{ formData.engineNumber }}</p>
                  </div>
                  <div>
                    <p class="text-gray-500">Chassis No.</p>
                    <p class="font-medium">{{ formData.chassisNumber }}</p>
                  </div>
                </div>
              </div>
              
              <div class="border-b pb-4 mb-4">
                <h4 class="font-bold text-gray-800 mb-2">Estimated Fees</h4>
                <div class="space-y-2">
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">Registration Fee</span>
                    <span class="font-medium">₱1,500.00</span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">Plate Issuance Fee</span>
                    <span class="font-medium">₱450.00</span>
                  </div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-600">Computer Fee</span>
                    <span class="font-medium">₱170.00</span>
                  </div>
                  <div class="flex justify-between text-sm font-bold pt-2 border-t">
                    <span>Total</span>
                    <span>₱2,120.00</span>
                  </div>
                </div>
              </div>
              
              <div class="bg-yellow-50 border-l-4 border-yellow-400 p-4 mb-4">
                <div class="flex">
                  <div class="flex-shrink-0">
                    <font-awesome-icon :icon="['fas', 'exclamation-triangle']" class="text-yellow-400" />
                  </div>
                  <div class="ml-3">
                    <p class="text-sm text-yellow-700">
                      <strong>Important:</strong> Please visit an LTO office with this reference slip and all original documents to complete your registration.
                    </p>
                  </div>
                </div>
              </div>
              
              <div class="text-center">
                <button class="bg-dark-blue text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors">
                  <font-awesome-icon :icon="['fas', 'download']" class="mr-2" />
                  Download Reference Slip
                </button>
              </div>
            </div>
          </div>
          
          <div class="bg-gray-50 p-4 rounded-lg">
            <h3 class="font-medium text-gray-800 mb-3">Next Steps</h3>
            <ol class="list-decimal pl-5 space-y-2 text-gray-600">
              <li>Visit any LTO office with your reference slip</li>
              <li>Submit original hard copies of all required documents</li>
              <li>Pay the registration fees</li>
              <li>Complete verification with an LTO officer</li>
              <li>Receive your official registration certificate and plate number</li>
            </ol>
          </div>
        </div>
        
        <!-- Navigation buttons -->
        <div class="mt-8 flex justify-between">
          <button 
            v-if="currentStep > 1"
            @click="prevStep"
            class="px-6 py-2 bg-gray-200 text-gray-700 rounded-lg hover:bg-gray-300 transition-colors"
          >
            <font-awesome-icon :icon="['fas', 'arrow-left']" class="mr-2" />
            Previous
          </button>
          <div v-else></div>
          
          <button 
            v-if="currentStep < totalSteps"
            @click="nextStep"
            class="px-6 py-2 bg-dark-blue text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            Next
            <font-awesome-icon :icon="['fas', 'arrow-right']" class="ml-2" />
          </button>
          <button 
            v-else
            @click="submitRegistration"
            class="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
          >
            Submit Registration
            <font-awesome-icon :icon="['fas', 'check']" class="ml-2" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.fade-in {
  animation: fadeIn 0.3s ease-in-out;
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
</style>
