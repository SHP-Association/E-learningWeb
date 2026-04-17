<template>
  <div class="ui-page-shell">
    <AppCard size="md" padding="lg" elevated centered>
      <template #header>
        <div class="text-center">
          <h1 class="text-3xl font-bold text-blue-800">Set New Password</h1>
          <p class="mt-3 text-gray-600">Enter your new password below.</p>
        </div>
      </template>

      <div v-if="message || error" class="mb-6 space-y-3">
        <AlertMessage
          v-if="message"
          type="success"
          :message="message"
          compact
        />
        <AlertMessage
          v-if="error"
          type="error"
          :message="error"
          compact
        />
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-4">
        <AppInput
          id="new_password"
          v-model="newPassword"
          label="New Password"
          name="new_password"
          type="password"
          autocomplete="new-password"
          required
          tone="primary"
        />

        <AppInput
          id="confirm_password"
          v-model="confirmPassword"
          label="Confirm New Password"
          name="confirm_password"
          type="password"
          autocomplete="new-password"
          required
          tone="primary"
        />

        <AppButton
          type="submit"
          tone="primary"
          :loading="isSubmitting"
          loading-label="Resetting..."
          :disabled="isSubmitting"
          block
        >
          Reset Password
        </AppButton>
      </form>
    </AppCard>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import AlertMessage from '../components/AlertMessage.vue';
import AppButton from '../components/ui/AppButton.vue';
import AppCard from '../components/ui/AppCard.vue';
import AppInput from '../components/ui/AppInput.vue';
import { apiService } from '../services/api.service';

const props = defineProps<{
  uid: string;
  token: string;
}>();

const router = useRouter();

const newPassword = ref('');
const confirmPassword = ref('');
const error = ref('');
const message = ref('');
const isSubmitting = ref(false);

const handleSubmit = async () => {
  error.value = '';
  message.value = '';

  if (newPassword.value.length < 6) {
    error.value = 'New password must be at least 6 characters long.';
    return;
  }

  if (newPassword.value !== confirmPassword.value) {
    error.value = 'Passwords do not match.';
    return;
  }

  isSubmitting.value = true;

  try {
    const data = await apiService.confirmPasswordReset(props.uid, props.token, newPassword.value);

    message.value = data.message || 'Your password has been reset successfully.';
    setTimeout(() => {
      router.push('/password_reset/complete');
    }, 1200);
  } catch (err: any) {
    error.value = err.message || 'Password reset failed. The link may be invalid or expired.';
  } finally {
    isSubmitting.value = false;
  }
};
</script>
