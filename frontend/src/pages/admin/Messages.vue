<script setup lang="ts">
import { computed,ref } from 'vue'

import CreateChannelModal from '../../components/admin/messages/CreateChannelModal.vue'
import MembersModal from '../../components/admin/messages/MembersModal.vue'
import MessageChatView from '../../components/admin/messages/MessageChatView.vue'
import MessageThreadList from '../../components/admin/messages/MessageThreadList.vue'
import {
  currentUser,
  mockMessages,
  mockThreads,
  mockUsers,
} from '../../stores/mockMessages'

// State
const activeThreadId = ref<string>('th-1')
const showCreateModal = ref(false)
const showMembersModal = ref(false)

// Computed
const activeThread = computed(() => mockThreads.value.find((t) => t.id === activeThreadId.value))

const threadMessages = computed(
  () => mockMessages.value.filter((m) => m.threadId === activeThreadId.value),
  // Simple sort (mock data should be ordered, but safeguards)
  // .sort(...)
)

const getUser = (userId: string) =>
  mockUsers.value.find((u) => u.id === userId) || { name: 'Unknown', avatar: '?' }

const sortedChannels = computed(() => {
  return mockThreads.value.filter((t) => t.type === 'channel')
})

const sortedInquiries = computed(() => {
  return mockThreads.value.filter((t) => t.type === 'inquiry')
})

// Actions
const selectThread = (id: string) => {
  activeThreadId.value = id
}

const handleSendMessage = (text: string) => {
  mockMessages.value.push({
    id: `m-${Date.now()}`,
    threadId: activeThreadId.value,
    userId: currentUser.value.id,
    text,
    timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
    type: 'text',
  })
}

const handleCreateThread = (name: string, desc: string) => {
  const id = `th-${Date.now()}`
  mockThreads.value.push({
    id,
    type: 'channel',
    name: name.toLowerCase().replace(/\s+/g, '-'),
    description: desc || 'No description',
    icon: '#️⃣',
  })

  showCreateModal.value = false
  selectThread(id)
}

const deleteActiveThread = () => {
  if (!activeThread.value) return
  if (confirm(`Are you sure you want to delete #${activeThread.value.name}?`)) {
    mockThreads.value = mockThreads.value.filter((t) => t.id !== activeThreadId.value)
    activeThreadId.value = mockThreads.value[0]?.id || ''
  }
}

// User Management Actions
const removeMember = (userId: string) => {
   
  console.log('removing user', userId)
  if (confirm('Remove this user from the thread?')) {
    alert('User removed (mock action)')
  }
}

const addMember = () => {
  alert('In a real app, this would open a user picker!')
}
</script>

<template>
  <div class="messages-page">
    <MessageThreadList
      :currentUser="currentUser"
      :channels="sortedChannels"
      :inquiries="sortedInquiries"
      :activeThreadId="activeThreadId"
      @select="selectThread"
      @create="showCreateModal = true"
    />

    <MessageChatView
      :activeThread="activeThread"
      :messages="threadMessages"
      :currentUser="currentUser"
      :getUser="getUser"
      @send="handleSendMessage"
      @delete="deleteActiveThread"
      @manageMembers="showMembersModal = true"
    />

    <CreateChannelModal
      :isOpen="showCreateModal"
      @close="showCreateModal = false"
      @create="handleCreateThread"
    />

    <MembersModal
      :isOpen="showMembersModal"
      :activeThread="activeThread"
      :users="mockUsers"
      :currentUser="currentUser"
      @close="showMembersModal = false"
      @addMember="addMember"
      @removeMember="removeMember"
    />
  </div>
</template>

<style scoped>
.messages-page {
  display: flex;
  height: calc(100vh - 140px); /* Adjust based on dashboard header/nav */
  background: white;
  border-radius: 12px;
  overflow: hidden;
  border: 1px solid #e5e7eb;
}
</style>
