import { computed,ref } from 'vue'

export interface ITransaction {
  id: string
  date: string
  donorName: string
  amount: number
  type: 'adoption_fee' | 'donation' | 'sponsorship' | 'merch'
  method: 'Credit Card' | 'Cash' | 'Check' | 'PayPal'
  status: 'completed' | 'pending' | 'failed' | 'refunded'
  notes?: string
  relatedPetName?: string
}

export const mockTransactions = ref<ITransaction[]>([
  {
    id: 'tx-1001',
    date: '2023-10-28',
    donorName: 'Alice Johnson',
    amount: 150.0,
    type: 'adoption_fee',
    method: 'Credit Card',
    status: 'completed',
    relatedPetName: 'Rex',
    notes: 'Adoption fee for Rex',
  },
  {
    id: 'tx-1002',
    date: '2023-10-28',
    donorName: 'Bob Smith',
    amount: 50.0,
    type: 'donation',
    method: 'PayPal',
    status: 'completed',
    notes: 'In honor of my late dog, Spot',
  },
  {
    id: 'tx-1003',
    date: '2023-10-27',
    donorName: 'Corporate Sponsors Inc.',
    amount: 500.0,
    type: 'sponsorship',
    method: 'Check',
    status: 'pending',
    relatedPetName: 'Luna',
    notes: 'Yearly sponsorship for Luna',
  },
  {
    id: 'tx-1004',
    date: '2023-10-27',
    donorName: 'Walk-in Guest',
    amount: 20.0,
    type: 'donation',
    method: 'Cash',
    status: 'completed',
    notes: 'Donation jar from front desk',
  },
  {
    id: 'tx-1005',
    date: '2023-10-26',
    donorName: 'Sarah Jenkins',
    amount: 25.0,
    type: 'merch',
    method: 'Credit Card',
    status: 'completed',
    notes: 'T-shirt purchase',
  },
])

export const donationStats = computed(() => {
  const total = mockTransactions.value
    .filter((t) => t.status === 'completed')
    .reduce((sum, t) => sum + t.amount, 0)

  const thisMonth = mockTransactions.value
    .filter((t) => t.date.includes('2023-10')) 
    .reduce((sum, t) => sum + t.amount, 0)

  return {
    totalRevenue: total,
    monthRevenue: thisMonth,
    recentCount: mockTransactions.value.length,
  }
})
