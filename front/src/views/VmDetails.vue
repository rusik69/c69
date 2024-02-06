<template>
  <body>
  <h1>VM Details</h1>
  <div id="details">
    <p>Details for VM {{ vm.id }} - {{ vm.name }}</p>
    <vue-vnc :url="vncURL"></vue-vnc>
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
    vncURL() {
      return vm.vncurl
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