<script setup lang="ts">
import type { ITrip } from '../../../stores/mockTransport'
import { Capsules } from '../../common/ui'

defineProps<{
  trips: ITrip[]
  selectedTripId: string | null
}>()

defineEmits<{
  select: [trip: ITrip]
}>()

const tripDirectionLabel = (direction: ITrip['direction']) => {
  return direction === 'to_vet' ? 'Shelter ➝ Vet (Drop-off)' : 'Vet ➝ Shelter (Pick-up)'
}

const getStatusColor = (status: ITrip['status']) => {
  if (status === 'incident') return '#fee2e2' // Red (Emergency)
  if (status.includes('en_route')) return '#bfdbfe' // Blue (Moving)
  if (status.includes('at_')) return '#fde68a' // Yellow (Waiting/Location)
  if (status === 'completed') return '#d1fae5' // Green
  if (status === 'delayed') return '#ffedd5' // Orange
  return '#f3f4f6' // Gray
}

const getStatusLabel = (status: ITrip['status']) => {
  switch (status) {
    case 'en_route_vet':
      return 'En Route to Vet'
    case 'at_vet':
      return 'At Vet'
    case 'en_route_shelter':
      return 'En Route to Shelter'
    case 'at_shelter':
      return 'Back at Shelter'
    case 'pending':
      return 'Needs Driver'
    case 'incident':
      return 'Incident Reported'
    default:
      return status.replace('_', ' ')
  }
}
</script>

<template>
  <div class="shifts-list">
    <h3>Upcoming Shifts</h3>
    <div
      v-for="trip in trips"
      :key="trip.id"
      class="trip-card"
      :class="{
        selected: selectedTripId === trip.id,
        'active-trip': trip.status.includes('en_route'),
        'urgent-trip': trip.status === 'incident',
      }"
      @click="$emit('select', trip)"
    >
      <div class="card-top">
        <span class="direction-badge" :class="trip.direction">
          {{ tripDirectionLabel(trip.direction) }}
        </span>
        <Capsules
          :label="getStatusLabel(trip.status)"
          :color="getStatusColor(trip.status)"
          size="sm"
        />
      </div>

      <div class="trip-time">⏰ {{ trip.pickupTime }}</div>

      <div class="pet-avatars">
        <div v-for="pet in trip.pets.slice(0, 4)" :key="pet.id" class="pet-capsule">
          🐾 {{ pet.name }}
        </div>
        <div v-if="trip.pets.length > 4" class="more-pets">+{{ trip.pets.length - 4 }}</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.shifts-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
  overflow-y: auto;
  padding-right: 4px;

  h3 {
    margin: 0;
    font-size: 1.1rem;
    color: hsl(from var(--color-neutral) h s 50%);
  }
}

.trip-card {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 16px;
  cursor: pointer;
  transition: all 0.2s;
  border-left: 4px solid transparent;

  &:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
    border-color: #d1d5db;
  }

  &.selected {
    border-color: var(--color-secondary);
    background: hsl(from var(--color-secondary) h s 98%);
  }

  &.active-trip {
    border-left-color: var(--color-primary);
  }

  &.urgent-trip {
    border-left-color: var(--color-danger);
    background: hsl(from var(--color-danger) h s 98%);
  }
}

.card-top {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 12px;
}

.direction-badge {
  font-size: 0.8rem;
  font-weight: 700;
  text-transform: uppercase;

  &.to_vet {
    color: var(--color-secondary);
  }
  &.from_vet {
    color: var(--color-tertiary);
  }
}

.trip-time {
  font-size: 1.1rem;
  font-weight: 700;
  color: var(--text-primary);
  margin-bottom: 12px;
}

.pet-avatars {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.pet-capsule {
  background: hsl(from var(--color-neutral) h s 95%);
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 0.85rem;
  color: var(--text-primary);
}

.more-pets {
  font-size: 0.8rem;
  color: hsl(from var(--color-neutral) h s 50%);
  align-self: center;
}
</style>
