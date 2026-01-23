<script setup lang="ts">
import { ref, watch } from 'vue'

import { Button } from '../../common/ui'

const props = defineProps<{
  isOpen: boolean
  mockPets: any[] // eslint-disable-line @typescript-eslint/no-explicit-any
}>()

const emit = defineEmits<{
  close: []
  confirm: [petId: string]
}>()

const assignPetId = ref('')

watch(
  () => props.isOpen,
  (val) => {
    if (val) assignPetId.value = ''
  },
)

const confirm = () => {
  if (assignPetId.value) {
    emit('confirm', assignPetId.value)
  }
}
</script>

<template>
  <div v-if="isOpen" class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-card">
      <h3>Assign Pet</h3>
      <div class="select-list">
        <div
          v-for="pet in mockPets"
          :key="pet.id"
          class="select-item"
          @click="assignPetId = pet.id"
          :class="{ selected: assignPetId === pet.id }"
        >
          {{ pet.name }} ({{ pet.breed }})
        </div>
      </div>
      <div class="modal-actions">
        <Button title="Cancel" color="white" :onClick="() => $emit('close')" />
        <Button title="Confirm Link" color="white" :onClick="confirm" />
      </div>
    </div>
  </div>
</template>

<style scoped>
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
  width: 400px;
  max-height: 80vh;
  display: flex;
  flex-direction: column;
}

h3 {
  margin-top: 0;
}

.select-list {
  overflow-y: auto;
  max-height: 300px;
  margin: 16px 0;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
}
.select-item {
  padding: 12px;
  cursor: pointer;
  &:hover {
    background: #f8fafc;
  }
  &.selected {
    background: hsl(from var(--color-secondary) h s 95%);
    color: var(--color-secondary);
    font-weight: 600;
  }
}
.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>
