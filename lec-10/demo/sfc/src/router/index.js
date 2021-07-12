import Vue from 'vue';
import VueRouter from 'vue-router';

const Home = () => import('../pages/Home.vue');
const Cart = () => import('../pages/Cart.vue');
const Product = () => import('../pages/Product.vue');

Vue.use(VueRouter);

const routes = [
  {
    name: 'home',
    path: '/',
    component: Home,
  },
  {
    name: 'cart',
    path: '/cart',
    component: Cart,
  },
  {
    name: 'product',
    path: '/product/:handle',
    component: Product,
  },
  {
    name: 'sign-in',
    path: '/account/sign-in',
    component: Home,
    meta: {
      layout: 'account'
    }
  },
  {
    name: 'sign-up',
    path: '/account/sign-up',
    component: Home,
    meta: {
      layout: 'account'
    }
  },
]

const router = new VueRouter({
  mode: 'history',
  routes,
});

export default router;
