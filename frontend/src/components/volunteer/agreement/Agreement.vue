<script setup lang="ts">

import InputField from '../../common/ui/InputField.vue'
import InputSignature from '../../common/ui/InputSignature.vue'

const {
  name,
  fullName,
  age,
  signature,
  signatureDate,
  parentName,
  parentSignature,
  parentDate,

} = defineProps<{
  age: number | null
  fullName: string
  name: string
  parentDate: string
  parentName: string
  parentSignature: string | null
  signature: string | null
  signatureDate: string

}>()

const emit = defineEmits<{
  'update:fullName': [value: string]
  'update:parentName': [value: string]
}>()
</script>

<template>
  <fieldset class="waiver-container">
    <legend id="waiv" class="section-title">Agreement</legend>
    <p class="waiver">
      I, {{ name === ' ' ? '(volunteer name)' : name }}, hereby volunteer to assist in various tasks
      to support IDOHR. I understand that IDOHR and partners are not responsible for any illness or
      injury caused during volunteer work. I agree to hold harmless IDOHR and partners should I
      become sick or injured from any animals as a result of my volunteer work.
    </p>

    <div class="acknowledgement">
      <div class="name-date-container">
        <InputField
          label="Name"
          placeholder=""
          :modelValue="fullName"
          @update:modelValue="(val) => emit('update:fullName', String(val))"
        />
        <InputField label="Date" placeholder="" type="date" :modelValue="signatureDate" />
      </div>
      <InputSignature label="Signature" placeholder="" :modelValue="signature" />
    </div>

    <label v-if="age !== null && age < 21" for="parental-consent" class="label"
      >If under 21, I ({{ parentName === '' ? 'parent/guardian name' : parentName }}) give
      permission for my child to volunteer with IDOHR and agree to the above waiver.</label
    >
    <div v-if="age !== null && age < 21" class="parentGuardian">
      <div class="name-date-container">
        <InputField
          label="Parent/Guardian Name"
          placeholder=""
          :modelValue="parentName"
          @update:modelValue="(val) => emit('update:parentName', String(val))"
        />
        <InputField label="Date" placeholder="" type="date" :modelValue="parentDate" />
      </div>
      <InputSignature
        label="Parent/Guardian Signature"
        placeholder=""
        :modelValue="parentSignature"
      />
    </div>
  </fieldset>
</template>

<style scoped lang="css">
.waiver-container {
  margin-bottom: 24px;
  width: 100%;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 12px;

  & .acknowledgement,
  .parentGuardian {
    display: flex;
    flex-direction: column;
    grid-column: 1 / -1;

    & .name-date-container {
      display: flex;
      gap: 12px;
      & > :nth-child(1) {
        flex: 1;
      }
      & > :nth-child(2) {
        flex: 0 0 33%;
      }
    }
    @media (max-width: 440px) {
      .name-date-container {
        flex-direction: column;
        gap: 0px;
        & > :nth-child(2) {
          flex: none;
        }
      }
    }
  }

  label {
    border-top: 2px solid gray;
    padding-top: 12px;
    margin-top: 12px;
    font-weight: 600;
    grid-column: 1 / -1;
    text-align: center;
  }

  @media (max-width: 440px) {
    grid-template-columns: repeat(1, minmax(0, 1fr));
    gap: 8px;
  }
}

.waiver {
  color: black;
  background: #fff;
  border: 1px dashed #cfe2ff;
  padding: 12px;
  border-radius: 10px;
  font-size: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  line-height: 1.4;
  margin-bottom: 12px;
  width: 100%;
  max-width: 100%;
  grid-column: 1 / -1;

  @media (max-width: 440px) {
    font-size: 0.9rem;
  }
}
</style>
