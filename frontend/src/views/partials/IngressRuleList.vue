<script setup>
import { useMutation, useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { computed } from 'vue'
import { useToast } from 'vue-toastification'
import FilledButton from '../components/FilledButton.vue'

const props = defineProps({
  applicationId: {
    type: String,
    required: false,
    default: ''
  }
})
const toast = useToast()

// Delete ingress rule
const {
  mutate: deleteIngressRule,
  onDone: onIngressDeleteSuccess,
  onError: onIngressRuleDeleteFail
} = useMutation(
  gql`
    mutation ($id: Uint!) {
      deleteIngressRule(id: $id)
    }
  `,
  {
    variables: {
      id: 0
    }
  }
)

const deleteIngressRulesWithConfirmation = (ingress_rule) => {
  if (confirm('Are you sure you want to delete this ingress rule ?')) {
    deleteIngressRule({
      id: ingress_rule.id
    })
  }
}

onIngressDeleteSuccess(() => {
  toast.success('Ingress Rule will be deleted shortly\nThis can take upto 5 minutes to reflect in the system')
  refetchIngressRules()
})

onIngressRuleDeleteFail((err) => {
  toast.error(err.message)
})

// Queries
const fetchAllIngressRulesQuery = gql`
  query {
    ingressRules {
      id
      status
      protocol
      domain {
        name
      }
      port
      targetType
      externalService
      application {
        name
      }
      targetPort
    }
  }
`
const applicationSpecificIngressRulesQuery = gql`
  query ($id: String!) {
    application(id: $id) {
      ingressRules {
        id
        status
        protocol
        domain {
          name
        }
        port
        application {
          name
        }
        targetType
        targetPort
      }
    }
  }
`
// Fetch ingress rules
const {
  result: ingressRulesRaw,
  refetch: refetchIngressRules,
  loading: isIngressRulesLoading,
  onError: onIngressRulesError
} = useQuery(
  props.applicationId ? applicationSpecificIngressRulesQuery : fetchAllIngressRulesQuery,
  {
    id: props.applicationId
  },
  {
    pollInterval: 10000
  }
)

const ingressRules = computed(() =>
  props.applicationId
    ? ingressRulesRaw.value?.application?.ingressRules ?? []
    : ingressRulesRaw.value?.ingressRules ?? []
)

onIngressRulesError((err) => {
  toast.error(err.message)
})

defineExpose({
  refetchIngressRules,
  isIngressRulesLoading
})
</script>

<template>
  <div class="grid grid-col gap-2 lg:gap-8 auto-cols-max grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 mt-4">
    <div class="indicator" v-for="ingressRule in ingressRules" :key="ingressRule.id">
      <span class="indicator-item badge badge-sm" :class="{
        'badge-warning': ingressRule.status === 'pending',
        'badge-success': ingressRule.status === 'applied',
        'badge-error': ingressRule.status === 'failed' || ingressRule.status === 'deleting',
      }"></span>

      <div class="card bg-base-300 shadow">
        <div class="card-body p-4">
          <div class="card-title items-center justify-between">
            <h3 v-if="ingressRule.protocol === 'http' || ingressRule.protocol === 'https'">
              {{ ingressRule.protocol }}://{{ ingressRule.domain.name }}:{{ ingressRule.port }}
            </h3>
            <h3 v-else-if="ingressRule.protocol === 'tcp'">tcp://&lt;server-ip&gt;:{{
              ingressRule.port }}</h3>
            <h3 v-else-if="ingressRule.protocol === 'udp'">udp://&lt;server-ip&gt;:{{
              ingressRule.port }}</h3>
            <h3 v-else><i>Unknown</i></h3>
          </div>

          <div class="flex flex-col gap-2">
            <!-- Status -->
            <div class="flex items-center gap-2">
              <span>Status:</span>
              <span v-if="ingressRule.status === 'pending'" class="font-semibold">Pending</span>
              <span v-else-if="ingressRule.status === 'applied'" class="font-semibold">Applied</span>
              <span v-else-if="ingressRule.status === 'failed'" class="font-semibold">Failed</span>
              <span v-else-if="ingressRule.status === 'deleting'" class="font-semibold">Deleting</span>
            </div>

            <!-- Target -->
            <div class="flex items-center gap-2">
              <span>Target:</span>
              <span v-if="ingressRule.targetType === 'application'" class="font-semibold">
                {{ ingressRule.application.name }}
              </span>
              <span v-else>{{ ingressRule.externalService }}</span>
            </div>
          </div>

          <!-- Target type -->
          <div class="flex items-center gap-2">
            <span>Target type:</span>
            <span v-if="ingressRule.targetType === 'externalService'" class="font-semibold">External Service</span>
            <span v-else-if="ingressRule.targetType === 'application'" class="font-semibold">Application</span>
          </div>

          <!-- Target port -->
          <div class="flex items-center gap-2">
            <span>Target port:</span>
            <span class="font-semibold">{{ ingressRule.targetPort }}</span>
          </div>

          <div class="card-actions justify-end">
            <FilledButton :click="() => deleteIngressRulesWithConfirmation(ingressRule)" type="error" size="sm">
              Delete
            </FilledButton>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
