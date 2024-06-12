<script setup>
import { useToast } from 'vue-toastification'
import { ref } from 'vue'

defineProps({
  content: {
    type: Array
  }
})

const toast = useToast()
const showCopyBorder = ref(false)
const textDivRef = ref(null)
const copyToClipboard = () => {
  if (textDivRef.value === null) {
    return
  }
  let isSuccess
  if ('clipboard' in navigator) {
    navigator.clipboard.writeText(textDivRef.value.innerText)
    isSuccess = true
  } else {
    const textArea = document.createElement('textarea')
    textArea.value = textDivRef.value.innerText
    textArea.style.opacity = '0'
    document.body.appendChild(textArea)
    textArea.focus()
    textArea.select()
    try {
      isSuccess = document.execCommand('copy')
    } catch (err) {
      isSuccess = false
    }
    document.body.removeChild(textArea)
  }
  if (isSuccess) {
    toast.success('Copied to clipboard')
    showCopyBorder.value = true
    setTimeout(() => {
      showCopyBorder.value = false
    }, 2000)
  } else {
    toast.error('Failed to copy to clipboard')
  }
}
</script>

<template>
  <div class="mockup-code">
    <div ref="textDivRef">
      <pre v-for="contentText in content ">{{ contentText }}</pre>
    </div>
    <div @click="copyToClipboard"
      class="absolute right-0 top-0 m-3 flex h-10 w-10 cursor-pointer items-center justify-center rounded-lg border-2 border-base-300 bg-base-content transition-all">
      <font-awesome-icon icon="fa-solid fa-clipboard" class="text-xl" />
    </div>
  </div>
</template>