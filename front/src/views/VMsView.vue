<template>
  <div id="vms" style="vms">
    <h1>VM Stats</h1>
    <table class="bordered-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
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
          <td>{{ vm.id }}</td>
          <td>{{ vm.name }}</td>
          <td>{{ vm.ip }}</td>
          <td>{{ vm.host }}</td>
          <td>{{ vm.status }}</td>
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
  </div>
</template>

<script>
import VmDetails from "@/views/VmDetails.vue";
export default {
  data() {
    return {
      vms: [],
      selectedvm : null,
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
      fetch("http://govnocloud-master.rusik69.lol:7070/api/v1/vms")
        .then((response) => response.json())
        .then((data) => {
          this.vms = data;
        })
        .catch((error) => {
          console.error("Error fetching vms:", error);
        });
    },
  },
};
</script>

<style scoped>
.vms {
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