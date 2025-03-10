<script setup>
import { ref, onMounted, onUnmounted, computed, defineAsyncComponent } from 'vue'
const DashboardContent = defineAsyncComponent(() => import('@/components/DashboardContent.vue'))
const VehicleInformation = defineAsyncComponent(() => import('@/components/VehicleInformation.vue'))
const PlateInformation = defineAsyncComponent(() => import('@/components/PlateInformation.vue'))
const RegistrationContent = defineAsyncComponent(
  () => import('@/components/RegistrationContent.vue'),
)
// Sidebar state
const isSidebarOpen = ref(false)
const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}

// User profile data (mock data)
const user = {
  name: 'Stanleigh Morales',
  email: 'stanleighmorales@gmail.com',
  avatar: '/Land_Transportation_Office.webp',
}

// Notifications (mock data)
const notifications = ref(3)
const notificationItems = ref([
  { id: 1, title: 'New plate available', message: 'Your new plate is ready for pickup', time: '2 hours ago', read: false },
  { id: 2, title: 'Registration reminder', message: 'Your vehicle registration expires in 7 days', time: '1 day ago', read: false },
  { id: 3, title: 'System maintenance', message: 'Scheduled maintenance on Saturday', time: '3 days ago', read: true }
])

// Notification dropdown state
const isNotificationDropdownOpen = ref(false)

// Profile dropdown state
const isProfileDropdownOpen = ref(false)

// Active content management
const activeMenuItem = ref('Dashboard')

// Sidebar menu items
const menuItems = ref([
  { name: 'Dashboard', icon: 'home', active: true },
  { name: 'Vehicles', icon: 'car', active: false },
  { name: 'Plates', icon: 'id-card', active: false },
  { name: 'Registration', icon: 'file-contract', active: false },
])

// Set active menu item
const setActiveMenuItem = (itemName) => {
  menuItems.value.forEach((item) => {
    item.active = item.name === itemName
  })
  activeMenuItem.value = itemName
  // Close sidebar on mobile after selection
  if (window.innerWidth < 1024) {
    isSidebarOpen.value = false
  }
}
// Component mapping
const componentMap = {
  Dashboard: DashboardContent,
  Vehicles: VehicleInformation,
  Plates: PlateInformation,
  Registration: RegistrationContent,
}

// Active component computed property
const activeComponent = computed(() => componentMap[activeMenuItem.value] || DashboardContent)

// Close sidebar on escape key
const handleEscKey = (event) => {
  if (event.key === 'Escape') {
    if (isSidebarOpen.value) {
      isSidebarOpen.value = false
    }
    if (isProfileDropdownOpen.value) {
      isProfileDropdownOpen.value = false
    }
    if (isNotificationDropdownOpen.value) {
      isNotificationDropdownOpen.value = false
    }
  }
}

// Close sidebar and dropdown on outside click
const handleOutsideClick = (event) => {
  const sidebar = document.getElementById('sidebar')
  const sidebarToggle = document.getElementById('sidebar-toggle')
  const profileDropdown = document.getElementById('profile-dropdown')
  const profileToggle = document.getElementById('profile-toggle')
  const notificationDropdown = document.getElementById('notification-dropdown')
  const notificationToggle = document.getElementById('notification-toggle')

  // Handle sidebar outside click
  if (
    isSidebarOpen.value &&
    sidebar &&
    !sidebar.contains(event.target) &&
    sidebarToggle &&
    !sidebarToggle.contains(event.target)
  ) {
    isSidebarOpen.value = false
  }

  // Handle profile dropdown outside click
  if (
    isProfileDropdownOpen.value &&
    profileDropdown &&
    !profileDropdown.contains(event.target) &&
    profileToggle &&
    !profileToggle.contains(event.target)
  ) {
    isProfileDropdownOpen.value = false
  }

  // Handle notification dropdown outside click
  if (
    isNotificationDropdownOpen.value &&
    notificationDropdown &&
    !notificationDropdown.contains(event.target) &&
    notificationToggle &&
    !notificationToggle.contains(event.target)
  ) {
    isNotificationDropdownOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleEscKey)
  document.addEventListener('click', handleOutsideClick)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleEscKey)
  document.removeEventListener('click', handleOutsideClick)
})
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Overlay for mobile when sidebar is open -->
    <div
      v-if="isSidebarOpen"
      class="fixed inset-0 bg-black bg-opacity-50 z-20 lg:hidden transition-opacity duration-300 ease-in-out"
      @click="toggleSidebar"
    ></div>

    <!-- Sidebar -->
    <aside
      id="sidebar"
      :class="[
        'fixed top-0 left-0 z-30 h-screen bg-gradient-to-b from-dark-blue to-blue-900 text-white w-64 transition-all duration-300 ease-in-out transform',
        isSidebarOpen ? 'translate-x-0' : '-translate-x-full',
      ]"
    >
      <!-- Sidebar Header -->
      <div class="flex items-center justify-between p-4 border-b border-blue-800">
        <div class="flex items-center space-x-3">
          <div class="w-10 h-10 rounded-full bg-red flex items-center justify-center">
            <span class="text-white font-bold text-xl">SP</span>
          </div>
          <h1 class="text-xl font-bold">SmartPlate</h1>
        </div>
        <button
          class="lg:hidden text-white hover:text-red transition-colors duration-200"
          @click="toggleSidebar"
        >
          <font-awesome-icon :icon="['fas', 'xmark']" class="w-5 h-5" />
        </button>
      </div>

      <!-- Sidebar Menu -->
      <nav class="mt-6 px-4">
        <ul class="space-y-2">
          <li v-for="(item, index) in menuItems" :key="index">
            <a
              href="#"
              :class="[
                'flex items-center p-3 rounded-lg transition-all duration-200 group hover:bg-blue-800',
                item.active ? 'bg-blue-800' : '',
              ]"
              @click.prevent="setActiveMenuItem(item.name)"
            >
              <div
                class="w-8 h-8 flex items-center justify-center text-blue-300 group-hover:text-white transition-colors duration-200"
              >
                <font-awesome-icon :icon="['fas', item.icon]" class="w-5 h-5" />
              </div>
              <span
                class="ml-3 font-medium group-hover:text-white transition-colors duration-200"
                >{{ item.name }}</span
              >
              <span
                class="ml-auto transform transition-transform duration-200 group-hover:translate-x-1"
              >
                <font-awesome-icon
                  :icon="['fas', 'chevron-right']"
                  class="w-3 h-3 opacity-0 group-hover:opacity-100"
                />
              </span>
            </a>
          </li>

          <!-- Logout Button -->
          <li class="mt-8">
            <a
              href="#"
              class="flex items-center p-3 rounded-lg text-red hover:bg-red hover:bg-opacity-10 transition-all duration-200 group"
            >
              <div class="w-8 h-8 flex items-center justify-center">
                <font-awesome-icon :icon="['fas', 'sign-out-alt']" class="w-5 h-5" />
              </div>
              <span class="ml-3 font-medium">Log Out</span>
            </a>
          </li>
        </ul>
      </nav>

      <!-- Sidebar Footer -->
      <div class="absolute bottom-0 left-0 right-0 p-4 border-t border-blue-800">
        <div class="flex items-center space-x-3">
          <img
            :src="user.avatar"
            alt="User Avatar"
            class="w-10 h-10 rounded-full border-2 border-light-blue"
          />
          <div class="overflow-hidden">
            <h4 class="text-sm font-semibold truncate">{{ user.name }}</h4>
            <p class="text-xs text-blue-300 truncate">{{ user.email }}</p>
          </div>
        </div>
      </div>
    </aside>

    <!-- Main Content -->
    <main
      :class="[
        'min-h-screen transition-all duration-300 ease-in-out',
        isSidebarOpen ? 'ml-64' : '',
      ]"
    >
      <!-- Header -->
      <header class="bg-white shadow-sm sticky top-0 z-10">
        <div class="flex items-center justify-between px-4 py-3">
          <!-- Left: Menu Toggle & Search -->
          <div class="flex items-center space-x-4">
            <button
              id="sidebar-toggle"
              class="p-2 rounded-lg text-gray hover:bg-gray-100 transition-colors duration-200 focus:outline-none"
              @click="toggleSidebar"
            >
              <font-awesome-icon :icon="['fas', 'bars']" class="w-5 h-5" />
            </button>

            <div class="relative hidden md:block">
              <input
                type="text"
                placeholder="Search..."
                class="w-64 pl-10 pr-4 py-2 rounded-lg border border-gray-200 focus:outline-none focus:ring-2 focus:ring-dark-blue/20 transition-all"
              />
              <div class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray">
                <font-awesome-icon :icon="['fas', 'search']" class="w-4 h-4" />
              </div>
            </div>
          </div>

          <!-- Right: Notifications & Profile -->
          <div class="flex items-center space-x-4">
            <!-- Notifications -->
            <div class="relative">
              <button
                id="notification-toggle"
                class="p-2 rounded-full text-gray hover:bg-gray-100 transition-colors duration-200 focus:outline-none"
                @click="isNotificationDropdownOpen = !isNotificationDropdownOpen"
              >
                <font-awesome-icon :icon="['fas', 'bell']" class="w-5 h-5" />
                <span
                  v-if="notifications > 0"
                  class="absolute top-0 right-0 w-4 h-4 bg-red rounded-full text-white text-xs flex items-center justify-center transform translate-x-1 -translate-y-1 animate-pulse"
                >
                  {{ notifications }}
                </span>
              </button>
              
              <!-- Notification Dropdown -->
              <div
                id="notification-dropdown"
                v-if="isNotificationDropdownOpen"
                class="absolute right-0 mt-2 w-80 bg-white rounded-lg shadow-lg py-2 z-50 border border-gray-200 overflow-hidden"
                style="animation: fadeIn 0.2s ease-out;"
              >
                <div class="px-4 py-2 border-b border-gray-100 flex items-center justify-between">
                  <h3 class="font-semibold text-gray-800">Notifications</h3>
                  <span class="text-xs text-blue-600 hover:text-blue-800 cursor-pointer">Mark all as read</span>
                </div>
                
                <div class="max-h-80 overflow-y-auto">
                  <div v-if="notificationItems.length === 0" class="py-4 text-center text-gray-500">
                    No notifications
                  </div>
                  
                  <div
                    v-for="item in notificationItems"
                    :key="item.id"
                    :class="['px-4 py-3 hover:bg-gray-50 transition-colors duration-150 cursor-pointer border-l-4', 
                             item.read ? 'border-transparent' : 'border-blue-500']"
                  >
                    <div class="flex justify-between items-start">
                      <h4 class="font-medium text-gray-800 text-sm">{{ item.title }}</h4>
                      <span class="text-xs text-gray-500">{{ item.time }}</span>
                    </div>
                    <p class="text-sm text-gray-600 mt-1">{{ item.message }}</p>
                  </div>
                </div>
                
                <div class="px-4 py-2 border-t border-gray-100 text-center">
                  <router-link 
                    to="/notifications" 
                    class="text-sm text-blue-600 hover:text-blue-800"
                    @click="isNotificationDropdownOpen = false"
                  >
                    View all notifications
                  </router-link>
                </div>
              </div>
            </div>

            <!-- Profile -->
            <div class="relative">
              <button
                id="profile-toggle"
                class="flex items-center space-x-2 focus:outline-none"
                @click="isProfileDropdownOpen = !isProfileDropdownOpen"
              >
                <img
                  :src="user.avatar"
                  alt="User Avatar"
                  class="w-8 h-8 rounded-full border-2 border-light-blue"
                />
                <span class="hidden md:block text-sm font-medium">{{
                  user.name.split(' ')[0]
                }}</span>
                <font-awesome-icon :icon="['fas', 'chevron-down']" class="w-3 h-3 text-gray" />
              </button>

              <!-- Profile Dropdown -->
              <transition
                enter-active-class="transition ease-out duration-200"
                enter-from-class="transform opacity-0 scale-95"
                enter-to-class="transform opacity-100 scale-100"
                leave-active-class="transition ease-in duration-150"
                leave-from-class="transform opacity-100 scale-100"
                leave-to-class="transform opacity-0 scale-95"
              >
                <div
                  id="profile-dropdown"
                  v-if="isProfileDropdownOpen"
                  class="absolute right-0 mt-2 w-60 bg-white rounded-lg shadow-lg py-2 z-50 border border-gray-100"
                >
                  <!-- Profile Header -->
                  <div class="px-4 py-3 border-b border-gray-100">
                    <div class="flex items-center space-x-3">
                      <img
                        :src="user.avatar"
                        alt="User Avatar"
                        class="w-10 h-10 rounded-full border-2 border-light-blue"
                      />
                      <div>
                        <h4 class="text-sm font-semibold text-gray-800">{{ user.name }}</h4>
                        <p class="text-xs text-gray-500">{{ user.email }}</p>
                      </div>
                    </div>
                  </div>

                  <!-- Dropdown Menu Items -->
                  <div class="py-1">
                    <router-link
                      to="/profile"
                      class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors"
                      @click="isProfileDropdownOpen = false"
                    >
                      <div class="flex items-center">
                        <font-awesome-icon
                          :icon="['fas', 'user']"
                          class="w-4 h-4 mr-3 text-gray-400"
                        />
                        <span>My Profile</span>
                      </div>
                    </router-link>
                    <a
                      href="#"
                      class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors"
                    >
                      <div class="flex items-center">
                        <font-awesome-icon
                          :icon="['fas', 'gear']"
                          class="w-4 h-4 mr-3 text-gray-400"
                        />
                        <span>Account Settings</span>
                      </div>
                    </a>
                  </div>

                  <!-- Divider -->
                  <div class="border-t border-gray-100 my-1"></div>

                  <!-- Logout Option -->
                  <a
                    href="#"
                    class="block px-4 py-2 text-sm text-red hover:bg-red hover:bg-opacity-10 transition-colors"
                  >
                    <div class="flex items-center">
                      <font-awesome-icon :icon="['fas', 'sign-out-alt']" class="w-4 h-4 mr-3" />
                      <span>Log Out</span>
                    </div>
                  </a>
                </div>
              </transition>
            </div>
          </div>
        </div>
      </header>

      <!-- Dynamic Content Component -->
      <component :is="activeComponent" />
    </main>
  </div>
</template>

<style scoped>
/* Additional animations and transitions */
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.fade-in {
  animation: fadeIn 0.5s ease-in-out;
}

/* Custom scrollbar for sidebar */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.1);
}

::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 2px;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.3);
}
</style>
