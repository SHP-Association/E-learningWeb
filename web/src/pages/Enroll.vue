<template>
  <div class="min-h-screen flex items-start justify-center bg-gray-100 p-4">
    <p v-if="!user" class="text-center text-xl mt-10">Please log in to enroll in courses.</p>
    <p v-else-if="!course" class="text-center text-xl mt-10">Course not found.</p>
    <div v-else class="bg-white p-8 rounded-lg shadow-xl border border-gray-200 w-full max-w-md mt-10">
      <h1 class="text-3xl font-bold mb-4 text-blue-800">Enroll in {{ course.title }}</h1>
      <p class="text-gray-700 mb-4">
        You are about to enroll in <strong class="font-semibold">{{ course.title }}</strong>.
      </p>
      <p class="text-gray-700 mb-6">
        <strong class="font-semibold">Price:</strong>
        <span class="ml-1">{{ course.is_free ? 'Free' : `₹${course.price}` }}</span>
      </p>

      <form @submit.prevent="handleConfirmEnrollment">
        <div class="flex justify-between space-x-4">
          <button type="submit" class="bg-blue-900 text-white hover:bg-blue-700 px-6 py-3 rounded-lg font-semibold shadow transition flex-1">
            Confirm Enrollment
          </button>
          <button
            type="button"
            @click="navigate(`/course/${course.slug}`)"
            class="bg-gray-300 text-gray-800 hover:bg-gray-400 px-6 py-3 rounded-lg font-semibold transition flex-1"
          >
            Cancel
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject } from 'vue';

// Define placeholder types
interface User {
  id: number;
  username: string;
}
interface Course {
  id: number;
  title: string;
  slug: string;
  is_free: boolean;
  price?: number;
}
type AddEnrollmentFn = (user: User, course: Course) => boolean;

const props = defineProps<{
  course: Course | null;
  user: User | null;
  addEnrollment: AddEnrollmentFn;
  navigate: (path: string) => void;
}>();

const handleConfirmEnrollment = () => {
  if (!props.user || !props.course) return;

  const enrolled = props.addEnrollment(props.user, props.course);

  if (enrolled) {
    alert(`You have successfully enrolled in ${props.course.title}!`);
    props.navigate(`/course/${props.course.slug}`);
  } else {
    alert(`You are already enrolled in ${props.course.title}.`);
    props.navigate(`/course/${props.course.slug}`);
  }
};
</script>