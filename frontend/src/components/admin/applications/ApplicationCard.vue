<script setup lang="ts">
import { computed } from 'vue'

import { Capsules } from '../../common/ui'
import Button from '../../common/ui/Button.vue'

export interface IApplicationItem {
  id: string
  type: string
  status: 'pending' | 'approved' | 'denied' | 'needs_info' | 'autodeleted' | 'all' | 'video_approved'
  date: string
  applicantName: string
  email: string
  details: {
    petName?: string | null
    role?: string | null
    reason?: string | null
  }
  fullApplication: Record<string, unknown>
}

const props = defineProps<{
  app: IApplicationItem
  expanded: boolean
  isExpandedId: boolean
  isResending?: boolean
  isResendSuccess?: boolean
}>()

defineEmits<{
  toggle: []
  'update-status': [app: IApplicationItem, status: IApplicationItem['status']]
  'view-original': [id: string]
  'resend-email': [id: string]
}>()

import {
  formatDate,
  formatValue,
  getDaysPending,
  getDisplayFields,
  getSignatureSrc,
  getStatusColor,
  getStatusText,
} from './utils'

const displayFields = computed(() => {
  return getDisplayFields(props.app.fullApplication)
})


</script>

<template>
  <div
    class="app-card"
    :class="{
      unread: app.status === 'pending',
      expanded: expanded,
      'no-action': app.status === 'autodeleted'
    }"
    @click="app.status !== 'autodeleted' && $emit('toggle')"
  >

    <div class="card-summary">
      <div class="app-main">
        <div class="app-header">
          <h3>{{ app.applicantName }}</h3>
          <span class="date">{{ formatDate(app.date) }}</span>
        </div>
        <p class="email">{{ app.email }}</p>

        <div class="details-preview" v-if="!expanded">
          <span v-if="app.details.petName"><strong>Pet:</strong> {{ app.details.petName }}</span>
          <span v-if="app.details.role"><strong>Role:</strong> {{ app.details.role }}</span>
          <span v-if="app.details.reason"><strong>Reason:</strong> {{ app.details.reason }}</span>
        </div>
      </div>

      <div class="app-status">
        <Capsules
          :label="getStatusText(app.status)"
          :color="getStatusColor(app.status)"
          size="sm"
        />
        <div v-if="app.status !== 'autodeleted'" class="expand-icon" :class="{ rotated: expanded }">▼</div>
      </div>
    </div>

    <div v-if="expanded" class="expanded-content" @click.stop>
      <hr class="divider" />

      <div class="action-bar">
        <div class="action-group left">
          <!-- Pending State Actions -->
          <template v-if="app.status === 'pending'">
            <Button
              title="Approve"
              size="small"
              variant="primary"
              theme="primary"
              :onClick="() => $emit('update-status', app, 'approved')"
            />
            <Button
              title="Deny"
              size="small"
              variant="secondary"
              theme="danger"
              :onClick="() => $emit('update-status', app, 'denied')"
            />
          </template>

          <!-- Approved State Actions (Video Tour - Adoption Only) -->
          <template v-if="app.status === 'approved' && app.type === 'adoption'">
            <Button
              title="Confirm Video Tour"
              size="small"
              variant="primary"
              theme="primary"
              :onClick="() => $emit('update-status', app, 'video_approved')"
            />
            <Button
              title="Reject Video Tour"
              size="small"
              variant="secondary"
              theme="danger"
              :onClick="() => $emit('update-status', app, 'denied')"
            />
          </template>

          <!-- Video Approved State -->
          <template v-if="app.status === 'video_approved'">
            <span class="text-sm text-green-600 font-medium px-2">✓ Video Verified</span>
          </template>
        </div>

        <div class="action-group right">
          <Button
            :title="isResendSuccess ? 'Sent!' : 'Resend Email'"
            size="small"
            :variant="isResendSuccess ? 'primary' : 'tertiary'"
            :color="isResendSuccess ? 'green' : undefined"
            :loading="isResending"
            :onClick="() => $emit('resend-email', app.id)"
          />
          <Button
            title="View Original Application"
            size="small"
            variant="tertiary"
            :onClick="() => $emit('view-original', app.id)"
          />
        </div>
      </div>

      <div
        v-if="app.status === 'pending' && getDaysPending(app.date) > 5"
        class="mt-4 p-3 bg-red-50 border border-red-200 rounded-md text-sm text-red-700 flex items-center gap-2"
        style="margin-bottom: 20px"
      >
        ⚠️ <strong>Warning:</strong> This application will be automatically denied and deleted in
        {{ 7 - getDaysPending(app.date) }} day(s) due to the 7-day retention policy.
      </div>

      <div
        v-if="app.status === 'denied'"
        class="mt-4 p-3 bg-red-50 border border-red-200 rounded-md text-sm text-red-700 flex items-center gap-2"
        style="margin-bottom: 20px"
      >
        ⚠️ <strong>Warning:</strong> This application and all its data will be permanently deleted in 24
        hours.
      </div>

      <div class="full-application" v-if="app.fullApplication && app.status !== 'autodeleted'">
        <h4>Application Information</h4>
        <div class="qa-grid">
          <div v-for="field in displayFields" :key="field.key" class="qa-item">
            <span class="question">{{ field.label }}</span>
            <span class="answer">
              {{ formatValue(field.value) }}
            </span>
          </div>

          <div v-if="app.fullApplication.signatureData" class="qa-item span-full">
            <span class="question">Signature</span>
            <img
              v-if="getSignatureSrc(app.fullApplication.signatureData as string)"
              :src="getSignatureSrc(app.fullApplication.signatureData as string)"
              alt="Applicant Signature"
              class="max-h-24 w-auto border rounded bg-white p-2 mt-1"
            />
            <p v-else class="text-sm text-gray-400 italic">Invalid Signature Data</p>
          </div>

          <div v-if="app.fullApplication.parentSignatureData" class="qa-item span-full">
            <span class="question">Parent Signature</span>
            <img
              v-if="getSignatureSrc(app.fullApplication.parentSignatureData as string)"
              :src="getSignatureSrc(app.fullApplication.parentSignatureData as string)"
              alt="Parent Signature"
              class="max-h-24 w-auto border rounded bg-white p-2 mt-1"
            />
            <p v-else class="text-sm text-gray-400 italic">Invalid Signature Data</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.app-card {
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  overflow: hidden;
  transition: all 0.2s;
  cursor: pointer;
  border-left: 4px solid transparent;
}

.app-card:hover {
  box-shadow: 0 4px 12px rgb(0 0 0 / 5%);
  border-color: #d1d5db;
}

.app-card.unread {
  border-left-color: var(--color-secondary);
  background: var(--color-secondary-surface);
}

.app-card.expanded {
  background: #fff;
  border-color: #d1d5db;
  box-shadow: 0 8px 24px rgb(0 0 0 / 8%);
  cursor: default;
}

.app-card.no-action {
  cursor: default;
  background: #fafafa;
}

.app-card.no-action .card-summary {
  cursor: default;
}

.card-summary {
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  cursor: pointer;
}

.app-main {
  flex: 1;
}

.app-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 4px;
}

.app-header h3 {
  margin: 0;
  font-size: 1.1rem;
  color: var(--text-primary);
}

.app-header .date {
  font-size: 0.85rem;
  color: var(--color-neutral-text-soft);
}

.email {
  margin: 0 0 8px;
  font-size: 0.9rem;
  color: var(--color-neutral-text-soft);
}

.details-preview {
  font-size: 0.9rem;
  color: var(--color-neutral-text-soft);
  display: flex;
  gap: 16px;
}

.details-preview strong {
  color: var(--text-primary);
  font-weight: 600;
}

.app-status {
  display: flex;
  align-items: center;
  gap: 16px;
}

.expand-icon {
  font-size: 0.8rem;
  color: var(--color-neutral-text-soft);
  transition: transform 0.2s;
}

.rotated {
  transform: rotate(180deg);
}

.expanded-content {
  padding: 0 24px 24px;
  animation: slideDown 0.2s ease-out;
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.divider {
  border: 0;
  border-top: 1px solid #f3f4f6;
  margin: 0 0 20px;
}

.action-bar {
  display: flex;
  justify-content: space-between;
  margin-bottom: 24px;
  align-items: center;
}

.action-group {
  display: flex;
  gap: 12px;
  align-items: center;
}

.full-application {
  background: #f9fafb;
  border-radius: 8px;
  padding: 24px;
  border: 1px solid #f3f4f6;
}

.full-application h4 {
  margin: 0 0 16px;
  font-size: 1rem;
  color: var(--text-primary);
  font-weight: 600;
}

.qa-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

@media (width <= 1200px) {
  .qa-grid {
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  }
}

.qa-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.span-full {
  grid-column: 1 / span 2;
  margin-top: 16px;
}

.question {
  font-size: 0.85rem;
  font-weight: 600;
  color: var(--color-neutral-text-soft);
}


</style>

<style scoped>
@media (width <= 768px) {
  .card-summary {
    flex-direction: column;
    gap: 12px;
  }

  .app-status {
    width: 100%;
    justify-content: space-between;
  }

  .details-preview {
    flex-wrap: wrap;
    gap: 8px;
  }

  /* Reset grid to single column on mobile */
  .qa-grid {
    grid-template-columns: 1fr;
  }

  .span-full {
    grid-column: 1;
  }

  .action-bar {
    flex-direction: column;
    align-items: stretch;
    gap: 16px;
  }

  .action-group {
    flex-direction: column;
    width: 100%;
    justify-content: stretch;
    gap: 12px;
  }

  .action-group button {
    width: 100%;
  }

  /* Ensure left/right groups both take full width */
  .action-group.right {
    flex-direction: column-reverse; /* Resend/View Original order on mobile */
  }
}
</style>
