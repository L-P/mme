<template>
  <div class="container">
    <a :href="uri">
      <img :src="uri">
    </a>
  </div>
</template>

<script>
import config from '@/config';

export default {
  data() {
    return {
      uri: `${config.API_URI}/api/colormap`,
    };
  },
};
</script>
