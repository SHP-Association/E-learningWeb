<!-- App.vue -->
<template>
  <BaseLayout>
    <router-view />
  </BaseLayout>
</template>

<script lang="ts">
import { defineComponent, provide, onMounted } from 'vue';
import BaseLayout from './components/BaseLayout.vue';
import { useUserStore } from './stores/userStore';

export default defineComponent({
  name: 'App',
  components: { BaseLayout },
  setup() {
    const userStore = useUserStore();

    // Provide user context for pages that need it (Login, Register)
    provide('user-context', {
      handleLogin: (user: any) => userStore.saveUser(user),
      handleLogout: () => userStore.logout(),
    });

    // Load user from localStorage on app init
    onMounted(() => {
      userStore.loadUser();
    });
  },
});
</script>

<style>
/* Add any global styles here */
</style>
