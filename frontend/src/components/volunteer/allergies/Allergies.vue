<script setup lang="ts">
import ButtonToggle from '../../common/ui/ButtonToggle.vue'

defineOptions({
  inheritAttrs: false,
})

defineProps<{
  modelValue?: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()
</script>

<template>
  <div class="field" v-bind="$attrs">
    <ButtonToggle
      label="Allergies?"
      @update:modelValue="emit('update:modelValue', $event)"
      :options="[
        { label: 'No', value: false },
        { label: 'Yes', value: true },
      ]"
      :value="modelValue"
    />
  </div>
  <div class="field spacer"></div>
</template>

<style scoped lang="css">
.spacer {
  display: none;
}

@media (min-width: 768px) {
  .spacer {
    display: block;
  }
}
</style>
