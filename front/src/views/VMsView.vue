<template>
  <div id="vms">
    <table class="bordered-table">
      <thead>
        <tr>
          <th>Name</th>
          <th>ID</th>
          <th>IP</th>
          <th>TailscaleIP</th>
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
          <td>{{ vm.tailscaleip }}</td>
          <td>{{ vm.host }}</td>
          <td>{{ vm.state }}</td>
          <td>{{ vm.image }}</td>
          <td>{{ vm.flavor }}</td>
          <td>
            <ul>
              <li v-for="volume in vm.Volumes" :key="volume.ID">
                {{ volume.Name }}
              </li>
            </ul>
          </td>
          <td>{{ vm.Committed }}</td>
        </tr>
      </tbody>
    </table>
    <div id="details" v-if="selectedvm !== null">
      <vm-details :vm="selectedvm"> </vm-details>
    </div>
    <button @click="showCreateDialog = true">Create VM</button>
    <div v-if="showCreateDialog">
      <h2>Create VM</h2>
      <label>
        VM Name:
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
import VmDetails from "@/views/VmDetails.vue";
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
    }
  },

  components: {
    "vm-details": VmDetails,
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