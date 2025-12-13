<template>
  <div v-if="!course" class="text-center text-xl mt-10 text-red-600">
    Course not found.
  </div>
  <div v-else class="bg-gradient-to-br from-blue-50 to-white p-0 sm:p-0 rounded-2xl shadow-2xl border border-gray-100 max-w-4xl mx-auto mt-10">
    <!-- Header -->
    <div class="flex flex-col md:flex-row gap-8 p-8 border-b border-gray-100 bg-white rounded-t-2xl">
      <!-- Thumbnail -->
      <div class="flex-shrink-0 w-full md:w-80 flex flex-col items-center">
        <img
          v-if="course.thumbnail"
          :src="course.thumbnail"
          :alt="course.title"
          class="rounded-xl border shadow max-w-full"
          style="max-height: 320px; object-fit: contain; background: #f3f4f6"
        />
        <img
          v-else-if="course.category?.image"
          :src="course.category.image"
          :alt="course.category.name"
          class="rounded-xl border shadow max-w-full"
          style="max-height: 320px; object-fit: contain; background: #f3f4f6"
        />
        <div
          v-else
          class="w-full h-44 md:h-72 flex items-center justify-center bg-gray-100 rounded-xl border text-gray-400"
        >
          No Image
        </div>
        <a
          v-if="course.promo_video_url"
          :href="course.promo_video_url"
          target="_blank"
          rel="noopener noreferrer"
          class="block mt-4 bg-blue-100 text-blue-700 px-4 py-2 rounded-lg font-medium hover:bg-blue-200 transition text-center"
        >
          ▶ Watch Promo Video
        </a>
      </div>

      <!-- Main Info -->
      <div class="flex-1 flex flex-col">
        <h1 class="text-3xl sm:text-4xl font-extrabold text-blue-800 mb-2">{{ course.title }}</h1>
        <p class="text-gray-700 mb-3 text-lg">{{ course.short_description || course.description }}</p>

        <!-- Tags -->
        <div class="flex flex-wrap gap-3 items-center mb-3">
          <span class="inline-flex items-center bg-blue-50 text-blue-800 px-3 py-1 rounded-full text-sm font-medium">
            <strong>Category:</strong> {{ course.category?.name }}
          </span>
          <span class="inline-flex items-center bg-green-50 text-green-800 px-3 py-1 rounded-full text-sm font-medium">
            <strong>Level:</strong> {{ course.level }}
          </span>
          <span class="inline-flex items-center bg-yellow-50 text-yellow-800 px-3 py-1 rounded-full text-sm font-medium">
            <strong>Lectures:</strong> {{ course.total_lectures }}
          </span>
          <span class="inline-flex items-center bg-purple-50 text-purple-800 px-3 py-1 rounded-full text-sm font-medium">
            <strong>Rating:</strong> {{ (course.average_rating || 0).toFixed(2) }}
          </span>
          <span class="inline-flex items-center bg-gray-50 text-gray-800 px-3 py-1 rounded-full text-sm font-medium">
            <strong>Reviews:</strong> {{ course.number_of_reviews }}
          </span>
          <span class="inline-flex items-center bg-pink-50 text-pink-800 px-3 py-1 rounded-full text-sm font-medium">
            <strong>Price:</strong> {{ course.is_free ? 'Free of Cost' : `₹${course.price}` }}
          </span>
        </div>

        <!-- Instructor Info -->
        <div class="flex items-center gap-3 mb-3">
          <span class="font-semibold text-gray-800">Instructor:</span>
          <span class="flex items-center gap-2">
            <img
              v-if="course.instructor?.profile_picture"
              :src="course.instructor.profile_picture"
              :alt="course.instructor.username"
              class="w-8 h-8 rounded-full object-cover border"
            />
            <span v-else class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center text-blue-700 font-bold text-lg">
              {{ course.instructor?.username?.[0]?.toUpperCase() || '?' }}
            </span>
            <span class="font-medium">{{ course.instructor?.username }}</span>
            <span
              v-if="course.instructor?.instructor_rating > 0"
              class="ml-2 text-yellow-600 text-xs bg-yellow-100 px-2 py-0.5 rounded-full"
            >
              ⭐ {{ course.instructor.instructor_rating.toFixed(1) }}
            </span>
          </span>
        </div>

        <!-- Published/Duration -->
        <div class="flex flex-wrap gap-4 mb-2">
          <span class="text-gray-500 text-sm">
            <strong>Published:</strong> {{ formatDate(course.created_at) }}
          </span>
          <span v-if="course.duration" class="text-gray-700 text-sm">
            <strong>Duration:</strong> {{ course.duration }}
          </span>
        </div>
      </div>
    </div>

    <!-- Description Section -->
    <div class="px-8 py-6 border-b border-gray-100 bg-white">
      <h2 class="text-xl font-semibold text-blue-900 mb-2">Course Overview</h2>
      <ul v-if="course.description" class="list-disc pl-6 text-gray-800 text-base leading-relaxed">
        <li v-for="(line, idx) in formattedDescription" :key="idx">{{ line }}</li>
      </ul>
      <p v-else class="text-gray-600">No detailed description available.</p>
    </div>

    <!-- What you'll learn / Requirements / Target Audience -->
    <div v-if="course.what_you_will_learn || course.requirements || course.target_audience" class="grid grid-cols-1 md:grid-cols-3 gap-6 px-8 py-6 border-b border-gray-100 bg-white">
      <section v-if="course.what_you_will_learn" class="bg-blue-50 rounded-xl p-4">
        <h2 class="text-lg font-semibold text-blue-900 mb-1">What you'll learn</h2>
        <p class="text-gray-700 text-sm">{{ course.what_you_will_learn }}</p>
      </section>
      <section v-if="course.requirements" class="bg-green-50 rounded-xl p-4">
        <h2 class="text-lg font-semibold text-green-900 mb-1">Requirements</h2>
        <p class="text-gray-700 text-sm">{{ course.requirements }}</p>
      </section>
      <section v-if="course.target_audience" class="bg-yellow-50 rounded-xl p-4">
        <h2 class="text-lg font-semibold text-yellow-900 mb-1">Target Audience</h2>
        <p class="text-gray-700 text-sm">{{ course.target_audience }}</p>
      </section>
    </div>

    <!-- Enrollment/Progress Section -->
    <div class="px-8 py-6 border-b border-gray-100 bg-white">
      <template v-if="user">
        <div v-if="isEnrolled" class="bg-green-50 border border-green-300 p-4 rounded-xl flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
          <div>
            <p class="text-green-700 font-semibold flex items-center gap-2">🎉 You are enrolled in this course!</p>
            <p v-if="enrollment?.enrolled_at" class="text-sm text-gray-700 mt-1">
              <strong>Enrolled At:</strong> {{ formatDateTime(enrollment.enrolled_at) }}
            </p>
            <p class="text-sm text-gray-700"><strong>Progress:</strong> {{ enrollment?.progress ?? 0 }}%</p>
          </div>
          <button disabled class="bg-green-600 text-white px-6 py-2 rounded-lg opacity-80 cursor-not-allowed font-semibold">
            Enrolled
          </button>
        </div>
        <button v-else @click="navigate(`/enroll/${course.slug}`)" class="bg-blue-700 hover:bg-blue-600 text-white font-semibold px-8 py-3 rounded-lg transition duration-300 shadow">
          Enroll Now
        </button>
      </template>
      <p v-else class="text-gray-700">
        Please
        <button @click="navigate('/login')" class="text-blue-600 font-medium hover:underline">log in</button>
        to enroll in this course.
      </p>
    </div>

    <!-- Lessons Section -->
    <div class="px-8 py-6 border-b border-gray-100 bg-white">
      <h2 class="text-2xl font-semibold text-gray-800 mb-3">Lessons</h2>
      <div v-if="isEnrolled">
        <div v-if="course.lessons?.length" class="space-y-3">
          <div v-for="lesson in course.lessons" :key="lesson.order" class="bg-gray-50 border border-gray-200 p-4 rounded-lg flex items-center justify-between">
            <p class="font-medium text-gray-900">Lesson {{ lesson.order }}: {{ lesson.title }}</p>
            <a v-if="lesson.video_url" :href="lesson.video_url" target="_blank" rel="noopener noreferrer" class="text-blue-600 hover:underline font-semibold">▶ Watch</a>
          </div>
        </div>
        <p v-else class="text-gray-600">No lessons available yet.</p>
      </div>
      <div v-else class="mt-4 bg-yellow-50 border-l-4 border-yellow-400 text-yellow-700 p-4 rounded">
        You need to enroll in this course to access the lessons.
      </div>
    </div>

    <!-- More Details Section -->
    <div class="px-8 py-8 grid grid-cols-1 md:grid-cols-2 gap-6 bg-white rounded-b-2xl">
      <div class="bg-gray-50 rounded-xl p-5 border">
        <h3 class="font-semibold text-gray-800 mb-2">Course Details</h3>
        <ul class="text-gray-700 text-sm space-y-1">
          <li><strong>Course ID:</strong> {{ course.id }}</li>
          <li><strong>Slug:</strong> {{ course.slug }}</li>
          <li><strong>Published:</strong> {{ course.is_published ? 'Yes' : 'No' }}</li>
          <li><strong>Created:</strong> {{ formatDateTime(course.created_at) }}</li>
          <li><strong>Updated:</strong> {{ formatDateTime(course.updated_at) }}</li>
        </ul>
      </div>
      <div class="bg-gray-50 rounded-xl p-5 border">
        <h3 class="font-semibold text-gray-800 mb-2">Instructor Details</h3>
        <ul class="text-gray-700 text-sm space-y-1">
          <li><strong>Username:</strong> {{ course.instructor?.username }}</li>
          <li><strong>Email:</strong> {{ course.instructor?.email || 'N/A' }}</li>
          <li><strong>Bio:</strong> {{ course.instructor?.bio || 'N/A' }}</li>
          <li><strong>Joined:</strong> {{ formatDate(course.instructor?.date_joined) }}</li>
          <li><strong>Role:</strong> {{ course.instructor?.role }}</li>
          <li><strong>Rating:</strong> {{ course.instructor?.instructor_rating ?? 'N/A' }}</li>
          <li><strong>Total Reviews:</strong> {{ course.instructor?.total_reviews ?? 0 }}</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useCourseStore } from '../stores/courseStore';
import { useEnrollmentStore } from '../stores/enrollmentStore';
import { useUserStore } from '../stores/userStore';
import { useCourseSEO } from '../composables/useSEO';

export default defineComponent({
  name: 'CourseDetail',
  setup() {
    const route = useRoute();
    const router = useRouter();
    const courseStore = useCourseStore();
    const enrollmentStore = useEnrollmentStore();
    const userStore = useUserStore();

    const slug = route.params.slug as string;

    onMounted(async () => {
      await courseStore.fetchCourseBySlug(slug);
      if (userStore.isLoggedIn) {
        await enrollmentStore.fetchEnrollments();
      }
    });

    const course = computed(() => courseStore.currentCourse);
    const isEnrolled = computed(() => enrollmentStore.isEnrolled(slug));
    const enrollment = computed(() => enrollmentStore.getEnrollment(slug));
    const user = computed(() => userStore.user);

    // Set SEO when course data is available
    watch(course, (newCourse) => {
      if (newCourse) {
        useCourseSEO({
          title: newCourse.title,
          description: newCourse.short_description || newCourse.description,
          instructor: newCourse.instructor?.username || 'Unknown',
          category: newCourse.category?.name || 'Uncategorized',
          level: newCourse.level,
          price: newCourse.price,
          isFree: newCourse.is_free,
          rating: newCourse.average_rating || 0,
          reviewCount: newCourse.number_of_reviews || 0,
          image: newCourse.thumbnail,
        });
      }
    }, { immediate: true });

    const formattedDescription = computed(() => {
      if (!course.value?.description) return [];
      return course.value.description
        .split(/\r?\n|• /)
        .filter(Boolean)
        .map((line: string) => line.replace(/^[-•\s]+/, ''));
    });

    const formatDate = (dateStr?: string) => {
      if (!dateStr) return 'N/A';
      return new Date(dateStr).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' });
    };

    const formatDateTime = (dateStr?: string) => {
      if (!dateStr) return 'N/A';
      return new Date(dateStr).toLocaleString('en-US', { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' });
    };

    const navigate = (path: string) => router.push(path);

    return { 
      course, 
      isEnrolled, 
      enrollment,
      user, 
      formattedDescription, 
      formatDate, 
      formatDateTime,
      navigate,
    };
  },
});
</script>


