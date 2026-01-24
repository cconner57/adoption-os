<script setup lang="ts">
import SurrenderCat from '../components/about/surrender/SurrenderCat.vue'
import Candid from '../components/common/candid-award/Candid.vue'
import { useScrollReveal } from '../composables/useScrollReveal.ts'

const { vScrollReveal } = useScrollReveal()
</script>

<template>
  <main class="about">
    <section class="hero" v-scroll-reveal>
      <div class="content-wrapper">
        <h1>About IDOHR</h1>
        <p class="lead">
          I Dream of Home Rescue (IDOHR) is a volunteer-powered nonprofit helping homeless cats and
          dogs heal, thrive, and find loving homes across Southern California.
        </p>
      </div>
    </section>

    <section class="story">
      <div class="content-wrapper">
        <content v-scroll-reveal class="reveal-left">
          <img
            src="/images/watercolor.jpeg"
            alt="IDOHR watercolor inspiration"
            width="600"
            height="466"
          />
        </content>
        <div class="story-text reveal-right" v-scroll-reveal>
          <h2>Our Story</h2>
          <p>
            IDOHR began with a simple dream: to give every abandoned, neglected, and unwanted animal
            a safe place to call home. What started as a small circle of fosters has grown into a
            community of adopters, volunteers, and supporters who believe compassion changes
            lives—both animal and human.
          </p>
        </div>
      </div>
    </section>

    <section class="mission">
      <div class="content-wrapper">
        <content v-scroll-reveal>
          <h2 class="text-on-teal">Our Mission</h2>
          <p class="text-on-teal">
            We rescue, rehabilitate, and rehome animals in need—one tail, one paw, one heart at a
            time. Every animal receives medical care, a safe foster home, love-based socialization,
            and a path to a forever family.
          </p>
          <ul class="ticks text-on-teal">
            <li>Vaccinations, microchipping, and spay/neuter</li>
            <li>Medical treatment and behavioral support</li>
            <li>Adopter education and lifetime support</li>
          </ul>
        </content>
        <div class="image-wrapper reveal-delay-200" v-scroll-reveal>
          <img
            src="/images/mission.png"
            alt="Happy adopted dog and cat"
            class="mission-image"
            width="600"
            height="900"
          />
        </div>
      </div>
    </section>

    <section class="transparency">
      <div class="content-wrapper">
        <div class="awards" v-scroll-reveal>
          <Candid type="Gold" year="2024" class="award-item" />
          <Candid type="Gold" year="2023" class="award-item" style="transition-delay: 0.1s" />
          <Candid type="Silver" year="2022" class="award-item" style="transition-delay: 0.2s" />
        </div>
        <div class="notice reveal-left" v-scroll-reveal>
          <h2>Transparency</h2>
          <p>
            IDOHR is a 501(c)(3) nonprofit. Donations are tax-deductible as allowed by law. We
            publish annual updates and operate with financial transparency.
          </p>
          <p class="small">EIN: 81-0780050 • PO Box 7612, La Verne, CA 91750</p>
        </div>
      </div>
    </section>
  </main>
  <SurrenderCat />
</template>

<style scoped lang="css">

:deep(.reveal) {
  opacity: 0;
  transform: translateY(30px);
  transition: all 0.8s cubic-bezier(0.2, 0.8, 0.2, 1);
}

:deep(.reveal.active) {
  opacity: 1;
  transform: translateY(0);
}

:deep(.reveal-left) {
  transform: translateX(-30px);
}

:deep(.reveal-right) {
  transform: translateX(30px);
}

:deep(.reveal-left.active),
:deep(.reveal-right.active) {
  transform: translateX(0);
}

:deep(.reveal-delay-200.active) {
  transition-delay: 0.2s;
}

:deep(.award-item) {
  opacity: 0;
  transform: translateY(20px);
  transition: all 0.6s cubic-bezier(0.2, 0.8, 0.2, 1);
}

:deep(.reveal.active .award-item) {
  opacity: 1;
  transform: translateY(0);
}

.about {
  width: 100%;
  overflow: hidden;
  color: var(--text-primary);

  .content-wrapper {
    width: 100%;
    max-width: 1600px;
    margin: 0 auto;
    box-sizing: border-box;
  }

  .hero {
    padding: 150px var(--layout-padding-side) 50px;
    color: var(--text-inverse);
    background-color: var(--color-primary);
    display: flex;
    justify-content: center;

    .content-wrapper {
      text-align: center;
      display: flex;
      flex-direction: column;
      align-items: center;
    }

    h1 {
      font-size: 2.5rem;
      color: var(--text-inverse);
    }

    p {
      font-size: 1.25rem;
      max-width: 600px;
      margin-top: 12px;
      font-weight: 400; 
      color: var(--text-inverse);
    }
  }

  .story {
    background-color: var(--text-inverse);
    display: flex;
    justify-content: center;
    padding: 80px var(--layout-padding-side);

    .content-wrapper {
      display: flex;
      justify-content: space-between;
      align-items: center;
      gap: 4rem;
    }

    img {
      width: 100%;
      max-width: 450px;
      height: auto;
    }

    .story-text {
      max-width: 600px;

      p {
        font-size: 1.15rem;
        margin-top: 12px;
      }

      h2 {
        font-size: 2.5rem;
      }
    }
  }

  .mission {
    background-color: var(--color-primary);
    color: var(--text-inverse);
    display: flex;
    justify-content: center;
    padding: 80px var(--layout-padding-side);

    .content-wrapper {
      display: flex;
      justify-content: space-between;
      align-items: center;
      gap: 4rem;
    }

    .image-wrapper {
      display: flex;
      justify-content: flex-end;
      width: 100%;
      max-width: 450px;
    }

    h2 {
      font-size: 2.5rem;
      color: var(--text-inverse);
    }

    p {
      font-size: 1.15rem;
      max-width: 600px;
      margin-top: 12px;
      font-weight: 400; 
      color: var(--text-inverse);
    }

    img {
      width: 225px;
      max-width: 100%;
      height: auto;
    }

    & .ticks {
      list-style: none;
      padding: 0;
      margin: 12px 0 0;
      font-weight: 400; 
      color: var(--text-inverse);
    }

    & .ticks li {
      position: relative;
      padding-left: 24px;
      margin: 6px 0;
    }

    & .ticks li::before {
      content: '✓';
      position: absolute;
      left: 0;
      color: hsl(from var(--color-primary) h s 85%);
      font-weight: 700;
    }
  }

  img {
    border-radius: 8px;
    box-shadow: 0 4px 6px rgb(0 0 0 / 25%);
  }

  .transparency {
    color: var(--text-primary);
    background-color: var(--text-inverse);
    display: flex;
    justify-content: center;
    padding: 80px var(--layout-padding-side);

    .content-wrapper {
      display: flex;
      justify-content: space-between;
      align-items: center;
      gap: 4rem;
    }

    p {
      font-size: 1.15rem;
      max-width: 600px;
      margin-top: 12px;
    }

    .notice {
      text-align: right; 
      display: flex;
      flex-direction: column;
      align-items: flex-end; 

      & h2 {
        font-size: 2.5rem;
      }

      .small {
        font-size: 0.85rem;
        color: hsl(from var(--color-neutral) h s 40%);
        margin-top: 12px;
      }
    }

    .awards {
      display: flex;
      gap: 3rem;
      align-items: center;
      justify-content: center; 
    }
  }

  @media (width <= 1024px) {
    .story .content-wrapper,
    .mission .content-wrapper,
    .transparency .content-wrapper {
      flex-direction: column;
      gap: 40px;
      text-align: center;
    }

    .story img,
    .mission .image-wrapper {
      justify-content: center;
      margin: 0 auto;
    }

    .story-text,
    .mission content,
    .transparency .notice {
      max-width: 100%;
      align-items: center; 
    }

    .story-text p,
    .mission p,
    .transparency p,
    .mission .ticks {
      margin-left: auto;
      margin-right: auto;
      text-align: left; 
    }

    .mission .ticks {
      display: inline-block;
      text-align: left;
    }

    .transparency .awards {
      order: -1; 
      margin-bottom: 2rem;
    }

    .transparency .notice {
      text-align: center;
      align-items: center;
    }
  }

  @media (width >= 321px) and (width <= 430px) {
    .hero {
      padding: 100px var(--layout-padding-side) 50px;

      h1 {
        font-size: 2rem;
      }

      p {
        font-size: 1.15rem;
      }
    }

    .transparency .awards {
      flex-direction: column;
      gap: 3.5rem;
      padding-top: 1rem;
      padding-bottom: 1rem;
    }
  }
}
</style>
