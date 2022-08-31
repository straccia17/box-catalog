<script setup lang="ts">
import { RouterView } from "vue-router";
</script>

<template>
  <nav>
    <ul>
      <li>
        <RouterLink to="/">Home</RouterLink>
      </li>
      <li>
        <RouterLink to="/items">Items</RouterLink>
      </li>
    </ul>
  </nav>
  <RouterView />
</template>
