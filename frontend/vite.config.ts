import { fileURLToPath, URL } from 'node:url'
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import dns from 'node:dns'

dns.setDefaultResultOrder('ipv4first')

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')

  return {
    plugins: [vue()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    server: {
      proxy: {
        '/v1': {
          target: env.VITE_API_URL || 'http://192.168.12.102:8080',
          changeOrigin: true,
          secure: false,
          rewrite: (path) => path.replace(/^\/v1/, ''),
        },
      },
    },
  }
})
