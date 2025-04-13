import { defineStore } from 'pinia'
import router from '@/router/index'
import type { User, UserState } from '@/types/user'

// Mock users data
const mockUsersData: User[] = [
  {
    // Account Information
    ltoClientId: 'LTO-2023-78945',
    lastName: 'Morales',
    firstName: 'Stanleigh',
    middleName: 'Garcia',
    role: 'user',
    password: 'password123',
    status: 'active',

    // Contact Information
    email: 'stanleighmorales@gmail.com',
    telephoneNumber: '(02) 8123-4567',
    intAreaCode: '+63',
    mobileNumber: '912 345 6789',

    // Personal Information - General
    nationality: 'Filipino',
    civilStatus: 'Single',
    dateOfBirth: '1990-05-15',
    placeOfBirth: 'Manila, Philippines',
    educationalAttainment: 'College Graduate',
    tin: '123-456-789-000',

    // Personal Information - Medical
    gender: 'Male',
    bloodType: 'O+',
    complexion: 'Fair',
    bodyType: 'Medium',
    eyeColor: 'Brown',
    hairColor: 'Black',
    weight: 70,
    height: 175,
    organDonor: false,

    // People - Emergency Contact
    emergencyContactName: 'Maria Morales',
    emergencyContactNumber: '+63 917 123 4567',
    emergencyContactAddress: '123 Main Street, Quezon City',

    // People - Employer
    employerName: 'ABC Corporation',
    employerAddress: '789 Corporate Ave, Makati City',

    // People - Mother's Maiden Name
    motherLastName: 'Santos',
    motherFirstName: 'Elena',
    motherMiddleName: 'Cruz',

    // People - Father
    fatherLastName: 'Morales',
    fatherFirstName: 'Roberto',
    fatherMiddleName: 'Reyes',

    // Address
    houseNo: '123',
    street: 'Main Street',
    province: 'Metro Manila',
    city: 'Quezon City',
    barangay: 'Barangay 123',
    zipCode: '1100',

    // Other
    avatar: '/Land_Transportation_Office.webp',
  },
  {
    ltoClientId: 'LTO-ADMIN-001',
    lastName: 'Admin',
    firstName: 'System',
    middleName: '',
    role: 'admin',
    email: 'admin@smartplate.com',
    password: 'admin123',
    status: 'active',
    avatar: '/Land_Transportation_Office.webp',
  },
  {
    ltoClientId: 'LTO-OFFICER-001',
    lastName: 'OFFICER',
    firstName: 'TESTER',
    middleName: '',
    role: 'LTO Officer',
    email: 'adminofficer@smartplate.com',
    password: 'officer123',
    status: 'active',
    avatar: '/Land_Transportation_Office.webp',
  },
  {
    ltoClientId: 'LTO-2023-12345',
    lastName: 'Santos',
    firstName: 'Maria',
    middleName: 'Cruz',
    role: 'user',
    email: 'maria.santos@gmail.com',
    password: 'maria123',
    status: 'active',
    avatar: '/Land_Transportation_Office.webp',
  },
  {
    ltoClientId: 'LTO-2023-67890',
    lastName: 'Reyes',
    firstName: 'Juan',
    middleName: 'Dela',
    role: 'user',
    email: 'juan.reyes@gmail.com',
    password: 'juan123',
    status: 'active',
    avatar: '/Land_Transportation_Office.webp',
  },
]

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    currentUser: null,
    users: mockUsersData,
    isAuthenticated: false,
    token: localStorage.getItem('token') || null,
    loading: false,
    error: null,
    isRegistering: false,
    registrationData: null,
  }),

  getters: {
    fullName: (state): string => {
      if (!state.currentUser) return ''
      return `${state.currentUser.firstName} ${state.currentUser.middleName ? state.currentUser.middleName + ' ' : ''}${state.currentUser.lastName}`
    },

    userRole: (state): string => state.currentUser?.role || 'guest',

    isAdmin: (state): boolean => state.currentUser?.role === 'admin',
  },

  actions: {
    async login(email: string, password: string, isAdminLogin: boolean = false): Promise<User> {
      this.loading = true
      this.error = null

      try {
        // Find user with matching email and password
        const foundUser = this.users.find((user) => {
          // For admin login, only match admin users
          if (isAdminLogin) {
            return (
              (user.role === 'admin' || user.role === 'LTO Officer') &&
              user.email === email &&
              user.password === password
            )
          }
          // For regular login, only match non-admin users
          return user.role === 'user' && user.email === email && user.password === password
        })

        if (foundUser) {
          // Set user data (without the password)
          const userWithoutPassword: User = { ...foundUser }
          delete userWithoutPassword.password
          this.currentUser = userWithoutPassword
          this.isAuthenticated = true

          // Generate token and store it
          const tokenPrefix = foundUser.role === 'user' ? 'user_' : 'admin_'
          const fakeToken = tokenPrefix + Math.random().toString(36).substring(2)
          this.token = fakeToken
          localStorage.setItem('token', fakeToken)

          this.loading = false
          return this.currentUser
        } else {
          this.error = 'Invalid email or password'
          this.loading = false
          throw new Error('Invalid email or password')
        }
      } catch (err) {
        this.loading = false
        this.error = err instanceof Error ? err.message : 'An error occurred during login'
        throw err
      }
    },

    startRegistration(initialData?: Partial<User>): void {
      this.isRegistering = true
      this.registrationData = initialData || {}
    },

    cancelRegistration(): void {
      this.isRegistering = false
      this.registrationData = null
      router.push('/login')
    },

    async register(userData: Partial<User>): Promise<User> {
      this.loading = true
      this.error = null

      try {
        const newUser: User = {
          ...userData,
          role: 'user',
          status: 'active',
          ltoClientId: `LTO-${new Date().getFullYear()}-${Math.floor(10000 + Math.random() * 90000)}`,
        } as User

        // Add the new user to users array
        const password = this.registrationData?.password || 'defaultpassword'
        this.users.push({
          ...newUser,
          password,
        })

        // Set user data (without the password)
        const userWithoutPassword: User = { ...newUser }
        delete userWithoutPassword.password
        this.currentUser = userWithoutPassword
        this.isAuthenticated = true

        // Generate a fake token and store it
        const fakeToken = 'user_' + Math.random().toString(36).substring(2)
        this.token = fakeToken
        localStorage.setItem('token', fakeToken)

        // Clear registration state
        this.isRegistering = false
        this.registrationData = null

        this.loading = false
        return this.currentUser
      } catch (err) {
        this.loading = false
        this.error = err instanceof Error ? err.message : 'An error occurred during registration'
        throw err
      }
    },

    logout(): void {
      this.currentUser = null
      this.isAuthenticated = false
      this.token = null
      localStorage.removeItem('token')
      router.push('/login')
    },

    checkAuth(): boolean {
      const storedToken = localStorage.getItem('token')

      if (storedToken) {
        if (storedToken.startsWith('admin_')) {
          // Find an admin user
          const adminUser = this.users.find((user) => user.role === 'admin')
          if (adminUser) {
            const userWithoutPassword: User = { ...adminUser }
            delete userWithoutPassword.password
            this.currentUser = userWithoutPassword
          }
        } else {
          // Find a regular user (first one)
          const regularUser = this.users.find((user) => user.role === 'user')
          if (regularUser) {
            const userWithoutPassword: User = { ...regularUser }
            delete userWithoutPassword.password
            this.currentUser = userWithoutPassword
          }
        }

        this.isAuthenticated = true
        this.token = storedToken
        return true
      }

      return false
    },

    async updateUserProfile(updatedData: Partial<User>): Promise<User | null> {
      if (!this.currentUser) return null

      this.currentUser = {
        ...this.currentUser,
        ...updatedData,
      }
      return Promise.resolve(this.currentUser)
    },
  },
})
