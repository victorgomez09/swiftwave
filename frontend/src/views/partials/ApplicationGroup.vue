<script setup>
import router from '@/router/index.js';
import { getRandomBackgroundAndBorderColourClass } from '@/vendor/utils.js';
import moment from 'moment';
import { computed, onMounted, ref } from 'vue';

const props = defineProps({
  groupIndex: {
    type: Number,
    default: null
  },
  group: {
    type: String,
    default: ''
  },
  applications: {
    type: Array,
    default: () => []
  }
})

const isExpanded = ref(false)
const groupReplicasPercentage = computed(() => {
  let liveApps = (props.applications || []).filter(
    (item) => item.latestDeployment && item.latestDeployment.status === 'live'
  ).length
  return (liveApps / props.applications.length) * 100
})

const placeGroupDiv = () => {
  if (!isApplicationListVisible.value) return
  if (props.applications.length === 0) return
  if (props.group === '') return
  const firstTableRowId = `group_${props.group}_${props.applications[0].id}`
  const firstTableRowElement = document.getElementById(firstTableRowId)
  const lastTableRowId = `group_${props.group}_${props.applications[props.applications.length - 1].id}`
  const lastTableRowElement = document.getElementById(lastTableRowId)
  const height = lastTableRowElement.getBoundingClientRect().bottom - firstTableRowElement.getBoundingClientRect().top
  const yStart = firstTableRowElement.getBoundingClientRect().top
  const xEnd = window.innerWidth - firstTableRowElement.getBoundingClientRect().x

  const groupElement = document.getElementById(`group_${props.group}_tag`)
  if (!groupElement) return
  const colorClasses = getRandomBackgroundAndBorderColourClass(props.groupIndex)
  groupElement.classList.add(colorClasses[0])
  groupElement.classList.add(colorClasses[1])
  groupElement.style.height = `${height}px`
  groupElement.style.top = `${yStart}px`
  groupElement.style.right = `${xEnd}px`
  groupElement.classList.remove('hidden')
}

const removeGroupDiv = () => {
  const groupElement = document.getElementById(`group_${props.group}_tag`)
  if (!groupElement) return
  groupElement.classList.add('hidden')
}

const initGroupDiv = () => {
  placeGroupDiv()
  window.addEventListener('resize', placeGroupDiv)
}

const isApplicationListVisible = computed(() => {
  return isExpanded.value || props.group === ''
})

const isGroupTagVisible = computed(() => {
  if (props.group === '') return false
  return isExpanded.value
})

const showApplicationList = () => {
  isExpanded.value = true
  setTimeout(() => {
    placeGroupDiv()
  }, 100)
}

const hideApplicationList = () => {
  isExpanded.value = false
  removeGroupDiv()
}

onMounted(() => {
  initGroupDiv()
})

const columns = [
  {
    name: 'name',
    label: 'Name',
    align: 'left',
    field: row => row.name,
    format: val => `${val}`,
    sortable: true
  },
  {
    name: 'status',
    label: 'Status',
    align: 'left',
    field: row => row.latestDeployment.status,
    sortable: true
  },
  {
    name: 'source',
    label: 'Source',
    align: 'left',
    field: row => row.latestDeployment.upstreamType,
    sortable: true
  },
  {
    name: 'replicas',
    label: 'Replicas',
    align: 'left',
    field: row => row.latestDeployment.status,
    sortable: true
  },
  {
    name: 'createdAt',
    label: 'Created At',
    align: 'left',
    field: row => row.latestDeployment.createdAt,
    format: row => moment(row.latestDeployment.createdAt).format('DD/MM/YYYY HH:mm'),
    sortable: true
  },
]

const getBadgeStatusColor = (status) => {
  switch (status) {
    case "pending":
    case "deployPending":
    case "stopped":
      return "yellow";

    case "live":
      return "green";

    default:
      return "red";
  }
}

const createdAtFormatted = (app) => {
  return moment(app.latestDeployment.createdAt).format('DD/MM/YYYY HH:mm')
}

const viewApplicationDetails = (application) => {
  router.push(`/application/${application.id}/deployments`)
}
</script>

<template>
  <div v-if="isGroupTagVisible" :id="`group_${group}_tag`" style="writing-mode: vertical-rl"
    class="absolute hidden rotate-180 cursor-pointer select-none overflow-hidden truncate text-nowrap rounded-r-md !border-2 !border-l-0 px-0.5 py-2 text-center text-sm transition-all hover:text-wrap">
    {{ decodeURI(group) }}
  </div>
  <div v-if="group !== ''" @click.prevent="isExpanded ? hideApplicationList() : showApplicationList()"
    class="cursor-pointer">
    wip
    <!-- <TableRow align="left">
      <div class="text-sm font-medium text-gray-900">
        {{ decodeURI(group) }}
      </div>
    </TableRow>
    <TableRow align="center" class="text-sm text-gray-700"> {{ groupReplicasPercentage }}% live</TableRow>
    <TableRow align="center" flex>
      <UptimeChart :label="`${groupReplicasPercentage}%`" :percentage="groupReplicasPercentage" :hide-label="true"
        :small="true" :hide-hover="true" />
    </TableRow>
    <TableRow align="center" class="text-sm text-gray-700">
      <font-awesome-icon icon="fa-solid fa-layer-group" class="mr-1" />
      {{ applications.length }} applications
    </TableRow>
    <TableRow align="center">-</TableRow>
    <TableRow align="right" flex>
      <FilledButton slim type="primary" v-if="isExpanded" :click="hideApplicationList">Hide Apps</FilledButton>
      <FilledButton slim type="primary" v-else :click="showApplicationList">View Apps</FilledButton>
    </TableRow> -->
  </div>

  <q-table flat bordered :rows="applications" :columns="columns" row-key="name" class="w-full h-full mt-10">
    <template v-slot:body="props">
      <q-tr :props="props">
        <q-td key="name" :props="props" @click="viewApplicationDetails(props.row)">
          <q-btn flat color="primary" :label="props.row.name" />
        </q-td>
        <q-td key="status" :props="props">
          <q-badge :color="getBadgeStatusColor(props.row.latestDeployment.status)">
            {{ props.row.latestDeployment.status }}
          </q-badge>
        </q-td>
        <q-td key="source" :props="props">
          <span v-if="props.row.latestDeployment.upstreamType === 'git'">
            {{ props.row.latestDeployment.gitProvider }}
          </span>
          <span v-else-if="props.row.latestDeployment.upstreamType === 'image'">
            Docker Image
          </span>
          <span v-else-if="props.row.latestDeployment.upstreamType === 'sourceCode'">
            Source Code
          </span>
          <span v-else>N/A</span>

          <font-awesome-icon v-if="props.row.latestDeployment.upstreamType === 'git'" icon="fa-solid fa-code-branch"
            class="mx-2" />
          <font-awesome-icon v-else-if="props.row.latestDeployment.upstreamType === 'image'" class="mx-2"
            icon="fa-brands fa-docker" />
          <font-awesome-icon v-else-if="props.row.latestDeployment.upstreamType === 'sourceCode'" class="mx-2"
            icon="fa-solid fa-upload" />
        </q-td>
        <q-td key="replicas" :props="props">
          <div v-if="!props.row.isSleeping && props.row.realtimeInfo.InfoFound">
            <div v-if="props.row.deploymentMode === 'replicated'">
              {{ props.row.realtimeInfo.RunningReplicas }} / {{ props.row.realtimeInfo.DesiredReplicas }}
            </div>
            <div v-else-if="props.row.deploymentMode === 'global'">Global</div>
            <div v-else>----</div>
          </div>
          <div v-else-if="props.row.isSleeping">
            <div class="text-warning">
              <div class="flex flex-row items-center gap-1.5">
                <font-awesome-icon icon="fa-solid fa-bed" />
                Sleeping
              </div>
            </div>
          </div>
          <div v-else align="center">
            <div>----</div>
          </div>
        </q-td>

        <q-td key="createdAt" :props="props">
          {{ createdAtFormatted(props.row) }}
        </q-td>
      </q-tr>
    </template>
  </q-table>
</template>
