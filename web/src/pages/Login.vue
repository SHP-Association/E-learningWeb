<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100 p-4 sm:p-6 font-sans">
    <div class="bg-white p-6 sm:p-8 rounded-2xl shadow-2xl border border-gray-100 w-full max-w-sm sm:max-w-md transform transition-all duration-300 hover:shadow-3xl">
      <h1 class="text-3xl sm:text-4xl font-extrabold mb-7 text-center text-blue-700 animate-fade-in-down">
        Welcome Back!
      </h1>
      <p class="text-center text-gray-600 mb-6 text-md sm:text-lg animate-fade-in delay-200">
        Log in to access your courses and profile.
      </p>

      <div v-if="errorMessage" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg mb-6 text-sm sm:text-base flex items-center space-x-3 shadow-md animate-fade-in">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 flex-shrink-0 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <p class="font-semibold">{{ errorMessage }}</p>
      </div>

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
              isSubmitting
                ? 'bg-blue-400 cursor-not-allowed'
                : 'bg-blue-600 text-white hover:bg-blue-700 focus:ring-blue-500 focus:ring-opacity-75 shadow-md hover:shadow-lg'
            ]"
            :disabled="isSubmitting"
          >
            <span v-if="isSubmitting" class="flex items-center">
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
import { ref, inject } from 'vue';
import { useRouter } from 'vue-router';

interface User {
  id: number;
  username: string;
  email: string;
  role: 'student' | 'instructor' | string;
  is_staff: boolean;
}

interface InjectedContext {
  handleLogin: (user: User) => void;
}

// Use Vue Router
const router = useRouter();

// Injected login handler
const { handleLogin } = inject('user-context', {} as InjectedContext);

const BACKEND_URL = import.meta.env.VITE_APP_BACKEND_URL;

const username = ref('');
const password = ref('');
const errorMessage = ref('');
const isSubmitting = ref(false);

function getCookie(name: string): string | null {
  let cookieValue = null;
  if (document.cookie && document.cookie !== '') {
    const cookies = document.cookie.split(';');
    for (let i = 0; i < cookies.length; i++) {
      const cookie = cookies[i].trim();
      if (cookie.substring(0, name.length + 1) === (name + '=')) {
        cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
        break;
      }
    }
  }
  return cookieValue;
}

const handleSubmit = async () => {
  errorMessage.value = '';
  isSubmitting.value = true;

  try {
    const csrfToken = getCookie('csrftoken');
    if (!csrfToken) {
      errorMessage.value = 'CSRF token not found. Please refresh and try again.';
      isSubmitting.value = false;
      return;
    }

    const response = await fetch(`${BACKEND_URL}/api/login/`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'X-CSRFToken': csrfToken,
      },
      credentials: 'include',
      body: JSON.stringify({ username: username.value, password: password.value }),
    });

    if (!response.ok) {
      errorMessage.value = 'Invalid username or password.';
      isSubmitting.value = false;
      return;
    }

    const userRes = await fetch(`${BACKEND_URL}/api/users/`, { credentials: 'include' });
    if (!userRes.ok) {
      errorMessage.value = 'Failed to retrieve user profile.';
      isSubmitting.value = false;
      return;
    }

    const users: User[] = await userRes.json();
    const user = users.find(u => u.username === username.value);

    if (!user) {
      errorMessage.value = 'User profile not found.';
      isSubmitting.value = false;
      return;
    }

    handleLogin(user);
    router.push('/profile');

  } catch (err) {
    console.error("Login error:", err);
    errorMessage.value = 'An unexpected error occurred.';
  } finally {
    isSubmitting.value = false;
  }
};
</script>
