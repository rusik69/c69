<template>
  <body>
  <h1>VM Details</h1>
  <div id="details">
    <p>Details for VM {{ vm.id }} - {{ vm.name }}</p>
    <vue-vnc :url="url"></vue-vnc>
  </div>
  </body>
</template>

<script>
import VueVnc from "vue-vnc";
const VmDetails = {
  props: ["vm"],
  data() {
    return {
      vm: "",
    };
  },
  computed: {
    url() {
      if (!this.vm) {
        return "";
      }
      return `ws://${this.vm.nodehostname}://${this.vm.vnc_port}`;
    }
  },
  name: "VmDetails",
  components: {
    'vue-vnc': VueVnc,
  },
  methods: {
   
  },
};
export default VmDetails;
</script>
