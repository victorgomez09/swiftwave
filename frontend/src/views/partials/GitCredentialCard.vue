<script setup>
import FilledButton from '@/views/components/FilledButton.vue'
import { ref } from 'vue'
import UpdateGitCredentialModal from '@/views/partials/UpdateGitCredentialModal.vue'
import GitCredentialDetailsModal from '@/views/partials/GitCredentialDetailsModal.vue'

defineProps({
  gitCredential: {
    type: Object,
    required: true
  },
  deleteGitCredential: {
    type: Function,
    required: true
  },
  onUpdateGitCredential: {
    type: Function,
    required: false,
    default: () => { }
  }
})

const updateGitCredentialModalRef = ref(null)
const gitCredentialDetailsModalRef = ref(null)

const openEditModal = () => {
  if (updateGitCredentialModalRef.value) updateGitCredentialModalRef.value.openModal()
}
const openDetailsModal = () => {
  if (gitCredentialDetailsModalRef.value) gitCredentialDetailsModalRef.value.openModal()
}
</script>

<template>
  <!-- Update Git Credential modal -->
  <UpdateGitCredentialModal :git-credential-id="gitCredential.id" ref="updateGitCredentialModalRef"
    :callback-on-pop="onUpdateGitCredential" />
  <!-- Git Credential details modal -->
  <GitCredentialDetailsModal :git-credential-id="gitCredential.id" ref="gitCredentialDetailsModalRef" />

  <div class="indicator">
    <span class="indicator-item badge badge-sm" :class="{
      'badge-success': gitCredential.type === 'http',
      'badge-warning': gitCredential.type === 'ssh',
      'badge-error': gitCredential.type !== 'http' && gitCredential.type !== 'ssh'
    }"></span>

    <div class="card bg-base-300 shadow">
      <div class="card-body p-4">
        <h3 class="card-title">{{ gitCredential.name }}</h3>

        <div class="flex flex-col gap-2">
          <!-- Status -->
          <div class="flex items-center gap-2">
            <span>Authentication:</span>
            <span v-if="gitCredential.type === 'http'" class="font-semibold">HTTP Authentication</span>
            <span v-else-if="gitCredential.type === 'ssh'" class="font-semibold">SSH Authentication</span>
            <span v-else class="font-semibold">Unknown</span>
          </div>
        </div>

        <div class="card-actions justify-end">
          <FilledButton slim :onclick="openDetailsModal" size="sm">Show Details</FilledButton>
          <FilledButton type="secondary" slim :onclick="openEditModal" size="sm">Edit Details</FilledButton>
          <FilledButton :click="() => deleteGitCredential(gitCredential)" type="error" size="sm">Delete</FilledButton>
        </div>
      </div>
    </div>
  </div>
</template>
