import type { IHealthLogEntry } from '../models/common'
import { mockPetsData } from './mockPetData'

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
  
]

export const mockHealthLogs: IHealthLogEntry[] = rawLogs
  .map((log) => {
    
    const pet = mockPetsData.find(
      (p) => p.name.trim().toLowerCase() === log.name.trim().toLowerCase(),
    )

    const weight = parseInt(log.weightG.replace(/g/i, ''), 10) || 0

    return {
      id: `log_csv_${log.id}`,
      petId: pet ? pet.id : 'unknown_pet', 
      date: new Date(log.date).toISOString(),
      weight: weight,
      temperature: undefined, 
      eating: null,
      drinking: null,
      activity: null,
      urination: null,
      defecation: null,
      notes: '',
      recordedBy: log.recordedBy ? log.recordedBy.trim() : 'System',
    }
  })
  .filter((log) => log.petId !== 'unknown_pet') 
