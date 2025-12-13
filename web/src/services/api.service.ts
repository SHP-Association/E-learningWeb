// Centralized API Service

const BACKEND_URL = import.meta.env.VITE_APP_BACKEND_URL || 'http://localhost:8000';

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
    private baseURL: string;

    constructor(baseURL: string) {
        this.baseURL = baseURL;
    }

    private async request<T>(
        endpoint: string,
        options: RequestOptions = {}
    ): Promise<T> {
        const { skipAuth, ...fetchOptions } = options;

        const headers: HeadersInit = {
            'Content-Type': 'application/json',
            ...fetchOptions.headers,
        };

        // Add CSRF token for non-GET requests
        if (fetchOptions.method && fetchOptions.method !== 'GET') {
            const csrfToken = getCookie('csrftoken');
            if (csrfToken) {
                headers['X-CSRFToken'] = csrfToken;
            }
        }

        const config: RequestInit = {
            ...fetchOptions,
            headers,
            credentials: 'include', // Always include credentials for session management
        };

        try {
            const response = await fetch(`${this.baseURL}${endpoint}`, config);

            // Handle non-JSON responses
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

    async post<T>(
        endpoint: string,
        data?: unknown,
        options?: RequestOptions
    ): Promise<T> {
        return this.request<T>(endpoint, {
            ...options,
            method: 'POST',
            body: data ? JSON.stringify(data) : undefined,
        });
    }

    async put<T>(
        endpoint: string,
        data?: unknown,
        options?: RequestOptions
    ): Promise<T> {
        return this.request<T>(endpoint, {
            ...options,
            method: 'PUT',
            body: data ? JSON.stringify(data) : undefined,
        });
    }

    async patch<T>(
        endpoint: string,
        data?: unknown,
        options?: RequestOptions
    ): Promise<T> {
        return this.request<T>(endpoint, {
            ...options,
            method: 'PATCH',
            body: data ? JSON.stringify(data) : undefined,
        });
    }

    async delete<T>(endpoint: string, options?: RequestOptions): Promise<T> {
        return this.request<T>(endpoint, { ...options, method: 'DELETE' });
    }
}

// Export singleton instance
export const apiService = new ApiService(BACKEND_URL);

// Export helper function for CSRF token (for forms that need it)
export { getCookie };
