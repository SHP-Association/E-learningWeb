<template>
  <div class="flex items-center">
    <svg
      v-for="i in 5"
      :key="i"
      class="w-4 h-4"
      :class="i <= Math.round(Number(rating || 0)) ? 'text-yellow-400' : 'text-gray-300'"
      fill="currentColor"
      viewBox="0 0 20 20"
    >
      <path
        d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.538 1.118l-2.8-2.034a1 1 0 00-1.176 0l-2.8 2.034c-.783.57-1.838-.197-1.538-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.381-1.81.588-1.81h3.462a1 1 0 00.95-.69l1.07-3.292z"
      />
    </svg>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';

export default defineComponent({
  name: 'StarRating',
  props: {
    rating: {
      type: [Number, String],
      required: true,
    },
    maxRating: {
      type: Number,
      default: 5,
    },
  },
  setup(props) {
    const numericRating = computed(() => Number(props.rating || 0));
    const fullStars = computed(() => Math.floor(numericRating.value));
    const hasHalfStar = computed(() => numericRating.value % 1 >= 0.5);
    const emptyStars = computed(() => props.maxRating - fullStars.value - (hasHalfStar.value ? 1 : 0));

    return {
      fullStars,
      hasHalfStar,
      emptyStars,
    };
  },
});
</script>
