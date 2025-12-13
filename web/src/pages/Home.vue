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
          @click="scrollToCourses"
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

    <LoadingSpinner v-if="courseStore.loading" message="Loading courses..." />

    <AlertMessage
      v-else-if="courseStore.error"
      type="error"
      :message="courseStore.error"
      details="Please check your internet connection or try again later."
    />

    <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 xl:grid-cols-4 gap-6 md:gap-8">
      <CourseCard
        v-for="(course, idx) in courseStore.courses"
        :key="course.slug"
        :course="course"
        :featured="idx === 0"
        @click="router.push(`/course/${course.slug}`)"
      />

      <p v-if="courseStore.courses.length === 0" class="text-gray-600 col-span-full text-center py-8 md:py-12 text-lg md:text-xl italic">
        No courses available at the moment. Please check back soon!
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useCourseStore } from '../stores/courseStore';
import CourseCard from '../components/CourseCard.vue';
import LoadingSpinner from '../components/LoadingSpinner.vue';
import AlertMessage from '../components/AlertMessage.vue';

const router = useRouter();
const courseStore = useCourseStore();

onMounted(async () => {
  await courseStore.fetchCourses();
});

function scrollToCourses() {
  const coursesSection = document.querySelector('.grid');
  coursesSection?.scrollIntoView({ behavior: 'smooth' });
}
</script>
