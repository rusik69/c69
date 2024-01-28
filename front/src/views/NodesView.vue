<template>
  <div class="about">
    <h1>Nodes</h1>
    <ul>
      <li v-for="node in nodes">{{ node }}</li>
    </ul>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>

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