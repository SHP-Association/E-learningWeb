<template>
  <div class="min-h-screen bg-gray-50 p-4 sm:p-8">
    <div v-if="!userStore.user" class="min-h-screen flex items-center justify-center bg-gray-50 p-4">
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
              :src="userStore.user.profile_picture || '/placeholder.jpg'"
              alt="Profile"
              class="w-32 h-32 rounded-full object-cover border-4 border-white shadow-lg"
            />
            <span :class="['absolute bottom-0 right-0 h-4 w-4 rounded-full ring-2 ring-white', userStore.user.is_active ? 'bg-green-400' : 'bg-red-400']"></span>
          </div>

          <div class="text-center sm:text-left">
            <h2 class="text-3xl font-bold text-blue-800">{{ userStore.user.first_name || userStore.user.username }} {{ userStore.user.last_name }}</h2>
            <p class="text-lg text-gray-600 italic">{{ userStore.user.role === 'instructor' ? 'Instructor' : 'Student' }}</p>
            <p class="text-gray-500 text-sm mt-1">Joined: {{ formatDate(userStore.user.date_joined) }}</p>
            
            <button 
              @click="editing = !editing" 
              :class="['mt-3 px-4 py-2 rounded-lg font-semibold transition', editing ? 'bg-gray-300 text-gray-800 hover:bg-gray-400' : 'bg-blue-600 text-white hover:bg-blue-700 shadow-md']"
            >
              {{ editing ? 'Cancel Edit' : 'Edit Profile' }}
            </button>
          </div>
        </div>

        <form @submit.prevent="handleEditSubmit" class="space-y-6">
          <!-- Profile Picture Section -->
          <section v-if="editing" class="pb-6 border-b border-gray-100">
            <h3 class="text-xl font-bold text-gray-800 mb-4">Profile Picture</h3>
            <div class="flex items-center space-x-6">
              <div class="relative">
                <img
                  :src="profilePicturePreview || userStore.user?.profile_picture || '/placeholder.jpg'"
                  alt="Profile Preview"
                  class="w-32 h-32 rounded-full object-cover border-4 border-blue-500 shadow-lg"
                />
              </div>
              <div class="flex-1">
                <label class="block text-sm font-medium text-gray-700 mb-2">Upload New Picture</label>
                <input
                  type="file"
                  accept="image/*"
                  @change="handleFileChange"
                  class="w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100 cursor-pointer"
                />
                <p class="mt-2 text-xs text-gray-500">JPG, PNG or GIF (MAX. 5MB)</p>
              </div>
            </div>
          </section>

          <section>
            <h3 class="text-xl font-bold text-gray-800 mb-4 border-b pb-2">Personal Information</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div v-for="field in editableFields" :key="field.key">
                <div v-if="editing">
                  <label :for="field.key" class="block text-sm font-medium text-gray-700 mb-1">{{ field.label }}</label>
                  
                  <input
                    v-if="['text', 'date', 'url', 'email'].includes(field.type)"
                    :type="field.type"
                    :id="field.key"
                    v-model="form[field.key]"
                    class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white text-gray-800"
                  />
                  <textarea
                    v-else-if="field.type === 'textarea'"
                    :id="field.key"
                    v-model="form[field.key]"
                    class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white text-gray-800 resize-y h-24"
                  ></textarea>
                  <select
                    v-else-if="field.type === 'select' && field.options"
                    :id="field.key"
                    v-model="form[field.key]"
                    class="w-full p-3 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white text-gray-800"
                  >
                    <option value="" disabled>Select {{ field.label }}</option>
                    <option v-for="option in field.options" :key="option" :value="option">{{ option }}</option>
                  </select>
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

        <section v-if="userStore.isStudent && enrollmentStore.enrollments.length > 0" class="pt-8 border-t border-gray-100">
          <h3 class="text-2xl font-bold text-gray-800 mb-4">Enrolled Courses</h3>
          <ul class="space-y-4">
            <li
              v-for="enrollment in enrollmentStore.enrollments"
              :key="enrollment.id"
              class="p-4 bg-gray-50 rounded-lg border border-gray-200 hover:shadow-lg transition duration-200"
            >
              <div class="flex items-start justify-between">
                <router-link
                  :to="`/course/${enrollment.course.slug}`"
                  class="text-blue-700 hover:underline font-semibold text-lg"
                >
                  {{ enrollment.course.title }}
                </router-link>
                <span class="text-sm text-gray-500 ml-4">
                  Enrolled: {{ formatDate(enrollment.enrolled_at) }}
                </span>
              </div>
              <div class="flex items-center justify-between mt-2">
                <span class="bg-blue-100 text-blue-800 text-xs font-semibold px-3 py-1 rounded-full">
                  {{ enrollment.progress }}% Complete
                </span>
              </div>
            </li>
          </ul>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch, onMounted } from 'vue';
import { useUserStore } from '../stores/userStore';
import { useEnrollmentStore } from '../stores/enrollmentStore';

const userStore = useUserStore();
const enrollmentStore = useEnrollmentStore();

const editing = ref(false);
const isSubmitting = ref(false);
const form = reactive<Record<string, any>>({});
const profilePictureFile = ref<File | null>(null);
const profilePicturePreview = ref<string>('');

// Initialize form with user data
watch(() => userStore.user, (newUser) => {
  if (newUser) {
    Object.assign(form, newUser);
  }
}, { immediate: true });

// Fetch enrollments on mount
onMounted(() => {
  if (userStore.isStudent) {
    enrollmentStore.fetchEnrollments();
  }
});

// Define editable fields
const studentFields = [
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

const editableFields = computed(() => userStore.isStudent ? studentFields : instructorFields);

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files[0]) {
    const file = target.files[0];
    
    // Validate file size (5MB max)
    if (file.size > 5 * 1024 * 1024) {
      alert('File size must be less than 5MB');
      return;
    }
    
    // Validate file type
    if (!file.type.startsWith('image/')) {
      alert('Please select an image file');
      return;
    }
    
    profilePictureFile.value = file;
    
    // Create preview
    const reader = new FileReader();
    reader.onload = (e) => {
      profilePicturePreview.value = e.target?.result as string;
    };
    reader.readAsDataURL(file);
  }
};

const handleEditSubmit = async () => {
  if (!userStore.user) return;

  isSubmitting.value = true;
  
  try {
    const updates: Record<string, any> = {};
    editableFields.value.forEach(field => {
      if (form[field.key] !== undefined && form[field.key] !== null) {
        updates[field.key] = form[field.key];
      }
    });

    // Handle profile picture upload separately if file is selected
    if (profilePictureFile.value) {
      const formData = new FormData();
      formData.append('profile_picture', profilePictureFile.value);
      
      // Add other fields to formData
      Object.keys(updates).forEach(key => {
        formData.append(key, updates[key]);
      });
      
      // Use fetch for file upload
      const response = await fetch(`${import.meta.env.VITE_APP_BACKEND_URL}/api/users/${userStore.user.id}/`, {
        method: 'PATCH',
        credentials: 'include',
        body: formData,
      });
      
      if (response.ok) {
        const updatedUser = await response.json();
        userStore.saveUser(updatedUser);
        editing.value = false;
        profilePictureFile.value = null;
        profilePicturePreview.value = '';
      } else {
        const error = await response.json();
        alert(error.detail || 'Failed to update profile');
      }
    } else {
      // No file upload, use regular update
      const success = await userStore.updateProfile(updates);
      if (success) {
        editing.value = false;
      } else {
        alert(userStore.error || 'Failed to update profile');
      }
    }
  } catch (err: any) {
    alert(err.message || 'An error occurred while updating the profile');
  } finally {
    isSubmitting.value = false;
  }
};

const formatDate = (dateStr?: string) => {
  if (!dateStr) return 'N/A';
  try {
    return new Date(dateStr).toLocaleDateString();
  } catch {
    return dateStr;
  }
};

const displayValue = (key: string, value: any) => {
  if (value === null || value === undefined || value === '') return 'N/A';
  if (key === 'date_of_birth') return formatDate(value);
  if (typeof value === 'boolean') return value ? 'Yes' : 'No';
  return value;
};
</script>
