import type { IApplicationItem } from '../../components/admin/applications/ApplicationCard.vue'

export function formatRole(prefs: unknown): string | null {
  if (!prefs) return null
  if (Array.isArray(prefs)) return prefs.join(', ')
  if (typeof prefs === 'object') return '' // Avoid [object Object]
  return String(prefs)
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function mapApplicationToItem(app: { id: string; type: 'adoption' | 'volunteer' | 'surrender'; status: IApplicationItem['status']; created_at: string; data: any }): IApplicationItem {
  const d = app.data || {}
  let name = 'Unknown'

  if (app.type === 'volunteer') {
    name =
      d.firstName && d.lastName
        ? `${d.firstName} ${d.lastName}`
        : String(d.nameFull || 'Volunteer Applicant')
  } else if (app.type === 'adoption') {
    name = d.firstName && d.lastName ? `${d.firstName} ${d.lastName}` : 'Adoption Applicant'
  } else if (app.type === 'surrender') {
    name = d.firstName && d.lastName ? `${d.firstName} ${d.lastName}` : 'Surrender Applicant'
  }

  return {
    id: String(app.id),
    type: app.type,
    status: app.status,
    date: app.created_at,
    applicantName: name,
    email: String(d.email || d.Email || ''),
    details: {
      petName:
        String(d.petName || d.catPreferenceName ||
        d.animalName ||
        (d.catPreferenceBreed ? `Breed: ${d.catPreferenceBreed}` : null) || ''),
      role: formatRole(d.positionPreferences),
      reason: String(d.interestReason || d.adoptionReason || d.animalWhySurrendered || ''),
    },
    fullApplication: { ...d, createdAt: app.created_at },
  }
}

export function getStatusText(status: string): string {
    return (
      {
        pending: 'Pending',
        approved: 'Approved',
        denied: 'Denied',
        needs_info: 'Needs Info',
        autodeleted: 'Auto-Deleted',
        video_approved: 'Video Approved',
      }[status] || status
    )
  }

  export function getStatusColor(status: string): 'gray' | 'green' | 'red' | 'yellow' | 'blue' {
    const map: Record<string, 'gray' | 'green' | 'red' | 'yellow' | 'blue'> = {
      pending: 'blue',
      approved: 'green',
      denied: 'red',
      needs_info: 'yellow',
      autodeleted: 'gray',
      video_approved: 'green',
    }
    return map[status] || 'gray'
  }
