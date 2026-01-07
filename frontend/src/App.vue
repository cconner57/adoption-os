<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import NavBar from './components/common/nav-bar/NavBar.vue'
import '../src/styles/reset.css'
import '../src/styles/base.css'
import '../src/styles/shared.css'

const route = useRoute()
const showNavBar = computed(() => {
  const hiddenRoutes = ['/login', '/admin']
  return !hiddenRoutes.some((path) => route.path.startsWith(path))
})
</script>

<template>
  <NavBar v-if="showNavBar" />
  <main>
    <RouterView />
  </main>
</template>

<style scoped></style>
