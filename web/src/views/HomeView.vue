<script setup lang="ts">
import ItemList from "../components/ItemList.vue";
import AddLocation from "../components/AddLocation.vue";
import AddCategory from "../components/AddCategory.vue";
import AddBox from "../components/AddBox.vue";
import AddItem from "../components/AddItem.vue";
</script>

<template>
  <main>
    <h1>Home</h1>
    <AddCategory></AddCategory>
    <AddLocation></AddLocation>
    <AddBox></AddBox>
    <AddItem></AddItem>
    <ItemList></ItemList>
  </main>
</template>
