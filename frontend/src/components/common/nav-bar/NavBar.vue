<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { goToDonate } from '../../../utils/navigate.ts'
import Button from '../ui/Button.vue'
import NavDrawer from './NavDrawer.vue'
import { useIsMobile, useIsTablet } from '../../../utils/useIsMobile.ts'

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
      <NavDrawer v-model="menuOpen" :size="28" style="color: #fff" />
    </nav>
    <nav v-else class="nav-blurred" aria-label="Primary navigation">
      <div class="nav-container">
        <section class="nav-links">
          <RouterLink to="/" class="nav-item" active-class="active"><p>Home</p></RouterLink>
          <RouterLink
            to="/about"
            class="nav-item"
            active-class="active"
            :class="{ active: $route.path.startsWith('/surrender') }"
            ><p>About</p></RouterLink
          >
          <RouterLink
            to="/adopt"
            class="nav-item"
            active-class="active"
            :class="{
              active: $route.path.startsWith('/adopt') || $route.path.startsWith('/pet-adoption'),
            }"
            ><p>Adopt</p></RouterLink
          >
          <RouterLink to="/volunteer" class="nav-item" active-class="active"
            ><p>Volunteer</p></RouterLink
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
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 1rem var(--layout-padding-side);
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1000;
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
        color: var(--font-color-light);
        text-decoration: none;

        p {
          font-size: 1.2rem;
          font-weight: 500;
        }
      }
      & .nav-item.active {
        font-weight: 600;
        border-bottom: 2px solid var(--font-color-light);
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
    background-color: hsl(182, 100%, 23%, 0.75);
    backdrop-filter: blur(10px);
    -webkit-backdrop-filter: blur(10px);
    transition:
      background-color 0.3s ease,
      backdrop-filter 0.3s ease;
  }

  @media (min-width: 321px) and (max-width: 430px) {
    nav {
      margin: 0;
      padding: calc(1rem + var(--safe-top)) 1rem 1rem 1rem;
      gap: 0.5rem;
      width: 100dvw;
      justify-content: space-between;

      & .nav-logo {
        & a {
          display: flex;
          align-items: center;
          gap: 0.5rem;
          & img {
            height: 40px;
            width: auto;
            border-radius: 8px;
          }

          & h1 {
            font-size: 1.2rem;
            color: var(--font-color-light);
          }
        }
      }
    }
  }

  @media (min-width: 431px) and (max-width: 768px) {
    nav {
      gap: 2rem;
      padding: calc(1rem + var(--safe-top)) 1.5rem 1rem 1.5rem;
      width: 100dvw;
      justify-content: space-between;
      .nav-logo {
        a {
          display: flex;
          align-items: center;
          gap: 0.75rem;
          img {
            height: 45px;
          }
          h1 {
            font-size: 1.5rem;
          }
        }
      }
    }
  }
}
</style>
