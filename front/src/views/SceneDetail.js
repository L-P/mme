export default {
  data() {
    return {
      scene: {},
    };
  },

  mounted() {
    this.$axios.get(`/api/scenes/${this.$route.params.start}`).then((res) => {
      this.scene = res.data;
    });
  },
};
