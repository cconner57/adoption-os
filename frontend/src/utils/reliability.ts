import type { IShift } from '../stores/mockVolunteerData'

export function calculateReliabilityScore(shifts: IShift[]): number {
  let score = 0
  shifts.forEach((shift) => {
    
    if (shift.status === 'scheduled') return

    switch (shift.status) {
      case 'completed':
      case 'all_good':
        score += 20
        break
      case 'late':
        score += 10
        break
      case 'missed':
      case 'no_show':
        score -= 50
        break
      case 'covered':
      case 'covered 24h':
      case 'covered_24h':
        score -= 5
        break
      case 'covered late':
      case 'covered_late':
      case 'covered <24h notice':
      case 'covered_less_24h':
        score -= 10
        break
      case 'covered <1h notice':
      case 'covered_less_1h':
        score -= 20
        break
    }

    if (shift.notes && shift.notes.includes('Covering for')) {
      if (shift.notes.includes('>24h notice')) {
        score += 10
      } else if (shift.notes.includes('<24h notice')) {
        score += 20
      }
    }
  })
  return score
}

export function calculateTotalHours(shifts: IShift[]): number {
  let hours = 0
  shifts.forEach((shift) => {
    if (shift.status === 'scheduled') return
    
    if (['completed', 'all_good', 'late'].includes(shift.status)) {
      const start = parseInt(shift.startTime.split(':')[0])
      const end = parseInt(shift.endTime.split(':')[0])
      let duration = end - start
      if (duration < 0) duration += 24
      hours += duration
    }
  })
  return hours
}

export function calculateStreak(shifts: IShift[]): number {
  
  const sorted = [...shifts]
    .filter((s) => s.status !== 'scheduled')
    .sort((a, b) => b.date.localeCompare(a.date))

  let streak = 0
  for (const shift of sorted) {
    if (['completed', 'all_good', 'late'].includes(shift.status)) {
      streak++
    } else {
      
      break
    }
  }
  return streak
}

export function calculateMaxStreak(shifts: IShift[]): number {
  
  const sorted = [...shifts]
    .filter((s) => s.status !== 'scheduled')
    .sort((a, b) => a.date.localeCompare(b.date))

  let maxStreak = 0
  let currentStreak = 0

  for (const shift of sorted) {
    if (['completed', 'all_good', 'late'].includes(shift.status)) {
      currentStreak++
      if (currentStreak > maxStreak) {
        maxStreak = currentStreak
      }
    } else {
      currentStreak = 0
    }
  }
  return maxStreak
}
