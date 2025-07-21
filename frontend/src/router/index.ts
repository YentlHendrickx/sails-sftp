// Router index, file for the main router of the application
import { Router } from "vue-router";
import { createRouter, createWebHistory } from "vue-router";

import { routes } from "./routes";

export const router: Router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});
