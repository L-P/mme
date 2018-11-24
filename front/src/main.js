import Vue from 'vue';
import './plugins/axios';
import App from './App.vue';
import router from './router';

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App),
}).$mount('#app');

Vue.filter('hex', (v, width) => {
  const hex = v.toString(16).toUpperCase().padStart(width, '0');
  return `0x${hex}`;
});
