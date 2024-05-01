<template>
  <div>
    <table class="bordered-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Total milliCPUs</th>
          <th>Free milliCPUs</th>
          <th>Total Memory</th>
          <th>Free Memory</th>
          <th>Total Disk</th>
          <th>Free Disk</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(nodeStats, nodeName) in nodes" :key="nodeName">
          <td>{{ nodeName }}</td>
          <td>{{ nodeStats.total_millicpus }}</td>
          <td>{{ nodeStats.free_millicpus }}</td>
          <td>{{ humanize(nodeStats.total_mem) }}</td>
          <td>{{ humanize(nodeStats.free_mem) }}</td>
          <td>{{ humanize(nodeStats.total_disk) }}</td>
          <td>{{ humanize(nodeStats.free_disk) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      nodes: []
    };
  },
  created() {
    this.fetchNodes();
  },
  methods: {
    fetchNodes() {
      fetch('http://master.govno.cloud:6969/api/v1/nodes')
        .then(response => response.json())
        .then(data => {
          this.nodes = data;
        })
        .catch(error => {
          console.error('Error fetching nodes:', error);
        });
    },
    humanize(value) {
      const units = ['B', 'KB', 'MB', 'GB', 'TB'];
      let index = 0;
      while (value >= 1024 && index < units.length - 1) {
        value /= 1024;
        index++;
      }
      return `${value.toFixed(2)} ${units[index]}`;
    }
  }
}
</script>

<style scoped>
.nodes {
  border: 2px solid green;
  padding: 20px;
}

.bordered-table {
  border-collapse: collapse;
  width: 100%;
}

.bordered-table th, .bordered-table td {
  border: 1px solid green;
  padding: 10px
}
</style>