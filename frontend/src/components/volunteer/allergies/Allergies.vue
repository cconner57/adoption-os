<script setup lang="ts">
import Toggle from '../../common/ui/Toggle.vue'

defineProps<{
  modelValue?: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()
</script>

<template>
  <div class="field col-span-2">
    <label class="label">Allergies?</label>
    <Toggle
      leftLabel="No"
      rightLabel="Yes"
      :modelValue="modelValue"
      @update:modelValue="emit('update:modelValue', $event)"
    />
  </div>
</template>

<style scoped lang="css">
.col-span-2 {
  grid-column: span 2;
}
</style>
