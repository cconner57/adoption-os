<script setup lang="ts">
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'
import { usePets } from '../composables/usePets.ts'

import BannerButton from '../components/common/ui/BannerButton.vue'
import Footer from '../components/common/footer/Footer.vue'
import AdoptionSpotlight from '../components/home/spotlight/AdoptionSpotlight.vue'
import HeroSection from '../components/home/hero-section/HeroSection.vue'
import Mission from '../components/home/mission/Mission.vue'
import Impact from '../components/home/impact/Impact.vue'

const router = useRouter()

const { spotlightPets, loading, error, fetchSpotlight } = usePets()

onMounted(() => {
  fetchSpotlight()
})
</script>

<template>
  <div class="container">
    <HeroSection />
    <main class="section-1">
      <AdoptionSpotlight :pets="spotlightPets" :loading="loading" :error="error" />
      <Mission />
    </main>
    <main class="section-2">
      <Impact />
      <section class="call-to-action">
        <BannerButton
          imgSrc="/images/paw.svg"
          title="Adopt a Pet"
          subtitle="Find your perfect companion"
          color="blue"
          @click="() => router.push('/adopt')"
        />
        <BannerButton
          imgSrc="/images/hand.svg"
          title="Get Involved"
          subtitle="Volunteer, foster, or support us"
          color="purple"
          @click="() => router.push('/volunteer')"
        />
        <BannerButton
          imgSrc="/images/heart.svg"
          title="Donate"
          subtitle="Help us save more lives"
          color="green"
          @click="() => router.push('/donate')"
        />
      </section>
    </main>

    <Footer />
  </div>
</template>

<style scoped lang="css">
main {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 64px;

  @media (min-width: 321px) and (max-width: 430px) {
    padding: 0 1rem;
    gap: 40px;
  }
}

.section-1 {
  background-color: var(--background);
  padding-bottom: 220px;
  margin-top: -20px;
  height: 975px;

  @media (min-width: 321px) and (max-width: 430px) {
    margin: 0px;
    padding-bottom: 60px;
    height: auto;
  }
}

.section-2 {
  background-color: var(--white);
  padding-top: 150px;
  padding-bottom: 220px;
  height: 550px;

  @media (min-width: 321px) and (max-width: 430px) {
    padding: 40px 2rem 0;
    height: auto;
  }
}

.call-to-action {
  display: flex;
  margin: 20px 0 80px;
  max-width: var(--desktop-breakpoint);
  padding: 0 60px;
  gap: 60px;

  @media (min-width: 321px) and (max-width: 430px) {
    flex-direction: column;
    gap: 20px;
    margin: 10px 0 40px;
    padding: 0 2rem;
  }
}

.adopt-now-button {
  background-color: var(--green);
  height: 48px;
  min-width: 160px;
  padding: 0 24px;
  font-size: 1rem;
  font-weight: 600;
  border-radius: 6px;
}
</style>
