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
      // Call backend login API
      const response = await apiService.post<{ detail: string }>('/api/login/', credentials);

      if (response.detail === 'Login successful') {
        // Fetch user profile after successful login
        await fetchUserProfile();
        return true;
      }

      error.value = 'Login failed';
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
      const response = await apiService.post<{ message: string }>('/api/register/', data);

      if (response.message === 'Registration successful!') {
        // Auto-login after registration
        return await login({
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
      await apiService.post('/api/logout/', {});
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
      // Fetch current user profile
      const response = await apiService.get<any>('/api/users/');

      // Handle paginated response
      let users: User[];
      if (response.results && Array.isArray(response.results)) {
        users = response.results;
      } else if (Array.isArray(response)) {
        users = response;
      } else {
        // Single user object
        saveUser(response);
        return;
      }

      // Get first user (should be current user)
      if (users.length > 0) {
        saveUser(users[0]);
      }
    } catch (err: any) {
      error.value = err.message || 'Failed to fetch user profile';
      console.error('Error fetching user profile:', err);
    } finally {
      loading.value = false;
    }
  }

  async function updateProfile(updates: Partial<User>): Promise<boolean> {
    if (!user.value) {
      error.value = 'No user logged in';
      return false;
    }

    loading.value = true;
    error.value = null;

    try {
      const updatedUser = await apiService.patch<User>(
        `/api/users/${user.value.id}/`,
        updates
      );
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
