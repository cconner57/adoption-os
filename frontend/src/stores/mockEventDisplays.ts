import { ref } from 'vue'

export interface IEventDevice {
  id: string
  name: string
  model: string
  batteryLevel: number
  status: 'online' | 'offline' | 'updating'
  template: 'featured-grid' | 'welcome-sign' | 'wayfinding' | 'kennel-card'
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
    model: '13.3" E-Ink',
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
    model: '13.3" E-Ink',
    batteryLevel: 78,
    status: 'online',
    template: 'featured-grid',
    config: {
      title: 'Meet Our Seniors',
      subtitle: 'Gentle companions looking for love',
      featuredPetIds: ['2', '4', '8'], 
    },
  },
  {
    id: 'disp-ev-3',
    name: 'Pop-up Wayfinding',
    model: '13.3" E-Ink',
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
  {
    id: 'disp-ev-4',
    name: 'Kennel Card 01',
    model: '4.2" E-Ink',
    batteryLevel: 100,
    status: 'online',
    template: 'kennel-card',
    config: {
      title: 'Kennel Card', 
      featuredPetIds: ['5'], 
    },
  },
])
