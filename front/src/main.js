import Vue from 'vue';
import Buefy from 'buefy';
import 'buefy/dist/buefy.css';

import './plugins/axios';
import config from './config';
import router from './router';
import * as filters from './filters';

import App from './App.vue';

Vue.config.productionTip = false;

Vue.use(Buefy);

new Vue({
  router,
  render: h => h(App),
}).$mount('#app');

Object.keys(filters.default).forEach((key) => {
  Vue.filter(key, filters.default[key]);
});

Vue.filter('apiURI', v => `${config.API_URI}${v}`);
