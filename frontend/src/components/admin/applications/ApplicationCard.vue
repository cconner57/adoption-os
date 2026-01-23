<script setup lang="ts">
import { computed } from 'vue'

import { Capsules } from '../../common/ui'
import Button from '../../common/ui/Button.vue'

export interface IApplicationItem {
  id: string
  type: string
  status: 'pending' | 'approved' | 'denied' | 'needs_info' | 'autodeleted' | 'all'
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
}>()

defineEmits<{
  toggle: []
  'update-status': [app: IApplicationItem, status: IApplicationItem['status']]
  'view-original': [id: string]
}>()

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString(undefined, {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

const getStatusColor = (status: string) => {
  switch (status) {
    case 'approved':
      return '#d1fae5'
    case 'denied':
      return '#fee2e2'
    case 'needs_info':
      return '#ffedd5'
    default:
      return '#f3f4f6'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'approved':
      return 'Approved'
    case 'denied':
      return 'Denied'
    case 'needs_info':
      return 'Needs Info'
    default:
      return 'Pending'
  }
}

const getDaysPending = (dateStr: string) => {
  if (!dateStr) return 0
  const created = new Date(dateStr)
  const now = new Date()
  const diffTime = Math.abs(now.getTime() - created.getTime())
  return Math.ceil(diffTime / (1000 * 60 * 60 * 24))
}

const formatKey = (key: string) => {
  const withSpaces = key.replace(/([a-z])([A-Z])/g, '$1 $2').replace(/_/g, ' ')
  return withSpaces
    .split(' ')
    .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
    .join(' ')
}

const getSignatureSrc = (data: string) => {
  if (!data) return ''
  if (data === 'base64data') {
    return 'data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iNTAiPjxwYXRoIGQ9Ik0xMCwyNSBRNDAsNSA3MCwyNSBUMTMwLDI1IiBzdHJva2U9ImJsYWNrIiBzdHJva2Utd2lkdGg9IjIiIGZpbGw9Im5vbmUiLz48L3N2Zz4='
  }
  if (data.length < 50) return ''
  if (data.startsWith('data:image')) return data
  return `data:image/png;base64,${data}`
}

const displayFields = computed(() => {
  const data = props.app.fullApplication
  if (!data) return []
  const entries = Object.entries(data)

  const filtered = entries.filter(([key, value]) => {
    if (key === 'id') return false
    if ((key === 'parentSignatureData' || key === 'parentSignatureDate') && !value) return false
    if (key === 'signatureData' || key === 'parentSignatureData') return false
    if (value === null || value === '' || value === undefined) return false
    return true
  })

  filtered.sort(([keyA], [keyB]) => {
    const priority = [
      'firstName',
      'lastName',
      'age',
      'birthday',
      'address',
      'city',
      'zip',
      'createdAt',
      'nameFull',
      'signatureDate',
      'allergies',
      'phoneNumber',
      'availability',
      'positionPreferences',
      'interestReason',
      'volunteerExperience',
      'emergencyContactName',
      'emergencyContactPhone',
    ]
    const idxA = priority.indexOf(keyA)
    const idxB = priority.indexOf(keyB)

    if (idxA !== -1 && idxB !== -1) return idxA - idxB
    if (idxA !== -1) return -1
    if (idxB !== -1) return 1

    return 0
  })

  return filtered.map(([key, value]) => {
    let label = formatKey(key)
    let displayValue: unknown = value

    if (key === 'createdAt') {
      label = 'Submitted At'
    }

    const dateKeys = ['birthday', 'signatureDate', 'parentSignatureDate']
    if (dateKeys.includes(key) && typeof value === 'string') {
      if (value.startsWith('0001-01-01')) {
        displayValue = 'N/A'
      } else {
        displayValue = formatDate(value)
      }
    }

    if (key === 'createdAt' && typeof value === 'string') {
      if (value.startsWith('0001-01-01')) {
        displayValue = 'N/A'
      } else {
        const date = new Date(value)
        displayValue = `${formatDate(value)} at ${date.toLocaleTimeString('en-US', {
          hour: 'numeric',
          minute: '2-digit',
        })}`
      }
    }

    if (value === true || value === 'true') displayValue = 'Yes'
    if (value === false || value === 'false') displayValue = 'No'

    return {
      key,
      label,
      value: displayValue,
    }
  })
})
</script>

<template>
  <div
    class="app-card"
    :class="{
      unread: app.status === 'pending',
      expanded: expanded,
    }"
    @click="$emit('toggle')"
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
        <div class="expand-icon" :class="{ rotated: expanded }">▼</div>
      </div>
    </div>

    <div v-if="expanded" class="expanded-content" @click.stop>
      <hr class="divider" />

      <div class="action-bar">
        <Button
          title="Approve"
          size="small"
          variant="primary"
          theme="primary"
          :disabled="app.status === 'approved'"
          :onClick="() => $emit('update-status', app, 'approved')"
        />
        <Button
          title="Request Info"
          size="small"
          variant="primary"
          theme="warning"
          :disabled="app.status === 'needs_info'"
          :onClick="() => $emit('update-status', app, 'needs_info')"
        />
        <Button
          title="Deny"
          size="small"
          variant="secondary"
          theme="danger"
          :disabled="app.status === 'denied'"
          :onClick="() => $emit('update-status', app, 'denied')"
        />
        <Button
          title="View Original Application"
          size="small"
          variant="tertiary"
          :onClick="() => $emit('view-original', app.id)"
        />
      </div>

      <div
        v-if="app.status === 'pending' && getDaysPending(app.date) > 5"
        class="mt-4 p-3 bg-red-50 border border-red-200 rounded-md text-sm text-red-700 flex items-center gap-2"
        style="margin-bottom: 20px"
      >
        ⚠️ <strong>Warning:</strong> This application will be automatically denied and deleted in
        {{ 7 - getDaysPending(app.date) }} day(s) due to the 7-day retention policy.
      </div>

      <div class="full-application" v-if="app.fullApplication">
        <h4>Application Information</h4>
        <div class="qa-grid">
          <div v-for="field in displayFields" :key="field.key" class="qa-item">
            <span class="question">{{ field.label }}</span>
            <span class="answer">
              {{ Array.isArray(field.value) ? field.value.join(', ') : field.value }}
            </span>
          </div>

          <div v-if="app.fullApplication.signatureData" class="qa-item span-full">
            <span class="question">Signature</span>
            <img
              v-if="getSignatureSrc(app.fullApplication.signatureData as string)"
              :src="getSignatureSrc(app.fullApplication.signatureData as string)"
              class="max-h-24 w-auto border rounded bg-white p-2 mt-1"
            />
            <p v-else class="text-sm text-gray-400 italic">Invalid Signature Data</p>
          </div>

          <div v-if="app.fullApplication.parentSignatureData" class="qa-item span-full">
            <span class="question">Parent Signature</span>
            <img
              v-if="getSignatureSrc(app.fullApplication.parentSignatureData as string)"
              :src="getSignatureSrc(app.fullApplication.parentSignatureData as string)"
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
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  overflow: hidden;
  transition: all 0.2s;
  cursor: pointer;
  border-left: 4px solid transparent;
}

.app-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  border-color: #d1d5db;
}

.app-card.unread {
  border-left-color: var(--color-secondary);
  background: hsl(from var(--color-secondary) h s 98%);
}

.app-card.expanded {
  background: white;
  border-color: #d1d5db;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
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
  color: hsl(from var(--color-neutral) h s 50%);
}

.email {
  margin: 0 0 8px 0;
  font-size: 0.9rem;
  color: hsl(from var(--color-neutral) h s 50%);
}

.details-preview {
  font-size: 0.9rem;
  color: hsl(from var(--color-neutral) h s 50%);
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
  color: hsl(from var(--color-neutral) h s 50%);
  transition: transform 0.2s;
}

.rotated {
  transform: rotate(180deg);
}

.expanded-content {
  padding: 0 24px 24px 24px;
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
  margin: 0 0 20px 0;
}

.action-bar {
  display: flex;
  gap: 12px;
  margin-bottom: 24px;
  align-items: center;
}

.action-bar > :last-child {
  margin-left: auto;
}

.full-application {
  background: #f9fafb;
  border-radius: 8px;
  padding: 24px;
  border: 1px solid #f3f4f6;
}

.full-application h4 {
  margin: 0 0 16px 0;
  font-size: 1rem;
  color: var(--text-primary);
  font-weight: 600;
}

.qa-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

@media (max-width: 1200px) {
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
  color: hsl(from var(--color-neutral) h s 50%);
}

.answer {
  
}
</style>
