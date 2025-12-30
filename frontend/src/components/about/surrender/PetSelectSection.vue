<script setup lang="ts">

const { formError, selectedAnimal } = defineProps<{
  formError: boolean
  selectedAnimal: 'dog' | 'cat' | null
}>()

const emit = defineEmits(['update:selectedAnimal'])

const updateSelectedAnimal = (value: 'dog' | 'cat') => {
  emit('update:selectedAnimal', value)
}
</script>

<template>
  <h5>Will you be surrendering Dog(s) or Cat(s)?</h5>
  <div class="times">
    <label class="time-card">
      <input
        type="radio"
        name="animal"
        value="dog"
        :checked="selectedAnimal === 'dog'"
        @change="updateSelectedAnimal('dog')"
      />
      <div class="time-card__content">
        <strong>Dog(s)</strong>
        <small>Select if you are surrendering dog(s) or puppies</small>
      </div>
    </label>

    <label class="time-card">
      <input
        type="radio"
        name="animal"
        value="cat"
        :checked="selectedAnimal === 'cat'"
        @change="updateSelectedAnimal('cat')"
      />
      <div class="time-card__content">
        <strong>Cat(s)</strong>
        <small>Select if you are surrendering cat(s) or kittens</small>
      </div>
    </label>
  </div>
  <p v-if="formError" class="error">Please select an option above</p>
</template>

<style scoped lang="css">
h5 {
  margin-top: 3rem;
  margin-bottom: 12px;
}
.times {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 12px;
  margin-bottom: 4rem;
}
.time-card {
  position: relative;
  display: flex;
  align-items: flex-start;
  gap: 10px;
  padding: 12px;
  border: 1px solid #e7ebf0;
  border-radius: 12px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  cursor: pointer;
  user-select: none;
  transition:
    background 0.2s,
    border-color 0.2s,
    box-shadow 0.2s;
}
.time-card > input {
  position: absolute;
  opacity: 0;
  width: 1px;
  height: 1px;
  pointer-events: none;
}
.time-card:has(> input:checked) {
  background: color-mix(in srgb, var(--green) 10%, white);
  border-color: var(--green);
  box-shadow: 0 0 0 1px var(--green) inset;
}
.time-card:has(> input:focus-visible) {
  box-shadow: 0 0 0 3px rgba(30, 99, 217, 0.45);
}
.time-card__content small {
  color: #6b7b8a;
  display: block;
  margin-top: 4px;
}
@supports not (selector(:has(*))) {
  .time-card > input:checked + .time-card__content {
    background: #e8f1ff;
    border-radius: 10px;
    margin: -12px;
    padding: 12px;
    box-shadow: 0 0 0 2px #bfd0ff inset;
  }
}
.error {
  color: var(--red);
  font-size: 1.5rem;
  margin-top: -2rem;
}
</style>
