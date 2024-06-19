<script setup>
import FilledButton from '@/views/components/FilledButton.vue'
import PageBar from '@/views/components/PageBar.vue'
import CreateServerModal from '@/views/partials/CreateServerModal.vue'
import ServerCard from '@/views/partials/ServerCard.vue'
import { useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { computed, ref } from 'vue'
import { useToast } from 'vue-toastification'

const toast = useToast()
const createServerModal = ref(null)

const {
  result: serversResult,
  loading: isServersLoading,
  refetch: refetchServers,
  onError: onServersError
} = useQuery(
  gql`
    query {
      servers {
        id
        ip
        hostname
        user
        ssh_port
        swarmMode
        swarmNodeStatus
        maintenanceMode
        scheduleDeployments
        dockerUnixSocketPath
        proxyEnabled
        proxyType
        status
      }
    }
  `,
  null,
  {
    pollInterval: 10000
  }
)

onServersError((err) => {
  toast.error(err.message)
})

const servers = computed(() => serversResult.value?.servers ?? [])
const openCreateServerModal = () => {
  if (createServerModal.value) createServerModal.value.openModal()
}
</script>

<template>
  <!-- Modal to create server  -->
  <CreateServerModal :callback-on-create="refetchServers" ref="createServerModal" />
  <section class="mx-auto w-full max-w-7xl">
    <!-- Top Page bar   -->
    <PageBar>
      <template v-slot:title>Registered Servers</template>
      <template v-slot:subtitle>Take control of your servers</template>
      <template v-slot:buttons>
        <FilledButton type="primary" :click="openCreateServerModal">
          <font-awesome-icon icon="fa-solid fa-plus" />
          &nbsp;&nbsp; Add Server
        </FilledButton>
        <FilledButton type="ghost" :click="refetchServers">
          <font-awesome-icon icon="fa-solid fa-arrows-rotate" :class="{
            'animate-spin ': isServersLoading
          }" />&nbsp;&nbsp; Refresh List
        </FilledButton>
      </template>
    </PageBar>

    <div
      class="grid grid-col gap-2 lg:gap-8 auto-cols-max grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 mt-4">
      <ServerCard v-for="server in servers" :key="server.id" :server="server" :refetch-servers="refetchServers" />
    </div>

  </section>
</template>
