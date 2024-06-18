<script setup>
import FilledButton from '../components/FilledButton.vue'

defineProps({
  user: {
    type: Object,
    required: true
  },
  deleteUser: {
    type: Function,
    required: true
  }
})
</script>

<template>
  <div class="card bg-base-300 shadow">
    <div class="card-body p-4">
      <h3 class="card-title">{{ user.username }}</h3>

      <div class="flex flex-col gap-2">
        <div class="flex gap-2">
          <span>Status:</span>
          <span class="font-semibold">Active</span>
        </div>

        <div class="flex gap-2">
          <span>Role:</span>
          <span class="font-semibold">Administrator</span>
        </div>
      </div>

      <div class="card-actions justify-end">
        <FilledButton :click="() => deleteUser(user)" type="error" size="sm"> Delete </FilledButton>
      </div>
    </div>
  </div>
</template>
