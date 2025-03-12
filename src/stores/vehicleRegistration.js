import { defineStore } from 'pinia'
import { useUserStore } from './user'

// Mock data for vehicles, plates, and registrations
const mockVehiclesData = [
  {
    id: 1,
    make: 'Toyota',
    model: 'Corolla',
    year: 2022,
    plateNumber: 'ABC-123',
    plateType: 'Regular',
    mvFileNumber: 'MV-2024-10001',
    color: 'White',
    status: 'Active',
    lastUpdated: '2025-02-15',
    ownerId: 'LTO-2023-78945',
  },
  {
    id: 2,
    make: 'Honda',
    model: 'Civic',
    year: 2021,
    plateNumber: 'XYZ-789',
    plateType: 'Temporary',
    mvFileNumber: 'MV-2024-10002',
    color: 'Black',
    status: 'Active',
    lastUpdated: '2025-01-20',
    ownerId: 'LTO-2023-78945',
  },
  {
    id: 3,
    make: 'Ford',
    model: 'Mustang',
    year: 2023,
    plateNumber: 'DEF-456',
    plateType: 'Improvised',
    mvFileNumber: 'MV-2025-10003',
    color: 'Red',
    status: 'Pending',
    lastUpdated: '2025-03-01',
    ownerId: 'LTO-2023-78945',
  },
  {
    id: 4,
    make: 'Toyota',
    model: 'Vios',
    year: 2022,
    plateNumber: 'GHI-789',
    plateType: 'Regular',
    mvFileNumber: 'MV-2024-10004',
    color: 'Silver',
    status: 'Active',
    lastUpdated: '2025-02-10',
    ownerId: 'LTO-2023-12345',
  },
  {
    id: 5,
    make: 'Mitsubishi',
    model: 'Montero',
    year: 2023,
    plateNumber: 'JKL-012',
    plateType: 'Regular',
    mvFileNumber: 'MV-2024-10005',
    color: 'Black',
    status: 'Active',
    lastUpdated: '2025-01-15',
    ownerId: 'LTO-2023-67890',
  },
]

const mockPlatesData = [
  {
    id: 1,
    plateNumber: 'ABC-123',
    vehicleId: 1,
    registrationDate: '2024-10-15',
    expiryDate: '2025-10-15',
    status: 'Active',
    type: 'Private',
    plateType: 'Regular',
    mvFileNumber: 'MV-2024-10001',
  },
  {
    id: 2,
    plateNumber: 'XYZ-789',
    vehicleId: 2,
    registrationDate: '2024-08-20',
    expiryDate: '2025-08-20',
    status: 'Active',
    type: 'Private',
    plateType: 'Temporary',
    mvFileNumber: 'MV-2024-10002',
  },
  {
    id: 3,
    plateNumber: 'DEF-456',
    vehicleId: 3,
    registrationDate: '2025-01-05',
    expiryDate: '2026-01-05',
    status: 'Pending',
    type: 'Private',
    plateType: 'Improvised',
    mvFileNumber: 'MV-2025-10003',
  },
  {
    id: 4,
    plateNumber: 'GHI-789',
    vehicleId: 4,
    registrationDate: '2024-10-10',
    expiryDate: '2025-10-10',
    status: 'Active',
    type: 'Private',
    plateType: 'Regular',
    mvFileNumber: 'MV-2024-10004',
  },
  {
    id: 5,
    plateNumber: 'JKL-012',
    vehicleId: 5,
    registrationDate: '2024-09-15',
    expiryDate: '2025-09-15',
    status: 'Active',
    type: 'Private',
    plateType: 'Regular',
    mvFileNumber: 'MV-2024-10005',
  },
]

const mockRegistrationsData = [
  {
    id: 1,
    vehicleId: 1,
    plateId: 1,
    registrationType: 'New Registration',
    submissionDate: '2024-10-05',
    expiryDate: '2025-10-05',
    status: 'Approved',
    documents: ['Proof of Ownership', 'Insurance Certificate', 'Emission Test'],
    fees: {
      registrationFee: 1500,
      plateIssuanceFee: 450,
      computerFee: 170,
      total: 2120,
    },
  },
  {
    id: 2,
    vehicleId: 2,
    plateId: 2,
    registrationType: 'Renewal',
    submissionDate: '2024-08-10',
    expiryDate: '2025-08-10',
    status: 'Approved',
    documents: ['Insurance Certificate', 'Emission Test'],
    fees: {
      registrationFee: 1500,
      computerFee: 170,
      total: 1670,
    },
  },
  {
    id: 3,
    vehicleId: 3,
    plateId: 3,
    registrationType: 'New Registration',
    submissionDate: '2025-01-05',
    expiryDate: '2026-01-05',
    status: 'Pending',
    documents: ['Proof of Ownership', 'Insurance Certificate', 'Emission Test'],
    fees: {
      registrationFee: 1500,
      plateIssuanceFee: 450,
      computerFee: 170,
      total: 2120,
    },
  },
  {
    id: 4,
    vehicleId: 4,
    plateId: 4,
    registrationType: 'New Registration',
    submissionDate: '2024-10-10',
    expiryDate: '2025-10-10',
    status: 'Approved',
    documents: ['Proof of Ownership', 'Insurance Certificate', 'Emission Test'],
    fees: {
      registrationFee: 1500,
      plateIssuanceFee: 450,
      computerFee: 170,
      total: 2120,
    },
  },
  {
    id: 5,
    vehicleId: 5,
    plateId: 5,
    registrationType: 'New Registration',
    submissionDate: '2024-09-15',
    expiryDate: '2025-09-15',
    status: 'Approved',
    documents: ['Proof of Ownership', 'Insurance Certificate', 'Emission Test'],
    fees: {
      registrationFee: 1500,
      plateIssuanceFee: 450,
      computerFee: 170,
      total: 2120,
    },
  },
]

export const useVehicleRegistrationStore = defineStore('vehicleRegistration', {
  // State
  state: () => ({
    vehicles: [...mockVehiclesData],
    plates: [...mockPlatesData],
    registrations: [...mockRegistrationsData],
    loading: false,
    error: null,
  }),

  // Getters
  getters: {
    // Get vehicles for the current user
    userVehicles: (state) => {
      const userStore = useUserStore()
      if (!userStore.user) return []

      return state.vehicles.filter((vehicle) => vehicle.ownerId === userStore.user.ltoClientId)
    },

    // Get plates for the current user's vehicles
    userPlates: (state) => {
      const userVehicleIds = state.userVehicles.map((vehicle) => vehicle.id)
      return state.plates.filter((plate) => userVehicleIds.includes(plate.vehicleId))
    },

    // Get registrations for the current user's vehicles
    userRegistrations: (state) => {
      const userVehicleIds = state.userVehicles.map((vehicle) => vehicle.id)
      return state.registrations.filter((registration) =>
        userVehicleIds.includes(registration.vehicleId),
      )
    },

    getVehicleById: (state) => (id) => {
      return state.vehicles.find((vehicle) => vehicle.id === id)
    },

    getPlateById: (state) => (id) => {
      return state.plates.find((plate) => plate.id === id)
    },

    getRegistrationById: (state) => (id) => {
      return state.registrations.find((registration) => registration.id === id)
    },

    getPlateByVehicleId: (state) => (vehicleId) => {
      return state.plates.find((plate) => plate.vehicleId === vehicleId)
    },

    getRegistrationByVehicleId: (state) => (vehicleId) => {
      return state.registrations.find((registration) => registration.vehicleId === vehicleId)
    },

    vehiclesWithOwnerInfo: (state) => {
      const userStore = useUserStore()
      return state.vehicles.map((vehicle) => {
        const owner = userStore.mockUsers.find((user) => user.ltoClientId === vehicle.ownerId)

        // Get the plate for this vehicle
        const plate = state.plates.find((p) => p.vehicleId === vehicle.id)

        return {
          ...vehicle,
          owner: owner ? `${owner.firstName} ${owner.lastName}` : 'Unknown',
          plateDetails: plate || null,
        }
      })
    },

    // Get plates with vehicle information
    platesWithVehicleInfo: (state) => {
      return state.plates.map((plate) => {
        const vehicle = state.vehicles.find((v) => v.id === plate.vehicleId)
        return {
          ...plate,
          vehicleMake: vehicle?.make || '',
          vehicleModel: vehicle?.model || '',
          vehicleYear: vehicle?.year || '',
          vehicleColor: vehicle?.color || '',
        }
      })
    },

    // Get registrations with vehicle and plate information
    registrationsWithDetails: (state) => {
      return state.registrations.map((registration) => {
        const vehicle = state.vehicles.find((v) => v.id === registration.vehicleId)
        const plate = state.plates.find((p) => p.id === registration.plateId)

        return {
          ...registration,
          vehicleInfo: vehicle ? `${vehicle.make} ${vehicle.model} ${vehicle.year}` : '',
          plateNumber: plate?.plateNumber || '',
        }
      })
    },

    // Get active registrations
    activeRegistrations: (state) => {
      return state.registrations.filter((reg) => reg.status === 'Approved')
    },

    // Get pending registrations
    pendingRegistrations: (state) => {
      return state.registrations.filter((reg) => reg.status === 'Pending')
    },

    // Get expired registrations (where expiryDate is before today)
    expiredRegistrations: (state) => {
      const today = new Date()
      return state.registrations.filter((reg) => {
        const expiryDate = new Date(reg.expiryDate)
        return expiryDate < today
      })
    },

    // Get soon to expire registrations (within 30 days)
    soonToExpireRegistrations: (state) => {
      const today = new Date()
      const thirtyDaysFromNow = new Date()
      thirtyDaysFromNow.setDate(today.getDate() + 30)

      return state.registrations.filter((reg) => {
        const expiryDate = new Date(reg.expiryDate)
        return expiryDate >= today && expiryDate <= thirtyDaysFromNow
      })
    },
  },

  // Actions
  actions: {
    // Calculate days remaining until expiry for a registration or plate
    getDaysRemaining(expiryDateStr) {
      const today = new Date()
      const expiryDate = new Date(expiryDateStr)
      const diffTime = expiryDate - today
      const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
      return diffDays
    },

    // Get expiry status (Expired, Expiring Soon, Valid)
    getExpiryStatus(expiryDateStr) {
      const daysRemaining = this.getDaysRemaining(expiryDateStr)

      if (daysRemaining < 0) return 'Expired'
      if (daysRemaining < 30) return 'Expiring Soon'
      return 'Valid'
    },
  },
})
