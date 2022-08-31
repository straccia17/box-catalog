<script setup lang="ts">
import type { MenuOption } from "naive-ui";
import { NMenu } from "naive-ui";
import { h } from "vue";
import { RouterLink, RouterView } from "vue-router";

const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            path: '/'
          }
        },
        { default: () => 'Home' }
      ),
    key: 'go-back-home',
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            path: '/items'
          }
        },
        { default: () => 'Items' }
      ),
    key: 'go-to-items',
  }
]
</script>

<template>
  <n-menu :options="menuOptions" mode="horizontal" />
  <RouterView />
</template>
