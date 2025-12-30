<script setup lang="ts">
import InputTextArea from '../../common/ui/InputTextArea.vue'
import type { SurrenderFormState } from '../../../models/common'

const { formState } = defineProps<{
  formState: SurrenderFormState
}>()
</script>

<template>
  <div class="behavior-section">
    <h5>Behavior</h5>
    <div class="behavior-grid">
      <InputTextArea
        label="How does the cat usually behave towards people they do know?"
        placeholder="Describe behavior"
        :spanFull="false"
        :modelValue="formState.catsBehaviorTowardsKnownPeople"
        @update:modelValue="(val) => (formState.catsBehaviorTowardsKnownPeople = val)"
      />
      <InputTextArea
        label="How does the cat usually behave towards animals they do not know?"
        placeholder="Describe behavior"
        :spanFull="false"
        :modelValue="formState.catsBehaviorTowardsStrangers"
        @update:modelValue="(val) => (formState.catsBehaviorTowardsStrangers = val)"
      />
      <InputTextArea
        label="How does the cat usually behave towards animals they know"
        placeholder="Describe behavior"
        :spanFull="false"
        :modelValue="formState.catsBehaviorTowardsKnownAnimals"
        @update:modelValue="(val) => (formState.catsBehaviorTowardsKnownAnimals = val)"
      />
      <InputTextArea
        label="Comments on answers above"
        placeholder="Any comments"
        :spanFull="false"
        :modelValue="formState.commentsOnBehavior"
        @update:modelValue="(val) => (formState.commentsOnBehavior = val)"
      />
      <InputTextArea
        label="How does the cat usually react when an unfamiliar person approaches or enters the yard or house?"
        placeholder="Describe reaction"
        :spanFull="false"
        :modelValue="formState.catsReactionToNewPeople"
        @update:modelValue="(val) => (formState.catsReactionToNewPeople = val)"
      />
      <InputTextArea
        label="Is the cat housetrained?"
        placeholder="Yes/No"
        :spanFull="false"
        :modelValue="formState.catHouseTrained"
        @update:modelValue="(val) => (formState.catHouseTrained = val)"
      />
      <InputTextArea
        label="Where does the cat spend the majority of their time?"
        placeholder="e.g. Indoors"
        :spanFull="false"
        :modelValue="formState.catSpendMajorityOfTime"
        @update:modelValue="(val) => (formState.catSpendMajorityOfTime = val)"
      />
      <InputTextArea
        label="How many hours a day is the cat left alone without a human?"
        placeholder="e.g. 5 hours"
        :spanFull="false"
        :modelValue="formState.catLeftAloneDuration"
        @update:modelValue="(val) => (formState.catLeftAloneDuration = val)"
      />
      <InputTextArea
        label="When left alone, is he/she..."
        placeholder="e.g. Sleeps"
        :spanFull="false"
        :modelValue="formState.catWhenLeftAlone"
        @update:modelValue="(val) => (formState.catWhenLeftAlone = val)"
      />
      <InputTextArea
        label="When left alone, does the cat usually show any of the following behaviors?"
        placeholder="Describe behaviors"
        :spanFull="false"
        :modelValue="formState.catLeftAloneBehaviors"
        @update:modelValue="(val) => (formState.catLeftAloneBehaviors = val)"
      />
      <InputTextArea
        label="When the cat plays, do they typically..."
        placeholder="Describe play"
        :spanFull="false"
        :modelValue="formState.catHowItPlays"
        @update:modelValue="(val) => (formState.catHowItPlays = val)"
      />
      <InputTextArea
        label="What toys does the cat like?"
        placeholder="List toys"
        :spanFull="false"
        :modelValue="formState.catToysItLikes"
        @update:modelValue="(val) => (formState.catToysItLikes = val)"
      />
      <InputTextArea
        label="What games does the cat like?"
        placeholder="List games"
        :spanFull="false"
        :modelValue="formState.catGamesItLikes"
        @update:modelValue="(val) => (formState.catGamesItLikes = val)"
      />
      <InputTextArea
        label="Is the cat scared of anything?"
        placeholder="Yes/No (List fears)"
        :spanFull="false"
        :modelValue="formState.catScaredOfAnything"
        @update:modelValue="(val) => (formState.catScaredOfAnything = val)"
      />
      <InputTextArea
        label="If yes, please explain"
        placeholder="Explanation"
        :spanFull="false"
        :modelValue="formState.catScaredOfAnythingExplanation"
        @update:modelValue="(val) => (formState.catScaredOfAnythingExplanation = val)"
      />
      <InputTextArea
        label="Please tell us the cats bad habits"
        placeholder="Describe habits"
        :spanFull="false"
        :modelValue="formState.catBadHabits"
        @update:modelValue="(val) => (formState.catBadHabits = val)"
      />
      <InputTextArea
        label="Is the cat allowed on furniture?"
        placeholder="Yes/No"
        :spanFull="false"
        :modelValue="formState.catAllowedOnFurniture"
        @update:modelValue="(val) => (formState.catAllowedOnFurniture = val)"
      />
      <InputTextArea
        label="Where does the pet usually sleep at overnight?"
        placeholder="e.g. Bedroom"
        :spanFull="false"
        :modelValue="formState.catSleepAtNight"
        @update:modelValue="(val) => (formState.catSleepAtNight = val)"
      />
      <InputTextArea
        label="Describe the pets reaction if a person or other cat gets near or attempts to take food he/she is eating"
        placeholder="Describe reaction"
        :spanFull="false"
        :modelValue="formState.catBehaviorFoodOthers"
        @update:modelValue="(val) => (formState.catBehaviorFoodOthers = val)"
      />
      <InputTextArea
        label="Describe the pets reaction if a person or other cat attempts to take a toy they are playing with"
        placeholder="Describe reaction"
        :spanFull="false"
        :modelValue="formState.catBehaviorToysOthers"
        @update:modelValue="(val) => (formState.catBehaviorToysOthers = val)"
      />
      <InputTextArea
        label="Does the cat have problems riding in cars?"
        placeholder="Yes/No"
        :spanFull="false"
        :modelValue="formState.catProblemsRidingInCar"
        @update:modelValue="(val) => (formState.catProblemsRidingInCar = val)"
      />
      <InputTextArea
        label="If yes, please explain"
        placeholder="Explanation"
        :spanFull="false"
        :modelValue="formState.catProblemsRidingInCarExplanation"
        @update:modelValue="(val) => (formState.catProblemsRidingInCarExplanation = val)"
      />
      <InputTextArea
        label="Has the pet escaped your property 2 or more times in the last 6 months?"
        placeholder="Yes/No"
        :spanFull="false"
        :modelValue="formState.catEscapedBefore"
        @update:modelValue="(val) => (formState.catEscapedBefore = val)"
      />
      <InputTextArea
        label="If yes, please explain"
        placeholder="Explanation"
        :spanFull="false"
        :modelValue="formState.catEscapedBeforeExplanation"
        @update:modelValue="(val) => (formState.catEscapedBeforeExplanation = val)"
      />
    </div>
  </div>
</template>

<style scoped lang="css">
.behavior-section h5 {
  margin-bottom: 24px;
}

.behavior-grid {
  display: flex;
  flex-direction: column;
  gap: 16px;

  @media (min-width: 768px) {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    column-gap: 24px;
    row-gap: 16px;
  }
}

.full-width {
  @media (min-width: 768px) {
    grid-column: 1 / -1;
  }
}
</style>
