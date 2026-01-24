<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'

import Footer from '../components/common/footer/Footer.vue'
import BannerButton from '../components/common/ui/BannerButton.vue'
import HeroSection from '../components/home/hero-section/HeroSection.vue'
import Impact from '../components/home/impact/Impact.vue'
import Mission from '../components/home/mission/Mission.vue'
import AdoptionSpotlight from '../components/home/spotlight/AdoptionSpotlight.vue'
import { usePets } from '../composables/usePets.ts'
import { useScrollReveal } from '../composables/useScrollReveal.ts'

const router = useRouter()

const { spotlightPets, loading, error, fetchSpotlight } = usePets()
const { vScrollReveal } = useScrollReveal()

onMounted(() => {
  fetchSpotlight()
})
</script>

<template>
  <div class="container">
    <div v-scroll-reveal>
      <HeroSection />
    </div>
    <main class="section-1">
      <div class="content-wrapper">
        <div v-scroll-reveal class="reveal-delay-100">
          <AdoptionSpotlight :pets="spotlightPets" :loading="loading" :error="error" />
        </div>
        <div v-scroll-reveal class="reveal-delay-200">
          <Mission />
        </div>
      </div>
    </main>
    <main class="section-2">
      <div class="content-wrapper">
        <div v-scroll-reveal>
          <Impact />
        </div>
        <section class="call-to-action" v-scroll-reveal>
          <BannerButton
            imgSrc="/images/paw.svg"
            title="Adopt a Pet"
            subtitle="Find your perfect companion"
            color="blue"
            @click="() => router.push('/adopt')"
            class="hover-scale"
          />
          <BannerButton
            imgSrc="/images/hand.svg"
            title="Get Involved"
            subtitle="Volunteer, foster, or support us"
            color="purple"
            @click="() => router.push('/volunteer')"
            class="hover-scale"
          />
          <BannerButton
            imgSrc="/images/heart.svg"
            title="Donate"
            subtitle="Help us save more lives"
            color="green"
            @click="() => router.push('/donate')"
            class="hover-scale"
          />
        </section>
      </div>
    </main>

    <Footer />
  </div>
</template>

<style scoped lang="css">
.container {
  width: 100%;
}

:deep(.reveal) {
  opacity: 0;
  transform: translateY(30px);
  transition: all 0.8s cubic-bezier(0.2, 0.8, 0.2, 1);
}

:deep(.reveal.active) {
  opacity: 1;
  transform: translateY(0);
}

:deep(.reveal-delay-100.active) {
  transition-delay: 0.1s;
}

:deep(.reveal-delay-200.active) {
  transition-delay: 0.2s;
}

:deep(.hover-scale) {
  transition: transform 0.3s ease;
}

:deep(.hover-scale:hover) {
  transform: scale(1.05);
}

.content-wrapper {
  width: 100%;
  max-width: 1600px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 64px;
}

main {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.content-wrapper {
  display: flex;
  flex-direction: column;
  gap: 64px;

  @media (width >= 0) and (width <= 320px) {
    gap: 24px;
  }

  @media (width >= 321px) and (width <= 430px) {
    gap: 40px;
  }

  @media (width >= 431px) and (width <= 768px) {
    gap: 48px;
  }

  @media (width >= 769px) and (width <= 1024px) {
    gap: 56px;
  }
}

.section-1 {
  background-color: var(--color-primary); 
  padding: 0 var(--layout-padding-side) 220px;
  margin-top: -20px;
  min-height: auto;

  @media (width >= 0) and (width <= 320px) {
    margin-top: 0;
    padding-bottom: 40px;
  }

  @media (width >= 321px) and (width <= 430px) {
    margin-top: 0;
    padding-bottom: 60px;
  }

  @media (width >= 431px) and (width <= 768px) {
    margin-top: 0;
    padding-bottom: 80px;
  }

  @media (width >= 769px) and (width <= 1024px) {
    padding-bottom: 120px;
  }
}

.section-2 {
  background-color: var(--text-inverse);
  padding: 60px var(--layout-padding-side) 80px;
  height: auto;

  @media (width >= 0) and (width <= 320px) {
    padding-top: 40px;
    padding-bottom: 24px;
  }

  @media (width >= 321px) and (width <= 430px) {
    padding-top: 60px;
    padding-bottom: 32px;
  }

  @media (width >= 431px) and (width <= 768px) {
    padding-top: 80px;
    padding-bottom: 100px;
  }

  @media (width >= 769px) and (width <= 1024px) {
    padding-top: 100px;
    padding-bottom: 150px;
  }
}

.call-to-action {
  display: flex;
  margin: 20px 0 0;
  gap: 60px;
  width: 100%;
  justify-content: flex-start;

  & > * {
    flex: 1;
  }

  @media (width >= 0) and (width <= 320px) {
    flex-direction: column;
    gap: 16px;
    margin: 20px 0 0;
  }

  @media (width >= 321px) and (width <= 430px) {
    flex-direction: column;
    gap: 20px;
    margin: 20px 0 0;
  }

  @media (width >= 431px) and (width <= 768px) {
    flex-direction: column;
    gap: 32px;
  }

  @media (width >= 769px) and (width <= 1024px) {
    gap: 40px;
  }
}

.adopt-now-button {
  background-color: var(--color-primary);
  height: 48px;
  min-width: 160px;
  padding: 0 24px;
  font-size: 1rem;
  font-weight: 600;
  border-radius: 6px;
}
</style>
