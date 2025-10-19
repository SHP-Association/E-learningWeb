import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useUserStore = defineStore('user', () => {
  // State
  const user = ref<any>(null);

  // Getters
  const isLoggedIn = () => !!user.value;
  const isAdmin = () => user.value?.is_staff ?? false;

  // Actions
  function setUser(userData: any) {
    user.value = userData;
  }

  function clearUser() {
    user.value = null;
  }

  function logout() {
    clearUser();
    // Optionally, remove token from localStorage/cookies
  }

  return {
    user,
    isLoggedIn,
    isAdmin,
    setUser,
    clearUser,
    logout,
  };
});
