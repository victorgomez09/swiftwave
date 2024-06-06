<script setup>
import FilledButton from '@/views/components/FilledButton.vue'
import { ref } from 'vue'
import { useToast } from 'vue-toastification'

const toast = useToast()

const props = defineProps({
  domain: {
    type: Object,
    required: true
  },
  deleteDomain: {
    type: Function,
    required: true
  },
  verifyDns: {
    type: Function,
    required: true
  },
  issueSsl: {
    type: Function,
    required: true
  },
  viewSsl: {
    type: Function,
    required: true
  }
})

const verifyingDns = ref(false)

const verifyDnsPointing = async () => {
  verifyingDns.value = true
  toast.info(`Verifying DNS for ${props.domain.name}`)
  const isPointed = await props.verifyDns(props.domain.name)
  verifyingDns.value = false
  if (isPointed) {
    toast.success(`DNS for ${props.domain.name} is pointed correctly`)
  } else {
    toast.error(`DNS for ${props.domain.name} is not pointed correctly`)
  }
}
</script>

<template>
  <div class="indicator w-full">
    <span class="indicator-item badge badge-sm" :class="{
      'badge-secondary': domain.sslStatus === 'none',
      'badge-warning': domain.sslStatus === 'pending',
      'badge-success': domain.sslStatus === 'issued',
      'badge-error': domain.sslStatus === 'failed'
    }"></span>

    <div class="card bg-base-300 w-full">
      <div class="card-body p-4">
        <h3 class="card-title">{{ domain.name }}</h3>

        <div class="flex flex-col gap-2">
          <!-- Status -->
          <div class="flex gap-2">
            <span>Status:</span>
            <span v-if="domain.sslStatus === 'none'" class="font-semibold">N/A</span>
            <span v-else-if="domain.sslStatus === 'pending'" class="font-semibold">Pending</span>
            <span v-else-if="domain.sslStatus === 'issued'" class="font-semibold">Issued</span>
            <span v-else-if="domain.sslStatus === 'failed'" class="font-semibold">Failed</span>
          </div>

          <!-- Issuer -->
          <div class="flex gap-2">
            <span>Issuer:</span>
            <span class="font-semibold">{{ domain.sslIssuer !== '' ? domain.sslIssuer : '---' }}</span>
          </div>

          <!-- Auto renew -->
          <div class="flex gap-2">
            <span>Auto renew:</span>
            <span v-if="domain.sslAutoRenew" class="font-semibold">Enabled</span>
            <span v-else class="font-semibold">Disabled</span>
          </div>

          <div class="card-actions justify-end items-center">
            <FilledButton :click="() => viewSsl(domain.id)" :disabled="domain.sslStatus !== 'issued'" slim
              type="secondary" size="sm">
              View SSL
            </FilledButton>

            <!-- More actions -->
            <div class="dropdown">
              <FilledButton tabindex="0" role="button" size="sm">
                <font-awesome-icon :icon="['fas', 'ellipsis-vertical']" />
              </FilledButton>
              <ul tabindex="0" class="dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-52">
                <li @click="issueSsl(domain)"><a>Issue SSL</a></li>
                <li @click="verifyDnsPointing"><a>Verify DNS</a></li>
                <div class="divider divider-vertical m-0 p-0"></div>
                <li @click="deleteDomain(domain)" class="text-error"><a>Delete</a></li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
