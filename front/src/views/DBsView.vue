<template>
  <div>
    <table class="bordered-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>ID</th>
          <th>Type</th>
          <th>Flavor</th>
          <th>Node</th>
          <th>IP</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="db in dbs" :key="db.name">
          <td>{{ db.name }}</td>
          <td>{{ db.id }}</td>
          <td>{{ db.type }}</td>
          <td>{{ db.container.flavor }}</td>
          <td>{{ k8.container.node }}</td>
          <td>{{ k8.container.ip }}</td>
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
      fetch('http://master.govno.cloud:6969/api/v1/db')
        .then(response => response.json())
        .then(data => {
          this.dbs = data;
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