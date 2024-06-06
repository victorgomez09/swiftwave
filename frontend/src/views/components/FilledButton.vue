<script setup>
import { computed } from 'vue'

const props = defineProps({
  loading: {
    default: false
  },
  disabled: {
    default: false
  },
  click: {
    type: Function,
    default: () => { }
  },
  type: {
    type: String,
    default: 'primary',
    validator(value) {
      return ['primary', 'success', 'warning', 'danger', 'secondary', 'info', 'ghost'].includes(value)
    }
  },
  slim: {
    type: Boolean,
    default: false
  },
  rounded: {
    type: Boolean,
    default: false
  },
  size: {
    type: String,
    default: '',
    validator(value) {
      return ['', 'sm', 'lg'].includes(value)
    }
  },
})

const isDisabled = computed(() => {
  return props.disabled || props.loading
})
</script>

<template>
  <button :class="{
    'btn-primary': type === 'primary',
    'btn-secondary': type === 'secondary',
    'btn-success': type === 'success',
    'btn-warning': type === 'warning',
    'btn-danger': type === 'danger',
    'btn-info': type === 'info',
    'btn-ghost': type === 'ghost',
    'rounded-full': rounded,
    'rounded-md': !rounded,
    'btn-sm': size === 'sm',
    'btn-lg': size === 'lg'
  }" :disabled="isDisabled" class="btn" type="button" @click.stop="click">
    <span v-if="loading" class="loading loading-spinner"></span>
    <!-- text -->
    <slot></slot>
  </button>
</template>

<style scoped></style>
