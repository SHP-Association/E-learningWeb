<!-- App.vue -->
<template>
  <BaseLayout :user="user" @logout="handleLogout" :navigate="navigate">
    <Router
      :currentRoute="currentRoute"
      :user="user"
      :courses="mockCourses"
      :enrollments="mockEnrollments"
      @login="handleLogin"
      :navigate="navigate"
      @addEnrollment="addEnrollment"
    />
  </BaseLayout>
</template>

<script lang="ts">
import { defineComponent, ref, reactive, onMounted, watch } from 'vue';
import BaseLayout from './components/BaseLayout.vue';
import Router from './router'; // Make sure this is your Vue router or custom component

export default defineComponent({
  name: 'App',
  components: { BaseLayout, Router },
  setup() {
    // State
    const currentRoute = ref(window.location.pathname);
    const user = ref(null);

    const mockCourses = reactive([
      /* ...same as your previous courses... */
    ]);

    const mockEnrollments = ref(
      JSON.parse(localStorage.getItem('mockEnrollments') || '[]')
    );

    // Load user from localStorage
    onMounted(() => {
      const storedUser = localStorage.getItem('currentUser');
      if (storedUser) user.value = JSON.parse(storedUser);

      const handlePopState = () => (currentRoute.value = window.location.pathname);
      window.addEventListener('popstate', handlePopState);

      // Cleanup
      return () => window.removeEventListener('popstate', handlePopState);
    });

    // Watch enrollments to save in localStorage
    watch(
      mockEnrollments,
      (newVal) => localStorage.setItem('mockEnrollments', JSON.stringify(newVal)),
      { deep: true }
    );

    // Navigation
    const navigate = (path: string) => {
      window.history.pushState({}, '', path);
      currentRoute.value = path;
    };

    // Login / Logout
    const handleLogin = (userData: any) => {
      user.value = userData;
      localStorage.setItem('currentUser', JSON.stringify(userData));
      navigate('/');
    };

    const handleLogout = () => {
      user.value = null;
      localStorage.removeItem('currentUser');
      navigate('/');
    };

    // Add enrollment
    const addEnrollment = (student: any, course: any) => {
      const exists = mockEnrollments.value.some(
        (e) => e.student.username === student.username && e.course.slug === course.slug
      );
      if (!exists) {
        mockEnrollments.value.push({ student, course, progress: 0 });
        return true;
      }
      return false;
    };

    // Provide state for child components if needed
    return {
      currentRoute,
      user,
      mockCourses,
      mockEnrollments,
      navigate,
      handleLogin,
      handleLogout,
      addEnrollment,
    };
  },
});
</script>

<style>
/* Add any global styles here */
</style>
