<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute,useRouter } from 'vue-router'

import { goToDonate } from '../../../utils/navigate.ts'
import { useIsMobile, useIsTablet } from '../../../utils/useIsMobile.ts'
import Button from '../ui/Button.vue'
import NavDrawer from './NavDrawer.vue'

const router = useRouter()
const route = useRoute()

const menuOpen = ref(false)

watch(
  () => route.fullPath,
  () => {
    menuOpen.value = false
  },
)

const isMobile = useIsMobile()
const isTablet = useIsTablet()
</script>

<template>
  <div class="nav-bar">
    <nav v-if="isMobile || isTablet" class="nav-blurred" aria-label="Primary navigation">
      <div class="nav-logo">
        <RouterLink to="/" class="nav-item">
          <img src="/images/idohr-logo.jpg" alt="" />
          <h1>I Dream of Home Rescue</h1>
        </RouterLink>
      </div>
      <NavDrawer v-model="menuOpen" :size="28" style="color: var(--text-inverse)" />
    </nav>
    <nav v-else class="nav-blurred" aria-label="Primary navigation">
      <div class="nav-container">
        <section class="nav-links">
          <RouterLink to="/" class="nav-item" active-class="active"><p data-text="Home">Home</p></RouterLink>
          <RouterLink
            to="/about"
            class="nav-item"
            active-class="active"
            :class="{ active: route.path.startsWith('/surrender') }"
            ><p data-text="About">About</p></RouterLink
          >
          <RouterLink
            to="/adopt"
            class="nav-item"
            active-class="active"
            :class="{
              active: route.path.startsWith('/adopt') || route.path.startsWith('/pet-adoption'),
            }"
            ><p data-text="Adopt">Adopt</p></RouterLink
          >
          <RouterLink to="/volunteer" class="nav-item" active-class="active"
            ><p data-text="Volunteer">Volunteer</p></RouterLink
          >
        </section>
        <Button title="Donate" color="green" @click="goToDonate(router)" />
      </div>
    </nav>
  </div>
</template>

<style scoped lang="css">
.nav-bar {
  nav {
    view-transition-name: page-navbar;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1rem var(--layout-padding-side);
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 2000;
    width: 100%;
    margin: 0 auto;
    overflow: hidden;
    align-self: center;
    transition: background-color 0.3s ease;

    .nav-container {
      width: 100%;
      max-width: 1600px;
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .nav-logo {
      .nav-item {
        text-decoration: none;
      }
    }

    & .nav-links {
      display: flex;
      gap: 4rem;

      & .nav-item {
        color: var(--text-inverse);
        text-decoration: none;

        p {
          font-size: 1.2rem;
          font-weight: 500;
          display: flex;
          flex-direction: column;
          align-items: center;
        }

        p::after {
          content: attr(data-text);
          font-weight: 600;
          height: 0;
          overflow: hidden;
          visibility: hidden;
        }
      }

      & .nav-item.active {
        font-weight: 600;
        border-bottom: 2px solid var(--text-inverse);
        padding-bottom: 4px;

        p {
          font-size: 1.2rem;
        }

        &:hover {
          cursor: pointer;
        }
      }
    }
  }

  .nav-blurred {
    background-color: oklch(52% 0.10 205deg / 85.1%); /* Final User-specified OKLCH (Medium-Dark Teal) */
    backdrop-filter: blur(10px);
    /* stylelint-disable-next-line property-no-vendor-prefix */
    -webkit-backdrop-filter: blur(10px);
    transition:
      background-color 0.3s ease,
      backdrop-filter 0.3s ease;
  }

  @media (width >= 321px) and (width <= 430px) {
    nav {
      margin: 0;
      padding: calc(1rem + var(--safe-top)) 1rem 1rem;
      gap: 0.5rem;
      width: 100dvw;
      justify-content: space-between;

      /* Flattened .nav-logo nesting */
      & .nav-logo a {
        display: flex;
        align-items: center;
        gap: 0.5rem;
      }

      & .nav-logo img {
        height: 40px;
        width: auto;
        border-radius: 8px;
      }

      & .nav-logo h1 {
        font-size: 1.2rem;
        color: var(--text-inverse);
      }
    }
  }

  @media (width >= 431px) and (width <= 768px) {
    nav {
      gap: 2rem;
      padding: calc(1rem + var(--safe-top)) 1.5rem 1rem;
      width: 100dvw;
      justify-content: space-between;

      /* Flattened .nav-logo nesting */
      .nav-logo a {
        display: flex;
        align-items: center;
        gap: 0.75rem;
      }

      .nav-logo img {
        height: 45px;
      }

      .nav-logo h1 {
        font-size: 1.5rem;
      }
    }
  }
}
</style>
