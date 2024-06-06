<script setup>
import FilledButton from '@/views/components/FilledButton.vue'
import ModalDialog from '@/views/components/ModalDialog.vue'
import { computed, reactive, ref } from 'vue'
import { useLazyQuery, useMutation } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { useToast } from 'vue-toastification'
import CreateDomainModal from '@/views/partials/CreateDomainModal.vue'

const props = defineProps({
  callbackOnCreate: {
    type: Function,
    required: false,
    default: () => { }
  },
  callbackOnPop: {
    type: Function,
    required: false,
    default: () => { }
  }
})

const toast = useToast()

const isModalOpen = ref(false)
const specificApplicationId = ref('')
const isSpecificApplicationChosen = computed(() => specificApplicationId.value !== '')

// Create ingress rule
const newIngressRuleDetails = reactive({
  targetType: 'application',
  protocol: 'http',
  domainId: 0,
  port: 80,
  targetPort: 80,
  applicationId: '',
  externalService: ''
})

const openModal = (appId) => {
  newIngressRuleDetails.protocol = 'http'
  newIngressRuleDetails.domainId = 0
  newIngressRuleDetails.port = 80
  newIngressRuleDetails.applicationId = ''
  newIngressRuleDetails.targetPort = 80
  isModalOpen.value = true
  if (typeof appId === 'string' && appId !== '') {
    specificApplicationId.value = appId
    newIngressRuleDetails.applicationId = appId
  } else {
    specificApplicationId.value = ''
  }
  fetchApplications()
  fetchDomains()
}
const closeModal = () => {
  isModalOpen.value = false
  props.callbackOnPop()
}

const configureExternalService = () => {
  newIngressRuleDetails.targetType = 'externalService'
  newIngressRuleDetails.applicationId = ''
}

const configureApplication = () => {
  newIngressRuleDetails.targetType = 'application'
  newIngressRuleDetails.externalService = ''
}

const {
  mutate: createIngressRule,
  loading: isIngressRuleCreating,
  onDone: onIngressRuleCreateSuccess,
  onError: onIngressRuleCreateFail
} = useMutation(
  gql`
    mutation ($input: IngressRuleInput!) {
      createIngressRule(input: $input) {
        id
      }
    }
  `,
  {
    variables: {
      input: newIngressRuleDetails
    }
  }
)

onIngressRuleCreateSuccess(() => {
  closeModal()
  toast.success('Ingress Rule created successfully')
  props.callbackOnCreate()
})

onIngressRuleCreateFail((err) => {
  toast.error(err.message)
})

const onChangeProtocol = () => {
  if (newIngressRuleDetails.protocol === 'https') {
    newIngressRuleDetails.port = 443
  } else if (newIngressRuleDetails.protocol === 'http') {
    newIngressRuleDetails.port = 80
  } else {
    newIngressRuleDetails.port = 81
    newIngressRuleDetails.domainId = 0
  }
}

// Fetch domains from the server
const {
  result: domainListResult,
  load: loadDomains,
  refetch: refetchDomains
} = useLazyQuery(gql`
  query {
    domains {
      id
      name
    }
  }
`)
const domains = computed(() => domainListResult.value?.domains ?? [])

const fetchDomains = () => {
  if (loadDomains() === false) {
    refetchDomains()
  }
}

// Fetch applications from the server
const {
  result: applicationListResult,
  load: loadApplications,
  refetch: refetchApplications
} = useLazyQuery(gql`
  query {
    applications {
      id
      name
    }
  }
`)
const applications = computed(() => applicationListResult.value?.applications ?? [])

const fetchApplications = () => {
  if (loadApplications() === false) {
    refetchApplications()
  }
}

const reopenModal = () => {
  isModalOpen.value = true
}

// Create Domain
const createDomainModalRef = ref(null)
const openNewDomainModal = () => {
  if (!createDomainModalRef.value?.openModal) return
  isModalOpen.value = false
  createDomainModalRef.value.openModal()
}

defineExpose({
  openModal,
  closeModal
})
</script>

<template>
  <CreateDomainModal ref="createDomainModalRef" :callback-on-create="fetchDomains" :callback-on-pop="reopenModal" />
  <teleport to="body">
    <ModalDialog :close-modal="closeModal" :is-open="isModalOpen" width="lg">
      <template v-slot:header>Create Ingress Rule</template>
      <template v-slot:body>
        Enter the details of the new ingress rule.
        <form @submit.prevent="">
          <!-- Domains -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Ingress Info</span>
              </div>
              <div class="flex space-x-2">
                <select v-model="newIngressRuleDetails.protocol" class="select select-bordered w-full"
                  @change="onChangeProtocol">
                  <option value="http">HTTP</option>
                  <option value="https">HTTPS</option>
                  <option value="tcp">TCP</option>
                  <option value="udp">UDP</option>
                </select>
                <select v-show="newIngressRuleDetails.protocol === 'http' || newIngressRuleDetails.protocol === 'https'"
                  v-model="newIngressRuleDetails.domainId" class="select select-bordered w-full">
                  <option value="0">Select a domain</option>
                  <option v-for="domain in domains" :key="domain.id" :value="domain.id">{{ domain.name }}</option>
                </select>
                <p v-show="newIngressRuleDetails.protocol === 'tcp' || newIngressRuleDetails.protocol === 'udp'"
                  class="flex w-full items-center justify-end pr-2 text-base font-medium text-gray-700">
                  Listen on port
                </p>
                <input v-model="newIngressRuleDetails.port" :readonly="newIngressRuleDetails.protocol === 'https'"
                  autocomplete="off" class="input input-bordered w-full" placeholder="Port" type="number" />
              </div>
              <p v-if="newIngressRuleDetails.protocol === 'tcp' || newIngressRuleDetails.protocol === 'udp'"
                class="mt-2 text-sm text-danger-500">
                <b>NOTE: </b>You don't need to specify the domain for TCP and UDP protocols. While connecting to the
                server, use the server IP address instead of the domain name.
              </p>
              <p v-else class="mt-2 flex items-center text-sm">
                Need to create a domain?
                <a @click="openNewDomainModal" class="text-primary ml-1.5 cursor-pointer font-bold">Register New
                  Domain</a>
              </p>
            </label>
          </div>

          <div class="mt-4 w-full text-center text-xl">
            <font-awesome-icon icon="fa-solid fa-arrow-down" />
          </div>

          <div class="mt-2" v-show="newIngressRuleDetails.targetType === 'application'">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Application Name</span>
              </div>
              <div class="flex space-x-2">
                <select v-model="newIngressRuleDetails.applicationId" :disabled="isSpecificApplicationChosen"
                  class="select select-bordered w-full" :class="{
                    'cursor-not-allowed': isSpecificApplicationChosen
                  }">
                  <option value="">Select application name</option>
                  <option v-for="application in applications" :key="application.id" :value="application.id">
                    {{ application.name }}
                  </option>
                </select>
                <input v-model="newIngressRuleDetails.targetPort" class="input input-bordered w-full" placeholder="Port"
                  type="number" />
              </div>
            </label>
          </div>
          <p v-show="newIngressRuleDetails.targetType === 'application' && !isSpecificApplicationChosen"
            class="mt-2 flex items-center text-sm">
            Need to expose some external services ?
            <a @click="configureExternalService" class="text-primary ml-1.5 cursor-pointer font-bold">Click here</a>
          </p>
          <div class="mt-2" v-show="newIngressRuleDetails.targetType === 'externalService'">
            <p class="block text-sm font-medium text-gray-700">External Service Name</p>
            <div class="mt-1 flex space-x-2">
              <input v-model="newIngressRuleDetails.externalService" class="input input-bordered w-full" />
              <input v-model="newIngressRuleDetails.targetPort" class="input input-bordered w-full" placeholder="Port"
                type="number" />
            </div>
          </div>
          <p v-show="newIngressRuleDetails.targetType === 'externalService'" class="mt-2 flex items-center text-sm">
            Need to expose deployed application ?
            <a @click="configureApplication" class="text-primary ml-1.5 cursor-pointer font-bold">Click here</a>
          </p>
        </form>
      </template>
      <template v-slot:footer>
        <FilledButton :click="createIngressRule" :loading="isIngressRuleCreating" type="primary">
          Create Now
        </FilledButton>
      </template>
    </ModalDialog>
  </teleport>
</template>
