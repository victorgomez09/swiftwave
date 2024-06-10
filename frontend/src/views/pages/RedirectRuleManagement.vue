<script setup>
import FilledButton from '@/views/components/FilledButton.vue'
import ModalDialog from '@/views/components/ModalDialog.vue'
import PageBar from '@/views/components/PageBar.vue'
import CreateDomainModal from '@/views/partials/CreateDomainModal.vue'
import { useMutation, useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { computed, reactive, ref } from 'vue'
import { useToast } from 'vue-toastification'

const toast = useToast()
const isModalOpen = ref(false)
const openModal = () => {
  isModalOpen.value = true
}
const closeModal = () => {
  isModalOpen.value = false
}

// Create redirect rule
const newRedirectRuleDetails = reactive({
  domainId: 0,
  protocol: 'http',
  redirectURL: ''
})

const {
  mutate: createRedirectRule,
  loading: isRedirectRuleCreating,
  onDone: onRedirectRuleCreateSuccess,
  onError: onRedirectRuleCreateFail
} = useMutation(
  gql`
    mutation ($input: RedirectRuleInput!) {
      createRedirectRule(input: $input) {
        id
      }
    }
  `,
  {
    variables: {
      input: newRedirectRuleDetails
    }
  }
)

onRedirectRuleCreateSuccess(() => {
  closeModal()
  newRedirectRuleDetails.name = ''
  refetchRedirectRules()
})

onRedirectRuleCreateFail((err) => {
  toast.error(err.message)
})

// Fetch domains from the server
const { result: domainListResult, refetch: refetchDomains } = useQuery(
  gql`
    query {
      domains {
        id
        name
      }
    }
  `,
  null,
  {
    pollInterval: 10000
  }
)
const domains = computed(() => domainListResult.value?.domains ?? [])

// Delete redirect rule
const {
  mutate: deleteRedirectRule,
  onDone: onRedirectDeleteSuccess,
  onError: onRedirectRuleDeleteFail
} = useMutation(
  gql`
    mutation ($id: Uint!) {
      deleteRedirectRule(id: $id)
    }
  `,
  {
    variables: {
      id: 0
    }
  }
)

const deleteRedirectRulesWithConfirmation = (redirect_rules) => {
  if (confirm('Are you sure you want to delete this redirect rule?')) {
    deleteRedirectRule({
      id: redirect_rules.id
    })
  }
}

onRedirectDeleteSuccess(() => {
  toast.success('Redirect Rule deleted successfully')
  refetchRedirectRules()
})

onRedirectRuleDeleteFail((err) => {
  toast.error(err.message)
})

// Fetch redirect rules
const {
  result: redirectRulesRaw,
  refetch: refetchRedirectRules,
  loading: isRedirectRulesLoading,
  onError: onRedirectRulesError
} = useQuery(
  gql`
    query {
      redirectRules {
        id
        domain {
          name
        }
        protocol
        redirectURL
        status
      }
    }
  `,
  null,
  {
    pollInterval: 10000
  }
)

const redirectRules = computed(() => redirectRulesRaw.value?.redirectRules ?? [])

onRedirectRulesError((err) => {
  toast.error(err.message)
})

const redirectRuleFrontURL = (redirectRule) => {
  return `${redirectRule.protocol}://${redirectRule.domain.name}`
}

// Create Domain
const createDomainModalRef = ref(null)
const openNewDomainModal = () => {
  if (!createDomainModalRef.value?.openModal) return
  isModalOpen.value = false
  createDomainModalRef.value.openModal()
}

const openRedirectRuleRegistrationModal = () => {
  isModalOpen.value = true
}
</script>

<template>
  <CreateDomainModal ref="createDomainModalRef" :callback-on-create="refetchDomains"
    :callback-on-pop="openRedirectRuleRegistrationModal" />
  <section class="mx-auto w-full max-w-7xl">
    <!-- Modal for create redirect rules -->
    <ModalDialog :close-modal="closeModal" :is-open="isModalOpen">
      <template v-slot:header>Create Redirect Rule</template>
      <template v-slot:body>
        Enter the details of the new redirect rule.
        <form @submit.prevent="createRedirectRule">
          <!-- Domains -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Select Domain and Protocol</span>
              </div>
              <div class="flex gap-2">
                <select class="select select-bordered" v-model="newRedirectRuleDetails.protocol">
                  <option value="http">HTTP</option>
                  <option value="https">HTTPS</option>
                </select>
                <select id="domain" v-model="newRedirectRuleDetails.domainId" class="select select-bordered w-full">
                  <option value="0">Select a domain</option>
                  <option v-for="domain in domains" :key="domain.id" :value="domain.id">{{ domain.name }}</option>
                </select>
              </div>
            </label>
          </div>
          <p class="mt-2 flex items-center text-sm">
            Need to create a domain?
            <a @click="openNewDomainModal" class="text-primary ml-1.5 cursor-pointer font-bold">Register New
              Domain</a>
          </p>

          <!--  Redirected URL   -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Redirected URL</span>
              </div>
              <input id="name" v-model="newRedirectRuleDetails.redirectURL" autocomplete="off"
                class="input input-bordered w-full" name="name" placeholder="Name of redirected URL" type="text" />
            </label>
          </div>
        </form>
      </template>
      <template v-slot:footer>
        <FilledButton :click="createRedirectRule" :loading="isRedirectRuleCreating" type="primary">
          Register
        </FilledButton>
      </template>
    </ModalDialog>

    <!-- Top Page bar   -->
    <PageBar>
      <template v-slot:title>Redirect Rules</template>
      <template v-slot:subtitle>Manage Redirect Rules</template>
      <template v-slot:buttons>
        <FilledButton :click="openModal" type="primary">
          <font-awesome-icon icon="fa-solid fa-plus" class="mr-2" />
          Add New
        </FilledButton>
        <FilledButton type="ghost" :click="refetchRedirectRules">
          <font-awesome-icon icon="fa-solid fa-arrows-rotate" :class="{
            'animate-spin ': isRedirectRulesLoading
          }" />&nbsp;&nbsp; Refresh List
        </FilledButton>
      </template>
    </PageBar>

    <div
      class="grid grid-col gap-2 lg:gap-8 auto-cols-max grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 mt-4">

      <div class="indicator" v-for="redirectRule in redirectRules" :key="redirectRule.id">
        <span class="indicator-item badge" :class="{
          'badge-warning': redirectRule.status === 'pending',
          'badge-success': redirectRule.status === 'applied',
          'badge-error': redirectRule.status === 'failed' || redirectRule.status === 'deleting'
        }"></span>

        <div class="card bg-base-300 shadow">
          <div class="card-body p-4">
            <div class="card-title items-center justify-between">
              <div class="flex items-center gap-2">
                <h2>{{ redirectRuleFrontURL(redirectRule) }}</h2>
                <div class="badge mr-1" v-if="volume.type === 'local'">Local</div>
                <div class="mr-1" v-if="volume.type === 'nfs'">&nbsp;&nbsp;NFS&nbsp;&nbsp;</div>
                <div class="mr-1" v-if="volume.type === 'cifs'">&nbsp;CIFS&nbsp;</div>
              </div>

              <font-awesome-icon class="cursor-pointer" icon="fa-solid fa-eye" size="sm" @click="showDetails" />
            </div>

            <div class="flex flex-col gap-2">
              <!-- Status -->
              <div class="flex items-center gap-2">
                <span>Status:</span>
                <span v-if="redirectRule.status === 'pending'" class="font-semibold">Pending</span>
                <span v-else-if="redirectRule.status === 'applied'" class="font-semibold">Applied</span>
                <span v-else-if="redirectRule.status === 'failed'" class="font-semibold">Failed</span>
                <span v-else-if="redirectRule.status === 'deleting'" class="font-semibold">Deleting</span>
              </div>

              <!-- Target -->
              <div class="flex items-center gap-2">
                <span>Target:</span>
                <span class="font-semibold">{{ redirectRule.redirectURL }}</span>
              </div>
            </div>

            <div class="card-actions justify-end">
              <FilledButton :click="() => deleteRedirectRulesWithConfirmation(redirectRule)" type="error">
                Delete
              </FilledButton>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
