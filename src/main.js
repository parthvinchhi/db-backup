import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";
import "./style.css";
import App from "./App.vue";

import Home from "./components/Home.vue";
import AboutUs from "./components/AboutUs.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: Home },
    { path: "/home", component: Home },
    { path: "/aboutus", component: AboutUs },
  ],
});

const app = createApp(App);
app.use(router);
app.mount("#app");
