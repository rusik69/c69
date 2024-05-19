<template>
  <div id="vms">
    <table class="bordered-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>ID</th>
          <th>IP</th>
          <th>Node</th>
          <th>State</th>
          <th>Image</th>
          <th>Flavor</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="vm in vms"
          :key="vm.ID"
          v-on:click="expand(vm)"
          :class="{ selected_row: selectedvm && selectedvm.id === vm.id }"
        >
          <td>{{ vm.name }}</td>
          <td>{{ vm.id }}</td>
          <td>{{ vm.ip }}</td>
          <td>{{ vm.host }}</td>
          <td>{{ vm.state }}</td>
          <td>{{ vm.image }}</td>
          <td>{{ vm.flavor }}</td>
          <td> <button @click="startVM(vm.name)"> Start</button> </td>
          <td> <button @click="stopVM(vm.name)"> Stop</button> </td>
          <td> <button @click="terminateVM(vm.name)">Terminate</button> </td>
        </tr>
      </tbody>
    </table>
    <button @click="showCreateDialog = true">Create VM</button>
    <div v-if="showCreateDialog">
      <h2>Create VM</h2>
      <label>
        Name:
        <input v-model="newVm.name" type="text" />
      </label>
      <label>
        Image:
        <select v-model="newVm.image">
          <option v-for="image in images" :key="image" :value="image">{{ image }}</option>
        </select>
      </label>
      <label>
        Flavor:
        <select v-model="newVm.flavor">
          <option v-for="flavor in flavors" :key="flavor" :value="flavor">{{ flavor }}</option>
        </select>
      </label>
      <button @click="createVM">Create</button>
      <button @click="showCreateDialog = false">Cancel</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      vms: [],
      selectedvm : null,
      showCreateDialog: false,
      newVm: {
        name: '',
        image: '',
        flavor: '',
      },
      images: ['ubuntu22.04'],
      flavors: ['small', 'medium', 'large', 'xlarge', '2xlarge'],
      intervalId: null,
    }
  },
  created() {
    this.fetchVMs();
  },
  methods: {
    expand(vm) {
      this.selectedvm = vm;
    },
    fetchVMs() {
      fetch("http://master.govno.cloud:6969/api/v1/vms")
        .then((response) => response.json())
        .then((data) => {
          this.vms = data;
        })
        .catch((error) => {
          console.error("Error fetching vms:", error);
        });
    },
    mounted() {
      this.fetchVMs();
      this.intervalId = setInterval(this.fetchVMs, 5000); // Fetch VMs every 5 seconds
    },
    beforeDestroy() {
      clearInterval(this.intervalId); // Clear the interval when the component is destroyed
    },
    createVM() {
      fetch("http://master.govno.cloud:6969/api/v1/vms", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: this.newVm.name,
          image: this.newVm.image,
          flavor: this.newVm.flavor,
          state: "creating",
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          this.vms.push(data);
        })
        .catch((error) => {
          console.error("Error creating vm:", error);
        });
        this.newVm = { name: '', image: '', flavor: '' };
        this.showCreateDialog = false;
    },
    startVM(vmName) {
      fetch("http://master.govno.cloud:6969/api/v1/vmstart/" + vmName, {
        method: "GET",
      })
        .then((response) => response.json())
        .then((data) => {
          this.vms.push(data);
        })
        .catch((error) => {
          console.error("Error starting vm:", error);
        });
    },
    stopVM(vmName) {
      fetch("http://master.govno.cloud:6969/api/v1/vmstop/" + vmName, {
        method: "GET",
      })
        .then((response) => response.json())
        .then((data) => {
          this.vms.push(data);
        })
        .catch((error) => {
          console.error("Error stopping vm:", error);
        });
    },
    terminateVM(vmName) {
      fetch("http://master.govno.cloud:6969/api/v1/vm/" + vmName, {
        method: "DELETE",
      })
        .then((response) => response.json())
        .then((data) => {
          this.vms = this.vms.filter((v) => v.id !== vm.id);
        })
        .catch((error) => {
          console.error("Error deleting vm:", error);
        });
    },
  },
};
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