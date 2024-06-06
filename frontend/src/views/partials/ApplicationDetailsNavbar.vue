<script setup></script>

<template>
  <div class="menu bg-base-200 w-56 rounded-box h-full">
    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Deployments',
        params: { id: $route.params.id }
      }">
        Deployments
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Resource Stats',
        params: { id: $route.params.id }
      }">
        Analytics
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Runtime Logs',
        params: { id: $route.params.id }
      }">
        Runtime Log
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Ingress Rules',
        params: { id: $route.params.id }
      }">
        Ingress Rules
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Persistent Volumes',
        params: { id: $route.params.id }
      }">
        Persistent Volume
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Environment Variables',
        params: { id: $route.params.id }
      }">
        Environment Variables
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Update Source',
        params: { id: $route.params.id }
      }">
        Update App Source
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Deployment Config',
        params: { id: $route.params.id }
      }">
        Deploy Config
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Webhook CI',
        params: { id: $route.params.id }
      }">
        Configure Webhook
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element" :to="{
        name: 'Application Details Manage',
        params: { id: $route.params.id }
      }">
        Actions
      </RouterLink>
    </li>

    <li>
      <RouterLink class="nav-element !text-red-500" :to="{
        name: 'Application Details Danger Zone',
        params: { id: $route.params.id }
      }">
        Destroy
      </RouterLink>
    </li>
  </div>
</template>

<style scoped>
.router-link-exact-active {
  @apply active;
}
</style>
