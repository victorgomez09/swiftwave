<script setup>
import router from '@/router/index.js';
import FilledButton from '@/views/components/FilledButton.vue';
import moment from 'moment';
import { computed } from 'vue';

const props = defineProps({
  application: {
    class: Object,
    required: true
  },
  isVisible: {
    class: Boolean,
    default: true
  }
})

const createdAtFormatted = computed(() => {
  return moment(props.application.latestDeployment.createdAt).format('DD/MM/YYYY HH:mm')
})
const viewApplicationDetails = () => {
  router.push(`/application/${props.application.id}/deployments`)
}
</script>

<template>
  <div class="indicator w-full">
    <span class="indicator-item badge" :class="{
      'badge-warning': application.latestDeployment.status === 'pending',
      'badge-warning': application.latestDeployment.status === 'deployPending',
      'badge-warning': application.latestDeployment.status === 'deploying',
      'badge-success': application.latestDeployment.status === 'live',
      'badge-warning': application.latestDeployment.status === 'stopped',
      'badge-error': application.latestDeployment.status === 'failed'
    }"></span>

    <div class="card bg-base-300 w-full hover:shadow" v-show="isVisible">
      <div class="card-body p-4">
        <h2 class="card-title">
          <font-awesome-icon v-if="application.latestDeployment.upstreamType === 'git'" icon="fa-solid fa-code-branch"
            class="mx-2" />
          <font-awesome-icon v-else-if="application.latestDeployment.upstreamType === 'image'" class="mx-2"
            icon="fa-brands fa-docker" />
          <font-awesome-icon v-else-if="application.latestDeployment.upstreamType === 'sourceCode'" class="mx-2"
            icon="fa-solid fa-upload" />
          {{ application.name }}
        </h2>

        <div class="flex flex-col gap-2">
          <!-- Status -->
          <div class="flex items-center gap-2">
            <span>Status:</span>
            <div v-if="application.latestDeployment.status === 'pending'" class="font-semibold">Pending
            </div>
            <div v-else-if="application.latestDeployment.status === 'deployPending'" class="font-semibold">
              Deploy Pending
            </div>
            <div v-else-if="application.latestDeployment.status === 'deploying'" class="font-semibold">
              Deploying
            </div>
            <div v-else-if="application.latestDeployment.status === 'live'" class="font-semibold">
              Live
            </div>
            <div v-else-if="application.latestDeployment.status === 'stopped'" class="font-semibold">
              Stopped
            </div>
            <div v-else-if="application.latestDeployment.status === 'failed'" class="font-semibold">
              Failed
            </div>
          </div>

          <!-- Replicas -->
          <div class="flex items-center gap-2">
            <span>Replicas:</span>
            <div v-if="!application.isSleeping && application.realtimeInfo.InfoFound" align="center"
              class="font-semibold">
              <div v-if="application.deploymentMode === 'replicated'" class="text-sm text-base-content">
                {{ application.realtimeInfo.RunningReplicas }} / {{ application.realtimeInfo.DesiredReplicas }}
              </div>
              <div v-else-if="application.deploymentMode === 'global'" class="text-sm text-base-content">Global</div>
              <div v-else class="text-sm text-base-content">----</div>
            </div>
            <div v-else-if="application.isSleeping" align="center">
              <div class="text-warning">
                <div class="flex flex-row items-center gap-1.5">
                  <font-awesome-icon icon="fa-solid fa-bed" />
                  Sleeping
                </div>
              </div>
            </div>
            <div v-else align="center">
              <div class="text-sm text-base-content">----</div>
            </div>
          </div>

          <!-- Source -->
          <div class="flex items-center gap-2">
            <span>Source:</span>
            <span v-if="application.latestDeployment.upstreamType === 'git'"
              class="text-sm capitalize text-base-content font-semibold">
              {{ application.latestDeployment.gitProvider }}
            </span>
            <span v-else-if="application.latestDeployment.upstreamType === 'image'"
              class="text-sm text-base-content  font-semibold">
              Docker Image
            </span>
            <span v-else-if="application.latestDeployment.upstreamType === 'sourceCode'"
              class="text-sm text-base-content font-semibold">
              Source Code
            </span>
            <span v-else class="text-sm text-base-content font-semibold">N/A</span>
          </div>

          <!-- Dates -->
          <div class="flex items-center gap-2">
            <span>Created:</span>
            <span class="text-sm text-base-content font-semibold"> {{ createdAtFormatted }} </span>
          </div>
        </div>

        <div class="card-actions justify-end">
          <FilledButton :click="viewApplicationDetails" slim class="primary" size="sm">View Details</FilledButton>
        </div>
      </div>
    </div>
  </div>
</template>
