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
  server: {
    // everything the go server provides, including the env driven theme
    proxy: Object.fromEntries(
      ["/api", "/theme.css", "/logo", "/favicon.ico"].map((path) => [
        path,
        { target: "http://localhost:8088", changeOrigin: true },
      ])
    ),
  },
}
)
