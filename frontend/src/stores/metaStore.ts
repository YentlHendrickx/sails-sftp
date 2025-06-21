// Pinia store
import { defineStore } from "pinia";

import { Meta } from "@interfaces/meta";

export const useMetaStore = defineStore("meta", {
  state: (): Meta => {
    // TODO: Manage this into a 'settings' or 'preferences' module
    const darkMode = localStorage.getItem("darkMode") === "true";

    return {
      darkMode: darkMode,
      showSidebar: true,
      showHeader: true,
    };
  },
  getters: {
    isDarkMode: (state) => state.darkMode,
  },
  actions: {
    setDarkMode(value: boolean) {
      this.darkMode = value;
      console.log("Dark mode set to:", value);
      localStorage.setItem("darkMode", String(value)); // Sync to localStorage
    },
    toggleDarkMode() {
      this.setDarkMode(!this.darkMode);
    },
  },
});
