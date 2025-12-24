<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-purple-50 to-pink-100 p-4 sm:p-6 font-sans">
    <div class="bg-white p-6 sm:p-8 rounded-2xl shadow-2xl border border-gray-100 w-full max-w-sm sm:max-w-md transform transition-all duration-300 hover:shadow-3xl">
      <h1 class="text-3xl sm:text-4xl font-extrabold mb-7 text-center text-purple-700 animate-fade-in-down">
        Join SHP-Learner!
      </h1>
      <p class="text-center text-gray-600 mb-6 text-md sm:text-lg animate-fade-in delay-200">
        Create your account and start learning today.
      </p>

      <!-- Error Messages -->
      <div v-if="hasErrors" class="bg-red-50 border border-red-300 text-red-700 px-4 py-3 rounded-lg mb-6 text-sm sm:text-base flex flex-col space-y-1 shadow-md animate-fade-in">
        <p class="font-semibold">Please correct the following errors:</p>
        <ul class="list-disc list-inside mt-1">
          <li v-for="(error, key) in errors" :key="key">{{ error }}</li>
        </ul>
      </div>

      <!-- Success Message -->
      <div v-if="successMessage" class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded-lg mb-6 text-sm sm:text-base shadow-md animate-fade-in">
        <p class="font-semibold text-center">{{ successMessage }}</p>
      </div>

      <!-- Registration Form -->
      <form @submit.prevent="handleSubmit" class="space-y-5 animate-fade-in delay-300">
        <!-- Role Selection -->
        <div class="form-group">
          <label class="block text-sm font-medium text-gray-700 mb-1">I want to register as:</label>
          <div class="flex space-x-4">
            <label class="inline-flex items-center">
              <input type="radio" name="role" value="student" v-model="role" class="form-radio h-5 w-5 text-purple-600 border-gray-300 focus:ring-purple-500" />
              <span class="ml-2 text-gray-700 font-medium">Student</span>
            </label>
            <label class="inline-flex items-center">
              <input type="radio" name="role" value="instructor" v-model="role" class="form-radio h-5 w-5 text-purple-600 border-gray-300 focus:ring-purple-500" />
              <span class="ml-2 text-gray-700 font-medium">Instructor</span>
            </label>
          </div>
        </div>

        <!-- Inputs -->
        <div class="form-group">
          <label for="username" class="block text-sm font-medium text-gray-700 mb-1">Username</label>
          <input id="username" v-model="username" class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500 bg-gray-50 text-gray-800 transition duration-150" required />
        </div>

        <div class="form-group">
          <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email Address</label>
          <input id="email" type="email" v-model="email" class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500 bg-gray-50 text-gray-800 transition duration-150" required />
        </div>

        <div class="form-group">
          <label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
          <input id="password" type="password" v-model="password" class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500 bg-gray-50 text-gray-800 transition duration-150" required />
          <p class="text-xs text-gray-500 mt-1">Min 8 chars, incl. uppercase, lowercase, number, special char.</p>
        </div>

        <div class="form-group">
          <label for="confirmPassword" class="block text-sm font-medium text-gray-700 mb-1">Confirm Password</label>
          <input id="confirmPassword" type="password" v-model="confirmPassword" class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500 bg-gray-50 text-gray-800 transition duration-150" required />
        </div>

        <!-- Submit Button -->
        <div>
          <button
            type="submit"
            :class="[
              'w-full py-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-offset-2 transition duration-200 text-base sm:text-lg font-semibold flex items-center justify-center mt-7',
              userStore.loading
                ? 'bg-purple-400 cursor-not-allowed'
                : 'bg-purple-600 text-white hover:bg-purple-700 focus:ring-purple-500 focus:ring-opacity-75 shadow-md hover:shadow-lg'
            ]"
            :disabled="userStore.loading"
          >
            <span v-if="userStore.loading" class="flex items-center">
              <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Registering...
            </span>
            <span v-else>Register</span>
          </button>
        </div>
      </form>

      <!-- Login Link -->
      <p class="mt-8 text-center text-gray-700 text-sm sm:text-base animate-fade-in delay-500">
        Already have an account?
        <button @click="router.push('/login')" class="text-purple-600 hover:text-purple-800 hover:underline font-bold transition duration-200">
          Log in
        </button>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '../stores/userStore';
import { useRegisterSEO } from '../composables/useSEO';

const router = useRouter();
const userStore = useUserStore();

// Set SEO meta tags for register page
useRegisterSEO();

const username = ref('');
const email = ref('');
const password = ref('');
const confirmPassword = ref('');
const role = ref<'student' | 'instructor'>('student');
const errors = reactive<Record<string, string>>({});
const successMessage = ref('');

const hasErrors = computed(() => Object.keys(errors).length > 0);

const validateForm = (): boolean => {
  Object.keys(errors).forEach(key => delete errors[key]);

  if (!username.value.trim()) errors.username = 'Username is required.';
  else if (username.value.trim().length < 3) errors.username = 'Username must be at least 3 characters.';

  if (!email.value.trim()) errors.email = 'Email is required.';
  else if (!/\S+@\S+\.\S+/.test(email.value)) errors.email = 'Invalid email address.';

  if (!password.value) errors.password = 'Password is required.';
  else if (password.value.length < 8) errors.password = 'Password must be at least 8 characters.';
  else if (!/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])/.test(password.value))
    errors.password = 'Password must include uppercase, lowercase, number, special character.';

  if (password.value !== confirmPassword.value) errors.confirmPassword = 'Passwords do not match.';

  return Object.keys(errors).length === 0;
};

const handleSubmit = async () => {
  successMessage.value = '';
  Object.keys(errors).forEach(key => delete errors[key]);

  if (!validateForm()) return;

  try {
    await userStore.registerUser({
      username: username.value,
      email: email.value,
      password: password.value,
      role: role.value,
    });

    successMessage.value = 'Registration successful! Redirecting to profile...';

    setTimeout(() => {
      router.push('/profile');
    }, 1500);

  } catch (err: any) {
    if (err.errors) {
      if (err.errors.username) errors.username = Array.isArray(err.errors.username) ? err.errors.username[0] : err.errors.username;
      if (err.errors.email) errors.email = Array.isArray(err.errors.email) ? err.errors.email[0] : err.errors.email;
    } else {
      errors.apiError = err.message || 'Registration failed.';
    }
  }
};
</script>

<style scoped>
@keyframes fade-in-down { from {opacity:0; transform:translateY(-20px);} to {opacity:1; transform:translateY(0);} }
@keyframes fade-in { from {opacity:0;} to {opacity:1;} }
.animate-fade-in-down { animation: fade-in-down 0.8s ease-out forwards; }
.animate-fade-in { animation: fade-in 0.6s ease-out forwards; }
.delay-200 { animation-delay: 0.2s; }
.delay-300 { animation-delay: 0.3s; }
.delay-500 { animation-delay: 0.5s; }
.hover\:shadow-3xl:hover { box-shadow: 0 25px 50px -12px rgba(0,0,0,0.25); }
</style>
