
export interface IShiftData {
  date: string
  startTime: string
  endTime: string
  role: string
  isRecurring?: boolean
  endDate?: string
  frequency?: 'weekly' | 'biweekly' | 'monthly'
}

export interface IShiftCreation {
  volunteerId: string
  date: string
  startTime: string
  endTime: string
  role: string
}

export function generateShifts(shiftData: IShiftData, selectedId: string): IShiftCreation[] {
  const newShifts: IShiftCreation[] = []
  const baseDate = new Date(shiftData.date)

  if (shiftData.isRecurring) {
    const endDate = shiftData.endDate
      ? new Date(shiftData.endDate)
      : new Date(baseDate.getTime() + 90 * 24 * 60 * 60 * 1000)

    const currentDate = new Date(baseDate)
    let count = 0

    while (currentDate <= endDate && count < 50) {
      newShifts.push({
        volunteerId: selectedId,
        date: currentDate.toISOString().split('T')[0],
        startTime: shiftData.startTime,
        endTime: shiftData.endTime,
        role: shiftData.role,
      })

      if (shiftData.frequency === 'weekly') {
        currentDate.setDate(currentDate.getDate() + 7)
      } else if (shiftData.frequency === 'biweekly') {
        currentDate.setDate(currentDate.getDate() + 14)
      } else if (shiftData.frequency === 'monthly') {
        currentDate.setMonth(currentDate.getMonth() + 1)
      } else {
        break
      }
      count++
    }
  } else {
    newShifts.push({
      volunteerId: selectedId,
      date: shiftData.date,
      startTime: shiftData.startTime,
      endTime: shiftData.endTime,
      role: shiftData.role,
    })
  }
  return newShifts
}
