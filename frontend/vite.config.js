import { fileURLToPath } from 'url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// src/views/Home.vue
// https://vitejs.dev/config/


export default defineConfig({
  plugins: [
    vue({
      reactivityTransform: true,
    })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  devServer: {
    proxy: {
      "^/api": {
        target: "http://localhost:8088",
        changeOrigin: true,
      },
    },
  },
}
)
