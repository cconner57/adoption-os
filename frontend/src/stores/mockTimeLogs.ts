import { ref } from 'vue'

export interface ITimeLog {
  id: string
  volunteerName: string
  date: string
  timeIn: string
  timeOut: string | null
  duration: string | null
  taskDescription: string
  status: 'approved' | 'pending' | 'flagged'
}

export interface IIncident {
  id: string
  date: string
  time: string
  reportedBy: string
  type: 'injury' | 'behavioral' | 'damage' | 'other'
  description: string
  severity: 'low' | 'medium' | 'high' | 'critical'
  status: 'open' | 'investigating' | 'resolved'
}

export const mockTimeLogs = ref<ITimeLog[]>([
  {
    id: 'tl-1',
    volunteerName: 'Sarah Jenkins',
    date: '2023-10-27',
    timeIn: '08:00 AM',
    timeOut: '12:00 PM',
    duration: '4h 0m',
    taskDescription: 'Morning shifts - dog walking and kennel cleaning.',
    status: 'approved',
  },
  {
    id: 'tl-2',
    volunteerName: 'Mike Ross',
    date: '2023-10-27',
    timeIn: '09:00 AM',
    timeOut: '02:00 PM',
    duration: '5h 0m',
    taskDescription: 'Assisted with adoption event setup.',
    status: 'pending',
  },
  {
    id: 'tl-3',
    volunteerName: 'Jessica Pearson',
    date: '2023-10-26',
    timeIn: '10:00 AM',
    timeOut: '11:30 AM',
    duration: '1h 30m',
    taskDescription: 'Cat socialization.',
    status: 'approved',
  },
  {
    id: 'tl-4',
    volunteerName: 'Harvey Specter',
    date: '2023-10-28',
    timeIn: '08:00 AM',
    timeOut: null, // Currently clocked in
    duration: null,
    taskDescription: 'Lobby greeter.',
    status: 'pending',
  },
])

export const mockIncidents = ref<IIncident[]>([
  {
    id: 'inc-1',
    date: '2023-10-25',
    time: '02:30 PM',
    reportedBy: 'Mike Ross',
    type: 'behavioral',
    description:
      'Dog "Buster" showed aggression towards another dog in the play yard. Separated immediately.',
    severity: 'medium',
    status: 'investigating',
  },
  {
    id: 'inc-2',
    date: '2023-10-20',
    time: '09:00 AM',
    reportedBy: 'Sarah Jenkins',
    type: 'injury',
    description: 'Slipped on wet floor in kennel 4. Minor scrape on knee.',
    severity: 'low',
    status: 'resolved',
  },
])
