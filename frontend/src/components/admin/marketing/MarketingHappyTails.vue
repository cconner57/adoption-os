<script setup lang="ts">
import type { IHappyTail } from '../../../stores/mockMarketing'
import { mockHappyTails } from '../../../stores/mockMarketing'
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

const approveStory = (story: IHappyTail) => {
  story.status = 'published'
  alert(`Published story for ${story.petName}!`)
}
</script>

<template>
  <div class="marketing-happytails">
    <div class="actions-row">
      <h2>Submitted Stories</h2>
    </div>

    <div class="stories-grid">
      <div v-for="story in mockHappyTails" :key="story.id" class="story-card">
        <div class="story-header">
          <span class="pet-name">🐾 {{ story.petName }}</span>
          <span class="date">{{ story.date }}</span>
        </div>
        <p class="adopter">From: {{ story.adopterName }}</p>

        <div class="story-body">"{{ story.story }}"</div>

        <div class="story-actions">
          <Capsules :label="story.status" :color="getStatusColor(story.status)" size="sm" />
          <Button
            v-if="story.status === 'pending'"
            title="Approve"
            size="small"
            color="white"
            :onClick="() => approveStory(story)"
          />
        </div>
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

.stories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.story-card {
  background: #fff;
  padding: 20px;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
}

.story-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}

.pet-name {
  font-weight: 700;
  font-size: 1.1rem;
  color: var(--color-secondary);
}

.date {
  font-size: 0.8rem;
  color: hsl(from var(--color-neutral) h s 50%);
}

.adopter {
  font-size: 0.9rem;
  margin: 0 0 12px;
  color: var(--text-primary);
}

.story-body {
  font-style: italic;
  color: hsl(from var(--color-neutral) h s 40%);
  margin-bottom: 16px;
  line-height: 1.5;
  flex: 1;
}

.story-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: auto;
  padding-top: 12px;
  border-top: 1px solid var(--border-color);
}
</style>
