<script setup>
import FilledButton from '@/views/components/FilledButton.vue'
import UpdateImageRegistryCredentialModal from '@/views/partials/UpdateImageRegistryCredentialModal.vue'
import { ref } from 'vue'

defineProps({
  imageRegistryCredential: {
    type: Object,
    required: true
  },
  deleteImageRegistryCredential: {
    type: Function,
    required: true
  },
  onUpdateImageRegistryCredential: {
    type: Function,
    required: false,
    default: () => { }
  }
})

const updateImageRegistryCredential = ref(null)
const openEditModal = () => {
  if (updateImageRegistryCredential.value) updateImageRegistryCredential.value.openModal()
}
</script>

<template>
  <!-- Update Image Registry Credential modal -->
  <UpdateImageRegistryCredentialModal :image-registry-credential-id="imageRegistryCredential.id"
    ref="updateImageRegistryCredential" :callback-on-pop="onUpdateImageRegistryCredential" />

  <div class="card bg-base-300 shadow">
    <div class="card-body p-4">
      <h3 class="card-title">{{ imageRegistryCredential.url }}</h3>

      <div class="flex flex-col gap-2">
        <!-- Username -->
        <div class="flex items-center gap-2">
          <span>Username:</span>
          <span>{{ imageRegistryCredential.username }}</span>
        </div>

        <!-- Password -->
        <div class="flex items-center gap-2">
          <span>Password:</span>
          <span>{{ imageRegistryCredential.password }}</span>
        </div>
      </div>

      <div class="card-actions justify-end">
        <FilledButton type="secondary" slim :click="openEditModal" size="sm">Edit Details</FilledButton>

        <FilledButton :click="() => deleteImageRegistryCredential(imageRegistryCredential)" type="error" size="sm">
          Delete
        </FilledButton>
      </div>
    </div>
  </div>
</template>
