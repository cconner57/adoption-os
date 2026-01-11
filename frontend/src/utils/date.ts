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
 * returns strings like "2 yr 3 mo", "5 mo", "< 1 mo".
 */
export function calculateAge(dateOfBirth?: string | null): string {
  if (!dateOfBirth) return '-'

  const birthDate = new Date(dateOfBirth)
  const today = new Date()

  if (isNaN(birthDate.getTime())) return '-'

  let years = today.getFullYear() - birthDate.getFullYear()
  let months = today.getMonth() - birthDate.getMonth()

  if (months < 0 || (months === 0 && today.getDate() < birthDate.getDate())) {
    years--
    months += 12
  }
  if (today.getDate() < birthDate.getDate()) months--

  if (years === 0 && months === 0) return '< 1 mo'
  if (years === 0) return `${months} mo`
  if (months === 0) return `${years} yr`
  return `${years} yr ${months} mo`
}
