<script setup lang="ts">
import { useRouter } from 'vue-router'
import { goToAdopt } from '../../../utils/navigate.ts'
import Button from '../../common/ui/Button.vue'

const router = useRouter()

function handleAdopt() {
  goToAdopt(router)
}
</script>

<template>
  <section class="hero-section">
    <h1>Rescue. Love. Adopt.</h1>
    <Button color="blue" title="Adopt Now" @click="handleAdopt" size="large" />
  </section>
</template>

<style scoped lang="css">
.hero-section {
  box-sizing: border-box;
  inline-size: 100%;
  min-height: 600px;
  background: url('/images/hero.png') center 70% / cover no-repeat;
  padding: 100px var(--layout-padding-side) 40px;
  overflow: hidden;

  & h1 {
    font-size: 6rem;
    color: var(--white);
    max-width: 600px;
    margin-bottom: 24px;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
    line-height: 100px;
    user-select: none;
  }

  @media (min-width: 321px) and (max-width: 430px) {
    padding-top: 8rem;
    padding-bottom: 2rem;
    min-height: 400px;
    background-position: 50% 20%;

    & h1 {
      font-size: 2.5rem;
      line-height: 3rem;
      width: 220px;
      max-width: 100%;
      margin-bottom: 16px;
    }
  }

  @media (min-width: 431px) and (max-width: 768px) {
    padding-top: 8rem;
    padding-bottom: 2rem;
    min-height: 400px;
    background-position: 50% 20%;

    & h1 {
      font-size: 2.5rem;
      line-height: 3rem;
      width: 200px;
      max-width: 100%;
      margin-bottom: 16px;
    }
  }

  @media (min-width: 1025px) and (max-width: 1440px) {
    padding-top: 8rem;
    padding-bottom: 2rem;
    & h1 {
      font-size: 5rem;
      line-height: 80px;
    }
  }
  @media (min-width: 1441px) {
    padding-top: 8rem;
    padding-bottom: 2rem;
    & h1 {
      font-size: 6rem;
      line-height: 100px;
    }
  }
}
</style>
