export interface IShift {
  id: number | string
  volunteerId?: string
  volunteerName: string
  role: string
  startTime: string 
  endTime: string 
  date: string 
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
