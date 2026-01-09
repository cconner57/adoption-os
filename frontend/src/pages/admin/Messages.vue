<script setup lang="ts">
import { ref, computed, nextTick, watch } from 'vue'
import {
  mockThreads,
  mockMessages,
  mockUsers,
  currentUser,
  type IThread,
  type IMessage,
} from '../../stores/mockMessages'
import { InputField, Button, InputSelectGroup } from '../../components/common/ui'

// State
const activeThreadId = ref<string>('th-1')
const newMessageText = ref('')
const showCreateModal = ref(false)
const showMembersModal = ref(false)
const chatContainer = ref<HTMLElement | null>(null)

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
  scrollToBottom()
}

const sendMessage = () => {
  if (!newMessageText.value.trim()) return

  mockMessages.value.push({
    id: `m-${Date.now()}`,
    threadId: activeThreadId.value,
    userId: currentUser.value.id,
    text: newMessageText.value,
    timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
    type: 'text',
  })

  newMessageText.value = ''
  scrollToBottom()
}

const scrollToBottom = async () => {
  await nextTick()
  if (chatContainer.value) {
    chatContainer.value.scrollTop = chatContainer.value.scrollHeight
  }
}

// Admin Actions
const newThreadName = ref('')
const newThreadDesc = ref('')

const createThread = () => {
  if (!newThreadName.value) return

  const id = `th-${Date.now()}`
  mockThreads.value.push({
    id,
    type: 'channel',
    name: newThreadName.value.toLowerCase().replace(/\s+/g, '-'),
    description: newThreadDesc.value || 'No description',
    icon: '#️⃣',
  })

  showCreateModal.value = false
  newThreadName.value = ''
  newThreadDesc.value = ''
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
  if (confirm('Remove this user from the thread?')) {
    alert('User removed (mock action)')
  }
}

const addMember = () => {
  alert('In a real app, this would open a user picker!')
}

watch(activeThreadId, scrollToBottom)
</script>

<template>
  <div class="messages-page">
    <!-- SIDEBAR -->
    <div class="sidebar">
      <div class="sidebar-header">
        <h2>AdoptionOS</h2>
        <span class="user-badge">{{ currentUser.name }}</span>
      </div>

      <div class="channels-section">
        <!-- VOLUNTEER CHANNELS -->
        <div class="section-title">
          <span>VOLUNTEER CHANNELS</span>
          <button
            v-if="currentUser.role === 'admin'"
            class="add-btn"
            @click="showCreateModal = true"
            title="Create Channel"
          >
            +
          </button>
        </div>

        <div class="channel-list">
          <button
            v-for="thread in sortedChannels"
            :key="thread.id"
            class="channel-btn"
            :class="{ active: activeThreadId === thread.id }"
            @click="selectThread(thread.id)"
          >
            <span class="channel-icon">{{ thread.icon }}</span>
            <span class="channel-name">{{ thread.name }}</span>
            <span v-if="thread.isPrivate" class="lock-icon">🔒</span>
          </button>
        </div>

        <!-- ADOPTER INQUIRIES -->
        <div class="section-title mt-6">
          <span>ADOPTER INQUIRIES</span>
        </div>
        <div class="channel-list">
          <button
            v-for="thread in sortedInquiries"
            :key="thread.id"
            class="channel-btn"
            :class="{ active: activeThreadId === thread.id }"
            @click="selectThread(thread.id)"
          >
            <span class="channel-icon">{{ thread.icon }}</span>
            <div class="inquiry-meta">
              <span class="channel-name">{{ thread.name }}</span>
              <span class="method-badge" :class="thread.contactMethod">
                {{ thread.contactMethod?.toUpperCase() }}
              </span>
            </div>
          </button>
        </div>
      </div>
    </div>

    <!-- MAIN CHAT AREA -->
    <div class="main-chat">
      <!-- HEADER -->
      <div class="chat-header" v-if="activeThread">
        <div class="header-info">
          <h3>
            <span class="header-icon">{{ activeThread.icon }}</span>
            {{ activeThread.name }}
          </h3>
          <p class="topic">{{ activeThread.description }}</p>
        </div>

        <div class="header-actions" v-if="currentUser.role === 'admin'">
          <Button
            title="Manage Members"
            size="small"
            color="white"
            :onClick="() => (showMembersModal = true)"
          />
          <button class="delete-thread-btn" @click="deleteActiveThread" title="Delete Channel">
            🗑️
          </button>
        </div>
      </div>

      <!-- MESSAGES -->
      <div class="messages-container" ref="chatContainer">
        <div v-if="activeThread" class="message-stream">
          <div class="intro-message">
            <div class="intro-icon">{{ activeThread.icon }}</div>
            <h2>Welcome to #{{ activeThread.name }}!</h2>
            <p>
              This is the start of the
              <span class="highlight">#{{ activeThread.name }}</span> channel.
            </p>
          </div>

          <div
            v-for="msg in threadMessages"
            :key="msg.id"
            class="message-row"
            :class="{ 'own-message': msg.userId === currentUser.id }"
          >
            <div class="avatar">
              {{ getUser(msg.userId).avatar }}
            </div>
            <div class="message-content">
              <div class="message-meta">
                <span class="username">{{ getUser(msg.userId).name }}</span>
                <span class="timestamp">{{ msg.timestamp }}</span>
              </div>
              <div class="message-text">
                {{ msg.text }}
              </div>
            </div>
          </div>
        </div>
        <div v-else class="no-thread-selected">
          <p>Please select a channel</p>
        </div>
      </div>

      <!-- INPUT AREA -->
      <div class="input-area">
        <div class="input-wrapper">
          <input
            v-model="newMessageText"
            :placeholder="`Message #${activeThread?.name || 'channel'}`"
            @keyup.enter="sendMessage"
          />
          <button class="send-btn" @click="sendMessage">➤</button>
        </div>
      </div>
    </div>

    <!-- CREATE CHANNEL MODAL -->
    <div v-if="showCreateModal" class="modal-overlay" @click.self="showCreateModal = false">
      <div class="modal-card">
        <h3>Create New Channel</h3>
        <div class="form-group">
          <label>Channel Name</label>
          <InputField v-model="newThreadName" placeholder="e.g. shelter-events" />
        </div>
        <div class="form-group">
          <label>Description</label>
          <InputField v-model="newThreadDesc" placeholder="e.g. Planning upcoming events" />
        </div>
        <div class="modal-actions">
          <Button title="Cancel" color="white" :onClick="() => (showCreateModal = false)" />
          <Button title="Create Channel" color="black" :onClick="createThread" />
        </div>
      </div>
    </div>

    <!-- MEMBERS MODAL -->
    <div v-if="showMembersModal" class="modal-overlay" @click.self="showMembersModal = false">
      <div class="modal-card">
        <h3>Manage Members - #{{ activeThread?.name }}</h3>
        <p class="subtitle">Only admins can add or remove members.</p>

        <div class="members-list">
          <div v-for="user in mockUsers" :key="user.id" class="member-row">
            <div class="m-info">
              <span class="m-avatar">{{ user.avatar }}</span>
              <span class="m-name">{{ user.name }}</span>
              <span v-if="user.role === 'admin'" class="admin-badge">Admin</span>
            </div>
            <button
              v-if="user.id !== currentUser.id"
              class="remove-btn"
              @click="removeMember(user.id)"
            >
              Remove
            </button>
          </div>
        </div>

        <div class="modal-actions">
          <Button title="Add People" color="purple" :onClick="addMember" />
          <Button title="Done" color="black" :onClick="() => (showMembersModal = false)" />
        </div>
      </div>
    </div>
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

/* SIDEBAR */
.sidebar {
  width: 260px;
  background: #f8fafc; /* Slate-50 */
  border-right: 1px solid #e2e8f0;
  display: flex;
  flex-direction: column;
}

.sidebar-header {
  padding: 16px;
  border-bottom: 1px solid #e2e8f0;

  h2 {
    margin: 0;
    font-size: 1.1rem;
    color: var(--text-primary);
  }
}

.user-badge {
  font-size: 0.8rem;
  color: hsl(from var(--color-neutral) h s 50%);
  display: block;
  margin-top: 4px;
}

.channels-section {
  padding: 16px 8px;
  flex: 1;
  overflow-y: auto;
}

.section-title {
  padding: 0 8px 8px 8px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.75rem;
  font-weight: 700;
  color: hsl(from var(--color-neutral) h s 50%);

  .add-btn {
    background: none;
    border: none;
    color: hsl(from var(--color-neutral) h s 50%);
    cursor: pointer;
    font-size: 1.2rem;
    line-height: 1;
    &:hover {
      color: var(--text-primary);
    }
  }
}

.channel-list {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.mt-6 {
  margin-top: 24px;
}

.channel-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  background: none;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  color: #64748b;
  width: 100%;
  text-align: left;
  font-size: 0.95rem;

  &:hover {
    background: hsl(from var(--color-neutral) h s 90%);
    color: var(--text-primary);
  }

  &.active {
    background: hsl(from var(--color-secondary) h s 95%);
    color: var(--color-secondary);
    font-weight: 500;
  }
}

.inquiry-meta {
  display: flex;
  flex-direction: column;
  line-height: 1.2;
}

.method-badge {
  font-size: 0.65rem;
  font-weight: 700;

  &.sms {
    color: var(--color-primary);
  }
  &.email {
    color: var(--color-warning);
  }
}

.lock-icon {
  margin-left: auto;
  font-size: 0.7rem;
}

/* MAIN CHAT */
.main-chat {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: white;
}

.chat-header {
  padding: 12px 20px;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  justify-content: space-between;
  align-items: center;

  h3 {
    margin: 0;
    font-size: 1.1rem;
    display: flex;
    align-items: center;
    gap: 8px;
  }
  .topic {
    margin: 4px 0 0 0;
    font-size: 0.85rem;
    color: hsl(from var(--color-neutral) h s 50%);
  }
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 12px;

  .delete-thread-btn {
    background: none;
    border: none;
    cursor: pointer;
    font-size: 1.1rem;
    opacity: 0.6;
    &:hover {
      opacity: 1;
    }
  }
}

.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  display: flex;
  flex-direction: column;
}

.message-stream {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.input-area {
  padding: 20px;
}

.input-wrapper {
  display: flex;
  gap: 8px;
  background: white;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  padding: 8px 12px;

  input {
    flex: 1;
    border: none;
    outline: none;
    font-size: 1rem;
    background: transparent;
  }

  .send-btn {
    background: var(--color-secondary);
    color: var(--text-inverse);
    border: none;
    width: 32px;
    height: 32px;
    border-radius: 4px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;

    &:hover {
      opacity: 0.9;
    }
  }
}

/* MESSAGES */
.intro-message {
  padding: 40px 0;
  border-bottom: 1px solid #f1f5f9;
  margin-bottom: 20px;

  .intro-icon {
    font-size: 3rem;
    background: #f1f5f9;
    width: 80px;
    height: 80px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 20px;
    margin-bottom: 16px;
  }

  h2 {
    margin: 0 0 8px 0;
  }
  p {
    margin: 0;
    color: hsl(from var(--color-neutral) h s 50%);
  }
  .highlight {
    color: var(--color-secondary);
    font-weight: 600;
  }
}

.message-row {
  display: flex;
  gap: 16px;

  &:hover {
    background: hsl(from var(--color-neutral) h s 98%);
    margin: 0 -20px;
    padding: 4px 20px;
  }
}

.avatar {
  width: 40px;
  height: 40px;
  background: hsl(from var(--color-neutral) h s 95%);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.2rem;
}

.message-content {
  flex: 1;
}

.message-meta {
  display: flex;
  gap: 8px;
  align-items: baseline;
  margin-bottom: 4px;

  .username {
    font-weight: 700;
    color: var(--text-primary);
  }
  .timestamp {
    font-size: 0.75rem;
    color: hsl(from var(--color-neutral) h s 60%);
  }
}

.message-text {
  color: var(--text-primary);
  line-height: 1.5;
}

/* MODAL */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
}

.modal-card {
  background: white;
  padding: 24px;
  border-radius: 12px;
  width: 100%;
  max-width: 450px;
  box-shadow:
    0 20px 25px -5px rgba(0, 0, 0, 0.1),
    0 10px 10px -5px rgba(0, 0, 0, 0.04);

  h3 {
    margin-top: 0;
  }
  .subtitle {
    color: hsl(from var(--color-neutral) h s 50%);
    margin-bottom: 20px;
    font-size: 0.9rem;
  }
}

.form-group {
  margin-bottom: 16px;
  label {
    display: block;
    margin-bottom: 6px;
    font-weight: 500;
    font-size: 0.9rem;
  }
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 24px;
}

.members-list {
  max-height: 300px;
  overflow-y: auto;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
}

.member-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 12px;
  border-bottom: 1px solid #f1f5f9;

  &:last-child {
    border-bottom: none;
  }
}

.m-info {
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 500;
}

.m-avatar {
  font-size: 1.2rem;
}
.admin-badge {
  background: hsl(from var(--color-secondary) h s 95%);
  color: var(--color-secondary);
  font-size: 0.7rem;
  padding: 2px 6px;
  border-radius: 4px;
  text-transform: uppercase;
  font-weight: 700;
}

.remove-btn {
  color: var(--color-danger);
  background: none;
  border: none;
  font-size: 0.85rem;
  cursor: pointer;
  &:hover {
    text-decoration: underline;
  }
}
</style>
