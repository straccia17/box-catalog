<script setup lang="ts">
import { useItems } from "@/stores/item.store";
import { storeToRefs } from "pinia";
import { ref } from "vue";
const itemStore = useItems();

const label = ref("");
const locationId = ref(-1);
const loading = ref(false);
const { locations } = storeToRefs(itemStore);

const addBox = async () => {
  if (label.value === "" || locationId.value < 0) return;
  try {
    loading.value = true;
    await itemStore.addBox(label.value, locationId.value);
    label.value = "";
  } catch (e) {
    console.error(e);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <h3>Add box</h3>
  <div v-if="loading">Loading...</div>
  <div v-else>
    <label for="label">Label</label>
    <input type="text" name="label" id="label" v-model="label" />
    <label for="locationId">Location</label>
    <select v-model="locationId">
      <option v-for="loc in locations" :value="loc.id">
        {{ loc.label }}
      </option>
    </select>
    <button type="button" @click="addBox()">Add box</button>
  </div>
</template>
