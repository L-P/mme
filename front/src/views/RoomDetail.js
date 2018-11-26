export default {
  data() {
    return {
      room: {},
    };
  },

  mounted() {
    this.$axios.get(`/api/rooms/${this.$route.params.start}`).then((res) => {
      this.room = res.data;
    });
  },
};
