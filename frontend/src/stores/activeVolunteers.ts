import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useActiveVolunteersStore = defineStore('activeVolunteers', () => {
  const activeCount = ref(0)
  const isFetching = ref(false)

  const fetchActiveCount = async () => {
    isFetching.value = true
    try {
      // Fetch all volunteers. The API supports pagination but defaults to 20.
      // We likely want a count of ALL active volunteers.
      // If the API supports a 'status' filter, we should use it.
      // Looking at the handler, it supports ?status=...

      const response = await fetch(
        `${import.meta.env.VITE_API_URL}/v1/volunteers?status=active&page_size=1000`,
        {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        },
      )

      if (!response.ok) throw new Error('Failed to fetch volunteers')

      const data = await response.json()
      // Structure is { volunteers: [], metadata: {} }
      const volunteers = data.volunteers || []

      activeCount.value = volunteers.length
    } catch (error) {
      console.error('Error fetching active volunteers:', error)
      // Fallback or leave as 0
    } finally {
      isFetching.value = false
    }
  }

  return {
    activeCount,
    isFetching,
    fetchActiveCount,
  }
})
