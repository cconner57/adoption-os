<script setup lang="ts">
import { onUnmounted,ref, watch } from 'vue'

const props = defineProps<{
  imageFile: File
  aspectRatio?: number // default 1
}>()

const emit = defineEmits(['crop', 'cancel'])

const containerRef = ref<HTMLElement | null>(null)
const imageRef = ref<HTMLImageElement | null>(null)
const imageUrl = ref('')

// State
const scale = ref(1)
const position = ref({ x: 0, y: 0 })
const isDragging = ref(false)
const dragStart = ref({ x: 0, y: 0 })


// Config
const MIN_SCALE = 1
const MAX_SCALE = 3

function resetState() {
  scale.value = 1
  position.value = { x: 0, y: 0 }
}

// Load image
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





// --- Drag Logic ---
function onMouseMove(e: MouseEvent | TouchEvent) {
  if (!isDragging.value) return
  e.preventDefault() // prevent scroll on touch

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

  // Snap back bounds check could happen here or continuously in draw

  // Simply clamping:
  /*
     Ideally we want the image to fill the crop box.
     Implementation detail: clamp position so edges don't come inside crop box if scale allows covering
  */
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



// --- Crop Logic ---
function crop() {
  if (!imageRef.value || !containerRef.value) return

  const canvas = document.createElement('canvas')
  const ctx = canvas.getContext('2d')
  if (!ctx) return

  // We want the output to be the size of the viewport on screen (or natural resolution equivalent)
  // Let's go for high-ish resolution: 1080x1080 usually good for aspect 1
  const outputWidth = 1200
  const outputHeight = outputWidth / (props.aspectRatio || 1)

  canvas.width = outputWidth
  canvas.height = outputHeight

  // Draw logic
  // We need to map the "Visible Viewport" pixels to the "Natural Image" pixels

  /*
    The generic constraints:
    The container (viewport) is usually fixed size on screen, e.g. 300px x 300px.
    The image is transformed by CSS: scale(S) translate(X, Y).

    We need to reverse this to find which part of the image is under the viewport.
  */

  // 1. Get viewport size in pixels
  const rect = containerRef.value.getBoundingClientRect()
  const vw = rect.width


  // 2. Map coordinates relative to image center/origin?
  // Easier: The image is drawn at `position` with size `naturalWidth * scale` ? No.
  // The CSS usually sets width: 100% or similar.
  // Let's assume the image rendered width is `rect.width` (initial) * scale.
  // Wait, let's look at CSS. Image is typically min-width: 100%, min-height: 100% object-fit: cover?
  // For manual panning, we usually set width/height specifically or use transform on a base size.

  // Let's use computed style for current rendered width/height (unscaled by transform)
  const imgRect = imageRef.value.getBoundingClientRect()
  // imgRect includes the transform scaling!

  // Ratio of Rendering-to-Output
  const renderScale = outputWidth / vw

  ctx.fillStyle = '#ffffff'
  ctx.fillRect(0, 0, outputWidth, outputHeight)

  // Draw the image onto the canvas
  // We can draw the FULL image into the canvas, transformed appropriately
  // The viewport is (0,0) to (vw, vh) in screen coordinates
  // The image is at (imgRect.left, imgRect.top)
  // Relative to viewport: x = imgRect.left - rect.left, y = imgRect.top - rect.top

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
      <!-- Overlay grid lines (optional) -->
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
  background: white;
  border: 1px solid var(--border-color);
  border-radius: 8px;
  padding: 16px;
  margin-top: 16px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
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
  max-width: 400px; /* Limit size on desktop */
  aspect-ratio: 1 / 1;
  background: #f1f5f9;
  margin: 0 auto;
  overflow: hidden;
  position: relative;
  border-radius: 4px;
  user-select: none;
  touch-action: none; /* Important for preventing scroll */
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
  /* Ensure image covers the box initially or fits?
     For cropping, typically we want 'contain' initially so they can zoom in,
     or 'cover' so it's already full. 'contain' is safer index.
  */
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
  /* Actually for transform based pan/zoom we might want auto size?
     If we use object-fit, the transform applies to the container div.
  */
  display: block;
}

.grid-overlay {
  position: absolute;
  inset: 0;
  pointer-events: none;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: inset 0 0 20px rgba(0, 0, 0, 0.1);
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
  font-weight: bold;
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
  background: white;
  border: 1px solid var(--border-color);
  color: var(--text-base);
}

.btn-secondary:hover {
  background: var(--bg-hover);
}

.btn-primary {
  background: var(--color-primary);
  border: 1px solid var(--color-primary);
  color: white;
}

.btn-primary:hover {
  filter: brightness(1.1);
}
</style>
