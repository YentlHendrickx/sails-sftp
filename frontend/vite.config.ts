import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
// @ts-ignore
import tailwindcss from "@tailwindcss/vite";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), tailwindcss()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      types: path.resolve(__dirname, "./src/types"),
      "@interfaces": path.resolve(__dirname, "./src/interfaces"),
      "@stores": path.resolve(__dirname, "./src/stores"),
    },
  },
});
