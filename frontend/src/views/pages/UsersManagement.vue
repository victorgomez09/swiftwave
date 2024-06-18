<script setup>
import { useMutation, useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { computed, reactive, ref } from 'vue'
import ModalDialog from '@/views/components/ModalDialog.vue'
import FilledButton from '@/views/components/FilledButton.vue'
import PageBar from '@/views/components/PageBar.vue'
import { useToast } from 'vue-toastification'
import UserCard from '@/views/partials/UserCard.vue'
import { preventSpaceInput } from '@/vendor/utils.js'

const toast = useToast()
const isModalOpen = ref(false)
const openModal = () => {
  isModalOpen.value = true
}
const closeModal = () => {
  isModalOpen.value = false
}

// New user form state
const newUser = reactive({
  username: '',
  password: ''
})

const {
  mutate: createUser,
  loading: isUserCreating,
  onDone: onUserCreateSuccess,
  onError: onUserCreateFail
} = useMutation(
  gql`
    mutation ($input: UserInput!) {
      createUser(input: $input) {
        id
        username
      }
    }
  `,
  {
    variables: {
      input: newUser
    }
  }
)

onUserCreateSuccess(() => {
  closeModal()
  refetchUserList()
  newUser.username = ''
  newUser.password = ''
  toast.success('User created successfully')
})

onUserCreateFail((err) => {
  toast.error(err.message)
})

// Delete user mutation
const {
  mutate: deleteUser,
  onDone: onUserDeleteSuccess,
  onError: onUserDeleteFail
} = useMutation(gql`
  mutation ($id: Uint!) {
    deleteUser(id: $id)
  }
`)

const deleteUserWithConfirmation = (user) => {
  if (confirm(`Are you sure you want to delete user ${user.username}?`)) {
    deleteUser({ id: user.id })
  }
}

onUserDeleteSuccess(() => {
  refetchUserList()
  toast.success('User deleted successfully')
})

onUserDeleteFail((err) => {
  toast.error(err.message)
})

// User list query
const {
  result: userListResult,
  refetch: refetchUserList,
  onError: onUserListFetchFailed
} = useQuery(
  gql`
    query {
      users {
        id
        username
      }
    }
  `,
  null,
  {
    pollInterval: 10000
  }
)
const users = computed(() => userListResult.value?.users)

onUserListFetchFailed((err) => {
  toast.error(err.message)
})
</script>

<template>
  <section class="mx-auto w-full max-w-7xl">
    <!-- Modal for new user -->
    <ModalDialog :close-modal="closeModal" :is-open="isModalOpen">
      <template v-slot:header>Create new user</template>
      <template v-slot:body>
        Enter the username and password for the new user.
        <form @submit.prevent="createUser">
          <!-- Username Field -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Username</span>
              </div>
              <input
                id="username"
                v-model="newUser.username"
                @keydown="preventSpaceInput"
                autocomplete="off"
                class="input input-bordered w-full"
                name="username"
                placeholder="Username"
                type="text" />
            </label>
          </div>
          <!-- Password Field -->
          <div class="mt-4">
            <label class="form-control w-full">
              <div class="label">
                <span class="label-text">Password</span>
              </div>
              <input
                id="password"
                v-model="newUser.password"
                autocomplete="new-password"
                class="input input-bordered w-full"
                name="password"
                placeholder="Password"
                type="password" />
            </label>
          </div>
        </form>
      </template>
      <template v-slot:footer>
        <FilledButton :click="createUser" :loading="isUserCreating" type="primary">Create </FilledButton>
      </template>
    </ModalDialog>

    <!-- Top Page bar   -->
    <PageBar>
      <template v-slot:title>Users</template>
      <template v-slot:subtitle>
        Registered users can access the Vira Deploy dashboard, allowing them to perform all actions and access all
        features
      </template>
      <template v-slot:buttons>
        <FilledButton :click="openModal" type="primary">Create User</FilledButton>
      </template>
    </PageBar>

    <div
      class="grid-col mt-4 grid auto-cols-max grid-cols-1 gap-2 md:grid-cols-2 lg:grid-cols-3 lg:gap-8 xl:grid-cols-4">
      <UserCard v-for="user in users" v-bind:key="user.id" :delete-user="deleteUserWithConfirmation" :user="user" />
    </div>
  </section>
</template>
