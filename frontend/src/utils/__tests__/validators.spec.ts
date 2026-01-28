import { describe, expect, it } from 'vitest'

import { formatPhoneNumber, sanitizeAddress, sanitizeCity, sanitizeName, sanitizeZip } from '../validators.js'

describe('validators', () => {
  describe('formatPhoneNumber', () => {
    it('returns empty string for empty input', () => {
      expect(formatPhoneNumber(null)).toBe('')
      expect(formatPhoneNumber('')).toBe('')
    })

    it('formats partial numbers', () => {
      expect(formatPhoneNumber('123')).toBe('(123')
      expect(formatPhoneNumber('123456')).toBe('(123)456')
    })

    it('formats full 10-digit numbers', () => {
      expect(formatPhoneNumber('1234567890')).toBe('(123)456-7890')
    })

    it('ignores non-digits', () => {
        expect(formatPhoneNumber('123-abc-456')).toBe('(123)456')
    })
  })

  describe('sanitizeName', () => {
    it('removes special characters but keeps spaces', () => {
      expect(sanitizeName('John Doe!')).toBe('John Doe')
      expect(sanitizeName('O\'Connor')).toBe('OConnor') // Current impl removes quote
      expect(sanitizeName('Jane-Marie')).toBe('JaneMarie') // Current impl removes dash
    })

    it('handles empty input', () => {
      expect(sanitizeName(null)).toBe('')
    })
  })

  describe('sanitizeCity', () => {
    it('allows letters, numbers, spaces, and dashes', () => {
      expect(sanitizeCity('New York')).toBe('New York')
      expect(sanitizeCity('Winston-Salem')).toBe('Winston-Salem')
    })

    it('removes other special chars', () => {
      expect(sanitizeCity('City!')).toBe('City')
      expect(sanitizeCity('St. Louis')).toBe('St Louis') // Current impl removes dot
    })
  })

  describe('sanitizeZip', () => {
    it('limits to 5 digits', () => {
      expect(sanitizeZip('123456')).toBe('12345')
    })

    it('removes non-digits', () => {
      expect(sanitizeZip('123ab')).toBe('123')
    })
  })

  describe('sanitizeAddress', () => {
    it('allows alphanumeric, spaces, and dashes', () => {
      expect(sanitizeAddress('123 Main St-West')).toBe('123 Main St-West')
    })

    it('removes special characters', () => {
      expect(sanitizeAddress('Apt #4')).toBe('Apt 4')
    })
  })
})
