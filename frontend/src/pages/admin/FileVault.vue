<script setup lang="ts">
import { computed, ref } from 'vue'

import { Button, Icon, InputField } from '../../components/common/ui'

interface IFile {
  id: string
  name: string
  type: 'pdf' | 'folder' | 'img'
  size: string
  date: string
  owner?: string
}

const searchQuery = ref('')

const files = ref<IFile[]>([
  { id: '1', name: 'Adoption Contracts 2025', type: 'folder', size: '12 items', date: 'Jan 20, 2026' },
  { id: '2', name: 'Vet Records - Bella', type: 'pdf', size: '2.4 MB', date: 'Jan 18, 2026', owner: 'Bella (Dog)' },
  { id: '3', name: 'Surrender Form - Luna', type: 'pdf', size: '1.1 MB', date: 'Jan 15, 2026', owner: 'Luna (Cat)' },
  { id: '4', name: 'Volunteer Waiver Template', type: 'pdf', size: '540 KB', date: 'Dec 10, 2025' },
  { id: '5', name: 'Event Photos - Winter Gala', type: 'folder', size: '45 items', date: 'Dec 05, 2025' },
  { id: '6', name: 'Transport Manifest #442', type: 'pdf', size: '890 KB', date: 'Nov 22, 2025' },
  { id: '7', name: 'Kennel Card Template', type: 'pdf', size: '1.2 MB', date: 'Oct 01, 2025' },
  { id: '8', name: 'Expense Report Q4', type: 'pdf', size: '3.5 MB', date: 'Jan 02, 2026' },
])

const filteredFiles = computed(() => {
  if (!searchQuery.value) return files.value
  const q = searchQuery.value.toLowerCase()
  return files.value.filter((f) => f.name.toLowerCase().includes(q))
})

const getIcon = (type: string) => {
  if (type === 'folder') return 'folder'
  if (type === 'pdf') return 'fileText'
  return 'fileText'
}

const getColor = (type: string) => {
  if (type === 'folder') return 'var(--color-secondary)'
  if (type === 'pdf') return 'var(--color-danger)'
  return 'var(--text-secondary)'
}
</script>

<template>
  <div class="file-vault-page">
    <div class="page-header">
      <div class="header-left">
        <h1>File Vault</h1>
        <span class="count-badge">{{ filteredFiles.length }} items</span>
      </div>
      <div class="header-actions">
        <div class="search-wrapper">
          <InputField v-model="searchQuery" placeholder="Search files..." />
        </div>
        <Button title="Upload File" variant="primary" icon="plus" />
      </div>
    </div>

    <div class="files-grid">
      <div v-for="file in filteredFiles" :key="file.id" class="file-card">
        <div class="file-icon-wrapper" :style="{ color: getColor(file.type) }">
          <Icon :name="getIcon(file.type)" size="48" />
        </div>
        <div class="file-info">
          <h3 class="file-name" :title="file.name">{{ file.name }}</h3>
          <div class="file-meta">
            <span>{{ file.date }}</span>
            <span>•</span>
            <span>{{ file.size }}</span>
          </div>
          <div class="file-owner" v-if="file.owner">
            Related: {{ file.owner }}
          </div>
        </div>
        <div class="file-actions">
          <button class="action-btn" title="Download">
            <Icon name="arrowRight" size="16" style="transform: rotate(90deg)" />
          </button>
        </div>
      </div>
    </div>

    <div v-if="filteredFiles.length === 0" class="empty-state">
      No files found matching "{{ searchQuery }}"
    </div>
  </div>
</template>

<style scoped>
.file-vault-page {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 16px;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 12px;

  h1 {
    font-size: 1.8rem;
    color: var(--text-primary);
    margin: 0;
  }
}

.count-badge {
  background: var(--color-neutral-weak);
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.9rem;
  color: var(--color-neutral-text-soft);
  font-weight: 600;
}

.header-actions {
  display: flex;
  gap: 16px;
  align-items: center;
}

.search-wrapper {
  width: 300px;
}

.search-wrapper :deep(.field) {
  margin-bottom: 0;
}

.files-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 24px;
}

.file-card {
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  border-radius: 16px;
  padding: 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  gap: 16px;
  transition: transform 0.2s, box-shadow 0.2s;
  cursor: pointer;
  position: relative;

  &:hover {
    transform: translateY(-4px);
    box-shadow: 0 12px 24px -8px rgb(0 0 0 / 15%);
    border-color: var(--color-primary-border);
  }
}

.file-icon-wrapper {
  background: var(--color-neutral-surface);
  width: 80px;
  height: 80px;
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 8px;
}

.file-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
  width: 100%;
}

.file-name {
  font-size: 1rem;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  width: 100%;
}

.file-meta {
  font-size: 0.8rem;
  color: var(--text-secondary);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
}

.file-owner {
  font-size: 0.75rem;
  background: var(--color-neutral-weak);
  padding: 2px 8px;
  border-radius: 8px;
  color: var(--text-secondary);
  margin-top: 4px;
}

.file-actions {
  position: absolute;
  top: 12px;
  right: 12px;
  opacity: 0;
  transition: opacity 0.2s;
}

.file-card:hover .file-actions {
  opacity: 1;
}

.action-btn {
  background: var(--text-inverse);
  border: 1px solid var(--border-color);
  border-radius: 50%;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: var(--text-secondary);

  &:hover {
    background: var(--color-primary-weak);
    color: var(--color-primary);
    border-color: var(--color-primary);
  }
}

.empty-state {
  text-align: center;
  padding: 48px;
  color: var(--text-secondary);
  font-style: italic;
  font-size: 1.1rem;
}

@media (width <= 768px) {
  .header-actions {
    flex-direction: column;
    align-items: stretch;
    width: 100%;
  }

  .search-wrapper {
    width: 100%;
  }

  .files-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 16px;
  }

  .file-card {
    padding: 16px;
  }

  .file-icon-wrapper {
    width: 60px;
    height: 60px;
  }
}
</style>
