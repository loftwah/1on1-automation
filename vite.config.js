import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
  root: path.resolve(__dirname, './frontend'),
  base: '/',
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './frontend'),
    },
  },
  plugins: [],
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@import "./frontend/scss/_base.scss";`,
      },
    },
  },
  build: {
    outDir: path.resolve(__dirname, './dist'),
    emptyOutDir: true,
    rollupOptions: {
      input: {
        main: path.resolve(__dirname, 'frontend/index.html'),
        // include other HTML files if they are entry points
      },
    },
  },
  server: {
    port: 3002,
    open: true,
  },
});
