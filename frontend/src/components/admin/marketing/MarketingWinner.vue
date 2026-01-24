<script setup lang="ts">
import { computed } from 'vue'

import { useMarketingStore } from '../../../stores/marketing'
import type { ICampaign } from '../../../stores/mockMarketing'
import { Button } from '../../common/ui'

const props = defineProps<{
  campaign: ICampaign
}>()

const store = useMarketingStore()

const winnerDetails = computed(() => {
  if (!props.campaign.winnerId || !props.campaign.participants) return null
  return props.campaign.participants.find(p => p.id === props.campaign.winnerId)
})

const hasEmail = computed(() => {
  const contact = winnerDetails.value?.contact || ''
  return contact.includes('@')
})

const hasPhone = computed(() => {
  const contact = winnerDetails.value?.contact || ''
  return contact.replace(/\D/g, '').length >= 7
})

const pickWinner = () => {
  if (!props.campaign.participants?.length) return

  const participants = props.campaign.participants
  const randomIndex = Math.floor(Math.random() * participants.length)
  const winner = participants[randomIndex]

  // Mutate local prop object (reactive via store ref passing)
  props.campaign.winnerId = winner.id
  store.updateCampaign(props.campaign)

  alert(`The winner is ${winner.name} (Ticket #${winner.ticketNumber})!`)
}

const notifyWinner = (method: 'email' | 'text' | 'call') => {
  const winner = props.campaign.participants?.find(p => p.id === props.campaign.winnerId)
  if (!winner) return

  if (method === 'call') {
     alert(`Please call ${winner.name} at ${winner.contact}`)
  } else {
     alert(`Simulating ${method} notification to ${winner.name}... Sent!`)
  }
}
</script>

<template>
  <div class="winner-section">
      <h2>Winner Circle</h2>
      <div class="winner-card" v-if="winnerDetails">
        <div class="trophy-icon">🏆</div>
        <h3>{{ winnerDetails.name }}</h3>
        <p>Ticket #{{ winnerDetails.ticketNumber }}</p>

        <div class="notify-actions">
          <Button v-if="hasEmail" size="small" title="Email Winner" color="blue" @click="notifyWinner('email')" />
          <Button v-if="hasPhone" size="small" title="Text Winner" color="green" @click="notifyWinner('text')" />
          <Button size="small" title="Log Call" color="white" @click="notifyWinner('call')" />
        </div>
      </div>

      <div class="empty-winner" v-else>
        <p>No winner selected yet.</p>
        <Button title="🎲 Pick Random Winner" color="purple" @click="pickWinner" size="large" />
      </div>
  </div>
</template>

<style scoped>
.winner-section {
    background: #fff;
    padding: 24px;
    border-radius: 12px;
    border: 1px solid #e5e7eb;
    height: fit-content;

    h2 {
        margin-top: 0;
        text-align: center;
        font-size: 1.2rem;
        margin-bottom: 24px;
    }
}

.winner-card {
    text-align: center;
    background: hsl(from var(--color-warning) h s 60%);
    padding: 32px 24px;
    border-radius: 8px;
    border: 1px solid #f3f4f6;

    .trophy-icon {
        font-size: 3rem;
        margin-bottom: 12px;
    }

    h3 {
        margin: 0;
        font-size: 1.4rem;
        color: #fff;
    }

    p {
        margin: 8px 0 24px;
        color: var(--text-secondary);
        font-weight: 600;
    }
}

.notify-actions {
    display: flex;
    justify-content: center;
    gap: 8px;
    flex-wrap: wrap;
}

.empty-winner {
    text-align: center;
    padding: 32px 0;
    color: var(--text-secondary);

    p { margin-bottom: 24px; }
}
</style>
