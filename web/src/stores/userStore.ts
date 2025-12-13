import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { apiService } from '../services/api.service';
import { SESSION, STORAGE_KEYS } from '../config/constants';
import type { User, LoginCredentials, RegisterData } from '../types/api.types';

export const useUserStore = defineStore('user', () => {
  // State
  const user = ref<User | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);
  let sessionTimeoutId: ReturnType<typeof setTimeout> | null = null;

  // Getters
  const isLoggedIn = computed(() => !!user.value);
  const isAdmin = computed(() => user.value?.is_staff ?? false);
  const isInstructor = computed(() => user.value?.role === 'instructor');

  // Session timeout management
  function resetSessionTimeout() {
    if (sessionTimeoutId) {
      clearTimeout(sessionTimeoutId);
    }

    if (user.value) {
      sessionTimeoutId = setTimeout(() => {
        console.warn('Session expired due to inactivity');
        logout();
        // Optionally redirect to login with timeout message
        window.location.href = '/login?timeout=true';
      }, SESSION.TIMEOUT);
    }
  }

  // Actions
  function loadUser() {
    const storedData = localStorage.getItem(STORAGE_KEYS.USER);
    if (storedData) {
      try {
        const data = JSON.parse(storedData);

        // Check if data has expired
        if (data.expiry && Date.now() > data.expiry) {
          console.warn('Stored user data expired');
          localStorage.removeItem(STORAGE_KEYS.USER);
          return;
        }

        user.value = data.user;
        resetSessionTimeout();
      } catch (e) {
        console.error('Failed to parse stored user:', e);
        localStorage.removeItem(STORAGE_KEYS.USER);
      }
    }
  }

  function setUser(userData: User) {
    const data = {
      user: userData,
      expiry: Date.now() + (24 * 60 * 60 * 1000), // 24 hours
    };
    user.value = userData;
    localStorage.setItem(STORAGE_KEYS.USER, JSON.stringify(data));
    resetSessionTimeout();
  }

  function clearUser() {
    user.value = null;
    localStorage.removeItem(STORAGE_KEYS.USER);
    if (sessionTimeoutId) {
      clearTimeout(sessionTimeoutId);
      sessionTimeoutId = null;
    }
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

  // Reset session timeout on user activity
  if (typeof window !== 'undefined') {
    window.addEventListener('click', resetSessionTimeout);
    window.addEventListener('keypress', resetSessionTimeout);
    window.addEventListener('scroll', resetSessionTimeout);
  }

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
    resetSessionTimeout,
  };
});
