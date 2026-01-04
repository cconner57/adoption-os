<script setup lang="ts">
import type { FormState } from '../../../models/adopt-form.ts'
import InputField from '../../common/ui/InputField.vue'

defineProps<{
  modelValue: FormState
  touched: Record<string, boolean>
  handleBlur: (field: string) => void
}>()
</script>

<template>
  <div class="summary-section">
    <p>
      Thank you for the time to complete this application. Please note: Filling out this application
      does not guarantee an adoption. We have thoroughly evaluated all of the animals in our care,
      and try very hard to match the right pet to the right person and situation. Adoptions are not
      on a first-come-first-served basis. We reserve the right not to adopt. Please remember that we
      are all volunteers and work regular full-time jobs like everyone else. We will reply to your
      application as soon as it has been processed (normally 3-4 days). Thank you for considering a
      rescue cat!
    </p>

    <p>
      You are making a major commitment when you adopt any pet. Please remember that cats can live
      for up to 20 or more years. Thousands of cats are killed at animals shelters each year because
      their owners did not plan for the future. Cats can get sick and require expensive medical
      treatment during the course of their life. Cats need affection, attention and understanding.
      You may have to adjust your lifestyle to accommodate a new pet. If you are ready to make this
      commitment, please type your name below:
    </p>
    <InputField
      v-model="modelValue.agreementSignature1"
      label="Type your name"
      name="agreementSignature1"
      placeholder="Type your name"
      :hasError="touched.agreementSignature1 && !modelValue.agreementSignature1"
      @blur="handleBlur('agreementSignature1')"
    />
    <p>
      Adoptions require a homecheck/delivery, in the event that my application is accepted, I agree
      and consent to have a volunteer deliver the pet to my home and inspect my indoor and outdoor
      space for safety. All members of the household - including roommates and children - must be
      present for the homecheck/delivery. If you understand and agree, please type your name below:
    </p>
    <InputField
      v-model="modelValue.agreementSignature2"
      label="Type your name"
      name="agreementSignature2"
      placeholder="Type your name"
      :hasError="touched.agreementSignature2 && !modelValue.agreementSignature2"
      @blur="handleBlur('agreementSignature2')"
    />
    <p>
      Please type your name below to indicate that you have gone to our website and read our "Are
      you Ready to Adopt" page.
    </p>
    <InputField
      v-model="modelValue.agreementSignature3"
      label="Type your name"
      name="agreementSignature3"
      placeholder="Type your name"
      :hasError="touched.agreementSignature3 && !modelValue.agreementSignature3"
      @blur="handleBlur('agreementSignature3')"
    />
  </div>
</template>

<style scoped lang="css"></style>
