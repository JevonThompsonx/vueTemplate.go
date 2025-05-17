import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      // String shorthand: '/api' -> 'http://localhost:8080/api'
      "/api": {
        target: "http://localhost:8080", // Your Go backend address
        changeOrigin: true,
        // Optional: rewrite path if your backend expects a different path structure
        // rewrite: (path) => path.replace(/^\/api/, '')
      },
    },
  },
});
