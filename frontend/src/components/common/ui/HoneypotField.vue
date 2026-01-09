<script setup lang="ts">
defineProps<{
  modelValue: string | null
}>()

defineEmits(['update:modelValue'])
</script>

<template>
  <div class="fax-field" aria-hidden="true">
    <label for="fax_number">Fax Number</label>
    <input
      id="fax_number"
      :value="modelValue"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      type="text"
      name="fax_number"
      tabindex="-1"
      autocomplete="off"
    />
  </div>
</template>

<style scoped>
.fax-field {
  opacity: 0;
  position: absolute;
  top: 0;
  left: 0;
  height: 0;
  width: 0;
  z-index: -1;
  overflow: hidden;
}
</style>
