<script setup>
import FilledButton from '@/views/components/FilledButton.vue'
import ModalDialog from '@/views/components/ModalDialog.vue'
import { ref } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { useToast } from 'vue-toastification'
import { preventSpaceInput } from '@/vendor/utils.js'
import { useRouter } from 'vue-router'

const props = defineProps({
  serverId: {
    type: Number,
    required: true
  },
  serverSshPort: {
    type: Number,
    required: true
  }
})

const toast = useToast()
const router = useRouter()
const isModalOpen = ref(false)
const ipChanged = ref(false)
const newServerSSHPort = ref('')

const openModal = () => {
  newServerSSHPort.value = props.serverSshPort
  ipChanged.value = false
  isModalOpen.value = true
}
const closeModal = () => {
  isModalOpen.value = false
}

const {
  mutate: changeServerSSHPortRaw,
  loading: isRequestRunning,
  onDone: onChangeServerSSHPortSuccess,
  onError: onChangeServerSSHPortFail
} = useMutation(gql`
  mutation ($id: Uint!, $port: Int!) {
    changeServerSSHPort(id: $id, port: $port)
  }
`)

onChangeServerSSHPortSuccess((val) => {
  if (val.data.changeServerSSHPort) {
    ipChanged.value = true
    startCountDown()
    closeModal()
  }
})

onChangeServerSSHPortFail((err) => {
  toast.error(err.message)
})

const changeServerSSHPort = () => {
  changeServerSSHPortRaw({
    id: props.serverId,
    port: newServerSSHPort.value
  })
}

const timeCount = ref(5)

const startCountDown = () => {
  const interval = setInterval(() => {
    timeCount.value--
    if (timeCount.value === 0) {
      clearInterval(interval)
      ipChanged.value = false
      router.push({ name: 'Maintenance', query: { redirect: router.currentRoute.value.path } })
    }
  }, 1000)
}

defineExpose({
  openModal,
  closeModal
})
</script>

<template>
  <teleport to="body">
    <ModalDialog :is-open="ipChanged" non-cancelable>
      <template v-slot:header>
        <span>🚀 Server SSH Port Changed</span>
      </template>
      <template v-slot:body>
        <p class="mb-2">SSH Port changed successfully! Swiftwave needs to restart.</p>
        <p>
          Redirecting to Maintenance Page in <b>{{ timeCount }}</b> seconds
        </p>
      </template>
    </ModalDialog>
    <ModalDialog :close-modal="closeModal" :is-open="isModalOpen">
      <template v-slot:header>Change Server SSH Port</template>
      <template v-slot:body>
        Note: Changing the server SSH Port address will restart the swiftwave service automatically.
        <form @submit.prevent="">
          <!--  IP Field   -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">SSH Port</span>
              </div>
              <input id="ssh_port" v-model="newServerSSHPort" @keydown="preventSpaceInput" autocomplete="off"
                class="input input-bordered w-full" type="text" />
            </label>
          </div>
        </form>
      </template>
      <template v-slot:footer>
        <FilledButton :click="changeServerSSHPort" :loading="isRequestRunning" type="primary"
          :disabled="newServerSSHPort === serverSshPort || newServerSSHPort === ''">Change Server SSH Port
        </FilledButton>
      </template>
    </ModalDialog>
  </teleport>
</template>

<style scoped></style>
