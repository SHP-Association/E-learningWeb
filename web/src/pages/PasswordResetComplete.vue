<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-gray-100 px-4">
    <div class="bg-white p-8 rounded-lg shadow-xl border border-gray-200 w-full max-w-md text-center">
      <h1 class="text-4xl font-bold mb-6 text-blue-800">Password Reset Complete</h1>
      <p class="text-gray-700 mb-6">
        Your password has been set. You may now log in with your new password.
      </p>
      <button
        @click="navigate('/login')"
        class="bg-blue-900 text-white hover:bg-blue-700 px-6 py-3 rounded-lg font-semibold shadow transition"
      >
        Log in
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  navigate: (path: string) => void;
}>();
</script>