<template>
  <div class="about">
    <h1>Files</h1>
    <ul>
      <li v-for="file in files">{{ file }}</li>
    </ul>
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
    this.fetchFiles();
  },
  methods: {
    fetchFiles() {
      fetch('http://govnocloud-master.rusik69.lol:7070/api/v1/files')
        .then(response => response.json())
        .then(data => {
          this.files = data;
        })
        .catch(error => {
          console.error('Error fetching files:', error);
        });
    }
  }
}
</script>