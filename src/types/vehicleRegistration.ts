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
  id: string
  userId: string
  vehicleId?: string
  plateId?: string
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
  inspectionCode: string
  paymentStatus: 'pending' | 'approved' | 'rejected'
  paymentCode: string
  verificationStatus: 'pending' | 'approved' | 'rejected'
  status: 'pending' | 'approved' | 'rejected'
  submissionDate: string
  expiryDate?: string
  plateNumber?: string
  registrationType: 'New Vehicle' | 'Renewal'
  privacyConsent: boolean
  declarationConsent: boolean
  applicantName?: string
  applicantEmail?: string
  applicantPhone?: string
}

export interface VehicleRegistrationState {
  hasUnsavedChanges: boolean
  forms: VehicleRegistrationForm[]
  formData: VehicleRegistrationForm
  errors: VehicleRegistrationErrors
  isSubmitting: boolean
  currentStep: number
}
