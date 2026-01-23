<script setup lang="ts">
import { computed } from 'vue'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const props = defineProps<{ name: string; pet: any }>()

const formatDate = (dateString: string) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleDateString()
}

const vaccinationRecord = computed(() => props.name === 'rabies')
const vaccinationSeries = computed(() => props.name === 'felineDistemper')
</script>

<template>
  <div class="vaccination-item">
    <template v-if="vaccinationRecord">
      <li v-if="pet.medical?.vaccinations?.rabies">
            <p>Rabies</p>
            <p>
              {{
                pet.medical.vaccinations.rabies.dateAdministered
                  ? 'Administered on ' + formatDate(pet.medical.vaccinations.rabies.dateAdministered)
                  : 'Not administered'
              }}
            </p>
          </li>
    </template>
    <template v-if="vaccinationSeries">
      <li v-if="pet.medical?.vaccinations?.felineDistemper">
            <p>Feline Distemper ({{ pet.medical.vaccinations.felineDistemper.isComplete ? 'Complete' : 'Incomplete' }})</p>
            <p>
              <span v-if="pet.medical.vaccinations.felineDistemper.round1">
                Round 1: {{ formatDate(pet.medical.vaccinations.felineDistemper.round1.dateAdministered) }}
              </span>
              <span v-if="pet.medical.vaccinations.felineDistemper.round2">
                , Round 2: {{ formatDate(pet.medical.vaccinations.felineDistemper.round2.dateAdministered) }}
              </span>
              <span v-if="pet.medical.vaccinations.felineDistemper.round3">
                , Round 3: {{ formatDate(pet.medical.vaccinations.felineDistemper.round3.dateAdministered) }}
              </span>
            </p>
          </li>
        </template>
  </div>
</template>

<style scoped lang="css">
.vaccination-item {
}
</style>
