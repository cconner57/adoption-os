import { onMounted, onUnmounted } from 'vue'

export function usePullToRefresh(onRefresh: () => Promise<void> | void) {
  let touchStartY = 0

  const handleTouchStart = (e: TouchEvent) => {
    // Only track if at the top of the page
    if (globalThis.scrollY === 0) {
      touchStartY = e.touches[0].clientY
    }
  }

  const handleTouchEnd = async (e: TouchEvent) => {
    if (globalThis.scrollY === 0) {
      const touchEndY = e.changedTouches[0].clientY
      // If pulled down more than 150px
      if (touchEndY - touchStartY > 150) {
        // Trigger refresh
        await onRefresh()
        // Optional: Vibrate if haptics available
        if (typeof navigator !== 'undefined' && navigator.vibrate) {
            navigator.vibrate(50)
        }
        // Force reload to ensure fresh data/state if simple refresh isn't enough,
        // or just rely on onRefresh() callback
        globalThis.location.reload()
      }
    }
    touchStartY = 0
  }

  onMounted(() => {
    // 1. Check if running in standalone mode (PWA)
    const isStandalone = globalThis.matchMedia('(display-mode: standalone)').matches

    // 2. If valid, add listeners
    if (isStandalone) {
      document.addEventListener('touchstart', handleTouchStart, { passive: true })
      document.addEventListener('touchend', handleTouchEnd, { passive: true })
    }
  })

  onUnmounted(() => {
     document.removeEventListener('touchstart', handleTouchStart)
     document.removeEventListener('touchend', handleTouchEnd)
  })
}
