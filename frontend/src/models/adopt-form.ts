export interface CurrentPet {
  name: string
  speciesBreedSize: string
  age: string
  source: string
  spayedNeutered: string
  likesDogs: string
}

export interface PastPet {
  name: string
  speciesBreedSize: string
  age: string
  source: string
  spayedNeutered: string
  passedAwayReason: string
}

export interface FormState {
  fax_number: string | null // Honeypot
  firstName: string | null
  lastName: string | null
  age: number | null
  spouseFirstName: string | null
  spouseLastName: string | null
  roommatesNames: string[]
  childrenNamesAges: { name: string; age: string }[]
  currentPets: CurrentPet[]
  currentlyHavePets: string | null

  pastPets: PastPet[]
  ownPetsBefore: string | null
  email: string | null
  address: string | null
  addressLine2: string | null
  city: string | null
  state: string | null
  zip: string | null
  phoneNumber: string | null
  cellPhoneNumber: string | null
  adultMembersAgreed: string | null
  homeType: string | null
  homeOwnership: string | null
  landlordName: string | null
  landlordPhoneNumber: string | null
  allowPets: string | null
  breedRestrictionsWeightLimit: string | null
  monthlyFee: string | null
  allergies: string | null
  primaryOwner: boolean | null
  yearsAtAddress: string | null
  previousAddress: string | null
  expectToMove: string | null
  travelPlan: string | null
  catAccess: string | null
  catIndoorOutdoor: string | null
  catPreferenceBreed: string | null
  catPreferencePhysical: string | null
  catPreferencePersonality: string | null
  catPreferenceNotWant: string | null
  whyInterested: string | null
  adoptionReason: string | null
  ownCatBefore: string | null
  ownKittenBefore: string | null
  alreadyHaveVeterinarian: string | null
  catAllowedHomeArea: string | null
  catHomeAloneHours: string | null
  catDisciplineType: string | null
  catEscapeSteps: string | null

  bredAnimalDescription: string | null

  ownedDeclawedOrDebarked: string | null
  movedWithPet: string | null
  ownedSpecialNeedsPet: string | null
  mobilityDevice: string | null
  surrenderConditions: string[]
  surrenderPlan: string | null

  foodTypeBrand: string | null
  affordVetCare: string | null
  affordEmergencyCost: string | null

  agreementSignature1: string | null

  agreementSignature2: string | null
  agreementSignature3: string | null
  signatureData: string | null
}
