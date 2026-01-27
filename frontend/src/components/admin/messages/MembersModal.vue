<script setup lang="ts">
import { Button } from '../../common/ui'

defineProps<{
  isOpen: boolean
  activeThread: any // eslint-disable-line @typescript-eslint/no-explicit-any
  users: any[] // eslint-disable-line @typescript-eslint/no-explicit-any
  currentUser: any // eslint-disable-line @typescript-eslint/no-explicit-any
}>()

defineEmits<{
  close: []
  addMember: []
  removeMember: [userId: string]
}>()
</script>

<template>
  <div v-if="isOpen" class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-card">
      <h3>Manage Members - #{{ activeThread?.name }}</h3>
      <p class="subtitle">Only admins can add or remove members.</p>

      <div class="members-list">
        <div v-for="user in users" :key="user.id" class="member-row">
          <div class="m-info">
            <span class="m-avatar">{{ user.avatar }}</span>
            <span class="m-name">{{ user.name }}</span>
            <span v-if="user.role === 'admin'" class="admin-badge">Admin</span>
          </div>
          <button
            v-if="user.id !== currentUser.id"
            class="remove-btn"
            @click="$emit('removeMember', user.id)"
          >
            Remove
          </button>
        </div>
      </div>

      <div class="modal-actions">
        <Button title="Add People" color="purple" :onClick="() => $emit('addMember')" />
        <Button title="Done" color="white" :onClick="() => $emit('close')" />
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

  .subtitle {
    color: var(--color-neutral-text-soft);
    margin-bottom: 20px;
    font-size: 0.9rem;
  }
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 24px;
}

.members-list {
  max-height: 300px;
  overflow-y: auto;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
}

.member-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 12px;
  border-bottom: 1px solid #f1f5f9;

  &:last-child {
    border-bottom: none;
  }
}

.m-info {
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 500;
}

.m-avatar {
  font-size: 1.2rem;
}

.admin-badge {
  background: var(--color-secondary-weak);
  color: var(--color-secondary);
  font-size: 0.7rem;
  padding: 2px 6px;
  border-radius: 4px;
  text-transform: uppercase;
  font-weight: 700;
}

.remove-btn {
  color: var(--color-danger);
  background: none;
  border: none;
  font-size: 0.85rem;
  cursor: pointer;

  &:hover {
    text-decoration: underline;
  }
}
</style>
