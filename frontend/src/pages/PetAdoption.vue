<script setup lang="ts">
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useAdoptionStore } from '../stores/adoption'
import { usePetStore } from '../stores/pets'

import HomeSection from '../components/pet-adoption/cat-adoption/HomeSection.vue'
import NewCatSection from '../components/pet-adoption/cat-adoption/NewCatSection.vue'
import GeneralSection from '../components/pet-adoption/cat-adoption/GeneralSection.vue'
import CurrentPetsSection from '../components/pet-adoption/cat-adoption/CurrentPetsSection.vue'
import PastPetsSection from '../components/pet-adoption/cat-adoption/PastPetsSection.vue'
import OtherSection from '../components/pet-adoption/cat-adoption/OtherSection.vue'
import SummarySection from '../components/pet-adoption/cat-adoption/SummarySection.vue'
import Button from '../components/common/ui/Button.vue'
import AdoptionSteps from '../components/pet-adoption/adoption-steps/AdoptionSteps.vue'
import ApplicationHeader from '../components/volunteer/application-header/ApplicationHeader.vue'
import FormSubmitted from '../components/common/form-submitted/FormSubmitted.vue'

const router = useRouter()
const adoptionStore = useAdoptionStore()
const petStore = usePetStore()

const { formState, step, isSubmitted, hasAttemptedSubmit, validationErrors } =
  storeToRefs(adoptionStore)
const { selectedPet } = storeToRefs(petStore)

const { nextStep, prevStep, resetForm } = adoptionStore

import { reactive } from 'vue'
const touched = reactive<Record<string, boolean>>({})

const handleBlur = (field: string) => {
  touched[field] = true
}

const handleSubmit = async () => {
  if (!adoptionStore.isStepValid) {
    hasAttemptedSubmit.value = true
    globalThis.scrollTo({ top: document.body.scrollHeight, behavior: 'smooth' })
    return
  }

  if (step.value < 6) {
    adoptionStore.nextStep()
    globalThis.scrollTo({ top: 0, behavior: 'smooth' })
  } else {
    console.log('Submitting form...')
    await adoptionStore.submitApplication()

    petStore.clearSelectedPet()
    globalThis.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const handleReset = () => {
  resetForm()
  router.push('/')
}
</script>

<template>
  <section class="page-shell">
    <section v-if="!isSubmitted" class="form-card" aria-labelledby="form-title">
      <ApplicationHeader
        :header-title="selectedPet?.species === 'cat' ? 'Cat' : 'Dog'"
        :header-text="
          selectedPet?.species === 'cat'
            ? 'This application is intended as a means to match the right cat with the right home. The more detail you provide, the better.  All of our adoptable pets are spayed/neutered, vaccinated, and microchipped. Typical adoption fees are $300 for kittens and $250 for adults. Adoption fees are tax-deductible donations, not purchase prices. Thank you for considering adoption!'
            : 'This application is intended as a means to match the right dog with the right home. The more detail you provide, the better.  All of our adoptable pets are spayed/neutered, vaccinated, and microchipped. Typical adoption fees are $450 for puppies, $400 for adults, and $350 for seniors. Adoption fees are tax-deductible donations, not purchase prices. Thank you for considering adoption!'
        "
      />
      <AdoptionSteps :formStep="step" selectedAnimal="cat" />
      <div class="cat-name-display">
        <h2>Adopting Pet:</h2>
        <p>{{ selectedPet?.petName }}</p>
      </div>
      <GeneralSection
        v-show="step === 0"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <HomeSection
        v-show="step === 1"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <NewCatSection
        v-show="step === 2"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <CurrentPetsSection
        v-show="step === 3"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <PastPetsSection
        v-show="step === 4"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <OtherSection
        v-show="step === 5"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />
      <SummarySection
        v-show="step === 6"
        v-model="formState"
        :touched="touched"
        :handleBlur="handleBlur"
      />

      <div v-if="hasAttemptedSubmit && validationErrors.length > 0" class="validation-summary">
        <p class="summary-title">Please complete the following required fields:</p>
        <div class="tags">
          <span v-for="err in validationErrors" :key="err" class="tag is-danger">{{ err }}</span>
        </div>
      </div>

      <div class="actions">
        <Button
          @click="prevStep"
          title="Back"
          color="white"
          size="large"
          style="border: 1px solid var(--green); color: var(--green)"
          :disabled="step === 0 || isSubmitted"
        />
        <Button
          @click="handleSubmit"
          type="submit"
          :title="step < 6 ? 'Next' : 'Submit Application'"
          color="green"
          size="large"
          :disabled="isSubmitted"
        />
      </div>
    </section>

    <FormSubmitted
      v-else
      @reset="handleReset"
      title="Application Submitted"
      text="Thank you for your application! We will review it as soon as possible."
      form-type="adoption"
    />
  </section>
</template>

<style scoped lang="css">
.page-shell {
  min-height: 100vh;
  background-color: var(--green);
  padding: 9rem var(--layout-padding-side) 64px;
  .form-card {
    max-width: 1600px;
    margin: 0 auto;
    background: var(--white);
    color: var(--font-color-dark);
    border-radius: 24px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    padding: 48px 48px 32px;
    fieldset {
      border: 0;
      margin: 24px 0;
      padding: 0;
      .section-title {
        font-weight: 700;
        font-size: 18px;
        margin: 18px 0 12px;
      }
    }
    .actions {
      display: flex;
      justify-content: center;
      gap: 16px;
      margin-top: 20px;
    }
    .grid {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 16px;
    }
  }
  @media (min-width: 321px) and (max-width: 430px) {
    padding: 6rem 16px 32px;
  }
  @media (min-width: 431px) and (max-width: 768px) {
  }
}
.cat-name-display {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  margin-bottom: 2rem;
}
.cat-name-display h2 {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--font-color-dark);
}
.cat-name-display p {
  font-size: 1.5rem;
  font-weight: 600;
  color: var(--green);
}
.validation-summary {
  background-color: #fef2f2;
  border: 1px solid #ef4444;
  border-radius: 12px;
  padding: 16px;
  margin: 24px 0;
  text-align: center;

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
</style>
