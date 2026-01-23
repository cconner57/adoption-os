
export function formatPhoneNumber(phone?: string | null): string {
  if (!phone) return '-'

  const cleaned = (`${  phone}`).replace(/\D/g, '')

  if (cleaned.length === 10) {
    return `${cleaned.slice(0, 3)}-${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }

  if (cleaned.length === 11 && cleaned.startsWith('1')) {
    const main = cleaned.slice(1)
    return `${main.slice(0, 3)}-${main.slice(3, 6)}-${main.slice(6)}`
  }

  if (cleaned.length > 6) {
    return `${cleaned.slice(0, 3)}-${cleaned.slice(3, 6)}-${cleaned.slice(6)}`
  }

  return phone
}

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

export function formatZipInput(value: string): string {
  return value.replace(/\D/g, '').slice(0, 5)
}
