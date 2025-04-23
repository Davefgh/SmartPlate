<script setup lang="ts">
import { ref, computed, defineAsyncComponent, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { useRouter } from 'vue-router'

const Dashboard = defineAsyncComponent(() => import('@/components/lto/LTODashboard.vue'))
const Registrations = defineAsyncComponent(() => import('@/components/lto/LTORegistrations.vue'))
const Plates = defineAsyncComponent(() => import('@/components/lto/LTOPlates.vue'))
const Violations = defineAsyncComponent(() => import('@/components/lto/LTOViolations.vue'))

const userStore = useUserStore()
const router = useRouter()
const activeSection = ref('dashboard')
const registrationStatus = ref('pending')
const isRegistrationDropdownOpen = ref(false)

// Ensure authentication is checked when the component is mounted
onMounted(() => {
  console.log('LTOView mounted, checking authentication...')

  // Force authentication check
  const isAuthenticated = userStore.isAuthenticated || userStore.checkAuth()
  const userRole = userStore.userRole

  console.log('Authentication state:', { isAuthenticated, userRole })

  // If not authenticated or not an LTO Officer (except admin)
  if (!isAuthenticated) {
    console.log('Not authenticated, redirecting to login')
    router.push('/login')
    return
  }

  if (userRole !== 'LTO Officer' && userRole !== 'admin') {
    console.log('Not authorized as LTO Officer, redirecting to appropriate view')
    if (userRole === 'admin') {
      router.push('/admin')
    } else {
      router.push('/home')
    }
    return
  }

  console.log('Authentication check passed, user can access LTO View')
})

const componentMap: Record<
  string,
  typeof Dashboard | typeof Registrations | typeof Plates | typeof Violations
> = {
  dashboard: Dashboard,
  'registrations-pending': Registrations,
  'registrations-processing': Registrations,
  'registrations-completed': Registrations,
  plates: Plates,
  violations: Violations,
}

const currentUser = computed(() => userStore.currentUser)

const navigateTo = (section: string, status?: string) => {
  if (status) {
    registrationStatus.value = status
    activeSection.value = `registrations-${status}`
  } else {
    activeSection.value = section
  }
  isRegistrationDropdownOpen.value = false
}

const toggleRegistrationDropdown = () => {
  isRegistrationDropdownOpen.value = !isRegistrationDropdownOpen.value
}
</script>

<template>
  <div class="flex h-screen bg-gray-50 overflow-hidden">
    <!-- Sidebar -->
    <aside class="w-64 bg-gradient-to-b from-blue-900 to-blue-800 text-white shadow-xl">
      <div class="flex items-center p-6">
        <div class="flex items-center space-x-3">
          <img
            src="/Land_Transportation_Office.webp"
            alt="LTO Logo"
            class="h-10 w-10 rounded-lg shadow-lg"
          />
          <span class="text-xl font-bold tracking-wide">LTO Portal</span>
        </div>
      </div>

      <nav class="space-y-2 px-4">
        <button
          @click="navigateTo('dashboard')"
          :class="[
            'w-full flex items-center py-2.5 px-4 rounded transition duration-200',
            activeSection === 'dashboard' ? 'bg-blue-800' : 'hover:bg-blue-800',
            'focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50',
          ]"
        >
          <font-awesome-icon icon="home" class="h-5 w-5 mr-3" />
          <span>Dashboard</span>
        </button>
        <div class="relative">
          <button
            @click="toggleRegistrationDropdown"
            :class="[
              'w-full flex items-center justify-between py-2.5 px-4 rounded transition duration-200',
              activeSection.includes('registrations') ? 'bg-blue-800' : 'hover:bg-blue-800',
              'focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50',
            ]"
          >
            <div class="flex items-center">
              <font-awesome-icon icon="file-alt" class="h-5 w-5 mr-3" />
              <span>Vehicle Registrations</span>
            </div>
            <font-awesome-icon
              :icon="isRegistrationDropdownOpen ? 'chevron-up' : 'chevron-down'"
              class="h-4 w-4 transform transition-transform duration-200"
              :class="{ 'rotate-180': isRegistrationDropdownOpen }"
            />
          </button>
          <transition
            enter-active-class="transition ease-out duration-200"
            enter-from-class="transform opacity-0 -translate-y-2"
            enter-to-class="transform opacity-100 translate-y-0"
            leave-active-class="transition ease-in duration-150"
            leave-from-class="transform opacity-100 translate-y-0"
            leave-to-class="transform opacity-0 -translate-y-2"
          >
            <div v-if="isRegistrationDropdownOpen" class="pl-8 mt-1 space-y-1">
              <button
                v-for="status in ['pending', 'processing', 'completed']"
                :key="status"
                @click="navigateTo('registrations', status)"
                :class="[
                  'w-full flex items-center py-2 px-4 rounded text-sm transition duration-200',
                  activeSection === `registrations-${status}` ? 'bg-blue-700' : 'hover:bg-blue-700',
                  'focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50',
                ]"
              >
                {{ status.charAt(0).toUpperCase() + status.slice(1) }}
              </button>
            </div>
          </transition>
        </div>
        <button
          @click="navigateTo('plates')"
          :class="[
            'w-full flex items-center py-2.5 px-4 rounded transition duration-200',
            activeSection === 'plates' ? 'bg-blue-800' : 'hover:bg-blue-800',
            'focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50',
          ]"
        >
          <font-awesome-icon icon="id-card" class="h-5 w-5 mr-3" />
          <span>Plates</span>
        </button>
        <button
          @click="navigateTo('violations')"
          :class="[
            'w-full flex items-center py-2.5 px-4 rounded transition duration-200',
            activeSection === 'violations' ? 'bg-blue-800' : 'hover:bg-blue-800',
            'focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50',
          ]"
        >
          <font-awesome-icon icon="exclamation-triangle" class="h-5 w-5 mr-3" />
          <span>Violations</span>
        </button>
      </nav>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- Header -->
      <header class="bg-white shadow-sm border-b border-gray-200">
        <div class="flex items-center justify-end px-8 py-4">
          <div class="flex items-center space-x-4">
            <div class="flex items-center space-x-3">
              <img
                :src="currentUser?.avatar"
                alt="User avatar"
                class="h-8 w-8 rounded-full object-cover"
              />
              <div class="flex flex-col">
                <span class="text-sm font-medium text-gray-800">
                  {{ currentUser?.firstName }} {{ currentUser?.lastName }}
                </span>
                <span class="text-xs text-gray-500">{{ currentUser?.role }}</span>
              </div>
            </div>
            <button
              class="bg-blue-900 text-white px-4 py-2 rounded-md hover:bg-blue-800 transition duration-200 flex items-center space-x-2"
              @click="userStore.logout()"
            >
              <font-awesome-icon icon="sign-out-alt" class="h-4 w-4" />
              <span>Logout</span>
            </button>
          </div>
        </div>
      </header>

      <!-- Main Content Area -->
      <main class="flex-1 overflow-y-auto bg-gray-50 p-6">
        <div class="container mx-auto">
          <component
            :is="componentMap[activeSection]"
            :registration-status="registrationStatus"
            v-if="componentMap[activeSection]"
          />
        </div>
      </main>
    </div>
  </div>
</template>
