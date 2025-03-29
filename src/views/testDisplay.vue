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
<template>
  <div class="container">
    <h1>User List</h1>
    <!-- Changed from Vehicle List to User List -->
    <div v-if="loading" class="status-message loading">Loading...</div>
    <div v-else-if="error" class="status-message error">
      Error: {{ error }}
      <button @click="retryFetch" class="retry-btn">Retry</button>
    </div>
    <div v-else>
      <div v-if="users.length === 0" class="status-message info">No users found</div>
      <div v-else class="table-container">
        <table class="user-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Email</th>
              <th>Role</th>
              <th>Status</th>
              <th>Contact Info</th>
              <th>Emergency Contact</th>
              <th>Created</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.user_id">
              <td>{{ user.user_id }}</td>
              <td>
                {{ user.first_name }}
                {{ user.middle_name }}
                {{ user.last_name }}
              </td>
              <td>{{ user.email }}</td>
              <td>{{ user.role }}</td>
              <td class="status-cell">
                <span :class="user.status.toLowerCase()">{{ user.status }}</span>
              </td>
              <td>
                <div v-if="user.contact">
                  <div>📞 {{ user.contact.mobile_number || 'N/A' }}</div>
                  <div>☎️ {{ user.contact.telephone_number || 'N/A' }}</div>
                </div>
                <div v-else>N/A</div>
              </td>
              <td>
                <div v-if="user.contact?.emergency_contact_name">
                  {{ user.contact.emergency_contact_name }}<br />
                  {{ user.contact.emergency_contact_number }}<br />
                  ({{ user.contact.emergency_contact_relationship }})
                </div>
                <div v-else>N/A</div>
              </td>
              <td>{{ formatDate(user.created) }}</td>
            </tr>
          </tbody>
        </table>
        <div class="table-footer">Showing {{ users.length }} users</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  padding: 1rem;
  max-width: 1200px;
  margin: 0 auto;
}

.status-message {
  padding: 1rem;
  border-radius: 4px;
  margin: 1rem 0;
  text-align: center;
}

.loading {
  background: #e9f5ff;
  color: #004085;
  border: 1px solid #b8daff;
}

.error {
  background: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}

.info {
  background: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.retry-btn {
  margin-left: 1rem;
  padding: 0.25rem 1rem;
  background: #dc3545;
  color: white;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  transition: opacity 0.2s;
}

.retry-btn:hover {
  opacity: 0.8;
}

.table-container {
  border: 1px solid #eee;
  border-radius: 8px;
  overflow-x: auto;
}

.user-table {
  width: 100%;
  border-collapse: collapse;
  min-width: 800px;
}

.user-table th,
.user-table td {
  padding: 12px 15px;
  text-align: left;
  border-bottom: 1px solid #eee;
}

.user-table th {
  background-color: #f8f9fa;
  font-weight: 600;
}

.user-table tr:hover {
  background-color: #f8f9fa;
}

.status-cell span {
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 0.85em;
}

.active {
  background: #d4edda;
  color: #155724;
}

.inactive {
  background: #fff3cd;
  color: #856404;
}

.table-footer {
  padding: 12px;
  background: #f8f9fa;
  color: #666;
  font-size: 0.9em;
  text-align: center;
  border-top: 1px solid #eee;
}
</style>
