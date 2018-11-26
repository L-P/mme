import Vue from 'vue';
import Router from 'vue-router';

import ColorMap from './views/ColorMap.vue';
import Files from './views/Files.vue';
import Home from './views/Home.vue';
import Messages from './views/Messages.vue';
import RoomDetail from './views/RoomDetail.vue';
import SceneDetail from './views/SceneDetail.vue';
import Scenes from './views/Scenes.vue';

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
    {
      path: '/scenes/:start',
      component: SceneDetail,
      name: 'SceneDetail',
    },
    {
      path: '/room/:start',
      component: RoomDetail,
      name: 'RoomDetail',
    },
  ],
});
