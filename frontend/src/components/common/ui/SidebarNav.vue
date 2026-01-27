<script setup lang="ts">
import Icon from './Icon.vue'

export interface NavItem {
  id?: string | number
  label: string
  icon?: string
  viewBox?: string
  to?: string
  exact?: boolean
}

const props = withDefaults(
  defineProps<{
    items: NavItem[]
    modelValue?: string | number
    variant?: 'default' | 'dashboard' | 'editor'
  }>(),
  {
    variant: 'default',
  },
)

const emit = defineEmits<{
  'update:modelValue': [id: string | number]
  click: [item: NavItem]
}>()

// eslint-disable-next-line no-unused-vars
function handleClick(item: NavItem, navigate?: (e: MouseEvent) => void) {
  if (item.id !== undefined) {
    emit('update:modelValue', item.id)
  }

  if (navigate) {
    // We need to pass the event if possible, but navigate() from router-link usually handles it.
    // However, since we are inside a custom click handler, we can just call it.
    // We pass a synthetic event or just call it as a function if strict typing allows.
    // Vue Router's navigate function expects a MouseEvent usually, but often works without if wrapped.
    // Let's create a fake event or just cast.
    // Actually, simply calling navigate() triggers the push.
    navigate({} as MouseEvent)
  }

  // Delay closing the menu slightly to ensure navigation registers
  // and to provide a better UX (visual feedback of click)
  setTimeout(() => {
    emit('click', item)
  }, 150)
}
</script>

<template>
  <nav class="sidebar-nav" :class="[`variant-${props.variant}`]">
    <template v-for="item in items" :key="item.label">

      <router-link
        v-if="item.to"
        :to="item.to"
        custom
        v-slot="{ href, navigate, isActive, isExactActive }"
      >
        <a
          :href="href"
          class="nav-item"
          :class="{ 'active': item.exact ? isExactActive : isActive }"
          @click="handleClick(item, navigate)"
        >
          <span class="icon" v-if="item.icon">
            <Icon :name="item.icon" :size="20" :viewBox="item.viewBox" />
          </span>
          <span class="label">{{ item.label }}</span>
        </a>
      </router-link>

      <button
        v-else
        class="nav-item"
        type="button"
        :class="{ active: modelValue === item.id }"
        @click="handleClick(item)"
      >
        <span class="icon" v-if="item.icon">
          <Icon :name="item.icon" :size="20" :viewBox="item.viewBox" />
        </span>
        <span class="label">{{ item.label }}</span>
      </button>
    </template>
  </nav>
</template>

<style scoped>
.sidebar-nav {
  padding: 16px 0;
  display: flex;
  flex-direction: column;
  gap: 2px;
  overflow-y: auto;
}

.sidebar-nav.variant-dashboard {
  flex: 1;
  gap: 8px;
}

.sidebar-nav.variant-editor {
  flex: none;
  gap: 8px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 24px;
  border-radius: 0;
  border-left: 3px solid transparent;
  cursor: pointer;
  text-decoration: none;
  background: none;
  border-top: none;
  border-right: none;
  border-bottom: none;
  width: 100%;
  text-align: left;
  color: var(--color-neutral-text-soft);
  font-weight: 500;
  font-size: 0.95rem;
  font-family: inherit;
  transition: all 0.2s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.sidebar-nav.variant-editor .nav-item {
  padding: 12px 16px;
}

.nav-item:hover {
  background-color: var(--color-neutral-weak);
  color: var(--text-primary);
}

.nav-item.active {
  background-color: var(--color-neutral-weak);
  color: var(--text-primary);
  border-left-color: var(--color-primary);
  font-weight: 700;
}

.nav-item.active .icon {
  color: var(--color-primary); /* Use primary brand color for icon */
  opacity: 1;
}

.nav-item .icon {
  font-size: 1.25rem;
  width: 24px;
  display: flex;
  justify-content: center;
}
</style>
