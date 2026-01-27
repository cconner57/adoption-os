<script setup lang="ts">
import { ref, watch } from 'vue'

import type { IInventoryItem } from '../../../stores/mockInventory'
import { Button } from '../../common/ui'

const props = defineProps<{
  isOpen: boolean
  item: IInventoryItem | null
}>()

const emit = defineEmits<{
  close: []
  save: [item: IInventoryItem, quantity: number]
}>()

const tempQuantity = ref(0)

watch(
  () => props.item,
  (newItem) => {
    if (newItem) {
      tempQuantity.value = newItem.quantity
    }
  },
  { immediate: true },
)

const save = () => {
  if (props.item) {
    emit('save', props.item, tempQuantity.value)
  }
}
</script>

<template>
  <div v-if="isOpen && item" class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-card">
      <h3>Adjust Stock: {{ item.name }}</h3>
      <p class="subtitle">Update current quantity on hand.</p>

      <div class="stock-adjuster">
        <button class="adjust-btn" @click="tempQuantity = Math.max(0, tempQuantity - 1)">-</button>
        <input v-model.number="tempQuantity" type="number" class="qty-input" />
        <button class="adjust-btn" @click="tempQuantity++">+</button>
      </div>

      <p class="threshold-info">Minimum Threshold: {{ item.minThreshold }} {{ item.unit }}</p>

      <div class="modal-actions">
        <Button title="Cancel" color="white" :onClick="() => $emit('close')" />
        <Button title="Save Changes" color="white" :onClick="save" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgb(0 0 0 / 50%);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
}

.modal-card {
  background: #fff;
  padding: 24px;
  border-radius: 12px;
  width: 100%;
  max-width: 400px;
  box-shadow: 0 20px 25px -5px rgb(0 0 0 / 10%);
  text-align: center;

  h3 {
    margin: 0 0 8px;
  }
}

.subtitle {
  margin: 0 0 24px;
  color: var(--color-neutral-text-soft);
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
  background: #fff;
  font-size: 1.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.adjust-btn:hover {
  background: var(--color-neutral-surface);
}

.qty-input {
  width: 80px;
  text-align: center;
  font-size: 1.5rem;
  font-weight: 700;
  border: none;
  border-bottom: 2px solid var(--color-secondary);
  outline: none;
}

.threshold-info {
  font-size: 0.85rem;
  color: var(--color-neutral-text-soft);
  margin-bottom: 24px;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>
