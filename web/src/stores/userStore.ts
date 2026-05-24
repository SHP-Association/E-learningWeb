import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { apiService } from '../services/api.service';
import type {
  User,
  LoginCredentials,
  RegisterData,
  VerifyOTPData,
  OnboardingData,
} from '../types/api.types';

const USER_STORAGE_KEY = 'user';
const ONBOARDING_REMINDER_KEY = 'onboardingReminderDismissed';
const PENDING_SIGNUP_STORAGE_KEY = 'pendingSignupData';

interface PendingSignupData {
  username: string;
  email: string;
  password: string;
}

function computeNeedsOnboarding(userData: User | null): boolean {
  if (!userData || userData.role !== 'student') {
    return false;
  }

  return (
    !userData.first_name?.trim() ||
    !userData.last_name?.trim() ||
    !userData.contact_number?.trim() ||
    !userData.country?.trim()
  );
}

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);
  const pendingVerificationEmail = ref<string | null>(null);
  const pendingSignupData = ref<PendingSignupData | null>(null);
  const onboardingReminderDismissed = ref(false);

  const isLoggedIn = computed(() => user.value !== null);
  const isInstructor = computed(() => user.value?.role === 'instructor');
  const isAdmin = computed(() => user.value?.role === 'admin' || user.value?.is_staff === true);
  const isStudent = computed(() => user.value?.role === 'student');
  const needsOnboarding = computed(() => computeNeedsOnboarding(user.value));
  const showOnboardingReminder = computed(() => isLoggedIn.value && needsOnboarding.value && !onboardingReminderDismissed.value);

  function loadUser() {
    const storedUser = localStorage.getItem(USER_STORAGE_KEY);
    const reminderDismissed = localStorage.getItem(ONBOARDING_REMINDER_KEY);
    const storedPendingSignupData = localStorage.getItem(PENDING_SIGNUP_STORAGE_KEY);

    if (reminderDismissed === 'true') {
      onboardingReminderDismissed.value = true;
    }

    if (storedPendingSignupData) {
      try {
        pendingSignupData.value = JSON.parse(storedPendingSignupData);
        pendingVerificationEmail.value = pendingSignupData.value?.email || null;
      } catch (e) {
        console.error('Failed to parse pending signup data:', e);
        localStorage.removeItem(PENDING_SIGNUP_STORAGE_KEY);
      }
    }

    if (!storedUser) return;

    try {
      user.value = JSON.parse(storedUser);
    } catch (e) {
      console.error('Failed to parse stored user:', e);
      localStorage.removeItem(USER_STORAGE_KEY);
    }
  }

  function saveUser(userData: User) {
    user.value = {
      ...userData,
      onboarding_required: computeNeedsOnboarding(userData),
    };
    localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(user.value));
  }

  function clearUser() {
    user.value = null;
    localStorage.removeItem(USER_STORAGE_KEY);
    clearPendingVerification();
    onboardingReminderDismissed.value = false;
    localStorage.removeItem(ONBOARDING_REMINDER_KEY);
  }

  function clearPendingVerification() {
    pendingVerificationEmail.value = null;
    pendingSignupData.value = null;
    localStorage.removeItem(PENDING_SIGNUP_STORAGE_KEY);
  }

  function dismissOnboardingReminder() {
    onboardingReminderDismissed.value = true;
    localStorage.setItem(ONBOARDING_REMINDER_KEY, 'true');
  }

  function resetOnboardingReminder() {
    onboardingReminderDismissed.value = false;
    localStorage.removeItem(ONBOARDING_REMINDER_KEY);
  }

  async function login(credentials: LoginCredentials): Promise<boolean> {
    loading.value = true;
    error.value = null;

    try {
      const response = await apiService.login({
        email: credentials.email || credentials.username || '',
        password: credentials.password,
      });

      const loginSuccessful =
        response.success ||
        response.message === 'logged in successfully' ||
        response.message === 'login successful' ||
        response.message === 'Login successful';

      if (!loginSuccessful) {
        error.value = response.message || 'Login failed';
        return false;
      }

      clearPendingVerification();
      await fetchUserProfile();
      resetOnboardingReminder();
      return true;
    } catch (err: any) {
      const payload = err?.errors || {};
      if (payload.verification_pending === true) {
        pendingVerificationEmail.value = payload.email || null;
        error.value = 'Email verification required. Please verify your email with OTP.';
      } else {
        error.value = err.message || 'Login failed';
      }
      console.error('Login error:', err);
      return false;
    } finally {
      loading.value = false;
    }
  }

  async function registerStudent(data: RegisterData): Promise<boolean> {
    loading.value = true;
    error.value = null;

    try {
      const response = await apiService.register({
        username: data.username,
        email: data.email,
        password: data.password,
      });

      if (response.verification_required === true) {
        pendingVerificationEmail.value = response.email || data.email;
        pendingSignupData.value = {
          username: data.username,
          email: data.email,
          password: data.password,
        };
        localStorage.setItem(PENDING_SIGNUP_STORAGE_KEY, JSON.stringify(pendingSignupData.value));
        return true;
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

  async function verifyOtp(data: VerifyOTPData): Promise<boolean> {
    loading.value = true;
    error.value = null;

    try {
      const response = await apiService.verifyOTP(data);
      const verified =
        response?.message === 'verification successful, logged in' ||
        response?.message === 'email verified successfully, please log in manually';

      if (!verified) {
        error.value = response?.message || 'OTP verification failed';
        return false;
      }

      clearPendingVerification();
      await fetchUserProfile();
      resetOnboardingReminder();
      return true;
    } catch (err: any) {
      error.value = err.message || 'OTP verification failed';
      console.error('Verify OTP error:', err);
      return false;
    } finally {
      loading.value = false;
    }
  }

  async function resendOtp(): Promise<boolean> {
    if (!pendingSignupData.value) {
      error.value = 'Signup details are missing. Please register again.';
      return false;
    }

    return registerStudent(pendingSignupData.value);
  }

  async function submitOnboarding(data: OnboardingData): Promise<boolean> {
    if (!user.value) {
      error.value = 'No user logged in';
      return false;
    }

    loading.value = true;
    error.value = null;

    try {
      const response = await apiService.submitOnboarding(data);
      const updatedUser = (response?.user || response?.data || response) as User;
      saveUser(updatedUser);
      resetOnboardingReminder();
      return true;
    } catch (err: any) {
      error.value = err.message || 'Failed to save onboarding details';
      console.error('Onboarding submit error:', err);
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
      const response = await apiService.getCurrentUser();
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

  const loginUser = login;
  const register = registerStudent;
  const registerUser = registerStudent;
  const handleLogout = logout;

  return {
    user,
    loading,
    error,
    pendingVerificationEmail,
    pendingSignupData,
    isLoggedIn,
    isInstructor,
    isAdmin,
    isStudent,
    needsOnboarding,
    showOnboardingReminder,
    loadUser,
    saveUser,
    clearUser,
    clearPendingVerification,
    dismissOnboardingReminder,
    resetOnboardingReminder,
    login,
    loginUser,
    register,
    registerUser,
    registerStudent,
    verifyOtp,
    resendOtp,
    submitOnboarding,
    logout,
    handleLogout,
    fetchUserProfile,
    updateProfile,
    clearError,
  };
});
