<script setup>
import { computed, ref, watch } from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

const props = defineProps({
  value: {
    type: String,
    required: true
  },
  options: {
    type: Array,
    required: true
  },
  onChange: {
    type: Function,
    required: false,
    default: () => { }
  }
})

const selectedOption = ref(props.value)
const query = ref('')
const optionsIncludingQuery = computed(() => {
  if (query.value === '') return props.options
  if (props.options.includes(query.value)) return props.options
  return [query.value, ...props.options]
})

const filteredOptions = computed(() =>
  query.value === ''
    ? optionsIncludingQuery.value
    : optionsIncludingQuery.value.filter((op) => {
      return op.toLowerCase().includes(query.value.toLowerCase())
    })
)

watch(selectedOption, (newValue) => {
  props.onChange(newValue)
})
</script>

<template>
  <select class="select select-bordered w-full" v-model="selectedOption">
    <!-- <input type="text" class="input input-bordered w-full" @change="query = $event.target.value"> -->
    <!-- <ComboboxInput @change="query = $event.target.value" placeholder="Start typing to filter..."
        class="w-full border-none focus:ring-0" />
      <ComboboxButton class="absolute inset-y-0 right-0 flex items-center pr-2">
        <font-awesome-icon icon="fa-solid fa-caret-down" class="h-5 w-5 text-gray-400" aria-hidden="true" />
      </ComboboxButton> -->

    <option v-for="op in filteredOptions" :key="op" :value="op">
      {{ op }} <font-awesome-icon icon="fa-solid fa-check" v-show="op === selectedOption" />
    </option>

    <!-- <ComboboxOptions
      class="scrollbox border-secondary mt-1 max-h-40 overflow-y-auto overflow-x-hidden rounded-md border-2 shadow">
      <ComboboxOption v-for="op in filteredOptions" :key="op" :value="op"
        class="hover:bg-primary flex items-center justify-between px-3 py-2 text-sm font-medium hover:text-white">
        <span>{{ op }}</span>
        <font-awesome-icon icon="fa-solid fa-check" v-show="op === selectedOption" />
      </ComboboxOption>
    </ComboboxOptions> -->
  </select>
</template>
