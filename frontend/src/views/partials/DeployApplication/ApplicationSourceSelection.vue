<script setup>
import { ref } from 'vue'
import FilledButton from '@/views/components/FilledButton.vue'

defineProps({
  finalizeApplicationSourceAndMoveToNextTab: {
    type: Function,
    required: true
  }
})

const sourceType = ref('')
const setSourceType = (type) => {
  sourceType.value = type
}
</script>

<template>
  <div :key="1" class="flex h-full flex-col items-center justify-center space-y-6">
    <p class="text-center text-lg font-medium text-base-content">Select the source of your application</p>
    <div class="flex flex-col items-center justify-center gap-4">
      <!--  Docker Image  -->
      <div :class="{
        'border-primary': sourceType === 'image'
      }" class="card border cursor-pointer hover:shadow" @click.prevent="() => setSourceType('image')">
        <div class="card-body flex-row items-center p-4">
          <font-awesome-icon class="w-10 h-10" icon="fa-brands fa-docker" />
          <div class="flex flex-col ml-2">
            <h3 class="card-title align-self-start">Docker Image</h3>
            <p class="font-thin">Select a Docker image from Docker Hub</p>
          </div>
        </div>
      </div>

      <!--  Git   -->
      <div :class="{
        'border-primary': sourceType === 'git'
      }" class="card border cursor-pointer hover:shadow" @click.prevent="() => setSourceType('git')">
        <div class="card-body flex-row items-center p-4">
          <font-awesome-icon class="w-10 h-10" icon="fa-brands fa-git" />
          <div class="flex flex-col ml-2">
            <h3 class="card-title align-self-start">Git Repository</h3>
            <p class="font-thin">Point to an existing repo and make a copy</p>
          </div>
        </div>
      </div>

      <!--  Upload Code -->
      <div :class="{
        'border-primary': sourceType === 'sourceCode'
      }" class="card border cursor-pointer w-full hover:shadow" @click.prevent="() => setSourceType('sourceCode')">
        <div class="card-body flex-row items-center p-4">
          <font-awesome-icon class="w-10 h-10" icon="fa-solid fa-upload" />
          <div class="flex flex-col ml-2">
            <h3 class="card-title align-self-start">Source Code</h3>
            <p class="font-thin text-ellipsis">Upload your application code</p>
          </div>
        </div>
      </div>
    </div>

    <FilledButton :disabled="sourceType === ''" class="w-full" type="primary"
      @click="() => finalizeApplicationSourceAndMoveToNextTab(sourceType)">
      Proceed to Next Page
    </FilledButton>
  </div>
</template>