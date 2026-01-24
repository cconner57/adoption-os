/// <reference types="node" />
import dns from 'node:dns'
import { fileURLToPath, URL } from 'node:url'

import vue from '@vitejs/plugin-vue'
import { defineConfig, loadEnv } from 'vite'
import { VitePWA } from 'vite-plugin-pwa'

dns.setDefaultResultOrder('ipv4first')

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')

  return {
    plugins: [vue(), VitePWA({
      registerType: 'autoUpdate',
      includeAssets: ['favicon.ico', 'images/paw.svg'],
      manifest: {
        name: 'Adoption OS',
        short_name: 'AdoptionOS',
        description: 'Adoption Management System',
        theme_color: '#00a5ad',
        background_color: '#ffffff',
        icons: [
          {
            src: '/images/paw.svg',
            sizes: '192x192',
            type: 'image/svg+xml'
          },
          {
            src: '/images/paw.svg',
            sizes: '512x512',
            type: 'image/svg+xml'
          }
        ]
      },
      workbox: {
        runtimeCaching: [
          {
            urlPattern: ({ url }) => url.pathname.startsWith('/v1'),
            method: 'GET',
            handler: 'NetworkFirst',
            options: {
              cacheName: 'api-cache-get',
              expiration: {
                maxEntries: 100,
                maxAgeSeconds: 72 * 60 * 60 // 3 days
              }
            }
          },
          {
            urlPattern: ({ url }) => url.pathname.startsWith('/v1'),
            method: 'POST',
            handler: 'NetworkOnly',
            options: {
              backgroundSync: {
                name: 'offline-mutations-post',
                options: {
                  maxRetentionTime: 24 * 60 // 24 hours
                }
              }
            }
          },
          {
            urlPattern: ({ url }) => url.pathname.startsWith('/v1'),
            method: 'PUT',
            handler: 'NetworkOnly',
            options: {
              backgroundSync: {
                name: 'offline-mutations-put',
                options: {
                  maxRetentionTime: 24 * 60
                }
              }
            }
          },
          {
            urlPattern: ({ url }) => url.pathname.startsWith('/v1'),
            method: 'DELETE',
            handler: 'NetworkOnly',
            options: {
              backgroundSync: {
                name: 'offline-mutations-delete',
                options: {
                  maxRetentionTime: 24 * 60
                }
              }
            }
          }
        ]
      }
    })],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    server: {
      allowedHosts: ['www.idohr.app', 'idohr.app'],
      proxy: {
        // Proxy all backend routes to local backend while maintaining relative paths
        // This allows access via localhost OR tunnel (idohr.app) transparently
        '/v1': {
          target: env.VITE_API_URL || 'http://127.0.0.1:8080',
          changeOrigin: true,
          secure: false,
        },
        '/api': {
          target: env.VITE_API_URL || 'http://127.0.0.1:8080',
          changeOrigin: true,
          secure: false,
        },
        '/pets': {
          target: env.VITE_API_URL || 'http://127.0.0.1:8080',
          changeOrigin: true,
          secure: false,
        },
        '/applications': {
          target: env.VITE_API_URL || 'http://127.0.0.1:8080',
          changeOrigin: true,
          secure: false,
        },
        '/metrics': {
          target: env.VITE_API_URL || 'http://127.0.0.1:8080',
          changeOrigin: true,
          secure: false,
        },
        '/healthz': {
          target: env.VITE_API_URL || 'http://127.0.0.1:8080',
          changeOrigin: true,
          secure: false,
        },
        '/sitemap.xml': {
          target: env.VITE_API_URL || 'http://127.0.0.1:8080',
          changeOrigin: true,
          secure: false,
        },
        '/uploads': {
          target: env.VITE_API_URL || 'http://127.0.0.1:8080',
          changeOrigin: true,
          secure: false,
        },
        '/pet-photos': {
          target: env.VITE_API_URL || 'http://127.0.0.1:8080',
          changeOrigin: true,
          secure: false,
        },
      },
    },
  }
})
