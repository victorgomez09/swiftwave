<script setup>
import PageBar from '@/views/components/PageBar.vue'
import ApplicationGroup from '@/views/partials/ApplicationGroup.vue'
import { useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'

const router = useRouter()
const toast = useToast()

const deployNewApplication = () => {
  router.push('/deploy/application')
}

const {
  result: applicationsResult,
  refetch: refetchApplications,
  loading: isApplicationsLoading,
  onError: onApplicationsError
} = useQuery(
  gql`
    query {
      applications {
        id
        name
        deploymentMode
        replicas
        isSleeping
        realtimeInfo {
          InfoFound
          DesiredReplicas
          RunningReplicas
          DeploymentMode
        }
        latestDeployment {
          status
          upstreamType
          gitProvider
          createdAt
        }
        group
      }
    }
  `,
  null,
  {
    pollInterval: 30000
  }
)

onApplicationsError((err) => {
  toast.error(err.message)
})

const applications = computed(() => applicationsResult.value?.applications ?? [])
const applicationGroupWise = computed(() => {
  let applications = applicationsResult.value?.applications ?? []
  let groupedApplications = {}
  applications.forEach((application) => {
    if (!groupedApplications[encodeURI(application.group)]) {
      groupedApplications[encodeURI(application.group)] = []
    }
    groupedApplications[encodeURI(application.group)].push(application)
  })
  return groupedApplications
})
</script>

<template>
  <section class="mx-auto w-full max-w-7xl">
    <!-- Top Page bar   -->
    <PageBar>
      <template v-slot:title>Deployed Applications</template>
      <template v-slot:subtitle>Take control of your deployed applications</template>
      <template v-slot:buttons>
        <q-btn color="primary" @click="deployNewApplication()">
          <font-awesome-icon icon="fa-solid fa-hammer" class="mr-2" />
          Deploy App
        </q-btn>
        <q-btn flat color="primary" @click="refetchApplications()">
          <font-awesome-icon icon="fa-solid fa-arrows-rotate" :class="{
            'animate-spin ': isApplicationsLoading
          }" />&nbsp;&nbsp; Refresh List
        </q-btn>
      </template>
    </PageBar>

    <!-- Table -->
    <p v-if="applications.length === 0">
      No deployed applications found.<br />
      Click on the "Deploy New" button to deploy a new application.
    </p>
    <p v-if="isApplicationsLoading"> Loading deployed applications...</p>

    <ApplicationGroup v-for="(applications, group, index) in applicationGroupWise" :key="group" :group="group"
      :group-index="index" :applications="applications" />
  </section>
</template>
