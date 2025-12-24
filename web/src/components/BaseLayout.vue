<template>
  <div class="min-h-screen flex flex-col bg-gray-100 font-sans">
    <!-- Navbar -->
    <nav class="bg-blue-900 p-4 shadow-lg">
      <div class="container mx-auto flex justify-between items-center">
        <!-- Logo -->
        <button
          class="flex items-center text-white text-xl md:text-2xl font-bold rounded-lg px-2 py-1 md:px-3 md:py-2 hover:bg-blue-800 focus:outline-none"
          @click="navigate('/')"
          aria-label="Go to homepage"
        >
          <img :src="logo" alt="SHP-Learner Logo" class="w-8 h-8 md:w-10 md:h-10 mr-2 rounded-full" />
          SHP-Learner
        </button>

        <!-- Navigation Buttons -->
        <div class="flex items-center space-x-4 md:space-x-8">
          <button
            @click="navigate('/')"
            class="text-white hover:text-blue-200 hover:bg-blue-800 rounded-lg px-4 py-2 text-base md:text-lg hidden md:block"
          >
            Home
          </button>

          <button
            v-if="user?.is_staff"
            @click="openAdmin"
            class="text-white hover:text-blue-200 hover:bg-blue-800 rounded-lg px-4 py-2 text-base md:text-lg hidden md:block"
          >
            Admin
          </button>

          <button
            v-if="!user"
            @click="navigate('/login')"
            class="text-white hover:text-blue-200 hover:bg-blue-800 rounded-lg px-4 py-2 text-base md:text-lg hidden md:block"
          >
            Login
          </button>

          <!-- Profile Dropdown -->
          <div v-if="user" class="relative" ref="profileBtnRef">
            <button @click="toggleDropdown" class="focus:outline-none" aria-label="Open profile menu">
              <img
                v-if="user.profile_picture"
                :src="user.profile_picture"
                alt="Profile"
                class="w-9 md:w-11 h-9 md:h-11 rounded-full border-2 border-yellow-400 object-cover shadow"
              />
              <span v-else class="flex items-center justify-center bg-yellow-400 rounded-full w-9 h-9 md:w-11 md:h-11">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-6 md:w-7 h-6 md:h-7 text-blue-900" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <circle cx="12" cy="8" r="4" stroke="currentColor" stroke-width="2" fill="none"/>
                  <path stroke="currentColor" stroke-width="2" d="M4 20c0-4 4-7 8-7s8 3 8 7"/>
                </svg>
              </span>
            </button>

            <div
              v-if="showProfileDropdown"
              ref="profileDropdownRef"
              class="absolute right-0 mt-2 w-44 bg-white rounded-lg shadow-lg z-50 border border-gray-200"
            >
              <button @click="goProfile" class="block w-full text-left px-4 py-2 text-gray-800 hover:bg-blue-100 hover:text-blue-900 rounded-t-lg">
                Profile
              </button>
              <button @click="logout" class="block w-full text-left px-4 py-2 text-gray-800 hover:bg-blue-100 hover:text-blue-900 rounded-b-lg">
                Logout
              </button>
            </div>
          </div>

          <!-- Mobile login button -->
          <button
            v-else
            @click="navigate('/login')"
            class="text-white bg-yellow-400 hover:bg-yellow-500 rounded-full px-5 py-2 font-semibold text-base shadow focus:outline-none md:hidden"
          >
            Login
          </button>
        </div>
      </div>
    </nav>

    <!-- Main content slot -->
    <main class="container mx-auto p-4 md:p-6 flex-grow">
      <router-view />
    </main>

    <!-- Footer -->
    <footer class="bg-blue-900 text-white py-10 mt-auto shadow-inner">
      <div class="container mx-auto grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-10 px-4">
        <div>
          <h4 class="text-2xl font-bold mb-4 flex items-center gap-2">
            <span class="bg-yellow-400 text-blue-900 px-2 py-1 rounded-lg font-black">SHP</span>
            About Us
          </h4>
          <p class="text-gray-300 text-base leading-relaxed">
            SHP-Learnering Platform is dedicated to providing high-quality courses...
          </p>
        </div>

        <div>
          <h4 class="text-2xl font-bold mb-4">Quick Links</h4>
          <ul>
            <li class="mb-2"><button @click="navigate('/')" class="text-gray-300 hover:text-yellow-300">Home</button></li>
            <li class="mb-2"><a href="https://wa.me/9399613606" target="_blank" class="text-gray-300 hover:text-yellow-300">Contact Us</a></li>
            <li class="mb-2"><button @click="navigate('/faq')" class="text-gray-300 hover:text-yellow-300">FAQ</button></li>
          </ul>
        </div>

        <div>
          <h4 class="text-2xl font-bold mb-4">Contact Us</h4>
          <p class="text-gray-300">Email: <a href="mailto:sandeshpatel.sp.93@gmail.com" class="hover:text-yellow-300">sandeshpatel.sp.93@gmail.com</a></p>
          <p class="text-gray-300">Phone: <a href="tel:+919399613606" class="hover:text-yellow-300">+91 9399613606</a></p>
          <p class="text-gray-300">Address: 420 Kareli St, MP Narsinghpur, India</p>
        </div>

        <div>
          <h4 class="text-2xl font-bold mb-4">Follow Us</h4>
          <div class="flex space-x-5 mt-2">
            <a href="https://x.com/SandeshPat007" target="_blank" class="text-gray-300 hover:text-yellow-300 text-3xl">X</a>
            <a href="https://www.linkedin.com/in/sandesh-patel07" target="_blank" class="text-gray-300 hover:text-yellow-300 text-3xl">IN</a>
            <a href="https://www.instagram.com/sandesh_patel007" target="_blank" class="text-gray-300 hover:text-yellow-300 text-3xl">IG</a>
          </div>
        </div>
      </div>
      <div class="footer-bottom text-center text-gray-400 mt-10 border-t border-blue-800 pt-5 px-4">
        <p>&copy; {{ new Date().getFullYear() }} SHP-Learnering Platform. All Rights Reserved.</p>
      </div>
    </footer>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, onBeforeUnmount, inject } from 'vue';
import { useRouter } from 'vue-router';
import logo from '../assets/logo.png';
import { useUserStore } from '../stores/userStore';

export default defineComponent({
  name: 'BaseLayout',
  setup() {
    const router = useRouter();
    // const userStore = inject('userStore') as any; // assume provide() in App.vue
    const user = useUserStore?.user;
    const handleLogout = useUserStore?.handleLogout;

    const showProfileDropdown = ref(false);
    const profileBtnRef = ref<HTMLElement | null>(null);
    const profileDropdownRef = ref<HTMLElement | null>(null);

    const navigate = (path: string) => router.push(path);

    const toggleDropdown = () => {
      showProfileDropdown.value = !showProfileDropdown.value;
    };

    const goProfile = () => {
      navigate('/profile');
      showProfileDropdown.value = false;
    };

    const logout = () => {
      handleLogout?.();
      showProfileDropdown.value = false;
    };

    const openAdmin = () => {
      window.open(`${import.meta.env.VITE_APP_BACKEND_URL}/admin`, '_blank');
    };

    const handleClickOutside = (event: MouseEvent) => {
      if (
        profileBtnRef.value &&
        !profileBtnRef.value.contains(event.target as Node) &&
        profileDropdownRef.value &&
        !profileDropdownRef.value.contains(event.target as Node)
      ) {
        showProfileDropdown.value = false;
      }
    };

    onMounted(() => document.addEventListener('mousedown', handleClickOutside));
    onBeforeUnmount(() => document.removeEventListener('mousedown', handleClickOutside));

    return {
      logo,
      user,
      showProfileDropdown,
      profileBtnRef,
      profileDropdownRef,
      navigate,
      toggleDropdown,
      goProfile,
      logout,
      openAdmin,
    };
  },
});
</script>
