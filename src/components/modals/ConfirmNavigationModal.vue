<script setup>
import { defineProps, defineEmits } from 'vue'

defineProps({
  isOpen: {
    type: Boolean,
    required: true,
  },
})

const emit = defineEmits(['confirm', 'cancel'])

const handleConfirm = () => {
  emit('confirm')
}

const handleCancel = () => {
  emit('cancel')
}
</script>

<template>
  <div v-if="isOpen" class="fixed inset-0 z-[9999] overflow-y-auto">
    <div
      class="flex min-h-screen items-center justify-center px-4 py-12 text-center sm:block sm:p-0"
    >
      <!-- Background overlay with increased opacity -->
      <div class="fixed inset-0 transition-opacity" aria-hidden="true">
        <div class="absolute inset-0 bg-gray-900 opacity-90"></div>
      </div>

      <!-- Modal panel with refined spacing and positioning -->
      <div
        class="inline-block w-full transform overflow-hidden rounded-xl border-2 border-gray-200 bg-white text-left align-bottom shadow-2xl transition-all sm:my-8 sm:w-full sm:max-w-2xl sm:align-middle"
      >
        <div class="bg-white px-8 py-7">
          <div class="sm:flex sm:items-start">
            <div
              class="mx-auto flex h-14 w-14 flex-shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-12 sm:w-12"
            >
              <font-awesome-icon
                :icon="['fas', 'exclamation-triangle']"
                class="h-8 w-8 text-red-600"
              />
            </div>
            <div class="mt-3 text-center sm:ml-6 sm:mt-0 sm:text-left">
              <h3 class="mb-3 text-2xl font-semibold leading-7 text-gray-900">
                Leave Registration Form?
              </h3>
              <div>
                <p class="text-base leading-relaxed text-gray-600">
                  Are you sure you want to leave? Your form progress will be lost and cannot be
                  recovered.
                </p>
              </div>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-8 py-5 sm:flex sm:flex-row-reverse sm:gap-4">
          <button
            type="button"
            class="w-full transform rounded-lg bg-red px-6 py-3 text-base font-medium text-black shadow-sm transition-all duration-200 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 sm:ml-4 sm:w-auto"
            @click="handleConfirm"
          >
            Leave
          </button>
          <button
            type="button"
            class="mt-3 w-full transform rounded-lg border border-gray-300 bg-white px-6 py-3 text-base font-medium text-gray-700 shadow-sm transition-all duration-200 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-dark-blue focus:ring-offset-2 sm:mt-0 sm:w-auto"
            @click="handleCancel"
          >
            Stay
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
