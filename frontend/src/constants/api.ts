export const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

export const API_ENDPOINTS = {
  VOLUNTEER_APPLICATION: `${API_BASE_URL}/applications/volunteer`,
  ADOPTION_APPLICATION: `${API_BASE_URL}/applications/adoption`,
  ADOPTED_PETS_COUNT: `${API_BASE_URL}/pets/adopted-count`,
  PET_SPOTLIGHT: `${API_BASE_URL}/pets/spotlight`,
  SURRENDER_APPLICATION: `${API_BASE_URL}/applications/surrender`,
}
