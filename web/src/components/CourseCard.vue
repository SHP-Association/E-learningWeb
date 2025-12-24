<template>
  <div
    class="yt-card group bg-white rounded-xl shadow-md hover:shadow-2xl border border-gray-100 transition-all duration-300 flex flex-col cursor-pointer overflow-hidden"
    @click="$emit('click')"
  >
    <!-- Thumbnail -->
    <div class="relative w-full aspect-video bg-gray-100 overflow-hidden">
      <img
        v-if="course.thumbnail"
        :src="course.thumbnail"
        :alt="course.title"
        class="w-full h-full object-contain bg-white group-hover:scale-105 transition-transform duration-300"
        style="max-height: 220px"
      />
      <img
        v-else-if="course.category?.image"
        :src="course.category.image"
        :alt="course.category.name"
        class="w-full h-full object-cover opacity-80"
      />
      <div
        v-else
        class="flex items-center justify-center h-full w-full text-gray-400"
      >
        <svg class="w-12 h-12" fill="currentColor" viewBox="0 0 24 24">
          <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8z" />
        </svg>
      </div>

      <!-- Badges -->
      <span
        v-if="featured"
        class="absolute top-2 left-2 bg-yellow-400 text-blue-900 text-xs font-semibold px-2 py-1 rounded z-10 shadow animate-pulse"
      >
        🌟 FEATURED
      </span>
      <span
        v-if="course.is_free"
        class="absolute top-2 right-2 bg-green-500 text-white text-xs font-semibold px-2 py-1 rounded z-10 shadow"
      >
        FREE
      </span>
    </div>

    <!-- Card Body -->
    <div class="flex flex-row p-3 gap-3">
      <div class="flex-shrink-0">
        <img
          v-if="course.instructor?.profile_picture"
          :src="course.instructor.profile_picture"
          :alt="course.instructor.username"
          class="w-10 h-10 rounded-full object-cover border border-gray-200"
        />
        <div
          v-else
          class="w-10 h-10 rounded-full bg-blue-100 flex items-center justify-center text-blue-700 font-bold text-lg"
        >
          {{ course.instructor?.username?.[0]?.toUpperCase() || '?' }}
        </div>
      </div>

      <div class="flex-1 min-w-0">
        <h3
          class="text-base md:text-lg font-bold text-blue-900 mb-1 truncate group-hover:text-blue-700 transition-colors"
        >
          {{ course.title }}
        </h3>
        <div class="flex items-center text-xs text-gray-500 mb-1">
          <StarRating :rating="course.average_rating || 0" />
          <span class="ml-1 font-semibold">{{ (course.average_rating || 0).toFixed(1) }}</span>
          <span class="ml-1">({{ course.number_of_reviews || 0 }})</span>
        </div>

        <div class="flex flex-wrap gap-1 text-xs mb-1">
          <span class="bg-blue-50 text-blue-700 px-2 py-0.5 rounded">
            {{ course.category?.name || 'Uncategorized' }}
          </span>
          <span class="bg-indigo-50 text-indigo-700 px-2 py-0.5 rounded">
            {{ course.level ? course.level.charAt(0).toUpperCase() + course.level.slice(1) : 'N/A' }}
          </span>
          <span class="bg-gray-50 text-gray-700 px-2 py-0.5 rounded">
            {{ course.total_lectures || 0 }} Lectures
          </span>
        </div>
        <p class="text-xs text-gray-700 mt-1 line-clamp-2">
          {{ course.short_description || truncateDescription(course.description) }}
        </p>
      </div>
    </div>

    <!-- Card Footer -->
    <div class="flex items-center justify-between px-3 pb-3">
      <span class="text-xs text-gray-500">
        By <span class="font-semibold text-blue-700">{{ course.instructor?.username || 'N/A' }}</span>
      </span>
      <span :class="course.is_free ? 'text-green-600 font-bold text-sm' : 'text-blue-800 font-bold text-sm'">
        {{ course.is_free ? 'Free' : formatPrice(course.price) }}
      </span>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import StarRating from './StarRating.vue';
import type { Course } from '../types/api.types';

export default defineComponent({
  name: 'CourseCard',
  components: { StarRating },
  props: {
    course: {
      type: Object as PropType<Course>,
      required: true,
    },
    featured: {
      type: Boolean,
      default: false,
    },
  },
  emits: ['click'],
  setup() {
    const truncateDescription = (desc: string) => {
      if (!desc) return 'No description available.';
      const words = desc.split(' ').slice(0, 20).join(' ');
      return desc.split(' ').length > 20 ? words + '...' : words;
    };

    const formatPrice = (price: number | string) => {
      const num = typeof price === 'number' ? price : parseFloat(price as string);
      return `₹${num.toFixed(2)}`;
    };

    return {
      truncateDescription,
      formatPrice,
    };
  },
});
</script>
