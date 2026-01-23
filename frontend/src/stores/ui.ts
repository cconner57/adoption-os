import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUIStore = defineStore('ui', () => {
  const isLoading = ref(false)

  function startLoading() {
    isLoading.value = true
  }

  function stopLoading() {
    isLoading.value = false
  }

  return {
    isLoading,
    startLoading,
    stopLoading,
  }
})
