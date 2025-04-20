<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { useVehicleRegistrationStore } from '@/stores/vehicleRegistration'
import { useVehicleRegistrationFormStore } from '@/stores/vehicleRegistrationForm'
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import type { Vehicle, Plate } from '@/types/vehicle'

interface User {
  name: string
  email: string
  avatar: string
}

const router = useRouter()
const userStore = useUserStore()
const vehicleStore = useVehicleRegistrationStore()
const registrationFormStore = useVehicleRegistrationFormStore()

const user = computed<User>(
  (): User => ({
    name: userStore.fullName,
    email: userStore.currentUser?.email || '',
    avatar: userStore.currentUser?.avatar || '/Land_Transportation_Office.webp',
  }),
)

const totalVehicles = computed(() => vehicleStore.userVehicles.length)
const newVehiclesThisMonth = computed(() => {
  const today = new Date()
  const firstDayOfMonth = new Date(today.getFullYear(), today.getMonth(), 1)
  return vehicleStore.userVehicles.filter((vehicle: Vehicle) => {
    const lastUpdated = new Date(vehicle.lastRenewalDate)
    return lastUpdated >= firstDayOfMonth
  }).length
})

const activePlates = computed(
  () => vehicleStore.userPlates.filter((plate: Plate) => plate.status === 'Active').length,
)
const allPlatesUpToDate = computed(() => {
  const today = new Date()
  return vehicleStore.userPlates.every((plate: Plate) => {
    const expiryDate = new Date(plate.plate_expiration_date)
    return expiryDate > today
  })
})

const pendingRenewals = computed(() => vehicleStore.soonToExpireRegistrations.length)

// Get in-progress vehicle registration forms
const inProgressForms = computed(() => {
  const userEmail = userStore.currentUser?.email || ''
  return registrationFormStore.forms.filter(
    (form) => form.userId === userEmail && form.status === 'pending',
  )
})

// Check if there is an active registration in progress
const hasActiveRegistration = computed(() => inProgressForms.value.length > 0)

// Get the most recent in-progress form
const latestInProgressForm = computed(() => {
  if (inProgressForms.value.length === 0) return null

  // Sort by id which contains timestamp (newest first)
  return inProgressForms.value.sort((a, b) => {
    const timeA = Number(a.id.substring(4))
    const timeB = Number(b.id.substring(4))
    return timeB - timeA
  })[0]
})

// Navigation functions
const emit = defineEmits(['navigate'])

const navigateToVehicles = () => {
  emit('navigate', 'Vehicles')
}

const navigateToPlates = () => {
  emit('navigate', 'Plates')
}

const navigateToVehicleRegistrationForm = () => {
  // If there's an active registration, load the most recent one
  if (hasActiveRegistration.value && latestInProgressForm.value) {
    registrationFormStore.loadForm(latestInProgressForm.value.id)
  }
  router.push('/vehicle-registration')
}

const continueRegistration = (formId: string) => {
  registrationFormStore.loadForm(formId)
  router.push('/vehicle-registration')
}

// Function to handle the registration button click
const handleRegistrationButtonClick = () => {
  if (hasActiveRegistration.value && latestInProgressForm.value) {
    // If there's a pending registration, load that form
    registrationFormStore.loadForm(latestInProgressForm.value.id)
    router.push('/vehicle-registration')
  } else {
    // Start a new registration
    router.push('/vehicle-registration')
  }
}
</script>

<template>
  <div class="dashboard-content p-6">
    <div class="max-w-7xl mx-auto">
      <!-- Welcome Banner -->
      <div
        class="bg-gradient-to-r from-dark-blue to-light-blue rounded-2xl p-6 mb-8 text-white shadow-lg transform transition-all duration-500 hover:scale-[1.01] hover:shadow-xl"
      >
        <div class="flex flex-col md:flex-row items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold mb-2">Welcome back, {{ user.name.split(' ')[0] }}!</h1>
            <p class="text-blue-100">Track and manage your vehicle plates with ease.</p>
          </div>
          <div class="mt-4 md:mt-0">
            <button
              @click="handleRegistrationButtonClick"
              class="bg-white text-dark-blue px-6 py-3 rounded-lg font-medium hover:bg-blue-50 transition-colors shadow-md flex items-center"
              :class="{ 'bg-amber-100 hover:bg-amber-200': hasActiveRegistration }"
            >
              <font-awesome-icon
                :icon="['fas', hasActiveRegistration ? 'clipboard-list' : 'car-side']"
                class="mr-2"
              />
              {{
                hasActiveRegistration ? 'View Registration Progress' : 'Add Vehicle Registration'
              }}
            </button>
          </div>
        </div>
      </div>

      <!-- Dashboard Content -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-10">
        <!-- Stats Card 1 - Vehicles -->
        <div
          @click="navigateToVehicles"
          class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-2xl shadow-md p-8 transform transition-all duration-300 hover:shadow-xl hover:-translate-y-2 border-t-4 border-dark-blue cursor-pointer"
        >
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-xl font-bold text-dark-blue">Total Vehicles</h3>
            <div
              class="w-14 h-14 rounded-full bg-dark-blue bg-opacity-10 flex items-center justify-center text-dark-blue"
            >
              <font-awesome-icon :icon="['fas', 'car']" class="w-7 h-7" />
            </div>
          </div>
          <div class="text-4xl font-bold text-dark-blue mb-3">{{ totalVehicles }}</div>
          <div class="flex items-center text-green-600">
            <font-awesome-icon :icon="['fas', 'arrow-up']" class="w-3 h-3 mr-1" />
            <p class="text-sm">+{{ newVehiclesThisMonth }} new this month</p>
          </div>
        </div>

        <!-- Stats Card 2 - Plates -->
        <div
          @click="navigateToPlates"
          class="bg-gradient-to-br from-green-50 to-green-100 rounded-2xl shadow-md p-8 transform transition-all duration-300 hover:shadow-xl hover:-translate-y-2 border-t-4 border-green-600 cursor-pointer"
        >
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-xl font-bold text-gray-800">Active Plates</h3>
            <div
              class="w-14 h-14 rounded-full bg-green-600 bg-opacity-10 flex items-center justify-center text-green-600"
            >
              <font-awesome-icon :icon="['fas', 'id-card']" class="w-7 h-7" />
            </div>
          </div>
          <div class="text-4xl font-bold text-gray-800 mb-3">{{ activePlates }}</div>
          <div
            class="flex items-center"
            :class="allPlatesUpToDate ? 'text-green-600' : 'text-amber-600'"
          >
            <font-awesome-icon
              :icon="['fas', allPlatesUpToDate ? 'check-circle' : 'exclamation-circle']"
              class="w-4 h-4 mr-1"
            />
            <p class="text-sm">
              {{ allPlatesUpToDate ? 'All plates are up to date' : 'Some plates need renewal' }}
            </p>
          </div>
        </div>

        <!-- Stats Card 3 - Registration -->
        <div
          @click="emit('navigate', 'Registration')"
          class="bg-gradient-to-br from-amber-50 to-amber-100 rounded-2xl shadow-md p-8 transform transition-all duration-300 hover:shadow-xl hover:-translate-y-2 border-t-4 border-amber-600 cursor-pointer"
        >
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-xl font-bold text-gray-800">Pending Renewals</h3>
            <div
              class="w-14 h-14 rounded-full bg-amber-600 bg-opacity-10 flex items-center justify-center text-amber-600"
            >
              <font-awesome-icon :icon="['fas', 'file-contract']" class="w-7 h-7" />
            </div>
          </div>
          <div class="text-4xl font-bold text-gray-800 mb-3">{{ pendingRenewals }}</div>
          <div
            class="flex items-center"
            :class="pendingRenewals > 0 ? 'text-red-500' : 'text-green-600'"
          >
            <font-awesome-icon
              :icon="['fas', pendingRenewals > 0 ? 'exclamation-circle' : 'check-circle']"
              class="w-4 h-4 mr-1"
            />
            <p class="text-sm">
              {{ pendingRenewals > 0 ? 'Renewal needed within 30 days' : 'No pending renewals' }}
            </p>
          </div>
        </div>
      </div>

      <!-- Pending Registration Process Section (replaces Recent Activity) -->
      <div class="mt-8 bg-white rounded-2xl shadow-md overflow-hidden backdrop-blur-sm bg-white/90">
        <div class="bg-gradient-to-r from-amber-600 to-amber-500 px-6 py-5 text-white">
          <h2 class="text-xl font-bold flex items-center">
            <font-awesome-icon :icon="['fas', 'clipboard-list']" class="w-5 h-5 mr-3" />
            Pending Registration Process
          </h2>
        </div>

        <div class="divide-y divide-gray-100">
          <!-- In-Progress Registration Items -->
          <div
            v-for="(form, index) in inProgressForms"
            :key="index"
            class="p-6 transition-colors duration-200 hover:bg-amber-50"
          >
            <div class="flex items-start">
              <div
                class="w-12 h-12 rounded-full flex items-center justify-center mr-4 shadow-sm bg-amber-100 text-amber-600"
              >
                <font-awesome-icon :icon="['fas', 'car']" class="w-6 h-6" />
              </div>
              <div class="flex-1">
                <div class="flex justify-between items-start">
                  <h4 class="text-base font-semibold text-dark-blue">
                    {{ form.make ? form.make + ' ' + form.model : 'Vehicle' }} Registration
                  </h4>
                  <span class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded-full">
                    Started on
                    {{
                      new Date(
                        parseInt(form.id.split('-')[1] || form.id.substring(4)),
                      ).toLocaleDateString()
                    }}
                  </span>
                </div>
                <p class="text-sm text-gray-600 mt-1">
                  {{ form.isNewVehicle ? 'New Vehicle' : 'Used Vehicle' }} |
                  {{ form.vehicleType || 'Not specified' }} |
                  <span
                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-yellow-100 text-yellow-800"
                  >
                    <font-awesome-icon :icon="['fas', 'clock']" class="mr-1 w-3 h-3" />
                    {{ form.status === 'pending' ? 'Pending Approval' : form.status }}
                  </span>
                </p>
                <div class="mt-3 flex gap-2">
                  <button
                    @click="continueRegistration(form.id)"
                    class="px-4 py-2 bg-amber-600 text-white rounded-lg hover:bg-amber-700 transition-colors text-sm flex items-center"
                  >
                    <font-awesome-icon :icon="['fas', 'arrow-right']" class="w-3 h-3 mr-2" />
                    Continue Registration
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Empty state if no pending registrations -->
          <div v-if="inProgressForms.length === 0" class="p-6 text-center text-gray-500">
            <div class="flex flex-col items-center justify-center py-8">
              <font-awesome-icon
                :icon="['fas', 'clipboard-check']"
                class="w-12 h-12 text-gray-300 mb-4"
              />
              <p>No pending registration processes</p>
              <button
                @click="handleRegistrationButtonClick"
                class="mt-4 px-4 py-2 bg-dark-blue text-white rounded-lg hover:bg-blue-700 transition-colors text-sm flex items-center"
              >
                <font-awesome-icon :icon="['fas', 'plus']" class="w-3 h-3 mr-2" />
                Start New Registration
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard-content {
  background-image:
    radial-gradient(circle at 20% 35%, rgba(76, 29, 149, 0.05) 0%, transparent 50%),
    radial-gradient(circle at 75% 44%, rgba(14, 165, 233, 0.05) 0%, transparent 50%),
    linear-gradient(
      135deg,
      rgba(240, 249, 255, 0.7) 0%,
      rgba(249, 250, 251, 0.8) 50%,
      rgba(243, 244, 246, 0.7) 100%
    );
  min-height: calc(100vh - 60px);
  position: relative;
  overflow: hidden;
}

.dashboard-content::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: url("data:image/svg+xml,%3Csvg width='100' height='100' viewBox='0 0 100 100' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M11 18c3.866 0 7-3.134 7-7s-3.134-7-7-7 3.134-7 7 3.134 7 7 7zm48 25c3.866 0 7-3.134 7-7s-3.134-7-7-7 3.134-7 7 3.134 7 7 7zm-43-7c1.657 0 3-1.343 3-3s-.895-3-2-3-3 1.343-3 3 1.343 3 3 3zm63 31c1.657 0 3-1.343 3-3s-.895-3-2-3-3 1.343-3 3 1.343 3 3 3zM34 90c1.657 0 3-1.343 3-3s-.895-3-2-3-3 1.343-3 3 1.343 3 3 3zm56-76c1.657 0 3-1.343 3-3s-.895-3-2-3-3 1.343-3 3 1.343 3 3 3zM12 86c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm28-65c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm23-11c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zm-6 60c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm29 22c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zM32 63c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zm57-13c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zm-9-21c1.105 0 2-.895 2-2s-.895-2-2-2-2 .895-2 2 .895 2 2 2zM60 91c1.105 0 2-.895 2-2s-.895-2-2-2-2 .895-2 2 .895 2 2 2zM35 41c1.105 0 2-.895 2-2s-.895-2-2-2-2 .895-2 2 .895 2 2 2zM12 60c1.105 0 2-.895 2-2s-.895-2-2-2-2 .895-2 2 .895 2 2 2z' fill='%23bfdbfe' fill-opacity='0.2' fill-rule='evenodd'/%3E%3C/svg%3E");
  opacity: 0.5;
  z-index: -1;
}

.dashboard-content::after {
  content: '';
  position: absolute;
  top: -50%;
  right: -50%;
  width: 90%;
  height: 90%;
  background: radial-gradient(
    circle,
    rgba(56, 189, 248, 0.03) 0%,
    rgba(59, 130, 246, 0.02) 50%,
    transparent 70%
  );
  border-radius: 50%;
  z-index: -1;
}
</style>
