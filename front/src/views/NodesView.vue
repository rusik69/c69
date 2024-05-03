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
        <tr v-for="node in nodes" :key="nodeName">
          <td>{{ node.name }}</td>
          <td>{{ node.stats.total_millicpus }}</td>
          <td>{{ node.stats.free_millicpus }}</td>
          <td>{{ humanize(node.stats.total_mem) }}</td>
          <td>{{ humanize(node.stats.free_mem) }}</td>
          <td>{{ humanize(node.stats.total_disk) }}</td>
          <td>{{ humanize(node.stats.free_disk) }}</td>
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
      const units = ['MB', 'GB', 'TB'];
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