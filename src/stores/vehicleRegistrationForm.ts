import { defineStore } from 'pinia'
import type {
  VehicleRegistrationForm,
  VehicleRegistrationState,
  VehicleDocuments,
} from '@/types/vehicleRegistration'

// Mock data for vehicle registration forms
const mockRegistrationForms: VehicleRegistrationForm[] = [
  {
    id: 'REG-2024-001',
    userId: 'LTO-2023-78945',
    isNewVehicle: true,
    vehicleType: 'Sedan',
    make: 'Toyota',
    model: 'Corolla Altis',
    year: '2024',
    engineNumber: 'ENG-2024-001',
    chassisNumber: 'CHS-2024-001',
    color: 'Pearl White',
    status: 'approved',
    submissionDate: '2024-01-15',
    registrationType: 'New Vehicle',
    documents: {
      csr: { name: 'csr.pdf', size: 1024 },
      salesInvoice: { name: 'invoice.pdf', size: 2048 },
      insurance: { name: 'insurance.pdf', size: 1536 },
      orCr: null,
      deedOfSale: null,
      pnpHpgClearance: null,
    },
    appointmentDate: '2024-02-15',
    appointmentTime: '10:00',
    referenceCode: 'SP-24-001',
    inspectionStatus: 'approved',
    inspectionCode: 'INS-2024-001',
    paymentCode: 'PAY-2024-001',
    paymentStatus: 'approved',
    verificationStatus: 'approved',
    privacyConsent: true,
    declarationConsent: true,
  },
  {
    id: 'REG-2024-002',
    userId: 'LTO-2023-12345',
    isNewVehicle: false,
    vehicleType: 'SUV',
    make: 'Honda',
    model: 'CR-V',
    year: '2020',
    engineNumber: 'ENG-2020-002',
    chassisNumber: 'CHS-2020-002',
    color: 'Cosmic Blue',
    status: 'pending',
    submissionDate: '2024-02-01',
    registrationType: 'Renewal',
    documents: {
      csr: null,
      salesInvoice: null,
      insurance: null,
      orCr: { name: 'or_cr.pdf', size: 2048 },
      deedOfSale: { name: 'deed.pdf', size: 1536 },
      pnpHpgClearance: { name: 'clearance.pdf', size: 1024 },
    },
    appointmentDate: '2024-02-20',
    appointmentTime: '14:30',
    referenceCode: 'SP-24-002',
    inspectionStatus: 'pending',
    inspectionCode: 'INS-2024-002',
    paymentStatus: 'pending',
    paymentCode: 'PAY-2024-002',
    verificationStatus: 'pending',
    privacyConsent: true,
    declarationConsent: true,
  },
  {
    id: 'REG-2024-003',
    userId: 'LTO-2023-67890',
    isNewVehicle: true,
    vehicleType: 'Motorcycle',
    make: 'Yamaha',
    model: 'MT-07',
    year: '2024',
    engineNumber: 'ENG-2024-003',
    chassisNumber: 'CHS-2024-003',
    color: 'Racing Blue',
    status: 'rejected',
    submissionDate: '2024-02-10',
    registrationType: 'New Vehicle',
    documents: {
      csr: { name: 'csr.pdf', size: 1024 },
      salesInvoice: { name: 'invoice.pdf', size: 2048 },
      insurance: null,
      orCr: null,
      deedOfSale: null,
      pnpHpgClearance: null,
    },
    appointmentDate: null,
    appointmentTime: null,
    referenceCode: '',
    inspectionStatus: 'rejected',
    inspectionCode: 'INS-2024-003',
    paymentCode: 'PAY-2024-003',
    paymentStatus: 'pending',
    verificationStatus: 'pending',
    privacyConsent: true,
    declarationConsent: false,
  },
]

export const useVehicleRegistrationFormStore = defineStore('vehicleRegistrationForm', {
  persist: true,
  state: (): VehicleRegistrationState => ({
    hasUnsavedChanges: false,
    forms: [...mockRegistrationForms],
    formData: {
      id: '',
      userId: '',
      isNewVehicle: true,
      vehicleType: '',
      make: '',
      model: '',
      year: '',
      engineNumber: '',
      chassisNumber: '',
      color: '',
      status: 'pending',
      submissionDate: new Date().toISOString().split('T')[0],
      registrationType: 'New Vehicle',
      documents: {
        csr: null,
        salesInvoice: null,
        insurance: null,
        orCr: null,
        deedOfSale: null,
        pnpHpgClearance: null,
      },
      appointmentDate: null,
      appointmentTime: null,
      referenceCode: '',
      inspectionStatus: 'pending',
      inspectionCode: '',
      paymentCode: '',
      paymentStatus: 'pending',
      verificationStatus: 'pending',
      privacyConsent: false,
      declarationConsent: false,
    },
    errors: {
      vehicleType: '',
      make: '',
      model: '',
      year: '',
      engineNumber: '',
      chassisNumber: '',
      color: '',
      documents: {
        csr: '',
        salesInvoice: '',
        insurance: '',
        orCr: '',
        deedOfSale: '',
        pnpHpgClearance: '',
      },
      referenceSlip: '',
      privacyConsent: '',
      declarationConsent: '',
    },
    isSubmitting: false,
    currentStep: 1,
  }),

  getters: {
    requiredDocuments: (state): Array<keyof VehicleDocuments> => {
      if (state.formData.isNewVehicle) {
        return ['csr', 'salesInvoice', 'insurance']
      } else {
        return ['orCr', 'deedOfSale', 'pnpHpgClearance']
      }
    },
    isFormComplete: (state): boolean => {
      return (
        state.formData.vehicleType !== '' &&
        state.formData.make !== '' &&
        state.formData.model !== '' &&
        state.formData.year !== '' &&
        state.formData.engineNumber !== '' &&
        state.formData.chassisNumber !== '' &&
        state.formData.color !== ''
      )
    },
    hasValidationErrors: (state): boolean => {
      return Object.values(state.errors).some((error) => error !== '')
    },
    pendingForms: (state): VehicleRegistrationForm[] => {
      return state.forms.filter((form) => form.inspectionStatus === 'pending')
    },
    approvedForms: (state): VehicleRegistrationForm[] => {
      return state.forms.filter((form) => form.inspectionStatus === 'approved')
    },
    rejectedForms: (state): VehicleRegistrationForm[] => {
      return state.forms.filter((form) => form.inspectionStatus === 'rejected')
    },
    getFormsByUserId:
      (state) =>
      (userId: string): VehicleRegistrationForm[] => {
        return state.forms.filter((form) => form.userId === userId)
      },
  },

  actions: {
    validateVehicleInfo(): boolean {
      let isValid = true

      // Reset errors
      this.errors.vehicleType = ''
      this.errors.make = ''
      this.errors.model = ''
      this.errors.year = ''
      this.errors.engineNumber = ''
      this.errors.chassisNumber = ''
      this.errors.color = ''

      // Vehicle type validation
      if (!this.formData.vehicleType) {
        this.errors.vehicleType = 'Please select a vehicle type'
        isValid = false
      }

      // Make validation
      if (!this.formData.make) {
        this.errors.make = 'Vehicle make is required'
        isValid = false
      } else if (this.formData.make.length < 2) {
        this.errors.make = 'Vehicle make must be at least 2 characters'
        isValid = false
      }

      // Model validation
      if (!this.formData.model) {
        this.errors.model = 'Vehicle model is required'
        isValid = false
      } else if (this.formData.model.length < 2) {
        this.errors.model = 'Vehicle model must be at least 2 characters'
        isValid = false
      }

      // Year validation
      const currentYear = new Date().getFullYear()
      const yearValue = Number(this.formData.year)
      if (!this.formData.year) {
        this.errors.year = 'Vehicle year is required'
        isValid = false
      } else if (isNaN(yearValue) || yearValue < 1900 || yearValue > currentYear + 1) {
        this.errors.year = `Year must be between 1900 and ${currentYear + 1}`
        isValid = false
      }

      // Engine number validation
      if (!this.formData.engineNumber) {
        this.errors.engineNumber = 'Engine number is required'
        isValid = false
      }

      // Chassis number validation
      if (!this.formData.chassisNumber) {
        this.errors.chassisNumber = 'Chassis number is required'
        isValid = false
      }

      // Color validation
      if (!this.formData.color) {
        this.errors.color = 'Vehicle color is required'
        isValid = false
      } else if (this.formData.color.length < 2) {
        this.errors.color = 'Vehicle color must be at least 2 characters'
        isValid = false
      }

      return isValid
    },

    setUnsavedChanges(value: boolean): void {
      this.hasUnsavedChanges = value
    },

    async submitRegistration(): Promise<boolean> {
      try {
        this.isSubmitting = true
        // TODO: Implement API call to submit registration
        await new Promise((resolve) => setTimeout(resolve, 1000)) // Simulated API call
        return true
      } catch (error) {
        console.error('Registration submission failed:', error)
        return false
      } finally {
        this.isSubmitting = false
      }
    },

    resetForm(): void {
      this.hasUnsavedChanges = false
      this.formData = {
        isNewVehicle: true,
        vehicleType: '',
        make: '',
        model: '',
        year: '',
        engineNumber: '',
        chassisNumber: '',
        color: '',
        documents: {
          csr: null,
          salesInvoice: null,
          insurance: null,
          orCr: null,
          deedOfSale: null,
          pnpHpgClearance: null,
        },
        appointmentDate: null,
        appointmentTime: null,
        referenceCode: '',
        inspectionStatus: 'pending' as const,
        inspectionCode: '',
        paymentStatus: 'pending' as const,
        paymentCode: '',
        verificationStatus: 'pending' as const,
        privacyConsent: false,
        declarationConsent: false,
      } as VehicleRegistrationForm
      this.errors = {
        vehicleType: '',
        make: '',
        model: '',
        year: '',
        engineNumber: '',
        chassisNumber: '',
        color: '',
        documents: {
          csr: '',
          salesInvoice: '',
          insurance: '',
          orCr: '',
          deedOfSale: '',
          pnpHpgClearance: '',
        },
        referenceSlip: '',
        privacyConsent: '',
        declarationConsent: '',
      }
      this.currentStep = 1
      this.isSubmitting = false
    },

    validateDocuments(): boolean {
      let isValid = true

      // Reset document errors
      Object.keys(this.formData.documents).forEach((doc) => {
        this.errors.documents[doc as keyof typeof this.errors.documents] = ''
      })

      // Check required documents
      this.requiredDocuments.forEach((docType) => {
        if (!this.formData.documents[docType as keyof typeof this.formData.documents]) {
          this.errors.documents[docType as keyof typeof this.errors.documents] =
            'This document is required'
          isValid = false
        }
      })

      return isValid
    },

    validateAppointment(): boolean {
      let isValid = true

      if (!this.formData.appointmentDate || !this.formData.appointmentTime) {
        alert('Please schedule an appointment first')
        isValid = false
      }

      return isValid
    },

    validatePayment(): boolean {
      let isValid = true

      if (!this.formData.privacyConsent) {
        this.errors.privacyConsent = 'Please consent to data privacy'
        isValid = false
      }

      if (!this.formData.declarationConsent) {
        this.errors.declarationConsent = 'Please declare the accuracy of your information'
        isValid = false
      }

      this.hasUnsavedChanges = true
      return isValid
    },
  },
})
