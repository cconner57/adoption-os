<script setup lang="ts">
const props = defineProps<{
  leftLabel?: string
  rightLabel?: string
  modelValue?: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()
</script>

<template>
  <div class="switch-row">
    <span class="switch-text">{{ leftLabel }}</span>
    <label class="switch">
      <input
        id="allergies"
        type="checkbox"
        class="switch-input"
        :checked="modelValue"
        @change="emit('update:modelValue', ($event.target as HTMLInputElement).checked)"
      />
      <span class="switch__track"></span>
      <span class="switch__thumb"></span>
    </label>

    <span class="switch-text">{{ rightLabel }}</span>
  </div>
</template>

<style scoped lang="css">
.switch-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: 6px;
}
.switch-text {
  color: black;
  font-weight: 600;
}

.switch {
  position: relative;
  display: inline-block;
  width: 46px;
  height: 28px;
  border-radius: 999px;
  cursor: pointer;
}

.switch-input {
  position: absolute;
  inset: 0;
  width: 100%;
  height: 100%;
  opacity: 0;
  margin: 0;
  cursor: pointer;
}

.switch__track {
  position: absolute;
  inset: 0;
  background: #e7ebf0;
  border-radius: 999px;
  transition: background 0.2s ease;
  box-shadow: inset 0 0 0 1px rgba(0, 0, 0, 0.03);
}
.switch__thumb {
  position: absolute;
  top: 3px;
  left: 3px;
  width: 22px;
  height: 22px;
  border-radius: 999px;
  background: #fff;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
  transition: transform 0.2s ease;
}

.switch-input:checked ~ .switch__track {
  background: #00a5ad;
}
.switch-input:checked ~ .switch__thumb {
  transform: translateX(18px);
}

.switch-input:focus-visible ~ .switch__track {
  box-shadow: 0 0 0 3px rgba(0, 165, 173, 0.4);
}
.switch-input:checked + .switch .switch__track {
  background: #00a5ad;
}
.switch-input:checked + .switch .switch__thumb {
  transform: translateX(18px);
}

.switch-input:focus-visible + .switch .switch__track {
  outline: none;
  box-shadow: 0 0 0 3px rgba(0, 165, 173, 0.4);
}
</style>
