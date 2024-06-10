<script setup>
import FilledButton from '@/views/components/FilledButton.vue'
import PageBar from '@/views/components/PageBar.vue'
import CreateImageRegistryCredentialModal from '@/views/partials/CreateImageRegistryCredentialModal.vue'
import ImageRegistryCredentialCard from '@/views/partials/ImageRegistryCredentialCard.vue'
import { useMutation, useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { computed, ref } from 'vue'
import { useToast } from 'vue-toastification'

const toast = useToast()

// Create Image Registry Credential
const createImageRegistryCredentialModalRef = ref(null)
const openCreateImageRegistryCredentialModal = computed(
  () => createImageRegistryCredentialModalRef.value?.openModal ?? (() => { })
)

// Delete Image Registry Credential mutation
const {
  mutate: deleteImageRegistryCredential,
  onError: onImageRegistryCredentialDeleteError,
  onDone: onImageRegistryCredentialDeleteSuccess
} = useMutation(
  gql`
    mutation ($id: Uint!) {
      deleteImageRegistryCredential(id: $id)
    }
  `,
  {
    variables: {
      id: ''
    }
  }
)

onImageRegistryCredentialDeleteError((err) => {
  toast.error(err.message)
})

onImageRegistryCredentialDeleteSuccess(() => {
  refetchImageRegistryCredentialList()
  toast.success('Image Registry Credential deleted successfully')
})

const deleteImageRegistryCredentialWithConfirmation = (imageRegistryCredential) => {
  if (
    confirm(
      `Are you sure you want to delete Image Registry Credential ?\nExisting deployments using this Image Registry Credential can't use this credential anymore.`
    )
  ) {
    deleteImageRegistryCredential({ id: imageRegistryCredential.id })
  }
}

// List Image Registry Credentials query
const {
  result: imageRegistryCredentialList,
  refetch: refetchImageRegistryCredentialList,
  loading: isImageRegistryCredentialListLoading,
  onError: onImageRegistryCredentialListError
} = useQuery(
  gql`
    query {
      imageRegistryCredentials {
        id
        url
        username
        password
      }
    }
  `,
  null,
  {
    pollInterval: 30000
  }
)
const imageRegistryCredentials = computed(() => imageRegistryCredentialList.value?.imageRegistryCredentials ?? [])

onImageRegistryCredentialListError((err) => {
  toast.error(err.message)
})
</script>

<template>
  <section class="mx-auto w-full max-w-7xl">
    <!-- Modal for create -->
    <CreateImageRegistryCredentialModal ref="createImageRegistryCredentialModalRef"
      :callback-on-create="refetchImageRegistryCredentialList" />

    <!-- Top Page bar   -->
    <PageBar>
      <template v-slot:title>Image Registry Credentials</template>
      <template v-slot:subtitle> Manage Image Registry Credentials and usage in deployments</template>
      <template v-slot:buttons>
        <FilledButton :click="openCreateImageRegistryCredentialModal" type="primary">
          <font-awesome-icon icon="fa-solid fa-plus" class="mr-2" />
          Add New
        </FilledButton>
        <FilledButton type="ghost" :click="refetchImageRegistryCredentialList">
          <font-awesome-icon icon="fa-solid fa-arrows-rotate" :class="{
            'animate-spin ': isImageRegistryCredentialListLoading
          }" />&nbsp;&nbsp; Refresh List
        </FilledButton>
      </template>
    </PageBar>

    <div
      class="grid grid-col gap-2 lg:gap-8 auto-cols-max grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 mt-4">
      <ImageRegistryCredentialCard v-for="imageRegistryCredential in imageRegistryCredentials"
        v-bind:key="imageRegistryCredential.id"
        :delete-image-registry-credential="deleteImageRegistryCredentialWithConfirmation"
        :image-registry-credential="imageRegistryCredential"
        :on-update-image-registry-credential="refetchImageRegistryCredentialList" />
    </div>
  </section>
</template>
