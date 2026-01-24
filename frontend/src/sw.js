import { clientsClaim } from 'workbox-core'
import { cleanupOutdatedCaches, precacheAndRoute } from 'workbox-precaching'

cleanupOutdatedCaches()
precacheAndRoute(self.__WB_MANIFEST)

self.addEventListener('message', (event) => {
  if (event.data && event.data.type === 'SKIP_WAITING') {
    self.skipWaiting()
  }
})

clientsClaim()

// Push Notification Listener
self.addEventListener('push', (event) => {
  if (!event.data) return

  const data = event.data.json()
  const title = data.title || 'AdoptionOS'
  const options = {
    body: data.body || '',
    icon: data.icon || '/images/paw.svg',
    badge: '/images/paw.svg', // Monochrome icon for status bar
    vibrate: [100, 50, 100],
    data: {
      url: data.url || '/'
    }
  }

  event.waitUntil(self.registration.showNotification(title, options))
})

// Notification Click Listener
self.addEventListener('notificationclick', (event) => {
  event.notification.close()
  event.waitUntil(
    clients.matchAll({ type: 'window', includeUncontrolled: true }).then((clientList) => {
      if (clientList.length > 0) {
        let client = clientList[0]
        for (let i = 0; i < clientList.length; i++) {
          if (clientList[i].focused) {
            client = clientList[i]
          }
        }
        return client.focus()
      }
      return clients.openWindow(event.notification.data.url)
    })
  )
})
