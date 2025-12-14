import type { Course, Enrollment, User } from '../types/api.types';

/**
 * Transform API response data to ensure correct types
 * Django serializes DecimalField as strings, but frontend expects numbers
 */

export function transformCourse(course: any): Course {
    return {
        ...course,
        // Convert string numbers to actual numbers
        average_rating: course.average_rating ? Number(course.average_rating) : 0,
        price: course.price ? Number(course.price) : 0,
        total_lectures: course.total_lectures ? Number(course.total_lectures) : 0,
        number_of_reviews: course.number_of_reviews ? Number(course.number_of_reviews) : 0,
        total_reviews: course.total_reviews ? Number(course.total_reviews) : 0,

        // Ensure nested objects are properly typed
        instructor: course.instructor ? {
            ...course.instructor,
            instructor_rating: course.instructor.instructor_rating ? Number(course.instructor.instructor_rating) : 0,
            total_reviews: course.instructor.total_reviews ? Number(course.instructor.total_reviews) : 0,
        } : undefined,
    };
}

export function transformCourses(courses: any[]): Course[] {
    return courses.map(transformCourse);
}

export function transformEnrollment(enrollment: any): Enrollment {
    return {
        ...enrollment,
        progress: enrollment.progress ? Number(enrollment.progress) : 0,
        course: enrollment.course ? transformCourse(enrollment.course) : enrollment.course,
    };
}

export function transformEnrollments(enrollments: any[]): Enrollment[] {
    return enrollments.map(transformEnrollment);
}

export function transformUser(user: any): User {
    return {
        ...user,
        instructor_rating: user.instructor_rating ? Number(user.instructor_rating) : undefined,
        total_reviews: user.total_reviews ? Number(user.total_reviews) : undefined,
    };
}

/**
 * Transform paginated API response
 */
export function transformPaginatedResponse<T>(
    response: any,
    transformer: (item: any) => T
): T[] {
    if (response.results && Array.isArray(response.results)) {
        // Paginated response
        return response.results.map(transformer);
    } else if (Array.isArray(response)) {
        // Non-paginated response
        return response.map(transformer);
    } else {
        console.error('Unexpected API response format:', response);
        return [];
    }
}
