export default {
  data() {
    return {
      rom: {},
    };
  },

  mounted() {
    this.$axios.get('/api/rom').then((res) => {
      this.rom = res.data;
    });
  },
};
