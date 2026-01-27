/**
 * Safely triggers device vibration if supported by the browser.
 * @param ms Duration of vibration in milliseconds
 */
export const vibrate = (ms: number) => {
  if (typeof navigator !== 'undefined' && navigator.vibrate) {
    try {
      navigator.vibrate(ms)
    } catch {
      // Ignore errors if vibration fails or is blocked
    }
  }
}
