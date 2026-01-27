<script setup lang="ts">
import { computed, ref } from 'vue'

import type { IPet } from '../../../models/common'
import { usePetStore } from '../../../stores/pets'
import { Button, Checkbox, Icon } from '../../common/ui'
import Drawer from '../../common/ui/Drawer.vue'

const petStore = usePetStore()

// --- Modal/Drawer State ---
const isDrawerOpen = ref(false)
const selectedPet = ref<IPet | null>(null)
const actionType = ref<'spayNeuter' | 'rabies' | null>(null)
const resolutionDate = ref('')
const markUnknownDate = ref(false)

// --- Pagination State ---
const currentPage = ref(1)
const itemsPerPage = 5

// --- Action Logic ---
const actionItems = computed(() => {
  const allPets = petStore.currentPets
  const items: { pet: IPet; type: 'spayNeuter' | 'rabies'; label: string }[] = []

  allPets.forEach((pet) => {
    // Check Spay/Neuter
    const isFixed = pet.medical.spayedOrNeutered
    const hasFixDate = !!pet.medical.spayedOrNeuteredDate

    if (!isFixed || (isFixed && !hasFixDate)) {
      items.push({
        pet,
        type: 'spayNeuter',
        label: isFixed ? 'Missing S/N Date' : 'Neuter/Spay Required',
      })
    }

    // Check Rabies
    const hasRabies = !!pet.medical.vaccinations?.rabies?.dateAdministered
    if (!hasRabies) {
      items.push({
        pet,
        type: 'rabies',
        label: 'Rabies Vax Required',
      })
    }
  })

  // Sort by name for stability
  return items.sort((a, b) => a.pet.name.localeCompare(b.pet.name))
})

const paginatedItems = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return actionItems.value.slice(start, end)
})

const totalPages = computed(() => Math.ceil(actionItems.value.length / itemsPerPage))

const nextPage = () => {
  if (currentPage.value < totalPages.value) currentPage.value++
}

const prevPage = () => {
  if (currentPage.value > 1) currentPage.value--
}

// --- Handler Logic ---
const openResolveDrawer = (item: typeof actionItems.value[0]) => {
  selectedPet.value = item.pet
  actionType.value = item.type
  resolutionDate.value = '' // Reset to empty, user must enter
  markUnknownDate.value = false // Default unchecked
  isDrawerOpen.value = true
}

const closeDrawer = () => {
  isDrawerOpen.value = false
  selectedPet.value = null
  actionType.value = null
}

const canSave = computed(() => {
  if (actionType.value === 'spayNeuter') {
    // Valid if date is entered OR unknown is checked
    return !!resolutionDate.value || markUnknownDate.value
  } else if (actionType.value === 'rabies') {
    // Rabies needs a date
    return !!resolutionDate.value
  }
  return false
})

const saveResolution = async () => {
  if (!selectedPet.value || !actionType.value || !canSave.value) return

  const petToUpdate = structuredClone(selectedPet.value) as IPet

  if (actionType.value === 'spayNeuter') {
    petToUpdate.medical.spayedOrNeutered = true // Always true if we are resolving this
    if (resolutionDate.value) {
      petToUpdate.medical.spayedOrNeuteredDate = resolutionDate.value
    } else if (markUnknownDate.value) {
      // Explicitly keeping it null/undefined but ensuring boolean is true
      petToUpdate.medical.spayedOrNeuteredDate = undefined
    }
  } else if (actionType.value === 'rabies') {
    if (!petToUpdate.medical.vaccinations.rabies) {
        petToUpdate.medical.vaccinations.rabies = { dateAdministered: '' }
    }
    if (resolutionDate.value) {
        petToUpdate.medical.vaccinations.rabies.dateAdministered = resolutionDate.value
    }
  }

  await petStore.updatePet(petToUpdate)
  closeDrawer()
}
</script>

<template>
  <div class="widget action-items-widget">
    <div class="widget-header">
        <h3>Action Items</h3>
        <span class="count-badge" v-if="actionItems.length > 0">{{ actionItems.length }} Pending</span>
    </div>

    <div class="items-list" v-if="paginatedItems.length > 0">
      <div v-for="(item, index) in paginatedItems" :key="`${item.pet.id}-${item.type}`" class="action-item">
        <div class="item-info">
          <span class="pet-name">{{ item.pet.name }}</span>
          <span class="issue-label" :class="item.type">{{ item.label }}</span>
        </div>
        <Button
          size="small"
          variant="secondary"
          @click="openResolveDrawer(item)"
          title="Resolve Action"
        >
          <Icon name="check" size="16" />
        </Button>
      </div>
    </div>

    <div v-else class="empty-state">
      <p>🎉 All caught up! No urgent actions.</p>
    </div>

    <div class="pagination-controls" v-if="totalPages > 1">
      <button class="nav-btn" :disabled="currentPage === 1" @click="prevPage">
        <Icon name="arrowRight" size="16" style="transform: rotate(180deg)" />
      </button>
      <span class="page-count">{{ currentPage }} / {{ totalPages }}</span>
      <button class="nav-btn" :disabled="currentPage === totalPages" @click="nextPage">
        <Icon name="arrowRight" size="16" />
      </button>
    </div>

    <!-- Resolution Drawer -->
    <Drawer
      :isOpen="isDrawerOpen"
      :title="`Resolve: ${selectedPet?.name}`"
      @close="closeDrawer"
    >
      <div class="drawer-content-inner">
        <p class="drawer-desc">
          Update medical records for <strong>{{ actionType === 'spayNeuter' ? 'Spay/Neuter' : 'Rabies Vaccination' }}</strong>.
        </p>

        <div class="form-group">
            <label for="resolutionDate">Date Completed</label>
            <input
              type="date"
              id="resolutionDate"
              v-model="resolutionDate"
              class="date-input"
              :disabled="markUnknownDate"
            />
        </div>

        <div class="form-group checkbox-group" v-if="actionType === 'spayNeuter'">
            <Checkbox
              id="markUnknown"
              v-model="markUnknownDate"
              label="Date is unknown (Mark as fixed)"
            />
        </div>

        <div class="info-box" v-if="markUnknownDate">
          <Icon name="info" size="16" />
          <span>Pet will be marked as Spayed/Neutered, but the date will remain unrecorded.</span>
        </div>
      </div>

      <template #footer>
        <Button variant="secondary" @click="closeDrawer">Cancel</Button>
        <Button @click="saveResolution" :disabled="!canSave">Save & Complete</Button>
      </template>
    </Drawer>
  </div>
</template>

<style scoped>
.widget {
  background: var(--text-inverse);
  padding: 16px; /* Reduced from 24px */
  border-radius: 16px;
  border: 1px solid var(--border-color);
  box-shadow: 0 4px 6px rgb(0 0 0 / 5%);
  display: flex;
  flex-direction: column;
  gap: 12px; /* Reduced from 16px */
  min-height: 450px; /* Enforce consistent height for 5 items */
}

.widget-header {
    display: flex;
    justify-content: space-between;
    align-items: center;

    h3 {
        font-size: 1rem; /* Slightly smaller */
        margin: 0;
        font-weight: 700;
        color: var(--text-primary);
    }
}

.count-badge {
    background: var(--color-neutral-weak); /* Changed from danger-weak */
    color: var(--text-primary); /* Changed from danger */
    font-size: 0.75rem;
    padding: 2px 8px;
    border-radius: 12px;
    font-weight: 600;
    border: 1px solid var(--border-color);
}

.items-list {
  display: flex;
  flex-direction: column;
  gap: 8px; /* Reduced from 12px */
  min-height: 340px; /* Reserve space for 5 items to prevent layout shift */
  justify-content: flex-start;
}

.action-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 12px; /* Reduced vertical padding from 12px */
  background: var(--color-neutral-surface);
  border-radius: 8px; /* Slightly tighter radius */
  border: 1px solid transparent;

  &:hover {
    border-color: var(--color-neutral-border);
  }
}

.item-info {
  display: flex;
  flex-direction: column;
  gap: 2px; /* Reduced from 4px */
}

.pet-name {
  font-weight: 600;
  color: var(--text-primary);
  font-size: 0.9rem; /* Slightly smaller */
}

.issue-label {
  font-size: 0.7rem; /* Slightly smaller */
  font-weight: 500;
}

.issue-label.spayNeuter { color: var(--color-warning-strong); }
.issue-label.rabies { color: var(--color-danger); }

.pagination-controls {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
  margin-top: 4px; /* Reduced from 8px */
  padding-top: 8px; /* Reduced from 12px */
  border-top: 1px solid var(--border-color);
}

.nav-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: var(--text-secondary);
  display: flex;
  align-items: center;
  padding: 4px;
  border-radius: 4px;

  &:disabled {
    opacity: 0.3;
    cursor: not-allowed;
  }

  &:not(:disabled):hover {
    background: var(--color-neutral-weak);
    color: var(--text-primary);
  }
}

.page-count {
  font-size: 0.85rem;
  color: var(--text-secondary);
}

.empty-state {
    padding: 24px;
    text-align: center;
    color: var(--text-secondary);
    font-style: italic;
}

/* Drawer Content Styles */
.drawer-content-inner {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.drawer-desc {
    font-size: 0.95rem;
    color: var(--text-secondary);
    margin: 0;
    line-height: 1.5;
}

.form-group {
    display: flex;
    flex-direction: column;
    gap: 8px;

    label {
        font-size: 0.85rem;
        font-weight: 500;
        color: var(--text-secondary);
    }
}

.checkbox-group {
    flex-direction: row;
    align-items: center;
    gap: 12px;
}

.date-input {
    padding: 12px;
    border-radius: 8px;
    border: 1px solid var(--border-color);
    font-family: inherit;
    font-size: 1rem;
    color: var(--text-primary);
    background: var(--text-inverse);

    &:focus {
        outline: none;
        border-color: var(--color-primary);
        box-shadow: 0 0 0 3px oklch(from var(--color-primary) l c h / 20%);
    }

    &:disabled {
      background: var(--color-neutral-surface);
      color: var(--text-secondary);
      cursor: not-allowed;
    }
}

.info-box {
  display: flex;
  gap: 12px;
  background: var(--color-neutral-surface);
  padding: 12px;
  border-radius: 8px;
  font-size: 0.85rem;
  color: var(--text-secondary);
  align-items: flex-start;

  :deep(svg) {
    flex-shrink: 0;
    color: var(--text-secondary);
  }
}
</style>
