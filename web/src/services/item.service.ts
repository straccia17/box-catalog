import type { Box } from "@/models/box.model";
import type { Category } from "@/models/category.model";
import type { Item } from "@/models/item.model";
import type { Location } from "@/models/location.model";
import { client } from "./http.service";

export function newLocation(label: string): Promise<void> {
    return client.post('/locations', {label})
}

export function getAllLocations(): Promise<Location[]> {
    return client.get<Location[]>('/locations').then( resp => resp.data)
}

export function newBox(label: string, locationId: number): Promise<void> {
    return client.post('/boxes', {label, locationId})
}

export function getAllBoxes(): Promise<Box[]> {
    return client.get<Box[]>('/boxes').then( resp => resp.data)
}

export function newCategory(title: string): Promise<void> {
    return client.post('/categories', {title})
}

export function getAllCategories(): Promise<Category[]> {
    return client.get<Category[]>('/categories').then( resp => resp.data)
}

export function newItem(item: string, boxId: number, categoryId: number): Promise<void> {
    return client.post('/items', {item, boxId, categoryId})
}

export function getAllItems(): Promise<Item[]> {
    return client.get<Item[]>('/items').then( resp => resp.data)
}