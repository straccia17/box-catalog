<script setup lang="ts">
import type { Item } from "@/models/item.model";
import { NCard, NTag } from "naive-ui";

defineProps<{
  item: Item;
}>();
</script>

<template>
  <n-card :title="item.item">
    <template #header-extra>
      <n-tag size="small" round>
        {{ item.category }}
      </n-tag>
    </template>
    {{ item.box }} - {{ item.location }}
  </n-card>
</template>

<style scoped>
.n-card {
  margin-bottom: 1rem;
}
</style>
