<template>
  <div class="container">
    <img :src="colormap" />
  </div>
</template>

<script>
import Vue from 'vue';

export default {
  data() {
    return {
      colormap: 'http://localhost:8064/api/colormap',
    }
  }
}
</script>
