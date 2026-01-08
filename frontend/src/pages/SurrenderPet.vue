<script setup lang="ts">
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useSurrenderStore } from '../stores/surrender'
import {
  AggressiveSection,
  FeedingSection,
  HouseholdSection,
  MedicalSection,
  BehaviorSection,
  OtherSection,
} from '../components/about/surrender/index.ts'
import Button from '../components/common/ui/Button.vue'
import SurrenderSteps from '../components/about/surrender/SurrenderSteps.vue'
import PetSelectSection from '../components/about/surrender/PetSelectSection.vue'
import FormSubmitted from '../components/common/form-submitted/FormSubmitted.vue'
import { reactive, computed, onMounted } from 'vue'
import { useMetrics } from '../composables/useMetrics'

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

      <div v-if="hasAttemptedSubmit && validationErrors.length > 0" class="validation-summary">
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
          style="border: 1px solid var(--green); color: var(--green)"
        />
        <Button
          @click="handleSubmit"
          type="submit"
          :title="step === 6 ? 'Submit' : 'Next'"
          color="green"
          size="large"
          :disabled="step === 0 && !selectedAnimal"
        />
      </div>
    </section>

    <FormSubmitted v-else @reset="handleReset" formType="surrender" />
  </section>
</template>

<style scoped lang="css">
.page-shell {
  min-height: 100vh;
  background-color: var(--green);
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
    background: var(--white);
    color: var(--font-color-dark);
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
      color: var(--green);

      h1 {
        font-size: 4.25rem;
        line-height: 1.2;
        letter-spacing: 0.2px;
        color: var(--green);
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
    }
  }

  /* Nested validation summary */
  .validation-summary {
    background-color: #fef2f2;
    border: 1px solid #ef4444;
    border-radius: 12px;
    padding: 16px;
    margin: 24px 0;
    text-align: center;
    display: flex;
    flex-direction: column;
    align-items: center;

    .summary-title {
      color: #b91c1c;
      font-weight: 600;
      margin-bottom: 12px;
    }

    .tags {
      display: flex;
      flex-wrap: wrap;
      gap: 8px;
      justify-content: center;
    }

    .tag.is-danger {
      background-color: #fee2e2;
      color: #b91c1c;
      padding: 4px 12px;
      border-radius: 16px;
      font-size: 0.875rem;
      font-weight: 500;
    }
  }
}
</style>
