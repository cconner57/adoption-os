export const formatDate = (dateStr: string) => {
  // If YYYY-MM-DD (Mock/Truncated), force UTC to avoid shift
  if (dateStr.length <= 10) {
    return new Date(dateStr).toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
      timeZone: 'UTC',
    })
  }
  // Otherwise (ISO), use Local time
  return new Date(dateStr).toLocaleDateString(undefined, {
    month: 'short',
    day: 'numeric',
    year: 'numeric',
  })
}

export const getStatusColor = (status: string) => {
  switch (status) {
    case 'approved':
      return 'green'
    case 'denied':
      return 'red'
    case 'needs_info':
      return 'orange'
    case 'autodeleted':
      return 'red'
    default:
      return 'neutral'
  }
}

export const getStatusText = (status: string) => {
  switch (status) {
    case 'approved':
      return 'Approved'
    case 'denied':
      return 'Denied'
    case 'needs_info':
      return 'Needs Info'
    case 'autodeleted':
      return 'Auto-Deleted'
    default:
      return 'Pending'
  }
}

export const getDaysPending = (dateStr: string) => {
  if (!dateStr) return 0
  const created = new Date(dateStr)
  const now = new Date()
  const diffTime = Math.abs(now.getTime() - created.getTime())
  return Math.ceil(diffTime / (1000 * 60 * 60 * 24))
}

export const formatKey = (key: string) => {
  const withSpaces = key.replace(/([a-z])([A-Z])/g, '$1 $2').replace(/_/g, ' ')
  return withSpaces
    .split(' ')
    .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
    .join(' ')
}

export const getSignatureSrc = (data: string) => {
  if (!data) return ''
  if (data === 'base64data') {
    return 'data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iNTAiPjxwYXRoIGQ9Ik0xMCwyNSBRNDAsNSA3MCwyNSBUMTMwLDI1IiBzdHJva2U9ImJsYWNrIiBzdHJva2Utd2lkdGg9IjIiIGZpbGw9Im5vbmUiLz48L3N2Zz4='
  }
  if (data.length < 50) return ''
  if (data.startsWith('data:image')) return data
  return `data:image/png;base64,${data}`
}

const formatPetObject = (val: unknown): string => {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const v = val as any
  const name = v.name || v.Name
  const breed = v.speciesBreedSize || v.SpeciesBreedSize || v.species || v.Species
  const age = v.age || v.Age

  // 1. Fully populated
  if (name) {
    const details = []
    if (breed) details.push(breed)
    if (age) details.push(age)
    return details.length > 0 ? `${name} (${details.join(', ')})` : name
  }

  // 2. Missing Name but has details
  if (!name && (breed || age)) {
    const type = breed || 'Unknown Pet'
    return age ? `${type} (${age})` : type
  }

  // 3. Simple Children check {name, age} (if mixed in same logic)
  if (v.Name && v.Age) return `${v.Name} (${v.Age})`
  if (v.name && v.age) return `${v.name} (${v.age})`

  return ''
}

export const formatValue = (val: unknown): string => {
  if (Array.isArray(val)) {
    return val
      .map((item) => {
        if (typeof item === 'object' && item !== null) {
           const formatted = formatPetObject(item)
           if (formatted) return formatted

           // Check if effectively empty to avoid displaying "{"age":""...}"
           const values = Object.values(item)
           const hasContent = values.some(v => v !== '' && v !== null && v !== undefined)
           if (!hasContent) return ''

          return JSON.stringify(item)
        }
        return String(item)
      })
      .filter(s => s !== '')
      .join('; ')
  }
  return String(val)
}

export const getDisplayFields = (fullApplication: Record<string, unknown>) => {
  if (!fullApplication) return []
  const entries = Object.entries(fullApplication)

  const filtered = entries.filter(([key, value]) => {
    if (key === 'id') return false
    if ((key === 'parentSignatureData' || key === 'parentSignatureDate') && !value) return false
    if (key === 'signatureData' || key === 'parentSignatureData') return false
    if (value === null || value === '' || value === undefined) return false
    return true
  })

  filtered.sort(([keyA], [keyB]) => {
    const priority = [
      'firstName',
      'lastName',
      'age',
      'birthday',
      'address',
      'city',
      'zip',
      'createdAt',
      'nameFull',
      'signatureDate',
      'allergies',
      'phoneNumber',
      'availability',
      'positionPreferences',
      'interestReason',
      'volunteerExperience',
      'emergencyContactName',
      'emergencyContactPhone',
    ]
    const idxA = priority.indexOf(keyA)
    const idxB = priority.indexOf(keyB)

    if (idxA !== -1 && idxB !== -1) return idxA - idxB
    if (idxA !== -1) return -1
    if (idxB !== -1) return 1

    return 0
  })

  return filtered.map(([key, value]) => {
    let label = formatKey(key)

    if (label.includes('Spouse')) {
      label = label.replace('Spouse', 'Spouse/Partner')
    }

    let displayValue: unknown = value

    if (key === 'createdAt') {
      label = 'Submitted At'
    }

    const dateKeys = ['birthday', 'signatureDate', 'parentSignatureDate']
    if (dateKeys.includes(key) && typeof value === 'string') {
      if (value.startsWith('0001-01-01')) {
        displayValue = 'N/A'
      } else {
        displayValue = formatDate(value)
      }
    }

    if (key === 'createdAt' && typeof value === 'string') {
      if (value.startsWith('0001-01-01')) {
        displayValue = 'N/A'
      } else if (value.length <= 10) {
        // Mock/Truncated: Date only
        displayValue = formatDate(value)
      } else {
        // Real/ISO: Date + Time (Local)
        const date = new Date(value)
        displayValue = `${formatDate(value)} at ${date.toLocaleTimeString('en-US', {
          hour: 'numeric',
          minute: '2-digit',
        })}`
      }
    }

    if (value === true || value === 'true') displayValue = 'Yes'
    if (value === false || value === 'false') displayValue = 'No'

    return {
      key,
      label,
      value: displayValue,
    }
  })
}
