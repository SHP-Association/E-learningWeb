<template>
  <div class="min-h-screen bg-gray-50 p-4 sm:p-8">
    <div v-if="!currentUser" class="min-h-screen flex items-center justify-center bg-gray-50 p-4">
      <p class="text-center text-xl text-gray-700 bg-white p-8 rounded-lg shadow-lg">
        Please log in to view your profile.
      </p>
    </div>

    <div v-else class="max-w-5xl mx-auto">
      <h1 class="text-4xl font-extrabold text-gray-900 mb-8 border-b pb-4">
        {{ editing ? 'Edit Your Profile' : 'My Profile' }}
      </h1>

      <div class="bg-white rounded-xl shadow-2xl border border-gray-100 p-6 sm:p-8 space-y-8">
        <div class="flex flex-col sm:flex-row items-center sm:items-start space-y-4 sm:space-y-0 sm:space-x-8 pb-6 border-b border-gray-100">
          <div class="relative">
            <img
              :src="currentUser.profile_picture || 'placeholder.jpg'"
              alt="Profile"
              class="w-32 h-32 rounded-full object-cover border-4 border-white shadow-lg"
            />
            <span :class="['absolute bottom-0 right-0 h-4 w-4 rounded-full ring-2 ring-white', currentUser.is_active ? 'bg-green-400' : 'bg-red-400']" :title="currentUser.is_active ? 'Active' : 'Inactive'"></span>
          </div>

          <div class="text-center sm:text-left">
            <h2 class="text-3xl font-bold text-blue-800">{{ currentUser.first_name || currentUser.username }} {{ currentUser.last_name }}</h2>
            <p class="text-lg text-gray-600 italic">{{ currentUser.role === 'instructor' ? 'Instructor' : 'Student' }}</p>
            <p class="text-gray-500 text-sm mt-1">Joined: {{ formatDate(currentUser.date_joined) }}</p>
            <p v-if="currentUser.last_activity" class="text-gray-500 text-sm">Last Activity: {{ formatDateTime(currentUser.last_activity) }}</p>
            
            <button 
              @click="setEditing(!editing)" 
              :class="['mt-3 px-4 py-2 rounded-lg font-semibold transition', editing ? 'bg-gray-300 text-gray-800 hover:bg-gray-400' : 'bg-blue-600 text-white hover:bg-blue-700 shadow-md']"
            >
              {{ editing ? 'Cancel Edit' : 'Edit Profile' }}
            </button>
          </div>
        </div>

        <form @submit.prevent="handleEditSubmit" class="space-y-6">
          <section>
            <h3 class="text-xl font-bold text-gray-800 mb-4 border-b pb-2">Personal Information</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div v-for="field in editableFields" :key="field.key">
                <div v-if="editing">
                  <label :for="field.key" class="form-label">{{ field.label }}</label>
                  
                  <input
                    v-if="['text', 'date', 'url', 'email', 'number'].includes(field.type)"
                    :type="field.type"
                    :id="field.key"
                    :name="field.key"
                    :value="form[field.key]"
                    @input="handleEditChange"
                    class="form-input"
                  />
                  <textarea
                    v-else-if="field.type === 'textarea'"
                    :id="field.key"
                    :name="field.key"
                    :value="form[field.key]"
                    @input="handleEditChange"
                    class="form-input resize-y h-24"
                  ></textarea>
                  <select
                    v-else-if="field.type === 'select' && field.options"
                    :id="field.key"
                    :name="field.key"
                    :value="form[field.key]"
                    @change="handleEditChange"
                    class="form-input"
                  >
                    <option value="" disabled>Select {{ field.label }}</option>
                    <option v-for="option in field.options" :key="option" :value="option">{{ option }}</option>
                  </select>
                  <input
                    v-else-if="field.type === 'file'"
                    type="file"
                    :id="field.key"
                    :name="field.key"
                    @change="handleFileChange"
                    class="w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100"
                  />
                </div>
                <div v-else>
                  <div class="flex items-start mb-3">
                    <span class="min-w-[120px] font-semibold text-gray-700">{{ field.label }}:</span>
                    <span class="text-gray-600 break-words ml-2">{{ displayValue(field.key, form[field.key]) }}</span>
                  </div>
                </div>
              </div>
            </div>
          </section>

          <div v-if="editing" class="pt-4 border-t border-gray-100 flex justify-end">
            <button
              type="submit"
              class="bg-green-600 text-white hover:bg-green-700 px-6 py-3 rounded-lg font-semibold shadow-md transition"
              :disabled="isSubmitting"
            >
              {{ isSubmitting ? 'Saving...' : 'Save Changes' }}
            </button>
          </div>
        </form>

        <section v-if="isStudent && enrollments && enrollments.length" class="pt-8 border-t border-gray-100">
          <h3 class="text-2xl font-bold text-gray-800 mb-4">Enrolled Courses</h3>
          <ul class="space-y-4">
            <li
              v-for="enrollment in enrollments"
              :key="enrollment.id"
              class="p-4 bg-gray-50 rounded-lg border border-gray-200 hover:shadow-lg transition duration-200 flex flex-col"
            >
              <div class="flex items-start justify-between">
                <button
                  @click="navigate(`/course/${enrollment.course.slug}`)"
                  :aria-label="`Go to ${enrollment.course.title} course details`"
                  class="text-blue-700 hover:underline font-semibold text-lg text-left"
                >
                  {{ enrollment.course.title }}
                </button>
                <span class="text-sm text-gray-500 ml-4 flex-shrink-0">
                  Enrolled: {{ formatDate(enrollment.enrolled_at) }}
                </span>
              </div>
              <div class="flex items-center justify-between mt-2">
                <span
                  class="progress-badge bg-blue-100 text-blue-800 text-xs font-semibold px-3 py-1 rounded-full"
                  :aria-label="`Progress: ${enrollment.progress} percent complete`"
                >
                  {{ enrollment.progress }}% Complete
                </span>
              </div>
            </li>
          </ul>
        </section>
        <section v-else-if="isStudent && enrollments && enrollments.length === 0" class="pt-8 border-t border-gray-100">
          <h3 class="text-2xl font-bold text-gray-800 mb-4">Enrolled Courses</h3>
          <li class="p-4 bg-white rounded-lg border border-gray-100 text-gray-600 text-base sm:text-lg italic text-center shadow">
            You are not enrolled in any courses yet.
          </li>
        </section>

        <div class="flex justify-center py-8">
          <button
            @click="navigate('/password_reset')"
            class="bg-yellow-500 text-white hover:bg-yellow-600 px-8 py-3 rounded-lg font-semibold text-lg shadow transition duration-200 focus:outline-none focus:ring-2 focus:ring-yellow-500 focus:ring-opacity-50"
          >
            Reset Password
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch, inject } from 'vue';

// Define complex types for clarity
interface BaseUser {
  id: number;
  username: string;
  email: string;
  role: 'student' | 'instructor' | string;
  is_active: boolean;
  date_joined: string;
  last_activity?: string;
  profile_picture?: string | File;
  first_name?: string;
  last_name?: string;
  bio?: string;
  institution?: string;
  skills?: string;
  linkedin_profile?: string;
  github_profile?: string;
  website?: string;
}

interface StudentUser extends BaseUser {
  date_of_birth?: string;
  gender?: 'male' | 'female' | 'other' | string;
  contact_number?: string;
  address?: string;
  country?: string;
  highest_qualification?: string;
}

interface Course {
  slug: string;
  title: string;
}

interface Enrollment {
  id: number;
  course: Course;
  progress: number;
  enrolled_at: string;
}

// Props
const props = defineProps<{
  user: BaseUser | null;
  enrollments: Enrollment[];
  navigate: (path: string) => void;
}>();

// Injected Context (Assuming App.vue provides this)
interface InjectedContext {
  user: BaseUser | null;
  handleLogin: (user: BaseUser) => void;
  // If navigate is also in context, you'd pull it here:
  // navigate: (path: string) => void;
}
// We will rely on props.navigate and props.user/enrollments but inject handleLogin from context
const { handleLogin, user: contextUser } = inject('user-context', {} as InjectedContext);

const BACKEND_URL = import.meta.env.VITE_APP_BACKEND_URL;

// State
const editing = ref(false);
const isSubmitting = ref(false);
const form = reactive<Record<string, any>>({});

// Determine the active user (prefer prop over injected context if both are present)
const currentUser = computed(() => props.user || contextUser.value);
const isStudent = computed(() => currentUser.value?.role === 'student');

// Initialize form state when currentUser changes
watch(currentUser, (newUser) => {
  if (newUser) {
    // Merge new user data into the reactive form object
    Object.assign(form, newUser);
  }
}, { immediate: true });

// Define editable fields based on user role
const studentFields = [
  { key: 'profile_picture', label: 'Profile Picture', type: 'file' },
  { key: 'first_name', label: 'First Name', type: 'text' },
  { key: 'last_name', label: 'Last Name', type: 'text' },
  { key: 'bio', label: 'Bio', type: 'textarea' },
  { key: 'date_of_birth', label: 'Date of Birth', type: 'date' },
  { key: 'gender', label: 'Gender', type: 'select', options: ['male', 'female', 'other'] },
  { key: 'contact_number', label: 'Contact Number', type: 'text' },
  { key: 'address', label: 'Address', type: 'textarea' },
  { key: 'country', label: 'Country', type: 'text' },
  { key: 'highest_qualification', label: 'Highest Qualification', type: 'text' },
  { key: 'institution', label: 'Institution', type: 'text' },
  { key: 'skills', label: 'Skills', type: 'text' },
  { key: 'linkedin_profile', label: 'LinkedIn', type: 'url' },
  { key: 'github_profile', label: 'GitHub', type: 'url' },
  { key: 'website', label: 'Website', type: 'url' },
];

const instructorFields = [
  { key: 'first_name', label: 'First Name', type: 'text' },
  { key: 'last_name', label: 'Last Name', type: 'text' },
  { key: 'bio', label: 'Bio', type: 'textarea' },
  { key: 'institution', label: 'Institution', type: 'text' },
  { key: 'skills', label: 'Skills', type: 'text' },
  { key: 'linkedin_profile', label: 'LinkedIn', type: 'url' },
  { key: 'github_profile', label: 'GitHub', type: 'url' },
  { key: 'website', label: 'Website', type: 'url' },
];

const editableFields = computed(() => isStudent.value ? studentFields : instructorFields);

// Event Handlers
const handleEditChange = (e: Event) => {
  const target = e.target as HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement;
  const { name, value } = target;
  form[name] = value;
};

const handleFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement;
  if (target.files && target.files[0]) {
    form.profile_picture = target.files[0];
  }
};

const handleEditSubmit = async () => {
  if (!currentUser.value) return;

  isSubmitting.value = true;
  
  try {
    // 1. Prepare FormData for file upload and PATCH request
    const body = new FormData();
    editableFields.value.forEach(field => {
      // Logic for profile picture file upload
      if (field.key === 'profile_picture') {
        // Only append if a new File object exists in the form state
        if (form.profile_picture instanceof File) {
          body.append('profile_picture', form.profile_picture);
        }
      } else if (form[field.key] !== undefined && form[field.key] !== null) {
        // Append all other fields
        body.append(field.key, form[field.key] === null ? '' : form[field.key]);
      }
    });

    // Get CSRF token from cookie (Preserved Django/React auth pattern)
    const csrfToken = (document.cookie.match(/csrftoken=([^;]+)/) || [])[1] || '';
    if (!csrfToken) {
      alert('CSRF token not found. Please refresh the page and try again.');
      isSubmitting.value = false;
      return;
    }

    const headers: HeadersInit = {
      'X-CSRFToken': csrfToken,
      // NOTE: Do NOT set Content-Type: 'application/json' when using FormData
    };

    // 2. Perform PATCH request
    const res = await fetch(`${BACKEND_URL}/api/users/${currentUser.value.id}/`, {
      method: 'PATCH',
      headers,
      credentials: 'include',
      body,
    });

    if (res.ok) {
      const updatedUser = await res.json();
      // Update global user state (assuming App.vue's handleLogin updates it)
      handleLogin(updatedUser);
      editing.value = false;
    } else {
      let msg = 'Failed to update profile.';
      try {
        const err = await res.json();
        // Check for common error structures
        if (err.detail) msg = err.detail;
        // Handle field-specific errors if available in the error object
        const fieldErrors = Object.values(err).flat();
        if (fieldErrors.length > 0 && typeof fieldErrors[0] === 'string') {
          msg = `Validation Error: ${fieldErrors.join('; ')}`;
        }
      } catch {}
      alert(msg);
    }
  } catch {
    alert('An unexpected error occurred while updating the profile.');
  } finally {
    isSubmitting.value = false;
  }
};

// Helper functions for display
const formatDate = (dateStr?: string) => {
  if (!dateStr) return 'N/A';
  try {
    return new Date(dateStr).toLocaleDateString();
  } catch {
    return dateStr;
  }
};

const formatDateTime = (dateStr?: string) => {
  if (!dateStr) return 'N/A';
  try {
    return new Date(dateStr).toLocaleString();
  } catch {
    return dateStr;
  }
};

const displayValue = (key: string, value: any) => {
  if (value === null || value === undefined || value === '') return 'N/A';
  
  if (key === 'date_of_birth') return formatDate(value);
  if (typeof value === 'boolean') return value ? 'Yes' : 'No';
  
  if (key.includes('profile')) return value.name || 'File Selected';

  return value;
};
</script>

<style scoped>
/* Simplified input style for Vue conversion */
.form-input {
  @apply w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white text-gray-800 transition duration-150;
}
.form-label {
  @apply block text-sm font-medium text-gray-700 mb-1;
}
</style>