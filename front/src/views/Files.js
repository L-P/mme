export default {
  data() {
    return {
      files: [],
    };
  },

  mounted() {
    this.$axios.get('/api/files').then((res) => {
      this.files = res.data;
    });
  },
};
