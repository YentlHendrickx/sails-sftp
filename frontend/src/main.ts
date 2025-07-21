import { createApp } from "vue";
import App from "./App.vue";
import "./style.css";
import Vue3Toastify from "vue3-toastify";
import "vue3-toastify/dist/index.css";
import { createPinia } from "pinia";

import { router } from "@/router/index";

const app = createApp(App);
app.use(Vue3Toastify);
app.use(createPinia());
app.use(router);
app.mount("#app");
