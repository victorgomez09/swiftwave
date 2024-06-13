<script setup>
import FilledButton from '@/views/components/FilledButton.vue'
import ModalDialog from '@/views/components/ModalDialog.vue'
import { onMounted, reactive } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { useToast } from 'vue-toastification'

const props = defineProps({
  isModalOpen: {
    type: Boolean,
    required: true
  },
  closeModal: {
    type: Function,
    required: true
  }
})

const toast = useToast()
const passwordDetails = reactive({
  newPassword: '',
  oldPassword: ''
})
const resetPasswordDetails = () => {
  passwordDetails.newPassword = ''
  passwordDetails.oldPassword = ''
}

onMounted(() => resetPasswordDetails())

const {
  mutate: changePassword,
  loading: isChangingPassword,
  onError: onPasswordChangeFail,
  onDone: onPasswordChangeSuccess
} = useMutation(
  gql`
    mutation ($input: PasswordUpdateInput) {
      changePassword(input: $input)
    }
  `,
  {
    fetchPolicy: 'no-cache',
    variables: {
      input: passwordDetails
    }
  }
)

onPasswordChangeSuccess((e) => {
  if (e.data.changePassword === true) {
    toast.success('Password changed successfully')
    props.closeModal()
  } else {
    toast.error('Password change failed')
  }
  resetPasswordDetails()
})

onPasswordChangeFail((error) => {
  toast.error(error.message)
  resetPasswordDetails()
})
</script>

<template>
  <Teleport to="body">
    <ModalDialog :close-modal="closeModal" :is-open="isModalOpen">
      <template v-slot:header> Change Password </template>
      <template v-slot:body>
        Enter details to change password.
        <form @submit.prevent="changePassword">
          <!-- Username Field -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Old Password</span>
              </div>
              <input id="oldPassword" v-model="passwordDetails.oldPassword" autocomplete="off"
                class="input input-bordered w-full" placeholder="Old Password" type="password" />
            </label>
          </div>
          <!-- Password Field -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">New Password</span>
              </div>
              <input id="newPassword" v-model="passwordDetails.newPassword" autocomplete="off"
                class="input input-bordered w-full" placeholder="New Password" type="password" />
            </label>
          </div>
        </form>
      </template>
      <template v-slot:footer>
        <FilledButton :click="changePassword" :loading="isChangingPassword" type="primary">Change Password
        </FilledButton>
      </template>
    </ModalDialog>
  </Teleport>
</template>
