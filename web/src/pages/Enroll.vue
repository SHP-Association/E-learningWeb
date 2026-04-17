<template>
  <div class="min-h-screen flex items-start justify-center bg-gray-100 p-4">
    <p v-if="!user" class="text-center text-xl mt-10">Please log in to enroll in courses.</p>
    <p v-else-if="!course" class="text-center text-xl mt-10">Course not found.</p>
    <AppCard v-else size="md" padding="lg" elevated centered class="mt-10">
      <h1 class="mb-4 text-3xl font-bold text-blue-800">Enroll in {{ course.title }}</h1>
      <p class="mb-4 text-gray-700">
        You are about to enroll in <strong class="font-semibold">{{ course.title }}</strong>.
      </p>
      <p class="mb-6 text-gray-700">
        <strong class="font-semibold">Price:</strong>
        <span class="ml-1">{{ course.is_free ? 'Free' : `₹${course.price}` }}</span>
      </p>

      <form @submit.prevent="handleConfirmEnrollment">
        <div class="flex justify-between space-x-4">
          <AppButton type="submit" tone="primary" class="flex-1">
            Confirm Enrollment
          </AppButton>
          <AppButton
            type="button"
            variant="outline"
            tone="primary"
            class="flex-1"
            @click="navigate(`/course/${course.slug}`)"
          >
            Cancel
          </AppButton>
        </div>
      </form>
    </AppCard>
  </div>
</template>

<script setup lang="ts">
import AppButton from '../components/ui/AppButton.vue';
import AppCard from '../components/ui/AppCard.vue';

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
