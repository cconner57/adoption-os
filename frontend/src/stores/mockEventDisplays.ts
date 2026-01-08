import { ref } from 'vue'

export interface IEventDevice {
  id: string
  name: string
  batteryLevel: number
  status: 'online' | 'offline' | 'updating'
  template: 'featured-grid' | 'welcome-sign' | 'wayfinding'
  config: {
    title: string
    subtitle?: string
    featuredPetIds: string[]
    direction?: 'left' | 'right' | 'straight'
  }
}

export const mockEventDisplays = ref<IEventDevice[]>([
  {
    id: 'disp-ev-1',
    name: 'Lobby Sign A (Large)',
    batteryLevel: 94,
    status: 'online',
    template: 'welcome-sign',
    config: {
      title: 'Welcome to Paws & Claws',
      subtitle: "Today's Adoption Event: 10AM - 4PM",
      featuredPetIds: [],
      direction: 'straight',
    },
  },
  {
    id: 'disp-ev-2',
    name: 'Table Sign 1',
    batteryLevel: 78,
    status: 'online',
    template: 'featured-grid',
    config: {
      title: 'Meet Our Seniors',
      subtitle: 'Gentle companions looking for love',
      featuredPetIds: ['2', '4', '8'], // IDs from mockPets
    },
  },
  {
    id: 'disp-ev-3',
    name: 'Pop-up Wayfinding',
    batteryLevel: 25,
    status: 'offline',
    template: 'wayfinding',
    config: {
      title: 'Cat Corner',
      subtitle: 'Meet the kittens!',
      featuredPetIds: [],
      direction: 'right',
    },
  },
])
