import { ref } from 'vue'

export interface IDisplayDevice {
  id: string
  macAddress: string
  batteryLevel: number 
  signalStrength: number 
  assignedPetId: string | null
  status: 'online' | 'offline' | 'syncing'
  lastSync: string
  template: 'standard' | 'medical' | 'urgent'
}

export const mockDevices = ref<IDisplayDevice[]>([
  {
    id: 'tag-001',
    macAddress: 'A1:B2:C3:D4:E5:F6',
    batteryLevel: 85,
    signalStrength: 92,
    assignedPetId: '1', 
    status: 'online',
    lastSync: '10 mins ago',
    template: 'standard',
  },
  {
    id: 'tag-002',
    macAddress: '1A:2B:3C:4D:5E:6F',
    batteryLevel: 42,
    signalStrength: 78,
    assignedPetId: '2',
    status: 'online',
    lastSync: '1 hour ago',
    template: 'urgent',
  },
  {
    id: 'tag-003',
    macAddress: 'AA:BB:CC:DD:EE:FF',
    batteryLevel: 12,
    signalStrength: 45,
    assignedPetId: null, 
    status: 'offline',
    lastSync: '2 days ago',
    template: 'standard',
  },
  {
    id: 'tag-004',
    macAddress: '11:22:33:44:55:66',
    batteryLevel: 98,
    signalStrength: 88,
    assignedPetId: '4', 
    status: 'syncing',
    lastSync: 'Just now',
    template: 'medical',
  },
])
