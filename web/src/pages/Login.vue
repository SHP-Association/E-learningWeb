<template>
  <AuthShell
    title="Welcome Back!"
    subtitle="Log in to access your courses and profile."
    tone="primary"
  >
    <template #status>
      <AlertMessage
        v-if="errorMessage && !userStore.error"
        type="error"
        :message="errorMessage"
        compact
      />
      <AppDialog
        v-model="isErrorDialogOpen"
        title="Login Failed"
      >
        <p class="text-base text-zinc-300">
          {{ errorMessage }}
        </p>
        <template #footer>
          <AppButton
            tone="primary"
            block
            @click="isErrorDialogOpen = false"
          >
            Try Again
          </AppButton>
        </template>
      </AppDialog>
    </template>

    <form @submit.prevent="handleSubmit" class="space-y-5">
      <AppInput
        id="username"
        v-model="username"
        label="Username"
        name="username"
        autocomplete="username"
        required
        tone="primary"
      />

      <AppInput
        id="password"
        v-model="password"
        label="Password"
        name="password"
        type="password"
        autocomplete="current-password"
        required
        tone="primary"
      />

      <div class="flex justify-end">
        <AppButton
          type="button"
          variant="ghost"
          tone="primary"
          @click="router.push('/password_reset')"
        >
          Forgot password?
        </AppButton>
      </div>

      <AppButton
        type="submit"
        tone="primary"
        :loading="userStore.loading"
        loading-label="Logging In..."
        :disabled="userStore.loading"
        block
      >
        Login
      </AppButton>
    </form>

    <template #footer>
      <p class="text-center text-sm text-gray-700 sm:text-base">
        New to SHP-Learner?
        <AppButton
          type="button"
          variant="ghost"
          tone="primary"
          @click="router.push('/register')"
        >
          Create an account
        </AppButton>
      </p>
    </template>
  </AuthShell>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useLoginSEO } from '../composables/useSEO';
import AlertMessage from '../components/AlertMessage.vue';
import AppButton from '../components/ui/AppButton.vue';
import AppInput from '../components/ui/AppInput.vue';
import AuthShell from '../components/ui/AuthShell.vue';
import AppDialog from '../components/ui/AppDialog.vue';
import { useUserStore } from '../stores/userStore';

const router = useRouter();
const userStore = useUserStore();

useLoginSEO();

const username = ref('');
const password = ref('');
const errorMessage = ref('');
const isErrorDialogOpen = ref(false);

const handleSubmit = async () => {
  errorMessage.value = '';
  isErrorDialogOpen.value = false;

  const success = await userStore.loginUser({
    username: username.value,
    password: password.value,
  });

  if (success) {
    router.push('/profile');
  } else {
    errorMessage.value = userStore.error || 'Invalid username or password.';
    isErrorDialogOpen.value = true;
  }
};
</script>
