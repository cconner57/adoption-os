<script setup lang="ts">
import Candid from '../../common/candid-award/Candid.vue'
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
        <span>
          <h5>500+</h5>
          <p>pets rescued in 2025</p>
        </span>
        <span>
          <h5>500+</h5>
          <p>pets rescued in 2026</p>
        </span>
      </div>
    </content>
  </section>
</template>

<style scoped lang="css">
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
