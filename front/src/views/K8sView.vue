<template>
  <div>
    <table class="bordered-table">
      <thead>
        <tr>
          <th>Name</th>
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
          <td>{{ k8.flavor }}</td>
          <td>{{ k8.vm.host }}</td>
          <td>{{ k8.vm.state }}</td>
          <td>{{ k8.vm.ip }}</td>
          <td>{{ k8.vm.tailscaleip }}</td>
          <td> <button @click="startK8S(k8.name)">Start</button> </td>
          <td> <button @click="stopK8S(k8.name)">Stop</button> </td>
          <td> <button @click="terminateK8S(k8.name)">Terminate</button> </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      k8s: [],
      intervalId: null,
    };
  },
  created() {
    this.fetchK8S();
  },
  methods: {
    fetchK8S() {
      fetch('http://master.govno.cloud:6969/api/v1/k8s')
        .then(response => response.json())
        .then(data => {
          this.k8s = data;
        })
        .catch(error => {
          console.error('Error fetching k8s:', error);
        });
    },
    startK8S(name) {
      fetch(`http://master.govno.cloud:6969/api/v1/k8sstart/${name}`, {
        method: 'GET'
      })
        .then(() => {
          this.fetchK8s();
        })
        .catch(error => {
          console.error('Error starting k8s:', error);
        });
    },
    stopK8S(name) {
      fetch(`http://master.govno.cloud:6969/api/v1/k8sstop/${name}`, {
        method: 'GET'
      })
        .then(() => {
          this.fetchK8s();
        })
        .catch(error => {
          console.error('Error stopping k8s:', error);
        });
    },
    terminateK8S(name) {
      fetch(`http://master.govno.cloud:6969/api/v1/k8s/${name}`, {
        method: 'DELETE'
      })
        .then(() => {
          this.fetchK8s();
        })
        .catch(error => {
          console.error('Error terminating k8s:', error);
        });
      },
      mounted() {
        this.fetchK8s();
        this.intervalId = setInterval(this.fetchK8s, 5000); // Fetch k8s every 5 seconds
      },
      beforeDestroy() {
        clearInterval(this.intervalId); // Clear the interval when the component is destroyed
      },
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