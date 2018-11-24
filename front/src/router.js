import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import ColorMap from './views/ColorMap.vue';
import Scenes from './views/Scenes.vue';
import Files from './views/Files.vue';
import Messages from './views/Messages.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Home,
    },
    {
      path: '/colormap',
      component: ColorMap,
    },
    {
      path: '/scenes',
      component: Scenes,
    },
    {
      path: '/files',
      component: Files,
    },
    {
      path: '/messages',
      component: Messages,
    },
  ],
});
