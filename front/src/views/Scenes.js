export default {
  data() {
    return {
      scenes: [],
    };
  },

  mounted() {
    this.$axios.get('/api/scenes').then((res) => {
      this.scenes = res.data;
    });
  },
};
