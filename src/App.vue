<script setup>
import { RouterView } from 'vue-router'
import { onMounted } from 'vue'
import { useUserStore } from './stores/user'
import { useVehicleRegistrationFormStore } from './stores/vehicleRegistrationForm'

const userStore = useUserStore()
const registrationFormStore = useVehicleRegistrationFormStore()

// On app startup, ensure we load all forms and restore user session
onMounted(() => {
  console.log('==== App mounted - initializing data ====')

  const isAuthenticated = userStore.checkAuth()
  console.log('User authenticated:', isAuthenticated)

  if (isAuthenticated && userStore.currentUser?.ltoClientId) {
    const userId = userStore.currentUser.ltoClientId
    console.log('Current authenticated user:', userId)

    // Log store state for debugging
    registrationFormStore.debugForms()
  } else {
    console.log('No authenticated user')
  }

  console.log('==== App initialization complete ====')
})
</script>

<template>
  <RouterView />
</template>
