<script setup>
import TableRow from '@/views/components/Table/TableRow.vue'
import { useLazyQuery, useMutation } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { computed, ref } from 'vue'
import FilledButton from '@/views/components/FilledButton.vue'
import { round } from 'lodash'
import { useToast } from 'vue-toastification'
import ModalDialog from '@/views/components/ModalDialog.vue'

const props = defineProps({
  volume: {
    type: Object,
    required: true
  },
  deletePersistentVolumeWithConfirmation: {
    type: Function,
    required: true
  },
  showBackups: {
    type: Function,
    required: true
  },
  showRestores: {
    type: Function,
    required: true
  },
  restoreNow: {
    type: Function,
    required: true
  },
  showDetails: {
    type: Function,
    required: true
  }
})

const toast = useToast()

const sizeFetched = ref(false)

const {
  result: persistentVolumeSizeResult,
  load: loadPersistentVolumeSize,
  refetch: refetchPersistentVolumeSize,
  loading: isPersistentVolumeSizeLoading,
  onResult: onPersistentVolumeSizeResult
} = useLazyQuery(
  gql`
    query ($id: Uint!) {
      persistentVolumeSizeMb(id: $id)
    }
  `,
  {
    id: props.volume.id
  },
  {
    fetchPolicy: 'no-cache',
    nextFetchPolicy: 'no-cache'
  }
)

onPersistentVolumeSizeResult(() => {
  sizeFetched.value = true
})

const fetchPersistentVolumeSize = async () => {
  if (!(await loadPersistentVolumeSize())) {
    await refetchPersistentVolumeSize()
  }
}
const persistentVolumeSize = computed(() => round(persistentVolumeSizeResult.value?.persistentVolumeSizeMb ?? 0, 2))

// Create backup
const backupType = ref('local')
const isBackupCreateModalOpen = ref(false)
const {
  mutate: createPersistentVolumeBackup,
  loading: isBackupCreating,
  onDone: onBackupCreateSuccess,
  onError: onBackupCreateFail
} = useMutation(gql`
  mutation ($input: PersistentVolumeBackupInput!) {
    backupPersistentVolume(input: $input) {
      id
    }
  }
`)

onBackupCreateSuccess(() => {
  toast.success('Backup will be created shortly ! Check the status in the backups list section.')
})

onBackupCreateFail((err) => {
  toast.error(err.message)
})

const closeBackupCreateModal = () => {
  isBackupCreateModalOpen.value = false
}

const openBackupCreateModal = () => {
  console.log('opening')
  isBackupCreateModalOpen.value = true
}

const createBackup = () => {
  closeBackupCreateModal()

  createPersistentVolumeBackup({
    input: {
      persistentVolumeId: props.volume.id,
      type: backupType.value
    }
  })
}
</script>

<template>
  <Teleport to="body">
    <ModalDialog :close-modal="closeBackupCreateModal" :is-open="isBackupCreateModalOpen">
      <template v-slot:header>Backup Volume</template>
      <template v-slot:body>
        <form @submit.prevent="createPersistentVolumeBackup">
          <!-- Domains -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Choose Backup Type</span>
              </div>
              <select class="select select-bordered w-full" v-model="backupType">
                <option value="local">Local</option>
                <option value="s3">S3</option>
              </select>
            </label>
          </div>
        </form>
      </template>
      <template v-slot:footer>
        <FilledButton :click="createBackup" :loading="isBackupCreating" type="primary">
          Create Backup
        </FilledButton>
      </template>
    </ModalDialog>
  </Teleport>

  <div class="card bg-base-200 hover:shadow" :key="volume.id">
    <div class="card-body p-4">
      <div class="card-title items-center justify-between">
        <div class="flex items-center gap-2">
          <h2>{{ volume.name }}</h2>
          <div class="badge mr-1" v-if="volume.type === 'local'">Local</div>
          <div class="mr-1" v-if="volume.type === 'nfs'">&nbsp;&nbsp;NFS&nbsp;&nbsp;</div>
          <div class="mr-1" v-if="volume.type === 'cifs'">&nbsp;CIFS&nbsp;</div>
        </div>

        <font-awesome-icon class="cursor-pointer" icon="fa-solid fa-eye" size="sm" @click="showDetails" />
      </div>

      <div class="flex flex-col gap-2">
        <!-- SIZE -->
        <div class="flex items-center gap-2">
          <span class="font-semibold">Size:</span>
          <span class="flex items-center">
            <p v-if="sizeFetched">{{ persistentVolumeSize }} MB</p>
            <p v-else>Fetch size to show data</p>
          </span>
        </div>
      </div>

      <!-- ACTIONS -->
      <div class="card-actions justify-end items-center">
        <div>
          <FilledButton :click="fetchPersistentVolumeSize" slim type="primary" :loading="isPersistentVolumeSizeLoading"
            size="sm">
            {{ sizeFetched ? 'Refresh Size' : 'Fetch Size' }}
          </FilledButton>
          <!-- <FilledButton slim type="secondary" : size="sm"> Show Details</FilledButton> -->
        </div>

        <div class="dropdown">
          <FilledButton slim size="sm" type="" role="button" tabindex="0">
            <font-awesome-icon icon="fa-solid fa-ellipsis-vertical" size="sm" />
          </FilledButton>
          <ul tabindex="0" class="dropdown-content absolute z-[100] menu p-2 shadow bg-base-100 rounded-box w-52">
            <li @click="openBackupCreateModal">
              <a>
                <span v-show="isBackupCreating" class="loading loading-spinner loading-xs"></span>
                <font-awesome-icon icon="fa-solid fa-plus" class="mr-2" size="sm" />
                Create Backup
              </a>
            </li>
            <li @click="showBackups">
              <a>
                <font-awesome-icon icon="fa-solid fa-list-check" class="mr-2" size="sm" />
                Show Backups
              </a>
            </li>
            <li @click="restoreNow">
              <a>
                <font-awesome-icon icon="fa-solid fa-upload" class="mr-2" />
                Restore Now
              </a>
            </li>
            <li @click="showRestores">
              <a>
                <font-awesome-icon icon="fa-solid fa-list-check" class="mr-2" />
                Restore History
              </a>
            </li>
            <div class="divider p-0 m-0"></div>
            <li @click="deletePersistentVolumeWithConfirmation(volume)">
              <a>
                Delete
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
