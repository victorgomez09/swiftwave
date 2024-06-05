import { createApp, h, markRaw, provide } from 'vue'
import { createPinia } from 'pinia'
import vueDebounce from 'vue-debounce'

import { DefaultApolloClient } from '@vue/apollo-composable'

import { ApolloClient, ApolloLink, createHttpLink, InMemoryCache, split } from '@apollo/client/core'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faAngleDown,
  faAngleUp,
  faArrowDown,
  faArrowRight,
  faArrowRotateRight,
  faArrowsRotate,
  faArrowUpRightFromSquare,
  faBed,
  faBook,
  faBox,
  faBug,
  faCalendarDays,
  faCaretDown,
  faChartColumn,
  faChartSimple,
  faCheck,
  faChevronDown,
  faChevronRight,
  faCircleCheck,
  faCircleDown,
  faCircleNotch,
  faCircleStop,
  faCircleXmark,
  faClipboard,
  faCloud,
  faCodeBranch,
  faCopy,
  faCubesStacked,
  faDiagramProject,
  faEllipsisVertical,
  faEnvelope,
  faEye,
  faEyeSlash,
  faFile,
  faFileLines,
  faFileWaveform,
  faFingerprint,
  faGear,
  faGlobe,
  faHammer,
  faHandshakeAngle,
  faHardDrive,
  faKey,
  faLayerGroup,
  faLink,
  faListCheck,
  faListUl,
  faLocationArrow,
  faMagnifyingGlass,
  faNetworkWired,
  faPeopleGroup,
  faPersonDigging,
  faPlay,
  faPlus,
  faPowerOff,
  faRightFromBracket,
  faRocket,
  faRotateRight,
  faRoute,
  faServer,
  faSignal,
  faSkullCrossbones,
  faSpinner,
  faStop,
  faStore,
  faTerminal,
  faTrash,
  faTriangleExclamation,
  faUpload,
  faUsers,
  faUserTie,
  faVault,
  faWrench,
  faXmark
} from '@fortawesome/free-solid-svg-icons'

import { useAuthStore } from '@/store/auth.js'
import Toast from 'vue-toastification'
import VueApexCharts from 'vue3-apexcharts'

import App from './App.vue'
import router from './router'
import './assets/css/base.css'
import 'vue-toastification/dist/index.css'
import { faDocker } from '@fortawesome/free-brands-svg-icons/faDocker'
import { faGit, faGithub } from '@fortawesome/free-brands-svg-icons'
import { createClient } from 'graphql-ws'
import { GraphQLWsLink } from '@apollo/client/link/subscriptions' // <-- This one uses graphql-ws
import { getMainDefinition } from '@apollo/client/utilities'
import { getGraphQlHttpBaseUrl, getGraphQlWsBaseUrl } from '@/vendor/utils.js'

// add icons to library
library.add(
  faHammer,
  faBox,
  faHardDrive,
  faCodeBranch,
  faCloud,
  faLink,
  faNetworkWired,
  faLocationArrow,
  faUsers,
  faRightFromBracket,
  faChevronDown,
  faChevronRight,
  faVault,
  faArrowRight,
  faUpload,
  faDocker,
  faGit,
  faCircleCheck,
  faCircleXmark,
  faTrash,
  faArrowDown,
  faCalendarDays,
  faFingerprint,
  faGear,
  faTriangleExclamation,
  faSkullCrossbones,
  faArrowUpRightFromSquare,
  faKey,
  faCopy,
  faEye,
  faEyeSlash,
  faPlus,
  faListCheck,
  faXmark,
  faCircleDown,
  faRotateRight,
  faFile,
  faPlay,
  faCircleStop,
  faCubesStacked,
  faStore,
  faMagnifyingGlass,
  faGlobe,
  faRocket,
  faWrench,
  faRoute,
  faServer,
  faListUl,
  faEllipsisVertical,
  faChartColumn,
  faTerminal,
  faBook,
  faArrowsRotate,
  faClipboard,
  faArrowRotateRight,
  faFileLines,
  faDiagramProject,
  faAngleDown,
  faAngleUp,
  faChartSimple,
  faSpinner,
  faStop,
  faFileWaveform,
  faBed,
  faSignal,
  faCaretDown,
  faCheck,
  faPowerOff,
  faUserTie,
  faGithub,
  faBug,
  faHandshakeAngle,
  faEnvelope,
  faPeopleGroup,
  faLayerGroup,
  faCircleNotch,
  faPersonDigging
)

// Environment variables
const GRAPHQL_HTTP_BASE_URL = getGraphQlHttpBaseUrl()
const GRAPHQL_WS_BASE_URL = getGraphQlWsBaseUrl()

// Setup apollo client
// create apollo link
const httpLink = createHttpLink({
  uri: `${GRAPHQL_HTTP_BASE_URL}/graphql`
})

const wsLink = new GraphQLWsLink(
  createClient({
    url: `${GRAPHQL_WS_BASE_URL}/graphql`,
    connectionParams: () => {
      const authStore = useAuthStore()
      return {
        authorization: authStore.FetchBearerToken()
      }
    }
  })
)

// create auth middleware
const apolloAuthMiddleware = new ApolloLink((operation, forward) => {
  // add the authorization to the headers
  const authStore = useAuthStore()
  operation.setContext({
    headers: {
      authorization: authStore.FetchBearerToken()
    }
  })
  return forward(operation)
})

const link = split(
  ({ query }) => {
    const definition = getMainDefinition(query)
    return definition.kind === 'OperationDefinition' && definition.operation === 'subscription'
  },
  wsLink,
  apolloAuthMiddleware.concat(httpLink)
)

// create apollo client
const apolloClient = new ApolloClient({
  link: link,
  defaultOptions: {
    query: {
      fetchPolicy: 'no-cache'
    },
    mutate: {
      fetchPolicy: 'no-cache'
    },
    watchQuery: {
      fetchPolicy: 'no-cache'
    }
  },
  cache: new InMemoryCache()
})

// create app
const app = createApp({
  setup() {
    // provide apollo client
    provide(DefaultApolloClient, apolloClient)
  },
  render: () => h(App)
})
app.component('font-awesome-icon', FontAwesomeIcon)
const pinia = createPinia()
pinia.use(({ store }) => {
  store.$router = markRaw(router)
})
app.use(router)
app.use(pinia)
app.use(Toast, {
  position: 'top-right',
  timeout: 5000,
  closeOnClick: true,
  pauseOnFocusLoss: true,
  pauseOnHover: true,
  draggable: true,
  draggablePercent: 0.6,
  showCloseButtonOnHover: false,
  hideProgressBar: false,
  closeButton: 'button',
  icon: true,
  rtl: false
})
app.use(VueApexCharts)
app.directive('debounce', vueDebounce({ lock: true }))
app.mount('#app')

// Protect routes
router.beforeEach(async (to) => {
  const authStore = useAuthStore()
  if ((to.name === 'Setup' && parseInt(to.query?.update ?? 0) === 0) || to.name === 'Maintenance') {
    return
  }
  if (!authStore.IsLoggedIn && to.name !== 'Login') {
    return { name: 'Login', query: { redirect: to.path } }
  }
  if (authStore.IsLoggedIn && to.name === 'Login') {
    return { name: 'Applications' }
  }
})
