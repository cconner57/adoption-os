import { describe, expect, it } from 'vitest'

import { formatPhoneInput, formatPhoneNumber, formatZipInput } from '../formatters.js'

describe('formatters', () => {
  describe('formatPhoneNumber', () => {
    it('returns "-" for empty input', () => {
      expect(formatPhoneNumber()).toBe('-')
      expect(formatPhoneNumber('')).toBe('-')
    })

    it('formats 10-digit numbers', () => {
      expect(formatPhoneNumber('1234567890')).toBe('123-456-7890')
    })

    it('formats 11-digit numbers starting with 1', () => {
      expect(formatPhoneNumber('11234567890')).toBe('123-456-7890')
    })

    it('formats partial numbers > 6 digits', () => {
      expect(formatPhoneNumber('1234567')).toBe('123-456-7')
    })

    it('returns original if not matching patterns', () => {
      expect(formatPhoneNumber('123')).toBe('123')
    })
  })

  describe('formatPhoneInput', () => {
    it('formats as you type (3 digits)', () => {
      expect(formatPhoneInput('123')).toBe('123')
    })

    it('formats as you type (>3 digits)', () => {
      expect(formatPhoneInput('1234')).toBe('123-4')
    })

    it('formats as you type (>6 digits)', () => {
      expect(formatPhoneInput('1234567')).toBe('123-456-7')
    })

    it('limits to 10 digits', () => {
      expect(formatPhoneInput('1234567890123')).toBe('123-456-7890')
    })
  })

  describe('formatZipInput', () => {
    it('truncates to 5 digits', () => {
      expect(formatZipInput('123456')).toBe('12345')
    })

    it('removes non-digits', () => {
      expect(formatZipInput('123ab')).toBe('123')
    })
  })
})
