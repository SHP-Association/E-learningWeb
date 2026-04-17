<template>
  <AuthShell
    title="Join SHP-Learner!"
    subtitle="Create your account and start learning today."
    tone="secondary"
  >
    <template #status>
      <div v-if="hasErrors" class="space-y-3">
        <AlertMessage
          type="error"
          title="Please correct the following errors:"
          :messages="errorList"
          compact
        />
      </div>
      <AppDialog
        v-model="isSuccessDialogOpen"
        title="Welcome Aboard!"
      >
        <p class="text-base text-zinc-300">{{ successMessage }}</p>
        <template #footer>
          <AppButton
            tone="primary"
            block
            @click="goToProfile"
          >
            Go to Profile
          </AppButton>
        </template>
      </AppDialog>
      <AppDialog
        v-model="isErrorDialogOpen"
        title="Registration Issue"
      >
        <p class="text-base text-zinc-300">
          We encountered a problem creating your account. Please check the form and try again.
        </p>
        <template #footer>
          <AppButton
            tone="danger"
            block
            @click="isErrorDialogOpen = false"
          >
            Close
          </AppButton>
        </template>
      </AppDialog>
    </template>

    <form @submit.prevent="handleSubmit" class="space-y-5">
      <div class="space-y-2">
        <p class="text-sm font-medium text-gray-700">I want to register as:</p>
        <div class="flex gap-4">
          <label class="inline-flex items-center gap-2 text-gray-700">
            <input
              v-model="role"
              type="radio"
              name="role"
              value="student"
              class="h-5 w-5 border-gray-300 text-purple-600 focus:ring-purple-500"
            />
            <span class="font-medium">Student</span>
          </label>
          <label class="inline-flex items-center gap-2 text-gray-700">
            <input
              v-model="role"
              type="radio"
              name="role"
              value="instructor"
              class="h-5 w-5 border-gray-300 text-purple-600 focus:ring-purple-500"
            />
            <span class="font-medium">Instructor</span>
          </label>
        </div>
      </div>

      <AppInput
        id="username"
        v-model="username"
        label="Username"
        name="username"
        required
        tone="secondary"
        :error="errors.username"
      />

      <AppInput
        id="email"
        v-model="email"
        label="Email Address"
        name="email"
        type="email"
        autocomplete="email"
        required
        tone="secondary"
        :error="errors.email"
      />

      <AppInput
        id="password"
        v-model="password"
        label="Password"
        name="password"
        type="password"
        autocomplete="new-password"
        required
        tone="secondary"
        hint="Min 8 chars, incl. uppercase, lowercase, number, special char."
        :error="errors.password"
      />

      <AppInput
        id="confirmPassword"
        v-model="confirmPassword"
        label="Confirm Password"
        name="confirmPassword"
        type="password"
        autocomplete="new-password"
        required
        tone="secondary"
        :error="errors.confirmPassword"
      />

      <AppButton
        type="submit"
        tone="secondary"
        :loading="userStore.loading"
        loading-label="Registering..."
        :disabled="userStore.loading"
        block
      >
        Register
      </AppButton>
    </form>

    <template #footer>
      <p class="text-center text-sm text-gray-700 sm:text-base">
        Already have an account?
        <AppButton
          type="button"
          variant="ghost"
          tone="secondary"
          @click="router.push('/login')"
        >
          Log in
        </AppButton>
      </p>
    </template>
  </AuthShell>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useRegisterSEO } from '../composables/useSEO';
import AlertMessage from '../components/AlertMessage.vue';
import AppButton from '../components/ui/AppButton.vue';
import AppInput from '../components/ui/AppInput.vue';
import AuthShell from '../components/ui/AuthShell.vue';
import AppDialog from '../components/ui/AppDialog.vue';
import { useUserStore } from '../stores/userStore';

const router = useRouter();
const userStore = useUserStore();

useRegisterSEO();

const username = ref('');
const email = ref('');
const password = ref('');
const confirmPassword = ref('');
const role = ref<'student' | 'instructor'>('student');
const errors = reactive<Record<string, string>>({});
const successMessage = ref('');
const isSuccessDialogOpen = ref(false);
const isErrorDialogOpen = ref(false);

const hasErrors = computed(() => Object.keys(errors).length > 0);
const errorList = computed(() => Object.values(errors));

const clearErrors = () => {
  Object.keys(errors).forEach((key) => delete errors[key]);
};

const goToProfile = () => {
  isSuccessDialogOpen.value = false;
  router.push('/profile');
};

const validateForm = (): boolean => {
  clearErrors();

  if (!username.value.trim()) errors.username = 'Username is required.';
  else if (username.value.trim().length < 3) errors.username = 'Username must be at least 3 characters.';

  if (!email.value.trim()) errors.email = 'Email is required.';
  else if (!/\S+@\S+\.\S+/.test(email.value)) errors.email = 'Invalid email address.';

  if (!password.value) errors.password = 'Password is required.';
  else if (password.value.length < 8) errors.password = 'Password must be at least 8 characters.';
  else if (!/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])/.test(password.value)) {
    errors.password = 'Password must include uppercase, lowercase, number, special character.';
  }

  if (!confirmPassword.value) {
    errors.confirmPassword = 'Please confirm your password.';
  } else if (password.value !== confirmPassword.value) {
    errors.confirmPassword = 'Passwords do not match.';
  }

  return Object.keys(errors).length === 0;
};

const handleSubmit = async () => {
  successMessage.value = '';
  isSuccessDialogOpen.value = false;
  isErrorDialogOpen.value = false;
  clearErrors();

  if (!validateForm()) return;

  try {
    await userStore.registerUser({
      username: username.value,
      email: email.value,
      password: password.value,
      role: role.value,
    });

    successMessage.value = 'Registration successful! We\'re excited to have you on board.';
    isSuccessDialogOpen.value = true;
  } catch (err: any) {
    if (err.errors) {
      if (err.errors.username) errors.username = Array.isArray(err.errors.username) ? err.errors.username[0] : err.errors.username;
      if (err.errors.email) errors.email = Array.isArray(err.errors.email) ? err.errors.email[0] : err.errors.email;
    } else {
      errors.apiError = err.message || 'Registration failed.';
    }
    isErrorDialogOpen.value = true;
  }
};
</script>
