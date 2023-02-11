import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig({
 server: {
    port: 3000,
    proxy: {
        "/server": {
          target: "https://probe.muhesh.studio",
          changeOrigin: true,
          secure: false,
          rewrite: (path) => path.replace(/^\/api/, ""),
        },
      },
},
  plugins: [react()],
})
