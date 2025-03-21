<script setup>
import { ref, onMounted, watch, defineAsyncComponent, computed } from 'vue'
import { useVehicleRegistrationStore } from '@/stores/vehicleRegistration'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'

const LogoutModal = defineAsyncComponent(() => import('@/components/modals/LogoutModal.vue'))
const UsersSection = defineAsyncComponent(() => import('@/components/admin/UsersSection.vue'))
const VehiclesSection = defineAsyncComponent(() => import('@/components/admin/VehiclesSection.vue'))
const PlatesSection = defineAsyncComponent(() => import('@/components/admin/PlatesSection.vue'))
const RegistrationsSection = defineAsyncComponent(
  () => import('@/components/admin/RegistrationsSection.vue'),
)
const PendingRegistrationsSection = defineAsyncComponent(
  () => import('@/components/admin/PendingRegistrationsSection.vue'),
)

const userStore = useUserStore()
const vehicleRegistrationStore = useVehicleRegistrationStore()
const router = useRouter()

// Logout modal state
const showLogoutModal = ref(false)

const confirmLogout = () => {
  showLogoutModal.value = true
}

const handleLogout = () => {
  userStore.logout()
  showLogoutModal.value = false
}

const cancelLogout = () => {
  showLogoutModal.value = false
}

// Check if user is admin
onMounted(() => {
  if (!userStore.isAdmin) {
    router.push('/home')
  }
})

// Sidebar state
const isSidebarOpen = ref(true)
const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}

// Active section management
const activeSection = ref('users')

// Dynamic data from stores
const totalUsers = computed(
  () => userStore.mockUsers.filter((user) => user.role !== 'admin').length,
)
const newRegistrations = computed(() => vehicleRegistrationStore.activeRegistrations.length)
const pendingApprovals = computed(() => vehicleRegistrationStore.pendingRegistrations.length)
const activeVehicles = computed(
  () => vehicleRegistrationStore.vehicles.filter((v) => v.status === 'Active').length,
)

// Mock user list
const users = ref([
  {
    id: 1,
    name: 'Stanleigh Morales',
    email: 'stanleighmorales@gmail.com',
    role: 'user',
    status: 'active',
    registrationDate: '2023-05-15',
    vehicles: 2,
  },
  {
    id: 2,
    name: 'Maria Santos',
    email: 'maria.santos@example.com',
    role: 'user',
    status: 'pending',
    registrationDate: '2023-06-22',
    vehicles: 1,
  },
  {
    id: 3,
    name: 'Juan Dela Cruz',
    email: 'juan.delacruz@example.com',
    role: 'user',
    status: 'active',
    registrationDate: '2023-04-10',
    vehicles: 3,
  },
  {
    id: 4,
    name: 'Ana Reyes',
    email: 'ana.reyes@example.com',
    role: 'user',
    status: 'inactive',
    registrationDate: '2023-03-05',
    vehicles: 0,
  },
  {
    id: 5,
    name: 'System Admin',
    email: 'admin@smartplate.com',
    role: 'admin',
    status: 'active',
    registrationDate: '2023-01-01',
    vehicles: 1,
  },
])

// Search functionality
const searchQuery = ref('')
const filteredUsers = ref([...users.value])

const searchUsers = () => {
  if (!searchQuery.value) {
    filteredUsers.value = [...users.value]
    return
  }

  const query = searchQuery.value.toLowerCase()
  filteredUsers.value = users.value.filter(
    (user) => user.name.toLowerCase().includes(query) || user.email.toLowerCase().includes(query),
  )
}

// Add watch effect to trigger search
watch(searchQuery, () => {
  searchUsers()
})
</script>

<template>
  <div class="flex h-screen bg-gray-100 overflow-hidden">
    <!-- Sidebar -->
    <div
      :class="[
        'bg-white shadow-lg transition-all duration-300 ease-in-out',
        isSidebarOpen ? 'w-64' : 'w-20',
      ]"
    >
      <div class="p-4">
        <div class="flex items-center justify-between mb-6">
          <button @click="toggleSidebar" class="p-2 rounded-lg hover:bg-gray-100">
            <font-awesome-icon :icon="['fas', 'bars']" />
          </button>
        </div>

        <!-- Navigation Menu -->
        <nav class="space-y-2">
          <button
            @click="activeSection = 'users'"
            :class="[
              'w-full flex items-center p-2 rounded-lg transition-colors',
              activeSection === 'users' ? 'bg-blue-100 text-blue-800' : 'hover:bg-gray-100',
            ]"
          >
            <font-awesome-icon :icon="['fas', 'users']" class="w-5 h-5" />
            <span v-if="isSidebarOpen" class="ml-3">Users</span>
          </button>

          <button
            @click="activeSection = 'vehicles'"
            :class="[
              'w-full flex items-center p-2 rounded-lg transition-colors',
              activeSection === 'vehicles' ? 'bg-blue-100 text-blue-800' : 'hover:bg-gray-100',
            ]"
          >
            <font-awesome-icon :icon="['fas', 'car']" class="w-5 h-5" />
            <span v-if="isSidebarOpen" class="ml-3">Vehicles</span>
          </button>

          <button
            @click="activeSection = 'plates'"
            :class="[
              'w-full flex items-center p-2 rounded-lg transition-colors',
              activeSection === 'plates' ? 'bg-blue-100 text-blue-800' : 'hover:bg-gray-100',
            ]"
          >
            <font-awesome-icon :icon="['fas', 'id-card']" class="w-5 h-5" />
            <span v-if="isSidebarOpen" class="ml-3">Plates</span>
          </button>

          <button
            @click="activeSection = 'registrations'"
            :class="[
              'w-full flex items-center p-2 rounded-lg transition-colors',
              activeSection === 'registrations' ? 'bg-blue-100 text-blue-800' : 'hover:bg-gray-100',
            ]"
          >
            <font-awesome-icon :icon="['fas', 'clipboard-list']" class="w-5 h-5" />
            <span v-if="isSidebarOpen" class="ml-3">Registrations</span>
          </button>

          <button
            @click="activeSection = 'pending-registrations'"
            :class="[
              'w-full flex items-center p-2 rounded-lg transition-colors',
              activeSection === 'pending-registrations'
                ? 'bg-blue-100 text-blue-800'
                : 'hover:bg-gray-100',
            ]"
          >
            <font-awesome-icon :icon="['fas', 'clock']" class="w-5 h-5" />
            <span v-if="isSidebarOpen" class="ml-3">Pending Registrations</span>
          </button>

          <!-- Logout Button -->
          <button
            @click="confirmLogout"
            class="w-full flex items-center p-2 rounded-lg text-red hover:bg-red hover:bg-opacity-10 transition-colors mt-8"
          >
            <font-awesome-icon :icon="['fas', 'sign-out-alt']" class="w-5 h-5" />
            <span v-if="isSidebarOpen" class="ml-3">Log Out</span>
          </button>
        </nav>
      </div>
    </div>

    <!-- Logout Modal -->
    <LogoutModal :show="showLogoutModal" @confirm="handleLogout" @cancel="cancelLogout" />

    <!-- Main Content -->
    <div class="flex-1 overflow-auto">
      <!-- Header -->
      <header class="bg-white shadow-sm">
        <div class="container mx-auto px-4 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3">
              <h1 class="text-xl font-bold">Admin Dashboard</h1>
            </div>
            <div class="flex items-center space-x-4">
              <span class="text-sm">Welcome, {{ userStore.fullName }}</span>
            </div>
          </div>
        </div>
      </header>

      <!-- Dashboard Content -->
      <main class="p-6">
        <!-- Stats Grid -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-6">
          <!-- Total Users -->
          <div class="bg-white rounded-xl shadow-md p-6">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-gray-500 text-sm">Total Users</p>
                <h3 class="text-2xl font-bold text-gray-900">{{ totalUsers }}</h3>
              </div>
              <div class="bg-blue-100 p-3 rounded-full">
                <font-awesome-icon :icon="['fas', 'users']" class="w-6 h-6 text-blue-600" />
              </div>
            </div>
          </div>

          <!-- New Registrations -->
          <div class="bg-white rounded-xl shadow-md p-6">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-gray-500 text-sm">New Registrations</p>
                <h3 class="text-2xl font-bold text-gray-900">{{ newRegistrations }}</h3>
              </div>
              <div class="bg-green-100 p-3 rounded-full">
                <font-awesome-icon
                  :icon="['fas', 'clipboard-list']"
                  class="w-6 h-6 text-green-600"
                />
              </div>
            </div>
          </div>

          <!-- Pending Approvals -->
          <div class="bg-white rounded-xl shadow-md p-6">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-gray-500 text-sm">Pending Approvals</p>
                <h3 class="text-2xl font-bold text-gray-900">{{ pendingApprovals }}</h3>
              </div>
              <div class="bg-yellow-100 p-3 rounded-full">
                <font-awesome-icon :icon="['fas', 'clock']" class="w-6 h-6 text-yellow-600" />
              </div>
            </div>
          </div>

          <!-- Active Vehicles -->
          <div class="bg-white rounded-xl shadow-md p-6">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-gray-500 text-sm">Active Vehicles</p>
                <h3 class="text-2xl font-bold text-gray-900">{{ activeVehicles }}</h3>
              </div>
              <div class="bg-red-100 p-3 rounded-full">
                <font-awesome-icon :icon="['fas', 'car']" class="w-6 h-6 text-red-600" />
              </div>
            </div>
          </div>
        </div>

        <!-- Dynamic Section Content -->
        <component
          :is="
            activeSection === 'users'
              ? UsersSection
              : activeSection === 'vehicles'
                ? VehiclesSection
                : activeSection === 'plates'
                  ? PlatesSection
                  : activeSection === 'pending-registrations'
                    ? PendingRegistrationsSection
                    : RegistrationsSection
          "
        />
      </main>
    </div>
  </div>
</template>
