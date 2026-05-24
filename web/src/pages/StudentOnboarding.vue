<template>
  <AuthShell
    title="Complete Student Onboarding"
    subtitle="Add a few required details. You can skip for now and update later."
    tone="secondary"
  >
    <template #status>
      <AlertMessage
        v-if="userStore.error"
        type="error"
        :message="userStore.error"
        compact
      />
      <AlertMessage
        v-else-if="statusMessage"
        type="success"
        :message="statusMessage"
        compact
      />
    </template>

    <div v-if="!userStore.isLoggedIn" class="space-y-4">
      <p class="text-sm text-zinc-300">Please log in first to continue onboarding.</p>
      <AppButton tone="secondary" block @click="router.push('/login')">
        Go to Login
      </AppButton>
    </div>

    <form v-else @submit.prevent="handleSubmit" class="space-y-5">
      <AppInput
        id="first_name"
        v-model="form.first_name"
        label="First Name"
        name="first_name"
        required
        tone="secondary"
      />

      <AppInput
        id="last_name"
        v-model="form.last_name"
        label="Last Name"
        name="last_name"
        required
        tone="secondary"
      />

      <AppInput
        id="contact_number"
        v-model="form.contact_number"
        label="Contact Number"
        name="contact_number"
        required
        tone="secondary"
      />

      <AppInput
        id="country"
        v-model="form.country"
        label="Country"
        name="country"
        required
        tone="secondary"
      />

      <div class="flex gap-3">
        <AppButton
          type="submit"
          tone="secondary"
          :loading="userStore.loading"
          loading-label="Saving..."
          :disabled="userStore.loading"
          block
        >
          Save and Continue
        </AppButton>

        <AppButton
          type="button"
          variant="ghost"
          tone="secondary"
          :disabled="userStore.loading"
          block
          @click="skipForNow"
        >
          Skip for now
        </AppButton>
      </div>
    </form>
  </AuthShell>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import AlertMessage from '../components/AlertMessage.vue';
import AppButton from '../components/ui/AppButton.vue';
import AppInput from '../components/ui/AppInput.vue';
import AuthShell from '../components/ui/AuthShell.vue';
import { useUserStore } from '../stores/userStore';

const router = useRouter();
const userStore = useUserStore();

const statusMessage = ref('');
const form = reactive({
  first_name: '',
  last_name: '',
  contact_number: '',
  country: '',
});

watch(
  () => userStore.user,
  (newUser) => {
    if (!newUser) return;
    form.first_name = newUser.first_name || '';
    form.last_name = newUser.last_name || '';
    form.contact_number = newUser.contact_number || '';
    form.country = newUser.country || '';
  },
  { immediate: true }
);

const handleSubmit = async () => {
  statusMessage.value = '';
  userStore.clearError();

  const ok = await userStore.submitOnboarding({
    first_name: form.first_name.trim(),
    last_name: form.last_name.trim(),
    contact_number: form.contact_number.trim(),
    country: form.country.trim(),
  });

  if (!ok) {
    return;
  }

  statusMessage.value = 'Onboarding details saved.';
  router.push('/profile');
};

const skipForNow = () => {
  router.push('/profile');
};
</script>
