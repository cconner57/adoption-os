<script setup lang="ts">
import { ref } from 'vue'

import MarketingCampaigns from '../../components/admin/marketing/MarketingCampaigns.vue'
import MarketingHappyTails from '../../components/admin/marketing/MarketingHappyTails.vue'
import MarketingNewsletters from '../../components/admin/marketing/MarketingNewsletters.vue'
import { Icon } from '../../components/common/ui'
import type { ICampaign } from '../../stores/mockMarketing'

const activeTab = ref<'campaigns' | 'newsletters' | 'stories'>('campaigns')
const selectedCampaign = ref<ICampaign | null>(null)
</script>

<template>
  <div class="marketing-page">
    <div class="page-header">
       <div v-if="selectedCampaign" class="header-breadcrumb">
        <button class="back-link" @click="selectedCampaign = null">← Marketing Center</button>
      </div>
      <h1>{{ selectedCampaign ? selectedCampaign.name : 'Marketing Center' }}</h1>

      <div v-if="!selectedCampaign" class="tabs">
        <button class="tab-btn" :class="{ active: activeTab === 'campaigns' }" @click="activeTab = 'campaigns'">
          <Icon name="megaphone" size="18" />
          Campaigns
        </button>
        <button class="tab-btn" :class="{ active: activeTab === 'newsletters' }" @click="activeTab = 'newsletters'">
          <Icon name="mail" size="18" />
          Newsletters
        </button>
        <button class="tab-btn" :class="{ active: activeTab === 'stories' }" @click="activeTab = 'stories'">
          <Icon name="paw" size="18" />
          Happy Tails
        </button>
      </div>
    </div>

    <!-- Tab Content -->
    <MarketingCampaigns
      v-if="activeTab === 'campaigns'"
      v-model="selectedCampaign"
    />

    <MarketingNewsletters
      v-if="activeTab === 'newsletters'"
    />

    <MarketingHappyTails
      v-if="activeTab === 'stories'"
    />
  </div>
</template>

<style scoped>
.marketing-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.page-header {
  display: flex;
  flex-direction: column;
  gap: 16px;

  h1 {
    margin: 0;
    font-size: 1.8rem;
    color: var(--text-primary);
  }
}

.back-link {
  background: none;
  border: none;
  color: var(--color-secondary);
  cursor: pointer;
  padding: 0;
  font-weight: 600;
  margin-bottom: 8px;
  &:hover { text-decoration: underline; }
}

.tabs {
  display: flex;
  gap: 24px;
  border-bottom: 1px solid #e5e7eb;
}

.tab-btn {
  background: none;
  border: none;
  padding: 12px 4px;
  font-size: 1rem;
  display: flex;
  align-items: center;
  gap: 8px;
  color: hsl(from var(--color-neutral) h s 50%);
  cursor: pointer;
  border-bottom: 2px solid transparent;
  font-weight: 500;

  &.active {
    color: var(--text-primary);
    border-bottom-color: var(--color-secondary);
    font-weight: 700;
  }
}
</style>
