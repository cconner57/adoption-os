/**
 * Formats a date string into a readable format (e.g., "Aug 11, 1990").
 * Handles both ISO strings and YYYY-MM-DD date strings appropriately to avoid timezone shifts.
 *
 * @param dateStr - The date string to format (ISO or YYYY-MM-DD)
 * @returns Formatted date string or '-' if input is invalid/empty
 */
export function formatDate(dateStr?: string | null): string {
  if (!dateStr) return '-'

  let date: Date

  // If YYYY-MM-DD (length 10), parse as local parts to avoid UTC shift
  if (dateStr.length === 10 && dateStr.includes('-')) {
    const [y, m, d] = dateStr.split('-').map(Number)
    date = new Date(y, m - 1, d)
  } else {
    date = new Date(dateStr)
  }

  // Check for invalid date
  if (isNaN(date.getTime())) return '-'

  return date.toLocaleDateString('en-US', {
    weekday: undefined, // Removed weekday for generic usage unless explicitly requested
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

/**
 * Calculates age from a date of birth string.
 * - If >= 1 year: "X yr Y mo"
 * - If < 1 year: "X mo Y days"
 */
export function calculateAge(dateOfBirth?: string | null): string {
  if (!dateOfBirth) return '-'

  const birthDate = new Date(dateOfBirth)
  const today = new Date()

  if (isNaN(birthDate.getTime())) return '-'

  let years = today.getFullYear() - birthDate.getFullYear()
  let months = today.getMonth() - birthDate.getMonth()
  let days = today.getDate() - birthDate.getDate()

  // Adjust for negative days (borrow from previous month)
  if (days < 0) {
    months--
    // Get days in the previous month relative to today
    // new Date(year, monthIndex, 0) gives the last day of the *previous* month
    const prevMonthLastDay = new Date(today.getFullYear(), today.getMonth(), 0).getDate()
    days += prevMonthLastDay
  }

  // Adjust for negative months (borrow from year)
  if (months < 0) {
    years--
    months += 12
  }

  // Safety check for future dates or same day
  if (years < 0) return '0 days'

  // Logic:
  // If >= 1 year -> Show Years + Months
  if (years >= 1) {
    if (months === 0) return `${years} yr`
    return `${years} yr ${months} mo`
  }

  // If < 1 year -> Show Months + Days
  // Edge cases: 0 mo, 0 days, etc.
  const parts: string[] = []
  if (months > 0) parts.push(`${months} mo`)
  if (days > 0) parts.push(`${days} day${days === 1 ? '' : 's'}`)

  if (parts.length === 0) return '0 days' // New born today
  return parts.join(' ')
}
