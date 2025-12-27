import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { apiService } from '../services/api.service';
import type { Course } from '../types/api.types';
import { transformCourse, transformPaginatedResponse } from '../utils/transformers';

export const useCourseStore = defineStore('course', () => {
    // State
    const courses = ref<Course[]>([]);
    const currentCourse = ref<Course | null>(null);
    const loading = ref(false);
    const error = ref<string | null>(null);

    // Cache to avoid redundant API calls
    const coursesCache = ref<Map<string, Course>>(new Map());
    const lastFetchTime = ref<number>(0);
    const CACHE_DURATION = 5 * 60 * 1000; // 5 minutes

    // Getters
    const featuredCourse = computed(() => courses.value[0] || null);
    const freeCourses = computed(() => courses.value.filter(c => c.is_free));
    const paidCourses = computed(() => courses.value.filter(c => !c.is_free));

    // Actions
    async function fetchCourses(forceRefresh = false): Promise<void> {
        const now = Date.now();

        // Use cache if available and not expired
        if (!forceRefresh && courses.value.length > 0 && (now - lastFetchTime.value) < CACHE_DURATION) {
            return;
        }

        loading.value = true;
        error.value = null;

        try {
            // Use new standardized API method
            const { items } = await apiService.getCourses();

            // Transform courses to ensure proper types (strings to numbers)
            const transformedCourses = items.map(transformCourse);

            courses.value = transformedCourses;
            lastFetchTime.value = now;

            // Update cache
            transformedCourses.forEach(course => {
                coursesCache.value.set(course.slug, course);
            });
        } catch (err: any) {
            error.value = err.message || 'Failed to load courses';
            console.error('Error fetching courses:', err);
        } finally {
            loading.value = false;
        }
    }

    async function fetchCourseBySlug(slug: string, forceRefresh = false): Promise<Course | null> {
        // Check cache first
        if (!forceRefresh && coursesCache.value.has(slug)) {
            currentCourse.value = coursesCache.value.get(slug)!;
            return currentCourse.value;
        }

        loading.value = true;
        error.value = null;

        try {
            // Use new standardized API method
            const courseData = await apiService.getCourse(slug);
            // Transform to ensure proper types
            const course = transformCourse(courseData);
            currentCourse.value = course;
            coursesCache.value.set(slug, course);
            return course;
        } catch (err: any) {
            error.value = err.message || 'Failed to load course';
            console.error('Error fetching course:', err);
            return null;
        } finally {
            loading.value = false;
        }
    }

    function getCourseBySlug(slug: string): Course | undefined {
        return courses.value.find(c => c.slug === slug) || coursesCache.value.get(slug);
    }

    function clearCurrentCourse() {
        currentCourse.value = null;
    }

    function clearCache() {
        coursesCache.value.clear();
        lastFetchTime.value = 0;
    }

    return {
        // State
        courses,
        currentCourse,
        loading,
        error,
        // Getters
        featuredCourse,
        freeCourses,
        paidCourses,
        // Actions
        fetchCourses,
        fetchCourseBySlug,
        getCourseBySlug,
        clearCurrentCourse,
        clearCache,
    };
});
