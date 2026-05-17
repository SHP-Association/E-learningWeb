import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { apiService } from '../services/api.service';
import type { User, LoginCredentials, RegisterData } from '../types/api.types';

export const useUserStore = defineStore('user', () => {
  // State
  const user = ref<User | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  // Getters
  const isLoggedIn = computed(() => user.value !== null);
  const isInstructor = computed(() => user.value?.role === 'instructor');
  const isAdmin = computed(() => user.value?.role === 'admin' || user.value?.is_staff === true);
  const isStudent = computed(() => user.value?.role === 'student');

  // Load user from localStorage on app start
  function loadUser() {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      try {
        user.value = JSON.parse(storedUser);
      } catch (e) {
        console.error('Failed to parse stored user:', e);
        localStorage.removeItem('user');
      }
    }
  }

  // Save user to localStorage
  function saveUser(userData: User) {
    user.value = userData;
    localStorage.setItem('user', JSON.stringify(userData));
  }

  // Clear user from localStorage
  function clearUser() {
    user.value = null;
    localStorage.removeItem('user');
  }

  // Actions
  async function login(credentials: LoginCredentials): Promise<boolean> {
    loading.value = true;
    error.value = null;

    try {
      // Call backend login API via the centralized apiService helper
      const response = await apiService.login({
        email: credentials.email || credentials.username || '',
        password: credentials.password
      });

      // Handle new standardized response format
      if (response.success || response.message === 'login successful' || response.message === 'Login successful') {
        // Fetch user profile after successful login
        await fetchUserProfile();
        return true;
      }

      error.value = response.message || 'Login failed';
      return false;
    } catch (err: any) {
      error.value = err.message || 'Login failed';
      console.error('Login error:', err);
      return false;
    } finally {
      loading.value = false;
    }
  }

  async function register(data: RegisterData): Promise<boolean> {
    loading.value = true;
    error.value = null;

    try {
      const response = await apiService.register({
        username: data.username,
        email: data.email,
        password: data.password
      });

      // Handle new standardized response format
      if (
        response.success || 
        response.message === 'Registration successful!' || 
        response.message === 'account created and logged in'
      ) {
        // Auto-login after registration
        return await login({
          email: data.email,
          username: data.username,
          password: data.password,
        });
      }

      error.value = response.message || 'Registration failed';
      return false;
    } catch (err: any) {
      error.value = err.message || 'Registration failed';
      console.error('Registration error:', err);
      return false;
    } finally {
      loading.value = false;
    }
  }

  async function logout(): Promise<void> {
    loading.value = true;
    error.value = null;

    try {
      await apiService.logout();
    } catch (err: any) {
      console.error('Logout error:', err);
    } finally {
      clearUser();
      loading.value = false;
    }
  }

  async function fetchUserProfile(): Promise<void> {
    loading.value = true;
    error.value = null;

    try {
      // Fetch current user profile using the centralized apiService helper
      const response = await apiService.getCurrentUser();

      // Handle new standardized response format
      if (response.data && typeof response.data === 'object' && !Array.isArray(response.data)) {
        saveUser(response.data);
      } else {
        saveUser(response);
      }
    } catch (err: any) {
      error.value = err.message || 'Failed to fetch user profile';
      console.error('Error fetching user profile:', err);
    } finally {
      loading.value = false;
    }
  }

  async function updateProfile(updates: Partial<User> | FormData): Promise<boolean> {
    if (!user.value) {
      error.value = 'No user logged in';
      return false;
    }

    loading.value = true;
    error.value = null;

    try {
      const response = await apiService.updateUser(user.value.id, updates);

      // Handle new standardized response format or direct object
      let updatedUser: User;
      if (response.data && typeof response.data === 'object' && !Array.isArray(response.data)) {
        updatedUser = response.data;
      } else {
        updatedUser = response;
      }

      saveUser(updatedUser);
      return true;
    } catch (err: any) {
      error.value = err.message || 'Failed to update profile';
      console.error('Error updating profile:', err);
      return false;
    } finally {
      loading.value = false;
    }
  }

  function clearError() {
    error.value = null;
  }

  // Aliases for compatibility
  const loginUser = login;
  const registerUser = register;
  const handleLogout = logout;

  return {
    // State
    user,
    loading,
    error,
    // Getters
    isLoggedIn,
    isInstructor,
    isAdmin,
    isStudent,
    // Actions
    loadUser,
    saveUser,
    clearUser,
    login,
    loginUser,
    register,
    registerUser,
    logout,
    handleLogout,
    fetchUserProfile,
    updateProfile,
    clearError,
  };
});
