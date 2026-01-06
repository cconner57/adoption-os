export type FormInput = string | number | null

export const formatPhoneNumber = (value: FormInput): string => {
  if (!value) return ''
  const digits = String(value).replace(/\D/g, '').substring(0, 10)
  if (digits.length === 0) return ''
  if (digits.length <= 3) return `(${digits}`
  if (digits.length <= 6) return `(${digits.slice(0, 3)})${digits.slice(3)}`
  return `(${digits.slice(0, 3)})${digits.slice(3, 6)}-${digits.slice(6, 10)}`
}

export const sanitizeName = (value: FormInput): string => {
  if (!value) return ''
  return String(value).replace(/[^a-zA-Z0-9 ]/g, '')
}

export const sanitizeCity = (value: FormInput): string => {
  if (!value) return ''
  return String(value).replace(/[^a-zA-Z0-9 -]/g, '')
}

export const sanitizeZip = (value: FormInput): string => {
  if (!value) return ''
  return String(value).replace(/\D/g, '').substring(0, 5)
}

export const sanitizeAddress = (value: FormInput): string => {
  if (!value) return ''
  return String(value).replace(/[^a-zA-Z0-9 -]/g, '')
}
