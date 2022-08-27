<script setup lang="ts">
import { useItems } from "@/stores/item.store";
import { storeToRefs } from "pinia";
import { onMounted, ref } from "vue";

const itemStore = useItems();
const loading = ref(true);
const { items } = storeToRefs(itemStore);

const loadItems = async () => {
  loading.value = true;
  try {
    await itemStore.getAllItems();
  } catch (e) {
    console.error(e);
  } finally {
    loading.value = false;
  }
};

onMounted(() => loadItems());
</script>

<template>
  <div>
    <h3>Items</h3>
    <p v-if="loading">Loading...</p>
    <div v-else>
      <div v-for="item in items">
        <h4>{{ item.item }}</h4>
        <span>{{ item.box }} - {{ item.location }}</span>
        <span>{{ item.category }}</span>
      </div>
    </div>
  </div>
</template>
