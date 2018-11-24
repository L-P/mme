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
      colormap: 'http://localhost:3000/api/colormap',
    }
  }
}
</script>
