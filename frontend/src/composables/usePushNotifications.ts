import { ref } from 'vue'

function urlBase64ToUint8Array(base64String: string) {
  const padding = '='.repeat((4 - (base64String.length % 4)) % 4)
  const base64 = (base64String + padding).replace(/-/g, '+').replace(/_/g, '/')
  const rawData = window.atob(base64)
  const outputArray = new Uint8Array(rawData.length)
  for (let i = 0; i < rawData.length; ++i) {
    outputArray[i] = rawData.charCodeAt(i)
  }
  return outputArray
}

export function usePushNotifications() {
  const isSupported = 'serviceWorker' in navigator && 'PushManager' in window
  const isSubscribed = ref(false)
  const error = ref<string | null>(null)
  const loading = ref(false)

  // Check initial status
  async function checkSubscription() {
    if (!isSupported) return
    const registration = await navigator.serviceWorker.ready
    const subscription = await registration.pushManager.getSubscription()
    isSubscribed.value = !!subscription
  }

  async function subscribe() {
    if (!isSupported) {
        error.value = 'Push notifications not supported in this browser'
        return
    }
    loading.value = true
    error.value = null
    console.log('AdoptionOS: Starting subscription...')

    try {
      // Fetch public key from backend
      console.log('AdoptionOS: Fetching public key...')
      const keyRes = await fetch(`/v1/notifications/public-key?t=${Date.now()}`)
      if (!keyRes.ok) {
        const text = await keyRes.text()
        throw new Error(`Failed to fetch VAPID key: ${keyRes.status} ${text}`)
      }
      const { public_key } = await keyRes.json()
      console.log('AdoptionOS: Public key received', public_key ? 'YES' : 'NO')

      const registration = await navigator.serviceWorker.ready
      console.log('AdoptionOS: Service Worker ready')

      const subscription = await registration.pushManager.subscribe({
        userVisibleOnly: true,
        applicationServerKey: urlBase64ToUint8Array(public_key),
      })
      console.log('AdoptionOS: Browser subscription created')

      // Send to backend
      const res = await fetch('/v1/notifications/subscribe', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(subscription),
      })

      if (!res.ok) {
        throw new Error('Failed to save subscription on server')
      }

      isSubscribed.value = true
    } catch (e: unknown) {
      console.error(e)
      if (e instanceof Error) {
        error.value = e.message
      } else {
        error.value = 'Failed to subscribe'
      }
      isSubscribed.value = false
    } finally {
      loading.value = false
    }
  }

  async function unsubscribe() {
    if (!isSupported) return
    loading.value = true
    try {
        const registration = await navigator.serviceWorker.ready
        const subscription = await registration.pushManager.getSubscription()
        if (subscription) {
            await subscription.unsubscribe()
            // Optional: Tell backend to delete
        }
        isSubscribed.value = false
    } catch(e: unknown) {
        console.error(e)
    } finally {
        loading.value = false
    }
  }

  async function sendTestNotification() {
    if (!isSubscribed.value) return
    loading.value = true
    try {
        const registration = await navigator.serviceWorker.ready
        const subscription = await registration.pushManager.getSubscription()
        if (subscription) {
            await fetch('/v1/notifications/test', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(subscription)
            })
        }
    } catch(e: unknown) {
        console.error(e)
        error.value = 'Failed to send test'
    } finally {
        loading.value = false
    }
  }

  return {
    isSupported,
    isSubscribed,
    subscribe,
    unsubscribe,
    sendTestNotification,
    checkSubscription,
    loading,
    error,
  }
}
