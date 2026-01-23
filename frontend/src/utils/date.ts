
export function formatDate(dateStr?: string | null): string {
  if (!dateStr) return '-'

  let date: Date

  if (dateStr.length === 10 && dateStr.includes('-')) {
    const [y, m, d] = dateStr.split('-').map(Number)
    date = new Date(y, m - 1, d)
  } else {
    date = new Date(dateStr)
  }

  if (isNaN(date.getTime())) return '-'

  return date.toLocaleDateString('en-US', {
    weekday: undefined, 
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

export function calculateAge(dateOfBirth?: string | null): string {
  if (!dateOfBirth) return '-'

  const birthDate = new Date(dateOfBirth)
  const today = new Date()

  if (isNaN(birthDate.getTime())) return '-'

  let years = today.getFullYear() - birthDate.getFullYear()
  let months = today.getMonth() - birthDate.getMonth()
  let days = today.getDate() - birthDate.getDate()

  if (days < 0) {
    months--
    
    const prevMonthLastDay = new Date(today.getFullYear(), today.getMonth(), 0).getDate()
    days += prevMonthLastDay
  }

  if (months < 0) {
    years--
    months += 12
  }

  if (years < 0) return '0 days'

  if (years >= 1) {
    if (months === 0) return `${years} yr`
    return `${years} yr ${months} mo`
  }

  const parts: string[] = []
  if (months > 0) parts.push(`${months} mo`)
  if (days > 0) parts.push(`${days} day${days === 1 ? '' : 's'}`)

  if (parts.length === 0) return '0 days' 
  return parts.join(' ')
}
