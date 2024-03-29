<template>
  <div>
    <table class="bordered-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>ID</th>
          <th>Flavor</th>
          <th>Node</th>
          <th>Status</th>
          <th>IP</th>
          <th>TailscaleIP</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="k8 in k8s" :key="k8.name">
          <td>{{ k8.name }}</td>
          <td>{{ k8.size }}</td>
          <td>{{ k8.flavor }}</td>
          <td>{{ k8.vm.node }}</td>
          <td>{{ k8.vm.host }}</td>
          <td>{{ k8.vm.ip }}</td>
          <td>{{ k8.vm.tailscaleip }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      files: []
    };
  },
  created() {
    this.fetchFiles();
  },
  methods: {
    fetchFiles() {
      fetch('http://master.govno.cloud:6969/api/v1/k8s')
        .then(response => response.json())
        .then(data => {
          this.files = data;
        })
        .catch(error => {
          console.error('Error fetching k8s:', error);
        });
    }
  }
}
</script>

<style scoped>
.bordered-table {
  border-collapse: collapse;
  width: 100%;
}

.bordered-table th, .bordered-table td {
  border: 1px solid green;
  padding: 10px
}
</style>