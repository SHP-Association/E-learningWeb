// Comprehensive API Service with dedicated methods for all endpoints

const BACKEND_URL = import.meta.env.VITE_APP_BACKEND_URL || 'http://localhost:8001';

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

        // Add CSRF token for non-GET requests
        if (fetchOptions.method && fetchOptions.method !== 'GET') {
            const csrfToken = getCookie('csrftoken');
            if (csrfToken) {
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
                    data.message || data.detail || 'An error occurred',
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
        if (response.data && 'items' in response.data && 'total' in response.data) {
            return response.data;
        }
        return { items: Array.isArray(response) ? response : [], total: 0 };
    }

    async getOne<T>(endpoint: string, options?: RequestOptions): Promise<T> {
        const response = await this.request<any>(endpoint, { ...options, method: 'GET' });
        if (response.data && typeof response.data === 'object' && !Array.isArray(response.data)) {
            return response.data as T;
        }
        return response as T;
    }

    // ==================== Authentication APIs ====================

    async login(credentials: { username: string; password: string }): Promise<any> {
        return this.post('/api/login/', credentials);
    }

    async logout(): Promise<any> {
        return this.post('/api/logout/', {});
    }

    async register(data: {
        username: string;
        email: string;
        password: string;
        first_name?: string;
        last_name?: string;
        role: 'student' | 'instructor';
    }): Promise<any> {
        return this.post('/api/register/', data);
    }

    async requestPasswordReset(email: string): Promise<any> {
        return this.post('/api/password_reset/', { email });
    }

    async confirmPasswordReset(uid: string, token: string, newPassword: string): Promise<any> {
        return this.post(`/api/password_reset/${uid}/${token}/`, { new_password: newPassword });
    }

    // ==================== User APIs ====================

    async getUsers(params?: { page?: number; search?: string }): Promise<{ items: any[]; total: number }> {
        const queryParams = new URLSearchParams();
        if (params?.page) queryParams.append('page', params.page.toString());
        if (params?.search) queryParams.append('search', params.search);
        const query = queryParams.toString();
        return this.getList(`/api/users/${query ? '?' + query : ''}`);
    }

    async getUser(userId: number): Promise<any> {
        return this.getOne(`/api/users/${userId}/`);
    }

    async updateUser(userId: number, data: any): Promise<any> {
        const isFormData = data instanceof FormData;
        const response = await this.patch<any>(`/api/users/${userId}/`, isFormData ? data : data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async uploadProfilePicture(userId: number, formData: FormData): Promise<any> {
        const response = await this.request<any>(`/api/users/${userId}/`, {
            method: 'PATCH',
            body: formData
        });
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async deleteUser(userId: number): Promise<any> {
        return this.delete(`/api/users/${userId}/`);
    }

    async getCurrentUser(): Promise<any> {
        const response = await this.get<any>('/api/users/me/');
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    // ==================== Course APIs ====================

    async getCourses(params?: { category?: string; level?: string; search?: string }): Promise<{ items: any[]; total: number }> {
        const queryParams = new URLSearchParams();
        if (params?.category) queryParams.append('category', params.category);
        if (params?.level) queryParams.append('level', params.level);
        if (params?.search) queryParams.append('search', params.search);
        const query = queryParams.toString();
        return this.getList(`/api/courses/${query ? '?' + query : ''}`);
    }

    async getCourse(slug: string): Promise<any> {
        return this.getOne(`/api/courses/${slug}/`);
    }

    async createCourse(data: any): Promise<any> {
        const response = await this.post<any>('/api/courses/', data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async updateCourse(slug: string, data: any): Promise<any> {
        const response = await this.patch<any>(`/api/courses/${slug}/`, data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async deleteCourse(slug: string): Promise<any> {
        return this.delete(`/api/courses/${slug}/`);
    }

    // ==================== Category APIs ====================

    async getCategories(): Promise<{ items: any[]; total: number }> {
        return this.getList('/api/categories/');
    }

    async getCategory(id: number): Promise<any> {
        return this.getOne(`/api/categories/${id}/`);
    }

    async createCategory(data: { name: string; description?: string; image?: string }): Promise<any> {
        const response = await this.post<any>('/api/categories/', data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async updateCategory(id: number, data: any): Promise<any> {
        const response = await this.patch<any>(`/api/categories/${id}/`, data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async deleteCategory(id: number): Promise<any> {
        return this.delete(`/api/categories/${id}/`);
    }

    // ==================== Lesson APIs ====================

    async getLessons(courseSlug?: string): Promise<{ items: any[]; total: number }> {
        const endpoint = courseSlug ? `/api/lessons/?course=${courseSlug}` : '/api/lessons/';
        return this.getList(endpoint);
    }

    async getLesson(id: number): Promise<any> {
        return this.getOne(`/api/lessons/${id}/`);
    }

    async createLesson(data: any): Promise<any> {
        const response = await this.post<any>('/api/lessons/', data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async updateLesson(id: number, data: any): Promise<any> {
        const response = await this.patch<any>(`/api/lessons/${id}/`, data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async deleteLesson(id: number): Promise<any> {
        return this.delete(`/api/lessons/${id}/`);
    }

    // ==================== Enrollment APIs ====================

    async getEnrollments(params?: { student?: number; course?: string }): Promise<{ items: any[]; total: number }> {
        const queryParams = new URLSearchParams();
        if (params?.student) queryParams.append('student', params.student.toString());
        if (params?.course) queryParams.append('course', params.course);
        const query = queryParams.toString();
        return this.getList(`/api/enrollments/${query ? '?' + query : ''}`);
    }

    async getEnrollment(id: number): Promise<any> {
        return this.getOne(`/api/enrollments/${id}/`);
    }

    async createEnrollment(data: { course_slug: string }): Promise<any> {
        const response = await this.post<any>('/api/enrollments/', data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async updateEnrollment(id: number, data: { progress?: number; completed?: boolean }): Promise<any> {
        const response = await this.patch<any>(`/api/enrollments/${id}/`, data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async deleteEnrollment(id: number): Promise<any> {
        return this.delete(`/api/enrollments/${id}/`);
    }

    async checkEnrollment(courseId: number): Promise<boolean> {
        try {
            const { items } = await this.getEnrollments();
            return items.some((enrollment: any) => enrollment.course.id === courseId);
        } catch {
            return false;
        }
    }

    // ==================== Quiz APIs ====================

    async getQuizzes(lessonId?: number): Promise<{ items: any[]; total: number }> {
        const endpoint = lessonId ? `/api/quizzes/?lesson=${lessonId}` : '/api/quizzes/';
        return this.getList(endpoint);
    }

    async getQuiz(quizId: number): Promise<any> {
        return this.getOne(`/api/quizzes/${quizId}/`);
    }

    async createQuiz(data: any): Promise<any> {
        const response = await this.post<any>('/api/quizzes/', data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async updateQuiz(quizId: number, data: any): Promise<any> {
        const response = await this.patch<any>(`/api/quizzes/${quizId}/`, data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async deleteQuiz(quizId: number): Promise<any> {
        return this.delete(`/api/quizzes/${quizId}/`);
    }

    async submitQuiz(quizId: number, answers: any): Promise<any> {
        const response = await this.post<any>(`/api/quizzes/${quizId}/submit/`, answers);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    // ==================== Question APIs ====================

    async getQuestions(quizId?: number): Promise<{ items: any[]; total: number }> {
        const endpoint = quizId ? `/api/questions/?quiz=${quizId}` : '/api/questions/';
        return this.getList(endpoint);
    }

    async getQuestion(questionId: number): Promise<any> {
        return this.getOne(`/api/questions/${questionId}/`);
    }

    async createQuestion(data: any): Promise<any> {
        const response = await this.post<any>('/api/questions/', data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async updateQuestion(questionId: number, data: any): Promise<any> {
        const response = await this.patch<any>(`/api/questions/${questionId}/`, data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async deleteQuestion(questionId: number): Promise<any> {
        return this.delete(`/api/questions/${questionId}/`);
    }

    // ==================== Review APIs ====================

    async getReviews(courseId?: number): Promise<{ items: any[]; total: number }> {
        const endpoint = courseId ? `/api/reviews/?course=${courseId}` : '/api/reviews/';
        return this.getList(endpoint);
    }

    async getReview(reviewId: number): Promise<any> {
        return this.getOne(`/api/reviews/${reviewId}/`);
    }

    async createReview(data: { course: number; rating: number; comment?: string }): Promise<any> {
        const response = await this.post<any>('/api/reviews/', data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async updateReview(reviewId: number, data: any): Promise<any> {
        const response = await this.patch<any>(`/api/reviews/${reviewId}/`, data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async deleteReview(reviewId: number): Promise<any> {
        return this.delete(`/api/reviews/${reviewId}/`);
    }

    // ==================== Certificate APIs ====================

    async getCertificates(enrollmentId?: number): Promise<{ items: any[]; total: number }> {
        const endpoint = enrollmentId ? `/api/certificates/?enrollment=${enrollmentId}` : '/api/certificates/';
        return this.getList(endpoint);
    }

    async getCertificate(certificateId: number): Promise<any> {
        return this.getOne(`/api/certificates/${certificateId}/`);
    }

    async createCertificate(data: { enrollment: number }): Promise<any> {
        const response = await this.post<any>('/api/certificates/', data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    // ==================== FAQ APIs ====================

    async getFAQs(category?: string): Promise<{ items: any[]; total: number }> {
        const endpoint = category ? `/api/faqs/?category=${category}` : '/api/faqs/';
        return this.getList(endpoint);
    }

    async getFAQ(faqId: number): Promise<any> {
        return this.getOne(`/api/faqs/${faqId}/`);
    }

    async createFAQ(data: { question: string; answer: string; category?: number }): Promise<any> {
        const response = await this.post<any>('/api/faqs/', data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async updateFAQ(faqId: number, data: any): Promise<any> {
        const response = await this.patch<any>(`/api/faqs/${faqId}/`, data);
        return response.data && typeof response.data === 'object' ? response.data : response;
    }

    async deleteFAQ(faqId: number): Promise<any> {
        return this.delete(`/api/faqs/${faqId}/`);
    }
}

// Export singleton instance
export const apiService = new ApiService(BACKEND_URL);

// Export helper function for CSRF token
export { getCookie };
