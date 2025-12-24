// Application Constants

export const API_CONFIG = {
    TIMEOUT: 30000,
    RETRY_ATTEMPTS: 3,
    CACHE_DURATION: 5 * 60 * 1000, // 5 minutes
} as const;

export const VALIDATION = {
    MIN_PASSWORD_LENGTH: 8,
    MIN_USERNAME_LENGTH: 3,
    MAX_USERNAME_LENGTH: 30,
    MAX_BIO_LENGTH: 500,
    PASSWORD_REGEX: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/,
} as const;

export const ROUTES = {
    HOME: '/',
    LOGIN: '/login',
    REGISTER: '/register',
    PROFILE: '/profile',
    FAQ: '/faq',
    COURSE_DETAIL: '/course/:slug',
} as const;

export const STORAGE_KEYS = {
    USER: 'currentUser',
    ENROLLMENTS: 'mockEnrollments',
    THEME: 'theme',
} as const;

export const SESSION = {
    TIMEOUT: 30 * 60 * 1000, // 30 minutes
    WARNING_TIME: 5 * 60 * 1000, // 5 minutes before timeout
} as const;

export const BREAKPOINTS = {
    SM: 640,
    MD: 768,
    LG: 1024,
    XL: 1280,
    '2XL': 1536,
} as const;

export const TOUCH_TARGET_SIZE = {
    MIN_WIDTH: 44,
    MIN_HEIGHT: 44,
} as const;
