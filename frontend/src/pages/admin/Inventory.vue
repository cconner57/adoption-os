<script setup lang="ts">
import { ref, computed } from 'vue'
import { mockInventory, inventoryStats, type IInventoryItem } from '../../stores/mockInventory'
import { Capsules, InputField, InputSelectGroup, Button } from '../../components/common/ui'

const searchQuery = ref('')
const filterCategory = ref<'All' | string>('All')
const filterStatus = ref<'All' | 'Low Stock'>('All')

// Computed
const categories = computed(() => ['All', ...new Set(mockInventory.value.map((i) => i.category))])

const filteredItems = computed(() => {
  return mockInventory.value
    .filter((item) => {
      // 1. Search
      const searchMatch =
        !searchQuery.value ||
        item.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        item.location.toLowerCase().includes(searchQuery.value.toLowerCase())

      // 2. Category
      const catMatch = filterCategory.value === 'All' || item.category === filterCategory.value

      // 3. Status
      const isLow = item.quantity <= item.minThreshold
      const statusMatch =
        filterStatus.value === 'All' || (filterStatus.value === 'Low Stock' && isLow)

      return searchMatch && catMatch && statusMatch
    })
    .sort((a, b) => {
      // Sort logic: Critical items first
      const aRatio = a.quantity / a.minThreshold
      const bRatio = b.quantity / b.minThreshold
      return aRatio - bRatio
    })
})

// Helpers
const getStockStatus = (item: IInventoryItem) => {
  if (item.quantity === 0) return { label: 'Out of Stock', color: '#fee2e2' } // Red
  if (item.quantity <= item.minThreshold) return { label: 'Low Stock', color: '#fef3c7' } // Yellow
  return { label: 'In Stock', color: '#d1fae5' } // Green
}

const getStockWidth = (item: IInventoryItem) => {
  // Visual percentage for progress bar, capped at 100%
  // Assuming 2x threshold is "full" enough for visualization
  const max = item.minThreshold * 2
  return Math.min((item.quantity / max) * 100, 100) + '%'
}

// Actions
const showEditModal = ref(false)
const editingItem = ref<IInventoryItem | null>(null)
const tempQuantity = ref(0) // For adjustment

const openEditModal = (item: IInventoryItem) => {
  editingItem.value = item
  tempQuantity.value = item.quantity
  showEditModal.value = true
}

const saveStock = () => {
  if (editingItem.value) {
    editingItem.value.quantity = tempQuantity.value
    editingItem.value.lastUpdated = new Date().toISOString().split('T')[0]
    showEditModal.value = false
    editingItem.value = null
  }
}

// Add Item Mock
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
</script>

<template>
  <div class="inventory-page">
    <!-- HEADER -->
    <div class="page-header">
      <h1>Inventory Management</h1>
      <Button title="+ Add Item" color="black" :onClick="addItem" />
    </div>

    <!-- ALERTS -->
    <div v-if="inventoryStats.lowStockCount > 0" class="alert-banner">
      <span class="alert-icon">⚠️</span>
      <span class="alert-text">
        <strong>Action Required:</strong> {{ inventoryStats.lowStockCount }} items are below stock
        threshold.
        <span class="link-text" @click="filterStatus = 'Low Stock'">View Items</span>
      </span>
    </div>

    <!-- STATS -->
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

    <!-- FILTERS -->
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

    <!-- TABLE -->
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

    <!-- EDIT ADJUST MODAL -->
    <div
      v-if="showEditModal && editingItem"
      class="modal-overlay"
      @click.self="showEditModal = false"
    >
      <div class="modal-card">
        <h3>Adjust Stock: {{ editingItem.name }}</h3>
        <p class="subtitle">Update current quantity on hand.</p>

        <div class="stock-adjuster">
          <button class="adjust-btn" @click="tempQuantity = Math.max(0, tempQuantity - 1)">
            -
          </button>
          <input v-model.number="tempQuantity" type="number" class="qty-input" />
          <button class="adjust-btn" @click="tempQuantity++">+</button>
        </div>

        <p class="threshold-info">
          Minimum Threshold: {{ editingItem.minThreshold }} {{ editingItem.unit }}
        </p>

        <div class="modal-actions">
          <Button title="Cancel" color="white" :onClick="() => (showEditModal = false)" />
          <Button title="Save Changes" color="black" :onClick="saveStock" />
        </div>
      </div>
    </div>
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
  h1 {
    margin: 0;
    font-size: 1.8rem;
    color: var(--font-color-dark);
  }
}

.alert-banner {
  background: #fff7ed;
  border: 1px solid #fed7aa;
  color: #c2410c;
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

/* STATS */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

.stat-card {
  background: white;
  padding: 20px;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  gap: 8px;

  &.warning-bg {
    border-color: #fca5a5;
    background: #fef2f2;
  }
}

.stat-label {
  font-size: 0.85rem;
  color: var(--font-color-medium);
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 2rem;
  font-weight: 700;
  color: var(--font-color-dark);
  &.warning-text {
    color: #b91c1c;
  }
}

/* FILTERS */
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
  background: white;
  font-size: 0.95rem;
  min-width: 140px;
}

/* TABLE */
.table-container {
  background: white;
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
    background: #f9fafb;
    color: var(--font-color-medium);
    font-weight: 600;
    font-size: 0.85rem;
    text-transform: uppercase;
    border-bottom: 1px solid #e5e7eb;
  }

  td {
    padding: 16px;
    border-bottom: 1px solid #f3f4f6;
    color: var(--font-color-dark);
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
  background: #f3f4f6;
  color: var(--font-color-dark);
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
  color: var(--font-color-medium);
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
  color: var(--font-color-medium);
}

/* MODAL */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
}

.modal-card {
  background: white;
  padding: 24px;
  border-radius: 12px;
  width: 100%;
  max-width: 400px;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1);
  text-align: center;

  h3 {
    margin: 0 0 8px 0;
  }
  .subtitle {
    margin: 0 0 24px 0;
    color: var(--font-color-medium);
  }
}

.stock-adjuster {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
  margin-bottom: 16px;
}

.adjust-btn {
  width: 40px;
  height: 40px;
  border-radius: 8px;
  border: 1px solid #e5e7eb;
  background: white;
  font-size: 1.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  &:hover {
    background: #f9fafb;
  }
}

.qty-input {
  width: 80px;
  text-align: center;
  font-size: 1.5rem;
  font-weight: 700;
  border: none;
  border-bottom: 2px solid var(--purple);
  outline: none;
}

.threshold-info {
  font-size: 0.85rem;
  color: var(--font-color-medium);
  margin-bottom: 24px;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>
