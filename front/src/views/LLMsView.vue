<template>
  <div>
    <table class="bordered-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>ID</th>
          <th>Model</th>
          <th>Flavor</th>
          <th>Node</th>
          <th>IP</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="llm in llms" :key="llm.name">
          <td>{{ llm.name }}</td>
          <td>{{ llm.id }}</td>
          <td>{{ llm.model }}</td>
          <td>{{ llm.container.flavor }}</td>
          <td>{{ k8.container.host }}</td>
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
      fetch('http://master.govno.cloud:6969/api/v1/llm')
        .then(response => response.json())
        .then(data => {
          this.llms = data;
        })
        .catch(error => {
          console.error('Error fetching llms:', error);
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