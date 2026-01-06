<script setup lang="ts">
import type { IPet } from '../../../models/common'

defineProps<{
  pet: IPet
}>()
</script>

<template>
  <div class="adopt-detail__adoption-process">
    <h2>Adoption Process</h2>
    <ol v-if="pet.profileSettings.adoptionProcess === 'cat'">
      <li>
        <span>Submit Your Application</span> <br />
        <p>
          Tell us about yourself, your home, and the kind of cat you’re hoping to welcome into your
          family.
        </p>
      </li>
      <li>
        <span> Share Your Home Photos or Video </span>
        <p>
          Help us get to know your space by sending a few photos or a quick video tour so we can
          make sure it’s a safe, happy place for your future feline friend.
        </p>
      </li>
      <li>
        <span>Adoption Fee & Final Steps</span>
        <p>
          Once your application is reviewed, we’ll work with you to finalize the adoption with a
          small fee that helps cover your cat’s care, vaccinations, and microchipping.
        </p>
      </li>
      <li>
        <span>Adoption Day! 🎉</span> <br />
        <p>
          Congratulations! Your new cat is ready to come home. We’ll be here to support you as they
          settle in and start exploring their new home.
        </p>
      </li>
    </ol>
    <ol v-else>
      <li>
        <span>Submit Your Application</span> <br />
        <p>
          Tell us a little about yourself, your home, and the kind of dog you’re hoping to welcome
          into your life.
        </p>
      </li>
      <li>
        <span>Chat With Our Team</span> <br />
        <p>
          After reviewing your application, a friendly volunteer will reach out for a quick
          conversation to get to know you and answer any questions.
        </p>
      </li>
      <li>
        <span>Virtual Home Visit</span> <br />
        <p>
          To ensure a safe and loving environment, we’ll take a virtual tour of your home via video
          call.
        </p>
      </li>
      <li>
        <span>Meet & Greet</span> <br />
        <p>
          Meet your potential new best friend! We’ll arrange a meet-and-greet with all household
          members (including pets) to ensure everyone gets along.
        </p>
      </li>
      <li>
        <span>Adoption Fee & Final Steps</span> <br />
        <p>
          Once it’s a match, we’ll finalize the adoption with a small fee that helps cover your
          dog’s care, vaccinations, and microchipping.
        </p>
      </li>
      <li>
        <span>Welcome Home! 🎉</span> <br />
        <p>
          Congratulations! Your new family member is ready to come home. We’re here to support you
          as your dog settles in and begins their new adventure.
        </p>
      </li>
    </ol>
  </div>
</template>

<style scoped lang="css">
.adopt-detail__adoption-process {
  width: 50%;
  margin-right: 20px;
  @media (max-width: 440px) {
    width: 100%;
    margin-right: 0;
  }
}
h2 {
  font-size: 1.5rem;
  margin-bottom: 16px;
  @media (max-width: 440px) {
    font-size: 1.25rem;
  }
}
ol {
  padding-left: 0;
  margin-bottom: 16px;
  list-style-type: decimal;
  list-style-position: inside;
}
li {
  margin-bottom: 8px;
  @media (max-width: 440px) {
    margin-bottom: 12px;
  }
}
li span {
  font-weight: bold;
  @media (max-width: 440px) {
    font-size: 1rem;
  }
}
li p {
  margin: 4px 0;
  color: var(--font-color-dark);
  line-height: 1.4;
  padding-left: 16px;
  font-weight: 400;
  @media (max-width: 440px) {
    font-size: 0.95rem;
  }
}
</style>
