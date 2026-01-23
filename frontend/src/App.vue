<script setup lang="ts">
import '../src/styles/reset.css'
import '../src/styles/base.css'
import '../src/styles/shared.css'

import { computed } from 'vue'
import { useRoute } from 'vue-router'

import NavBar from './components/common/nav-bar/NavBar.vue'
import PageLoader from './components/common/PageLoader.vue'

const route = useRoute()
const showNavBar = computed(() => {
  return !route.meta.hideNavbar
})
</script>

<template>
  <PageLoader />
  <NavBar v-if="showNavBar" />
  <main>
    <RouterView />
  </main>
</template>

<style scoped></style>
