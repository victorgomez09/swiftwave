<script setup>
import Drawer from '@/views/components/Drawer.vue'
import { ref, watch } from 'vue'
import { useLazyQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { useToast } from 'vue-toastification'
import Badge from '@/views/components/Badge.vue'
import { round } from 'lodash'
import FilledButton from '@/views/components/FilledButton.vue'
import DotLoader from '@/views/components/DotLoader.vue'

const props = defineProps({
  isDrawerOpen: Boolean,
  closeDrawer: Function,
  persistentVolumeId: Number,
  persistentVolumeName: String
})

const toast = useToast()
const persistentVolumeNameRef = ref(props.persistentVolumeName)
const backups = ref([])

watch(
  () => props.persistentVolumeName,
  (newValue) => {
    persistentVolumeNameRef.value = newValue
  }
)
watch(
  () => props.isDrawerOpen,
  (newValue) => {
    if (newValue) {
      backups.value = []
      fetchPersistentVolumeBackups()
    }
  }
)

// Fetch Persistent Volume Backups
const {
  load: loadPersistentVolumeBackups,
  refetch: refetchPersistentVolumeBackups,
  loading: isPersistentVolumeBackupsLoading,
  onError: onPersistentVolumeBackupsError,
  variables: persistentVolumeBackupsVariables,
  onResult: onPersistentVolumeBackupsResult
} = useLazyQuery(
  gql`
    query ($id: Uint!) {
      persistentVolume(id: $id) {
        backups {
          id
          type
          status
          sizeMb
          createdAt
          completedAt
        }
      }
    }
  `,
  null,
  {
    fetchPolicy: 'no-cache',
    nextFetchPolicy: 'no-cache'
  }
)

const fetchPersistentVolumeBackups = async () => {
  persistentVolumeBackupsVariables.value = { id: props.persistentVolumeId }
  if (!(await loadPersistentVolumeBackups())) {
    await refetchPersistentVolumeBackups()
  }
}

onPersistentVolumeBackupsResult((result) => {
  backups.value = result.data?.persistentVolume?.backups || []
})

onPersistentVolumeBackupsError((err) => {
  toast.error(err.message)
})

const downloadBackup = (backup_id) => {
  const downloadRoute = window.location.origin + import.meta.env.BASE_URL + `pv-backup-download/${backup_id}`
  window.open(downloadRoute, '_blank')
}
</script>

<template>
  <Drawer :close-drawer="closeDrawer" :is-open="isDrawerOpen" :width-size="3">
    <template v-slot:title>Manage your backups</template>
    <template v-slot:subtitle>Manage backups for volume {{ persistentVolumeNameRef }}</template>
    <template v-slot:body>
      <FilledButton slim class="mb-2 w-full" type="secondary" :click="fetchPersistentVolumeBackups">
        <font-awesome-icon icon="fa-solid fa-rotate-right" class="mr-2" />
        Refresh List
      </FilledButton>
      <div class="mb-2 text-sm text-gray-500">{{ backups.length }} backups found for this volume</div>
      <div class="my-4 flex justify-center" v-if="isPersistentVolumeBackupsLoading">
        <DotLoader />
      </div>
      <div class="flex flex-col space-y-2" v-if="!isPersistentVolumeBackupsLoading">
        <div v-for="backup in backups" :key="backup.id" class="card border">
          <div class="card-body p-4">
            <p>
            <div v-if="backup.status === 'pending'" class="badge badge-warning">Pending</div>
            <div v-if="backup.status === 'failed'" class="badge badge-error">Failed</div>
            <div v-if="backup.status === 'success'" class="badge badge-success">Success</div>
            <span class="ml-2 mr-1"><b>Initiated at</b> {{ new Date(backup.createdAt).toLocaleString() }}</span>
            </p>

            <div v-if="backup.status === 'success'" class="mt-1">
              <span class="mr-1">Backed up in</span>
              <Badge type="warning">{{ backup.type }}</Badge>
              <span class="ml-2 mr-1">of</span>
              <Badge type="warning">{{ round(backup.sizeMb, 4) }} MB</Badge>
            </div>
            <div v-if="backup.status === 'success'" class="mt-1">
              Completed at {{ new Date(backup.completedAt).toLocaleString() }}
            </div>
            <FilledButton v-if="backup.status === 'success'" class="mt-2 w-full" slim
              :click="() => downloadBackup(backup.id)">
              <font-awesome-icon icon="fa-solid fa-circle-down" class="mr-2" />
              Download Backup
            </FilledButton>
          </div>
        </div>
      </div>
    </template>
  </Drawer>
</template>

<style scoped></style>
