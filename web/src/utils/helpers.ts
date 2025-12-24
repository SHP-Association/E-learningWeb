// Security and Utility Functions

/**
 * Sanitize HTML to prevent XSS attacks
 */
export function sanitizeHtml(html: string): string {
    const div = document.createElement('div');
    div.textContent = html;
    return div.innerHTML;
}

/**
 * Sanitize user input for display
 */
export function sanitizeInput(input: string): string {
    return input
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/"/g, '&quot;')
        .replace(/'/g, '&#x27;')
        .replace(/\//g, '&#x2F;');
}

/**
 * Debounce function for search/filter inputs
 */
export function debounce<T extends (...args: any[]) => any>(
    fn: T,
    delay: number
): (...args: Parameters<T>) => void {
    let timeoutId: ReturnType<typeof setTimeout>;
    return (...args: Parameters<T>) => {
        clearTimeout(timeoutId);
        timeoutId = setTimeout(() => fn(...args), delay);
    };
}

/**
 * Handle API errors consistently
 */
export function handleApiError(error: unknown): string {
    if (error instanceof Error) {
        // Don't expose sensitive error details
        if (error.message.includes('Network')) {
            return 'Network error. Please check your connection.';
        }
        if (error.message.includes('404')) {
            return 'Resource not found.';
        }
        if (error.message.includes('401') || error.message.includes('403')) {
            return 'You are not authorized to perform this action.';
        }
        return 'An error occurred. Please try again.';
    }
    return 'An unexpected error occurred.';
}

/**
 * Validate email format
 */
export function isValidEmail(email: string): boolean {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
}

/**
 * Check password strength
 */
export function getPasswordStrength(password: string): {
    score: number;
    feedback: string;
} {
    let score = 0;
    const feedback: string[] = [];

    if (password.length >= 8) score++;
    else feedback.push('At least 8 characters');

    if (/[a-z]/.test(password)) score++;
    else feedback.push('Lowercase letter');

    if (/[A-Z]/.test(password)) score++;
    else feedback.push('Uppercase letter');

    if (/\d/.test(password)) score++;
    else feedback.push('Number');

    if (/[@$!%*?&]/.test(password)) score++;
    else feedback.push('Special character');

    const strengthLabels = ['Very Weak', 'Weak', 'Fair', 'Good', 'Strong'];
    return {
        score,
        feedback: feedback.length > 0
            ? `Missing: ${feedback.join(', ')}`
            : strengthLabels[score - 1] || 'Very Weak',
    };
}

/**
 * Format currency
 */
export function formatCurrency(amount: number, currency = 'INR'): string {
    return new Intl.NumberFormat('en-IN', {
        style: 'currency',
        currency,
        minimumFractionDigits: 0,
    }).format(amount);
}

/**
 * Truncate text
 */
export function truncateText(text: string, maxLength: number): string {
    if (text.length <= maxLength) return text;
    return text.substring(0, maxLength).trim() + '...';
}
