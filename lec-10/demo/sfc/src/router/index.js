import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../pages/Home.vue';
import Cart from '../pages/Cart.vue';
import Product from '../pages/Product.vue';

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
  }
]

const router = new VueRouter({
  mode: 'history',
  routes,
});

export default router;
