<template>
  <div id="vms">
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
        <tr
          v-for="vm in vms"
          :key="vm.ID"
          v-on:click="expand(vm)"
          :class="{ selected_row: detailsId === vm.id }"
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
    <div id="details" v-if="detailsId !== 0">
      <vm-details :id="detailsId" :name="detailsName"> </vm-details>
    </div>
  </div>
</template>

<script>
import VmDetails from "@/views/VmDetails.vue";
export default {
  data() {
    return {
      vms: [
        {
          id: 48,
          name: "test0",
          ip: "",
          host: "x220",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
        {
          id: 49,
          name: "test1",
          ip: "",
          host: "x220",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
        {
          id: 50,
          name: "test2",
          ip: "",
          host: "x220",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
        {
          id: 11,
          name: "test3",
          ip: "",
          host: "x230",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
        {
          id: 12,
          name: "test4",
          ip: "",
          host: "x230",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
        {
          id: 13,
          name: "test5",
          ip: "",
          host: "x230",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
        {
          id: 14,
          name: "test6",
          ip: "",
          host: "x230",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
        {
          id: 15,
          name: "test7",
          ip: "",
          host: "x230",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
        {
          id: 16,
          name: "test8",
          ip: "",
          host: "x230",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
        {
          id: 51,
          name: "test9",
          ip: "",
          host: "x220",
          status: "",
          image: "ubuntu22.04",
          flavor: "tiny",
          volumes: null,
          committed: true,
        },
      ],
      detailsId: 0,
      detailsName: "",
    };
  },

  components: {
    "vm-details": VmDetails,
  },
  created() {
    this.fetchVMs();
  },
  methods: {
    expand(vm) {
      console.log("Expanding VM", vm);
      this.detailsId = vm.id;
      this.detailsName = vm.name;
    },
    fetchVMs() {
      fetch("http://govnocloud-master.rusik69.lol:7070/api/v1/vms")
        .then((response) => response.json())
        .then((data) => {
          //this.vms = data;
        })
        .catch((error) => {
          console.error("Error fetching vms:", error);
        });
    },
  },
};
</script>
