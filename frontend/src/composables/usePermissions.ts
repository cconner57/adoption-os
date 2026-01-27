import { computed } from 'vue'

import { useAuthStore } from '../stores/auth'

export const ROLES = {
  SUPER_ADMIN: 50,
  ADMIN: 40,
  VOLUNTEER_2: 30,
  VOLUNTEER_1: 20,
  TEEN_VOLUNTEER: 10,
  NONE: 0,
}

// Map string roles (from DB/Auth) to numeric levels
const ROLE_MAP: Record<string, number> = {
  super_admin: ROLES.SUPER_ADMIN,
  admin: ROLES.ADMIN,
  'Tier 2': ROLES.VOLUNTEER_2, // Matches VolunteerEditor options
  'Tier 1': ROLES.VOLUNTEER_1, // Matches VolunteerEditor options
  Teen: ROLES.TEEN_VOLUNTEER, // Matches VolunteerEditor options
  // Fallbacks/Variations
  volunteer: ROLES.VOLUNTEER_1,
}

export function usePermissions() {
  const authStore = useAuthStore()

  const currentLevel = computed(() => {
    const roleStr = authStore.user?.Role || ''
    return ROLE_MAP[roleStr] || ROLES.NONE
  })

  function getLevel(roleParams: string | number): number {
    if (typeof roleParams === 'number') return roleParams
    return ROLE_MAP[roleParams] || ROLES.NONE
  }

  /**
   * Check if current user has at least the required level
   */
  function can(requiredLevel: number): boolean {
    return currentLevel.value >= requiredLevel
  }

  /**
   * Check if current user can edit a user with the target role
   */
  function canEditUser(targetRole: string | number): boolean {
    const myLevel = currentLevel.value
    const targetLevel = getLevel(targetRole)

    // Rule 1: SUPER_ADMIN can edit everyone
    if (myLevel >= ROLES.SUPER_ADMIN) return true

    // Rule 2: ADMIN (40) can edit anyone BELOW 40 (Volunteers)
    // They CANNOT edit themselves, other Admins, or Super Admins
    if (myLevel >= ROLES.ADMIN) {
      return targetLevel < ROLES.ADMIN
    }

    // Rule 3: VOLUNTEERS cannot edit anyone
    return false
  }

  return {
    ROLES,
    currentLevel,
    can,
    canEditUser,
  }
}
