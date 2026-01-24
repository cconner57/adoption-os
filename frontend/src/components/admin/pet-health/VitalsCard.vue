<script setup lang="ts">
import type { IHealthLogEntry } from '../../../models/common'

defineProps<{
  latestVitals: IHealthLogEntry | null
}>()

function getStatusColor(status?: string | null) {
  if (!status) return 'gray'
  if (['normal', 'energetic', 'calm'].includes(status)) return 'green'
  if (['low', 'lethargic'].includes(status)) return 'orange'
  if (['none', 'blood', 'diarrhea'].includes(status)) return 'red'
  return 'blue'
}
</script>

<template>
  <div class="stat-card vitals-snapshot">
    <h3>Latest Vitals</h3>

    <div class="vitals-row">
      <div class="vital-item">
        <label>Temp</label>
        <span class="vital-value">{{ latestVitals?.temperature || '-' }}°F</span>
      </div>
      <div class="vital-item">
        <label>Eating</label>
        <span class="status-pill" :class="getStatusColor(latestVitals?.eating)">
          {{ latestVitals?.eating || '-' }}
        </span>
      </div>
      <div class="vital-item">
        <label>Drinking</label>
        <span class="status-pill" :class="getStatusColor(latestVitals?.drinking)">
          {{ latestVitals?.drinking || '-' }}
        </span>
      </div>
    </div>

    <div class="vitals-row mt-4">
      <div class="vital-item">
        <label>Activity</label>
        <span class="status-pill" :class="getStatusColor(latestVitals?.activity)">
          {{ latestVitals?.activity || '-' }}
        </span>
      </div>
      <div class="vital-item">
        <label>Urination</label>
        <span class="status-pill" :class="getStatusColor(latestVitals?.urination)">
          {{ latestVitals?.urination || '-' }}
        </span>
      </div>
      <div class="vital-item">
        <label>Defecation</label>
        <span class="status-pill" :class="getStatusColor(latestVitals?.defecation)">
          {{ latestVitals?.defecation || '-' }}
        </span>
      </div>
    </div>

    <div class="vitals-note mt-6">
      <label>Latest Note</label>
      <div class="note-box">
        <p>{{ latestVitals?.notes || 'No notes available.' }}</p>
        <span class="note-author">— {{ latestVitals?.recordedBy || 'Unknown' }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.stat-card {
  background: var(--text-inverse);
  border-radius: 12px;
  box-shadow: 0 4px 6px rgb(0 0 0 / 5%);
  padding: 24px;
  border: 1px solid var(--border-color);
}

.vitals-snapshot {
  h3 {
    margin: 0 0 16px;
    font-size: 1rem;
    color: hsl(from var(--color-neutral) h s 50%);
  }
}

.vitals-row {
  display: flex;
  justify-content: space-between;
  gap: 12px;
}

.vital-item {
  display: flex;
  flex-direction: column;
  gap: 6px;

  label {
    font-size: 0.75rem;
    text-transform: uppercase;
    color: hsl(from var(--color-neutral) h s 50%);
    font-weight: 600;
  }
}

.vital-value {
  font-size: 1.1rem;
  font-weight: 600;
  color: var(--text-primary);
}

.status-pill {
  padding: 4px 10px;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: capitalize;

  &.green {
    background: hsl(from var(--color-primary) h s 95%);
    color: var(--color-primary);
  }

  &.orange {
    background: hsl(from var(--color-warning) h s 95%);
    color: var(--color-warning);
  }

  &.red {
    background: hsl(from var(--color-danger) h s 95%);
    color: var(--color-danger);
  }

  &.blue {
    background: hsl(from var(--color-secondary) h s 95%);
    color: var(--color-secondary);
  }

  &.gray {
    background: hsl(from var(--color-neutral) h s 95%);
    color: hsl(from var(--color-neutral) h s 50%);
  }
}

.mt-4 {
  margin-top: 16px;
}

.mt-6 {
  margin-top: 24px;
}

.vitals-note {
  label {
    font-size: 0.75rem;
    text-transform: uppercase;
    color: hsl(from var(--color-neutral) h s 50%);
    font-weight: 600;
    display: block;
    margin-bottom: 8px;
  }
}

.note-box {
  background: hsl(from var(--color-neutral) h s 98%);
  border-radius: 8px;
  padding: 12px;
  border: 1px solid var(--border-color);

  p {
    margin: 0 0 8px;
    font-size: 0.9rem;
    color: var(--text-primary);
    line-height: 1.4;
    font-style: italic;
  }
}

.note-author {
  font-size: 0.75rem;
  color: hsl(from var(--color-neutral) h s 50%);
  display: block;
  text-align: right;
  font-weight: 500;
}
</style>
