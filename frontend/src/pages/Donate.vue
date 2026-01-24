<script setup lang="ts">
import Footer from '../components/common/footer/Footer.vue'
</script>

<template>
  <div class="donate-page">
    <div class="header">
      <h1>Help Us Save Lives</h1>
      <p class="mission-text">
        You can save a life today! Your donation enables IDOHR to continue our lifesaving
        mission—finding loving homes for wonderful rescued shelter dogs and cats. As a small but
        dedicated operation, we face significant expenses including veterinary care, food, shelter,
        and daily care for all the animals in our rescue. Every contribution, no matter the size,
        makes a real difference. We are a 501(c)(3) non-profit organization, which means your
        donation is 100% tax-deductible.
      </p>
    </div>

    <main class="donate-container">
      <section class="donate-card cat-card">
        <h2>Support Our Cats</h2>
        <div class="illustration">
          <img src="/images/shadow.jpeg" alt="Cute Cat" width="500" height="375" />
        </div>

        <div class="payment-methods">
          <div class="qr-section">
            <div class="qr-placeholder">
              <span class="username">lorraine marrero</span>
              <img
                src="/images/cat-paypal.jpg"
                class="real-qr"
                alt="PayPal QR Code"
                width="400"
                height="402"
              />
              <span>Scan for PayPal</span>
            </div>
          </div>

          <a
            href="https://www.paypal.com/qrcodes/p2pqrc/FZYN4ZRPJWGXQ"
            target="_blank"
            rel="noopener noreferrer"
            class="donate-btn btn-paypal"
          >
            <i class="icon-paypal"></i> Donate via PayPal
          </a>

          <div class="qr-section">
            <div class="qr-placeholder">
              <span class="username">Lorraine Marrero-Baeza</span>
              <span class="handle">@Lorraine-Marrero-Baeza</span>
              <img
                src="/images/cat-venmo.jpg"
                class="real-qr"
                alt="Venmo QR Code"
                width="400"
                height="403"
              />
              <span>Scan for Venmo</span>
            </div>
          </div>

          <a
            href="https://venmo.com/code?user_id=3025264426614784345&created=1762028642"
            target="_blank"
            rel="noopener noreferrer"
            class="donate-btn btn-venmo"
          >
            <i class="icon-venmo"></i> Donate via Venmo
          </a>
        </div>
      </section>

      <section class="donate-card dog-card">
        <h2>Support Our Dogs</h2>
        <div class="illustration">
          <img src="/images/paw.svg" alt="Dog Paw" />
        </div>

        <div class="payment-methods">
          <div class="qr-section">
            <div class="qr-placeholder">
              <img src="/images/paw.svg" class="qr-icon" alt="QR Code" />
              <span>Scan for PayPal</span>
            </div>
          </div>

          <a
            href="https://www.paypal.com"
            target="_blank"
            rel="noopener noreferrer"
            class="donate-btn btn-paypal"
          >
            <i class="icon-paypal"></i> Donate via PayPal
          </a>

          <div class="qr-section">
            <div class="qr-placeholder">
              <img src="/images/paw.svg" class="qr-icon" alt="QR Code" />
              <span>Scan for Venmo</span>
            </div>
          </div>

          <a href="#" class="donate-btn btn-venmo"> <i class="icon-venmo"></i> Donate via Venmo </a>
        </div>
      </section>
    </main>

    <Footer />
  </div>
</template>

<style scoped>
.donate-page {
  background-color: var(--color-primary);
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  text-align: center;
  padding: 140px 20px 40px;
}

.header h1 {
  color: var(--text-inverse);
  font-size: 3rem;
  margin: 0;
  line-height: 1.2; 
}

.header .mission-text {
  text-align: center;
}

.mission-text {
  color: var(--text-inverse);
  font-size: 1.15rem;
  line-height: 1.6;
  max-width: 840px;
  margin: 24px auto 0;
  opacity: 0.95;
  font-weight: 500;
}

.donate-container {
  display: flex;
  justify-content: center;
  gap: 48px;
  padding: 0 var(--layout-padding-side) 100px;
  flex-wrap: wrap;
  flex: 1;
  max-width: 1400px;
  margin: 0 auto;
  width: 100%;
}

.donate-card {
  background: var(--text-inverse);
  border-radius: 24px;
  padding: 40px;
  flex: 1;
  min-width: 400px;
  max-width: 600px;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-shadow: 0 10px 30px rgb(0 0 0 / 10%);
  transition: all 0.3s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.donate-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 16px 40px rgb(0 0 0 / 12%);
}

.donate-card h2 {
  font-size: 2rem;
  margin-bottom: 24px;
  font-weight: 700;
  text-align: center; 
}

.illustration img {
  height: 200px;
  width: 200px;
  object-fit: cover;
  border-radius: 50%;
  margin-bottom: 32px;
  border: 6px solid;
  transition: transform 0.3s ease;
}

.donate-card:hover .illustration img {
  transform: scale(1.05);
}

.payment-methods {
  display: flex;
  flex-direction: column;
  gap: 16px;
  width: 100%;
}

.donate-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16px 24px;
  border-radius: 50px;
  text-decoration: none;
  font-weight: 600;
  font-size: 1.1rem;
  transition: all 0.2s cubic-bezier(0.2, 0.8, 0.2, 1);
  color: #fff;
  box-shadow:
    0 4px 6px rgb(50 50 93 / 11%),
    0 1px 3px rgb(0 0 0 / 8%);
  text-align: center;
}

.donate-btn:hover {
  transform: translateY(-2px);
  box-shadow:
    0 7px 14px rgb(50 50 93 / 10%),
    0 3px 6px rgb(0 0 0 / 8%);
  filter: brightness(1.1);
}

.donate-btn:active {
  transform: translateY(1px);
  box-shadow: none;
}

.qr-section {
  display: flex;
  justify-content: center;
  margin: 8px 0;
  width: 100%; 
}

.qr-placeholder {
  width: 100%;
  max-width: 240px;
  height: auto;
  padding: 16px;
  background: transparent;
  border-radius: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: auto;
  color: #666;
  font-size: 1rem;
  text-align: center;
  gap: 12px;
  border: none;
  box-shadow: inset 0 2px 6px rgb(0 0 0 / 3%);
}

.qr-placeholder > span:not(.username, .handle) {
  font-weight: 600;
  color: #222;
  letter-spacing: 0.5px;
}

.username {
  font-weight: 700;
  color: #333;
  font-size: 0.9rem;
  word-break: break-word; 
}

.handle {
  font-size: 0.85rem;
  color: #666;
  word-break: break-all; 
}

.qr-icon {
  width: 32px;
  height: 32px;
  opacity: 0.5;
}

.real-qr {
  width: 100%; 
  max-width: 180px; 
  height: auto; 
  object-fit: contain;
  border-radius: 4px;
}

.cat-card h2 {
  color: var(--color-tertiary);
}

.cat-card .illustration img {
  border-color: hsl(from var(--color-tertiary) h s 90%);
}

.cat-card .qr-placeholder {
  background-color: var(--color-tertiary-weak);
  color: var(--color-tertiary);
  border: 2px solid hsl(from var(--color-tertiary) h s 90%);
}

.btn-paypal {
  background-color: var(--color-tertiary);
}

.dog-card h2 {
  color: var(--color-secondary);
}

.dog-card .illustration img {
  border-color: var(--color-secondary-light);
  padding: 40px;
  box-sizing: border-box;
  background-color: var(--color-secondary-weak);
  border-width: 6px;
  box-shadow: inset 0 0 20px rgb(0 0 0 / 5%);
  filter: brightness(0) saturate(100%) invert(42%) sepia(97%) saturate(1466%) hue-rotate(198deg)
    brightness(99%) contrast(94%);
}

.dog-card .qr-placeholder {
  background-color: var(--color-secondary-weak);
  color: var(--color-secondary);
  padding: 24px;
  border: 2px solid var(--color-secondary-light);
}

.dog-card .qr-icon {
  width: 80px;
  height: 80px;
  opacity: 0.9;
  margin-bottom: 12px;
  filter: brightness(0) saturate(100%) invert(42%) sepia(97%) saturate(1466%) hue-rotate(198deg)
    brightness(99%) contrast(94%);
  border: 2px solid var(--color-secondary-light);
  padding: 12px;
  border-radius: 12px;
  background-color: #fff;
}

.dog-card .btn-paypal {
  background-color: var(--color-secondary);
}

.cat-card .btn-venmo {
  background-color: var(--color-tertiary);
}

.dog-card .btn-venmo {
  background-color: var(--color-secondary);
}

@media (width <= 1024px) {
  .donate-container {
    flex-direction: column;
    align-items: center;
    padding-bottom: 60px; 
  }

  .header {
    padding-top: 60px;
  }
}

@media (width <= 480px) {
  .header {
    padding: 110px 20px 30px; 
  }

  .header h1 {
    font-size: 2rem; 
    margin-top: 10px;
  }

  .mission-text {
    font-size: 1rem; 
    padding: 0 10px; 
  }

  .donate-container {
    padding: 0 16px 60px; 
  }

  .donate-card {
    padding: 24px 16px; 
    border-radius: 16px;
    min-width: 0; 
    width: 100%;
  }

  .donate-card h2 {
    font-size: 1.5rem;
  }

  .illustration img {
    height: 140px; 
    width: 140px;
    margin-bottom: 20px;
  }

  .qr-placeholder {
    min-height: 240px; 
  }
}
</style>
