<script setup lang="ts">
import { reactive } from 'vue'

import type { IPet } from '../../../models/common.ts'
import Drawer from '../../common/drawer/Drawer.vue'
import Button from '../../common/ui/Button.vue'
import InputField from '../../common/ui/InputField.vue'

const { pet, isDrawerOpen } = defineProps<{
  pet: IPet
  isDrawerOpen: boolean
}>()

const emit = defineEmits(['update:isDrawerOpen'])

const formData = reactive({
  firstName: '',
  lastName: '',
  email: '',
  phone: '',
  preferredDate: '',
  preferredTime: '',
})

const closeDrawer = () => {
  emit('update:isDrawerOpen', false)
}

const submitForm = () => {
  
  console.log('Form submitted:', formData)
  closeDrawer()
}
</script>

<template>
  <Drawer
    :modelValue="isDrawerOpen"
    @update:modelValue="closeDrawer"
    :header="`Schedule a Meet with ${pet.name}`"
  >
    <p>
      We’re excited to help you meet {{ pet.name }}! Our friendly volunteers are here to assist you
      in setting up a meet and greet. You’re welcome to visit our shelter during these hours:
    </p>
    <ul>
      <li>Monday through Friday: 10am to 12pm and 6pm to 8pm</li>
      <li>Weekends: 11am to 4pm</li>
    </ul>
    <p>
      If these times don’t work for you, no problem! Simply fill out the form below, and one of our
      volunteers will get in touch with you as soon as possible to arrange a time that works best
      for you.
    </p>

    <form>
      <InputField
        label="Your First Name:"
        placeholder="Enter your first name"
        type="text"
        v-model="formData.firstName"
      />
      <InputField
        label="Your Last Name:"
        placeholder="Enter your last name"
        type="text"
        v-model="formData.lastName"
      />
      <InputField
        label="Your Email:"
        placeholder="Enter your email"
        type="email"
        v-model="formData.email"
      />
      <InputField
        label="Your Phone Number:"
        placeholder="Enter your phone number"
        type="tel"
        v-model="formData.phone"
      />
      <InputField
        label="Preferred Date:"
        placeholder="Select a preferred date"
        type="date"
        v-model="formData.preferredDate"
      />
      <InputField
        label="Preferred Time:"
        placeholder="Select a preferred time"
        type="time"
        v-model="formData.preferredTime"
      />

      <button type="submit">Submit</button>
    </form>

    <p>We look forward to helping you find your new best friend!</p>

    <div class="actions">
      <Button title="Submit" color="green" @click="submitForm()" />
    </div>
  </Drawer>
</template>

<style scoped lang="css">
p {
  margin: 0;
  font-size: 1rem;
  color: var(--text-primary);
  margin-bottom: 1rem;
  font-weight: 300;
  &:last-of-type {
    margin-bottom: 2rem;
    text-align: center;
  }
}
ul {
  margin: 0 0 1rem 1.5rem;
  padding: 0;
  list-style-type: disc;
  color: var(--text-primary);
}
li {
  margin-bottom: 0.5rem;
}
.actions {
  display: flex;
  justify-content: center;
  margin-top: 1.5rem;
}
</style>
