<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'

import { Button,Capsules, InputField } from '../../components/common/ui'
import { type IInventoryItem,inventoryStats, mockInventory } from '../../stores/mockInventory'

const searchQuery = ref('')
const filterCategory = ref<'All' | string>('All')
const filterStatus = ref<'All' | 'Low Stock'>('All')

const categories = computed(() => ['All', ...new Set(mockInventory.value.map((i) => i.category))])

const filteredItems = computed(() => {
  return mockInventory.value
    .filter((item) => {

      const searchMatch =
        !searchQuery.value ||
        item.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        item.location.toLowerCase().includes(searchQuery.value.toLowerCase())

      const catMatch = filterCategory.value === 'All' || item.category === filterCategory.value

      const isLow = item.quantity <= item.minThreshold
      const statusMatch =
        filterStatus.value === 'All' || (filterStatus.value === 'Low Stock' && isLow)

      return searchMatch && catMatch && statusMatch
    })
    .sort((a, b) => {

      const aRatio = a.quantity / a.minThreshold
      const bRatio = b.quantity / b.minThreshold
      return aRatio - bRatio
    })
})

const getStockStatus = (item: IInventoryItem) => {
  if (item.quantity === 0) return { label: 'Out of Stock', color: '#fee2e2' }
  if (item.quantity <= item.minThreshold) return { label: 'Low Stock', color: '#fef3c7' }
  return { label: 'In Stock', color: '#d1fae5' }
}

const getStockWidth = (item: IInventoryItem) => {

  const max = item.minThreshold * 2
  return `${Math.min((item.quantity / max) * 100, 100)  }%`
}

import StockAdjustmentModal from '../../components/admin/inventory/StockAdjustmentModal.vue'

const showEditModal = ref(false)
const editingItem = ref<IInventoryItem | null>(null)

const openEditModal = (item: IInventoryItem) => {
  editingItem.value = item
  showEditModal.value = true
}

const handleSaveStock = (item: IInventoryItem, quantity: number) => {
  item.quantity = quantity
  item.lastUpdated = new Date().toISOString().split('T')[0]
  showEditModal.value = false
  editingItem.value = null
}

const addItem = () => {
  const name = prompt('Enter item name:')
  if (name) {
    mockInventory.value.push({
      id: `inv-${Date.now()}`,
      name,
      category: 'Other',
      quantity: 0,
      unit: 'units',
      minThreshold: 5,
      location: 'Unassigned',
      lastUpdated: new Date().toISOString().split('T')[0],
    })
  }
}

const isMounted = ref(false)
onMounted(() => {
  isMounted.value = true
})
</script>

<template>
  <div class="inventory-page">
    <Teleport v-if="isMounted" to="#mobile-header-target" :disabled="false">
      <h1 class="mobile-header-title">Inventory Management</h1>
    </Teleport>

    <div class="page-header">
      <h1 class="desktop-only">Inventory Management</h1>
      <Button title="+ Add Item" color="white" :onClick="addItem" />
    </div>

    <div v-if="inventoryStats.lowStockCount > 0" class="alert-banner">
      <span class="alert-icon">⚠️</span>
      <span class="alert-text">
        <strong>Action Required:</strong> {{ inventoryStats.lowStockCount }} items are below stock
        threshold.
        <span class="link-text" @click="filterStatus = 'Low Stock'">View Items</span>
      </span>
    </div>

    <div class="stats-grid">
      <div class="stat-card">
        <span class="stat-label">Total Items</span>
        <span class="stat-value">{{ inventoryStats.totalItems }}</span>
      </div>
      <div class="stat-card">
        <span class="stat-label">Categories</span>
        <span class="stat-value">{{ inventoryStats.categories }}</span>
      </div>
      <div class="stat-card" :class="{ 'warning-bg': inventoryStats.lowStockCount > 0 }">
        <span class="stat-label">Low Stock Alerts</span>
        <span class="stat-value" :class="{ 'warning-text': inventoryStats.lowStockCount > 0 }">
          {{ inventoryStats.lowStockCount }}
        </span>
      </div>
    </div>

    <div class="filters-bar">
      <div class="search-wrap">
        <InputField v-model="searchQuery" placeholder="Search items or location..." />
      </div>
      <div class="dropdowns">
        <select v-model="filterCategory" class="native-dropdown">
          <option v-for="cat in categories" :key="cat" :value="cat">{{ cat }}</option>
        </select>
        <select v-model="filterStatus" class="native-dropdown">
          <option value="All">All Statuses</option>
          <option value="Low Stock">Low Stock</option>
        </select>
      </div>
    </div>

    <div class="table-container">
      <table class="data-table">
        <thead>
          <tr>
            <th>Item Name</th>
            <th>Category</th>
            <th>Location</th>
            <th class="stock-col">Stock Level</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in filteredItems" :key="item.id">
            <td class="fw-bold">{{ item.name }}</td>
            <td>
              <span class="cat-badge">{{ item.category }}</span>
            </td>
            <td>{{ item.location }}</td>
            <td>
              <div class="stock-visual">
                <div class="stock-info">
                  <strong>{{ item.quantity }}</strong> {{ item.unit }}
                  <span class="threshold-sub">Min: {{ item.minThreshold }}</span>
                </div>
                <div class="progress-bar">
                  <div
                    class="fill"
                    :style="{
                      width: getStockWidth(item),
                      background:
                        getStockStatus(item).color === '#d1fae5'
                          ? '#10b981'
                          : getStockStatus(item).color === '#fee2e2'
                            ? '#ef4444'
                            : '#f59e0b',
                    }"
                  ></div>
                </div>
              </div>
            </td>
            <td>
              <Capsules
                :label="getStockStatus(item).label"
                :color="getStockStatus(item).color"
                size="sm"
              />
            </td>
            <td>
              <Button
                title="Adjust"
                size="small"
                color="white"
                :onClick="() => openEditModal(item)"
              />
            </td>
          </tr>
        </tbody>
      </table>
      <div v-if="filteredItems.length === 0" class="empty-state">
        No items found matching criteria.
      </div>
    </div>

    <StockAdjustmentModal
      :isOpen="showEditModal"
      :item="editingItem"
      @close="showEditModal = false"
      @save="handleSaveStock"
    />
  </div>
</template>

<style scoped>
.inventory-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  h1.desktop-only {
    margin: 0;
    font-size: 1.8rem;
    color: var(--text-primary);
  }
}

.mobile-header-title {
  display: none;
}

@media (width <= 768px) {
  .page-header h1.desktop-only {
    display: none;
  }

  .mobile-header-title {
    display: block;
    font-size: 1.25rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
  }
}

.alert-banner {
  background: var(--color-warning-weak);
  border: 1px solid var(--color-warning-border-strong);
  color: var(--color-warning-strong);
  padding: 12px 16px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  gap: 12px;
}

.link-text {
  text-decoration: underline;
  cursor: pointer;
  font-weight: 600;
  margin-left: 4px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

.stat-card {
  background: #fff;
  padding: 20px;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  gap: 8px;

  &.warning-bg {
    border-color: var(--color-danger-border-strong);
    background: var(--color-danger-weak);
  }
}

.stat-label {
  font-size: 0.85rem;
  color: var(--color-neutral-text-soft);
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 2rem;
  font-weight: 700;
  color: var(--text-primary);

  &.warning-text {
    color: var(--color-danger);
  }
}

.filters-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-wrap {
  width: 320px;
}

.dropdowns {
  display: flex;
  gap: 12px;
}

.native-dropdown {
  padding: 10px 16px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  background: #fff;
  font-size: 0.95rem;
  min-width: 140px;
}

.table-container {
  background: #fff;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  overflow: hidden;
}

.data-table {
  width: 100%;
  border-collapse: collapse;

  th {
    text-align: left;
    padding: 16px;
    background: var(--color-neutral-surface);
    color: var(--color-neutral-text-soft);
    font-weight: 600;
    font-size: 0.85rem;
    text-transform: uppercase;
    border-bottom: 1px solid var(--border-color);
  }

  td {
    padding: 16px;
    border-bottom: 1px solid var(--border-color);
    color: var(--text-primary);
    font-size: 0.95rem;
    vertical-align: middle;
  }

  tr:last-child td {
    border-bottom: none;
  }
}

.fw-bold {
  font-weight: 600;
}

.cat-badge {
  background: var(--color-neutral-weak);
  color: var(--text-primary);
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 0.8rem;
  font-weight: 500;
}

.stock-col {
  width: 250px;
}

.stock-visual {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.stock-info {
  display: flex;
  justify-content: space-between;
  font-size: 0.9rem;
}

.threshold-sub {
  font-size: 0.75rem;
  color: var(--color-neutral-text-soft);
}

.progress-bar {
  height: 6px;
  background: #e5e7eb;
  border-radius: 3px;
  overflow: hidden;
  width: 100%;
}

.fill {
  height: 100%;
  border-radius: 3px;
  transition: width 0.3s ease;
}

.empty-state {
  padding: 40px;
  text-align: center;
  color: var(--color-neutral-text-soft);
}

</style>
