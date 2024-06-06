<script setup>
defineProps({
  argKey: {
    type: String,
    required: true
  },
  description: {
    type: String,
    required: true
  },
  value: {
    type: String,
    required: true
  },
  updateBuildArg: {
    type: Function,
    required: true
  }
})
</script>

<template>
  <div class="mt-4" :for="argKey">
    <label class="form-control w-full">
      <div class="label">
        <span class="label-text">{{ argKey }}</span>
      </div>

      <!-- <br v-if="description !== ''" />
      <span class="text-xs font-normal text-gray-800">{{ description }}</span> -->
      <input :id="argKey" :name="argKey" :value="value" autocomplete="off" class="input input-bordered w-full"
        placeholder="{{description}}" type="text" @input="(e) => updateBuildArg(e.target.value)" />
    </label>
  </div>
</template>

<style scoped></style>
