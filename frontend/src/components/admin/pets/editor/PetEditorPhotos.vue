<script setup lang="ts">
import { computed, ref } from 'vue'
import type { IPet } from '../../../../../models/common'
import { Toast, ImageCropper } from '../../../common/ui'

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

const isUploading = ref(false)
const showToast = ref(false)
const toastMessage = ref('')
const toastType = ref<'success' | 'error'>('error')

// --- Cropping State ---
const pendingFile = ref<File | null>(null)

function onCancelCrop() {
  pendingFile.value = null
  if (fileInput.value) fileInput.value.value = ''
}

async function onCropComplete(croppedBlob: Blob) {
  // Convert blob to File
  const file = new File([croppedBlob], pendingFile.value?.name || 'photo.jpg', {
    type: 'image/jpeg',
  })

  await uploadFile(file)
  pendingFile.value = null
  if (fileInput.value) fileInput.value.value = ''
}

async function uploadFile(file: File) {
  // Validate Pet ID availability
  if (!formData.value.id) {
    toastMessage.value = 'Please save the pet first before uploading photos.'
    toastType.value = 'error'
    showToast.value = true
    return
  }

  // note: max photos check moved to selection time or here?
  // If we check here, user might crop then get rejected. Better to check before crop.
  // But redundant check is fine safety.

  isUploading.value = true

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

    const isFirstPhoto = formData.value.photos.length === 0

    formData.value.photos.push({
      url: data.url,
      thumbnailUrl: data.thumbnailUrl,
      isPrimary: isFirstPhoto,
      isSpotlight: isFirstPhoto,
      uploadedAt: new Date().toISOString(),
    })

    toastMessage.value = 'Photo uploaded successfully'
    toastType.value = 'success'
    showToast.value = true
  } catch (error: any) {
    console.error('Upload error:', error)
    toastMessage.value = error.message || 'Failed to upload photo'
    toastType.value = 'error'
    showToast.value = true
  } finally {
    isUploading.value = false
  }
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
    handleFileSelection(file)
  }
}

async function onFileSelected(event: Event) {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  if (!file) return
  handleFileSelection(file)
}

function handleFileSelection(file: File) {
  // Validate Max Photos BEFORE cropping
  if ((formData.value.photos?.length || 0) >= 5) {
    toastMessage.value = 'Maximum of 5 photos allowed.'
    toastType.value = 'error'
    showToast.value = true
    if (fileInput.value) fileInput.value.value = ''
    return
  }

  pendingFile.value = file
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
  formData.value.photos?.splice(index, 1) // already removed

  // If only 1 photo remains, enforce Main + Spotlight
  if (formData.value.photos?.length === 1) {
    formData.value.photos[0].isPrimary = true
    formData.value.photos[0].isSpotlight = true
  }
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

      <!-- Hide Add Button if Cropping -->
      <div
        class="add-photo-btn"
        :class="{ 'is-dragging': isDragging, 'is-loading': isUploading }"
        @click="!isUploading && !pendingFile && triggerFileInput()"
        @dragover.prevent="onDragOver"
        @dragleave.prevent="onDragLeave"
        @drop.prevent="onDrop"
        v-if="(formData.photos?.length || 0) < 5 && !pendingFile"
      >
        <div v-if="isUploading" class="spinner"></div>
        <template v-else>
          <span class="plus-icon">+</span>
          <span class="add-text">Add Photo</span>
          <span class="sub-text">{{ isDragging ? 'Drop Here' : 'Click or Drag & Drop' }}</span>
          <span class="count-text">{{ formData.photos?.length || 0 }}/5</span>
        </template>
        <input
          ref="fileInput"
          type="file"
          accept="image/*"
          style="display: none"
          @change="onFileSelected"
        />
      </div>
    </div>

    <!-- Inline Cropper -->
    <ImageCropper
      v-if="pendingFile"
      :imageFile="pendingFile"
      @cancel="onCancelCrop"
      @crop="onCropComplete"
    />

    <Toast :show="showToast" :message="toastMessage" :type="toastType" @close="showToast = false" />
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

.add-photo-btn.is-loading {
  cursor: wait;
  opacity: 0.7;
}

.spinner {
  width: 30px;
  height: 30px;
  border: 3px solid rgba(0, 0, 0, 0.1);
  border-radius: 50%;
  border-top-color: var(--primary-color, #3b82f6);
  animation: spin 1s ease-in-out infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
