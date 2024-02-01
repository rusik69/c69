<template>
  <div>
    <h1>VM Stats</h1>
    <table>
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
        <tr v-for="vm in vms" :key="vm.ID">
          <td>{{ vm.id }}</td>
          <td>{{ vm.name }}</td>
          <td>{{ vm.ip }}</td>
          <td>{{ vm.host }}</td>
          <td>{{ vm.status }}</td>
          <td>{{ vm.image }}</td>
          <td>{{ vm.flavor }}</td>
          <td>
            <ul>
              <li v-for="volume in vm.Volumes" :key="volume.ID">{{ volume.Name }}</li>
            </ul>
          </td>
          <td>{{ vm.Committed }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      vms: []
    };
  },
  created() {
    this.fetchVMs();
  },
  methods: {
    fetchVMs() {
      fetch('http://govnocloud-master.rusik69.lol:7070/api/v1/vms')
        .then(response => response.json())
        .then(data => {
          this.vms = data;
        })
        .catch(error => {
          console.error('Error fetching vms:', error);
        });
    }
  }
}
</script>