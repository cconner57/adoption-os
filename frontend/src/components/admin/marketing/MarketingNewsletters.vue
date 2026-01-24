<script setup lang="ts">
import { mockNewsletters } from '../../../stores/mockMarketing'
import { Button, Capsules } from '../../common/ui'

const getStatusColor = (status: string) => {
  switch (status) {
    case 'active':
    case 'published':
    case 'sent':
      return 'green'
    case 'draft':
    case 'pending':
      return 'neutral'
    case 'completed':
      return 'blue'
    case 'scheduled':
      return 'warning'
    case 'rejected':
      return 'danger'
    default:
      return 'neutral'
  }
}
</script>

<template>
  <div class="marketing-newsletters">
    <div class="actions-row">
      <h2>Email Blasts</h2>
      <Button title="Compose New" color="white" />
    </div>

    <div class="list-container">
      <div v-for="nl in mockNewsletters" :key="nl.id" class="list-item">
        <div class="item-main">
          <span class="subject">{{ nl.subject }}</span>
          <div class="meta">
            <span v-if="nl.sentDate">Sent: {{ nl.sentDate }}</span>
            <span v-if="nl.recipientCount"
              >• {{ nl.recipientCount.toLocaleString() }} Recipients</span
            >
          </div>
        </div>

        <div class="item-stats" v-if="nl.openRate">
          <span class="stat-lbl">Open Rate</span>
          <span class="stat-val">{{ nl.openRate }}%</span>
        </div>

        <Capsules :label="nl.status" :color="getStatusColor(nl.status)" size="sm" />
        <button class="icon-btn">⋮</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.actions-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;

  h2 {
    margin: 0;
    font-size: 1.2rem;
  }
}

.list-container {
  background: #fff;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  overflow: hidden;
}

.list-item {
  display: flex;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid var(--border-color);
  gap: 16px;

  &:last-child {
    border-bottom: none;
  }
}

.item-main {
  flex: 1;
}

.subject {
  font-weight: 600;
  display: block;
  margin-bottom: 4px;
}

.meta {
  font-size: 0.85rem;
  color: hsl(from var(--color-neutral) h s 50%);
}

.item-stats {
  text-align: right;
  margin-right: 16px;
}

.stat-lbl {
  display: block;
  font-size: 0.75rem;
  color: hsl(from var(--color-neutral) h s 50%);
  text-transform: uppercase;
}

.stat-val {
  font-weight: 700;
  font-size: 1.1rem;
}

.icon-btn {
  background: none;
  border: none;
  font-size: 1.2rem;
  cursor: pointer;
  color: hsl(from var(--color-neutral) h s 50%);
}
</style>
