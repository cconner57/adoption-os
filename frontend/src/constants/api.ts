export const API_BASE_URL = import.meta.env.VITE_API_URL || ''

export const API_ENDPOINTS = {
  VOLUNTEER_APPLICATION: `${API_BASE_URL}/applications/volunteer`,
  ADOPTION_APPLICATION: `${API_BASE_URL}/applications/adoption`,
  ADOPTED_PETS_COUNT: `${API_BASE_URL}/pets/adopted-count`,
  PET_SPOTLIGHT: `${API_BASE_URL}/pets/spotlight`,
  SURRENDER_APPLICATION: `${API_BASE_URL}/applications/surrender`,
  METRICS: `${API_BASE_URL}/metrics`,
  PETS: `${API_BASE_URL}/pets`,
  PET_PHOTOS: (id: string) => `${API_BASE_URL}/pets/${id}/photos`,
  MARKETING_CAMPAIGNS: `${API_BASE_URL}/v1/marketing/campaigns`, // Using v1
  APPLICATIONS: `${API_BASE_URL}/v1/applications`,
  VOLUNTEERS: `${API_BASE_URL}/v1/volunteers`,
  SHIFTS: `${API_BASE_URL}/v1/shifts`,
}
