import { fileURLToPath } from 'node:url'

import { configDefaults,defineConfig, mergeConfig } from 'vitest/config'

import viteConfig from './vite.config'

export default defineConfig(async (env) => {

  const viteConfigObj = (typeof viteConfig === 'function' ? viteConfig(env) : viteConfig)

  return mergeConfig(
    viteConfigObj,
    defineConfig({
      test: {
        environment: 'jsdom',
        exclude: [...configDefaults.exclude, 'e2e/**'],
        root: fileURLToPath(new URL('./', import.meta.url)),
        coverage: {
          provider: 'v8',
          reporter: ['text', 'json', 'html'],
          exclude: [
            ...configDefaults.exclude,
            'e2e/**',
            '*.config.ts',
            'src/manifest.ts',
            'src/main.ts',
            'src/vite-env.d.ts',
            'src/stores/mockPetData.ts',
          ],
        },
      },
    }),
  )
})
