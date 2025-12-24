<template>
  <div class="flex justify-center items-center py-16">
    <div class="loader ease-linear rounded-full border-8 border-t-8 border-gray-200 h-24 w-24"></div>
    <p class="text-center text-xl text-blue-600 ml-4">{{ message }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'LoadingSpinner',
  props: {
    message: {
      type: String,
      default: 'Loading...',
    },
  },
});
</script>

<style scoped>
.loader {
  border-top-color: #3b82f6;
  animation: spinner 1.5s linear infinite;
}

@keyframes spinner {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
