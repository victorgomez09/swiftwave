<script setup>
import PageBar from '@/views/components/PageBar.vue'
import FilledButton from '@/views/components/FilledButton.vue'
import { useToast } from 'vue-toastification'
import { computed, ref } from 'vue'
import gql from 'graphql-tag'
import { useMutation, useQuery } from '@vue/apollo-composable'
import GitCredentialCard from '@/views/partials/GitCredentialCard.vue'
import CreateGitCredentialModal from '@/views/partials/CreateGitCredentialModal.vue'

const toast = useToast()

// Create Git Credential
const createGitCredentialModalRef = ref(null)
const openCreateGitCredentialModal = computed(() => createGitCredentialModalRef.value?.openModal ?? (() => { }))

// Delete Git Credential mutation
const {
  mutate: deleteGitCredential,
  onError: onGitCredentialDeleteError,
  onDone: onGitCredentialDeleteSuccess
} = useMutation(
  gql`
    mutation ($id: Uint!) {
      deleteGitCredential(id: $id)
    }
  `,
  {
    variables: {
      id: ''
    }
  }
)

onGitCredentialDeleteError((err) => {
  toast.error(err.message)
})

onGitCredentialDeleteSuccess(() => {
  refetchGitCredentialList()
  toast.success('Git Credential deleted successfully')
})

const deleteGitCredentialWithConfirmation = (gitCredential) => {
  if (
    confirm(
      `Are you sure you want to delete Git Credential ${gitCredential.name}?\nExisting deployments using this Git Credential can't use this credential anymore.`
    )
  ) {
    deleteGitCredential({ id: gitCredential.id })
  }
}

// List Git Credentials query
const {
  result: gitCredentialList,
  refetch: refetchGitCredentialList,
  loading: isGitCredentialListLoading,
  onError: onGitCredentialListError
} = useQuery(
  gql`
    query {
      gitCredentials {
        id
        name
        type
      }
    }
  `,
  null,
  {
    pollInterval: 30000
  }
)
const gitCredentials = computed(() => gitCredentialList.value?.gitCredentials ?? [])

onGitCredentialListError((err) => {
  toast.error(err.message)
})
</script>

<template>
  <section class="mx-auto w-full max-w-7xl">
    <!-- Modal for create -->
    <CreateGitCredentialModal ref="createGitCredentialModalRef" :callback-on-create="refetchGitCredentialList" />
    <!-- Top Page bar   -->
    <PageBar>
      <template v-slot:title>Git Credentials</template>
      <template v-slot:subtitle> Manage Git Credentials and usage in deployments</template>
      <template v-slot:buttons>
        <FilledButton :click="openCreateGitCredentialModal" type="primary">
          <font-awesome-icon icon="fa-solid fa-plus" class="mr-2" />
          Add New
        </FilledButton>
        <FilledButton type="ghost" :click="refetchGitCredentialList">
          <font-awesome-icon icon="fa-solid fa-arrows-rotate" :class="{
            'animate-spin ': isGitCredentialListLoading
          }" />&nbsp;&nbsp; Refresh List
        </FilledButton>
      </template>
    </PageBar>

    <div
      class="grid grid-col gap-2 lg:gap-8 auto-cols-max grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 mt-4">
      <GitCredentialCard v-for="gitCredential in gitCredentials" v-bind:key="gitCredential.id"
        :delete-git-credential="deleteGitCredentialWithConfirmation" :git-credential="gitCredential"
        :on-update-git-credential="refetchGitCredentialList" />
    </div>
  </section>
</template>
