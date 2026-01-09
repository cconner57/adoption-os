<script setup lang="ts">
import { computed, ref, defineProps } from 'vue'

const props = defineProps<{
  data: { date: string; value: number }[]
  color?: string
  height?: number
}>()

const width = 600
const height = props.height || 300
const padding = 60

const hoveredPoint = ref<any>(null)

const sortedData = computed(() => {
  return [...props.data].sort((a, b) => new Date(a.date).getTime() - new Date(b.date).getTime())
})

const chartComputeds = computed(() => {
  if (sortedData.value.length === 0) return null

  const dates = sortedData.value.map((d) => new Date(d.date).getTime())
  const values = sortedData.value.map((d) => d.value)

  const minDate = Math.min(...dates)
  const maxDate = Math.max(...dates)
  const minValue = Math.min(...values)
  const maxValue = Math.max(...values)

  const dateRange = maxDate - minDate || 1
  const valueRange = maxValue - minValue || 1

  // Buffering
  const yMin = Math.max(0, minValue - valueRange * 0.2) // Ensure we don't go below 0 for weight
  const yMax = maxValue + valueRange * 0.1
  const yRange = yMax - yMin

  const getX = (dateStr: string) => {
    return padding + ((new Date(dateStr).getTime() - minDate) / dateRange) * (width - padding * 2)
  }

  const getY = (val: number) => {
    return height - padding - ((val - yMin) / yRange) * (height - padding * 2)
  }

  const pointsStr = sortedData.value.map((d) => `${getX(d.date)},${getY(d.value)}`).join(' ')

  // Area path for gradient (close the loop at the bottom)
  const areaPath = `M ${getX(sortedData.value[0].date)},${height - padding} L ${pointsStr} L ${getX(sortedData.value[sortedData.value.length - 1].date)},${height - padding} Z`

  // Generate Y Axis Labels (5 ticks)
  const yTicks = []
  for (let i = 0; i <= 4; i++) {
    const val = yMin + yRange * (i / 4)
    yTicks.push({
      value: val,
      y: getY(val),
    })
  }

  // Generate X Axis Labels (Time-based for even spacing)
  const xTicks = []
  const numXTicks = 5
  for (let i = 0; i < numXTicks; i++) {
    const time = minDate + dateRange * (i / (numXTicks - 1))
    // Round to nearest day to avoid weird fraction times if needed,
    // but formatted date usually handles it well.
    const dateStr = new Date(time).toISOString()
    xTicks.push({
      date: dateStr,
      x: getX(dateStr),
    })
  }

  const interactivePoints = sortedData.value.map((d) => ({
    x: getX(d.date),
    y: getY(d.value),
    value: d.value,
    date: d.date,
  }))

  return {
    linePath: `M ${pointsStr}`,
    areaPath,
    yTicks,
    xTicks,
    interactivePoints,
  }
})

const formatDate = (d: string) =>
  new Date(d).toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
const formatWeight = (w: number) => `${w.toFixed(2)}kg / ${(w * 2.20462).toFixed(2)}lbs`
</script>

<template>
  <div class="chart-container" @mouseleave="hoveredPoint = null">
    <svg
      v-if="chartComputeds"
      :viewBox="`0 0 ${width} ${height}`"
      preserveAspectRatio="xMidYMid meet"
    >
      <defs>
        <linearGradient id="line-gradient" x1="0" x2="0" y1="0" y2="1">
          <stop offset="0%" stop-color="var(--color-secondary)" stop-opacity="0.2" />
          <stop offset="100%" stop-color="var(--color-secondary)" stop-opacity="0" />
        </linearGradient>
      </defs>

      <!-- Grid & Axes -->
      <g class="grid">
        <!-- Y Axis Lines & Text -->
        <g v-for="tick in chartComputeds.yTicks" :key="tick.value">
          <line
            :x1="padding"
            :y1="tick.y"
            :x2="width - padding"
            :y2="tick.y"
            stroke="#f3f4f6"
            stroke-width="1"
          />
          <text :x="padding - 10" :y="tick.y + 4" text-anchor="end" font-size="10" fill="#9ca3af">
            {{ tick.value.toFixed(1) }}kg
          </text>
        </g>
        <!-- X Axis Text -->
        <g v-for="tick in chartComputeds.xTicks" :key="tick.date">
          <text
            :x="tick.x"
            :y="height - padding + 20"
            text-anchor="middle"
            font-size="10"
            fill="#9ca3af"
          >
            {{ formatDate(tick.date) }}
          </text>
        </g>
      </g>

      <!-- Data Area -->
      <path :d="chartComputeds.areaPath" fill="url(#line-gradient)" />

      <!-- Main Line -->
      <path
        :d="chartComputeds.linePath"
        fill="none"
        stroke="var(--color-secondary)"
        stroke-width="3"
        stroke-linecap="round"
        stroke-linejoin="round"
      />

      <!-- Interactive Points -->
      <g v-for="(point, idx) in chartComputeds.interactivePoints" :key="idx">
        <!-- Visible Dot -->
        <circle
          :cx="point.x"
          :cy="point.y"
          r="4"
          fill="white"
          stroke="var(--color-secondary)"
          stroke-width="2"
        />
        <!-- Invisible Hit Area -->
        <circle
          :cx="point.x"
          :cy="point.y"
          r="15"
          fill="transparent"
          @mouseenter="hoveredPoint = point"
          style="cursor: pointer"
        />
      </g>
    </svg>

    <!-- Empty State -->
    <div v-else class="empty-chart">No data available</div>

    <!-- Tooltip -->
    <div
      v-if="hoveredPoint"
      class="tooltip"
      :style="{
        left: `${(hoveredPoint.x / width) * 100}%`,
        top: `${(hoveredPoint.y / height) * 100}%`,
      }"
    >
      <div class="tooltip-date">{{ formatDate(hoveredPoint.date) }}</div>
      <div class="tooltip-val">{{ formatWeight(hoveredPoint.value) }}</div>
    </div>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
  height: 100%;
  position: relative;
  /* Use flex to center empty state */
  display: flex;
  flex-direction: column;
}

svg {
  width: 100%;
  height: 100%;
  overflow: visible;
}

.empty-chart {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: hsl(from var(--color-neutral) h s 50%);
  font-size: 0.9rem;
}

.tooltip {
  position: absolute;
  transform: translate(-50%, -120%);
  background: rgba(0, 0, 0, 0.8);
  color: white;
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 0.8rem;
  pointer-events: none;
  z-index: 10;
  white-space: nowrap;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.tooltip-date {
  opacity: 0.8;
  font-size: 0.7rem;
  margin-bottom: 2px;
}

.tooltip-val {
  font-weight: 600;
}
</style>
