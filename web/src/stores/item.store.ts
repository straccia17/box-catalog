import type { Box } from '@/models/box.model'
import type { Category } from '@/models/category.model'
import type { Item } from '@/models/item.model'
import type { Location } from '@/models/location.model'
import { getAllBoxes, getAllCategories, getAllItems, getAllLocations, newBox, newCategory, newItem, newLocation } from '@/services/item.service'
import { defineStore } from 'pinia'

export const useItems = defineStore({
  id: 'items',
  state: () => ({
    items: [] as Item[],
    locations: [] as Location[],
    boxes: [] as Box[],
    categories: [] as Category[],
  }),
  actions: {
    async getAllItems() {
      this.locations = await getAllLocations()
      this.categories = await getAllCategories()
      this.boxes = await getAllBoxes()
      this.items = await getAllItems()
    },

    async addLocation(label: string) {
      await newLocation(label)
      this.locations = await getAllLocations()
    },

    async addBox(label: string, locationId: number) {
      await newBox(label, locationId)
      this.boxes = await getAllBoxes()
    },

    async addCategory(title: string) {
      await newCategory(title)
      this.categories = await getAllCategories()
    },

    async addItem(item: string, boxId: number, categoryId: number) {
      await newItem(item, boxId, categoryId)
      this.items = await getAllItems()
    },
  }
})
