<script setup lang="ts">
import { ref } from 'vue'
import {
  mockCampaigns,
  mockNewsletters,
  mockHappyTails,
  type ICampaign,
  type INewsletter,
  type IHappyTail,
} from '../../stores/mockMarketing'
import { Capsules, Button, InputField } from '../../components/common/ui'

const activeTab = ref<'campaigns' | 'newsletters' | 'stories'>('campaigns')

// Helpers
const getProgressColor = (progress: number) => {
  if (progress >= 100) return '#10b981' // Green
  if (progress >= 50) return '#3b82f6' // Blue
  return '#f59e0b' // Orange
}

const getStatusColor = (status: string) => {
  switch (status) {
    case 'active':
    case 'published':
    case 'sent':
      return '#d1fae5'
    case 'draft':
    case 'pending':
      return '#f3f4f6'
    case 'completed':
      return '#e0e7ff'
    case 'scheduled':
      return '#fef3c7'
    case 'rejected':
      return '#fee2e2'
    default:
      return '#f3f4f6'
  }
}

// Actions
const createCampaign = () => {
  const name = prompt('Campaign Name:')
  if (name) {
    mockCampaigns.value.push({
      id: `cp-${Date.now()}`,
      name,
      status: 'draft',
      startDate: new Date().toISOString().split('T')[0],
      endDate: '',
      goal: 'TBD',
      progress: 0,
      metric: 'generic',
    })
  }
}

const approveStory = (story: IHappyTail) => {
  story.status = 'published'
  alert(`Published story for ${story.petName}!`)
}
</script>

<template>
  <div class="marketing-page">
    <!-- HEADER -->
    <div class="page-header">
      <h1>Marketing Center</h1>
      <div class="tabs">
        <button
          class="tab-btn"
          :class="{ active: activeTab === 'campaigns' }"
          @click="activeTab = 'campaigns'"
        >
          📣 Campaigns
        </button>
        <button
          class="tab-btn"
          :class="{ active: activeTab === 'newsletters' }"
          @click="activeTab = 'newsletters'"
        >
          📧 Newsletters
        </button>
        <button
          class="tab-btn"
          :class="{ active: activeTab === 'stories' }"
          @click="activeTab = 'stories'"
        >
          🐾 Happy Tails
        </button>
      </div>
    </div>

    <!-- CAMPAIGNS TAB -->
    <div v-if="activeTab === 'campaigns'" class="tab-content">
      <div class="actions-row">
        <h2>Active Campaigns</h2>
        <Button title="+ New Campaign" color="purple" :onClick="createCampaign" />
      </div>

      <div class="campaigns-grid">
        <div v-for="camp in mockCampaigns" :key="camp.id" class="camp-card">
          <div class="camp-header">
            <h3>{{ camp.name }}</h3>
            <Capsules :label="camp.status" :color="getStatusColor(camp.status)" size="sm" />
          </div>

          <div class="camp-dates">{{ camp.startDate }} — {{ camp.endDate || 'Ongoing' }}</div>

          <div class="camp-stats">
            <div class="stat-row">
              <span>Goal: {{ camp.goal }}</span>
              <strong>{{ camp.progress }}%</strong>
            </div>
            <div class="progress-bar">
              <div
                class="fill"
                :style="{ width: camp.progress + '%', background: getProgressColor(camp.progress) }"
              ></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- NEWSLETTERS TAB -->
    <div v-if="activeTab === 'newsletters'" class="tab-content">
      <div class="actions-row">
        <h2>Email Blasts</h2>
        <Button title="Compose New" color="black" />
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

    <!-- HAPPY TAILS TAB -->
    <div v-if="activeTab === 'stories'" class="tab-content">
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
              color="black"
              :onClick="() => approveStory(story)"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.marketing-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.page-header {
  display: flex;
  flex-direction: column;
  gap: 16px;
  h1 {
    margin: 0;
    font-size: 1.8rem;
    color: var(--font-color-dark);
  }
}

.tabs {
  display: flex;
  gap: 24px;
  border-bottom: 1px solid #e5e7eb;
}

.tab-btn {
  background: none;
  border: none;
  padding: 12px 4px;
  font-size: 1rem;
  color: var(--font-color-medium);
  cursor: pointer;
  border-bottom: 2px solid transparent;
  font-weight: 500;

  &.active {
    color: var(--font-color-dark);
    border-bottom-color: var(--purple);
    font-weight: 700;
  }
}

.actions-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  h2 {
    margin: 0;
    font-size: 1.2rem;
  }
}

/* CAMPAIGNS */
.campaigns-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.camp-card {
  background: white;
  padding: 20px;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
}

.camp-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 8px;
  h3 {
    margin: 0;
    font-size: 1.1rem;
  }
}

.camp-dates {
  font-size: 0.9rem;
  color: var(--font-color-medium);
  margin-bottom: 20px;
}

.progress-bar {
  height: 8px;
  background: #f3f4f6;
  border-radius: 4px;
  overflow: hidden;
  margin-top: 8px;
}

.fill {
  height: 100%;
  transition: width 0.5s ease;
}

.stat-row {
  display: flex;
  justify-content: space-between;
  font-size: 0.9rem;
}

/* NEWSLETTERS */
.list-container {
  background: white;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  overflow: hidden;
}

.list-item {
  display: flex;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid #f3f4f6;
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
  color: var(--font-color-medium);
}

.item-stats {
  text-align: right;
  margin-right: 16px;
}
.stat-lbl {
  display: block;
  font-size: 0.75rem;
  color: var(--font-color-medium);
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
  color: var(--font-color-medium);
}

/* STORIES */
.stories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.story-card {
  background: white;
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
  color: var(--purple);
}
.date {
  font-size: 0.8rem;
  color: var(--font-color-medium);
}
.adopter {
  font-size: 0.9rem;
  margin: 0 0 12px 0;
  color: var(--font-color-dark);
}

.story-body {
  font-style: italic;
  color: #4b5563;
  line-height: 1.5;
  margin-bottom: 20px;
  flex: 1;
}

.story-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-top: 1px solid #f3f4f6;
  padding-top: 16px;
}
</style>
