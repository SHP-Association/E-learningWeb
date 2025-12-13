import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { apiService } from '../services/api.service';
import type { Enrollment, Course } from '../types/api.types';
import { useUserStore } from './userStore';

export const useEnrollmentStore = defineStore('enrollment', () => {
    // State
    const enrollments = ref<Enrollment[]>([]);
    const loading = ref(false);
    const error = ref<string | null>(null);

    // Getters
    const enrolledCourses = computed(() =>
        enrollments.value.map(e => e.course)
    );

    const completedCourses = computed(() =>
        enrollments.value.filter(e => e.completed).map(e => e.course)
    );

    const inProgressCourses = computed(() =>
        enrollments.value.filter(e => !e.completed && e.progress > 0).map(e => e.course)
    );

    // Actions
    async function fetchEnrollments(): Promise<void> {
        const userStore = useUserStore();

        if (!userStore.isLoggedIn) {
            enrollments.value = [];
            return;
        }

        loading.value = true;
        error.value = null;

        try {
            // Try to fetch from localStorage first (for mock data)
            const storedEnrollments = localStorage.getItem('mockEnrollments');
            if (storedEnrollments) {
                enrollments.value = JSON.parse(storedEnrollments);
            } else {
                // Fetch from API
                const fetchedEnrollments = await apiService.get<Enrollment[]>('/api/enrollments/');
                enrollments.value = fetchedEnrollments;
            }
        } catch (err: any) {
            error.value = err.message || 'Failed to load enrollments';
            console.error('Error fetching enrollments:', err);
            // Fallback to empty array on error
            enrollments.value = [];
        } finally {
            loading.value = false;
        }
    }

    async function enrollInCourse(courseSlug: string): Promise<boolean> {
        const userStore = useUserStore();

        if (!userStore.user) {
            error.value = 'You must be logged in to enroll';
            return false;
        }

        loading.value = true;
        error.value = null;

        try {
            const enrollment = await apiService.post<Enrollment>('/api/enrollments/', {
                course_slug: courseSlug,
            });

            enrollments.value.push(enrollment);

            // Save to localStorage for persistence
            localStorage.setItem('mockEnrollments', JSON.stringify(enrollments.value));

            return true;
        } catch (err: any) {
            error.value = err.message || 'Enrollment failed';
            console.error('Error enrolling in course:', err);
            return false;
        } finally {
            loading.value = false;
        }
    }

    function isEnrolled(courseSlug: string): boolean {
        return enrollments.value.some(e => e.course.slug === courseSlug);
    }

    function getEnrollment(courseSlug: string): Enrollment | undefined {
        return enrollments.value.find(e => e.course.slug === courseSlug);
    }

    function updateProgress(courseSlug: string, progress: number): void {
        const enrollment = enrollments.value.find(e => e.course.slug === courseSlug);
        if (enrollment) {
            enrollment.progress = progress;
            enrollment.completed = progress >= 100;

            // Save to localStorage
            localStorage.setItem('mockEnrollments', JSON.stringify(enrollments.value));
        }
    }

    function clearEnrollments() {
        enrollments.value = [];
        localStorage.removeItem('mockEnrollments');
    }

    return {
        // State
        enrollments,
        loading,
        error,
        // Getters
        enrolledCourses,
        completedCourses,
        inProgressCourses,
        // Actions
        fetchEnrollments,
        enrollInCourse,
        isEnrolled,
        getEnrollment,
        updateProgress,
        clearEnrollments,
    };
});
