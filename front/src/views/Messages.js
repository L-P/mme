export default {
  data() {
    return {
      messages: [],
    };
  },

  mounted() {
    this.$axios.get('/api/messages').then((res) => {
      this.messages = res.data;
    });
  },
};
