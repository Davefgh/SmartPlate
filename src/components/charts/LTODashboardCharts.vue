<script setup lang="ts">
import { useVehicleRegistrationStore } from '@/stores/vehicleRegistration'
import { useVehicleRegistrationFormStore } from '@/stores/vehicleRegistrationForm'
import { computed, onMounted } from 'vue'
import { Chart, registerables } from 'chart.js'

Chart.register(...registerables)

const vehicleRegistrationStore = useVehicleRegistrationStore()
const registrationFormStore = useVehicleRegistrationFormStore()

const registrationStatusData = computed(() => {
  const statuses = ['pending', 'approved', 'rejected']
  return statuses.map(
    (status) => registrationFormStore.forms.filter((form) => form.status === status).length,
  )
})

const plateStatusData = computed(() => {
  const statuses = ['pending', 'ready', 'issued']
  return statuses.map(
    (status) =>
      vehicleRegistrationStore.registrations.filter((reg) => reg.status === status).length,
  )
})

const getDailyRegistrations = computed(() => {
  const last7Days = Array.from({ length: 7 }, (_, i) => {
    const date = new Date()
    date.setDate(date.getDate() - i)
    return date.toISOString().split('T')[0]
  }).reverse()

  return last7Days.map((date) => ({
    date,
    count: registrationFormStore.forms.filter((form) => form.submissionDate.split('T')[0] === date)
      .length,
  }))
})

onMounted(() => {
  // Registration Status Pie Chart
  new Chart(document.getElementById('registrationStatusChart') as HTMLCanvasElement, {
    type: 'pie',
    data: {
      labels: ['Pending', 'Approved', 'Rejected'],
      datasets: [
        {
          data: registrationStatusData.value,
          backgroundColor: ['#3B82F6', '#10B981', '#EF4444'],
        },
      ],
    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          position: 'bottom',
        },
      },
    },
  })

  // Plate Status Bar Chart
  new Chart(document.getElementById('plateStatusChart') as HTMLCanvasElement, {
    type: 'bar',
    data: {
      labels: ['Pending', 'Ready', 'Issued'],
      datasets: [
        {
          label: 'Number of Plates',
          data: plateStatusData.value,
          backgroundColor: '#3B82F6',
        },
      ],
    },
    options: {
      responsive: true,
      scales: {
        y: {
          beginAtZero: true,
          ticks: {
            stepSize: 1,
          },
        },
      },
    },
  })

  // Daily Registration Line Chart
  new Chart(document.getElementById('dailyRegistrationChart') as HTMLCanvasElement, {
    type: 'line',
    data: {
      labels: getDailyRegistrations.value.map((item) => item.date),
      datasets: [
        {
          label: 'Daily Registrations',
          data: getDailyRegistrations.value.map((item) => item.count),
          borderColor: '#3B82F6',
          tension: 0.1,
        },
      ],
    },
    options: {
      responsive: true,
      scales: {
        y: {
          beginAtZero: true,
          ticks: {
            stepSize: 1,
          },
        },
      },
    },
  })
})
</script>

<template>
  <div class="mt-8">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-xl font-semibold mb-4">Registration Status</h2>
        <canvas id="registrationStatusChart"></canvas>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-xl font-semibold mb-4">Plate Status</h2>
        <canvas id="plateStatusChart"></canvas>
      </div>
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-xl font-semibold mb-4">Daily Registrations</h2>
        <canvas id="dailyRegistrationChart"></canvas>
      </div>
    </div>
  </div>
</template>
