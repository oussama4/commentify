import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import postcssConfig from './postcss.config.js'

// https://vitejs.dev/config/
export default defineConfig({
  envDir: ".",
  css: {
    postcss: postcssConfig
  },
  plugins: [svelte()]
})
