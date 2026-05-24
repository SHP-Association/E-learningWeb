// Comprehensive API Service with dedicated methods for all endpoints
// In dev, Vite proxies /api requests to the backend (see vite.config.js).
// In production, a reverse proxy (nginx, etc.) routes /api to the backend.

export class ApiError extends Error {
    constructor(
        message: string,
        public status?: number,
        public errors?: Record<string, string[]>
    ) {
        super(message);
        this.name = 'ApiError';
    }
}

// Helper to get CSRF token from cookies
function getCookie(name: string): string | null {
    let cookieValue = null;
    if (document.cookie && document.cookie !== '') {
        const cookies = document.cookie.split(';');
        for (let i = 0; i < cookies.length; i++) {
            const cookie = cookies[i].trim();
            if (cookie.substring(0, name.length + 1) === (name + '=')) {
                cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
                break;
            }
        }
    }
    return cookieValue;
}

interface RequestOptions extends RequestInit {
    skipAuth?: boolean;
}

class ApiService {
    public baseURL: string;

    constructor(baseURL: string) {
        this.baseURL = baseURL;
    }

    // ==================== Core Request Methods ====================

    private async request<T>(
        endpoint: string,
        options: RequestOptions = {}
    ): Promise<T> {
        const { skipAuth, ...fetchOptions } = options;

        const headers: HeadersInit = {
            ...fetchOptions.headers,
        };

        // Only set Content-Type to JSON if body is not FormData
        if (!(fetchOptions.body instanceof FormData)) {
            if (!headers['Content-Type' as keyof HeadersInit]) {
                (headers as any)['Content-Type'] = 'application/json';
            }
        }

        // Add CSRF token for non-GET requests (checking both Echo _csrf and legacy Django csrftoken)
        if (fetchOptions.method && fetchOptions.method !== 'GET') {
            const csrfToken = getCookie('_csrf') || getCookie('csrftoken');
            if (csrfToken) {
                (headers as any)['X-CSRF-Token'] = csrfToken;
                (headers as any)['X-CSRFToken'] = csrfToken;
            }
        }

        const config: RequestInit = {
            ...fetchOptions,
            headers,
            credentials: 'include',
        };

        try {
            const response = await fetch(`${this.baseURL}${endpoint}`, config);

            const contentType = response.headers.get('content-type');
            if (!contentType || !contentType.includes('application/json')) {
                if (!response.ok) {
                    throw new ApiError(
                        `HTTP error! status: ${response.status}`,
                        response.status
                    );
                }
                return {} as T;
            }

            const data = await response.json();

            if (!response.ok) {
                throw new ApiError(
                    data.error || data.message || data.detail || 'An error occurred',
                    response.status,
                    data.errors || data
                );
            }

            return data;
        } catch (error) {
            if (error instanceof ApiError) {
                throw error;
            }
            throw new ApiError(
                error instanceof Error ? error.message : 'Network error occurred'
            );
        }
    }

    async get<T>(endpoint: string, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, { ...options, method: 'GET' });
    }

    async post<T>(endpoint: string, data?: unknown, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, {
            ...options,
            method: 'POST',
            body: data ? JSON.stringify(data) : undefined,
        });
    }

    async put<T>(endpoint: string, data?: unknown, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, {
            ...options,
            method: 'PUT',
            body: data ? JSON.stringify(data) : undefined,
        });
    }

    async patch<T>(endpoint: string, data?: unknown, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, {
            ...options,
            method: 'PATCH',
            body: data ? JSON.stringify(data) : undefined,
        });
    }

    async delete<T>(endpoint: string, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, { ...options, method: 'DELETE' });
    }

    // ==================== Response Unwrapping Helpers ====================

    async getList<T>(endpoint: string, options?: RequestOptions): Promise<{ items: T[]; total: number }> {
        const response = await this.request<any>(endpoint, { ...options, method: 'GET' });
        if (Array.isArray(response)) {
            return { items: response, total: response.length };
        }
        if (response.data && Array.isArray(response.data)) {
            return { items: response.data, total: response.data.length };
        }
        if (response.items && Array.isArray(response.items)) {
            return { items: response.items, total: response.total || response.items.length };
        }
        return { items: [], total: 0 };
    }

    async getOne<T>(endpoint: string, options?: RequestOptions): Promise<T> {
        const response = await this.request<any>(endpoint, { ...options, method: 'GET' });
        if (response.data && typeof response.data === 'object' && !Array.isArray(response.data)) {
            return response.data as T;
        }
        return response as T;
    }

    // ==================== Authentication APIs ====================

    async login(credentials: { email: string; password: string }): Promise<any> {
        // Echo supports both /api/auth/login and legacy /api/login (non-trailing slash)
        return this.post('/api/auth/login', credentials);
    }

    async logout(): Promise<any> {
        return this.post('/api/auth/logout', {});
    }

    async register(data: {
        username: string;
        email: string;
        password: string;
    }): Promise<any> {
        return this.post('/api/auth/register', data);
    }

    async verifyOTP(data: { email: string; otp: string }): Promise<any> {
        return this.post('/api/auth/verify-otp', data);
    }

    async getSignupConfig(): Promise<{ instructor_contact_email?: string }> {
        return this.get('/api/auth/signup-config');
    }

    async requestPasswordReset(email: string): Promise<any> {
        return this.post('/api/auth/password-reset', { email });
    }

    async confirmPasswordReset(uid: string, token: string, newPassword: string): Promise<any> {
        return this.post('/api/auth/password-reset/confirm', {
            uid,
            token,
            new_password: newPassword,
        });
    }

    // ==================== User & Profile APIs ====================

    async getCurrentUser(): Promise<any> {
        // Go backend maps student profile to /api/profile
        const response = await this.get<any>('/api/profile');
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async getProfile(): Promise<any> {
        return this.getCurrentUser();
    }

    async updateUser(userId: number, data: any): Promise<any> {
        // Fallback profile update mapping
        return this.patch<any>('/api/profile', data);
    }

    async uploadProfilePicture(userId: number, formData: FormData): Promise<any> {
        return this.request<any>('/api/profile', {
            method: 'PATCH',
            body: formData
        });
    }

    async submitOnboarding(data: {
        first_name: string;
        last_name: string;
        contact_number: string;
        country: string;
    }): Promise<any> {
        return this.patch('/api/profile/onboarding', data);
    }

    // ==================== Course APIs ====================

    async getCourses(params?: { category?: string; level?: string; search?: string }): Promise<{ items: any[]; total: number }> {
        const queryParams = new URLSearchParams();
        if (params?.category) queryParams.append('category', params.category);
        if (params?.level) queryParams.append('level', params.level);
        if (params?.search) queryParams.append('search', params.search);
        const query = queryParams.toString();
        return this.getList(`/api/courses${query ? '?' + query : ''}`);
    }

    async getCourse(id: number | string): Promise<any> {
        // Mapped to /api/courses/{id}
        return this.getOne(`/api/courses/${id}`);
    }

    // ==================== Category APIs ====================

    async getCategories(): Promise<{ items: any[]; total: number }> {
        return this.getList('/api/categories');
    }

    async getCategory(id: number): Promise<any> {
        return this.getOne(`/api/categories/${id}`);
    }

    // ==================== Lesson APIs ====================

    async getLesson(id: number): Promise<any> {
        return this.getOne(`/api/lessons/${id}`);
    }

    // ==================== Enrollment APIs ====================

    async getEnrollments(): Promise<{ items: any[]; total: number }> {
        return this.getList('/api/enrollments');
    }

    async createEnrollment(data: { course_slug: string }): Promise<any> {
        // Call the newly fixed POST /api/enroll/{slug}
        return this.post(`/api/enroll/${data.course_slug}`, {});
    }

    async enrollInCourse(slug: string): Promise<any> {
        return this.post(`/api/enroll/${slug}`, {});
    }

    async checkEnrollment(courseId: number): Promise<boolean> {
        try {
            const user = await this.getCurrentUser();
            const enrollments = user?.edges?.enrollments || [];
            return enrollments.some(
                (enrollment: any) => enrollment?.edges?.course?.id === courseId
            );
        } catch {
            return false;
        }
    }

    // ==================== Quiz APIs ====================

    async getQuizzes(): Promise<{ items: any[]; total: number }> {
        return this.getList('/api/quizzes');
    }

    async getQuiz(quizId: number): Promise<any> {
        return this.getOne(`/api/quizzes/${quizId}`);
    }

    async getQuestions(quizId: number): Promise<{ items: any[]; total: number }> {
        return this.getList(`/api/quizzes/${quizId}/questions`);
    }

    async submitQuiz(quizId: number, answers: any): Promise<any> {
        return this.post(`/api/quizzes/${quizId}/submit`, answers);
    }

    // ==================== Review APIs ====================

    async getReviews(courseId: number): Promise<{ items: any[]; total: number }> {
        return this.getList(`/api/courses/${courseId}/reviews`);
    }

    async createReview(data: { course: number; rating: number; comment?: string }): Promise<any> {
        return this.post(`/api/courses/${data.course}/reviews`, {
            rating: data.rating,
            comment: data.comment
        });
    }

    // ==================== FAQ APIs ====================

    async getFAQs(): Promise<{ items: any[]; total: number }> {
        return this.getList('/api/faqs');
    }
}

// Export singleton instance — uses relative paths; proxied by Vite in dev
export const apiService = new ApiService('');

// Export helper function for CSRF token
export { getCookie };
