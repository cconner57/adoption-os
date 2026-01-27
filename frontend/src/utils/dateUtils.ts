/**
 * Formats a date string to a human-readable format.
 * intended to fix timezone rollback issues by correctly parsing YYYY-MM-DD strings.
 */
export function formatDate(dateString?: string | null): string {
  if (!dateString) return '-'

  // Handle YYYY-MM-DD specifically to prevent timezone rollback
  // Treating it as a "Local Date" rather than UTC Midnight
  if (dateString.length === 10 && dateString.includes('-')) {
    const [year, month, day] = dateString.split('-').map(Number)
    const date = new Date(year, month - 1, day)

    return date.toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
    })
  }

  const date = new Date(dateString)
  if (Number.isNaN(date.getTime())) return '-'

  return date.toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

export function formatDateTime(dateString?: string | null): string {
    if (!dateString) return '-'
    if (dateString.startsWith('0001-01-01')) return 'N/A'

    const date = new Date(dateString)
    if (Number.isNaN(date.getTime())) return '-'

    // If it looks like a date-only string (length <= 10), use formatDate
    if (dateString.length <= 10) {
        return formatDate(dateString)
    }

    return date.toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
        year: 'numeric',
        hour: 'numeric',
        minute: '2-digit'
    })
}
