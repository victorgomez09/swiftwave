<script setup>
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'

const router = useRouter()
const toast = useToast()

const applicationId = router.currentRoute.value.params.id

// Fetch ingress rules
const { result: ingressRulesRaw, onError: onDeploymentsError } = useQuery(
  gql`
    query ($id: String!) {
      application(id: $id) {
        deployments {
          id
          status
          upstreamType
          gitProvider
          commitHash
          dockerImage
          repositoryName
          repositoryOwner
          repositoryBranch
          createdAt
        }
      }
    }
  `,
  {
    id: router.currentRoute.value.params.id
  },
  {
    pollInterval: 10000
  }
)

const deployments = computed(() => ingressRulesRaw.value?.application?.deployments ?? [])

onDeploymentsError((err) => {
  toast.error(err.message)
})

const formatdate = (date) => {
  const x = new Date(date)
  return x.toLocaleString()
}
</script>

<template>
  <div class="max-h-0">
    <p class="w-full text-center font-semibold text-gray-900">🛠 Click on a deployment to view more details 🛠</p>
    <div class="pt-2 pr-1 mt-4 flex max-h-[60vh] flex-col gap-2 overflow-y-auto">
      <RouterLink v-for="deployment in deployments" :key="deployment.id" :to="{
        name: 'Application Deployment Details',
        params: {
          id: applicationId,
          deployment_id: deployment.id
        }
      }" class="indicator w-full">
        <span class="indicator-item badge badge-sm" :class="{
          'badge-warning': deployment.status === 'pending',
          'badge-warning': deployment.status === 'deployPending',
          'badge-warning': deployment.status === 'deploying',
          'badge-success': deployment.status === 'live',
          'badge-warning': deployment.status === 'stopped',
          'badge-error': deployment.status === 'failed'
        }"></span>
        <div class="card bg-base-300 w-full hover:shadow">
          <div class="card-body p-4">
            <div class="flex items-center gap-2 font-bold">
              <font-awesome-icon icon="fa-solid fa-fingerprint" />
              <p class="mr-1">{{ deployment.id }}</p>
            </div>
            <div class="flex flex-row items-center gap-2">
              <font-awesome-icon v-if="deployment.upstreamType === 'git'" icon="fa-solid fa-code-branch" />
              <font-awesome-icon v-if="deployment.upstreamType === 'image'" icon="fa-brands fa-docker" />
              <font-awesome-icon v-if="deployment.upstreamType === 'sourceCode'" icon="fa-solid fa-upload" />

              <p v-if="deployment.upstreamType === 'git'">
                {{ deployment.gitProvider }}@{{ deployment.repositoryOwner }}/{{ deployment.repositoryName }}:{{
                  deployment.repositoryBranch
                }}
              </p>
              <p v-if="deployment.upstreamType === 'image'">{{ deployment.dockerImage }}</p>
              <p v-if="deployment.upstreamType === 'sourceCode'">Source-code uploaded manually</p>
            </div>
            <div class="flex items-center gap-2 font-normal">
              <font-awesome-icon icon="fa-solid fa-calendar-days" />
              <p>Deployment triggered on {{ formatdate(deployment.createdAt) }}</p>
            </div>
          </div>
        </div>
      </RouterLink>
    </div>
    <p class="text-secondary mt-4 w-full text-center text-sm">Scroll down to view more deployments(if any)</p>
  </div>
</template>

<style scoped>
.indicator :where(.indicator-item) {
  inset-inline-end: 5px;
  top: 5px;
}
</style>