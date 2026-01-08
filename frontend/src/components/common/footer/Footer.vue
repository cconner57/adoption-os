<script setup lang="ts">
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>

<template>
  <footer class="footer">
    <div class="footer__social">
      <h3 class="footer__title">Connect with us on:</h3>
      <a href="#" aria-label="Facebook">Facebook</a>
      <a href="#" aria-label="Instagram">Instagram</a>
    </div>
    <div class="footer__info">
      <p>I Dream of Home Rescue, Inc. is a 501(c)(3) NON-PROFIT ORGANIZATION. TAX ID: 81-0780050</p>
      <p>I Dream of Home Rescue PO Box 7612, La Verne, California 91750</p>
    </div>
    <div class="footer__copyright">
      <p>&copy; {{ currentYear }} I Dream of Home Rescue. All rights reserved.</p>
    </div>
  </footer>
</template>

<style scoped lang="css">
.footer {
  background-color: var(--green);
  color: var(--font-color-light);
  padding: 40px 20px;
  text-align: center;
  border-top: 1px solid #ddd;
  display: flex;
  flex-direction: column;
  align-items: center;
  overflow: hidden;

  @media (max-width: 440px) {
    padding: 20px 10px;
    font-size: 14px;
  }
}

.footer__social {
  margin-bottom: 15px;
  display: flex;
  align-items: center;
  a {
    margin: 0 10px;
    color: var(--font-color-light) !important;
    text-decoration: none;
    font-weight: 500;
    padding-bottom: 5px;
  }
  a:hover {
    color: #ddd;
  }
}

.footer__title {
  font-size: 20px;
  font-weight: 800;
  margin-bottom: 10px;
  color: var(--font-color-light) !important;

  @media (max-width: 440px) {
    font-size: 16px;
    margin-bottom: 5px;
  }
}

.footer__info {
  margin-bottom: 15px;
  p {
    line-height: 1.5;
    font-weight: 400;
    margin: 4px 0;
    color: var(--font-color-light) !important;
  }

  @media (max-width: 440px) {
    font-size: 14px;
    margin-bottom: 10px;
    width: 90%;
    p {
      margin: 2px 0;
      font-size: 12px;
    }
  }
}

.footer__copyright {
  font-size: 12px;
  color: var(--font-color-light);
  p {
    line-height: 1.5;
    font-weight: 400;
    color: var(--font-color-light) !important;
  }

  @media (max-width: 440px) {
    font-size: 10px;
    p {
      line-height: 1.4;
      font-size: 10px;
    }
  }
}
</style>
