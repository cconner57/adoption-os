<script setup lang="ts"></script>

<template>
  <section class="mission">
    <h4>Our Mission</h4>
    <div class="mission-content">
      <p>
        I Dream of Home Rescue's (IDOHR) mission is to rescue abandoned, unwanted, and surrendered
        animals. While we search for loving homes, we give them what all animals deserve: shelter,
        food, veterinary care, and love. IDOHR will also advocate and educate on the importance of
        pet adoption, spaying, neutering, and proper pet care.
      </p>
      <div class="image-container">
        <img
          src="/images/watercolor.jpeg"
          alt="watercolor painting of dog and cat"
          width="600"
          height="466"
        />
      </div>
    </div>
  </section>
</template>

<style scoped lang="css">
.mission {
  width: 100%;
  text-align: start;
  display: flex;
  flex-direction: column;
  gap: 16px;
  justify-content: center;

  & h4 {
    font-size: 2.75rem;
    color: var(--text-inverse);
    margin: 0;
  }

  .mission-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    gap: 40px;

    & .image-container {
      background-color: var(--text-inverse);
      padding: 16px;
      border-radius: 24px;
      flex-shrink: 0;
      display: flex;
      justify-content: center;
      align-items: center;
      box-shadow: 0 4px 6px rgb(0 0 0 / 5%);
    }

    & img {
      border-radius: 16px;
      height: auto;
      max-height: 300px;
      object-fit: cover;
      width: 100%;
      max-width: 400px;
    }

    & p {
      width: 100%;
      max-width: 650px;
      font-size: 1.2rem;
      color: var(--text-inverse);
      line-height: 1.6;
      font-weight: 400;
    }
  }

  @media (width <= 768px) {
    & h4 {
      font-size: 2rem;
    }

    .mission-content {
      flex-direction: column;
      align-items: center;
      gap: 24px;

      & img {
        width: 100%;
        max-width: 100%;
        margin: 0;
      }

      & p {
        width: 100%;
        font-size: 1rem;
        margin-top: 0;
      }
    }
  }

  @media (width >= 769px) and (width <= 1024px) {
    & h4 {
      font-size: 2.25rem;
    }

    .mission-content {
      gap: 32px;

      & p {
        flex: 1;
        width: auto;
        font-size: 1.1rem;
      }

      & img {
        width: 40%;
        max-width: 350px;
      }
    }
  }

  @media (width >= 1025px) and (width <= 1440px) {
    & h4 {
      font-size: 2.5rem;
    }

    .mission-content {
      & p {
        flex: 1;
        width: auto;
      }
    }
  }
}
</style>
