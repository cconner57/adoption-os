<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { computed, onMounted,reactive } from 'vue'
import { useRouter } from 'vue-router'

import {
  AggressiveSection,
  BehaviorSection,
  FeedingSection,
  HouseholdSection,
  MedicalSection,
  OtherSection,
} from '../components/about/surrender/index.ts'
import PetSelectSection from '../components/about/surrender/PetSelectSection.vue'
import SurrenderSteps from '../components/about/surrender/SurrenderSteps.vue'
import FormSubmitted from '../components/common/form-submitted/FormSubmitted.vue'
import Button from '../components/common/ui/Button.vue'
import { useMetrics } from '../composables/useMetrics'
import { useSurrenderStore } from '../stores/surrender'

const { submitMetric } = useMetrics()

onMounted(() => {
  submitMetric('form_start', { form: 'surrender' })
})

const router = useRouter()
const surrenderStore = useSurrenderStore()
const {
  formState,
  step,
  isSubmitted,
  hasAttemptedSubmit,
  selectedAnimal,
  validationErrors,
  isStepValid,
} = storeToRefs(surrenderStore)
const { nextStep, prevStep, resetForm } = surrenderStore

const touched = reactive<Record<string, boolean>>({})

const handleBlur = (field: string) => {
  touched[field] = true
}

const formError = computed(() => hasAttemptedSubmit.value && !isStepValid.value)

const handleSubmit = () => {
  if (!nextStep()) {
    globalThis.scrollTo({ top: document.body.scrollHeight, behavior: 'smooth' })

    return
  }
  globalThis.scrollTo({ top: 0, behavior: 'smooth' })
}

const handleReset = () => {
  resetForm()
  router.push('/')
}

const headerText = computed(() => {
  if (!selectedAnimal.value || step.value === 0) {
    return 'Surrender Pet'
  }
  return selectedAnimal.value === 'cat' ? 'Cat Surrender' : 'Dog Surrender'
})

const formattedAnimal = computed(() => {
  if (!selectedAnimal.value) return ''
  return selectedAnimal.value.charAt(0).toUpperCase() + selectedAnimal.value.slice(1)
})
</script>

<template>
  <section class="page-shell">
    <section v-if="!isSubmitted" class="form-card" aria-labelledby="form-title">
      <div class="form-header">
        <img
          v-if="selectedAnimal === 'cat' && step > 0"
          src="/images/cat.png"
          alt="cat"
          height="50"
          width="100"
        />
        <img
          v-if="selectedAnimal === 'dog' && step > 0"
          src="/images/dog.png"
          alt="cat"
          height="50"
          width="100"
        />
        <h1>{{ headerText }}</h1>
      </div>
      <SurrenderSteps
        v-if="selectedAnimal && step > 0"
        :formStep="step"
        :selectedAnimal="selectedAnimal"
      />
      <PetSelectSection
        v-if="step === 0"
        :formError="formError"
        :selectedAnimal="selectedAnimal"
        @update:selectedAnimal="(value) => (selectedAnimal = value)"
      />
      <HouseholdSection
        v-if="step === 1 && selectedAnimal"
        :formState="formState"
        :touched="touched"
        :handleBlur="handleBlur"
        :hasAttemptedSubmit="hasAttemptedSubmit"
        :selectedAnimal="formattedAnimal"
      />
      <BehaviorSection
        v-if="step === 2 && selectedAnimal"
        :formState="formState"
        :touched="touched"
        :handleBlur="handleBlur"
        :hasAttemptedSubmit="hasAttemptedSubmit"
        :selectedAnimal="formattedAnimal"
      />
      <AggressiveSection
        v-if="step === 3 && selectedAnimal"
        :formState="formState"
        :touched="touched"
        :handleBlur="handleBlur"
        :hasAttemptedSubmit="hasAttemptedSubmit"
        :selectedAnimal="formattedAnimal"
      />
      <MedicalSection
        v-if="step === 4 && selectedAnimal"
        :formState="formState"
        :touched="touched"
        :handleBlur="handleBlur"
        :hasAttemptedSubmit="hasAttemptedSubmit"
        :selectedAnimal="formattedAnimal"
      />
      <FeedingSection
        v-if="step === 5 && selectedAnimal"
        :formState="formState"
        :touched="touched"
        :handleBlur="handleBlur"
        :hasAttemptedSubmit="hasAttemptedSubmit"
        :selectedAnimal="formattedAnimal"
      />
      <OtherSection
        v-if="step === 6 && selectedAnimal"
        :formState="formState"
        :touched="touched"
        :handleBlur="handleBlur"
        :hasAttemptedSubmit="hasAttemptedSubmit"
        :selectedAnimal="formattedAnimal"
      />

      <div
        v-if="hasAttemptedSubmit && validationErrors.length > 0"
        class="validation-summary"
        role="alert"
        aria-live="polite"
      >
        <p class="summary-title">Please complete the following required fields:</p>
        <div class="tags">
          <span v-for="err in validationErrors" :key="err" class="tag is-danger">{{ err }}</span>
        </div>
      </div>

      <div class="actions">
        <Button
          v-if="step > 0"
          @click="prevStep"
          title="Back"
          color="white"
          size="large"
          style="border: 1px solid var(--color-primary); color: var(--color-primary)"
        />
        <Button
          @click="handleSubmit"
          type="submit"
          :title="step === 6 ? 'Submit' : 'Next'"
          color="green"
          size="large"
        />
      </div>
    </section>

    <FormSubmitted v-else @reset="handleReset" formType="surrender" />
  </section>
</template>

<style scoped lang="css">
.page-shell {
  min-height: 100vh;
  background-color: var(--color-primary);
  padding: 9rem var(--layout-padding-side) 64px;
  container-type: inline-size;
  container-name: shell;

  /* Viewport query for container's own padding */
  @media (max-width: 440px) {
    padding: 6rem 16px 32px;
  }

  .form-card {
    max-width: 1600px;
    margin: 0 auto;
    background: var(--text-inverse);
    color: var(--text-primary);
    border-radius: 24px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    padding: 48px 48px 32px;

    /* Logic based on shell width */
    @container shell (max-width: 800px) {
      padding: 32px 16px;
    }

    .form-header {
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 16px;
      margin-bottom: 4px;
      color: var(--color-primary);

      h1 {
        font-size: 4.25rem;
        line-height: 1.2;
        letter-spacing: 0.2px;
        color: var(--color-primary);
      }
      img {
        width: 100px;
      }

      /* Tablet/Mobile adjustments based on container width */
      @container shell (max-width: 800px) {
        flex-direction: column;
        align-items: center;
        gap: 0px;
        margin-bottom: 1rem;
        h1 {
          font-size: 2.25rem;
          text-align: center;
        }
        img {
          width: 60px;
          height: 60px;
        }
      }

      @container shell (max-width: 480px) {
        h1 {
          font-size: 1.75rem;
        }
      }
    }

    .actions {
      display: flex;
      justify-content: center;
      gap: 16px;
      margin-top: 20px;

      @media (max-width: 600px) {
        flex-direction: column;

        button {
          width: 100%;
        }
      }
    }
  }

  /* Nested validation summary */
  /* Nested validation summary matching Volunteer.vue */
  .validation-summary {
    background-color: #fff1f2;
    border: 1px solid #e11d48;
    color: #9f1239;
    border-radius: 12px;
    padding: 1.5rem;
    margin: 2rem 0;
    text-align: center;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);

    .summary-title {
      font-weight: 700;
      margin-bottom: 1rem;
      font-size: 1.1rem;
    }

    .tags {
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
      gap: 0.5rem;
    }

    .tag.is-danger {
      background-color: #fce7f3;
      color: #9f1239;
      padding: 0.5rem 1rem;
      border-radius: 999px;
      font-size: 0.9rem;
      font-weight: 600;
    }
  }
}
</style>
