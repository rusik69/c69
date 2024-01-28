<template>
  <div class="about">
    <h1>VMs</h1>
    <ul>
      <li v-for="vm in vms">{{ vm }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      vms: []
    };
  },
  created() {
    this.fetchVMs();
  },
  methods: {
    fetchVMs() {
      fetch('http://govnocloud-master.rusik69.lol:7070/api/v1/vms')
        .then(response => response.json())
        .then(data => {
          this.vms = data;
        })
        .catch(error => {
          console.error('Error fetching vms:', error);
        });
    }
  }
}
</script>