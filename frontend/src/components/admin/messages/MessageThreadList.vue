<script setup lang="ts">
import type { IThread } from '../../../stores/mockMessages'

defineProps<{
  currentUser: any // eslint-disable-line @typescript-eslint/no-explicit-any
  channels: IThread[]
  inquiries: IThread[]
  activeThreadId: string
}>()

defineEmits<{
  select: [id: string]
  create: []
}>()
</script>

<template>
  <div class="sidebar">
    <div class="sidebar-header">
      <h2>AdoptionOS</h2>
      <span class="user-badge">{{ currentUser.name }}</span>
    </div>

    <div class="channels-section">
      
      <div class="section-title">
        <span>VOLUNTEER CHANNELS</span>
        <button
          v-if="currentUser.role === 'admin'"
          class="add-btn"
          @click="$emit('create')"
          title="Create Channel"
        >
          +
        </button>
      </div>

      <div class="channel-list">
        <button
          v-for="thread in channels"
          :key="thread.id"
          class="channel-btn"
          :class="{ active: activeThreadId === thread.id }"
          @click="$emit('select', thread.id)"
        >
          <span class="channel-icon">{{ thread.icon }}</span>
          <span class="channel-name">{{ thread.name }}</span>
          <span v-if="thread.isPrivate" class="lock-icon">🔒</span>
        </button>
      </div>

      <div class="section-title mt-6">
        <span>ADOPTER INQUIRIES</span>
      </div>
      <div class="channel-list">
        <button
          v-for="thread in inquiries"
          :key="thread.id"
          class="channel-btn"
          :class="{ active: activeThreadId === thread.id }"
          @click="$emit('select', thread.id)"
        >
          <span class="channel-icon">{{ thread.icon }}</span>
          <div class="inquiry-meta">
            <span class="channel-name">{{ thread.name }}</span>
            <span class="method-badge" :class="thread.contactMethod">
              {{ thread.contactMethod?.toUpperCase() }}
            </span>
          </div>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.sidebar {
  width: 260px;
  background: #f8fafc; 
  border-right: 1px solid #e2e8f0;
  display: flex;
  flex-direction: column;
}

.sidebar-header {
  padding: 16px;
  border-bottom: 1px solid #e2e8f0;

  h2 {
    margin: 0;
    font-size: 1.1rem;
    color: var(--text-primary);
  }
}

.user-badge {
  font-size: 0.8rem;
  color: hsl(from var(--color-neutral) h s 50%);
  display: block;
  margin-top: 4px;
}

.channels-section {
  padding: 16px 8px;
  flex: 1;
  overflow-y: auto;
}

.section-title {
  padding: 0 8px 8px 8px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.75rem;
  font-weight: 700;
  color: hsl(from var(--color-neutral) h s 50%);

  .add-btn {
    background: none;
    border: none;
    color: hsl(from var(--color-neutral) h s 50%);
    cursor: pointer;
    font-size: 1.2rem;
    line-height: 1;
    &:hover {
      color: var(--text-primary);
    }
  }
}

.channel-list {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.mt-6 {
  margin-top: 24px;
}

.channel-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: none;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  color: #64748b;
  width: 100%;
  text-align: left;
  font-size: 0.95rem;

  &:hover {
    background: hsl(from var(--color-neutral) h s 90%);
    color: var(--text-primary);
  }

  &.active {
    background: hsl(from var(--color-secondary) h s 95%);
    color: var(--color-secondary);
    font-weight: 500;
  }
}

.inquiry-meta {
  display: flex;
  flex-direction: column;
  line-height: 1.2;
}

.method-badge {
  font-size: 0.65rem;
  font-weight: 700;

  &.sms {
    color: var(--color-primary);
  }
  &.email {
    color: var(--color-warning);
  }
}

.lock-icon {
  margin-left: auto;
  font-size: 0.7rem;
}
</style>
