<script setup>
import { ref, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'

const userStore = useUserStore()
const router = useRouter()

// Check if user is admin
onMounted(() => {
  if (!userStore.isAdmin) {
    router.push('/home')
  }
})

// Mock data for admin dashboard
const totalUsers = ref(1245)
const newRegistrations = ref(87)
const pendingApprovals = ref(32)
const activeVehicles = ref(3567)

// Mock user list
const users = ref([
  {
    id: 1,
    name: 'Stanleigh Morales',
    email: 'stanleighmorales@gmail.com',
    role: 'user',
    status: 'active',
    registrationDate: '2023-05-15',
    vehicles: 2
  },
  {
    id: 2,
    name: 'Maria Santos',
    email: 'maria.santos@example.com',
    role: 'user',
    status: 'pending',
    registrationDate: '2023-06-22',
    vehicles: 1
  },
  {
    id: 3,
    name: 'Juan Dela Cruz',
    email: 'juan.delacruz@example.com',
    role: 'user',
    status: 'active',
    registrationDate: '2023-04-10',
    vehicles: 3
  },
  {
    id: 4,
    name: 'Ana Reyes',
    email: 'ana.reyes@example.com',
    role: 'user',
    status: 'inactive',
    registrationDate: '2023-03-05',
    vehicles: 0
  },
  {
    id: 5,
    name: 'System Admin',
    email: 'admin@smartplate.com',
    role: 'admin',
    status: 'active',
    registrationDate: '2023-01-01',
    vehicles: 1
  }
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
  filteredUsers.value = users.value.filter(user => 
    user.name.toLowerCase().includes(query) || 
    user.email.toLowerCase().includes(query)
  )
}

// Status badge color
const getStatusColor = (status) => {
  switch (status) {
    case 'active':
      return 'bg-green-100 text-green-800'
    case 'pending':
      return 'bg-yellow-100 text-yellow-800'
    case 'inactive':
      return 'bg-red-100 text-red-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

// Role badge color
const getRoleColor = (role) => {
  switch (role) {
    case 'admin':
      return 'bg-purple-100 text-purple-800'
    case 'user':
      return 'bg-blue-100 text-blue-800'
    default:
      return 'bg-gray-100 text-gray-800'
  }
}

const goBack = () => {
  router.push('/home')
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <header class="bg-gradient-to-r from-dark-blue to-blue-900 text-white shadow-md">
      <div class="container mx-auto px-4 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <button
              @click="goBack"
              class="p-2 rounded-full hover:bg-blue-800 transition-colors duration-200"
            >
              <font-awesome-icon :icon="['fas', 'arrow-left']" class="w-5 h-5" />
            </button>
            <h1 class="text-xl font-bold">Admin Dashboard</h1>
          </div>
          <div class="flex items-center space-x-2">
            <span class="text-sm">Welcome, {{ userStore.fullName }}</span>
            <button
              @click="userStore.logout"
              class="p-2 rounded-full hover:bg-blue-800 transition-colors duration-200"
            >
              <font-awesome-icon :icon="['fas', 'sign-out-alt']" class="w-5 h-5" />
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="container mx-auto px-4 py-8">
      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <!-- Total Users -->
        <div class="bg-white rounded-xl shadow-md p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-500 text-sm">Total Users</p>
              <h3 class="text-2xl font-bold text-dark-blue">{{ totalUsers }}</h3>
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
              <h3 class="text-2xl font-bold text-dark-blue">{{ newRegistrations }}</h3>
              <p class="text-xs text-green-600">+12% from last month</p>
            </div>
            <div class="bg-green-100 p-3 rounded-full">
              <font-awesome-icon :icon="['fas', 'user-circle']" class="w-6 h-6 text-green-600" />
            </div>
          </div>
        </div>

        <!-- Pending Approvals -->
        <div class="bg-white rounded-xl shadow-md p-6">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-500 text-sm">Pending Approvals</p>
              <h3 class="text-2xl font-bold text-dark-blue">{{ pendingApprovals }}</h3>
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
              <h3 class="text-2xl font-bold text-dark-blue">{{ activeVehicles }}</h3>
            </div>
            <div class="bg-red-100 p-3 rounded-full">
              <font-awesome-icon :icon="['fas', 'car']" class="w-6 h-6 text-red-600" />
            </div>
          </div>
        </div>
      </div>

      <!-- User Management Section -->
      <div class="bg-white rounded-xl shadow-md overflow-hidden">
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-bold text-dark-blue">User Management</h2>
        </div>
        
        <!-- Search and Filter -->
        <div class="p-6 border-b border-gray-200">
          <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
            <div class="relative w-full md:w-96">
              <input
                type="text"
                v-model="searchQuery"
                @input="searchUsers"
                placeholder="Search users..."
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
              <font-awesome-icon
                :icon="['fas', 'search']"
                class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
              />
            </div>
            <div class="flex items-center gap-2">
              <button class="bg-dark-blue hover:bg-blue-800 text-white font-medium py-2 px-4 rounded-lg transition-colors duration-200">
                Add User
              </button>
            </div>
          </div>
        </div>
        
        <!-- User Table -->
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  User
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Role
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Status
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Registration Date
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Vehicles
                </th>
                <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="user in filteredUsers" :key="user.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="flex-shrink-0 h-10 w-10 bg-gray-200 rounded-full flex items-center justify-center">
                      <font-awesome-icon :icon="['fas', 'user']" class="text-gray-500" />
                    </div>
                    <div class="ml-4">
                      <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                      <div class="text-sm text-gray-500">{{ user.email }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="[getRoleColor(user.role), 'px-2 py-1 text-xs rounded-full']">
                    {{ user.role }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="[getStatusColor(user.status), 'px-2 py-1 text-xs rounded-full']">
                    {{ user.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ user.registrationDate }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ user.vehicles }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                  <button class="text-blue-600 hover:text-blue-900 mr-3">
                    <font-awesome-icon :icon="['fas', 'edit']" />
                  </button>
                  <button class="text-red-600 hover:text-red-900">
                    <font-awesome-icon :icon="['fas', 'trash-alt']" />
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        
        <!-- Pagination -->
        <div class="px-6 py-4 border-t border-gray-200 flex items-center justify-between">
          <div class="text-sm text-gray-700">
            Showing <span class="font-medium">1</span> to <span class="font-medium">{{ filteredUsers.length }}</span> of <span class="font-medium">{{ users.length }}</span> results
          </div>
          <div class="flex items-center space-x-2">
            <button class="bg-white border border-gray-300 rounded-md px-3 py-1 text-sm font-medium text-gray-700 hover:bg-gray-50">
              Previous
            </button>
            <button class="bg-white border border-gray-300 rounded-md px-3 py-1 text-sm font-medium text-gray-700 hover:bg-gray-50">
              Next
            </button>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
