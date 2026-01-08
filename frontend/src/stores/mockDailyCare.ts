import { ref } from 'vue'

export interface ICareTaskStatus {
  food: boolean
  pee: boolean
  poo: boolean
  clean: boolean
  play: boolean
  notes: string
}

export interface ICareEntry {
  id: string
  petId: string
  petName: string
  cageNumber: number | string
  am: ICareTaskStatus
  pm: ICareTaskStatus
}

export interface IGeneralTasks {
  sweep: boolean
  litterScoop: boolean
  bowls: boolean
  water: boolean
  trash: boolean
  windows: boolean
  mop: boolean
  donationsCollected: string
  applicationsInfo: string
  volunteerRequests: string
  additionalNotes: string
}

export interface IDailyLog {
  date: string // YYYY-MM-DD
  amVolunteer: string
  pmVolunteer: string
  entries: ICareEntry[]
  amTasks: IGeneralTasks
  pmTasks: IGeneralTasks
}

// Helper to create empty status
const emptyStatus = (): ICareTaskStatus => ({
  food: false,
  pee: false,
  poo: false,
  clean: false,
  play: false,
  notes: '',
})

const emptyTasks = (): IGeneralTasks => ({
  sweep: false,
  litterScoop: false,
  bowls: false,
  water: false,
  trash: false,
  windows: false,
  mop: false,
  donationsCollected: '',
  applicationsInfo: '',
  volunteerRequests: '',
  additionalNotes: '',
})

// Mock Data for "Today"
export const mockDailyLogs = ref<IDailyLog[]>([
  {
    date: new Date().toISOString().split('T')[0],
    amVolunteer: 'Alejandra',
    pmVolunteer: '',
    entries: [
      {
        id: 'ce-1',
        petId: 'p-1',
        petName: 'Allison',
        cageNumber: 1,
        am: { food: true, pee: true, poo: true, clean: true, play: false, notes: '' },
        pm: emptyStatus(),
      },
      {
        id: 'ce-2',
        petId: 'p-2',
        petName: 'Domi',
        cageNumber: 2,
        am: {
          food: true,
          pee: true,
          poo: true,
          clean: true,
          play: false,
          notes: 'Really wants attention ♡',
        },
        pm: emptyStatus(),
      },
      {
        id: 'ce-3',
        petId: 'p-3',
        petName: 'Sally',
        cageNumber: 3,
        am: {
          food: true,
          pee: true,
          poo: false,
          clean: true,
          play: true,
          notes: 'Seems down. No poop. Did not touch water.',
        },
        pm: emptyStatus(),
      },
      {
        id: 'ce-4',
        petId: 'p-4',
        petName: 'Colby',
        cageNumber: 4,
        am: { food: true, pee: true, poo: true, clean: true, play: true, notes: '' },
        pm: emptyStatus(),
      },
      {
        id: 'ce-5',
        petId: 'p-5',
        petName: 'Purina',
        cageNumber: 5,
        am: {
          food: true,
          pee: true,
          poo: false,
          clean: true,
          play: false,
          notes: 'Teary-eyed; looked like he was crying :(',
        },
        pm: emptyStatus(),
      },
      {
        id: 'ce-6',
        petId: 'p-6',
        petName: 'Champu & Churro',
        cageNumber: 6,
        am: { food: true, pee: true, poo: true, clean: true, play: true, notes: '' },
        pm: emptyStatus(),
      },
      {
        id: 'ce-7',
        petId: 'p-7',
        petName: 'Mylo',
        cageNumber: 7,
        am: { food: true, pee: true, poo: true, clean: true, play: true, notes: '' },
        pm: emptyStatus(),
      },
    ],
    amTasks: {
      sweep: true,
      litterScoop: true,
      bowls: true,
      water: true,
      trash: true,
      windows: false,
      mop: false,
      donationsCollected: '',
      applicationsInfo: '',
      volunteerRequests: '',
      additionalNotes:
        "Purina's eyes look very watery - either from irritation, allergies or is sad.",
    },
    pmTasks: emptyTasks(),
  },
])
