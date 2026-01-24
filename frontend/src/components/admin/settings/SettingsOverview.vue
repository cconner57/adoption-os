

import { ButtonToggle } from '../../common/ui'
import SettingsCard from './SettingsCard.vue'

defineProps<{
  settings: {
    overview: {
      showRecentActivity: boolean
      showPendingTasks: boolean
      showStatsGraph: boolean
    }
  }
}>()

defineEmits<{
  'update:settings': [value: Record<string, unknown>]
}>()
</script>

<template>
  <div class="settings-grid">
    <SettingsCard title="Dashboard Configuration" icon="📊">
      <div class="setting-row">
        <div class="setting-info">
          <label>Show Recent Activity</label>
          <p class="setting-desc">Display recent system actions on the dashboard.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle
            label=""
            v-model="settings.overview.showRecentActivity"
            true-label="Show"
            false-label="Hide"
          />
        </div>
      </div>
      <div class="setting-row">
        <div class="setting-info">
          <label>Show Pending Tasks</label>
          <p class="setting-desc">Show your to-do list widget.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle
            label=""
            v-model="settings.overview.showPendingTasks"
            true-label="Show"
            false-label="Hide"
          />
        </div>
      </div>
      <div class="setting-row">
        <div class="setting-info">
          <label>Show Stats Graph</label>
          <p class="setting-desc">Display the monthly adoption trends graph.</p>
        </div>
        <div class="toggle-wrapper">
          <ButtonToggle
            label=""
            v-model="settings.overview.showStatsGraph"
            true-label="Show"
            false-label="Hide"
          />
        </div>
      </div>
    </SettingsCard>
  </div>
</template>

<style scoped>
.settings-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 24px;
  align-items: start;
}

.setting-row {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  padding: 16px 0;
  border-bottom: 1px solid #f3f4f6;
}

.setting-row:last-child {
  border-bottom: none;
}

.setting-info label {
  display: block;
  font-weight: 500;
  color: var(--text-primary);
  margin-bottom: 4px;
}

.setting-desc {
  font-size: 0.85rem;
  color: #6b7280;
  margin: 0;
  line-height: 1.4;
}
</style>
