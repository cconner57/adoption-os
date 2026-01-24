<script setup lang="ts">
import { computed } from 'vue'

import type { ICampaign } from '../../../stores/mockMarketing'
import MarketingCampaignDetail from './MarketingCampaignDetail.vue'
import MarketingCampaignList from './MarketingCampaignList.vue'

const props = defineProps<{
  modelValue: ICampaign | null
}>()

const emit = defineEmits<{
  'update:modelValue': [value: ICampaign | null]
}>()

const handleSelect = (camp: ICampaign) => {
  emit('update:modelValue', camp)
}

const selectedCampaign = computed(() => props.modelValue)
</script>

<template>
  <div class="marketing-campaigns">
    <MarketingCampaignDetail
      v-if="selectedCampaign"
      :campaign="selectedCampaign"
    />
    <MarketingCampaignList
      v-else
      @select="handleSelect"
    />
  </div>
</template>

<style scoped>
.marketing-campaigns {
  display: flex;
  flex-direction: column;
}
</style>
