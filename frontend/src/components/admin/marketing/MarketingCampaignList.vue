<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { computed, onMounted, ref } from 'vue'

import { useMarketingStore } from '../../../stores/marketing'
import type { ICampaign } from '../../../stores/mockMarketing'
import { Button, Capsules } from '../../common/ui'
import CampaignEditor from './CampaignEditor.vue'

const emit = defineEmits<{
  'select': [campaign: ICampaign]
}>()

const store = useMarketingStore()
const { campaigns } = storeToRefs(store)

const calculateProgress = (camp: ICampaign) => {
  const goalNum = Number(camp.goal)
  if (!goalNum) return 0

  // 1. Use explicit progress from DB if available and seemingly valid
  if (typeof camp.progress === 'number' && camp.progress > 0) {
      // Assuming 'progress' in DB is the raw value (e.g. 150 entries), calculate percentage
      return Math.min(100, Math.round((camp.progress / goalNum) * 100))
  }

  // 2. Fallback to calculating from participants if array is present
  if (camp.participants && camp.participants.length > 0) {
      let current = 0
      if (camp.metric === 'dollars') {
          current = camp.participants
              .filter(p => p.paid)
              .reduce((sum, p) => {
                  const amt = parseFloat((p.amount || '0').toString().replace('$', ''))
                  return sum + (isNaN(amt) ? 0 : amt)
              }, 0)
      } else {
          current = camp.participants.length
      }
      return Math.min(100, Math.round((current / goalNum) * 100))
  }

  return 0
}

const getGoalText = (camp: ICampaign) => {
    if (camp.metric === 'dollars') {
        const goalNum = Number(camp.goal)
        return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD', maximumFractionDigits: 0 }).format(goalNum)
    }
    return `${camp.goal}`
}
const getProgressColor = (progress: number) => {
  if (progress >= 100) return '#10b981'
  if (progress >= 50) return '#3b82f6'
  return '#f59e0b'
}

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

// Drawer Logic
const isEditorOpen = ref(false)
const newCampaign = ref<ICampaign | null>(null)

const startCreate = () => {
  newCampaign.value = null // Resets editor to "add" mode
  isEditorOpen.value = true
}

const handleSaveNew = async (campaignData: ICampaign) => {
  // If editing an existing campaign (not likely here, but for safety)
  if (campaignData.id && !campaignData.id.startsWith('new-')) {
      await store.updateCampaign(campaignData)
  } else {
      // Create new
      const newId = `cp-${Date.now()}`
      campaigns.value.push({
          ...campaignData,
          id: newId,
          progress: 0,
          metric: 'entries',
          type: 'standard', // Default, or infer from data if added to editor
      })
  }
  isEditorOpen.value = false
}

onMounted(() => {
  store.fetchCampaigns()
})

const activeCampaigns = computed(() => {
  return campaigns.value.filter(c => ['active', 'published', 'scheduled'].includes(c.status))
})

const draftCampaigns = computed(() => {
  return campaigns.value.filter(c => ['draft', 'pending'].includes(c.status))
})

const pastCampaigns = computed(() => {
  return campaigns.value.filter(c => ['completed', 'archived', 'rejected'].includes(c.status))
})

const formatDate = (dateString: string) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  if (isNaN(date.getTime())) return dateString
  return new Intl.DateTimeFormat('en-US', { month: 'short', day: 'numeric', year: 'numeric' }).format(date)
}
</script>

<template>
  <div class="marketing-campaign-list">
    <div class="actions-row">
      <h2>Overview</h2>
      <Button title="+ New Campaign" color="purple" :onClick="startCreate" />
    </div>

    <div class="campaign-section">
      <h3>Active Campaigns</h3>
      <div class="campaigns-grid" v-if="activeCampaigns.length">
        <div
          v-for="camp in activeCampaigns"
          :key="camp.id"
          class="camp-card"
          @click="emit('select', camp)"
        >
          <div class="camp-header">
            <h3>{{ camp.name }}</h3>
            <Capsules :label="camp.status" :color="getStatusColor(camp.status)" size="sm" />
          </div>

          <div class="camp-dates">{{ formatDate(camp.startDate) }} — {{ camp.endDate ? formatDate(camp.endDate) : 'Ongoing' }}</div>

          <div class="camp-stats">
            <div class="stat-row">
              <span>Goal: {{ getGoalText(camp) }}</span>
              <strong>{{ calculateProgress(camp) }}%</strong>
            </div>
            <div class="progress-bar">
              <div
                class="fill"
                :style="{ width: calculateProgress(camp) + '%', background: getProgressColor(calculateProgress(camp)) }"
              ></div>
            </div>
          </div>

          <div class="camp-type-badge" v-if="camp.type === 'raffle'">
              🎟️ Raffle
          </div>
        </div>
      </div>
      <p v-else class="empty-text">No active campaigns.</p>
    </div>

    <div class="campaign-section">
      <h3>Drafts</h3>
      <div class="campaigns-grid" v-if="draftCampaigns.length">
        <div
          v-for="camp in draftCampaigns"
          :key="camp.id"
          class="camp-card draft"
          @click="emit('select', camp)"
        >
          <div class="camp-header">
            <h3>{{ camp.name }}</h3>
            <Capsules :label="camp.status" :color="getStatusColor(camp.status)" size="sm" />
          </div>
          <div class="camp-dates">{{ formatDate(camp.startDate) }}</div>

          <div class="camp-stats">
            <div class="stat-row">
              <span>Goal: {{ getGoalText(camp) }}</span>
              <strong>{{ calculateProgress(camp) }}%</strong>
            </div>
            <div class="progress-bar">
              <div
                class="fill"
                :style="{ width: calculateProgress(camp) + '%', background: getProgressColor(calculateProgress(camp)) }"
              ></div>
            </div>
          </div>

          <div class="camp-type-badge" v-if="camp.type === 'raffle'">
              🎟️ Raffle
          </div>
        </div>
      </div>
      <p v-else class="empty-text">No drafts.</p>
    </div>

    <div class="campaign-section">
      <h3>Past Campaigns</h3>
      <div class="campaigns-grid" v-if="pastCampaigns.length">
        <div
          v-for="camp in pastCampaigns"
          :key="camp.id"
          class="camp-card past"
          @click="emit('select', camp)"
        >
          <div class="camp-header">
            <h3>{{ camp.name }}</h3>
            <Capsules :label="camp.status" :color="getStatusColor(camp.status)" size="sm" />
          </div>
          <div class="camp-dates">{{ formatDate(camp.startDate) }} — {{ formatDate(camp.endDate) }}</div>
          <div class="camp-type-badge" v-if="camp.type === 'raffle'">
              🎟️ Raffle
          </div>
          <div class="camp-stats">
            <div class="stat-row">
              <span>Goal: {{ getGoalText(camp) }}</span>
              <strong>{{ calculateProgress(camp) }}%</strong>
            </div>
            <div class="progress-bar">
              <div
                class="fill"
                :style="{ width: calculateProgress(camp) + '%', background: getProgressColor(calculateProgress(camp)) }"
              ></div>
            </div>
          </div>
        </div>
      </div>
      <p v-else class="empty-text">No past campaigns.</p>
    </div>

    <!-- Create Campaign Drawer -->
    <CampaignEditor
        :is-open="isEditorOpen"
        :campaign="newCampaign"
        @close="isEditorOpen = false"
        @save="handleSaveNew"
    />
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

.campaign-section {
  margin-top: 24px;

  h3 {
    font-size: 1rem;
    color: var(--text-secondary);
    margin-bottom: 12px;
    font-weight: 600;
  }
}

.empty-text {
  font-style: italic;
  color: hsl(from var(--color-neutral) h s 60%);
  font-size: 0.9rem;
}

.campaigns-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.camp-card {
  background: #fff;
  padding: 24px;
  border-radius: 16px;
  border: 1px solid var(--border-color);
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
  display: flex;
  flex-direction: column;
  gap: 16px;

  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgb(0 0 0 / 5%);
  }
}

.camp-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 0;

  h3 {
    margin: 0;
    font-size: 1.1rem;
  }
}

.camp-dates {
  font-size: 0.9rem;
  color: hsl(from var(--color-neutral) h s 50%);
  margin-bottom: 0;
}

.progress-bar {
  height: 8px;
  background: hsl(from var(--color-neutral) h s 95%);
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

.camp-type-badge {
    margin-top: 12px;
    font-size: 0.8rem;
    color: var(--color-tertiary);
    font-weight: 600;
    display: flex;
    align-items: center;
    gap: 6px;
    background: hsl(from var(--color-tertiary) h s 96%);
    width: fit-content;
    padding: 2px 8px;
    border-radius: 8px;
}
</style>
