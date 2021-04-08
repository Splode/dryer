var test = new Vue({
  el: "#test",
  data: {
    statusProxy: 32+1.160,
  },










  computed: {
    status: {
      get() {
        return this.statusProxy === null ? true : this.statusProxy;
      },
      set(val) {
        this.statusProxy = val += 1;
      },
    },
  },
});
