/**
 * Formats a phone number string into XXX-XXX-XXXX format.
 * Strips non-numeric characters first.
 * @param phone The input phone string
 * @returns Formatted phone number
 */
export function formatPhoneNumber(phone?: string | null): string {
  if (!phone) return '-'

  // Strip all non-numeric characters
  const cleaned = ('' + phone).replace(/\D/g, '')

  // Check if it looks like a standard US number (10 digits)
  // If messy or not 10 digits, just return original or cleaned?
  // User asked for specific format 626-616-7068.

  if (cleaned.length === 10) {
    return `${cleaned.slice(0, 3)}-${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }

  // If 11 digits and starts with 1, strip it?
  if (cleaned.length === 11 && cleaned.startsWith('1')) {
    const main = cleaned.slice(1)
    return `${main.slice(0, 3)}-${main.slice(3, 6)}-${main.slice(6)}`
  }

  // Fallback: if we have some chunks, try to format, or return original
  if (cleaned.length > 6) {
    return `${cleaned.slice(0, 3)}-${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }

  return phone
}

/**
 * Formator for input fields (as-you-type).
 * Returns the cleaned and formatted string.
 */
export function formatPhoneInput(value: string): string {
  const cleaned = value.replace(/\D/g, '')

  if (cleaned.length >= 10) {
    return `${cleaned.slice(0, 3)}-${cleaned.slice(3, 6)}-${cleaned.slice(6, 10)}`
  }
  if (cleaned.length > 6) {
    return `${cleaned.slice(0, 3)}-${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }
  if (cleaned.length > 3) {
    return `${cleaned.slice(0, 3)}-${cleaned.slice(3)}`
  }

  return cleaned
}

/**
 * Formats zip input to only allow numbers and max 5 chars.
 */
export function formatZipInput(value: string): string {
  return value.replace(/\D/g, '').slice(0, 5)
}
