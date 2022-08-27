<script setup lang="ts">
import { useItems } from "@/stores/item.store";
import { storeToRefs } from "pinia";
import { ref } from "vue";
const itemStore = useItems();

const label = ref("");
const boxId = ref(-1);
const categoryId = ref(-1);
const loading = ref(false);
const { boxes, categories } = storeToRefs(itemStore);

const addItem = async () => {
  if (label.value === "" || boxId.value < 0 || categoryId.value < 0) return;
  try {
    loading.value = true;
    await itemStore.addItem(label.value, boxId.value, categoryId.value);
    label.value = "";
  } catch (e) {
    console.error(e);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <h3>Add item</h3>
  <div v-if="loading">Loading...</div>
  <div v-else>
    <label for="label">Label</label>
    <input type="text" name="label" id="label" v-model="label" />
    <label for="boxId">Box</label>
    <select v-model="boxId">
      <option v-for="box in boxes" :value="box.id">
        {{ box.label }} - {{ box.location }}
      </option>
    </select>
    <label for="categoryId">Category</label>
    <select v-model="categoryId">
      <option v-for="cat in categories" :value="cat.id">
        {{ cat.title }}
      </option>
    </select>
    <button type="button" @click="addItem()">Add item</button>
  </div>
</template>
