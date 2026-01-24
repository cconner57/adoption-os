<script setup lang="ts">
import { nextTick, ref, watch } from 'vue'

import type { IMessage,IThread } from '../../../stores/mockMessages'
import { Button } from '../../common/ui'

const props = defineProps<{
  activeThread: IThread | undefined
  messages: IMessage[]
  currentUser: any // eslint-disable-line @typescript-eslint/no-explicit-any
  getUser: (_id: string) => any // eslint-disable-line @typescript-eslint/no-explicit-any, no-unused-vars
}>()

const emit = defineEmits<{
  send: [text: string]
  delete: []
  manageMembers: []
}>()

const newMessageText = ref('')
const chatContainer = ref<HTMLElement | null>(null)

const scrollToBottom = async () => {
  await nextTick()
  if (chatContainer.value) {
    chatContainer.value.scrollTop = chatContainer.value.scrollHeight
  }
}

const sendMessage = () => {
  if (!newMessageText.value.trim()) return
  emit('send', newMessageText.value)
  newMessageText.value = ''
  scrollToBottom()
}

watch(
  () => props.activeThread,
  () => {
    scrollToBottom()
  },
)

watch(
  () => props.messages.length,
  () => {
    scrollToBottom()
  },
)
</script>

<template>
  <div class="main-chat">
    
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
          :onClick="() => $emit('manageMembers')"
        />
        <button class="delete-thread-btn" @click="$emit('delete')" title="Delete Channel">
          🗑️
        </button>
      </div>
    </div>

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
          v-for="msg in messages"
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
</template>

<style scoped>
.main-chat {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: #fff;
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
    margin: 4px 0 0;
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
  background: #fff;
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
    margin: 0 0 8px;
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

.no-thread-selected {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  color: hsl(from var(--color-neutral) h s 50%);
  font-style: italic;
}
</style>
