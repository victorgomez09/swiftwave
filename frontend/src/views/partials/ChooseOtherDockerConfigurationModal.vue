<script setup>
import ModalDialog from '@/views/components/ModalDialog.vue'
import { computed, ref } from 'vue'
import { useLazyQuery, useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { useToast } from 'vue-toastification'
import FilledButton from '@/views/components/FilledButton.vue'

const props = defineProps({
  onApplyConfiguration: {
    type: Function,
    required: true
  }
})

const toast = useToast()
const isModalOpen = ref(false)
const choosenDockerConfig = ref('')

const openModal = () => {
  choosenDockerConfig.value = ''
  isModalOpen.value = true
}
const closeModal = () => {
  isModalOpen.value = false
  choosenDockerConfig.value = ''
}

// All available docker configurations
const { result: availableDockerConfigsRawResult, onError: onAvailableDockerConfigsError } = useQuery(gql`
  query {
    availableDockerConfigs
  }
`)

onAvailableDockerConfigsError((err) => {
  toast.error(err.message)
})

const availableDockerConfigs = computed(() => availableDockerConfigsRawResult.value?.availableDockerConfigs ?? [])

// Fetch docker configurations
const {
  load: fetchDockerConfigurationsRaw,
  refetch: refetchDockerConfigurationsRaw,
  loading: isDockerConfigurationsLoading,
  onResult: onDockerConfigurationsResult,
  onError: onDockerConfigurationsError,
  variables: dockerConfigurationsVariables
} = useLazyQuery(gql`
  query ($serviceName: String!) {
    dockerConfigFromServiceName(serviceName: $serviceName) {
      detectedServiceName
      dockerFile
      dockerBuildArgs {
        description
        defaultValue
        key
      }
    }
  }
`)

const fetchDockerConfigurations = (serviceName) => {
  dockerConfigurationsVariables.value = { serviceName: serviceName }
  if (fetchDockerConfigurationsRaw() === false) {
    refetchDockerConfigurationsRaw()
  }
}

onDockerConfigurationsError((err) => {
  toast.error(err.message)
})

onDockerConfigurationsResult((result) => {
  if (!result.data || !result.data.dockerConfigFromServiceName) return
  props.onApplyConfiguration(result.data.dockerConfigFromServiceName)
  closeModal()
})

const fetchAndApplyConfiguration = () => {
  if (choosenDockerConfig.value === '') {
    return
  }
  fetchDockerConfigurations(choosenDockerConfig.value)
}

defineExpose({
  openModal,
  closeModal
})
</script>

<template>
  <teleport to="body">
    <ModalDialog :close-modal="closeModal" :is-open="isModalOpen" width="lg">
      <template v-slot:header>Pick Service Configuration</template>
      <template v-slot:body>
        <i>If you feel the detected service is incorrect, choose the correct one from the list. </i>
        <p class="mt-1">
          <b>Note:</b> Also you can just tap on <b>View / Modify Dockerfile</b> button to add custom configuration.
        </p>

        <!-- Type Field -->
        <div class="mt-4">
          <label class="form-control w-full">
            <div class="label">
              <span class="label-text">Available Service Configuration</span>
            </div>
            <select id="selected_docker_config" v-model="choosenDockerConfig" class="select select-bordered w-full">
              <option selected value="">Choose From List</option>
              <option v-for="dockerConfig in availableDockerConfigs" :key="dockerConfig" :value="dockerConfig">
                {{ dockerConfig }}
              </option>
            </select>
          </label>
        </div>
      </template>
      <template v-slot:footer>
        <FilledButton type="primary" :disabled="choosenDockerConfig === ''" :loading="isDockerConfigurationsLoading"
          class="w-full" :click="fetchAndApplyConfiguration">
          Fetch & Apply Configuration
        </FilledButton>
      </template>
    </ModalDialog>
  </teleport>
</template>

<style scoped></style>
