<template>
  <div>
    <h1>Vehicle List</h1>
    <div v-if="loading" class="loading">Loading... ({{ loadingDots }})</div>
    <div v-else-if="error" class="error">
      Error: {{ error }}
      <button @click="retryFetch">Retry</button>
    </div>
    <div v-else>
      <div v-if="users.length === 0" class="no-data">No vehicles found in database</div>
      <table v-else class="users-table">
        <thead>
          <tr>
            <th v-for="header in tableHeaders" :key="header">
              {{ header }}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.user_id">
            <td>{{ user.user_id }}</td>
            <td>{{ user.last_name }}</td>
            <td>{{ user.first_name }}</td>
            <td>{{ user.middle_name }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.role }}</td>
            <td>{{ user.status }}</td>
            <td>{{ user.lto_client_id }}</td>
            <td>{{ formatDate(user.created) }}</td>
            <td>{{ formatDate(user.updated) }}</td>
          </tr>
        </tbody>
      </table>
      <div class="debug-info">Loaded {{ users.length }} users</div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const users = ref([])
const loading = ref(true)
const error = ref(null)
const loadingDots = ref('')
const tableHeaders = [
  'USER_ID',
  'LAST_NAME',
  'FIRST_NAME',
  'MIDDLE_NAME',
  'EMAIL',
  'ROLE',
  'STATUS',
  'LTO_CLIENT_ID',
  'Created',
  'Updated',
]

// Animate loading dots
let interval = null

const formatDate = (dateString) => {
  try {
    const options = {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    }
    return new Date(dateString).toLocaleDateString(undefined, options)
  } catch (e) {
    console.error('Date formatting error:', e)
    return 'Invalid date'
  }
}

const startLoadingAnimation = () => {
  interval = setInterval(() => {
    loadingDots.value = '.'.repeat((Date.now() % 3) + 1)
  }, 500)
}

const fetchData = async () => {
  try {
    loading.value = true
    error.value = null
    startLoadingAnimation()

    console.log('Starting data fetch...')
    const response = await fetch('http://localhost:8081/users')
    console.log('Received response status:', response.status)

    if (!response.ok) {
      const errorText = await response.text()
      throw new Error(`HTTP ${response.status}: ${errorText || 'Unknown error'}`)
    }

    const data = await response.json()
    console.log('API response data:', data)

    // Handle different response formats
    if (Array.isArray(data)) {
      users.value = data
    } else if (data.items && Array.isArray(data.items)) {
      users.value = data.items
    } else {
      throw new Error('Unexpected response format. Expected array or items array')
    }

    console.log('Processed Users:', users.value)
  } catch (err) {
    error.value = `Failed to load: ${err.message}`
    console.error('Fetch error:', err)
  } finally {
    loading.value = false
    clearInterval(interval)
  }
}

const retryFetch = () => {
  fetchData()
}

onMounted(() => {
  startLoadingAnimation()
  fetchData()
})
</script>

<style scoped>
.loading {
  color: #666;
  padding: 1rem;
  border: 1px solid #ddd;
  margin: 1rem 0;
}

.error {
  color: #dc3545;
  background-color: #f8d7da;
  border: 1px solid #f5c6cb;
  padding: 1rem;
  margin: 1rem 0;
  border-radius: 4px;
}

.error button {
  margin-left: 1rem;
  padding: 0.25rem 0.5rem;
  background-color: #dc3545;
  color: white;
  border: none;
  border-radius: 3px;
  cursor: pointer;
}

.vehicle-table {
  width: 100%;
  border-collapse: collapse;
  margin: 1rem 0;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12);
}

.vehicle-table th,
.vehicle-table td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.vehicle-table th {
  background-color: #f8f9fa;
  font-weight: 600;
}

.vehicle-table tr:hover {
  background-color: #f8f9fa;
}

.debug-info {
  color: #666;
  font-size: 0.9rem;
  margin-top: 0.5rem;
}

.no-data {
  color: #666;
  padding: 1rem;
  border: 1px solid #ddd;
  margin: 1rem 0;
}
</style>
