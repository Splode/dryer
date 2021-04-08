var demo = new Vue({
  el: "#demo",
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
