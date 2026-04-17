<template>
  <div class="ui-page-shell">
    <AppCard size="md" padding="lg" elevated centered>
      <template #header>
        <div class="text-center">
          <h1 class="text-4xl font-bold text-blue-800">Password Reset Complete</h1>
        </div>
      </template>

      <p class="text-center text-gray-700">
        Your password has been set. You may now log in with your new password.
      </p>

      <template #footer>
        <AppButton tone="primary" block @click="router.push('/login')">
          Log in
        </AppButton>
      </template>
    </AppCard>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import AppButton from '../components/ui/AppButton.vue';
import AppCard from '../components/ui/AppCard.vue';

const router = useRouter();
</script>
