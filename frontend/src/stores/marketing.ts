import { defineStore } from 'pinia'
import { ref } from 'vue'

import { API_ENDPOINTS } from '../constants/api'
import type { ICampaign } from './mockMarketing'

export const useMarketingStore = defineStore('marketing', () => {
  const campaigns = ref<ICampaign[]>([])
  const isFetching = ref(false)

  const fetchCampaigns = async () => {
    isFetching.value = true
    try {
      const response = await fetch(API_ENDPOINTS.MARKETING_CAMPAIGNS)
      if (!response.ok) throw new Error('Failed to fetch campaigns')
      const data = await response.json()
      // Backend returns envelope {"campaigns": [...]}
      campaigns.value = data.campaigns || []
    } catch (error) {
      console.error('Error fetching campaigns:', error)
      // Fallback to empty or keep prev state?
      // For now, let's just log. Backend might be down or not migrated yet.
    } finally {
      isFetching.value = false
    }
  }

  const updateCampaign = async (campaign: ICampaign) => {
    try {
      const response = await fetch(`${API_ENDPOINTS.MARKETING_CAMPAIGNS}/${campaign.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(campaign),
      })
      if (!response.ok) throw new Error('Failed to update campaign')

      const data = await response.json()
      // Update local state
      const index = campaigns.value.findIndex(c => c.id === campaign.id)
      if (index !== -1) {
        campaigns.value[index] = data.campaign
      }
    } catch (error) {
      console.error('Error updating campaign:', error)
      throw error
    }
  }

  return {
    campaigns,
    isFetching,
    fetchCampaigns,
    updateCampaign
  }
})
