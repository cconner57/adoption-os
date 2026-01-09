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
    if (event instanceof MouseEvent) {
      return {
        offsetX: event.clientX - rect.left,

        offsetY: event.clientY - rect.top, // Adjust for scaling
      }
    } else {
      const touch = event.touches[0]
      return {
        offsetX: touch.clientX - rect.left,
        offsetY: touch.clientY - rect.top,
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
    if (rect.width === 0 || rect.height === 0) return

    const dpr = window.devicePixelRatio || 1
    canvasRef.value.width = rect.width * dpr
    canvasRef.value.height = rect.height * dpr

    context = canvasRef.value.getContext('2d')
    if (context) {
      context.scale(dpr, dpr)
      context.lineWidth = 2
      context.lineCap = 'round'
      context.strokeStyle = '#000'
      context.lineCap = 'round'
      context.strokeStyle = '#000'
    }
  }
}

let resizeObserver: ResizeObserver | null = null

onMounted(() => {
  if (canvasRef.value) {
    resizeObserver = new ResizeObserver(() => {
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
      role="img"
      :aria-label="label || 'Signature canvas'"
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

  .signature-label {
    margin-bottom: 10px;
    font-size: 16px;
    font-weight: bold;
    color: var(--text-primary);
  }

  .signature-canvas {
    border: 1px solid var(--border-color);
    border-radius: 4px;
    cursor: crosshair;
    background-color: #ffffff;
    width: 100%;
    height: 200px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    touch-action: none;
  }

  /* Nested state modifier */
  &.has-error .signature-canvas {
    border: 2px solid var(--color-danger);
  }

  .signature-actions {
    position: absolute;
    bottom: 10px;
    right: 10px;

    button {
      padding: 5px 10px;
      background-color: var(--color-secondary);
      color: var(--text-inverse);
      border: none;
      border-radius: 4px;
      cursor: pointer;

      &:hover {
        background-color: hsl(from var(--color-secondary) h s 40%);
      }
    }
  }
}
</style>
