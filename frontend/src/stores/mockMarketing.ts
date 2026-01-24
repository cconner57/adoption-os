import { ref } from 'vue'

export interface IRaffleParticipant {
  id: string
  ticketNumber: number
  name: string
  contact: string
  paid: boolean
  date?: string
  paymentMethod?: string
  amount?: string
  comments?: string
}


export interface ICampaign {
  id: string
  name: string
  status: 'active' | 'draft' | 'completed' | 'scheduled' | 'rejected' | 'pending' | 'sent' | 'published' | 'archived'
  startDate: string
  endDate: string
  goal: number
  progress: number
  metric: 'entries' | 'dollars'
  type?: 'raffle' | 'standard'
  prize?: string
  ticketPrice?: number
  participants?: IRaffleParticipant[]
  winnerId?: string
}


export interface INewsletter {
  id: string
  subject: string
  status: string
  sentDate?: string
  recipientCount?: number
  openRate?: number
}

export interface IHappyTail {
  id: string
  petName: string
  adopterName: string
  story: string
  date: string
  status: string
}

export const mockNewsletters = ref<INewsletter[]>([
  {
    id: 'nl-1',
    subject: 'October Shelter Updates 🎃',
    status: 'sent',
    sentDate: '2023-10-31',
    recipientCount: 1250,
    openRate: 42.5,
  },
  {
    id: 'nl-2',
    subject: 'Meet Our Longest Residents',
    status: 'draft',
    recipientCount: 1280,
  },
  {
    id: 'nl-3',
    subject: 'Urgent: Fosters Needed!',
    status: 'scheduled',
    sentDate: '2023-11-05',
    recipientCount: 1260,
  },
])

export const mockHappyTails = ref<IHappyTail[]>([
  {
    id: 'ht-1',
    petName: 'Bella',
    adopterName: 'The Smith Family',
    story: 'Bella has settled in perfectly! She loves playing with the kids in the backyard.',
    date: '2023-10-28',
    status: 'pending',
  },
  {
    id: 'ht-2',
    petName: 'Max',
    adopterName: 'John Doe',
    story: 'Max is the best hiking buddy I could ask for. Thank you for bringing us together.',
    date: '2023-10-15',
    status: 'published',
  },
])

// Mock Participants for Raffle
const mockParticipants = [
  {
    id: 'p-1',
    ticketNumber: 1,
    name: 'Andrea Gladson',
    contact: 'andrea@example.com',
    paid: true,
    date: '2025-12-01',
    paymentMethod: 'Credit Card',
    amount: '$10',
  },
  {
    id: 'p-2',
    ticketNumber: 2,
    name: 'John Smith',
    contact: 'john@example.com',
    paid: true,
    date: '2025-12-02',
    paymentMethod: 'Cash',
    amount: '$10',
  },
  // ... generating 148 more ideally, but we'll cheat for the progress bar by repeating
  ...Array.from({ length: 148 }, (_, i) => ({
    id: `p-${i + 3}`,
    ticketNumber: i + 3,
    name: `Participant ${i + 3}`,
    contact: 'user@example.com',
    paid: true,
    date: '2025-12-05',
    paymentMethod: 'Online',
    amount: '$10',
  }))
]

export const mockCampaigns = ref<ICampaign[]>([
  {
    id: 'cp-1',
    name: '65" Samsung Smart TV Raffle',
    status: 'completed',
    startDate: '2025-11-30',
    endDate: '2025-12-10',
    goal: 150,
    progress: 100, // Calculated dynamically usually, but initial state
    metric: 'entries',
    type: 'raffle',
    prize: '65" Samsung Smart TV',
    ticketPrice: 10,
    participants: mockParticipants,
    winnerId: 'p-1', // Andrea Gladson
  },
  {
    id: 'cp-2',
    name: 'Holiday Giving Drive',
    status: 'active',
    startDate: '2025-12-01',
    endDate: '2025-12-31',
    goal: 5000,
    progress: 45,
    metric: 'dollars',
    type: 'standard',
    participants: [],
  }
])
