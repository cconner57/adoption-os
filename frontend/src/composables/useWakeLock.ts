import { onUnmounted, ref } from 'vue'

export function useWakeLock() {
  const isSupported = typeof navigator !== 'undefined' && 'wakeLock' in navigator
  const isActive = ref(false)

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  let wakeLock: any | null = null

  const requestLock = async () => {
    if (!isSupported) {
      console.warn('Wake Lock API not supported.')
      return
    }

    try {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      wakeLock = await (navigator as any).wakeLock.request('screen')
      isActive.value = true

      wakeLock.addEventListener('release', () => {
        isActive.value = false
        wakeLock = null
        console.log('Wake Lock released (event).')
      })
      console.log('Wake Lock active.')
    } catch (err: unknown) {
      console.error('Wake Lock request failed:', err)
      isActive.value = false
    }
  }

  const releaseLock = async () => {
    if (wakeLock) {
      try {
        await wakeLock.release()
      } catch (err: unknown) {
        console.error('Wake Lock release failed:', err)
      } finally {
        wakeLock = null
        isActive.value = false
      }
    }
  }

  // Automatically release properly when component unmounts
  onUnmounted(() => {
    if (isActive.value) {
      releaseLock()
    }
  })

  return {
    isSupported,
    isActive,
    requestLock,
    releaseLock,
  }
}
