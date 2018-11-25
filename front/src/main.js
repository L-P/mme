import Vue from 'vue';
import './plugins/axios';
import App from './App.vue';
import router from './router';
import config from '@/config';

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App),
}).$mount('#app');

Vue.filter('hex', (v, width) => {
  const hex = v.toString(16).toUpperCase().padStart(width, '0');
  return `0x${hex}`;
});

Vue.filter('apiURI', v => `${config.API_URI}${v}`);

Vue.filter('humanizeBytes', (v) => {
  if (v === 0) {
    return '0 B';
  }

  const units = ['B', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'ZiB', 'YiB'];
  const i = Math.min(units.length - 1, Math.floor(Math.log(v) / Math.log(1024)));
  let rounded = v / (1024 ** i);
  rounded = +rounded.toFixed(2);


  return `${rounded} ${units[i]}`;
});
