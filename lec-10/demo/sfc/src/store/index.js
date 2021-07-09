import Vue from 'vue';
import Vuex from 'vuex';
import product from './product';
import cart from './cart';

Vue.use(Vuex);

const store = new Vuex.Store({
  strict: true,
  modules: {
    product,
    cart,
  }
});

export default store;
