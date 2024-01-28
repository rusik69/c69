<template>
  <div class="containers">
    <h1>Containers</h1>
    <ul>
      <li v-for="container in containers">{{ container }}</li>
    </ul>
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
      fetch('http://govnocloud-master.rusik69.lol:7070/api/v1/containers')
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