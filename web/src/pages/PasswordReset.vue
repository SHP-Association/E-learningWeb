<template>
  <AuthShell
    title="Reset Your Password"
    subtitle="Enter your email address and we'll send you a link to reset your password."
    tone="primary"
  >
    <template #status>
      <AlertMessage
        v-if="error"
        type="error"
        :message="error"
        compact
      />
      <AppDialog
        v-model="isSuccessDialogOpen"
        title="Email Sent"
      >
        <p class="text-base text-zinc-300">{{ message }}</p>
        <template #footer>
          <AppButton
            tone="primary"
            block
            @click="router.push('/login')"
          >
            Go to Login
          </AppButton>
        </template>
      </AppDialog>
      <AppDialog
        v-model="isErrorDialogOpen"
        title="Reset Failed"
      >
        <p class="text-base text-zinc-300">{{ error }}</p>
        <template #footer>
          <AppButton
            tone="danger"
            block
            @click="isErrorDialogOpen = false"
          >
            Close
          </AppButton>
        </template>
      </AppDialog>
    </template>

    <form @submit.prevent="handleSubmit" class="space-y-4">
      <AppInput
        id="email"
        v-model="email"
        label="Email"
        name="email"
        type="email"
        autocomplete="email"
        required
        :disabled="isSubmitting"
        tone="primary"
      />

      <AppButton
        type="submit"
        tone="primary"
        :loading="isSubmitting"
        loading-label="Sending..."
        :disabled="isSubmitting"
        block
      >
        Send Reset Link
      </AppButton>
    </form>

    <template #footer>
      <p class="text-center text-gray-700">
        Remembered your password?
        <AppButton
          type="button"
          variant="ghost"
          tone="primary"
          @click="router.push('/login')"
        >
          Log in
        </AppButton>
      </p>
    </template>
  </AuthShell>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import AlertMessage from '../components/AlertMessage.vue';
import AppButton from '../components/ui/AppButton.vue';
import AppInput from '../components/ui/AppInput.vue';
import AuthShell from '../components/ui/AuthShell.vue';
import AppDialog from '../components/ui/AppDialog.vue';
import { apiService } from '../services/api.service';

const router = useRouter();

const email = ref('');
const message = ref('');
const error = ref('');
const isSubmitting = ref(false);
const isSuccessDialogOpen = ref(false);
const isErrorDialogOpen = ref(false);

const handleSubmit = async () => {
  message.value = '';
  error.value = '';
  isSuccessDialogOpen.value = false;
  isErrorDialogOpen.value = false;

  if (!email.value.trim()) {
    error.value = 'Please enter your email address.';
    return;
  }

  isSubmitting.value = true;

  try {
    const data = await apiService.requestPasswordReset(email.value);

    message.value = data.message || 'If your email is registered, you will receive a password reset link.';
    isSuccessDialogOpen.value = true;
  } catch (err: any) {
    error.value = err.message || 'Failed to send password reset email. Please try again.';
    isErrorDialogOpen.value = true;
  } finally {
    isSubmitting.value = false;
  }
};
</script>
