import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";
import "./style.css";
import App from "./App.vue";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css"; // Import Element Plus CSS

import Home from "./components/Home.vue";
import AboutUs from "./components/AboutUs.vue";
import Services from "./components/Services.vue";
import DataBackup from "./components/DataBackup.vue";
import DataRestore from "./components/DataRestore.vue";
import DataMigration from "./components/DataMigration.vue";
import MongoDB from "./components/MongoDB.vue";
import MySQL from "./components/MySQL.vue";
import PostgreSQL from "./components/PostgreSQL.vue";
import BreadCrumb from "./components/BreadCrumb.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: Home },
    { path: "/home", component: Home },
    { path: "/aboutus", component: AboutUs },
    { path: "/services", component: Services },
    { path: "/services/data-backup", component: DataBackup },
    { path: "/services/data-restore", component: DataRestore },
    { path: "/services/data-migration", component: DataMigration },
    { path: "/MongoDB", component: MongoDB },
    { path: "/MySQL", component: MySQL },
    { path: "/services/PostgreSQL", component: PostgreSQL },
  ],
});

const app = createApp(App);
app.component("BreadCrumb", BreadCrumb);
app.use(ElementPlus);
app.use(router);
app.mount("#app");
