export interface IVolunteerFormState {
  fax_number?: string
  firstName: string
  lastName: string
  address: string
  city: string
  zip: string
  phoneNumber: string
  birthday: string
  age: number | null
  allergies: boolean | null
  emergencyContactName: string
  emergencyContactPhone: string
  volunteerExperience: string
  interestReason: string
  positionPreferences: string[]
  availability: string[]
  nameFull: string
  signatureData: string | null
  signatureDate: string
  parentName: string
  parentSignatureData: string | null
  parentSignatureDate: string
}
