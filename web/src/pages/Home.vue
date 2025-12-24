<template>
  <div class="container mx-auto px-4 py-8 md:py-12">
    <!-- Hero Section -->
    <section
      class="hero relative bg-gradient-to-r from-blue-600 to-indigo-700 text-white py-16 md:py-24 lg:py-32 text-center rounded-3xl shadow-2xl mb-12 md:mb-16 overflow-hidden"
    >
      <!-- Animated Background Circles -->
      <div class="absolute inset-0 pointer-events-none z-0">
        <div
          class="absolute top-[-60px] left-[-60px] w-72 h-72 bg-blue-400 opacity-20 rounded-full animate-pulse-slow"
        ></div>
        <div
          class="absolute bottom-[-80px] right-[-80px] w-96 h-96 bg-indigo-400 opacity-20 rounded-full animate-pulse-slower"
        ></div>
      </div>
      <div class="absolute inset-0 bg-black opacity-10 rounded-3xl z-0"></div>
      <div class="relative z-10">
        <h1
          class="text-4xl sm:text-5xl md:text-6xl lg:text-7xl font-extrabold mb-4 md:mb-6 leading-tight drop-shadow-lg animate-fade-in-down"
        >
          Welcome to <span class="text-yellow-300">SHP-Learner</span>
        </h1>
        <p
          class="text-lg sm:text-xl md:text-2xl lg:text-3xl mb-8 md:mb-10 opacity-90 max-w-xs sm:max-w-md md:max-w-2xl lg:max-w-3xl mx-auto px-4 animate-fade-in"
        >
          Explore our wide range of meticulously crafted courses to elevate your skills and
          career.
        </p>
        <button
          @click="navigate('/courses')"
          class="btn bg-white text-blue-700 hover:bg-blue-100 hover:text-blue-800 px-8 py-3 md:px-10 md:py-4 text-lg md:text-xl font-semibold rounded-full shadow-xl transform hover:-translate-y-1 transition duration-300 ease-in-out border-2 border-transparent hover:border-blue-700 animate-scale-in"
        >
          Browse All Courses
        </button>
      </div>
    </section>

    <!-- Section Divider -->
    <div class="flex items-center justify-center my-12">
      <span class="h-1 w-24 bg-blue-200 rounded-full"></span>
      <span class="mx-4 text-2xl font-bold text-blue-700 tracking-wide">Courses</span>
      <span class="h-1 w-24 bg-blue-200 rounded-full"></span>
    </div>

    <!-- Courses List -->
    <h2
      class="text-3xl sm:text-4xl lg:text-5xl font-extrabold mb-8 md:mb-12 text-gray-800 text-center relative"
    >
      <span class="relative inline-block pb-2">
        Discover Our Popular Courses
        <span
          class="absolute bottom-0 left-1/2 w-20 h-1 bg-blue-500 transform -translate-x-1/2 rounded-full"
        ></span>
      </span>
    </h2>

    <div v-if="loading" class="flex justify-center items-center py-16">
      <div class="loader ease-linear rounded-full border-8 border-t-8 border-gray-200 h-24 w-24"></div>
      <p class="text-center text-xl text-blue-600 ml-4">Loading courses...</p>
    </div>

    <div
      v-else-if="error"
      class="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded-lg text-center mx-auto max-w-md mt-8 shadow-md"
    >
      <p class="text-lg">{{ error }}</p>
      <p class="text-sm mt-2">Please check your internet connection or try again later.</p>
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 xl:grid-cols-4 gap-6 md:gap-8">
      <div
        v-for="(course, idx) in courses"
        :key="course.slug"
        class="yt-card group bg-white rounded-xl shadow-md hover:shadow-2xl border border-gray-100 transition-all duration-300 flex flex-col cursor-pointer overflow-hidden"
        @click="navigate(`/course/${course.slug}`)"
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
              <path d="M12 2C6.48 2 2 6.48...z" />
            </svg>
          </div>

          <!-- Badges -->
          <span
            v-if="idx === 0"
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
              <div class="flex items-center">
                <svg
                  v-for="i in 5"
                  :key="i"
                  class="w-4 h-4"
                  :class="i <= Math.round(course.average_rating || 0) ? 'text-yellow-400' : 'text-gray-300'"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.538 1.118l-2.8-2.034a1 1 0 00-1.176 0l-2.8 2.034c-.783.57-1.838-.197-1.538-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.381-1.81.588-1.81h3.462a1 1 0 00.95-.69l1.07-3.292z"
                  />
                </svg>
              </div>
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

      <p v-if="courses.length === 0" class="text-gray-600 col-span-full text-center py-8 md:py-12 text-lg md:text-xl italic">
        No courses available at the moment. Please check back soon!
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';

const BACKEND_URL = import.meta.env.VITE_APP_BACKEND_URL;

export default defineComponent({
  name: 'Home',
  props: {
    courses: {
      type: Array,
      default: () => [],
    },
    navigate: {
      type: Function,
      required: true,
    },
  },
  setup(props) {
    const courses = ref([...props.courses]);
    const loading = ref(false);
    const error = ref('');

    onMounted(async () => {
      loading.value = true;
      try {
        const res = await fetch(`${BACKEND_URL}/api/courses/`);
        if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`);
        courses.value = await res.json();
      } catch (err) {
        console.error('Error fetching courses:', err);
        error.value = 'Failed to load courses. Please try again later.';
      } finally {
        loading.value = false;
      }
    });

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
      courses,
      loading,
      error,
      truncateDescription,
      formatPrice,
      navigate: props.navigate,
    };
  },
});
</script>
