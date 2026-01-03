<script setup lang="ts">
import { ref } from 'vue'
import type { IPet } from '../../../models/common.ts'
import Button from '../../common/ui/Button.vue'
import Capsules from '../../common/ui/Capsules.vue'
import AdoptionFAQ from '../adopt-faq/AdoptionFAQ.vue'
import AdoptionProcess from '../adopt-process/AdoptionProcess.vue'
import { formatDate } from '../../../utils/common.ts'
import AdditionalInfo from '../additional-info/AdditionalInfo.vue'
import AdoptDrawer from './AdoptDrawer.vue'

const props = defineProps<{
  pet: IPet
}>()

const isDrawerOpen = ref(false)

const handleStartAdoption = () => {
  sessionStorage.setItem('adoption_pet_id', JSON.stringify({petId: props.pet.id, petName: props.pet.name}))
  globalThis.location.href = `/pet-adoption/${props.pet.id}`
}

// const handleScheduleMeet = () => {
//   isDrawerOpen.value = true
// }

// const handleRequestInformation = () => {
//   globalThis.location.href = `/pet-adoption/${props.pet.id}`
// }

const handleShare = () => {
  const shareData = {
    title: `Check out ${props.pet.name} for adoption!`,
    text: `I found ${props.pet.name} on IDOHR and thought you might be interested!`,
    url: globalThis.location.href,
  }
  if (navigator.share) {
    navigator
      .share(shareData)
      .then(() => console.log('Successful share'))
      .catch((error) => console.log('Error sharing', error))
  } else {
    alert('Sharing is not supported in this browser.')
  }
}

const imgError = ref(false)

function onImgError() {
  imgError.value = true
}
</script>

<template>
  <div class="adopt-detail">
    <div class="adopt-detail__main">
      <img
        v-if="!imgError"
        :src="pet.photos?.find((p) => p.isPrimary)?.url ?? ''"
        :alt="pet.name"
        @error="onImgError"
      />
      <div v-else class="img-fallback" aria-hidden="true"></div>
      <div class="adopt-detail__info">
        <div class="adopt-detail__info__main">
          <h1>{{ pet.name }}</h1>
          <div class="adopt-detail__traits">
            <Capsules v-if="pet?.species" :label="pet?.species" />
            <Capsules v-if="pet?.sex" :label="pet?.sex" />
            <Capsules
              v-if="pet?.physical?.dateOfBirth"
              :label="formatDate(pet?.physical?.dateOfBirth, true)"
            />
          </div>
          <p>{{ pet?.descriptions?.behavioral }}</p>
          <div class="adopt-detail__actions">
            <Button
              title="Start Adoption"
              color="blue"
              @click="handleStartAdoption"
              :fullWidth="true"
            />
            <!-- <Button
              title="Schedule a Meet"
              color="purple"
              @click="handleScheduleMeet"
              :fullWidth="true"
            /> -->
            <!-- <Button
              title="Request Information"
              color="orange"
              @click="handleRequestInformation"
              :fullWidth="true"
            /> -->
            <Button title="Share" color="green" @click="handleShare" :fullWidth="true" />
          </div>
        </div>
        <AdditionalInfo :pet="pet" />
      </div>
    </div>
    <div
      v-if="
        pet.descriptions?.fun ||
        pet.descriptions?.additionalInformation?.length ||
        pet.profileSettings.showAdditionalInformation
      "
      class="adopt-detail__about"
    >
      <div class="adopt-detail__about__content">
        <div v-if="pet.descriptions?.fun" class="adopt-detail__about__fun">
          <h2>From {{ pet.name }}</h2>
          <p>{{ pet.descriptions?.fun }}</p>
        </div>
        <div
          class="adopt-detail__about__additional-info"
          v-if="pet.profileSettings.showAdditionalInformation"
        >
          <h2>Additional Information</h2>
          <ul>
            <li v-for="(info, index) in pet.descriptions?.additionalInformation" :key="index">
              {{ info }}
            </li>
          </ul>
        </div>
      </div>
      <div class="adopt-detail__about__medical" v-if="pet.profileSettings.showMedicalHistory">
        <h2>Medical History</h2>
        <h3>Vaccinations</h3>
        <ul>
          <VaccinationItem name="Rabies" :date-administered="pet.medical?.vaccinations?.rabies?.dateAdministered" />
          <VaccinationItem name="Bordetella" :date-administered="pet.medical?.vaccinations?.bordetella?.dateAdministered" />
          <VaccinationItem name="Canine Distemper" :round1="pet.medical?.vaccinations?.canineDistemper?.round1?.dateAdministered" :round2="pet.medical?.vaccinations?.canineDistemper?.round2?.dateAdministered" :round3="pet.medical?.vaccinations?.canineDistemper?.round3?.dateAdministered" :isComplete="pet.medical?.vaccinations?.canineDistemper?.isComplete" />
          <VaccinationItem name="Feline Distemper" :round1="pet.medical?.vaccinations?.felineDistemper?.round1?.dateAdministered" :round2="pet.medical?.vaccinations?.felineDistemper?.round2?.dateAdministered" :round3="pet.medical?.vaccinations?.felineDistemper?.round3?.dateAdministered" :isComplete="pet.medical?.vaccinations?.felineDistemper?.isComplete" />
          <VaccinationItem name="Feline Leukemia" :round1="pet.medical?.vaccinations?.felineLeukemia?.round1?.dateAdministered" :round2="pet.medical?.vaccinations?.felineLeukemia?.round2?.dateAdministered" :round3="pet.medical?.vaccinations?.felineLeukemia?.round3?.dateAdministered" :isComplete="pet.medical?.vaccinations?.felineLeukemia?.isComplete" />
          <VaccinationItem name="Leptospira" :round1="pet.medical?.vaccinations?.leptospira?.round1?.dateAdministered" :round2="pet.medical?.vaccinations?.leptospira?.round2?.dateAdministered" :round3="pet.medical?.vaccinations?.leptospira?.round3?.dateAdministered" :isComplete="pet.medical?.vaccinations?.leptospira?.isComplete" />
          <VaccinationItem name="Other" :otherRounds="pet.medical?.vaccinations?.other" />
        </ul>

        <h3 v-if="pet.medical?.surgeries?.length">Surgeries</h3>
        <ul v-if="pet.medical?.surgeries?.length">
          <li v-for="(surgery, index) in pet.medical.surgeries" :key="'surgery-' + index">
            <p>{{ surgery.name }}</p>
            <p>
              {{
                surgery.date
                  ? 'Performed on ' + formatDate(surgery.date)
                  : 'Date not available'
              }}
            </p>
          </li>
        </ul>
      </div>
    </div>
    <div class="adopt-detail__adoption">
      <AdoptionProcess :pet="pet" />
      <AdoptionFAQ />
    </div>
  </div>
  <!-- <MoreFriends :pet="pet" /> -->
  <AdoptDrawer
    :pet="pet"
    :isDrawerOpen="isDrawerOpen"
    @update:isDrawerOpen="isDrawerOpen = $event"
  />
</template>

<style scoped lang="css">
.adopt-detail {
  width: 100%;

  .adopt-detail__main {
    display: flex;
    gap: 30px;
    width: 100%;

    img {
      flex: 3;
      width: 0;
      min-width: 0;
      height: 600px;
      object-fit: cover;
      border-radius: 16px;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
    }
    & .img-fallback {
      flex: 3;
      width: 0;
      min-width: 0;
      height: 600px;
      border-radius: 16px;
      background: url('/images/paw.svg') center/100px 100px no-repeat #add8e6;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
    }
    .adopt-detail__info {
      display: flex;
      flex-direction: column;
      gap: 20px;
      background-color: var(--white);
      color: var(--font-color-dark);
      padding: 32px;
      border-radius: 16px;
      flex: 2;
      width: 0;
      min-width: 0;
      height: auto;
      min-height: 600px;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
      @media (max-width: 1024px) {
        width: 100%;
        min-height: auto;
      }
      .adopt-detail__info__main {
        display: flex;
        flex-direction: column;
        gap: 12px;
        border-bottom: 1px solid rgb(178, 177, 177);
        padding-bottom: 20px;
        h1 {
          font-size: 2.5rem; /* Larger title */
        }
        @media (min-width: 321px) and (max-width: 430px) {
          h1 {
            font-size: 1.5rem;
          }
        }
        .adopt-detail__traits {
          display: flex;
          flex-direction: row;
          gap: 10px;
          flex-wrap: wrap;
          & p {
            background-color: var(--green-weak);
            padding: 4px 12px;
            border-radius: 16px;
          }
        }
        .adopt-detail__actions {
          display: grid;
          grid-template-columns: 1fr 1fr;
          gap: 14px;
          flex-wrap: wrap; /* Allow wrapping */
          @media (max-width: 440px) {
            display: flex;
            flex-direction: column;
          }
        }
      }
    }
    @media (max-width: 1024px) {
        flex-direction: column;
        img, .img-fallback {
            width: 100%;
            height: 400px;
            flex: auto;
        }
        .adopt-detail__info {
            width: 100%;
            flex: auto;
        }
    }
  }
  .adopt-detail__about {
    display: flex;
    margin-top: 20px;
    background-color: var(--white);
    padding: 32px;
    border-radius: 16px;
    color: var(--font-color-dark);
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
    width: 100%;
    & .adopt-detail__about__content {
      width: 50%;
      margin-right: 20px;
      & .adopt-detail__about__fun {
        width: 100%;
        height: 50%;
      }
      & .adopt-detail__about__additional-info {
        margin-top: 2rem;
        ul {
          padding-left: 20px;
          list-style: disc;
          li {
            margin-bottom: 8px;
          }
        }
      }
    }
    & .adopt-detail__about__medical {
      width: 50%;
      ul {
        margin-bottom: 16px;
        li {
          margin-bottom: 8px;
          display: flex;
          border-bottom: 1px solid rgb(178, 177, 177);
          width: 100%; /* Full width */
          justify-content: space-between; /* Space out */
          p {
            margin-bottom: 8px;
          }
          p:first-child {
            margin-right: 8px;
            /* width: 300px; removed fixed width */
            flex: 1;
          }
          p:last-child {
            font-weight: bold;
          }
        }
        li:last-of-type {
          border-bottom: none;
          margin-bottom: 0px;
        }
      }
    }
    h2 {
      font-size: 1.5rem;
      margin-bottom: 16px;
    }
    p {
      font-size: 1rem;
      line-height: 1.5;
      margin-bottom: 12px;
    }
    @media (max-width: 768px) {
      flex-direction: column;
      & .adopt-detail__about__content,
      & .adopt-detail__about__medical {
        width: 100%;
        margin-right: 0px;
      }
      & .adopt-detail__about__medical {
        margin-top: 2rem;
      }
    }
  }
  .adopt-detail__adoption {
    display: flex;
    margin-top: 30px;
    background-color: var(--white);
    color: var(--font-color-dark);
    padding: 32px;
    border-radius: 16px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.25);
    width: 100%;
    @media (max-width: 768px) {
      flex-direction: column;
    }
  }

  /* Mobile padding fix */
  @media (max-width: 430px) {
    padding: 0; /* Let parent handle padding */
  }
}
</style>
