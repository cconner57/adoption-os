<script setup lang="ts">
import type { IPet } from '../../../models/common'
import PetRow from './PetRow.vue'

defineProps<{
  pets: IPet[]
  visibleColumns: Record<string, boolean>  
  expandedPetId: string | null
  statusFilter: string
}>()

defineEmits<{
  toggleExpand: [pet: IPet]
  edit: [pet: IPet]
  archive: [pet: IPet]
  markAdopted: [pet: IPet]
}>()
</script>

<template>
  <div class="table-container">
    <table class="pets-table">
      <thead>
        <tr>
          <th class="expand-col"></th>
          <th v-if="visibleColumns.photo" class="col-photo">Photo</th>
          <th v-if="visibleColumns.name" class="col-name">Name</th>
          <th v-if="visibleColumns.breed" class="col-species">Species</th>
          <th v-if="visibleColumns.sex" class="col-sex">Sex</th>
          <th v-if="visibleColumns.sn" class="text-center col-sn">S/N</th>
          <th v-if="visibleColumns.microchip" class="col-microchip">Microchip</th>
          <th v-if="visibleColumns.age" class="col-age">Age</th>
          <th v-if="visibleColumns.dob" class="col-dob">DOB</th>
          <th v-if="visibleColumns.intake" class="col-intake">Intake</th>
          <th v-if="statusFilter === 'adopted'" class="col-dob">Adopted Date</th>
          <th v-if="statusFilter === 'foster'" class="col-dob">Foster Start</th>
          <th v-if="visibleColumns.status" class="col-status">Status</th>
          <th v-if="visibleColumns.actions" class="col-actions">Actions</th>
        </tr>
      </thead>
      <tbody>
        <template v-for="(pet, index) in pets" :key="pet.id">
          <PetRow
            :pet="pet"
            :index="index"
            :visible-columns="visibleColumns"
            :is-expanded="expandedPetId === pet.id"
            :status-filter="statusFilter"
            @toggle-expand="$emit('toggleExpand', pet)"
            @edit="$emit('edit', pet)"
            @archive="$emit('archive', pet)"
            @mark-adopted="$emit('markAdopted', pet)"
          />
        </template>
      </tbody>
    </table>
    <div v-if="pets.length === 0" class="empty-state">No pets found matching criteria.</div>
  </div>
</template>

<style scoped>
.table-container {
  background: var(--text-inverse);
  border-radius: 12px;
  box-shadow: 0 4px 6px rgb(0 0 0 / 5%);
  overflow: auto; 
  flex: 1;
  z-index: 1;

  &::-webkit-scrollbar {
    width: 8px;
    height: 8px;
  }

  &::-webkit-scrollbar-track {
    background: #f1f5f9;
  }

  &::-webkit-scrollbar-thumb {
    background: #cbd5e1;
    border-radius: 4px;
  }

  &::-webkit-scrollbar-thumb:hover {
    background: #94a3b8;
  }
}

.pets-table {
  width: 100%;
  border-collapse: collapse;
  text-align: left;
  table-layout: fixed; 

  th {
    background: hsl(from var(--color-neutral) h s 98%);
    padding: 12px 16px;
    font-weight: 600;
    color: hsl(from var(--color-neutral) h s 50%);
    font-size: 0.9rem;
    border-bottom: 1px solid var(--border-color);
    position: sticky;
    top: 0;
    z-index: 10;
    white-space: nowrap;
    overflow: hidden; 
    text-overflow: ellipsis;
  }
}

.expand-col {
  width: 48px;
}

.col-photo {
  width: 80px;
}

.col-name {
  width: 140px;
  min-width: 120px;
}

.col-species {
  width: 80px;
}

.col-sex {
  width: 80px;
}

.col-sn {
  width: 60px;
  text-align: center;
}

.col-microchip {
  width: 100px;
}

.col-age {
  width: 140px;
}

.col-dob {
  width: 120px;
}

.col-intake {
  width: 120px;
}

.col-status {
  width: 140px;
}

.col-actions {
  width: 100px;
  text-align: right;
}

.text-center {
  text-align: center;
}

.empty-state {
  padding: 40px;
  text-align: center;
  color: var(--text-secondary);
  font-style: italic;
}
</style>
