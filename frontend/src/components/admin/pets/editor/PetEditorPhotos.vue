<script setup lang="ts">
import { computed } from 'vue'
import type { IPet } from '../../../../../models/common'

const props = defineProps<{
  modelValue: Partial<IPet>
}>()

const emit = defineEmits(['update:modelValue'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})

function handlePhotoUpload() {
  if (!formData.value.photos) formData.value.photos = []

  if (formData.value.photos.length >= 5) {
    alert('Maximum 5 photos allowed.')
    return
  }

  const mockPhoto = {
    url: 'https://placekitten.com/300/300',
    thumbnailUrl: 'https://placekitten.com/100/100',
    isPrimary: formData.value.photos.length === 0,
    isSpotlight: false,
    uploadedAt: new Date().toISOString(),
  }
  formData.value.photos.push(mockPhoto)
}

function setPrimaryPhoto(index: number) {
  if (!formData.value.photos) return
  formData.value.photos.forEach((p, i) => {
    p.isPrimary = i === index
  })
}

function setSpotlightPhoto(index: number) {
  if (!formData.value.photos) return
  formData.value.photos.forEach((p, i) => {
    p.isSpotlight = i === index
  })
}

function removePhoto(index: number) {
  formData.value.photos?.splice(index, 1)
}
</script>

<template>
  <div class="form-section">
    <div class="photos-grid">
      <div class="photo-card" v-for="(photo, index) in formData.photos" :key="index">
        <img :src="photo.url" alt="Pet photo" />
        <button class="remove-photo" @click="removePhoto(index)" type="button">✕</button>

        <div class="photo-actions">
          <label class="photo-radio">
            <input
              type="radio"
              :name="'primary-photo'"
              :checked="photo.isPrimary"
              @change="setPrimaryPhoto(index)"
            />
            <span>Main</span>
          </label>
          <label class="photo-radio">
            <input
              type="radio"
              :name="'spotlight-photo'"
              :checked="photo.isSpotlight"
              @change="setSpotlightPhoto(index)"
            />
            <span>Spotlight</span>
          </label>
        </div>
      </div>
      <div
        class="add-photo-btn"
        @click="handlePhotoUpload"
        v-if="(formData.photos?.length || 0) < 5"
      >
        <span>+ Add Photo</span>
        <span style="font-size: 0.8rem; margin-top: 4px">{{ formData.photos?.length || 0 }}/5</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import './form.css';

.photos-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr)); /* Better responsiveness */
  gap: 16px;
}

.photo-card {
  position: relative;
  width: 100%;
  aspect-ratio: 1;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid var(--border-color);
  display: flex;
  flex-direction: column;
}

.photo-card img {
  width: 100%;
  height: 70%; /* Give space for controls */
  object-fit: cover;
}

.photo-actions {
  height: 30%;
  background: white;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 4px;
  padding: 0 8px;
  font-size: 0.8rem;
  border-top: 1px solid var(--border-color);
}

.photo-radio {
  display: flex;
  align-items: center;
  gap: 6px;
  cursor: pointer;
}

.photo-radio input {
  cursor: pointer;
}

.remove-photo {
  position: absolute;
  top: 4px;
  right: 4px;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.6);
  color: white;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  z-index: 2; /* Ensure on top */
}

.remove-photo:hover {
  background: rgba(220, 38, 38, 0.9);
}

.add-photo-btn {
  aspect-ratio: 1;
  border: 2px dashed var(--border-color);
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: var(--color-primary);
  transition: all 0.2s;
  background: hsla(from var(--color-primary) h s l / 0.05);
}

.add-photo-btn:hover {
  background: hsla(from var(--color-primary) h s l / 0.1);
}
</style>
