<script setup lang="ts">
import { useItems } from "@/stores/item.store";
import { ref } from "vue";

const category = ref("");
const loading = ref(false);
const itemStore = useItems();

const addCategory = async () => {
  if (category.value === "") return;
  try {
    loading.value = true;
    await itemStore.addCategory(category.value);
    category.value = "";
  } catch (e) {
    console.error(e);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <h3>Add category</h3>
  <div v-if="loading">Loading...</div>
  <div v-else>
    <label for="category">Category</label>
    <input type="text" name="category" id="category" v-model="category" />
    <button type="button" @click="addCategory()">Add category</button>
  </div>
</template>
