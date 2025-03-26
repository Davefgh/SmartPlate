<template>
  <div class="crud-container">
    <h1>Vehicle CRUD Operations</h1>

    <!-- Create/Edit Form -->
    <div class="vehicle-form">
      <h2>{{ editing ? 'Edit Vehicle' : 'Create New Vehicle' }}</h2>
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label>Make:</label>
          <input v-model="formData.make" required>
        </div>
        <div class="form-group">
          <label>Model:</label>
          <input v-model="formData.model" required>
        </div>
        <div class="form-group">
          <label>Plate:</label>
          <input v-model="formData.plate" required>
        </div>
        <div class="form-actions">
          <button type="submit">{{ editing ? 'Update' : 'Create' }}</button>
          <button v-if="editing" @click="cancelEdit">Cancel</button>
        </div>
      </form>
    </div>

    <!-- Vehicle List -->
    <div class="vehicle-list">
      <div v-if="loading">Loading vehicles...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else>
        <div v-if="vehicles.length === 0" class="no-data">No vehicles found</div>
        <div v-else class="vehicle-grid">
          <div v-for="vehicle in vehicles" :key="vehicle.id" class="vehicle-card">
            <div class="vehicle-info">
              <div><strong>ID:</strong> {{ vehicle.id }}</div>
              <div><strong>Make:</strong> {{ vehicle.make }}</div>
              <div><strong>Model:</strong> {{ vehicle.model }}</div>
              <div><strong>Plate:</strong> {{ vehicle.plate }}</div>
              <div><strong>Created:</strong> {{ formatDate(vehicle.created) }}</div>
              <div><strong>Updated:</strong> {{ formatDate(vehicle.updated) }}</div>
            </div>
            <div class="vehicle-actions">
              <button @click="startEdit(vehicle)">Edit</button>
              <button @click="deleteVehicle(vehicle.id)">Delete</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'

const vehicles = ref([])
const loading = ref(true)
const error = ref(null)
const editing = ref(false)
const currentVehicleId = ref(null)

const formData = reactive({
  make: '',
  model: '',
  plate: ''
})

// Format dates
const formatDate = (dateString) => {
  return new Date(dateString).toLocaleString()
}

// Fetch all vehicles
const fetchVehicles = async () => {
  try {
    loading.value = true
    const response = await fetch('http://localhost:8081/vehicles')
    if (!response.ok) throw new Error('Failed to fetch vehicles')
    vehicles.value = await response.json()
    error.value = null
  } catch (err) {
    error.value = err.message
    console.error('Fetch error:', err)
  } finally {
    loading.value = false
  }
}

// Handle form submission
const handleSubmit = async () => {
  try {
    const url = editing.value
      ? `http://localhost:8081/vehicles/${currentVehicleId.value}`
      : 'http://localhost:8081/vehicles'

    const method = editing.value ? 'PUT' : 'POST'

    const response = await fetch(url, {
      method,
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData)
    })

    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.error || 'Request failed')
    }

    await fetchVehicles()
    resetForm()
  } catch (err) {
    error.value = err.message
    console.error('Submission error:', err)
  }
}

// Start editing a vehicle
const startEdit = (vehicle) => {
  editing.value = true
  currentVehicleId.value = vehicle.id
  formData.make = vehicle.make
  formData.model = vehicle.model
  formData.plate = vehicle.plate
}

// Cancel editing
const cancelEdit = () => {
  resetForm()
}

// Delete a vehicle
const deleteVehicle = async (id) => {
  try {
    const response = await fetch(`http://localhost:8081/vehicles/${id}`, {
      method: 'DELETE'
    })

    if (!response.ok) throw new Error('Failed to delete vehicle')

    await fetchVehicles()
    error.value = null
  } catch (err) {
    error.value = err.message
    console.error('Delete error:', err)
  }
}

// Reset form
const resetForm = () => {
  editing.value = false
  currentVehicleId.value = null
  formData.make = ''
  formData.model = ''
  formData.plate = ''
}

// Initial load
onMounted(fetchVehicles)
</script>

<style scoped>
.crud-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.vehicle-form {
  background: #f5f5f5;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
}

.form-group input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.form-actions {
  margin-top: 15px;
}

button {
  padding: 8px 15px;
  margin-right: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  background-color: #007bff;
  color: white;
}

button:hover {
  background-color: #0056b3;
}

.vehicle-grid {
  display: grid;
  gap: 15px;
}

.vehicle-card {
  background: white;
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 15px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.vehicle-info {
  flex-grow: 1;
}

.vehicle-actions {
  display: flex;
  gap: 10px;
}

.error {
  color: #dc3545;
  padding: 10px;
  background-color: #f8d7da;
  border: 1px solid #f5c6cb;
  border-radius: 4px;
  margin: 10px 0;
}

.no-data {
  color: #666;
  padding: 20px;
  text-align: center;
  border: 1px solid #ddd;
  border-radius: 4px;
}
</style>
