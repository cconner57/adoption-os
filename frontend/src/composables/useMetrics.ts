

export function useMetrics() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars, no-unused-vars
  const submitMetric = async (eventType: string, eventData?: unknown) => {
    // Metrics disabled for now to prevent auth confusion
    return
    /*
    try {
      fetch(API_ENDPOINTS.METRICS, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          eventType,
          eventData,
        }),
      }).catch((err) => console.error('Metric logging failed:', err))
    } catch (e) {
      console.error(e)
    }
    */
  }

  return {
    submitMetric,
  }
}
