<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'

import { Button,Capsules, InputField, InputSelectGroup } from '../../components/common/ui'
import { donationStats, type ITransaction,mockTransactions } from '../../stores/mockDonations'

const searchQuery = ref('')
const filterType = ref<'all' | ITransaction['type']>('all')
const isMounted = ref(false)

onMounted(() => {
  isMounted.value = true
})

const filteredTransactions = computed(() => {
  return mockTransactions.value
    .filter((tx) => {

      const searchMatch =
        !searchQuery.value ||
        tx.donorName.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        tx.relatedPetName?.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        tx.notes?.toLowerCase().includes(searchQuery.value.toLowerCase())

      const typeMatch = filterType.value === 'all' || tx.type === filterType.value

      return searchMatch && typeMatch
    })
    .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
})

const formatCurrency = (amount: number) => {
  return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(amount)
}

const getTypeLabel = (type: ITransaction['type']) => {
  switch (type) {
    case 'adoption_fee':
      return '🏠 Adoption Fee'
    case 'donation':
      return '❤️ Donation'
    case 'sponsorship':
      return '🌟 Sponsorship'
    case 'merch':
      return '👕 Merchandise'
    default:
      return type
  }
}

const getStatusColor = (status: ITransaction['status']) => {
  switch (status) {
    case 'completed':
      return 'green' // Primary is teal/greenish
    case 'pending':
      return 'orange'
    case 'failed':
      return 'red'
    case 'refunded':
      return 'gray'
    default:
      return 'gray'
  }
}

const showLogModal = ref(false)
const newDonation = ref({
  donorName: '',
  amount: '',
  method: 'Cash',
  notes: '',
})

const handleLogDonation = () => {
  if (!newDonation.value.donorName || !newDonation.value.amount) return

  mockTransactions.value.unshift({
    id: `tx-${Date.now()}`,
    date: new Date().toISOString().split('T')[0],
    donorName: newDonation.value.donorName,
    amount: Number.parseFloat(newDonation.value.amount),
    type: 'donation',
    method: newDonation.value.method as any, // eslint-disable-line @typescript-eslint/no-explicit-any
    status: 'completed',
    notes: newDonation.value.notes,
  })

  showLogModal.value = false
  newDonation.value = { donorName: '', amount: '', method: 'Cash', notes: '' }
}
</script>

<template>
  <div class="donations-page">
    <Teleport v-if="isMounted" to="#mobile-header-target" :disabled="false">
      <h1 class="mobile-header-title">Donations & Revenue</h1>
    </Teleport>

    <div class="page-header">
      <h1 class="desktop-only">Donations & Revenue</h1>
      <Button title="Log Manual Donation" color="purple" :onClick="() => (showLogModal = true)" />
    </div>

    <div class="stats-grid">
      <div class="stat-card">
        <span class="stat-label">Total Revenue</span>
        <span class="stat-value">{{ formatCurrency(donationStats.totalRevenue) }}</span>
      </div>
      <div class="stat-card">
        <span class="stat-label">This Month</span>
        <span class="stat-value highlight">{{ formatCurrency(donationStats.monthRevenue) }}</span>
      </div>
      <div class="stat-card">
        <span class="stat-label">Transactions</span>
        <span class="stat-value">{{ donationStats.recentCount }}</span>
      </div>
    </div>

    <div class="filters-bar">
      <div class="search-wrap">
        <InputField v-model="searchQuery" placeholder="Search donors, pets, or notes..." />
      </div>
      <InputSelectGroup
        :modelValue="filterType"
        @update:modelValue="(val) => (filterType = val as any)"
        :options="[
          { label: 'All Types', value: 'all' },
          { label: 'Adoptions', value: 'adoption_fee' },
          { label: 'Donations', value: 'donation' },
          { label: 'Sponsorships', value: 'sponsorship' },
        ]"
      />
    </div>

    <div class="table-container">
      <table class="data-table">
        <thead>
          <tr>
            <th>Date</th>
            <th>Type</th>
            <th>Donor / Payer</th>
            <th>Details</th>
            <th>Method</th>
            <th>Amount</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="tx in filteredTransactions" :key="tx.id">
            <td class="date-cell">{{ tx.date }}</td>
            <td>
              <span class="type-capsule">{{ getTypeLabel(tx.type) }}</span>
            </td>
            <td class="fw-bold">{{ tx.donorName }}</td>
            <td class="details-cell">
              <span v-if="tx.relatedPetName" class="pet-badge">🐾 {{ tx.relatedPetName }}</span>
              <span v-if="tx.notes" class="notes-text">{{ tx.notes }}</span>
            </td>
            <td>{{ tx.method }}</td>
            <td class="amount-cell">{{ formatCurrency(tx.amount) }}</td>
            <td>
              <Capsules :label="tx.status" :color="getStatusColor(tx.status)" size="sm" />
            </td>
          </tr>
        </tbody>
      </table>

      <div v-if="filteredTransactions.length === 0" class="empty-state">No transactions found.</div>
    </div>

    <div v-if="showLogModal" class="modal-overlay" @click.self="showLogModal = false">
      <div class="modal-card">
        <h3>Log Manual Donation</h3>
        <p class="subtitle">Enter details for cash, check, or external donations.</p>

        <div class="form-group">
          <label>Donor Name</label>
          <InputField v-model="newDonation.donorName" aria-label="Donor Name" placeholder="e.g. John Doe" />
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Amount ($)</label>
            <InputField v-model="newDonation.amount" aria-label="Amount" type="number" placeholder="0.00" />
          </div>
          <div class="form-group">
            <label>Method</label>
            <select v-model="newDonation.method" class="native-select" aria-label="Method">
              <option>Cash</option>
              <option>Check</option>
            </select>
          </div>
        </div>

        <div class="form-group">
          <label>Notes</label>
          <InputField v-model="newDonation.notes" aria-label="Notes" placeholder="Optional notes..." />
        </div>

        <div class="modal-actions">
          <Button title="Cancel" color="white" :onClick="() => (showLogModal = false)" />
          <Button title="Log Donation" color="white" :onClick="handleLogDonation" />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.donations-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  h1.desktop-only {
    margin: 0;
    font-size: 1.8rem;
    color: var(--text-primary);
  }
}

.mobile-header-title {
  display: none;
}

@media (width <= 768px) {
  .page-header h1.desktop-only {
    display: none;
  }

  .mobile-header-title {
    display: block;
    font-size: 1.25rem;
    font-weight: 800;
    color: var(--text-primary);
    margin: 0;
  }
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

.stat-card {
  background: #fff;
  padding: 20px;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.stat-label {
  font-size: 0.85rem;
  color: hsl(from var(--color-neutral) h s 50%);
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 2rem;
  font-weight: 700;
  color: var(--text-primary);

  &.highlight {
    color: var(--color-primary);
  }
}

.filters-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-wrap {
  width: 320px;
}

.table-container {
  background: #fff;
  border-radius: 12px;
  border: 1px solid #e5e7eb;
  overflow: hidden;
}

.data-table {
  width: 100%;
  border-collapse: collapse;

  th {
    text-align: left;
    padding: 16px;
    background: hsl(from var(--color-neutral) h s 98%);
    color: hsl(from var(--color-neutral) h s 50%);
    font-weight: 600;
    font-size: 0.85rem;
    text-transform: uppercase;
    border-bottom: 1px solid var(--border-color);
  }

  td {
    padding: 16px;
    border-bottom: 1px solid var(--border-color);
    color: var(--text-primary);
    font-size: 0.95rem;
  }

  tr:last-child td {
    border-bottom: none;
  }
}

.date-cell {
  font-family: monospace;
  color: hsl(from var(--color-neutral) h s 50%);
}

.fw-bold {
  font-weight: 600;
}

.amount-cell {
  font-weight: 700;
  font-family: monospace;
}

.details-cell {
  display: flex;
  flex-direction: column;
  gap: 4px;
  font-size: 0.85rem;
}

.pet-badge {
  background: hsl(from var(--color-secondary) h s 95%);
  color: var(--color-secondary);
  padding: 2px 6px;
  border-radius: 4px;
  display: inline-block;
  width: fit-content;
  font-weight: 600;
  font-size: 0.75rem;
}

.notes-text {
  color: hsl(from var(--color-neutral) h s 50%);
  font-style: italic;
}

.type-capsule {
  font-size: 0.9rem;
}

.empty-state {
  padding: 40px;
  text-align: center;
  color: hsl(from var(--color-neutral) h s 50%);
}

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgb(0 0 0 / 50%);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
}

.modal-card {
  background: #fff;
  padding: 24px;
  border-radius: 12px;
  width: 100%;
  max-width: 450px;
  box-shadow: 0 20px 25px -5px rgb(0 0 0 / 10%);

  h3 {
    margin-top: 0;
    margin-bottom: 4px;
  }

  .subtitle {
    margin-top: 0;
    color: hsl(from var(--color-neutral) h s 50%);
    margin-bottom: 20px;
    font-size: 0.9rem;
  }
}

.form-group {
  margin-bottom: 16px;
  display: flex;
  flex-direction: column;
  gap: 6px;

  label {
    font-weight: 500;
    font-size: 0.9rem;
  }
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

.native-select {
  padding: 10px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  font-size: 0.95rem;
  background: #fff;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 24px;
}
</style>
