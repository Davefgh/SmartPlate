<script setup lang="ts">
import { ref, computed, defineAsyncComponent } from 'vue'
import { useUserStore } from '@/stores/user'

const Dashboard = defineAsyncComponent(() => import('@/components/lto/Dashboard.vue'))
const Registrations = defineAsyncComponent(() => import('@/components/lto/Registrations.vue'))
const Plates = defineAsyncComponent(() => import('@/components/lto/Plates.vue'))
const Violations = defineAsyncComponent(() => import('@/components/lto/Violations.vue'))

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
          <svg class="h-5 w-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
            />
          </svg>
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
          <svg class="h-5 w-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
            />
          </svg>
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
          <svg class="h-5 w-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"
            />
          </svg>
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
          <svg class="h-5 w-5 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
            />
          </svg>
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
              <svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                />
              </svg>
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
