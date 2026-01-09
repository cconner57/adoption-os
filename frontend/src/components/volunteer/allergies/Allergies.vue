<script setup lang="ts">
import ButtonToggle from '../../common/ui/ButtonToggle.vue'

defineOptions({
  inheritAttrs: false,
})

defineProps<{
  modelValue?: boolean | null
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()
</script>

<template>
  <div class="field" v-bind="$attrs">
    <ButtonToggle
      label="Allergies?"
      :modelValue="modelValue"
      @update:modelValue="emit('update:modelValue', $event as boolean)"
      :true-value="true"
      :false-value="false"
      true-label="Yes"
      false-label="No"
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
