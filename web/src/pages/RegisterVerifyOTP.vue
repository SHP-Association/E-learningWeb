<template>
  <AuthShell
    title="Verify Your Email"
    subtitle="Enter the 6-digit OTP sent to your email address."
    tone="secondary"
  >
    <template #status>
      <AlertMessage
        v-if="userStore.error"
        type="error"
        :message="userStore.error"
        compact
      />
      <AlertMessage
        v-else-if="statusMessage"
        type="success"
        :message="statusMessage"
        compact
      />
    </template>

    <form @submit.prevent="handleVerify" class="space-y-5">
      <AppInput
        id="email"
        v-model="email"
        label="Email Address"
        name="email"
        type="email"
        autocomplete="email"
        required
        tone="secondary"
      />

      <AppInput
        id="otp"
        v-model="otp"
        label="OTP Code"
        name="otp"
        required
        tone="secondary"
        hint="Enter the 6-digit code."
      />

      <AppButton
        type="submit"
        tone="secondary"
        :loading="userStore.loading"
        loading-label="Verifying..."
        :disabled="userStore.loading"
        block
      >
        Verify OTP
      </AppButton>
    </form>

    <div class="flex items-center justify-between gap-3 pt-2">
      <AppButton
        type="button"
        variant="ghost"
        tone="secondary"
        :disabled="resendDisabled || userStore.loading"
        @click="handleResend"
      >
        {{ resendDisabled ? `Resend in ${cooldownRemaining}s` : 'Resend OTP' }}
      </AppButton>

      <AppButton
        type="button"
        variant="ghost"
        tone="secondary"
        @click="router.push('/register')"
      >
        Back to Signup
      </AppButton>
    </div>
  </AuthShell>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import AlertMessage from '../components/AlertMessage.vue';
import AppButton from '../components/ui/AppButton.vue';
import AppInput from '../components/ui/AppInput.vue';
import AuthShell from '../components/ui/AuthShell.vue';
import { useUserStore } from '../stores/userStore';

const route = useRoute();
const router = useRouter();
const userStore = useUserStore();

const email = ref('');
const otp = ref('');
const statusMessage = ref('');
const cooldownRemaining = ref(0);
let cooldownTimer: number | null = null;

const resendDisabled = computed(() => cooldownRemaining.value > 0);

const startCooldown = (seconds = 60) => {
  cooldownRemaining.value = seconds;

  if (cooldownTimer !== null) {
    window.clearInterval(cooldownTimer);
  }

  cooldownTimer = window.setInterval(() => {
    if (cooldownRemaining.value <= 1) {
      cooldownRemaining.value = 0;
      if (cooldownTimer !== null) {
        window.clearInterval(cooldownTimer);
        cooldownTimer = null;
      }
      return;
    }

    cooldownRemaining.value -= 1;
  }, 1000);
};

const handleVerify = async () => {
  userStore.clearError();
  statusMessage.value = '';

  const ok = await userStore.verifyOtp({
    email: email.value.trim(),
    otp: otp.value.trim(),
  });

  if (!ok) {
    return;
  }

  statusMessage.value = 'Email verified successfully.';
  router.push('/onboarding');
};

const handleResend = async () => {
  userStore.clearError();
  statusMessage.value = '';

  const ok = await userStore.resendOtp();
  if (!ok) {
    return;
  }

  statusMessage.value = 'A new OTP has been sent to your email.';
  startCooldown(60);
};

onMounted(() => {
  const queryEmail = typeof route.query.email === 'string' ? route.query.email : '';
  email.value = queryEmail || userStore.pendingVerificationEmail || '';
});

onUnmounted(() => {
  if (cooldownTimer !== null) {
    window.clearInterval(cooldownTimer);
  }
});
</script>
