<template>
  <h1>VM Details</h1>
  <div id="details">
    <p>Details for VM {{ id }} - {{ name }}</p>
  </div>
</template>

<script>
const VmDetails = {
  props: ["id", "name"],

  name: "VmDetails",
  methods: {
    getVmDetails() {
      console.log("Getting details for VM", this.id);
    },
  },
};
export default VmDetails;
</script>
