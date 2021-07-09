import '../assets/styles/style.scss';
import Vue from 'vue';
import App from './App.vue';
import store from './store';
import router from './router';

const LoggerPlugin = {
  install(Vue, options = {}) {
    Vue.prototype.$log = function (msg) {
      if (!options.slient) {
        console.log('Log', msg);
      }
    }
  }
}

Vue.use(LoggerPlugin, { slient: true });

const vm = new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount('#app');

export default vm;
