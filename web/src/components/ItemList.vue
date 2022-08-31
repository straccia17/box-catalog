<script setup lang="ts">
import { useItems } from "@/stores/item.store";
import { storeToRefs } from "pinia";
import { onMounted, ref } from "vue";
import ItemDetail from "./ItemDetail.vue";

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
  <p v-if="loading">Loading...</p>
  <div v-else>
    <item-detail :item="item" v-for="item in items"></item-detail>
  </div>
</template>
