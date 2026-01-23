
export interface IVolunteer {
  id: string
  firstName: string
  lastName: string
  email: string
  phone: string
  status: 'active' | 'inactive' | 'pending'
  role: 'Tier 1' | 'Tier 2' | 'Teen' | 'Admin' | 'Super Admin'
  joinDate: string
  bio?: string
  photoUrl?: string
  skills: string[]
  reliabilityScore: number 

  address: string
  city: string
  zip: string
  birthday: string
  age?: number
  allergies: boolean
  emergencyContactName: string
  emergencyContactPhone: string
  volunteerExperience?: string
  interestReason?: string
  positionPreferences: string[]
  availability: string[] 

  totalHours: number
  streak: number 
  badges: string[] 
}

export interface IShift {
  id: string
  volunteerId: string
  date: string 
  startTime: string 
  endTime: string 
  role: string 
  status:
    | 'scheduled'
    | 'all_good'
    | 'late'
    | 'no_show'
    | 'cancelled'
    | 'completed'
    | 'missed'
    | 'covered'
    | 'covered 24h'
    | 'covered_24h'
    | 'covered late'
    | 'covered_late'
    | 'covered <24h notice'
    | 'covered_less_24h'
    | 'covered <1h notice'
    | 'covered_less_1h'
  notes?: string
}

export interface IIncident {
  id: string
  volunteerId: string
  date: string 
  severity: 'low' | 'medium' | 'high'
  description: string
  recordedBy: string
}

export const mockVolunteers: IVolunteer[] = []

export const mockShifts: IShift[] = [
  
  {
    id: 's-t2-1-1',
    volunteerId: 't2-1',
    date: '2023-05-15',
    startTime: '09:00',
    endTime: '13:00',
    role: 'Feeding/Cleaning',
    status: 'all_good',
  },
  {
    id: 's-t2-1-2',
    volunteerId: 't2-1',
    date: '2023-08-20',
    startTime: '09:00',
    endTime: '13:00',
    role: 'Feeding/Cleaning',
    status: 'all_good',
  },
  {
    id: 's-t2-1-3',
    volunteerId: 't2-1',
    date: '2024-01-10',
    startTime: '09:00',
    endTime: '13:00',
    role: 'Feeding/Cleaning',
    status: 'all_good',
  },
  {
    id: 's-t2-1-4',
    volunteerId: 't2-1',
    date: '2024-01-17',
    startTime: '09:00',
    endTime: '13:00',
    role: 'Feeding/Cleaning',
    status: 'scheduled',
  },

  {
    id: 's1',
    volunteerId: 'v1',
    date: '2024-01-05',
    startTime: '08:00',
    endTime: '12:00',
    role: 'Feeding/Cleaning',
    status: 'all_good',
  },
  {
    id: 's2',
    volunteerId: 'v1',
    date: '2024-01-12',
    startTime: '08:00',
    endTime: '12:00',
    role: 'Feeding/Cleaning',
    status: 'scheduled',
  },
  
  {
    id: 's3',
    volunteerId: 'v2',
    date: '2024-01-06',
    startTime: '13:00',
    endTime: '17:00',
    role: 'Cat Socializing',
    status: 'all_good',
  },
  {
    id: 's4',
    volunteerId: 'v2',
    date: '2024-01-07',
    startTime: '13:00',
    endTime: '17:00',
    role: 'Feeding/Cleaning',
    status: 'no_show',
    notes: 'Did not call.',
  },
]

export const mockIncidents: IIncident[] = [
  {
    id: 'i1',
    volunteerId: 'v2',
    date: '2024-01-07',
    severity: 'low',
    description: 'No-show for scheduled cleaning shift without prior notice.',
    recordedBy: 'Admin',
  },
]
export const availableBadges = [
  
  { id: 'b1', name: 'New Recruit', icon: '🌱', description: 'Completed your first shift!' },
  { id: 'b2', name: 'Rookie', icon: '⭐', description: 'Active for 1 month.' },
  { id: 'b3', name: 'Regular', icon: '🌟', description: 'Active for 6 months.' },
  { id: 'b4', name: 'Veteran', icon: '🎖️', description: 'Active for 1 year.' },
  { id: 'b5', name: 'Lifer', icon: '🏅', description: 'Active for 3 years.' },
  { id: 'b6', name: 'Legend', icon: '👑', description: 'Active for 5+ years.' },

  { id: 'h1', name: '10 Hours', icon: '🕙', description: 'Contributed 10 hours of service.' },
  { id: 'h2', name: '50 Hours', icon: '🕔', description: 'Contributed 50 hours of service.' },
  {
    id: 'h3',
    name: 'Bronze Service',
    icon: '🥉',
    description: 'Contributed 100 hours of service.',
  },
  {
    id: 'h4',
    name: 'Silver Service',
    icon: '🥈',
    description: 'Contributed 250 hours of service.',
  },
  { id: 'h5', name: 'Gold Service', icon: '🥇', description: 'Contributed 500 hours of service.' },
  {
    id: 'h6',
    name: 'Platinum Service',
    icon: '💎',
    description: 'Contributed 1000+ hours of service.',
  },

  { id: 'd1', name: 'Dog Walker', icon: '🐕', description: 'Completed 10 Dog Walking shifts.' },
  { id: 'd2', name: 'Power Walker', icon: '👟', description: 'Completed 50 Dog Walking shifts.' },
  {
    id: 'd3',
    name: 'Leash Master',
    icon: '🦮',
    description: 'Handled a difficult dog successfully.',
  },
  { id: 'd4', name: 'Poop Scoop Pro', icon: '💩', description: 'The unsung hero of the yard.' },
  { id: 'd5', name: 'Puppy Pal', icon: '🐾', description: 'Helped with puppy socialization.' },
  { id: 'd6', name: 'Bark Ranger', icon: '🏞️', description: 'Took a dog on a field trip.' },

  {
    id: 'c1',
    name: 'Cat Cuddler',
    icon: '🐱',
    description: 'Completed 10 Cat Socializing shifts.',
  },
  {
    id: 'c2',
    name: 'Cat Whisperer',
    icon: '🐈‍⬛',
    description: 'Completed 50 Cat Socializing shifts.',
  },
  { id: 'c3', name: 'Litter Box Hero', icon: '🧹', description: 'Cleaned the cat room 20 times.' },
  {
    id: 'c4',
    name: 'Laser Pointer Pro',
    icon: '🔴',
    description: 'Expert at feline entertainment.',
  },
  { id: 'c5', name: 'Kitty Comforter', icon: '🧶', description: 'Socialized shy cats.' },
  {
    id: 'c6',
    name: 'Meowtain Climber',
    icon: '🧗',
    description: 'Helped maximize vertical cat space.',
  },

  { id: 's1', name: 'Bunny Buddy', icon: '🐰', description: 'Cared for rabbits.' },
  { id: 's2', name: 'Pocket Pet Pal', icon: '🐹', description: 'Cared for hamsters/guinea pigs.' },
  { id: 'o1', name: 'Cleaning Crew', icon: '🧼', description: 'Completed 20 Cleaning shifts.' },
  { id: 'o2', name: 'Laundry Lord', icon: '🧺', description: 'Folded 100 loads of laundry.' },
  { id: 'o3', name: 'Dishwasher Dynamo', icon: '🍽️', description: 'Kept the bowls sparkling.' },
  {
    id: 'o4',
    name: 'Front Desk Ace',
    icon: '☎️',
    description: 'Managed the reception desk like a pro.',
  },

  { id: 't1', name: 'Early Bird', icon: '🌅', description: 'Shift started before 9 AM.' },
  { id: 't2', name: 'Night Owl', icon: '🌙', description: 'Shift ended after 6 PM.' },
  { id: 't3', name: 'Weekend Warrior', icon: '📅', description: '10 weekend shifts completed.' },
  { id: 't4', name: 'Holiday Hero', icon: '🎄', description: 'Volunteered on a major holiday.' },
  { id: 't5', name: 'Double Header', icon: '✌️', description: 'Completed 2 shifts in one day.' },
  {
    id: 't6',
    name: 'Last Minute Savior',
    icon: '🚑',
    description: 'Picked up a shift with < 24h notice.',
  },

  { id: 'r1', name: 'Reliable', icon: '🛡️', description: '100% attendance for 3 months.' },
  { id: 'r2', name: 'Perfect Year', icon: '📅', description: '100% attendance for 12 months.' },
  { id: 'r3', name: 'Mentor', icon: '🎓', description: 'Trained a new volunteer.' },
  { id: 'r4', name: 'Photographer', icon: '📸', description: 'Took adoption photos for pets.' },
  {
    id: 'r5',
    name: 'Matchmaker',
    icon: '💘',
    description: 'Helped a pet find their forever home!',
  },
  { id: 'r6', name: 'Shelter Guardian', icon: '👼', description: 'Donated supplies or funds.' },

  {
    id: 'f1',
    name: 'The Treat Dispenser',
    icon: '🍪',
    description: 'Dogs know you have the goods.',
  },
  {
    id: 'f2',
    name: 'Patient Saint',
    icon: '🧘',
    description: 'Spent >1 hour with a single shy animal.',
  },
  {
    id: 'f3',
    name: 'Rising Star',
    icon: '🚀',
    description: 'Completed first 20 hours of volunteering.',
  },
  { id: 'f4', name: 'Event Star', icon: '🎪', description: 'Helped at an adoption event.' },
  {
    id: 'f5',
    name: 'Transport Team',
    icon: '🚗',
    description: 'Drove an animal to the vet/rescue.',
  },
]
