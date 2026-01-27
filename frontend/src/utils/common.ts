export const formatDate = (dateString: string, birthday?: boolean) => {
  if (birthday) {
    const date = new Date(dateString)
    const now = new Date()
    let years = now.getFullYear() - date.getFullYear()
    let months = now.getMonth() - date.getMonth()

    if (months < 0) {
      years--
      months += 12
    }

    if (years > 0) {
      return `${years} year${years > 1 ? 's' : ''}`
    } else if (months > 0) {
      return `${months} month${months > 1 ? 's' : ''}`
    } else if (months === 0) {
      const days = Math.floor((now.getTime() - date.getTime()) / (1000 * 60 * 60 * 24))
      return `${days} day${days > 1 ? 's' : ''}`
    }
    return 'N/A'
  }

  const date = new Date(dateString)
  if (date.toString() === 'Invalid Date') {
    return 'N/A'
  }
  return date.toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

export const escapeRegExp = (string: string) => {
  return string.replace(/[.*+?^${}()|[\]\\]/g, String.raw`\$&`)
}
