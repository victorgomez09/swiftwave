<script setup>
defineProps({
  title: {
    type: String,
    required: true
  },
  subTitle: {
    type: String,
    default: ''
  },
  prefixText: {
    type: String,
    default: ''
  },
  type: {
    type: String,
    default: 'success',
    validator: (value) => {
      return ['primary', 'secondary', 'success', 'danger', 'warning'].includes(value)
    }
  },
  showBody: {
    type: Boolean,
    default: true
  }
})
</script>

<template>
  <div class="my-3 flex items-center gap-4">
    <div class="flex h-12 w-12 items-center justify-center rounded-box border text-xl font-bold text-white shadow-sm"
      :class="{
        'bg-primary border-primary': type === 'primary',
        'border-secondary bg-secondary': type === 'secondary',
        'border-success bg-success': type === 'success',
        'border-error bg-error': type === 'error',
        'border-warning bg-warning': type === 'warning'
      }">
      {{ prefixText }}
    </div>
    <div class="flex flex-col">
      <div class="text-base font-medium text-base-content">{{ title }}</div>
      <div class="text-sm font-thin text-base-content">{{ subTitle }}</div>
    </div>
  </div>
  <div class="ml-16 mt-4" v-if="showBody">
    <slot></slot>
  </div>
</template>

<style scoped></style>
