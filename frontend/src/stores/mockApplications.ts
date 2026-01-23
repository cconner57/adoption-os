import { ref } from 'vue'

export interface IApplication {
  id: string
  type: 'volunteer' | 'surrender' | 'adoption'
  applicantName: string
  email: string
  date: string
  status: 'pending' | 'approved' | 'denied' | 'needs_info'
  notes?: string
  details: Record<string, unknown>
  fullApplication: Record<string, unknown> // The raw form data
}

export const mockApplications = ref<IApplication[]>([
  // Volunteer Applications
  {
    id: 'v-app-1',
    type: 'volunteer',
    applicantName: 'Sarah Jenkins',
    email: 'sarah.j@example.com',
    date: '2023-10-25',
    status: 'pending',
    details: {
      role: 'Dog Walker',
      experience: 'Previous shelter volunteer',
      availability: 'Weekends',
    },
    fullApplication: {
      Phone: '555-0101',
      Address: '123 Maple St, Springfield',
      'Why do you want to volunteer?': 'I love dogs and have free time on weekends.',
      'Previous Experience': 'Volunteered at City Shelter for 2 years.',
      References: 'John Doe (555-0102)',
      'Emergency Contact': 'Jane Jenkins (Mother) - 555-0103',
    },
  },
  {
    id: 'v-app-2',
    type: 'volunteer',
    applicantName: 'Mike Ross',
    email: 'mike.ross@example.com',
    date: '2023-10-24',
    status: 'needs_info',
    notes: 'Reference check needed',
    details: {
      role: 'Event Helper',
      experience: 'None',
      availability: 'Evenings',
    },
    fullApplication: {
      Phone: '555-0202',
      Address: '456 Oak Ave, Springfield',
      'Why do you want to volunteer?': 'To meet people and help out.',
      'Previous Experience': 'None officially.',
      References: 'Harvey Specter (555-0203)',
      'Emergency Contact': 'Rachel Zane - 555-0204',
    },
  },

  // Surrender Applications
  {
    id: 's-app-1',
    type: 'surrender',
    applicantName: 'John Doe',
    email: 'john.doe@example.com',
    date: '2023-10-26',
    status: 'pending',
    details: {
      petName: 'Buddy',
      species: 'Dog',
      reason: 'Moving',
      urgent: true,
    },
    fullApplication: {
      'Pet Name': 'Buddy',
      Species: 'Dog',
      Breed: 'Golden Retriever Mix',
      Age: '4 years',
      Sex: 'Male',
      Fixed: 'Yes',
      Vaccinated: 'Yes',
      'Reason for Surrender': 'Moving to apartment that does not allow pets.',
      'Good with kids?': 'Yes',
      'Good with other dogs?': 'Yes',
      'Bite History': 'None',
    },
  },
  {
    id: 's-app-2',
    type: 'surrender',
    applicantName: 'Emily Clark',
    email: 'emily.c@example.com',
    date: '2023-10-22',
    status: 'denied',
    notes: 'Full capacity for cats',
    details: {
      petName: 'Whiskers',
      species: 'Cat',
      reason: 'Allergies',
      urgent: false,
    },
    fullApplication: {
      'Pet Name': 'Whiskers',
      Species: 'Cat',
      Breed: 'Domestic Shorthair',
      Age: '2 years',
      Sex: 'Female',
      Fixed: 'Yes',
      'Reason for Surrender': 'Developed severe allergies.',
      'Indoor/Outdoor': 'Indoor only',
    },
  },

  // Adoption Applications
  {
    id: 'a-app-1',
    type: 'adoption',
    applicantName: 'David Lee',
    email: 'david.lee@example.com',
    date: '2023-10-26',
    status: 'pending',
    details: {
      petName: 'Rex',
      petId: 'p-101',
      homeType: 'House with yard',
      otherPets: true,
    },
    fullApplication: {
      Applicant: 'David Lee',
      Phone: '555-0303',
      Address: '789 Pine Ln, Springfield',
      Housing: 'Own house with fenced yard',
      'Landlord Contact': 'N/A',
      'Household Members': 'Wife and 2 kids (ages 8, 10)',
      'Current Pets': '1 Cat (Spayed female)',
      'Vet Reference': 'Springfield Vet Clinic (Dr. Hibbert)',
      'Time alone': '4 hours per day',
    },
  },
  {
    id: 'a-app-2',
    type: 'adoption',
    applicantName: 'Amanda Smith',
    email: 'amanda.s@example.com',
    date: '2023-10-23',
    status: 'approved',
    details: {
      petName: 'Luna',
      petId: 'p-102',
      homeType: 'Apartment',
      otherPets: false,
    },
    fullApplication: {
      Applicant: 'Amanda Smith',
      Phone: '555-0404',
      Address: '321 Elm St, Apt 4B',
      Housing: 'Apartment (Pet friendly)',
      'Landlord Contact': 'Mr. Roper (555-0405)',
      'Household Members': 'Just me',
      'Current Pets': 'None',
      'Vet Reference': 'N/A',
      'Time alone': '8 hours (work from home 2 days)',
    },
  },
])
