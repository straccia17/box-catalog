<script setup lang="ts">
import { useItems } from "@/stores/item.store";
import { ref } from "vue";

const location = ref("");
const loading = ref(false);
const itemStore = useItems();

const addLocation = async () => {
  if (location.value === "") return;
  try {
    loading.value = true;
    await itemStore.addLocation(location.value);
    location.value = "";
  } catch (e) {
    console.error(e);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <h3>Add location</h3>
  <div v-if="loading">Loading...</div>
  <div v-else>
    <label for="location">Location</label>
    <input type="text" name="location" id="location" v-model="location" />
    <button type="button" @click="addLocation()">Add location</button>
  </div>
</template>
