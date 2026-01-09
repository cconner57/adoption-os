<script setup lang="ts">
import { useRouter } from 'vue-router'
import { goToAdopt } from '../../../utils/navigate.ts'
import Button from '../ui/Button.vue'
import { ref, type PropType } from 'vue'
import Capsules from '../ui/Capsules.vue'
import { useMetrics } from '../../../composables/useMetrics'

const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  id: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: false,
  },
  capsules: {
    type: Array as PropType<string[]>,
    required: false,
    default: () => [],
  },
  photo: {
    type: String as PropType<string | null>,
    required: false,
  },
  size: {
    type: String as PropType<'small' | 'medium' | 'large'>,
    required: false,
    default: 'medium',
  },
  priority: {
    type: Boolean,
    default: false,
  },
})
const router = useRouter()

const imgError = ref(false)

function onImgError() {
  imgError.value = true
}

const { submitMetric } = useMetrics()

function handleAdopt() {
  submitMetric('spotlight_click', { petId: props.id, petName: props.name })
  goToAdopt(router, props.id.toLowerCase())
}
</script>

<template>
  <div class="pet-item" :style="{ viewTransitionName: `pet-card-${props.id}` }">
    <img
      v-if="!imgError"
      :src="`/images/${props.photo ?? ''}`"
      :alt="props.name"
      height="250"
      width="240"
      :style="{ viewTransitionName: 'pet-' + props.id }"
      :fetchpriority="priority ? 'high' : 'auto'"
      @error="onImgError"
      @click="handleAdopt"
    />
    <div v-else class="img-fallback" aria-hidden="true" @click="handleAdopt"></div>
    <div class="info-section">
      <h3>{{ props.name }}</h3>
      <div v-if="props.capsules.length > 0" class="capsules">
        <template v-for="capText in props.capsules" :key="capText">
          <Capsules v-if="capText && capText !== 'Invalid Date'">{{ capText }}</Capsules>
        </template>
      </div>
      <p v-if="props.description">{{ props.description }}</p>
      <div class="adopt-button">
        <Button title="Adopt Me" color="blue" @click="handleAdopt" :fullWidth="true" />
      </div>
    </div>
  </div>
</template>

<style scoped lang="css">
.pet-item {
  display: flex;
  flex-direction: column;
  gap: 12px;
  width: 280px;
  border-radius: 8px;
  overflow: hidden;
  overflow: hidden;
  background-color: var(--text-inverse);
  color: var(--text-primary);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
  height: 400px;

  img {
    width: 100%;
    object-fit: cover;
    height: 180px;
    background: url('/images/paw.svg') 90px 60px/100px 100px no-repeat #add8e6;
    cursor: pointer;
  }

  .img-fallback {
    width: 100%;
    height: 180px;
    background: url('/images/paw.svg') 90px 60px/100px 100px no-repeat #add8e6;
    cursor: pointer;
  }

  .info-section {
    display: flex;
    flex-direction: column;
    padding: 0 20px 16px;
    flex: 1;
    overflow: hidden;
  }

  h3 {
    font-size: 1.5rem;
    margin-bottom: 4px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    line-height: 1.6;
    padding: 4px 0 4px 0;
    flex-shrink: 0;
  }

  .capsules {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
    margin-bottom: 16px;
  }

  p {
    font-size: 1rem;
    flex-grow: 0;
    margin-bottom: 8px;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
    overflow: hidden;
    line-height: 1.5;
    padding-bottom: 2px;
  }

  .adopt-button {
    margin-top: auto;
  }

  @media (min-width: 321px) and (max-width: 430px) {
    & .img-fallback {
      background: url('/images/paw.svg') center center/100px 100px no-repeat #add8e6;
    }
    & .info-section {
      & .capsules {
        margin-bottom: 8px;
        gap: 2px;
      }
    }
  }

  @media (min-width: 1025px) and (max-width: 1440px) {
    width: 240px;
    height: 360px;
    & .img-fallback {
      background: url('/images/paw.svg') center center/100px 100px no-repeat #add8e6;
    }
    & .info-section {
      & .capsules {
        margin-bottom: 12px;
        gap: 2px;
      }
    }
  }

  @media (min-width: 1441px) {
    width: 260px;
    height: 380px;
    & .img-fallback {
      background: url('/images/paw.svg') center center/100px 100px no-repeat #add8e6;
    }
  }
}
</style>
