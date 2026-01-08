import { ref } from 'vue'

export interface IDriver {
  id: string
  name: string
  phone: string
  vehicle: string // "Toyota RAV4 (Silver)"
  status: 'available' | 'busy' | 'offline'
  rating: number
  avatar?: string
}

export interface IChatMessage {
  id: string
  sender: 'driver' | 'dispatch' | 'system'
  text: string
  timestamp: string
}

export interface IPetPassenger {
  id: string
  name: string
  reason: string // e.g. "Surgery", "Checkup", "Dental"
  status: 'adopted' | 'shelter'
  description: string
}

export interface ITrip {
  id: string
  direction: 'to_vet' | 'from_vet' // Morning vs Afternoon
  status:
    | 'pending'
    | 'assigned'
    | 'en_route_vet'
    | 'at_vet'
    | 'en_route_shelter'
    | 'at_shelter'
    | 'completed'
    | 'delayed'
    | 'incident'
  date: string
  pickupTime: string // "08:00 AM" or "04:00 PM"
  pets: IPetPassenger[]
  driverId?: string
  driverNotes?: string // For vehicle description updates specific to this trip
  estimatedDuration: string
  messages: IChatMessage[]
}

export const mockDrivers = ref<IDriver[]>([
  {
    id: 'd-1',
    name: 'Alice Cooper',
    phone: '555-0101',
    vehicle: 'Toyota RAV4 (Silver)',
    status: 'available',
    rating: 4.8,
    avatar: 'AC',
  },
  {
    id: 'd-2',
    name: 'Bob Marley',
    phone: '555-0102',
    vehicle: 'Ford F-150 (Blue)',
    status: 'busy',
    rating: 5.0,
    avatar: 'BM',
  },
])

export const mockTrips = ref<ITrip[]>([
  {
    id: 't-201',
    direction: 'to_vet',
    status: 'en_route_vet',
    date: '2023-10-27',
    pickupTime: '08:30 AM',
    pets: [
      {
        id: 'p-1',
        name: 'Rex',
        reason: 'Neuter Surgery',
        status: 'shelter',
        description: 'Large Golden Retriever mix, friendly but pulls on leash.',
      },
      {
        id: 'p-2',
        name: 'Luna',
        reason: 'Dental Cleaning',
        status: 'adopted',
        description: 'Small Beagle, nervous around loud noises.',
      },
      {
        id: 'p-3',
        name: 'Spot',
        reason: 'Leg X-Ray',
        status: 'shelter',
        description: 'Dalmatian puppy, high energy.',
      },
    ],
    driverId: 'd-1',
    driverNotes: 'Driving the Silver RAV4 today.',
    estimatedDuration: '20 mins',
    messages: [
      { id: 'm-1', sender: 'system', text: 'Trip started at 8:35 AM', timestamp: '8:35 AM' },
    ],
  },
  {
    id: 't-202',
    direction: 'from_vet',
    status: 'pending',
    date: '2023-10-27',
    pickupTime: '04:00 PM',
    pets: [
      {
        id: 'p-1',
        name: 'Rex',
        reason: 'Post-Op Pickup',
        status: 'shelter',
        description: 'Large Golden Retriever mix.',
      },
      {
        id: 'p-2',
        name: 'Luna',
        reason: 'Post-Op Pickup',
        status: 'adopted',
        description: 'Small Beagle.',
      },
    ],
    driverId: undefined, // Needs volunteer sign up
    estimatedDuration: '25 mins',
    messages: [],
  },
  {
    id: 't-203',
    direction: 'to_vet',
    status: 'at_vet',
    date: '2023-10-27',
    pickupTime: '09:00 AM',
    pets: [
      {
        id: 'p-4',
        name: 'Bella',
        reason: 'Checkup',
        status: 'shelter',
        description: 'Senior cat, must stay in carrier.',
      },
    ],
    driverId: 'd-2',
    estimatedDuration: '15 mins',
    messages: [
      {
        id: 'm-2',
        sender: 'driver',
        text: 'Arrived safely. Checking her in now.',
        timestamp: '09:15 AM',
      },
    ],
  },
])
