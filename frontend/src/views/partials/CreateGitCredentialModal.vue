<script setup>
import ModalDialog from '@/views/components/ModalDialog.vue'
import { reactive, ref } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { useToast } from 'vue-toastification'
import FilledButton from '@/views/components/FilledButton.vue'
import { preventSpaceInput } from '@/vendor/utils.js'

const props = defineProps({
  callbackOnCreate: {
    type: Function,
    required: false,
    default: () => { }
  },
  callbackOnPop: {
    type: Function,
    required: false,
    default: () => { }
  }
})

const toast = useToast()

const isModalOpen = ref(false)
const resetForm = () => {
  newGitCredential.name = ''
  newGitCredential.type = ''
  newGitCredential.username = ''
  newGitCredential.password = ''
  newGitCredential.sshPrivateKey = ''
}
const openModal = () => {
  resetForm()
  isModalOpen.value = true
}
const closeModal = () => {
  isModalOpen.value = false
  props.callbackOnPop()
  resetForm()
}

// New Git Credential form state
const newGitCredential = reactive({
  name: '',
  type: '',
  username: '',
  password: '',
  sshPrivateKey: ''
})

const {
  mutate: createGitCredential,
  loading: isGitCredentialCreating,
  onDone: onGitCredentialCreateSuccess,
  onError: onGitCredentialCreateFail
} = useMutation(
  gql`
    mutation ($input: GitCredentialInput!) {
      createGitCredential(input: $input) {
        id
        name
      }
    }
  `,
  {
    variables: {
      input: newGitCredential
    }
  }
)

onGitCredentialCreateSuccess(() => {
  closeModal()
  toast.success('Git Credential created successfully')
  props.callbackOnCreate()
})

onGitCredentialCreateFail((err) => {
  toast.error(err.message)
})

defineExpose({
  openModal,
  closeModal
})
</script>

<template>
  <teleport to="body">
    <ModalDialog :close-modal="closeModal" :is-open="isModalOpen" width="lg">
      <template v-slot:header>Add Git Credential</template>
      <template v-slot:body>
        Enter the necessary information for configuring the new Git Credential.
        <form @submit.prevent="">
          <!--  Name Field   -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Name (Provide a name to identify the credential)</span>
              </div>
              <input id="name" v-model="newGitCredential.name" autocomplete="off" class="select select-bordered w-full"
                name="name" placeholder="Name" type="text" />
            </label>
          </div>
          <!-- Type Field -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Authentication Type</span>
              </div>
              <select id="git_credential" v-model="newGitCredential.type" class="select select-bordered w-full">
                <option selected value="">No Credential</option>
                <option value="http">HTTP</option>
                <option value="ssh">SSH</option>
              </select>
            </label>
          </div>
          <!-- Username Field -->
          <div class="mt-4" v-if="newGitCredential.type === 'http'">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Git Username</span>
              </div>
              <input id="username" v-model="newGitCredential.username" @keydown="preventSpaceInput" autocomplete="off"
                class="input input-bordered w-full" name="username" placeholder="Git Username" type="text" />
            </label>
          </div>
          <!-- Password Field -->
          <div class="mt-4" v-if="newGitCredential.type === 'http'">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Git Password / Auth Token</span>
              </div>
              <input id="password" v-model="newGitCredential.password" autocomplete="off"
                class="input input-bordered w-full" name="password" placeholder="Git Password" type="text" />
            </label>
          </div>
          <!-- Private Key Field -->
          <div class="mt-4" v-if="newGitCredential.type === 'ssh'">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">ED25519 Private Key (In OpenSSH Format)</span>
              </div>
              <textarea id="ssh_private_key" v-model="newGitCredential.sshPrivateKey"
                class="textarea textarea-bordered w-full" placeholder="ECDSA Private Key (In OpenSSH Format)"
                type="text" rows="5" />
              <p class="mt-1 text-sm text-gray-500">
                <b class="text-danger-500">NOTE:</b> Leave the input blank if you like to auto-generate the private key
              </p>
            </label>
          </div>
        </form>
      </template>
      <template v-slot:footer>
        <FilledButton :click="createGitCredential" :loading="isGitCredentialCreating" type="primary"
          :disabled="newGitCredential.type === '' || newGitCredential.name === ''">
          Add Now
        </FilledButton>
      </template>
    </ModalDialog>
  </teleport>
</template>

<style scoped></style>
