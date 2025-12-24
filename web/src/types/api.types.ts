// API Types and Interfaces

export interface User {
    id: number;
    username: string;
    email: string;
    role: 'student' | 'instructor';
    is_staff: boolean;
    profile_picture?: string;
    bio?: string;
    date_joined?: string;
    instructor_rating?: number;
    total_reviews?: number;
}

export interface Course {
    id: number;
    slug: string;
    title: string;
    description: string;
    short_description?: string;
    thumbnail?: string;
    category?: Category;
    level: 'beginner' | 'intermediate' | 'advanced';
    price: number;
    is_free: boolean;
    is_published: boolean;
    instructor: User;
    average_rating: number;
    number_of_reviews: number;
    total_lectures: number;
    duration?: string;
    what_you_will_learn?: string;
    requirements?: string;
    target_audience?: string;
    promo_video_url?: string;
    lessons?: Lesson[];
    created_at: string;
    updated_at: string;
}

export interface Category {
    id: number;
    name: string;
    image?: string;
}

export interface Lesson {
    id: number;
    order: number;
    title: string;
    video_url?: string;
    content?: string;
}

export interface Enrollment {
    id: number;
    student: User;
    course: Course;
    enrolled_at: string;
    progress: number;
    completed: boolean;
}

export interface FAQ {
    id: number;
    question: string;
    answer: string;
    category?: string;
    order?: number;
}

export interface LoginCredentials {
    username: string;
    password: string;
}

export interface RegisterData {
    username: string;
    email: string;
    password: string;
    role: 'student' | 'instructor';
}

export interface ApiError {
    message: string;
    errors?: Record<string, string[]>;
    status?: number;
}

export interface ApiResponse<T> {
    data?: T;
    error?: ApiError;
    success: boolean;
}
