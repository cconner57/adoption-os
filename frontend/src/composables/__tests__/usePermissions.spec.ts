import { beforeEach,describe, expect, it, vi } from 'vitest'
import { reactive } from 'vue'

import { usePermissions } from '../usePermissions'

// Use reactive to verify store reactivity like Pinia
const mockStore = reactive({
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    user: null as any
})

vi.mock('../../stores/auth', () => ({
    useAuthStore: () => mockStore
}))

describe('usePermissions', () => {
    beforeEach(() => {
        mockStore.user = null
    })

    it('calculates current level from user role', () => {
        const { currentLevel, ROLES } = usePermissions()

        mockStore.user = { Role: 'super_admin' }
        expect(currentLevel.value).toBe(ROLES.SUPER_ADMIN)

        mockStore.user = { Role: 'admin' }
        expect(currentLevel.value).toBe(ROLES.ADMIN)

        mockStore.user = { Role: 'volunteer' }
        expect(currentLevel.value).toBe(ROLES.VOLUNTEER_1)

        mockStore.user = null
        expect(currentLevel.value).toBe(ROLES.NONE)
    })

    it('checks capability with can()', () => {
        const { can, ROLES } = usePermissions()
        mockStore.user = { Role: 'admin' } // Level 40

        expect(can(ROLES.ADMIN)).toBe(true)
        expect(can(ROLES.VOLUNTEER_1)).toBe(true) // 40 >= 20
        expect(can(ROLES.SUPER_ADMIN)).toBe(false) // 40 < 50
    })

    it('checks edit capabilities', () => {
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const { canEditUser } = usePermissions() as any

        mockStore.user = { Role: 'super_admin' }
        expect(canEditUser('admin')).toBe(true)
        expect(canEditUser('super_admin')).toBe(true)

        // Admin
        mockStore.user = { Role: 'admin' }
        expect(canEditUser('volunteer')).toBe(true)
        expect(canEditUser('admin')).toBe(false) // Cannot edit peer
        expect(canEditUser('super_admin')).toBe(false)

        // Volunteer
        mockStore.user = { Role: 'volunteer' }
        expect(canEditUser('volunteer')).toBe(false)
    })
})
