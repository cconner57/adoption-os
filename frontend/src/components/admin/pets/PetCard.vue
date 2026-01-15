<script setup lang="ts">
import { ref, computed } from 'vue'
import type { IPet } from '../../../models/common'
import { useRouter } from 'vue-router'
import { formatDate, calculateAge } from '../../../utils/date'

const props = defineProps<{
  pet: IPet
  statusFilter: string
}>()

const emit = defineEmits<{
  (e: 'edit', pet: IPet): void
  (e: 'archive', pet: IPet): void
}>()

const isExpanded = ref(false)
const router = useRouter()

// --- Helpers ---
function getStatusColor(status: string) {
  const map: Record<string, string> = {
    available: 'green',
    'adoption-pending': 'orange',
    adopted: 'blue',
    foster: 'purple',
    hold: 'red',
    intake: 'gray',
    archived: 'gray',
  }
  return map[status] || 'gray'
}

function formatDoB(dateString?: string | null) {
  return formatDate(dateString)
}

function formatCurrency(amount?: number | null) {
  if (amount == null) return '-'
  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(amount)
}

function formatList(list?: string[] | null) {
  if (!list || list.length === 0) return '-'
  return list.join(', ')
}

function isTrue(val: boolean | string | undefined | null) {
  return val === 'true' || val === true
}

// --- Computed ---
const vaccineSummary = computed(() => {
  const v = props.pet.medical?.vaccinations
  if (!v) return 'No recent records'
  const list: string[] = []
  if (v.rabies?.dateAdministered) list.push(`Rabies: ${formatDoB(v.rabies.dateAdministered)}`)
  if (v.bordetella?.dateAdministered)
    list.push(`Bordetella: ${formatDoB(v.bordetella.dateAdministered)}`)
  if (v.canineDistemper?.round1?.dateAdministered)
    list.push(`Distemper: ${formatDoB(v.canineDistemper.round1.dateAdministered)}`)
  if (v.felineDistemper?.round1?.dateAdministered)
    list.push(`FVRCP: ${formatDoB(v.felineDistemper.round1.dateAdministered)}`)
  return list.length ? list.join(' • ') : 'No recent records'
})
</script>

<template>
  <div class="pet-card" :class="{ expanded: isExpanded }">
    <!-- Header: Avatar, Name, Status -->
    <div class="card-header" @click="isExpanded = !isExpanded">
      <div class="header-main">
        <div class="pet-avatar">
          <img
            v-if="pet.photos && pet.photos.length > 0"
            :src="`/pet-photos/${pet.photos.find((p) => p.isPrimary)?.url ?? ''}`"
            alt="Pet"
          />
          <div v-else class="avatar-placeholder">
            {{ pet.species === 'cat' ? '🐱' : '🐶' }}
          </div>
        </div>
        <div class="pet-info">
          <div class="name-row">
            <h3>{{ pet.name }}</h3>
            <span class="status-badge" :class="getStatusColor(pet.details.status)">
              {{ pet.details.status.replace('-', ' ') }}
            </span>
          </div>
          <div class="sub-info">
            <div class="info-pill">{{ pet.physical.breed || pet.species }}</div>
            <div class="info-pill">{{ pet.sex }}</div>
            <div class="info-pill">{{ calculateAge(pet.physical.dateOfBirth) }}</div>
          </div>
        </div>
      </div>
      <div class="expand-icon" :class="{ rotated: isExpanded }">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="20"
          height="20"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <polyline points="6 9 12 15 18 9"></polyline>
        </svg>
      </div>
    </div>

    <!-- Collapsed Quick Stats -->
    <div class="card-body" v-if="!isExpanded">
      <div class="quick-stats">
        <div class="stat">
          <span class="label">LOCATION</span>
          <span class="value">{{ pet.details.shelterLocation || '-' }}</span>
        </div>
        <div class="stat">
          <span class="label">INTAKE</span>
          <span class="value">{{ formatDoB(pet.details.intakeDate) }}</span>
        </div>
      </div>
    </div>

    <!-- Expanded Details -->
    <div v-if="isExpanded" class="card-details">
      <!-- Reuse similar grid layout but stacked for mobile -->
      <div class="detail-section">
        <h4>Basic Info</h4>
        <div class="detail-row">
          <span class="label">Litter:</span>
          <span class="value">{{ pet.litterName || '-' }}</span>
        </div>
        <div class="detail-row">
          <span class="label">Color:</span>
          <span class="value">{{ pet.physical.color || '-' }}</span>
        </div>
        <div class="detail-row">
          <span class="label">Weight:</span>
          <span class="value">{{
            pet.physical.currentWeight ? `${pet.physical.currentWeight} lbs` : '-'
          }}</span>
        </div>
        <div class="detail-row">
          <span class="label">Spayed/Neutered:</span>
          <span class="value">
            {{ pet.medical.spayedOrNeutered ? 'Yes' : 'No' }}
          </span>
        </div>
      </div>

      <div class="detail-section">
        <h4>Status & Location</h4>
        <div class="detail-row">
          <span class="label">Location:</span>
          <span class="value">{{ pet.details.shelterLocation || '-' }}</span>
        </div>
        <div class="detail-row">
          <span class="label">Intake Date:</span>
          <span class="value">{{ pet.details.intakeDate || '-' }}</span>
        </div>

        <template v-if="pet.details.status === 'foster'">
          <div class="detail-row">
            <span class="label">Foster Parent:</span>
            <span class="value">{{ pet.foster.parentName || '-' }}</span>
          </div>
        </template>
        <template v-if="pet.details.status === 'adopted'">
          <div class="detail-row">
            <span class="label">Adopter:</span>
            <span class="value">{{ pet.adoption.adoptedBy || '-' }}</span>
          </div>
        </template>
      </div>

      <div class="detail-section">
        <h4>Behavior</h4>
        <div class="detail-row">
          <span class="label">Energy:</span>
          <span class="value capitalize">{{ pet.behavior?.energyLevel || '-' }}</span>
        </div>
        <div class="tags-row">
          <span class="tag" :class="pet.behavior?.isGoodWithKids ? 'yes' : 'no'">Kids</span>
          <span class="tag" :class="pet.behavior?.isGoodWithCats ? 'yes' : 'no'">Cats</span>
          <span class="tag" :class="pet.behavior?.isGoodWithDogs ? 'yes' : 'no'">Dogs</span>
        </div>
      </div>

      <!-- Actions Footer -->
      <div class="card-actions">
        <button class="action-btn edit" @click.stop="emit('edit', pet)">Edit Pet</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.pet-card {
  background: var(--text-inverse);
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  overflow: hidden;
  border: 1px solid var(--border-color);
  margin-bottom: 16px;
  transition: all 0.2s;
}

.pet-card.expanded {
  border-color: var(--color-secondary);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.card-header {
  padding: 16px 16px 12px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  cursor: pointer;
}

.header-main {
  display: flex;
  gap: 12px;
}

.pet-avatar {
  width: 64px;
  height: 64px;
  border-radius: 12px;
  overflow: hidden;
  background: #f1f5f9;
  flex-shrink: 0;
}
.pet-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.avatar-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem;
}

.pet-info {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.name-row {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.pet-info h3 {
  margin: 0;
  font-size: 1.15rem;
  color: var(--text-primary);
  font-weight: 700;
}

.sub-info {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
}

.info-pill {
  font-size: 0.8rem;
  color: #64748b;
  background: #f8fafc;
  padding: 2px 8px;
  border-radius: 6px;
  border: 1px solid #e2e8f0;
  text-transform: capitalize;
}

.status-badge {
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 0.7rem;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}
/* Reused status colors */
.status-badge.green {
  background: #dcfce7;
  color: #166534;
}
.status-badge.orange {
  background: #ffedd5;
  color: #9a3412;
}
.status-badge.blue {
  background: #dbeafe;
  color: #1e40af;
}
.status-badge.purple {
  background: #f3e8ff;
  color: #6b21a8;
}
.status-badge.red {
  background: #fee2e2;
  color: #991b1b;
}
.status-badge.gray {
  background: #f1f5f9;
  color: #64748b;
}

.expand-icon {
  color: #94a3b8;
  transition: transform 0.3s;
  margin-top: 4px; /* Align with top mostly */
}
.expand-icon.rotated {
  transform: rotate(180deg);
}

.card-body {
  padding: 0 16px 16px;
  padding-left: 92px; /* Align with text, skipping avatar width (64) + gap (12) + padding (16) = 92 */
}

.quick-stats {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}
.stat {
  display: flex;
  flex-direction: column;
  gap: 2px;
}
.stat .label {
  font-size: 0.65rem;
  color: #94a3b8;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  font-weight: 600;
}
.stat .value {
  font-size: 0.9rem;
  color: #334155;
  font-weight: 500;
}

.card-details {
  padding: 20px;
  background: #f8fafc;
  border-top: 1px solid var(--border-color);
}

.detail-section {
  margin-bottom: 20px;
}
.detail-section h4 {
  font-size: 0.8rem;
  text-transform: uppercase;
  color: #94a3b8;
  margin-bottom: 10px;
  font-weight: 700;
}
.detail-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  font-size: 0.95rem;
}
.detail-row .label {
  color: #64748b;
}
.detail-row .value {
  color: #334155;
  font-weight: 500;
  text-align: right;
}

.tags-row {
  display: flex;
  gap: 8px;
  margin-top: 8px;
}
.tag {
  padding: 4px 8px;
  border-radius: 6px;
  font-size: 0.8rem;
  border: 1px solid #e2e8f0;
}
.tag.yes {
  background: #dcfce7;
  color: #166534;
  border-color: #bbf7d0;
}
.tag.no {
  background: #fee2e2;
  color: #991b1b;
  border-color: #fecaca;
  text-decoration: line-through;
  opacity: 0.7;
}

.card-actions {
  display: flex;
  gap: 12px;
  margin-top: 24px;
}
.action-btn {
  flex: 1;
  padding: 10px;
  border-radius: 8px;
  border: none;
  font-weight: 600;
  cursor: pointer;
  font-size: 0.95rem;
}
.action-btn.edit {
  background: var(--color-primary);
  color: white;
}
.action-btn.archive {
  background: white;
  border: 1px solid #cbd5e1;
  color: #64748b;
}
</style>
