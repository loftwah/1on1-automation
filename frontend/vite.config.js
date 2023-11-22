import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
  root: '.', // Current directory is 'frontend'
  base: '/',
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'), // Alias '@' points to 'src' within 'frontend'
    },
  },
  plugins: [],
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@import "src/scss/_base.scss";`, // Path relative to 'frontend'
      },
    },
  },
  build: {
    outDir: 'dist', // Output directory relative to 'frontend'
    emptyOutDir: true,
    rollupOptions: {
      input: path.resolve(__dirname, 'src/index.html'), // Main entry point inside 'src'
    },
  },
  server: {
    port: 3002,
    open: true,
  },
});
