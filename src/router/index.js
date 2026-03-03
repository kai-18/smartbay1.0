import { createRouter, createWebHistory } from "vue-router";
import DashboardPage from "src/pages/DashboardPage.vue";
import WaitingPage from "src/pages/WaitingPage.vue";
import MainLayout from "layouts/MainLayout.vue";

const routes = [
  {
    path: "/",
    component: MainLayout,
    children: [{ path: "", component: DashboardPage }],
    meta: { requiresAuth: true },
  },
  {
    path: "/waiting",
    component: MainLayout,
    children: [{ path: "", component: WaitingPage }],
    meta: { requiresAuth: true },
  }
];
  const router = createRouter({
    history: createWebHistory(),
    routes,
  });

export default router;
