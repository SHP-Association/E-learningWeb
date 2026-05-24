<template>
  <AuthShell
    title="Student Signup"
    subtitle="Create your student account, verify your email with OTP, then finish onboarding."
    tone="secondary"
  >
    <template #status>
      <AlertMessage
        v-if="hasErrors"
        type="error"
        title="Please correct the following errors:"
        :messages="errorList"
        compact
      />
      <AlertMessage
        v-if="userStore.error"
        type="error"
        :message="userStore.error"
        compact
      />
    </template>

    <form @submit.prevent="handleSubmit" class="space-y-5">
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
        loading-label="Creating account..."
        :disabled="userStore.loading"
        block
      >
        Continue
      </AppButton>
    </form>

    <div class="rounded-lg border border-zinc-700 bg-zinc-900/30 p-4 text-sm text-zinc-300">
      <p class="font-semibold text-zinc-100">Are you an instructor?</p>
      <p class="mt-1">
        Instructor signup is handled separately.
        <a :href="`mailto:${instructorContactEmail}`" class="text-cyan-300 underline">
          Contact us at {{ instructorContactEmail }}
        </a>
      </p>
    </div>

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
import { computed, onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useRegisterSEO } from '../composables/useSEO';
import AlertMessage from '../components/AlertMessage.vue';
import AppButton from '../components/ui/AppButton.vue';
import AppInput from '../components/ui/AppInput.vue';
import AuthShell from '../components/ui/AuthShell.vue';
import { useUserStore } from '../stores/userStore';
import { apiService } from '../services/api.service';

const router = useRouter();
const userStore = useUserStore();

useRegisterSEO();

const username = ref('');
const email = ref('');
const password = ref('');
const confirmPassword = ref('');
const instructorContactEmail = ref('admin@localhost');
const errors = reactive<Record<string, string>>({});

const hasErrors = computed(() => Object.keys(errors).length > 0);
const errorList = computed(() => Object.values(errors));

const clearErrors = () => {
  Object.keys(errors).forEach((key) => delete errors[key]);
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

const loadSignupConfig = async () => {
  try {
    const response = await apiService.getSignupConfig();
    if (response?.instructor_contact_email) {
      instructorContactEmail.value = response.instructor_contact_email;
    }
  } catch (err) {
    console.error('Failed to fetch signup config:', err);
  }
};

const handleSubmit = async () => {
  userStore.clearError();
  clearErrors();

  if (!validateForm()) return;

  const ok = await userStore.registerStudent({
    username: username.value.trim(),
    email: email.value.trim(),
    password: password.value,
  });

  if (!ok) {
    return;
  }

  router.push({
    path: '/register/verify',
    query: { email: userStore.pendingVerificationEmail || email.value.trim() },
  });
};

onMounted(() => {
  loadSignupConfig();
});
</script>
