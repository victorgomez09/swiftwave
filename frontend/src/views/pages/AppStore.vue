<script setup>
import { onMounted, ref, shallowRef, watch } from 'vue'
import ModalDialog from '@/views/components/ModalDialog.vue'
import OutlinedButton from '@/views/components/OutlinedButton.vue'
import DotLoader from '@/views/components/DotLoader.vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const storeEndpoints = shallowRef(['https://raw.githubusercontent.com/swiftwave-org/app-store/main/store.json'])
const apps = ref([])
const appsShown = ref([])
const searchText = ref('')
const isOptionsModalOpen = ref(false)
const selectedApp = ref({})
const isLoading = ref(false)

watch(apps, () => {
  if (!apps.value.length) return
  searchApps()
})

onMounted(() => {
  fetchApps()
})

const closeModal = () => {
  isOptionsModalOpen.value = false
}

const openModal = () => {
  isOptionsModalOpen.value = true
}

function fetchApps() {
  isLoading.value = true
  // for each endpoint, fetch apps
  storeEndpoints.value.forEach((endpoint) => {
    fetch(endpoint)
      .then((response) => response.json())
      .then((data) => {
        apps.value = apps.value.concat(data)
        isLoading.value = false
      })
  })
}

function searchApps() {
  if (!searchText.value) {
    appsShown.value = apps.value
    return
  }
  // split search text by space and search for each word
  const searchWords = searchText.value.split(' ')
  appsShown.value = apps.value.filter((app) => {
    return searchWords.every((word) => {
      return (
        app.title.toLowerCase().includes(word.toLowerCase()) ||
        app.description.toLowerCase().includes(word.toLowerCase())
      )
    })
  })
}

const chooseApp = (app) => {
  if (!app) return
  if (app.stacks.length === 1) {
    openStackFileForInstall(app.stacks[0])
    return
  }
  selectedApp.value = app
  openModal()
}

const openStackFileForInstall = (stack) => {
  router.push({
    name: 'Install from App Store',
    query: {
      stack: stack.stack
    }
  })
}
</script>

<template>
  <!-- Main -->
  <section class="mx-auto mt-2 flex w-full max-w-7xl flex-col items-center overflow-hidden">
    <!-- Header -->
    <div class="mb-2 flex w-full flex-row justify-center gap-3 p-2">
      <!--   Search Bar   -->
      <label class="input input-bordered flex items-center gap-2">
        <input type="text" class="grow border-none focus:outline-none" v-model="searchText"
          v-debounce:300ms="searchApps" placeholder="Search" />
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" class="w-4 h-4 opacity-70">
          <path fill-rule="evenodd"
            d="M9.965 11.026a5 5 0 1 1 1.06-1.06l2.755 2.754a.75.75 0 1 1-1.06 1.06l-2.755-2.754ZM10.5 7a3.5 3.5 0 1 1-7 0 3.5 3.5 0 0 1 7 0Z"
            clip-rule="evenodd" />
        </svg>
      </label>
      <!--   Settings   -->
      <!--  TODO: Will be added soon   -->
      <!--      <FilledButton class="my-1" type="secondary" rounded>-->
      <!--        <span class="flex items-center px-1.5"-->
      <!--          ><font-awesome-icon icon="fa-solid fa-gear" class="mr-2 text-base" /> Settings</span-->
      <!--        >-->
      <!--      </FilledButton>-->
    </div>
    <!-- If loading   -->
    <div v-if="isLoading" class="flex h-[500px] w-full items-center justify-center">
      <DotLoader />
    </div>
    <!--    No app available -->
    <div v-else-if="appsShown.length === 0" class="flex h-[500px] w-full flex-col items-center justify-center">
      <p class="text-5xl">🤔</p>
      <p class="ml-4 mt-10 text-xl font-medium">No apps found</p>
      <p class="mt-3">
        If you think the app should be here, Raise a request in
        <a href="https://github.com/swiftwave-org/app-store" target="_blank"
          class="text-primary font-semibold">Swiftwave App Store</a>.
      </p>
    </div>
    <!-- Apps List (If available) -->
    <div v-else class="scrollbox grid grid-cols-4 gap-4 overflow-auto pr-2">
      <!-- Component  -->
      <div @click="() => chooseApp(app)" v-for="app in appsShown" :key="app.id"
        class="card bg-base-300 cursor-pointer h-40 hover:shadow">
        <div class="card-body p-4">
          <!-- Header -->
          <div class="flex flex-row gap-3 border-b pb-2">
            <div class="h-12 w-12 rounded-md p-1.5">
              <img :src="app.logo" class="h-full w-full" :alt="app.title" />
            </div>
            <div>
              <p class="text-base font-semibold text-gray-800">{{ app.title }}</p>
              <p class="text-sm">{{ app.category }}</p>
            </div>
          </div>
          <!-- Description Body -->
          <p class="truncate">
            {{ app.description }}
          </p>
        </div>
      </div>
    </div>
  </section>
  <!-- Modal to show options -->
  <ModalDialog :close-modal="closeModal" :is-open="isOptionsModalOpen">
    <template v-slot:header>Install {{ selectedApp.title }}</template>
    <template v-slot:body>
      <p>Choose the preferred version -</p>
      <div class="mt-6 flex flex-col gap-2">
        <OutlinedButton :click="() => openStackFileForInstall(stack)" class="w-full" type="primary"
          v-for="stack in selectedApp.stacks" :key="stack.id">
          {{ stack.title }}
        </OutlinedButton>
      </div>
    </template>
  </ModalDialog>
</template>

<style scoped>
.scrollbox::-webkit-scrollbar {
  width: 9px;
}

.scrollbox::-webkit-scrollbar-track {
  @apply rounded-full bg-gray-200;
}

.scrollbox::-webkit-scrollbar-thumb {
  @apply bg-primary rounded-full;
}
</style>
