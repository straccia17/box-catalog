<script setup lang="ts">
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
  </main>
</template>
