import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

export const useSettingsStore = defineStore('settings', () => {
  const settings = ref({
    organization: {
      name: 'Happy Tails Shelter',
      email: 'contact@happytails.org',
      phone: '555-0123',
      timezone: 'PST',
    },
    volunteers: {
      enableGamification: true,
      autoApproveShifts: false,
      minHoursForTier1: 20,
      shiftReminderHours: '24',
      allowTeenVolunteers: true,
    },
    pets: {
      defaultSpecies: 'cat',
      defaultBreed: 'Domestic Short Hair',
      defaultIntakeStatus: 'intake',
      defaultEnvironment: 'indoor',
      measurementSystem: 'imperial',
      requireMicrochip: false,
      defaultAdoptionFee: 0,
      defaultShelterLocation: '',
      defaultShowMedical: false,
      defaultBioTemplate: '',
      defaultGoodWithCats: 'unknown',
      defaultGoodWithDogs: 'unknown',
      defaultGoodWithKids: 'unknown',
      requirePhotoForPublic: false,
      defaultAutoHoldPeriod: 0,
    },
    notifications: {
      emailDigests: 'daily',
      incidentAlerts: true,
      newApplicationAlerts: true,
    },
    forms: {
      volunteer: {
        enabled: true,
        emails: ['director@happytails.org', 'volunteers@happytails.org'],
        newEmail: '',
      },
      surrender: {
        enabled: true,
        emails: ['director@happytails.org', 'intake@happytails.org'],
        newEmail: '',
      },
      adoption: {
        enabled: true,
        emails: ['adoptions@happytails.org'],
        newEmail: '',
      },
    },
    overview: {
      showRecentActivity: true,
      showPendingTasks: true,
      showStatsGraph: true,
    },
  })

  // Load from LocalStorage
  const saved = localStorage.getItem('appSettings')
  if (saved) {
    try {
      const parsed = JSON.parse(saved)
      // Deep merge or just Object.assign?
      // Simple merge for top-level keys to allow schema evolution if we add new keys
      Object.keys(parsed).forEach((key) => {
        const target = settings.value as any // eslint-disable-line @typescript-eslint/no-explicit-any
        if (target[key] && typeof target[key] === 'object') {
          Object.assign(target[key], parsed[key])
        } else {
          target[key] = parsed[key]
        }
      })
    } catch (e) {
      console.error('Failed to load settings', e)
    }
  }

  // Persist to LocalStorage
  watch(
    settings,
    (newVal) => {
      localStorage.setItem('appSettings', JSON.stringify(newVal))
    },
    { deep: true },
  )

  return {
    settings,
  }
})
