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
          <div class="text-4xl font-bold text-dark-blue mb-3">12</div>
          <div class="flex items-center text-green-600">
            <font-awesome-icon :icon="['fas', 'arrow-up']" class="w-3 h-3 mr-1" />
            <p class="text-sm">+2 new this month</p>
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
          <div class="text-4xl font-bold text-gray-800 mb-3">8</div>
          <div class="flex items-center text-green-600">
            <font-awesome-icon :icon="['fas', 'check-circle']" class="w-4 h-4 mr-1" />
            <p class="text-sm">All plates are up to date</p>
          </div>
        </div>

        <!-- Stats Card 3 - Registration -->
        <div
          @click="navigateToRegistration"
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
          <div class="text-4xl font-bold text-gray-800 mb-3">2</div>
          <div class="flex items-center text-red-500">
            <font-awesome-icon :icon="['fas', 'exclamation-circle']" class="w-4 h-4 mr-1" />
            <p class="text-sm">Renewal needed within 30 days</p>
          </div>
        </div>
      </div>

      <!-- Recent Activity Section -->
      <div class="mt-8 bg-white rounded-2xl shadow-md overflow-hidden backdrop-blur-sm bg-white/90">
        <div class="bg-gradient-to-r from-dark-blue to-blue-700 px-6 py-5 text-white">
          <h2 class="text-xl font-bold flex items-center">
            <font-awesome-icon :icon="['fas', 'history']" class="w-5 h-5 mr-3" />
            Recent Activity
          </h2>
        </div>

        <div class="divide-y divide-gray-100">
          <!-- Activity Item 1 -->
          <div class="p-6 hover:bg-blue-50 transition-colors duration-200">
            <div class="flex items-start">
              <div
                class="w-12 h-12 rounded-full bg-blue-100 flex items-center justify-center text-dark-blue mr-4 shadow-sm"
              >
                <font-awesome-icon :icon="['fas', 'car']" class="w-6 h-6" />
              </div>
              <div class="flex-1">
                <div class="flex justify-between items-start">
                  <h4 class="text-base font-semibold text-dark-blue">New Vehicle Added</h4>
                  <span class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded-full"
                    >2 days ago</span
                  >
                </div>
                <p class="text-sm text-gray-600 mt-1">Toyota Camry with plate number ABC-123</p>
              </div>
            </div>
          </div>

          <!-- Activity Item 2 -->
          <div class="p-6 hover:bg-green-50 transition-colors duration-200">
            <div class="flex items-start">
              <div
                class="w-12 h-12 rounded-full bg-green-100 flex items-center justify-center text-green-600 mr-4 shadow-sm"
              >
                <font-awesome-icon :icon="['fas', 'id-card']" class="w-6 h-6" />
              </div>
              <div class="flex-1">
                <div class="flex justify-between items-start">
                  <h4 class="text-base font-semibold text-dark-blue">Plate Registration Renewed</h4>
                  <span class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded-full"
                    >1 week ago</span
                  >
                </div>
                <p class="text-sm text-gray-600 mt-1">Honda Civic with plate number XYZ-789</p>
              </div>
            </div>
          </div>

          <!-- Activity Item 3 -->
          <div class="p-6 hover:bg-amber-50 transition-colors duration-200">
            <div class="flex items-start">
              <div
                class="w-12 h-12 rounded-full bg-amber-100 flex items-center justify-center text-amber-600 mr-4 shadow-sm"
              >
                <font-awesome-icon :icon="['fas', 'file-contract']" class="w-6 h-6" />
              </div>
              <div class="flex-1">
                <div class="flex justify-between items-start">
                  <h4 class="text-base font-semibold text-dark-blue">Renewal Notification</h4>
                  <span class="text-xs text-gray-500 bg-gray-100 px-2 py-1 rounded-full"
                    >3 days ago</span
                  >
                </div>
                <p class="text-sm text-gray-600 mt-1">Ford F-150 plate expires in 30 days</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// User profile data (mock data)
const user = {
  name: 'Stanleigh Morales',
  email: 'stanleighmorales@gmail.com',
  avatar: '/Land_Transportation_Office.webp',
}

// Define emits
const emit = defineEmits(['navigate'])

// Navigation functions
const navigateToVehicles = () => {
  emit('navigate', 'Vehicles')
}

const navigateToPlates = () => {
  emit('navigate', 'Plates')
}

const navigateToRegistration = () => {
  emit('navigate', 'Registration')
}
</script>

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
  background-image: url("data:image/svg+xml,%3Csvg width='100' height='100' viewBox='0 0 100 100' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M11 18c3.866 0 7-3.134 7-7s-3.134-7-7-7-7 3.134-7 7 3.134 7 7 7zm48 25c3.866 0 7-3.134 7-7s-3.134-7-7-7-7 3.134-7 7 3.134 7 7 7zm-43-7c1.657 0 3-1.343 3-3s-1.343-3-3-3-3 1.343-3 3 1.343 3 3 3zm63 31c1.657 0 3-1.343 3-3s-1.343-3-3-3-3 1.343-3 3 1.343 3 3 3zM34 90c1.657 0 3-1.343 3-3s-1.343-3-3-3-3 1.343-3 3 1.343 3 3 3zm56-76c1.657 0 3-1.343 3-3s-1.343-3-3-3-3 1.343-3 3 1.343 3 3 3zM12 86c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm28-65c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm23-11c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zm-6 60c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm29 22c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zM32 63c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zm57-13c2.76 0 5-2.24 5-5s-2.24-5-5-5-5 2.24-5 5 2.24 5 5 5zm-9-21c1.105 0 2-.895 2-2s-.895-2-2-2-2 .895-2 2 .895 2 2 2zM60 91c1.105 0 2-.895 2-2s-.895-2-2-2-2 .895-2 2 .895 2 2 2zM35 41c1.105 0 2-.895 2-2s-.895-2-2-2-2 .895-2 2 .895 2 2 2zM12 60c1.105 0 2-.895 2-2s-.895-2-2-2-2 .895-2 2 .895 2 2 2z' fill='%23bfdbfe' fill-opacity='0.2' fill-rule='evenodd'/%3E%3C/svg%3E");
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
