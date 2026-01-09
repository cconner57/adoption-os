<script setup lang="ts">
import { computed } from 'vue'
import type { IPet } from '../../../models/common'
import { useRouter } from 'vue-router'

const props = defineProps<{
  pet: IPet
  index: number
  visibleColumns: Record<string, boolean>
  isExpanded: boolean
  statusFilter: string
}>()

const emit = defineEmits<{
  (e: 'toggle-expand', pet: IPet): void
  (e: 'edit', pet: IPet): void
  (e: 'archive', pet: IPet): void
}>()

const router = useRouter()

// --- Helpers (moved from Pets.vue or similar) ---

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
  // status tag text is already handled, simplified return for class
  return map[status] || 'gray'
}

function formatDoB(dateString?: string | null) {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

function calculateAge(dateOfBirth?: string | null) {
  if (!dateOfBirth) return '-'
  const birthDate = new Date(dateOfBirth)
  const today = new Date()
  let years = today.getFullYear() - birthDate.getFullYear()
  let months = today.getMonth() - birthDate.getMonth()
  if (months < 0 || (months === 0 && today.getDate() < birthDate.getDate())) {
    years--
    months += 12
  }
  if (today.getDate() < birthDate.getDate()) months--

  if (years === 0 && months === 0) return '< 1 mo'
  if (years === 0) return `${months} mo`
  if (months === 0) return `${years} yr`
  return `${years} yr ${months} mo`
}

function formatCurrency(amount?: number | null) {
  if (amount == null) return '-'
  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(amount)
}

function formatList(list?: string[] | null) {
  if (!list || list.length === 0) return '-'
  return list.join(', ')
}

// --- Computed Properties for Expanded Data ---

const hasVaccinations = computed(() => {
  const v = props.pet.medical.vaccinations
  return v && (v.rabies || v.bordetella || v.canineDistemper || v.felineDistemper)
})

const vaccineSummary = computed(() => {
  const v = props.pet.medical.vaccinations
  const list = []
  if (v.rabies?.dateAdministered) list.push(`Rabies: ${formatDoB(v.rabies.dateAdministered)}`)
  if (v.bordetella?.dateAdministered)
    list.push(`Bordetella: ${formatDoB(v.bordetella.dateAdministered)}`)
  // Add simplified checks for series
  if (v.canineDistemper?.round1?.dateAdministered)
    list.push(`Distemper (DHPP): ${formatDoB(v.canineDistemper.round1.dateAdministered)}`)
  if (v.felineDistemper?.round1?.dateAdministered)
    list.push(`FVRCP: ${formatDoB(v.felineDistemper.round1.dateAdministered)}`)

  return list.length ? list.join(' • ') : 'No recent records'
})
</script>

<template>
  <!-- Main Row -->
  <tr
    class="pet-row"
    :class="{ expanded: isExpanded, 'even-row': index % 2 === 1 }"
    @click="emit('toggle-expand', pet)"
  >
    <!-- Expand Arrow Column -->
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
          :src="`/images/${pet.photos?.find((p) => p.isPrimary)?.url ?? ''}`"
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
    <td v-if="visibleColumns.breed">{{ pet.physical.breed || 'Unknown' }}</td>
    <td v-if="visibleColumns.sex" class="capitalize">{{ pet.sex }}</td>
    <td v-if="visibleColumns.sn" class="text-center">
      <span v-if="pet.medical?.spayedOrNeutered" title="Spayed/Neutered">✅</span>
      <span v-else class="text-muted" title="Not Spayed/Neutered">❌</span>
    </td>
    <td v-if="visibleColumns.microchip">
      <span v-if="pet.medical?.microchip?.microchipID" class="mono-text">
        {{ pet.medical.microchip.microchipID }}
      </span>
      <span v-else class="text-muted">-</span>
    </td>
    <td v-if="visibleColumns.age">
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
        {{ formatDoB(pet.details.intakeDate) }}
      </span>
      <span v-else class="text-muted">-</span>
    </td>

    <!-- Dynamic Data -->
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

    <!-- Actions -->
    <td v-if="visibleColumns.actions" align="right">
      <div class="row-actions">
        <button class="icon-btn edit" @click.stop="emit('edit', pet)" title="Edit">✎</button>
        <button class="icon-btn archive" @click.stop="emit('archive', pet)" title="Archive">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="18"
            height="18"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <rect width="20" height="5" x="2" y="3" rx="1" />
            <path d="M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8" />
            <path d="M10 12h4" />
          </svg>
        </button>
      </div>
    </td>
  </tr>

  <!-- Expanded Details Row -->
  <tr v-if="isExpanded" class="details-row">
    <td :colspan="100">
      <div class="expanded-content">
        <div class="details-grid">
          <!-- Adoption & Sponsorship -->
          <div class="detail-section">
            <h4>Adoption & Status</h4>
            <div class="detail-item">
              <span class="label">Adopted Date:</span>
              <span class="value">{{ formatDoB(pet.adoption.date) }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Adopted By:</span>
              <span class="value">{{ pet.adoption.adoptedBy || '-' }}</span>
            </div>
            <div class="detail-item">
              <span class="label">New Name:</span>
              <span class="value">{{ pet.adoption.newAdoptedName || '-' }}</span>
            </div>
            <div class="detail-item">
              <span class="label">Adoption Fee:</span>
              <span class="value">{{ formatCurrency(pet.adoption.fee) }}</span>
            </div>
            <div class="detail-item" v-if="pet.sponsored.isSponsored">
              <span class="label">Sponsored By:</span>
              <span class="value"
                >{{ pet.sponsored.sponsoredBy }} ({{ formatCurrency(pet.sponsored.amount) }})</span
              >
            </div>
          </div>

          <!-- Medical Info -->
          <div class="detail-section">
            <h4>Medical Snapshot</h4>
            <div class="detail-item">
              <span class="label">Health Concerns:</span>
              <span class="value" :class="{ warning: pet.medical.healthConcerns?.length }">
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
          </div>

          <!-- Returns & History -->
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
              <span class="value">{{ formatDoB(pet.returned.date) }}</span>
            </div>
            <div class="detail-item" v-if="pet.returned.isReturned">
              <span class="label">Return Reason:</span>
              <span class="value">{{ pet.returned.reason || '-' }}</span>
            </div>
          </div>

          <!-- Additional Notes -->
          <div class="detail-section">
            <h4>Descriptions</h4>
            <p class="description-text">
              {{ pet.descriptions.primary || 'No primary description available.' }}
            </p>
          </div>
        </div>
      </div>
    </td>
  </tr>
</template>

<style scoped>
.pet-row {
  cursor: pointer;
  border-bottom: 2px solid var(--border-color); /* Strong outline for row separation */
  transition: background-color 0.2s;
}

.pet-row.even-row {
  background-color: hsl(from var(--color-neutral) h s 98%);
}

.pet-row:hover {
  background-color: hsl(from var(--color-neutral) h s 95%);
}

.pet-row.expanded {
  background-color: hsl(from var(--color-secondary) h s 95%);
  border-bottom: none; /* Merges with details row visually if needed, but we wanted outline */
}

td {
  padding: 16px 24px;
  vertical-align: middle;
  color: var(--text-primary);
}

/* Override padding for the expand icon column */
.expand-col {
  width: 40px;
  padding: 0 8px;
  text-align: center;
}

.details-row td {
  padding: 0;
  /* border-bottom: 2px solid var(--border-color);  Ensure the bottom of the details row closes the "card" */
  background-color: var(--text-inverse);
}

.expanded-content {
  padding: 24px;
  border-top: 1px dashed var(--border-color);
  background-color: hsl(from var(--color-neutral) h s 98%);
}

.details-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 24px;
}

.detail-section h4 {
  font-size: 0.85rem;
  color: hsl(from var(--color-neutral) h s 50%);
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
  color: hsl(from var(--color-neutral) h s 50%);
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
}

/* Inherited styles that need to be scoped here or global */
.pet-avatar {
  width: 48px;
  height: 48px;
  border-radius: 8px;
  overflow: hidden;
  background: hsl(from var(--color-neutral) h s 95%);
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
/* Re-include status badge colors or rely on global if they are global (they were scoped in Pets.vue) */
.status-badge.green {
  background-color: hsl(from var(--color-primary) h s 95%);
  color: var(--color-primary);
}
.status-badge.orange {
  background-color: hsl(from var(--color-warning) h s 95%);
  color: var(--color-warning);
}
.status-badge.blue {
  background-color: hsl(from var(--color-secondary) h s 95%);
  color: var(--color-secondary);
}
.status-badge.purple {
  background-color: hsl(from var(--color-secondary) h s 92%); /* Using secondary tint for purple */
  color: var(--color-secondary);
}
.status-badge.red {
  background-color: hsl(from var(--color-danger) h s 95%);
  color: var(--color-danger);
}
.status-badge.gray {
  background-color: hsl(from var(--color-neutral) h s 95%);
  color: hsl(from var(--color-neutral) h s 50%);
}

.mono-text {
  font-family: monospace;
  font-size: 0.9em;
  color: var(--text-primary);
  background: var(--text-inverse); /* Contrast against row background */
  padding: 2px 6px;
  border-radius: 4px;
  border: 1px solid var(--border-color);
}

.row-actions {
  display: flex;
  gap: 8px;
  justify-content: flex-end;
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
  background: hsl(from var(--color-secondary) h s 95%);
  color: var(--color-secondary);
}
.icon-btn.archive:hover {
  background: hsl(from var(--color-danger) h s 95%);
  color: var(--color-danger);
}

.expand-btn {
  transition: transform 0.2s;
  color: hsl(from var(--color-neutral) h s 50%);
}
.expand-btn:hover {
  background: hsl(from var(--color-neutral) h s 95%);
  color: var(--text-primary);
}
.expand-btn.rotated {
  transform: rotate(
    -90deg
  ); /* Points down then rotates up? Standard is 0>180 or -90>0. Let's see */
}

.pet-link {
  color: var(--text-primary);
  text-decoration: none;
}
.pet-link:hover {
  color: var(--color-secondary);
  text-decoration: underline;
}
</style>
