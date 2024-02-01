<template>
  <div>
    <h1>Nodes</h1>
    <table>
      <thead>
        <tr>
          <th>Node</th>
          <th>Total CPU</th>
          <th>Total Memory</th>
          <th>Free Memory</th>
          <th>Total Disk</th>
          <th>Free Disk</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(nodeStats, nodeName) in nodes" :key="nodeName">
          <td>{{ nodeName }}</td>
          <td>{{ nodeStats.cpus }}</td>
          <td>{{ humanize(nodeStats.total_mem) }}</td>
          <td>{{ humanize(nodeStats.mem) }}</td>
          <td>{{ humanize(nodeStats.total_disk) }}</td>
          <td>{{ humanize(nodeStats.disk) }}</td>
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
      fetch('http://govnocloud-master.rusik69.lol:7070/api/v1/nodes')
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