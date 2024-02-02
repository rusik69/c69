<template>
  <h1>VM Details</h1>
  <div id="details">
    <p>Details for VM {{ vm.id }} - {{ vm.name }}</p>
    <vue-vnc :url="vncurl"></vue-vnc>
  </div>
</template>

<script>
import VueVnc from "vue-vnc";
const VmDetails = {
  props: ["vm"],
  data() {
    return {
      vncurl: "",
    };
  },
  created: function () {
    this.getHostName();
  },
  name: "VmDetails",
  components: {
    'vue-vnc': VueVnc,
  },
  methods: {
    getHostName() {
      fetch(`http://govnocloud-master.rusik69.lol:7070/api/v1/node/${this.vm.host}`)
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          this.vncurl = `ws://${data.ip}:${this.vm.vnc_port}`;
          console.log(this.vncurl);
        });
    }
  },
};
export default VmDetails;
</script>
