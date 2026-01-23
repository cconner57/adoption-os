<script setup lang="ts">
export interface NavItem {
  id?: string | number
  label: string
  icon?: string
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

function handleClick(item: NavItem) {
  if (item.id !== undefined) {
    emit('update:modelValue', item.id)
  }
  emit('click', item)
}
</script>

<template>
  <nav class="sidebar-nav" :class="[`variant-${props.variant}`]">
    <template v-for="item in items" :key="item.label">
      
      <router-link
        v-if="item.to"
        :to="item.to"
        class="nav-item"
        :active-class="item.exact ? '' : 'active'"
        :exact-active-class="item.exact ? 'active' : ''"
        @click="handleClick(item)"
      >
        <span class="icon" v-if="item.icon">{{ item.icon }}</span>
        <span class="label">{{ item.label }}</span>
      </router-link>

      <button
        v-else
        class="nav-item"
        type="button"
        :class="{ active: modelValue === item.id }"
        @click="handleClick(item)"
      >
        <span class="icon" v-if="item.icon">{{ item.icon }}</span>
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

  color: hsl(from var(--color-neutral) h s 50%);
  font-weight: 500;
  font-size: 0.95rem;
  font-family: inherit;
  transition: all 0.2s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.sidebar-nav.variant-editor .nav-item {
  padding: 12px 16px; 
}

.nav-item:hover {
  background-color: hsl(from var(--color-neutral) h s 95%);
  color: var(--text-primary);
}

.nav-item.active {
  background-color: hsla(from var(--color-secondary) h s l / 0.12);
  color: var(--color-secondary);
  border-left-color: var(--color-secondary);
  font-weight: 600;
}

.nav-item .icon {
  font-size: 1.25rem;
  width: 24px;
  display: flex;
  justify-content: center;
}
</style>
