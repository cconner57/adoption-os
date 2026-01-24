<script setup lang="ts">
import { ref } from 'vue'

const faqs = ref([
  {
    question: 'Can I adopt if I rent my home?',
    answer:
      'Yes, you can adopt if you rent your home. We recommend checking with your landlord and providing proof of pet approval if necessary.',
    expanded: true,
  },
  {
    question: 'Do you have a return policy?',
    answer:
      'Yes, we have a no-questions-asked return policy within the first 30 days of adoption to ensure the best fit for both the pet and the adopter.',
    expanded: true,
  },
  {
    question: 'Are there any age restrictions for adopters?',
    answer: 'Adopters must be at least 21 years old to adopt a pet from us.',
    expanded: true,
  },
  {
    question: "Can I return the pet if it doesn't work out?",
    answer: 'Yes, we have a no-questions-asked return policy within the first 30 days of adoption.',
    expanded: false,
  },
  {
    question: 'Can I foster a pet before adopting?',
    answer:
      'Yes, we encourage fostering as it helps pets adjust to home life and allows you to see if the pet is a good fit for your family.',
    expanded: false,
  },
  {
    question: 'Do you offer support after adoption?',
    answer:
      'Yes, we provide lifetime support to our adopters, including behavioral advice and resources.',
    expanded: false,
  },
])

const toggleFaq = (faqIndex: number) => {
  faqs.value[faqIndex].expanded = !faqs.value[faqIndex].expanded
}
</script>

<template>
  <div class="adopt-detail__adoption-faq">
    <h2>Adoption FAQs</h2>
    <ul>
      <li v-for="(faqItem, faqIndex) in faqs" :key="faqIndex">
        <p @click="toggleFaq(faqIndex)" class="faq-question">
          <svg
            :class="['arrow', faqItem.expanded ? 'rotated-arrow' : '']"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
            aria-hidden="true"
            role="img"
          >
            <path
              d="M9 6l6 6-6 6"
              stroke="currentColor"
              stroke-width="2.2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
          {{ faqItem.question }}
        </p>
        <p v-if="faqItem.expanded" class="faq-answer">{{ faqItem.answer }}</p>
      </li>
    </ul>
  </div>
</template>

<style scoped lang="css">
.adopt-detail__adoption-faq {
  width: 50%;

  ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  h2 {
    font-size: 1.5rem;
    margin-bottom: 16px;

    @media (width <= 440px) {
      font-size: 1.25rem;
    }
  }

  .faq-question {
    font-weight: 700;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 8px;

    &:hover {
      color: var(--color-secondary);
    }
  }

  .faq-answer {
    margin: 8px 0 16px 34px;
    color: var(--text-primary);
    line-height: 1.4;
    font-weight: 400;
  }

  @media (width >= 321px) and (width <= 430px) {
    width: 100%;
    margin-right: 0;

    ul {
      padding: 0;

      li {
        margin-bottom: 12px;
      }
    }

    .faq-question {
      font-size: 1rem;
    }

    .faq-answer {
      font-size: 0.95rem;
      margin-left: 32px;
    }
  }
}

.arrow {
  transition: transform 0.2s;
}

.rotated-arrow {
  transform: rotate(90deg);
}
</style>
