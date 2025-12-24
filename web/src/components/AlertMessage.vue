<template>
  <div
    class="px-6 py-4 rounded-lg text-center mx-auto max-w-md mt-8 shadow-md"
    :class="alertClass"
  >
    <p class="text-lg">{{ message }}</p>
    <p v-if="details" class="text-sm mt-2">{{ details }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';

export default defineComponent({
  name: 'AlertMessage',
  props: {
    type: {
      type: String,
      default: 'error',
      validator: (value: string) => ['error', 'success', 'warning', 'info'].includes(value),
    },
    message: {
      type: String,
      required: true,
    },
    details: {
      type: String,
      default: '',
    },
  },
  setup(props) {
    const alertClass = computed(() => {
      const classes = {
        error: 'bg-red-100 border border-red-400 text-red-700',
        success: 'bg-green-100 border border-green-400 text-green-700',
        warning: 'bg-yellow-100 border border-yellow-400 text-yellow-700',
        info: 'bg-blue-100 border border-blue-400 text-blue-700',
      };
      return classes[props.type as keyof typeof classes] || classes.error;
    });

    return {
      alertClass,
    };
  },
});
</script>
