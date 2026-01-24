<script setup lang="ts">
import { ref } from 'vue'

import { Button, InputField } from '../../common/ui'

defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
  create: [name: string, desc: string]
}>()

const name = ref('')
const desc = ref('')

const create = () => {
  if (name.value) {
    emit('create', name.value, desc.value)
    name.value = ''
    desc.value = ''
  }
}
</script>

<template>
  <div v-if="isOpen" class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-card">
      <h3>Create New Channel</h3>
      <div class="form-group">
        <label>Channel Name</label>
        <InputField v-model="name" placeholder="e.g. shelter-events" />
      </div>
      <div class="form-group">
        <label>Description</label>
        <InputField v-model="desc" placeholder="e.g. Planning upcoming events" />
      </div>
      <div class="modal-actions">
        <Button title="Cancel" color="white" :onClick="() => $emit('close')" />
        <Button title="Create Channel" color="white" :onClick="create" />
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
  max-width: 450px;
  box-shadow:
    0 20px 25px -5px rgb(0 0 0 / 10%),
    0 10px 10px -5px rgb(0 0 0 / 4%);

  h3 {
    margin-top: 0;
  }
}

.form-group {
  margin-bottom: 16px;

  label {
    display: block;
    margin-bottom: 6px;
    font-weight: 500;
    font-size: 0.9rem;
  }
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 24px;
}
</style>
