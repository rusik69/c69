<template>
  <div>
    <table class="bordered-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>ID</th>
          <th>Image</th>
          <th>Status</th>
          <th>Flavor</th>
          <th>Node</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="container in containers" :key="container.ID">
          <td>{{ container.name }}</td>
          <td>{{ container.id.substring(0, 8) }}</td>
          <td>{{ container.image }}</td>
          <td>{{ container.state }}</td>
          <td>{{ container.flavor }}</td>
          <td>{{ container.host }}</td>
          <td> <button @click="startContainer(container.name)">Start</button> </td>
          <td> <button @click="stopContainer(vm.Name)">Stop</button> </td>
          <td> <button @click="terminateContainer(vm)">Terminate</button> </td>
        </tr>
      </tbody>
    </table>
    <button @click = "showCreateDialog = true"> Create Container</button>
    <div v-if="showCreateDialog">
      <h2>Create Container</h2>
      <label>
        Name:
        <input v-model="newContainer.name" type="text" />
      </label>
      <label>
        Image:
        <input v-model="newContainer.image" type="text" />
      </label>
      <label>
        Flavor:
        <select v-model="newContainer.flavor">
          <option v-for="flavor in flavors" :key="flavor" :value="flavor">{{ flavor }}</option>
        </select>
      </label>
      <button @click="createContainer">Create</button>
      <button @click="showCreateDialog = false">Cancel</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      containers: [],
      showCreateDialog: false,
      newContainer: {
        name: '',
        image: '',
        flavor: ''
      },
      flavors: ['small', 'medium', 'large', "xlarge", "2xlarge"]
    };
  },
  created() {
    this.fetchContainers();
  },
  methods: {
    fetchContainers() {
      fetch('http://master.govno.cloud:6969/api/v1/containers')
        .then(response => response.json())
        .then(data => {
          this.containers = data;
        })
        .catch(error => {
          console.error('Error fetching containers:', error);
        });
    },
    mounted(){
      this.IntervalId = setInterval(this.fetchContainers, 5000)
    },
    createContainer() {
      fetch('http://master.govno.cloud:6969/api/v1/containers', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          name: this.newContainer.name,
          image: this.newContainer.image,
          flavor: this.newContainer.flavor,
        })
      })
        .then(() => {
          this.fetchContainers();
          this.showCreateDialog = false;
        })
        .catch(error => {
          console.error('Error creating container:', error);
        });
    },
    startContainer(name) {
      fetch(`http://master.govno.cloud:6969/api/v1/containerstart/` + name, {
        method: 'GET'
      })
        .then(() => {
          this.fetchContainers();
        })
        .catch(error => {
          console.error('Error starting container:', error);
        });
    },
    stopContainer(name) {
      fetch(`http://master.govno.cloud:6969/api/v1/containerstop/` + name, {
        method: 'GET'
      })
        .then(() => {
          this.fetchContainers();
        })
        .catch(error => {
          console.error('Error stopping container:', error);
        });
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
