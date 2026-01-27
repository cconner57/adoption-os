import { describe, expect, it } from 'vitest'

import { formatDate, formatDateTime } from '../dateUtils.js'

describe('dateUtils', () => {
  describe('formatDate', () => {
    it('returns "-" for null, undefined, or empty string', () => {
      expect(formatDate(null)).toBe('-')
      expect(formatDate(undefined)).toBe('-')
      expect(formatDate('')).toBe('-')
    })

    it('returns "-" for invalid dates', () => {
      expect(formatDate('invalid-date')).toBe('-')
    })

    it('formats YYYY-MM-DD strings as local dates (preventing timezone rollback)', () => {
      // 2023-01-01 should match Jan 1, 2023 regardless of timezone
      // We check for "Jan 1, 2023" or similar depending on locale, but en-US is hardcoded in implementation
      expect(formatDate('2023-01-01')).toBe('Jan 1, 2023')
      expect(formatDate('2023-12-31')).toBe('Dec 31, 2023')
    })

    it('formats ISO ISO strings correctly', () => {
      // These use new Date() which parses in local timezone or UTC depending on string
      // "2023-01-01T12:00:00" is local
      expect(formatDate('2023-01-01T12:00:00')).toBe('Jan 1, 2023')
    })
  })

  describe('formatDateTime', () => {
    it('returns "-" for null or undefined', () => {
      expect(formatDateTime(null)).toBe('-')
      expect(formatDateTime(undefined)).toBe('-')
    })

    it('returns "N/A" for 0001-01-01', () => {
      expect(formatDateTime('0001-01-01T00:00:00Z')).toBe('N/A')
      expect(formatDateTime('0001-01-01')).toBe('N/A')
    })

    it('falls back to formatDate for short strings (YYYY-MM-DD)', () => {
      expect(formatDateTime('2023-05-20')).toBe('May 20, 2023')
    })

    it('includes time for full ISO strings', () => {
      // This test might be flaky across timezones if we check exact string
      // e.g. "Jan 1, 2023, 12:00 PM"
      const res = formatDateTime('2023-01-01T12:00:00')
      expect(res).toContain('Jan 1, 2023')
      // Check time format loosely to avoid strict timezone flakes if machine settings weird
      // But implementation uses 'en-US'
      expect(res).toContain('12:00 PM')
    })

    it('returns "-" for invalid date strings', () => {
      expect(formatDateTime('not-a-date')).toBe('-')
    })
  })
})
