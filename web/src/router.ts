// router.ts
import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';

const Home = () => import('./pages/Home.vue');
const CourseDetail = () => import('./pages/CourseDetail.vue');
const Enroll = () => import('./pages/Enroll.vue');
const FAQ = () => import('./pages/FAQ.vue');
const Login = () => import('./pages/Login.vue');
const Profile = () => import('./pages/Profile.vue');
const Register = () => import('./pages/Register.vue');
const PasswordReset = () => import('./pages/PasswordReset.vue');
const PasswordResetDone = () => import('./pages/PasswordResetDone.vue');
const PasswordResetConfirm = () => import('./pages/PasswordResetConfirm.vue');
const PasswordResetComplete = () => import('./pages/PasswordResetComplete.vue');
const QuizTake = () => import('./pages/QuizTake.vue');

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    props: true,
  },
  {
    path: '/course/:slug',
    name: 'CourseDetail',
    component: CourseDetail,
    props: (route) => ({ slug: route.params.slug }),
  },
  {
    path: '/enroll/:slug',
    name: 'Enroll',
    component: Enroll,
    props: (route) => ({ slug: route.params.slug }),
  },
  { path: '/profile', name: 'Profile', component: Profile },
  { path: '/login', name: 'Login', component: Login },
  { path: '/register', name: 'Register', component: Register },
  { path: '/faq', name: 'FAQ', component: FAQ },
  {
    path: '/quiz/:id',
    name: 'QuizTake',
    component: QuizTake,
    props: (route) => ({ id: route.params.id }),
  },
  {
    path: '/password_reset',
    name: 'PasswordReset',
    component: PasswordReset,
  },
  {
    path: '/password_reset/done',
    name: 'PasswordResetDone',
    component: PasswordResetDone,
  },
  {
    path: '/password_reset/:uid/:token',
    name: 'PasswordResetConfirm',
    component: PasswordResetConfirm,
    props: true,
  },
  {
    path: '/password_reset/complete',
    name: 'PasswordResetComplete',
    component: PasswordResetComplete,
  },
  {
    path: '/:catchAll(.*)',
    redirect: '/',
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Optional: Global navigation guards to fetch data or protect routes
router.beforeEach(async (to, from, next) => {
  // Example: fetch courses if not loaded yet
  // You can implement global store (Pinia) to handle courses/enrollments state
  next();
});

export default router;
