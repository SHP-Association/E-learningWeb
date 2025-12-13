<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100 p-4 sm:p-6 font-sans">
    <div class="bg-white p-6 sm:p-8 rounded-2xl shadow-2xl border border-gray-100 w-full max-w-sm sm:max-w-md transform transition-all duration-300 hover:shadow-3xl">
      <h1 class="text-3xl sm:text-4xl font-extrabold mb-7 text-center text-blue-700 animate-fade-in-down">
        Welcome Back!
      </h1>
      <p class="text-center text-gray-600 mb-6 text-md sm:text-lg animate-fade-in delay-200">
        Log in to access your courses and profile.
      </p>

      <AlertMessage
        v-if="errorMessage"
        type="error"
        :message="errorMessage"
      />

      <form @submit.prevent="handleSubmit" class="space-y-5 animate-fade-in delay-300">
        <div class="form-group">
          <label for="username" class="block text-sm font-medium text-gray-700 mb-1">Username</label>
          <input
            type="text"
            id="username"
            v-model="username"
            class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-gray-50 text-gray-800 transition duration-150"
            required
            autocomplete="username"
          />
        </div>

        <div class="form-group">
          <label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-gray-50 text-gray-800 transition duration-150"
            required
            autocomplete="current-password"
          />
        </div>

        <div class="flex justify-between items-center text-sm">
          <button
            type="button"
            @click="router.push('/password_reset')"
            class="text-blue-500 hover:text-blue-700 hover:underline transition duration-200 font-medium"
          >
            Forgot password?
          </button>
        </div>

        <div>
          <button
            type="submit"
            :class="[
              'w-full py-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-offset-2 transition duration-200 text-lg font-semibold flex items-center justify-center',
              userStore.loading
                ? 'bg-blue-400 cursor-not-allowed'
                : 'bg-blue-600 text-white hover:bg-blue-700 focus:ring-blue-500 focus:ring-opacity-75 shadow-md hover:shadow-lg'
            ]"
            :disabled="userStore.loading"
          >
            <span v-if="userStore.loading" class="flex items-center">
              <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Logging In...
            </span>
            <span v-else>Login</span>
          </button>
        </div>
      </form>

      <p class="mt-8 text-center text-gray-700 text-sm sm:text-base animate-fade-in delay-500">
        New to SHP-Learner?
        <button @click="router.push('/register')" class="text-blue-600 hover:text-blue-800 hover:underline font-bold transition duration-200">
          Create an account
        </button>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '../stores/userStore';
import AlertMessage from '../components/AlertMessage.vue';

const router = useRouter();
const userStore = useUserStore();

const username = ref('');
const password = ref('');
const errorMessage = ref('');

const handleSubmit = async () => {
  errorMessage.value = '';

  try {
    await userStore.loginUser({
      username: username.value,
      password: password.value,
    });
    
    router.push('/profile');
  } catch (err: any) {
    errorMessage.value = err.message || 'Invalid username or password.';
  }
};
</script>
