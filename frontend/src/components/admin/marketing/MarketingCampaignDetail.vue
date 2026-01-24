<script setup lang="ts">
import { computed, ref } from 'vue'

import { useMarketingStore } from '../../../stores/marketing'
import type { ICampaign, IRaffleParticipant } from '../../../stores/mockMarketing'
import { Button, Capsules, SettingsButton, TableColumnToggle } from '../../common/ui'
import CampaignEditor from './CampaignEditor.vue'
import MarketingWinner from './MarketingWinner.vue'
import ParticipantEditor from './ParticipantEditor.vue'

const props = defineProps<{
  campaign: ICampaign
}>()

const store = useMarketingStore()

const visibleColumns = ref({
  ticketNumber: true,
  name: true,
  date: true,
  paymentMethod: true,
  amount: true,
  status: true,
  contact: true,
  comments: true
})

const columnConfig = [
  { key: 'ticketNumber', label: 'Ticket #' },
  { key: 'name', label: 'Name' },
  { key: 'date', label: 'Date' },
  { key: 'paymentMethod', label: 'Payment' },
  { key: 'amount', label: 'Amount' },
  { key: 'status', label: 'Status' },
  { key: 'contact', label: 'Contact' },
  { key: 'comments', label: 'Comments' },
]

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

const formatDate = (dateString: string) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  if (isNaN(date.getTime())) return dateString
  return new Intl.DateTimeFormat('en-US', { month: 'short', day: 'numeric', year: 'numeric' }).format(date)
}

const raisedAmount = computed(() => {
  if (!props.campaign.participants) return '$0'

  const total = props.campaign.participants
    .filter(p => p.paid)
    .reduce((sum, p) => {
        const amt = parseFloat((p.amount || '0').toString().replace('$', ''))
        return sum + (isNaN(amt) ? 0 : amt)
    }, 0)

  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD', maximumFractionDigits: 0 }).format(total)
})

// Drawer Logic
const isEditorOpen = ref(false)
const editingParticipant = ref<IRaffleParticipant | null>(null)

const startEdit = (participant: IRaffleParticipant) => {
  editingParticipant.value = participant
  isEditorOpen.value = true
}

const startAdd = () => {
  editingParticipant.value = null
  isEditorOpen.value = true
}

const handleSaveParticipant = async (updated: IRaffleParticipant) => {
  if (!props.campaign.participants) {
      props.campaign.participants = []
  }

  const campaign = props.campaign

  if (updated.id && !updated.id.startsWith('new-')) {
      const idx = campaign.participants!.findIndex(p => p.id === updated.id)
      if (idx !== -1) {
        campaign.participants![idx] = updated
      }
  } else {
      campaign.participants!.push({
          ...updated,
          id: `p-${Date.now()}`,
          ticketNumber: campaign.participants!.length + 1
      })
  }

  await store.updateCampaign(campaign)
  isEditorOpen.value = false
}

const calculateProgress = (camp: ICampaign) => {
  if (!camp.participants || !camp.goal) return 0

  let current = 0
  if (camp.metric === 'dollars') {
      current = camp.participants
          .filter(p => p.paid)
          .reduce((sum, p) => {
              const amt = parseFloat((p.amount || '0').toString().replace('$', ''))
              return sum + (isNaN(amt) ? 0 : amt)
          }, 0)
  } else {
      // Default to entries
      current = camp.participants.length
  }

  const goalNum = Number(camp.goal)
  if (!goalNum) return 0
  return Math.min(100, Math.round((current / goalNum) * 100))
}

const getGoalText = (camp: ICampaign) => {
    if (camp.metric === 'dollars') {
        const goalNum = Number(camp.goal)
        return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD', maximumFractionDigits: 0 }).format(goalNum)
    }
    return `${camp.goal}`
}
// Campaign Editor Logic
const isCampaignEditorOpen = ref(false)

const handleSaveCampaign = async (updated: ICampaign) => {
  if (props.campaign.id === updated.id) {
    // Direct mutate prop object since it is a reactive object from parent store
    Object.assign(props.campaign, updated)
    await store.updateCampaign(props.campaign)
  }
  isCampaignEditorOpen.value = false
}
</script>

<template>
  <div class="marketing-campaign-detail">
    <div class="detail-header-card">
      <div class="status-row">
          <Capsules :label="campaign.status" :color="getStatusColor(campaign.status)" />
          <span class="dates">{{ formatDate(campaign.startDate) }} — {{ campaign.endDate ? formatDate(campaign.endDate) : 'Ongoing' }}</span>
          <div style="flex: 1"></div>
          <SettingsButton
            title="Edit Campaign Details"
            @click="isCampaignEditorOpen = true"
          />
      </div>

      <div class="stats-overview" v-if="campaign.type === 'raffle'">
            <div class="stat-box">
            <span class="label">Prize</span>
            <span class="value">{{ campaign.prize || '-' }}</span>
          </div>
          <div class="stat-box">
            <span class="label">Ticket Price</span>
            <span class="value">{{ campaign.ticketPrice ? '$' + campaign.ticketPrice : '-' }}</span>
          </div>
          <div class="stat-box">
            <span class="label">Entries</span>
            <span class="value">{{ campaign.participants?.length || 0 }}</span>
          </div>
            <div class="stat-box">
            <span class="label">Raised</span>
            <span class="value">{{ raisedAmount }}</span>
          </div>
            <div class="stat-box">
            <span class="label">Goal</span>
            <span class="value">{{ getGoalText(campaign) }}</span>
          </div>
          <div class="stat-box" style="grid-column: span 5; margin-top: 12px;">
             <span class="label" style="display:flex; justify-content: space-between;">
                <span>Progress</span>
                <span>{{ calculateProgress(campaign) }}%</span>
             </span>
             <div class="progress-bar">
              <div
                class="fill"
                :style="{ width: calculateProgress(campaign) + '%', background: calculateProgress(campaign) >= 100 ? '#10b981' : '#3b82f6' }"
              ></div>
            </div>
          </div>
      </div>
    </div>

    <div class="raffle-management" v-if="campaign.type === 'raffle'">
      <div class="participants-section">
        <div class="section-header">
          <h2>Entries</h2>
          <div class="add-entry-actions">
              <TableColumnToggle v-model="visibleColumns" :columns="columnConfig" />
              <Button title="+ Add Entry" color="purple" :onClick="startAdd" size="large" />
          </div>
        </div>

        <div class="participants-table-container">
          <table class="participants-table">
              <thead>
                <tr>
                  <th v-if="visibleColumns.ticketNumber">Ticket #</th>
                  <th v-if="visibleColumns.name">Name</th>
                  <th v-if="visibleColumns.date">Date</th>
                  <th v-if="visibleColumns.paymentMethod">Payment</th>
                  <th v-if="visibleColumns.amount">Amount</th>
                  <th v-if="visibleColumns.status">Status</th>
                  <th v-if="visibleColumns.contact">Contact</th>
                  <th v-if="visibleColumns.comments">Comments</th>
                  <th></th>
                </tr>
              </thead>
            <tbody>
              <tr v-for="p in campaign.participants" :key="p.id" :class="{ 'is-winner': campaign.winnerId === p.id }">
                <td v-if="visibleColumns.ticketNumber">#{{ p.ticketNumber }}</td>

                <!-- Name -->
                <td v-if="visibleColumns.name">
                    {{ p.name }}
                </td>

                <!-- Date -->
                <td v-if="visibleColumns.date">{{ p.date || '-' }}</td>

                <!-- Payment -->
                <td v-if="visibleColumns.paymentMethod">{{ p.paymentMethod || '-' }}</td>

                <!-- Amount -->
                <td v-if="visibleColumns.amount">{{ p.amount || '-' }}</td>

                <!-- Status -->
                  <td v-if="visibleColumns.status">
                    <span class="paid-badge" v-if="p.paid">PAID</span>
                  </td>

                <!-- Contact -->
                <td v-if="visibleColumns.contact">
                    {{ p.contact }}
                </td>

                <!-- Comments -->
                <td v-if="visibleColumns.comments">
                  <div class="comments-cell" :title="p.comments">{{ p.comments || '-' }}</div>
                </td>

                <!-- Actions -->
                <td>
                  <div class="actions-wrapper">
                        <span v-if="campaign.winnerId === p.id" class="winner-label">WINNER</span>
                      <div class="actions-cell" style="justify-content: flex-end;">
                        <Button
                          title="Edit"
                          color="white"
                          size="small"
                          :onClick="() => startEdit(p)"
                          style="padding: 4px 10px; height: 30px; font-size: 0.85rem"
                        />
                      </div>
                    </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <MarketingWinner :campaign="campaign" />
    </div>

    <!-- Fallback for standard campaigns -->
    <div v-else class="standard-campaign-view">
      <p>Standard campaign view not implemented yet.</p>
    </div>

    <!-- Participant Editor Drawer -->
    <ParticipantEditor
        :is-open="isEditorOpen"
        :participant="editingParticipant"
        @close="isEditorOpen = false"
        @save="handleSaveParticipant"
    />

    <!-- Campaign Editor Drawer -->
    <CampaignEditor
        :is-open="isCampaignEditorOpen"
        :campaign="campaign"
        @close="isCampaignEditorOpen = false"
        @save="handleSaveCampaign"
    />
  </div>
</template>

<style scoped>
/* Detail View Styles */
.detail-header-card {
    background: #fff;
    padding: 24px;
    border-radius: 12px;
    border: 1px solid #e5e7eb;
    margin-bottom: 24px;
}

.status-row {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 24px;

    .dates {
        color: var(--text-secondary);
    }
}

.stats-overview {
    display: grid;

    /* Custom grid: Prize (2.5fr), Price (1fr), Entries (1fr), Raised (1fr), Goal (1fr) */
    grid-template-columns: 2.5fr 1fr 1fr 1fr 1fr;
    gap: 24px;

    .stat-box {
        display: flex;
        flex-direction: column;
        gap: 4px;

        .label {
            font-size: 0.8rem;
            text-transform: uppercase;
            color: var(--text-secondary);
            font-weight: 600;
        }

        .value {
            font-size: 1.5rem;
            font-weight: 700;
            color: var(--text-primary);
        }
    }
}

.raffle-management {
    display: grid;
    grid-template-columns: 2fr 1fr;
    gap: 24px;

    @media(width <= 1024px) {
        grid-template-columns: 1fr;
    }
}

.participants-section {
    background: #fff;
    padding: 24px;
    border-radius: 12px;
    border: 1px solid #e5e7eb;

    .section-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 16px;

        h2 { margin: 0; font-size: 1.2rem; }
    }
}

.add-entry-actions {
    display: flex;
    gap: 12px;
}

.participants-table-container {
    overflow-x: auto;
}

.participants-table {
    width: 100%;
    border-collapse: collapse;

    th {
        text-align: left;
        padding: 12px;
        border-bottom: 2px solid #f3f4f6;
        font-size: 0.85rem;
        color: var(--text-secondary);
    }

    td {
        padding: 12px;
        border-bottom: 1px solid #f3f4f6;
        font-size: 0.95rem;
    }

    tr.is-winner {
        /* User requested matching color from winner circle (L60) */
        background-color: hsl(from var(--color-warning) h s 60%);
        color: #fff; /* Ensure text is readable on darker background */
        td { font-weight: 600; }
    }
}

.winner-label {
    /* Lighter gold for the badge to contrast with the row */
    background: hsl(from var(--color-warning) h s 80%);
    color: hsl(from var(--color-warning) h s 30%); /* Darker text for contrast */
    font-size: 0.65rem;
    padding: 3px 0; /* Removing horizontal padding, using width instead */
    width: 60px;    /* Fixed width to approximate button width */
    text-align: center;
    border-radius: 12px;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    box-shadow: 0 1px 2px rgb(0 0 0 / 10%);
}

.actions-wrapper {
    display: flex;
    flex-direction: column;
    align-items: center; /* Center align badge and button */
    gap: 8px;
    justify-content: center;
    height: 100%;
    width: fit-content;
    margin-left: auto; /* Push to right side of cell */
}

.paid-badge {
    background: #d1fae5;
    color: #065f46;
    font-size: 0.75rem;
    padding: 2px 8px;
    border-radius: 12px;
    font-weight: 600;
}

.comments-cell {
    max-width: 150px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    color: var(--text-secondary);
    font-size: 0.85rem;
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
</style>
