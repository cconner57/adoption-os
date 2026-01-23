import { ref } from 'vue'

export interface IKioskSettings {
  general: {
    kioskName: string
    welcomeMessage: string
    timeoutSeconds: number
  }
  theme: {
    primaryColor: string
    backgroundImage: string 
    logoUrl: string
  }
  features: {
    allowCheckIn: boolean
    showAvailablePets: boolean
    showEvents: boolean
    donationsEnabled: boolean
  }
}

export const mockKioskSettings = ref<IKioskSettings>({
  general: {
    kioskName: 'Lobby Kiosk 1',
    welcomeMessage: 'Welcome to Paws & Claws Adoption Center! 🐾',
    timeoutSeconds: 60,
  },
  theme: {
    primaryColor: '#8b5cf6', 
    backgroundImage:
      'https://images.unsplash.com/photo-1450778869180-41d0601e046e?q=80&w=2500&auto=format&fit=crop', 
    logoUrl: '', 
  },
  features: {
    allowCheckIn: true,
    showAvailablePets: true,
    showEvents: false,
    donationsEnabled: true,
  },
})
