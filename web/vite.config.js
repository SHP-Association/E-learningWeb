// Vite config with dev proxy for backend API
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  const backendURL = env.VITE_APP_BACKEND_URL || 'http://localhost:8000'

  return {
    plugins: [vue()],
    server: {
      proxy: {
        '/api': {
          target: backendURL,
          changeOrigin: true,
          secure: false,
        },
      },
    },
    build: {
      rollupOptions: {
        output: {
          manualChunks: undefined,
        },
      },
      chunkSizeWarningLimit: 1000,
    },
  }
})
