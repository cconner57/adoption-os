export interface IShift {
  id: number | string
  volunteerId?: string
  volunteerName: string
  role: string
  startTime: string // "HH:MM:SS" or ISO
  endTime: string // "HH:MM:SS" or ISO
  date: string // "YYYY-MM-DD"
  status?: 'scheduled' | 'imcomplete' | 'completed' | 'canceled'
}

export interface IVolunteer {
  id: string
  firstName: string
  lastName: string
  email: string
  phone?: string
  status: 'active' | 'inactive' | 'pending'
  roles?: string[]
}
