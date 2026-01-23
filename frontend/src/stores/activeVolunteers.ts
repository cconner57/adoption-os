import { defineStore } from 'pinia'
import { ref } from 'vue'

import type { IShift } from '../models/volunteers'

export const useActiveVolunteersStore = defineStore('activeVolunteers', () => {
  const activeCount = ref(0)
  const isFetching = ref(false)
  const weeklyShifts = ref<IShift[]>([])
  const error = ref<string | null>(null)

  const fetchActiveCount = async () => {
    isFetching.value = true
    error.value = null
    try {
      const apiUrl = import.meta.env.VITE_API_URL || ''
      // Request only 1 item since we just need the count from metadata
      const response = await fetch(`${apiUrl}/v1/volunteers?status=active&page_size=1`, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      })

      if (!response.ok) throw new Error(`vols: ${response.status} ${response.statusText}`)

      const json = await response.json()
      // Unwrap envelope: { status: "success", data: { volunteers: [], metadata: {} } }
      const data = json.data || {}

      // Use the totalRecords from metadata instead of counting the array
      activeCount.value = data.metadata?.totalRecords || 0
    } catch (e: unknown) {
      console.error('Error fetching active volunteers:', e)
      if (e instanceof Error) {
        error.value = e.message
      } else {
        error.value = String(e)
      }
    } finally {
      isFetching.value = false
    }
  }

  const fetchWeeklyShifts = async (start: string, end: string) => {
    try {
      const apiUrl = import.meta.env.VITE_API_URL || ''
      const response = await fetch(`${apiUrl}/v1/shifts?start=${start}&end=${end}`, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      })

      if (!response.ok) throw new Error(`shifts: ${response.status}`)
      const json = await response.json()
      // Unwrap: { status: "success", data: { shifts: [] } }
      const data = json.data || {}
      weeklyShifts.value = data.shifts || []
    } catch (e: unknown) {
      console.error('Error fetching weekly shifts:', e)
      const msg = e instanceof Error ? e.message : String(e)
      error.value = error.value ? `${error.value} | ${msg}` : msg
      weeklyShifts.value = []
    }
  }

  return {
    activeCount,
    isFetching,
    weeklyShifts,
    error,
    fetchActiveCount,
    fetchWeeklyShifts,
  }
})
