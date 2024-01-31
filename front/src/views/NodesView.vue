<template>
  <div class="nodes">
    <h1>Nodes</h1>
    <table>
      <thead>
        <tr>
          <th>Node</th>
          <th>Total CPU</th>
          <th>Free CPU</th>
          <th>Total Memory</th>
          <th>Free Memory</th>
          <th>Total Disk</th>
          <th>Free Disk</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="node in nodes" :key="node.id">
          <td>{{ node.name }}</td>
          <td>{{ node.cpu.total }}</td>
          <td>{{ node.cpu.free }}</td>
          <td>{{ node.memory.total }}</td>
          <td>{{ node.memory.free }}</td>
          <td>{{ node.disk.total }}</td>
          <td>{{ node.disk.free }}</td>
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
    }
  }
}
</script>