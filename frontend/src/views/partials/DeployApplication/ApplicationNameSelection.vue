<script setup>
import { preventSpaceInput } from '@/vendor/utils.js'
import FilledButton from '@/views/components/FilledButton.vue'
import { ref } from 'vue'
import { useLazyQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'

defineProps({
  finalApplicationNameAndMoveToNextTab: {
    type: Function,
    required: true
  }
})

// state
const newApplicationName = ref('')
// Is exist application name
const isExistApplicationName = ref(true)
const {
  loading: isExistApplicationNameLoading,
  load: isExistApplicationLoad,
  variables: isExistApplicationVariables,
  onResult: onIsExistApplicationNameResult
} = useLazyQuery(
  gql`
    query ($name: String!) {
      isExistApplicationName(name: $name)
    }
  `,
  {
    name: ''
  },
  {
    fetchPolicy: 'no-cache',
    nextFetchPolicy: 'no-cache'
  }
)

onIsExistApplicationNameResult((result) => {
  if (result.data?.isExistApplicationName !== null && result.data?.isExistApplicationName !== undefined) {
    isExistApplicationName.value = result.data?.isExistApplicationName === true
  }
})

const isExistApplicationNameCheck = () => {
  if (newApplicationName.value === '') {
    isExistApplicationName.value = true
    return
  }
  newApplicationName.value = newApplicationName.value.replace(/[^a-zA-Z0-9_-]/g, '')
  isExistApplicationVariables.value.name = newApplicationName.value
  isExistApplicationLoad()
}
</script>

<template>
  <div :key="0" class="mt-12 flex w-full max-w-md flex-col p-6">
    <!-- <img src="@/assets/images/deploy-app-image.png" class="mx-auto w-3/4" alt="deploy-app-image" /> -->
    <div class="mt-8">
      <label class="form-control w-full">
        <div class="label">
          <span class="label-text !font-semibold">Application Name</span>
        </div>
        <input id="application_name" v-model="newApplicationName" @keydown="preventSpaceInput" autocomplete="off"
          class="input input-bordered w-full" v-debounce:1000ms="isExistApplicationNameCheck" />
        <div v-if="
          newApplicationName !== '' &&
          !isExistApplicationNameLoading &&
          isExistApplicationVariables.name === newApplicationName
        ">
          <p v-if="isExistApplicationName" class="mt-1 text-sm text-error">
            <font-awesome-icon icon="fa-solid fa-circle-xmark" />
            Sorry, {{ newApplicationName }} name is not available
          </p>
          <p v-else-if="!isExistApplicationName" class="mt-1 text-sm text-success">
            <font-awesome-icon icon="fa-solid fa-circle-check" />
            {{ newApplicationName }} name is available
          </p>
        </div>
      </label>
    </div>
    <FilledButton :disabled="isExistApplicationName" :loading="isExistApplicationNameLoading" class="mt-4 w-full"
      type="primary" @click="() => finalApplicationNameAndMoveToNextTab(newApplicationName)">
      Check & Proceed to Next
    </FilledButton>
  </div>
</template>

<style scoped></style>
