<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const { label, hasError } = defineProps<{
  label?: string
  modelValue?: string | null
  hasError?: boolean
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string | null]
}>()

const canvasRef = ref<HTMLCanvasElement | null>(null)
const isDrawing = ref(false)
const signatureData = ref<string | null>(null)

let context: CanvasRenderingContext2D | null = null
const getEventPosition = (event: MouseEvent | TouchEvent) => {
  if (canvasRef.value) {
    const rect = canvasRef.value.getBoundingClientRect()
    const dpr = window.devicePixelRatio || 1 // Get the device pixel ratio
    if (event instanceof MouseEvent) {
      return {
        offsetX: (event.clientX - rect.left) / dpr, // Adjust for scaling
        offsetY: (event.clientY - rect.top) / dpr, // Adjust for scaling
      }
    } else {
      const touch = event.touches[0]
      return {
        offsetX: (touch.clientX - rect.left) / dpr, // Adjust for scaling
        offsetY: (touch.clientY - rect.top) / dpr, // Adjust for scaling
      }
    }
  }
  return { offsetX: 0, offsetY: 0 }
}

const startDrawing = (event: MouseEvent | TouchEvent) => {
  isDrawing.value = true
  const { offsetX, offsetY } = getEventPosition(event)
  if (context) {
    context.beginPath()
    context.moveTo(offsetX, offsetY)
  }
}

const draw = (event: MouseEvent | TouchEvent) => {
  if (!isDrawing.value || !context) return
  const { offsetX, offsetY } = getEventPosition(event)
  context.lineTo(offsetX, offsetY)
  context.stroke()
}

const stopDrawing = () => {
  isDrawing.value = false
  if (context) {
    context.closePath()
    const data = canvasRef.value?.toDataURL('image/png') || null
    signatureData.value = data
    emit('update:modelValue', data)
  }
}

const clearCanvas = () => {
  if (context && canvasRef.value) {
    context.fillStyle = '#ffffff'
    context.fillRect(0, 0, canvasRef.value.width, canvasRef.value.height)
    signatureData.value = null
    emit('update:modelValue', null)
  }
}

const scaleCanvas = () => {
  if (canvasRef.value) {
    const rect = canvasRef.value.getBoundingClientRect()
    if (rect.width === 0 || rect.height === 0) return // Don't scale if hidden

    const dpr = window.devicePixelRatio || 1
    canvasRef.value.width = rect.width * dpr
    canvasRef.value.height = rect.height * dpr

    context = canvasRef.value.getContext('2d')
    if (context) {
      context.scale(dpr, dpr)
      context.lineWidth = 2
      context.lineCap = 'round'
      context.strokeStyle = '#000'
      // Fill with white initially (transparent by default in some browsers, but white ensures contrast)
      // Note: We don't fillRect here because it might clear user data if called on resize.
      // Ideally we should persist data, but for now let's just ensure it's writable.
      // If we want to persist, we'd need to save the image data and put it back.
      // For this simple fix, we just want it to work when it becomes visible.
    }
  }
}

let resizeObserver: ResizeObserver | null = null

onMounted(() => {
  if (canvasRef.value) {
    resizeObserver = new ResizeObserver(() => {
      // Debounce or just call it?
      // Just calling it is fine for v-show toggles.
      // Check if context is already valid and has dimensions to avoid clearing if it's just a minor layout shift?
      // For v-show, width goes 0 -> actual.
      // If width goes actual -> actual (minor shift), we might clear content.
      // Let's only scale if internal width is 0 or mismatch significantly.
      scaleCanvas()
    })
    resizeObserver.observe(canvasRef.value)
  }
})

onUnmounted(() => {
  if (resizeObserver) {
    resizeObserver.disconnect()
  }
})
</script>

<template>
  <div class="signature-container" :class="{ 'has-error': hasError }">
    <label for="signatureCanvas" class="signature-label">{{ label || 'Please sign below:' }}</label>
    <canvas
      id="signatureCanvas"
      ref="canvasRef"
      class="signature-canvas"
      width="500"
      height="200"
      @mousedown="startDrawing"
      @mousemove="draw"
      @mouseup="stopDrawing"
      @mouseleave="stopDrawing"
      @touchstart.prevent="startDrawing"
      @touchmove.prevent="draw"
      @touchend.prevent="stopDrawing"
    ></canvas>
    <div class="signature-actions">
      <button @click.prevent="clearCanvas">Clear</button>
    </div>
  </div>
</template>

<style scoped lang="css">
.signature-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
  width: 100%;
}

.signature-label {
  margin-bottom: 10px;
  font-size: 16px;
  font-weight: bold;
  color: #333;
}

.signature-canvas {
  border: 1px solid #ccc;
  border-radius: 4px;
  cursor: crosshair;
  background-color: #fff;
  width: 100%;
  height: 200px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  touch-action: none; /* Critical for mobile scrolling prevention */
}

.has-error .signature-canvas {
  border: 2px solid #ef4444;
}

.signature-actions {
  position: absolute;
  bottom: 10px;
  right: 10px;
}

button {
  padding: 5px 10px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
