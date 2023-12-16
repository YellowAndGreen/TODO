import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    meta: {
      requireAuth: true // 添加该字段，表示进入这个路由是需要登录的
    },
    component: Home,
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/sign/login.vue"),
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("../views/sign/register.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to) => {
  const token = localStorage.getItem("token");
  // 如果访问非登录界面，且户会话信息不存在，代表未登录，则跳转到登录界面
  if (to.meta.requireAuth && !token) {
    return { name: 'Login' }
  }
  return true;
});
export default router;
