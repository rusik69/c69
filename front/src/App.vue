<script setup>
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <header>
    <div>
      <h2>Govnocloud</h2>
      <nav>
        <RouterLink to="/nodes" class="nav-item">Nodes</RouterLink>
        <RouterLink to="/vms" class="nav-item">VMs</RouterLink>
        <RouterLink to="/containers" class="nav-item">Containers</RouterLink>
        <RouterLink to="/files" class="nav-item">Files</RouterLink>
      </nav>
    </div>
  </header>
  <div id="router_body">
    <RouterView />
  </div>
</template>

<style scoped>
nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
  
}
.nav-item.selected {
  background-color: green
}

nav a:first-of-type {
  border: 0;
}

</style>