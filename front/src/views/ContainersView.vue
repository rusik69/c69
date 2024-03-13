<template>
  <div>
    <table class="bordered-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>ID</th>
          <th>Image</th>
          <th>Status</th>
          <th>Node</th>
          <th>Ports</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="container in containers" :key="container.ID">
          <td>{{ container.name }}</td>
          <td>{{ container.id.substring(0, 8) }}</td>
          <td>{{ container.image }}</td>
          <td>{{ container.state }}</td>
          <td>{{ container.host }}</td>
          <td>
            <ul>
              <li v-for="port in container.Ports" :key="port">{{ port }}</li>
            </ul>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      containers: []
    };
  },
  created() {
    this.fetchContainers();
  },
  methods: {
    fetchContainers() {
      fetch('http://master.govno.cloud:7070/api/v1/containers')
        .then(response => response.json())
        .then(data => {
          this.containers = data;
        })
        .catch(error => {
          console.error('Error fetching containers:', error);
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
