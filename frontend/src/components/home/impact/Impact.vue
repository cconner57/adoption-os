<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Candid from '../../common/candid-award/Candid.vue'

const currentYear = new Date().getFullYear()
const previousYear = currentYear - 1

const countCurrent = ref(0)
const countPrevious = ref(0)
const isLoading = ref(true)

const fetchCount = async (year: number) => {
  try {
    const response = await fetch(`http://localhost:8080/pets/adopted-count?year=${year}`)
    const data = await response.json()
    return data.count || 0
  } catch (error) {
    console.error(`Error fetching count for ${year}:`, error)
    return 0
  }
}

const getLabel = (count: number) => (count === 1 ? 'pet' : 'pets')

onMounted(async () => {
  // Parallel fetch
  const [prev, curr] = await Promise.all([fetchCount(previousYear), fetchCount(currentYear)])

  countPrevious.value = prev
  countCurrent.value = curr
  isLoading.value = false
})
</script>

<template>
  <section class="impact">
    <h4>Our Impact</h4>
    <content>
      <div class="awards">
        <Candid type="Gold" year="2024" />
        <Candid type="Gold" year="2023" />
        <Candid type="Silver" year="2022" />
      </div>
      <div class="divider"></div>
      <div class="stats">
        <div v-if="isLoading" class="loader-container">
          <div class="spinner"></div>
        </div>
        <template v-else>
          <span>
            <h5>{{ countPrevious }}</h5>
            <p>{{ getLabel(countPrevious) }} rescued in {{ previousYear }}</p>
          </span>
          <span>
            <h5>{{ countCurrent }}</h5>
            <p>{{ getLabel(countCurrent) }} rescued in {{ currentYear }}</p>
          </span>
        </template>
      </div>
    </content>
  </section>
</template>

<style scoped lang="css">
/* Loader Styles */
.loader-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
}

.spinner {
  width: 40px;
  height: 40px;
  border: 4px solid rgba(0, 0, 0, 0.1);
  border-left-color: var(--primary-color, #00c3c3);
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* Main Layout Styles */
.impact {
  width: 100%;
  background-color: var(--white);
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 24px 32px 50px 32px;
  border-radius: 12px;
  margin-top: -200px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
  container-type: inline-size;
  container-name: impact;

  & h4 {
    font-size: 2rem;
    color: var(--font-color-dark);
    margin-left: -15px;
  }
  & content {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    color: var(--font-color-dark);
    gap: 40px;
  }
  & .awards {
    display: flex;
    gap: 40px;
    justify-content: center;
  }
  .divider {
    border-left: 2px solid #e5e7eb;
    height: 120px;
    margin: 0 20px;
  }
  .stats {
    display: flex;
    gap: 40px;
    justify-content: center;
    width: 100%;
    max-width: 450px;
    min-height: 100px; /* Prevent layout jump during loading */

    & span {
      display: flex;
      flex-direction: column;
      gap: 8px;
      align-items: center;
      text-align: center;

      & h5 {
        font-size: 3rem;
        color: var(--font-color-dark);
      }
      & p {
        font-size: 1.1rem;
        color: var(--font-color-dark);
        max-width: 150px;
      }
    }
  }

  @container impact (max-width: 900px) {
    margin-top: -80px;
    padding: 30px;
    height: auto;

    & h4 {
      margin-left: 0;
      text-align: left;
      width: 100%;
      margin-bottom: 20px;
    }

    & content {
      flex-direction: column;
      gap: 40px;
      align-items: center;
    }

    .divider {
      width: 100%;
      height: 2px;
      border-left: none;
      border-top: 2px solid #e5e7eb;
      margin: 0;
    }

    & .awards {
      width: 100%;
      gap: 40px;
      margin-top: 0;
      justify-content: center;
      min-width: 0;
      flex-wrap: wrap;
    }

    .stats {
      width: 100%;
      gap: 30px;
      justify-content: center;
      min-width: 0;

      & span {
        align-items: center;
        text-align: center;
      }
    }
  }

  @container impact (max-width: 480px) {
    padding: 24px;

    .stats {
      flex-direction: column;
      gap: 30px;
      align-items: center;
    }
  }
}
</style>
