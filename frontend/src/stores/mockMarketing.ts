import { ref } from 'vue'

export interface ICampaign {
  id: string
  name: string
  status: 'active' | 'draft' | 'completed'
  startDate: string
  endDate: string
  goal: string
  progress: number // 0-100
  metric: string // e.g., "$ raised" or "adoptions"
}

export interface INewsletter {
  id: string
  subject: string
  status: 'sent' | 'draft' | 'scheduled'
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
  status: 'published' | 'pending' | 'rejected'
  photoUrl?: string // Placeholder for now
}

export const mockCampaigns = ref<ICampaign[]>([
  {
    id: 'camp-1',
    name: 'Holiday Adoption Drive',
    status: 'active',
    startDate: '2023-11-01',
    endDate: '2023-12-31',
    goal: '50 Adoptions',
    progress: 72,
    metric: 'adoptions',
  },
  {
    id: 'camp-2',
    name: 'Winter Fundraiser',
    status: 'draft',
    startDate: '2023-12-01',
    endDate: '2023-12-31',
    goal: '$10,000',
    progress: 0,
    metric: '$ raised',
  },
  {
    id: 'camp-3',
    name: 'Spay/Neuter Awareness',
    status: 'completed',
    startDate: '2023-08-01',
    endDate: '2023-08-31',
    goal: '200 Procedures',
    progress: 100,
    metric: 'procedures',
  },
])

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
