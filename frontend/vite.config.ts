import { fileURLToPath, URL } from 'node:url'
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

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
        // Any request starting with /v1 will be sent to the Backend
        '/v1': {
          target: env.VITE_API_URL || 'http://localhost:8080', // Fallback to localhost if missing
          changeOrigin: true,
          secure: false,
          rewrite: (path) => path.replace(/^\/v1/, ''),
        },
      },
    },
  }
})
