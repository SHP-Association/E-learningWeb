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
  const isLoggedIn = computed(() => !!user.value);
  const isAdmin = computed(() => user.value?.is_staff ?? false);
  const isInstructor = computed(() => user.value?.role === 'instructor');

  // Actions
  function loadUser() {
    const storedUser = localStorage.getItem('currentUser');
    if (storedUser) {
      try {
        user.value = JSON.parse(storedUser);
      } catch (e) {
        console.error('Failed to parse stored user:', e);
        localStorage.removeItem('currentUser');
      }
    }
  }

  function setUser(userData: User) {
    user.value = userData;
    localStorage.setItem('currentUser', JSON.stringify(userData));
  }

  function clearUser() {
    user.value = null;
    localStorage.removeItem('currentUser');
  }

  async function loginUser(credentials: LoginCredentials): Promise<void> {
    loading.value = true;
    error.value = null;

    try {
      // Login request
      await apiService.post('/api/login/', credentials);

      // Fetch user profile
      const users = await apiService.get<User[]>('/api/users/');
      const loggedInUser = users.find(u => u.username === credentials.username);

      if (!loggedInUser) {
        throw new Error('User profile not found');
      }

      setUser(loggedInUser);
    } catch (err: any) {
      error.value = err.message || 'Login failed';
      throw err;
    } finally {
      loading.value = false;
    }
  }

  async function registerUser(data: RegisterData): Promise<User> {
    loading.value = true;
    error.value = null;

    try {
      const response = await apiService.post<User>('/api/register/', data);

      // Set user after successful registration
      const newUser: User = {
        ...response,
        username: data.username,
        email: data.email,
        role: data.role,
        is_staff: data.role === 'instructor',
      };

      setUser(newUser);
      return newUser;
    } catch (err: any) {
      error.value = err.message || 'Registration failed';
      throw err;
    } finally {
      loading.value = false;
    }
  }

  async function updateProfile(updates: Partial<User>): Promise<void> {
    if (!user.value) {
      throw new Error('No user logged in');
    }

    loading.value = true;
    error.value = null;

    try {
      const updatedUser = await apiService.put<User>(
        `/api/users/${user.value.id}/`,
        updates
      );
      setUser(updatedUser);
    } catch (err: any) {
      error.value = err.message || 'Profile update failed';
      throw err;
    } finally {
      loading.value = false;
    }
  }

  function logout() {
    clearUser();
  }

  // Alias for compatibility
  const handleLogout = logout;

  return {
    // State
    user,
    loading,
    error,
    // Getters
    isLoggedIn,
    isAdmin,
    isInstructor,
    // Actions
    loadUser,
    setUser,
    clearUser,
    loginUser,
    registerUser,
    updateProfile,
    logout,
    handleLogout,
  };
});
