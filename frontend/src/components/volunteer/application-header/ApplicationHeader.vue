<script setup lang="ts">
const { headerTitle, headerText } = defineProps<{
  headerTitle: string
  headerText?: string
}>()
</script>

<template>
  <header class="form-header">
    <svg
      width="80"
      height="80"
      viewBox="0 0 128 128"
      xmlns="http://www.w3.org/2000/svg"
      role="img"
      aria-label="Dog paw icon"
      class="header-icon"
    >
      <g fill="var(--color-primary)">
        <ellipse cx="28" cy="44" rx="11" ry="14" />
        <ellipse cx="100" cy="44" rx="11" ry="14" />
        <ellipse cx="52" cy="28" rx="11" ry="14" />
        <ellipse cx="76" cy="28" rx="11" ry="14" />
        <path
          d="M64 56 c-20 0-36 15-36 32 c0 11 9 20 22 20h28 c13 0 22-9 22-20 c0-17-16-32-36-32z"
        />
      </g>
    </svg>

    <h1 id="form-title">{{ headerTitle }} Application</h1>
  </header>

  <p class="lede">
    {{ headerText }}
  </p>
</template>

<style scoped lang="css">
.form-header {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
  margin-bottom: 4px;
  color: var(--color-primary);

  & h1 {
    font-size: var(--font-size-h1);
    line-height: 1.2;
    letter-spacing: 0.2px;
    color: var(--color-primary);
  }

  /* SVG Icon Styling */
  & .header-icon {
    flex-shrink: 0;
    width: 80px;
    height: 80px;
  }

  @media (max-width: 440px) {
    flex-direction: column;
    align-items: center;
    gap: 8px;
    margin-bottom: 1rem;

    & h1 {
      font-size: 1.75rem;
      text-align: center;
    }

    & .header-icon {
      width: 50px;
      height: 50px;
    }
  }
}
.lede {
  color: var(--text-primary);
  margin: 6px auto 18px;
  font-size: 1.125rem;
  line-height: 1.6;
  max-width: 900px;
  user-select: none;
  font-weight: 400;
  @media (max-width: 440px) {
    font-size: 1rem;
    margin: 4px 0 12px;
  }
}
</style>
