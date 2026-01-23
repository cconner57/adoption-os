import type { IHealthLogEntry } from '../models/common'
import { mockPetsData } from './mockPetData'

// Raw data from "Master Pet List - Daily Health Logs.csv"
const rawLogs = [
  {
    id: '1',
    name: 'Merry',
    date: '11/7/2025',
    recordedBy: 'Lori',
    weightG: '1880g',
    pounds: '4.14',
  },
  {
    id: '2',
    name: 'Legolas',
    date: '11/7/2025',
    recordedBy: 'Lori',
    weightG: '1330g',
    pounds: '2.93',
  },
  // ... (Truncated for file size limits)
]

// Transform raw logs into application compatible logs
export const mockHealthLogs: IHealthLogEntry[] = rawLogs
  .map((log) => {
    // Find pet ID by name (case-insensitive trim)
    const pet = mockPetsData.find(
      (p) => p.name.trim().toLowerCase() === log.name.trim().toLowerCase(),
    )

    // Parse weight: remove 'g' and convert to number
    const weight = parseInt(log.weightG.replace(/g/i, ''), 10) || 0

    return {
      id: `log_csv_${log.id}`,
      petId: pet ? pet.id : 'unknown_pet', // Fallback if name not found
      date: new Date(log.date).toISOString(),
      weight: weight,
      temperature: undefined, // CSV was missing these for most part
      eating: null,
      drinking: null,
      activity: null,
      urination: null,
      defecation: null,
      notes: '',
      recordedBy: log.recordedBy ? log.recordedBy.trim() : 'System',
    }
  })
  .filter((log) => log.petId !== 'unknown_pet') // Filter out any logs where we couldn't match the pet
