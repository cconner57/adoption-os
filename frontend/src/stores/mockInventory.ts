import { computed,ref } from 'vue'

export interface IInventoryItem {
  id: string
  name: string
  category: 'Food' | 'Medical' | 'Cleaning' | 'Office' | 'Merch' | 'Other'
  quantity: number
  unit: string
  minThreshold: number
  location: string
  lastUpdated: string
}

export const mockInventory = ref<IInventoryItem[]>([
  {
    id: 'inv-1',
    name: 'Dry Dog Food (Adult)',
    category: 'Food',
    quantity: 45,
    unit: 'bags',
    minThreshold: 10,
    location: 'Warehouse A',
    lastUpdated: '2023-10-25',
  },
  {
    id: 'inv-2',
    name: 'Wet Cat Food (Cans)',
    category: 'Food',
    quantity: 8,
    unit: 'cases',
    minThreshold: 15, 
    location: 'Pantry',
    lastUpdated: '2023-10-28',
  },
  {
    id: 'inv-3',
    name: 'Bleach',
    category: 'Cleaning',
    quantity: 12,
    unit: 'gallons',
    minThreshold: 5,
    location: 'Janitorial Closet',
    lastUpdated: '2023-10-20',
  },
  {
    id: 'inv-4',
    name: 'Vaccines (Da2PP)',
    category: 'Medical',
    quantity: 2,
    unit: 'vials',
    minThreshold: 10, 
    location: 'Medical Fridge',
    lastUpdated: '2023-10-28',
  },
  {
    id: 'inv-5',
    name: 'Paper Towels',
    category: 'Cleaning',
    quantity: 50,
    unit: 'rolls',
    minThreshold: 20,
    location: 'Janitorial Closet',
    lastUpdated: '2023-10-22',
  },
  {
    id: 'inv-6',
    name: 'Printer Paper',
    category: 'Office',
    quantity: 4,
    unit: 'reams',
    minThreshold: 2,
    location: 'Office Shelf',
    lastUpdated: '2023-09-15',
  },
])

export const inventoryStats = computed(() => {
  const lowStock = mockInventory.value.filter((i) => i.quantity <= i.minThreshold)
  const critical = mockInventory.value.filter((i) => i.quantity <= i.minThreshold / 2) 

  return {
    totalItems: mockInventory.value.length,
    lowStockCount: lowStock.length,
    criticalCount: critical.length,
    categories: [...new Set(mockInventory.value.map((i) => i.category))].length,
  }
})
