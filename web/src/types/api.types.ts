// API Types and Interfaces

// Standardized API Response Format
export interface StandardizedListResponse<T> {
    ci_environment: string;
    data: {
        items: T[];
        total: number;
    };
}

export interface StandardizedResponse<T> {
    ci_environment: string;
    data: T;
    message?: string;
    success?: boolean;
}

export interface User {
    id: number;
    username: string;
    email: string;
    first_name?: string;
    last_name?: string;
    role: 'student' | 'instructor' | 'admin';
    is_staff: boolean;
    is_active: boolean;
    profile_picture?: string;
    bio?: string;
    date_joined?: string;
    last_login?: string;
    instructor_rating?: number;
    total_reviews?: number;
    date_of_birth?: string;
    gender?: string;
    contact_number?: string;
    address?: string;
    country?: string;
    highest_qualification?: string;
    institution?: string;
    skills?: string;
    linkedin_profile?: string;
    github_profile?: string;
    website?: string;
}

export interface Category {
    id: number;
    name: string;
    slug: string;
    description?: string;
    image?: string;
}

export interface Lesson {
    id: number;
    course?: number;
    order: number;
    title: string;
    slug: string;
    video_url?: string;
    content?: string;
    is_preview: boolean;
    created_at?: string;
    updated_at?: string;
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
    price: string | number;
    is_free: boolean;
    is_published: boolean;
    instructor: User;
    average_rating: string | number;
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

export interface Enrollment {
    id: number;
    student: User;
    course: Course;
    enrolled_at: string;
    completed_at?: string;
    progress: number;
    is_completed: boolean;
}

export interface AnswerChoice {
    id: number;
    choice_text: string;
    is_correct: boolean;
}

export interface Question {
    id: number;
    quiz: number;
    question_text: string;
    question_type: 'mcq' | 'text';
    order: number;
    choices?: AnswerChoice[];
}

export interface Quiz {
    id: number;
    lesson: number;
    title: string;
    description?: string;
    passing_score: number;
    created_at: string;
    updated_at: string;
}

export interface QuizAttempt {
    id: number;
    student: User;
    quiz: Quiz;
    score: number;
    passed: boolean;
    attempt_number: number;
    submitted_at: string;
}

export interface Review {
    id: number;
    course: Course;
    student: User;
    rating: number;
    comment?: string;
    created_at: string;
    is_approved: boolean;
}

export interface Certificate {
    id: number;
    enrollment: Enrollment;
    unique_id: string;
    issue_date: string;
}

export interface FAQ {
    id: number;
    question: string;
    answer: string;
    category?: Category;
    is_published: boolean;
    order?: number;
    created_at?: string;
    updated_at?: string;
}

export interface LoginCredentials {
    username: string;
    password: string;
}

export interface RegisterData {
    username: string;
    email: string;
    password: string;
    first_name?: string;
    last_name?: string;
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

// Quiz submission types
export interface QuizAnswer {
    question_id: number;
    selected_choice_id?: number;
    text_answer?: string;
}

export interface QuizSubmission {
    quiz_id: number;
    answers: QuizAnswer[];
}
