import type { IPet } from '../models/common'
import { formatDigitDate } from './date'

export function getStatusColor(status: string) {
  const map: Record<string, string> = {
    available: 'green',
    'adoption-pending': 'orange',
    adopted: 'blue',
    foster: 'purple',
    hold: 'red',
    intake: 'gray',
    archived: 'gray',
  }
  return map[status] || 'gray'
}

export function getSpayNeuterLabel(sex?: string) {
  const s = sex?.toLowerCase()
  if (s === 'female') return 'Spayed'
  if (s === 'male') return 'Neutered'
  return 'Spayed/Neutered'
}

export function formatSpayNeuter(pet: IPet) {
  if (!pet.medical.spayedOrNeutered) {
    return 'No'
  }

  const dateStr = pet.medical.spayedOrNeuteredDate
  if (dateStr) {
    return formatDigitDate(dateStr)
  }

  return 'Yes'
}
