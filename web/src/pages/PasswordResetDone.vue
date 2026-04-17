<template>
  <div class="ui-page-shell">
    <AppCard size="md" padding="lg" elevated centered>
      <template #header>
        <div class="text-center">
          <h1 class="text-4xl font-bold text-blue-800">Password Reset Email Sent</h1>
        </div>
      </template>

      <div class="space-y-4 text-center text-gray-700">
        <p>
          We've emailed you instructions for setting your password, if an account exists with the email you entered.
          You should receive them shortly.
        </p>
        <p>
          If you don't receive an email, please make sure you've entered the address you registered with,
          and check your spam folder.
        </p>
      </div>

      <template #footer>
        <AppButton tone="primary" block @click="router.push('/login')">
          Return to Login
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
