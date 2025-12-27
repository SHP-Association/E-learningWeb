<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-blue-50 px-4">
    <div class="bg-white p-8 rounded-lg shadow-xl border border-gray-200 w-full max-w-md">
      <h1 class="text-3xl font-bold mb-6 text-blue-800 text-center">Reset Your Password</h1>
      <p class="text-gray-600 mb-6 text-center">
        Enter your email address and we'll send you a link to reset your password.
      </p>

      <!-- Success message -->
      <div v-if="message" class="bg-blue-100 border border-blue-400 text-blue-700 px-4 py-3 rounded-lg mb-4" role="alert">
        {{ message }}
      </div>

      <!-- Error message -->
      <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg mb-4" role="alert">
        {{ error }}
      </div>

      <form @submit.prevent="handleSubmit" class="w-full space-y-4">
        <div class="form-group">
          <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email</label>
          <input
            type="email"
            name="email"
            id="email"
            v-model="email"
            required
            autocomplete="email"
            :disabled="isSubmitting"
            class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-gray-50 text-gray-800 transition duration-150"
          />
        </div>

        <button 
          type="submit"
          :disabled="isSubmitting"
          class="w-full py-3 mt-2 rounded-lg bg-blue-600 text-white font-semibold hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:outline-none transition duration-200 flex justify-center items-center"
        >
          <span v-if="isSubmitting" class="flex items-center space-x-2">
            <svg class="animate-spin h-5 w-5 text-white -ml-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            <span>Sending...</span>
          </span>
          <span v-else>Send Reset Link</span>
        </button>
      </form>

      <p class="mt-6 text-center text-gray-700">
        Remembered your password?
        <button @click="router.push('/login')" class="text-blue-600 hover:text-blue-800 hover:underline font-bold transition duration-200">
          Log in
        </button>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
import { apiService } from '../services/api.service';

const email = ref('');
const message = ref('');
const error = ref('');
const isSubmitting = ref(false);
const BACKEND_URL = apiService.baseURL;

const handleSubmit = async () => {
  message.value = '';
  error.value = '';

  if (!email.value.trim()) {
    error.value = 'Please enter your email address.';
    return;
  }

  isSubmitting.value = true;

  try {
    const data = await apiService.requestPasswordReset(email.value);

    message.value = data.message || 'If your email is registered, you will receive a password reset link.';
    setTimeout(() => {
      router.push('/password_reset/done');
    }, 1500);
  } catch (err: any) {
    error.value = err.message || 'Failed to send password reset email. Please try again.';
  } finally {
    error.value = 'An error occurred. Please try again later.';
  } finally {
    isSubmitting.value = false;
  }
};
</script>

<style scoped>
/* Optional: add smooth fade-in for messages */
</style>
