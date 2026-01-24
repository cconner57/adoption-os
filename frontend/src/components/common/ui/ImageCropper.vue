<script setup lang="ts">
import { onUnmounted,ref, watch } from 'vue'

const props = defineProps<{
  imageFile: File
  aspectRatio?: number 
}>()

const emit = defineEmits(['crop', 'cancel'])

const containerRef = ref<HTMLElement | null>(null)
const imageRef = ref<HTMLImageElement | null>(null)
const imageUrl = ref('')

const scale = ref(1)
const position = ref({ x: 0, y: 0 })
const isDragging = ref(false)
const dragStart = ref({ x: 0, y: 0 })

const MIN_SCALE = 1
const MAX_SCALE = 3

function resetState() {
  scale.value = 1
  position.value = { x: 0, y: 0 }
}

function onImageLoad() {
  resetState()
}

watch(
  () => props.imageFile,
  (file) => {
    if (file) {
      if (imageUrl.value) URL.revokeObjectURL(imageUrl.value)
      imageUrl.value = URL.createObjectURL(file)
      resetState()
    }
  },
  { immediate: true },
)

function onMouseMove(e: MouseEvent | TouchEvent) {
  if (!isDragging.value) return
  e.preventDefault() 

  const clientX = 'touches' in e ? e.touches[0].clientX : e.clientX
  const clientY = 'touches' in e ? e.touches[0].clientY : e.clientY

  const newX = clientX - dragStart.value.x
  const newY = clientY - dragStart.value.y

  position.value = { x: newX, y: newY }
}

function onMouseUp() {
  isDragging.value = false
  window.removeEventListener('mousemove', onMouseMove)
  window.removeEventListener('mouseup', onMouseUp)
  window.removeEventListener('touchmove', onMouseMove)
  window.removeEventListener('touchend', onMouseUp)

}

function onMouseDown(e: MouseEvent | TouchEvent) {
  e.preventDefault()
  isDragging.value = true
  const clientX = 'touches' in e ? e.touches[0].clientX : e.clientX
  const clientY = 'touches' in e ? e.touches[0].clientY : e.clientY
  dragStart.value = { x: clientX - position.value.x, y: clientY - position.value.y }

  window.addEventListener('mousemove', onMouseMove)
  window.addEventListener('mouseup', onMouseUp)
  window.addEventListener('touchmove', onMouseMove)
  window.addEventListener('touchend', onMouseUp)
}

function crop() {
  if (!imageRef.value || !containerRef.value) return

  const canvas = document.createElement('canvas')
  const ctx = canvas.getContext('2d')
  if (!ctx) return

  const outputWidth = 1200
  const outputHeight = outputWidth / (props.aspectRatio || 1)

  canvas.width = outputWidth
  canvas.height = outputHeight

  const rect = containerRef.value.getBoundingClientRect()
  const vw = rect.width

  const imgRect = imageRef.value.getBoundingClientRect()
  
  const renderScale = outputWidth / vw

  ctx.fillStyle = '#ffffff'
  ctx.fillRect(0, 0, outputWidth, outputHeight)

  const relativeX = (imgRect.left - rect.left) * renderScale
  const relativeY = (imgRect.top - rect.top) * renderScale
  const drawW = imgRect.width * renderScale
  const drawH = imgRect.height * renderScale

  ctx.drawImage(imageRef.value, relativeX, relativeY, drawW, drawH)

  canvas.toBlob(
    (blob) => {
      if (blob) emit('crop', blob)
    },
    'image/jpeg',
    0.9,
  )
}

onUnmounted(() => {
  if (imageUrl.value) URL.revokeObjectURL(imageUrl.value)
})
</script>

<template>
  <div class="cropper-container">
    <div class="cropper-header">
      <h3>Crop Photo</h3>
      <p>Drag to reposition, slider to zoom</p>
    </div>

    <div class="viewport-wrapper" ref="containerRef">
      <div
        class="image-layer"
        :style="{
          transform: `translate(${position.x}px, ${position.y}px) scale(${scale})`,
          cursor: isDragging ? 'grabbing' : 'grab',
        }"
        @mousedown="onMouseDown"
        @touchstart="onMouseDown"
      >
        <img :src="imageUrl" ref="imageRef" @load="onImageLoad" draggable="false" />
      </div>
      
      <div class="grid-overlay"></div>
    </div>

    <div class="controls">
      <span class="zoom-icon">-</span>
      <input
        type="range"
        :min="MIN_SCALE"
        :max="MAX_SCALE"
        step="0.01"
        v-model.number="scale"
        class="zoom-slider"
      />
      <span class="zoom-icon">+</span>
    </div>

    <div class="actions">
      <button class="btn btn-secondary" @click="$emit('cancel')">Cancel</button>
      <button class="btn btn-primary" @click="crop">Crop Photo</button>
    </div>
  </div>
</template>

<style scoped>
.cropper-container {
  background: #fff;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  padding: 16px;
  margin-top: 16px;
  box-shadow: 0 4px 6px -1px rgb(0 0 0 / 10%);
}

.cropper-header h3 {
  margin: 0;
  font-size: 1rem;
  font-weight: 600;
}

.cropper-header p {
  margin: 4px 0 12px;
  color: var(--text-secondary);
  font-size: 0.8rem;
}

.viewport-wrapper {
  width: 100%;
  max-width: 400px; 
  aspect-ratio: 1 / 1;
  background: #f1f5f9;
  margin: 0 auto;
  overflow: hidden;
  position: relative;
  border-radius: 4px;
  user-select: none;
  touch-action: none; 
}

.image-layer {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  transform-origin: center center;
}

.image-layer img {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
  display: block;
}

.grid-overlay {
  position: absolute;
  inset: 0;
  pointer-events: none;
  border: 1px solid rgb(255 255 255 / 30%);
  box-shadow: inset 0 0 20px rgb(0 0 0 / 10%);
}

.controls {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-top: 16px;
  margin-bottom: 16px;
}

.zoom-slider {
  flex: 1;
  max-width: 200px;
  accent-color: var(--color-primary);
}

.zoom-icon {
  color: var(--text-secondary);
  font-weight: 700;
}

.actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  border-top: 1px solid var(--border-color);
  padding-top: 16px;
}

.btn {
  padding: 8px 16px;
  border-radius: 6px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-secondary {
  background: #fff;
  border: 1px solid var(--border-color);
  color: var(--text-base);
}

.btn-secondary:hover {
  background: var(--bg-hover);
}

.btn-primary {
  background: var(--color-primary);
  border: 1px solid var(--color-primary);
  color: #fff;
}

.btn-primary:hover {
  filter: brightness(1.1);
}
</style>
