<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-gray-100 px-4">
    <div class="bg-white p-8 rounded-lg shadow-xl border border-gray-200 w-full max-w-md">
      <h1 class="text-3xl font-bold mb-6 text-blue-800 text-center">Set New Password</h1>
      <p class="text-gray-600 mb-6 text-center">
        Enter your new password below.
      </p>

      <div v-if="message" class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded-lg mb-4" role="alert">
        {{ message }}
      </div>
      <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg mb-4" role="alert">
        {{ error }}
      </div>

      <form @submit.prevent="handleSubmit" class="w-full space-y-4">
        <div class="form-group">
          <label for="new_password" class="block text-sm font-medium text-gray-700 mb-1">New Password</label>
          <input
            type="password"
            name="new_password"
            id="new_password"
            class="w-full p-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white text-gray-800"
            required
            autocomplete="new-password"
            v-model="newPassword"
          />
        </div>

        <div class="form-group">
          <label for="confirm_password" class="block text-sm font-medium text-gray-700 mb-1">Confirm New Password</label>
          <input
            type="password"
            name="confirm_password"
            id="confirm_password"
            class="w-full p-3 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white text-gray-800"
            required
            autocomplete="new-password"
            v-model="confirmPassword"
          />
        </div>

        <button 
          type="submit" 
          class="bg-blue-900 text-white hover:bg-blue-700 px-6 py-3 rounded-md font-semibold transition w-full mt-4"
          :disabled="isSubmitting"
        >
          {{ isSubmitting ? 'Resetting...' : 'Reset Password' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';

const props = defineProps<{
  navigate: (path: string) => void;
  uid: string;
  token: string;
}>();

import { apiService } from '../services/api.service';

const newPassword = ref('');
const confirmPassword = ref('');
const error = ref('');
const message = ref('');
const isSubmitting = ref(false);
const BACKEND_URL = apiService.baseURL;

// Optional: Initial validation check from JSX is omitted as the form handles submission
// onMounted(() => { /* initial validation logic */ });

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
    // Redirect to a final page
    props.navigate('/password_reset/complete');
  } catch (err: any) {
    error.value = err.message || 'Password reset failed. The link may be invalid or expired.';
  } finally {
    isSubmitting.value = false;
  }
};
</script>

