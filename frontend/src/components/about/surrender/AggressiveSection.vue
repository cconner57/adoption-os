<script setup lang="ts">
import { onMounted } from 'vue'
import InputTextArea from '../../common/ui/InputTextArea.vue'
import type { SurrenderFormState } from '../../../models/common'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()

onMounted(() => {
  if (!formState.catEverAttackedPeople) {
    formState.catEverAttackedPeople = 'No'
  }
  if (!formState.catEverAttackedOtherCats) {
    formState.catEverAttackedOtherCats = 'No'
  }
})
</script>

<template>
  <div class="aggressive-section">
    <h5>Aggressive Behavior</h5>
    <div class="aggressive-grid">
      <!-- Toggle for Attacked People -->
      <div class="field-group">
        <label class="field-label">Has the cat ever attacked or bit a person?</label>
        <div class="toggle-container">
          <button
            type="button"
            class="toggle-btn"
            :class="{ active: formState.catEverAttackedPeople === 'Yes' }"
            @click="formState.catEverAttackedPeople = 'Yes'"
          >
            Yes
          </button>
          <button
            type="button"
            class="toggle-btn"
            :class="{ active: formState.catEverAttackedPeople === 'No' }"
            @click="formState.catEverAttackedPeople = 'No'"
          >
            No
          </button>
        </div>
      </div>

      <InputTextArea
        label="If yes, please explain"
        placeholder="Explanation"
        :spanFull="false"
        :modelValue="formState.catEverAttackedPeopleExplanation"
        @update:modelValue="(val) => (formState.catEverAttackedPeopleExplanation = val)"
      />

      <!-- Toggle for Attacked Cats -->
      <div class="field-group">
        <label class="field-label">Has the cat ever attacked or bit a cat?</label>
        <div class="toggle-container">
          <button
            type="button"
            class="toggle-btn"
            :class="{ active: formState.catEverAttackedOtherCats === 'Yes' }"
            @click="formState.catEverAttackedOtherCats = 'Yes'"
          >
            Yes
          </button>
          <button
            type="button"
            class="toggle-btn"
            :class="{ active: formState.catEverAttackedOtherCats === 'No' }"
            @click="formState.catEverAttackedOtherCats = 'No'"
          >
            No
          </button>
        </div>
      </div>

      <InputTextArea
        label="If yes, please explain"
        placeholder="Explanation"
        :spanFull="false"
        :modelValue="formState.catEverAttackedOtherCatsExplanation"
        @update:modelValue="(val) => (formState.catEverAttackedOtherCatsExplanation = val)"
      />
    </div>
  </div>
</template>

<style scoped lang="css">
.aggressive-section h5 {
  margin-bottom: 24px;
}

.aggressive-grid {
  display: flex;
  flex-direction: column;
  gap: 16px;

  @media (min-width: 768px) {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    column-gap: 24px;
    row-gap: 16px;
  }
}

.field-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.field-label {
  font-size: 0.875rem;
  font-weight: 600;
  color: #374151;
  /* Match InputTextArea label style usually found in global or scoped */
  margin-bottom: 4px;
}

.toggle-container {
  display: flex;
  background: #fff;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  padding: 4px;
  height: 48px; /* Match typical input height */
}

.toggle-btn {
  flex: 1;
  border: none;
  background: transparent;
  border-radius: 6px;
  font-weight: 500;
  color: #64748b;
  cursor: pointer;
  padding: 0 4px;
  transition: all 0.2s;
  font-size: 1rem;

  &.active {
    background: color-mix(in srgb, var(--green) 10%, white);
    border: 1px solid var(--green);
    box-shadow: 0 0 0 1px var(--green) inset;
    color: var(--font-color-dark);
    font-weight: 600;
  }

  &:hover:not(.active) {
    background-color: #f1f5f9;
  }
}
</style>
