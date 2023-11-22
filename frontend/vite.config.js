// vite.config.js
import { defineConfig } from 'vite';

export default defineConfig({
  root: './src',
  server: {
    port: 3002,
    open: true,
    proxy: {
      '/dynamic-content': 'http://localhost:1323', // Replace with your Echo server's URL and port
    }
  }
});
