<script setup lang="ts">
import { ref, computed, defineAsyncComponent } from 'vue'
import { useUserStore } from '@/stores/user'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

const Dashboard = defineAsyncComponent(() => import('@/components/lto/LTODashboard.vue'))
const Registrations = defineAsyncComponent(() => import('@/components/lto/LTORegistrations.vue'))
const Plates = defineAsyncComponent(() => import('@/components/lto/LTOPlates.vue'))
const Violations = defineAsyncComponent(() => import('@/components/lto/LTOViolations.vue'))

const userStore = useUserStore()
const activeSection = ref('dashboard')
const componentMap: Record<
  string,
  typeof Dashboard | typeof Registrations | typeof Plates | typeof Violations
> = {
  dashboard: Dashboard,
  registrations: Registrations,
  plates: Plates,
  violations: Violations,
}

const currentUser = computed(() => userStore.currentUser)

const navigateTo = (section: string) => {
  activeSection.value = section
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
        <button
          @click="navigateTo('registrations')"
          :class="[
            'w-full flex items-center py-2.5 px-4 rounded transition duration-200',
            activeSection === 'registrations' ? 'bg-blue-800' : 'hover:bg-blue-800',
            'focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50',
          ]"
        >
          <font-awesome-icon icon="clipboard" class="h-5 w-5 mr-3" />
          <span>Vehicle Registrations</span>
        </button>
        <button
          @click="navigateTo('plates')"
          :class="[
            'w-full flex items-center py-2.5 px-4 rounded transition duration-200',
            activeSection === 'plates' ? 'bg-blue-800' : 'hover:bg-blue-800',
            'focus:outline-none focus:ring-2 focus:ring-blue-600 focus:ring-opacity-50',
          ]"
        >
          <font-awesome-icon icon="car" class="h-5 w-5 mr-3" />
          <span>Plate Issuance</span>
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
      <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-100 p-6">
        <div class="container mx-auto">
          <component :is="componentMap[activeSection]" />
        </div>
      </main>
    </div>
  </div>
</template>
