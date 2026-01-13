<script setup lang="ts">
import { computed, ref } from 'vue'
import type { IPet } from '../../../../../models/common'

const props = defineProps<{
  modelValue: Partial<IPet>
}>()

const emit = defineEmits(['update:modelValue'])

const formData = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
})

const fileInput = ref<HTMLInputElement | null>(null)
function triggerFileInput() {
  fileInput.value?.click()
}

const isDragging = ref(false)

function onDragOver(e: DragEvent) {
  isDragging.value = true
}

function onDragLeave(e: DragEvent) {
  isDragging.value = false
}

function onDrop(e: DragEvent) {
  isDragging.value = false
  const file = e.dataTransfer?.files[0]
  if (file) {
    uploadFile(file)
  }
}

async function uploadFile(file: File) {
  // Validate Pet ID availability
  if (!formData.value.id) {
    alert('Please save the pet first before uploading photos.')
    return
  }

  // Validate Max Photos
  if ((formData.value.photos?.length || 0) >= 5) {
    alert('Maximum of 5 photos allowed.')
    return
  }

  const fd = new FormData()
  fd.append('photo', file)

  try {
    const token = localStorage.getItem('token')
    const response = await fetch(
      `${import.meta.env.VITE_API_URL}/pets/${formData.value.id}/photos`,
      {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: fd,
      },
    )

    if (!response.ok) {
      const txt = await response.text()
      throw new Error(txt || 'Upload failed')
    }

    const data = await response.json() // { url, thumbnailUrl }

    if (!formData.value.photos) formData.value.photos = []

    formData.value.photos.push({
      url: data.url,
      thumbnailUrl: data.thumbnailUrl,
      isPrimary: formData.value.photos.length === 0,
      isSpotlight: false,
      uploadedAt: new Date().toISOString(),
    })
  } catch (error) {
    console.error('Upload error:', error)
    alert('Failed to upload photo')
  }
}

async function onFileSelected(event: Event) {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  if (!file) return
  await uploadFile(file)
  if (fileInput.value) fileInput.value.value = ''
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
        <img
          :src="photo.url.startsWith('http') ? photo.url : `/pet-photos/${photo.url}`"
          alt="Pet photo"
        />
        <button class="remove-photo" @click="removePhoto(index)" type="button">✕</button>

        <div class="photo-actions">
          <button
            type="button"
            class="action-pill"
            :class="{ active: photo.isPrimary }"
            @click="setPrimaryPhoto(index)"
          >
            Main
          </button>
          <button
            type="button"
            class="action-pill"
            :class="{ active: photo.isSpotlight }"
            @click="setSpotlightPhoto(index)"
          >
            Spotlight
          </button>
        </div>
      </div>
      <div
        class="add-photo-btn"
        :class="{ 'is-dragging': isDragging }"
        @click="triggerFileInput"
        @dragover.prevent="onDragOver"
        @dragleave.prevent="onDragLeave"
        @drop.prevent="onDrop"
        v-if="(formData.photos?.length || 0) < 5"
      >
        <span class="plus-icon">+</span>
        <span class="add-text">Add Photo</span>
        <span class="sub-text">{{ isDragging ? 'Drop Here' : 'Click or Drag & Drop' }}</span>
        <span class="count-text">{{ formData.photos?.length || 0 }}/5</span>
        <input
          ref="fileInput"
          type="file"
          accept="image/jpeg,image/png"
          style="display: none"
          @change="onFileSelected"
        />
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
  background: #f1f5f9;
}

.photo-actions {
  height: 30%;
  background: white;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-evenly;
  gap: 4px;
  padding: 0 4px;
  border-top: 1px solid var(--border-color);
}

.action-pill {
  border: 1px solid var(--border-color);
  background: transparent;
  color: var(--text-secondary);
  border-radius: 12px;
  padding: 2px 8px;
  font-size: 0.7rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  flex: 1;
  text-align: center;
}

.action-pill:hover {
  background: hsl(from var(--color-neutral) h s 95%);
}

.action-pill.active {
  background: var(--color-secondary);
  color: white;
  border-color: var(--color-secondary);
}

.remove-photo {
  position: absolute;
  top: 4px;
  right: 4px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.5);
  color: white;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  z-index: 2;
  transition: background 0.2s;
}

.remove-photo:hover {
  background: var(--color-danger);
}

.add-photo-btn {
  aspect-ratio: 1;
  border: 1px dashed var(--border-color);
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: var(--color-primary);
  transition: all 0.2s;
  background: hsla(from var(--color-primary) h s l / 0.02);
  padding: 8px;
  text-align: center;
}

.add-photo-btn:hover,
.add-photo-btn.is-dragging {
  background: hsla(from var(--color-primary) h s l / 0.08);
  border-color: var(--color-primary);
  border-style: solid;
}

.plus-icon {
  font-size: 1.5rem;
  line-height: 1;
  margin-bottom: 4px;
}

.add-text {
  font-weight: 600;
  font-size: 0.9rem;
}

.sub-text {
  font-size: 0.7rem;
  color: var(--text-secondary);
  margin-top: 2px;
}

.count-text {
  font-size: 0.7rem;
  color: var(--text-tertiary);
  margin-top: auto; /* Push to bottom */
}
</style>
