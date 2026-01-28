<script setup lang="ts">
/* eslint-disable max-lines */
import { computed } from 'vue'

import type { IPet } from '../../../models/common'
import { calculateAge, formatDigitDate, formatIntakeDate } from '../../../utils/date'
import { Button, Icon } from '../../common/ui'

const props = defineProps<{
  pet: IPet
  index: number
  visibleColumns: Record<string, boolean>
  isExpanded: boolean
  statusFilter: string
}>()

const emit = defineEmits<{
  'toggle-expand': [pet: IPet]
  edit: [pet: IPet]
  archive: [pet: IPet]
}>()

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
  return formatDigitDate(dateString)
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
  if (val === 'true') return true
  if (val === true) return true
  return false
}

const vaccineSummary = computed(() => {
  const v = props.pet.medical?.vaccinations
  if (!v) return 'No recent records'

  const list: string[] = []
  if (v.rabies?.dateAdministered) list.push(`Rabies: ${formatDoB(v.rabies.dateAdministered)}`)
  if (v.bordetella?.dateAdministered)
    list.push(`Bordetella: ${formatDoB(v.bordetella.dateAdministered)}`)

  if (v.canineDistemper?.round1?.dateAdministered)
    list.push(`Distemper (DHPP): ${formatDoB(v.canineDistemper.round1.dateAdministered)}`)
  if (v.felineDistemper?.round1?.dateAdministered)
    list.push(`FVRCP: ${formatDoB(v.felineDistemper.round1.dateAdministered)}`)

  return list.length ? list.join(' • ') : 'No recent records'
})

const colCount = computed(() => {
  let count = 1
  if (props.visibleColumns.photo) count++
  if (props.visibleColumns.name) count++
  if (props.visibleColumns.breed) count++
  if (props.visibleColumns.sex) count++
  if (props.visibleColumns.sn) count++
  if (props.visibleColumns.microchip) count++
  if (props.visibleColumns.age) count++
  if (props.visibleColumns.dob) count++
  if (props.visibleColumns.intake) count++

  if (props.statusFilter === 'adopted') count++
  if (props.statusFilter === 'foster') count++

  if (props.visibleColumns.status) count++
  if (props.visibleColumns.actions) count++
  return count
})
</script>

<template>

  <tr
    class="pet-row"
    :class="{ expanded: isExpanded, 'even-row': index % 2 === 1 }"
    @click="emit('toggle-expand', pet)"
  >

    <td class="expand-col">
      <button class="icon-btn expand-btn" :class="{ rotated: isExpanded }" title="Toggle Details">
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
      </button>
    </td>

    <td v-if="visibleColumns.photo">
      <div class="pet-avatar">
        <img
          v-if="pet.photos && pet.photos.length > 0"
          :src="`/pet-photos/${pet.photos?.find((p) => p.isPrimary)?.url ?? ''}`"
          alt="Pet"
        />
        <div v-else class="avatar-placeholder">
          {{ pet.species === 'cat' ? '🐱' : '🐶' }}
        </div>
      </div>
    </td>
    <td v-if="visibleColumns.name" class="font-bold">
      <router-link :to="`/adopt/${pet.id}`" target="_blank" class="pet-link" @click.stop>
        {{ pet.name }}
      </router-link>
    </td>
    <td v-if="visibleColumns.breed" class="capitalize">{{ pet.species }}</td>
    <td v-if="visibleColumns.sex" class="capitalize">{{ pet.sex }}</td>
    <td v-if="visibleColumns.sn" class="text-center">
      <Icon
        v-if="pet.medical?.spayedOrNeutered"
        name="check"
        size="22"
        style="color: var(--color-primary)"
        title="Spayed/Neutered"
      />
      <Icon
        v-else
        name="cross"
        size="22"
        class="text-muted"
        title="Not Spayed/Neutered"
      />
    </td>
    <td v-if="visibleColumns.microchip">
      <div v-if="pet.medical?.microchip?.microchipID" class="microchip-stack mono-text">
        <div v-for="(part, idx) in pet.medical.microchip.microchipID.split(' ')" :key="idx">
          {{ part }}
        </div>
      </div>
      <span v-else class="text-muted">-</span>
    </td>
    <td v-if="visibleColumns.age" style="white-space: nowrap">
      {{ calculateAge(pet.physical.dateOfBirth) }}
    </td>
    <td v-if="visibleColumns.dob">
      <span v-if="pet.physical.dateOfBirth">
        {{ formatDoB(pet.physical.dateOfBirth) }}
      </span>
      <span v-else class="text-muted">-</span>
    </td>
    <td v-if="visibleColumns.intake">
      <span v-if="pet.details?.intakeDate">
        {{ formatIntakeDate(pet.details.intakeDate) }}
      </span>
      <span v-else class="text-muted">-</span>
    </td>

    <td v-if="statusFilter === 'adopted'">
      <span v-if="pet.adoption?.date">{{ formatDoB(pet.adoption.date) }}</span>
      <span v-else class="text-muted">-</span>
    </td>
    <td v-if="statusFilter === 'foster'">
      <span v-if="pet.foster?.startDate">{{ formatDoB(pet.foster.startDate) }}</span>
      <span v-else class="text-muted">-</span>
    </td>

    <td v-if="visibleColumns.status">
      <span class="status-badge" :class="getStatusColor(pet.details.status)">
        {{ pet.details.status.replace('-', ' ') }}
      </span>
    </td>

    <td v-if="visibleColumns.actions">
      <div class="row-actions" @click.stop>
        <Button
          size="small"
          variant="secondary"
          color="white"
          title="Edit"
          class="row-btn"
          @click="emit('edit', pet)"
        >
          <span>Edit</span>
        </Button>
      </div>
    </td>
  </tr>

  <tr v-if="isExpanded" class="details-row">
    <td :colspan="colCount">
      <div class="expanded-content">
        <div class="details-grid">
          <div class="detail-section">
            <h4>Basic Info</h4>
            <div class="detail-item">
              <span class="label">Breed:</span>
              <span class="value">{{ pet.physical.breed || 'Unknown' }}</span>
            </div>
            <div class="detail-item" v-if="pet.litterName">
              <span class="label">Litter:</span>
              <span class="value">{{ pet.litterName }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Color:</span>
              <span class="value">{{ pet.physical.color || '-' }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Sex:</span>
              <span class="value capitalize">{{ pet.sex }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Age Group:</span>
              <span class="value capitalize">{{ pet.physical.ageGroup }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Birth Date:</span>
              <span class="value">
                {{ formatDoB(pet.physical.dateOfBirth) }}
                <span class="text-muted" v-if="pet.physical.dateOfBirth">
                  ({{ calculateAge(pet.physical.dateOfBirth) }})
                </span>
              </span>
            </div>
            <div class="detail-item">
              <span class="label">Size / Weight:</span>
              <span class="value capitalize">
                {{ pet.physical.size || '-' }}
                <span v-if="pet.physical.currentWeight">
                  / {{ pet.physical.currentWeight }} lbs
                </span>
              </span>
            </div>
            <div class="detail-item">
              <span class="label">Coat:</span>
              <span class="value capitalize">{{ pet.physical.coatLength }}</span>
            </div>
          </div>

          <div class="detail-section">
            <h4>Behavior & Compatibility</h4>
            <div class="detail-item">
              <span class="label">Energy Level:</span>
              <span class="value capitalize">{{ pet.behavior?.energyLevel || '-' }}</span>
            </div>
            <div class="detail-item">
              <span class="label">House Trained:</span>
              <span class="value">{{ pet.behavior?.isHouseTrained ? 'Yes' : 'No' }}</span>
            </div>
            <div class="detail-item compatibility-grid">
              <span class="label">Good With:</span>
              <div class="comp-tags">
                <span class="tag" :class="pet.behavior?.isGoodWithKids ? 'yes' : 'no'">Kids</span>
                <span class="tag" :class="pet.behavior?.isGoodWithCats ? 'yes' : 'no'">Cats</span>
                <span class="tag" :class="pet.behavior?.isGoodWithDogs ? 'yes' : 'no'">Dogs</span>
              </div>
            </div>
            <div class="detail-item" v-if="pet.behavior?.personalityTags?.length">
              <span class="label">Personality:</span>
              <span class="tag-list">
                <span v-for="tag in pet.behavior.personalityTags" :key="tag" class="p-tag">{{
                  tag
                }}</span>
              </span>
            </div>
            <div class="detail-item warning" v-if="pet.behavior?.prefersToBeAlone">
              <span class="label">Note:</span>
              <span>Prefers to be alone</span>
            </div>
          </div>

          <div class="detail-section">
            <h4>Adoption & Status</h4>
            <div class="detail-item">
              <span class="label">Intake Date:</span>
              <span class="value">{{ formatIntakeDate(pet.details.intakeDate) }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Location:</span>
              <span class="value">{{ pet.details.shelterLocation || '-' }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Environment:</span>
              <span class="value capitalize">{{ pet.details.environmentType || '-' }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Sponsored:</span>
              <div v-if="pet.sponsored.isSponsored">
                <span class="value warning">YES</span>
                <div class="text-muted text-small">
                  {{ pet.sponsored.sponsoredBy }} ({{ formatCurrency(pet.sponsored.amount) }})
                </div>
              </div>
              <span class="value" v-else>No</span>
            </div>

            <div class="detail-item mt-2">
              <span class="label">Profile Visibility:</span>
              <div class="settings-grid">
                <div
                  class="setting-tag"
                  :class="{ active: isTrue(pet.profileSettings?.isSpotlightFeatured) }"
                >
                  Spotlight: {{ isTrue(pet.profileSettings?.isSpotlightFeatured) ? 'ON' : 'OFF' }}
                </div>
                <div
                  class="setting-tag"
                  :class="{ active: isTrue(pet.profileSettings?.showMedicalHistory) }"
                >
                  Public Medical:
                  {{ isTrue(pet.profileSettings?.showMedicalHistory) ? 'ON' : 'OFF' }}
                </div>
                <div
                  class="setting-tag"
                  :class="{ active: isTrue(pet.profileSettings?.showAdditionalInformation) }"
                >
                  Add. Info:
                  {{ isTrue(pet.profileSettings?.showAdditionalInformation) ? 'ON' : 'OFF' }}
                </div>
              </div>
            </div>

            <template v-if="pet.details.status === 'adopted'">
              <div class="detail-item">
                <span class="label">Adopted Date:</span>
                <span class="value">{{ formatDoB(pet.adoption.date) }}</span>
              </div>
              <div class="detail-item">
                <span class="label">Adopted By:</span>
                <span class="value">
                  {{ pet.adoption.adoptedBy || '-' }}
                  <div v-if="pet.adoption.adopterContactInfo" class="text-muted text-small">
                    {{ pet.adoption.adopterContactInfo.email }}
                    <br />
                    {{ pet.adoption.adopterContactInfo.phone }}
                  </div>
                </span>
              </div>
              <div class="detail-item">
                <span class="label">Adoption Fee:</span>
                <span class="value">{{ formatCurrency(pet.adoption.fee) }}</span>
              </div>
              <div class="detail-item" v-if="pet.adoption.surveyCompleted !== undefined">
                <span class="label">Survey:</span>
                <span class="value">{{
                  pet.adoption.surveyCompleted ? 'Completed' : 'Pending'
                }}</span>
              </div>
            </template>

            <template v-if="pet.details.status === 'foster'">
              <div class="detail-item">
                <span class="label">Foster Parent:</span>
                <span class="value">
                  {{ pet.foster.parentName || '-' }}
                  <div v-if="pet.foster.fosterContactInfo" class="text-muted text-small">
                    {{ pet.foster.fosterContactInfo.email }}
                  </div>
                </span>
              </div>
              <div class="detail-item">
                <span class="label">Dates:</span>
                <span class="value">
                  {{ formatDoB(pet.foster.startDate) }} -
                  {{ pet.foster.endDate ? formatDoB(pet.foster.endDate) : 'Present' }}
                </span>
              </div>
            </template>
          </div>

          <div class="detail-section">
            <h4>Medical Snapshot</h4>
            <div class="detail-item">
              <span class="label">Spayed/Neutered:</span>
              <span class="value">
                {{ pet.medical.spayedOrNeutered ? 'Yes' : 'No' }}
                <span class="text-muted" v-if="pet.medical.spayedOrNeuteredDate">
                  ({{ formatDoB(pet.medical.spayedOrNeuteredDate) }})
                </span>
              </span>
            </div>
            <div class="detail-item">
              <span class="label">Microchip:</span>
              <span class="value mono-text" v-if="pet.medical?.microchip?.microchipID">
                {{ pet.medical.microchip.microchipID }}
                <span class="text-muted" v-if="pet.medical.microchip.microchipCompany">
                  ({{ pet.medical.microchip.microchipCompany }})
                </span>
              </span>
              <span class="value" v-else>-</span>
            </div>
            <div class="detail-item">
              <span class="label">Health Concerns:</span>
              <span
                class="value"
                :class="{ warning: pet.medical.healthConcerns?.length }"
                style="text-transform: capitalize"
              >
                {{ formatList(pet.medical.healthConcerns) }}
              </span>
            </div>
            <div class="detail-item">
              <span class="label">Current Meds:</span>
              <span class="value">{{ formatList(pet.medical.currentMedications) }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Vaccinations:</span>
              <span class="value text-small">{{ vaccineSummary }}</span>
            </div>
            <div class="detail-item" v-if="pet.medical.surgeries?.length">
              <span class="label">Surgeries:</span>
              <span class="value">{{ pet.medical.surgeries.map((s) => s.name).join(', ') }}</span>
            </div>
          </div>

          <div
            class="detail-section"
            v-if="pet.returned.isReturned || pet.details.status === 'intake'"
          >
            <h4>History & Returns</h4>
            <div class="detail-item">
              <span class="label">Returned?</span>
              <span class="value warning" v-if="pet.returned.isReturned">YES</span>
              <span class="value" v-else>No</span>
            </div>
            <div class="detail-item" v-if="pet.returned.isReturned">
              <span class="label">Return Date:</span>
              <span class="value">{{ formatDoB(pet.returned.history[0]?.date) || '-' }}</span>
            </div>
            <div class="detail-item" v-if="pet.returned.isReturned">
              <span class="label">Return Reason:</span>
              <span class="value">{{ pet.returned.history[0]?.reason || '-' }}</span>
            </div>
          </div>

          <div class="detail-section">
            <h4>Descriptions</h4>
            <div class="desc-group">
              <span class="label">Origin</span>
              <p class="description-text italic line-clamp-5">
                {{ pet.descriptions.origin ? `"${pet.descriptions.origin}"` : '-' }}
              </p>
            </div>
            <div class="desc-group">
              <span class="label">Primary Bio</span>
              <p class="description-text line-clamp-5">
                {{ pet.descriptions.primary || '-' }}
              </p>
            </div>
            <div class="desc-group">
              <span class="label">Spotlight</span>
              <p class="description-text line-clamp-5">
                {{ pet.descriptions.spotlight ? `"${pet.descriptions.spotlight}"` : '-' }}
              </p>
            </div>

            <div class="desc-group">
              <span class="label">Health Summary</span>
              <p class="description-text line-clamp-5">
                {{ pet.behavior?.healthSummary || '-' }}
              </p>
            </div>
            <div class="desc-group">
              <span class="label">Additional Info</span>
              <ul class="info-list" v-if="pet.descriptions.additionalInformation?.length">
                <li v-for="info in pet.descriptions.additionalInformation" :key="info">
                  {{ info }}
                </li>
              </ul>
              <p class="description-text" v-else>-</p>
            </div>
          </div>
        </div>
      </div>
    </td>
  </tr>
</template>

<style scoped>
.pet-row {
  cursor: pointer;
  border-bottom: 2px solid var(--border-color);
  transition: background-color 0.2s;
}

.pet-row.even-row {
  background-color: var(--color-neutral-weak);
}

.pet-row:hover {
  background-color: var(--color-neutral-weak);
}

.pet-row.expanded {
  background-color: var(--color-secondary-weak);
  border-bottom: none;
}

td {
  padding: 16px;
  vertical-align: middle;
  color: var(--text-primary);
}

.expand-col {
  width: 40px;
  padding: 0 8px;
  text-align: center;
}

.details-row td {
  padding: 0;
  background-color: var(--text-inverse);
}

.expanded-content {
  padding: 24px;
  border-top: 1px dashed var(--border-color);
  background-color: var(--color-neutral-surface);
  border-bottom: 2px solid var(--border-color);
}

.details-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 24px;
}

.detail-section h4 {
  font-size: 0.85rem;
  color: var(--color-neutral-text-soft);
  text-transform: uppercase;
  margin-bottom: 12px;
  font-weight: 700;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 4px;
}

.detail-item {
  display: flex;
  flex-direction: column;
  margin-bottom: 8px;
}

.label {
  font-size: 0.8rem;
  color: var(--color-neutral-text-soft);
}

.value {
  font-size: 0.95rem;
  color: var(--text-primary);
  font-weight: 500;
}

.warning {
  color: var(--color-danger);
  font-weight: 700;
}

.text-small {
  font-size: 0.85rem;
}

.description-text {
  font-size: 0.9rem;
  color: var(--text-primary);
  line-height: 1.5;
  margin-top: 2px;
}

.desc-group {
  margin-bottom: 12px;
}

.comp-tags {
  display: flex;
  gap: 4px;
  flex-wrap: wrap;
  margin-top: 4px;
}

.tag {
  font-size: 0.75rem;
  padding: 2px 6px;
  border-radius: 4px;
  background: #f1f5f9;
  color: #475569;
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

.tag-list {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-top: 4px;
}

.p-tag {
  background: var(--color-secondary-weak);
  color: var(--color-secondary);
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 0.8rem;
  font-weight: 500;
}

.info-list {
  margin: 4px 0 0 16px;
  padding: 0;
  font-size: 0.9rem;
  color: var(--text-primary);
}

.italic {
  font-style: italic;
}

.line-clamp-5 {
  display: -webkit-box;
  -webkit-line-clamp: 5;
  line-clamp: 5;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}

.mt-2 {
  margin-top: 8px;
}

.settings-grid {
  display: flex;
  flex-direction: column;
  gap: 4px;
  margin-top: 4px;
  align-items: flex-start;
}



.row-btn {
  min-width: 90px;
  justify-content: center;
}

.setting-tag {
  font-size: 0.7rem;
  padding: 2px 6px;
  border-radius: 4px;
  background: #f1f5f9;
  color: #475569;
  border: 1px solid #cbd5e1;
  white-space: nowrap;
}

.setting-tag.active {
  background: var(--color-secondary-weak);
  color: var(--color-secondary);
  border-color: #bbf7d0;
  border-color: var(--color-secondary-light);
  font-weight: 500;
}

.pet-avatar {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  overflow: hidden;
  background: var(--color-neutral-weak);
  display: flex;
  align-items: center;
  justify-content: center;
}

.pet-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.status-badge {
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
  white-space: nowrap;
}

.status-badge.green {
  background-color: var(--color-neutral-weak);
  color: var(--text-primary);
}

.status-badge.orange {
  background-color: var(--color-warning-weak);
  color: var(--color-warning);
}

.status-badge.blue {
  background-color: var(--color-secondary-weak);
  color: var(--color-secondary);
}

.status-badge.purple {
  background-color: var(--color-secondary-weak);
  color: var(--color-secondary);
}

.status-badge.red {
  background-color: var(--color-danger-weak);
  color: var(--color-danger);
}

.status-badge.gray {
  background-color: var(--color-neutral-weak);
  color: var(--color-neutral-text-soft);
}

.mono-text {
  font-family: monospace;
  font-size: 0.9em;
  color: var(--text-primary);
  background: var(--text-inverse);
  padding: 2px 6px;
  border-radius: 4px;
  border: 1px solid var(--border-color);
  display: inline-block;
  width: fit-content;
}

.microchip-stack {
  display: flex !important;
  flex-direction: column;
  align-items: flex-start;
  width: min-content;
  white-space: nowrap;
}

.row-actions {
  display: flex;
  flex-direction: row; /* Changed from column */
  gap: 8px;
  justify-content: flex-end; /* Align to right */
  align-items: center;
}

.icon-btn {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  border: 1px solid transparent;
  background: transparent;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.1rem;
  transition: all 0.2s;
}

.icon-btn.edit:hover {
  background: var(--color-secondary-weak);
  color: var(--color-secondary);
}

.action-btn {
  background: var(--color-neutral-surface);
  color: var(--text-secondary);
  border: 1px solid var(--border-color);
}

.action-btn:hover {
  background: var(--color-neutral-weak);
  color: var(--text-primary);
  border-color: var(--color-neutral-border-strong);
}

.action-btn.adopt {
  color: var(--color-primary);
  background: var(--color-primary-weak);
  border-color: transparent;
}

.action-btn.adopt:hover {
  background: var(--color-primary);
  color: var(--text-inverse);
}

.icon-btn.archive:hover {
  background: var(--color-danger-weak);
  color: var(--color-danger);
}

.expand-btn {
  transition: transform 0.2s;
  color: var(--color-neutral-text-soft);
}

.expand-btn:hover {
  background: var(--color-neutral-weak);
  color: var(--text-primary);
}

.expand-btn.rotated {
  transform: rotate(
    -90deg
  );
}

.pet-link {
  color: var(--text-primary);
  text-decoration: none;
}

.pet-link:hover {
  color: var(--color-secondary);
  text-decoration: underline;
}

.capitalize {
  text-transform: capitalize;
}
</style>
