export interface VehicleDocument {
  name: string
  size: number
}

export interface VehicleDocuments {
  csr: VehicleDocument | null
  salesInvoice: VehicleDocument | null
  insurance: VehicleDocument | null
  orCr: VehicleDocument | null
  deedOfSale: VehicleDocument | null
  pnpHpgClearance: VehicleDocument | null
}

export interface DocumentErrors {
  csr: string
  salesInvoice: string
  insurance: string
  orCr: string
  deedOfSale: string
  pnpHpgClearance: string
}

export interface VehicleRegistrationErrors {
  vehicleType: string
  make: string
  model: string
  year: string
  engineNumber: string
  chassisNumber: string
  color: string
  documents: DocumentErrors
  referenceSlip: string
  privacyConsent: string
  declarationConsent: string
}

export interface VehicleRegistrationForm {
  id?: string
  userId?: string
  isNewVehicle: boolean
  vehicleType: string
  make: string
  model: string
  year: string
  engineNumber: string
  chassisNumber: string
  color: string
  documents: VehicleDocuments
  appointmentDate: string | null
  appointmentTime: string | null
  referenceCode: string
  inspectionStatus: 'pending' | 'approved' | 'rejected'
  referenceSlip: VehicleDocument | null
  paymentStatus: 'pending' | 'paid'
  verificationStatus: 'pending' | 'verified'
  privacyConsent: boolean
  declarationConsent: boolean
}

export interface VehicleRegistrationState {
  hasUnsavedChanges: boolean
  forms: VehicleRegistrationForm[]
  formData: VehicleRegistrationForm
  errors: VehicleRegistrationErrors
  isSubmitting: boolean
  currentStep: number
}
