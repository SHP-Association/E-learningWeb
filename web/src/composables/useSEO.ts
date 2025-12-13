// SEO Helper Composable using @unhead/vue

import { useHead } from '@unhead/vue';
import type { SEOConfig, CoursePageSEO, ProfilePageSEO } from '../types/seo.types';

// Default SEO configuration
const DEFAULT_SEO: SEOConfig = {
    siteName: 'SHP-Learner',
    locale: 'en_US',
    type: 'website',
    author: 'SHP-Learner Team',
    twitterCard: 'summary_large_image',
    twitterSite: '@SandeshPat007',
    robots: 'index, follow',
};

// Base URL - should match your production URL
const BASE_URL = import.meta.env.VITE_APP_BASE_URL || 'https://shp-learner.com';

/**
 * Main SEO composable for setting page meta tags
 */
export function useSEO(config: SEOConfig = {}) {
    const {
        title = 'SHP-Learner - Online Learning Platform',
        description = 'Explore our wide range of meticulously crafted courses to elevate your skills and career. Learn from expert instructors.',
        keywords = ['online learning', 'courses', 'education', 'e-learning', 'skill development'],
        image = `${BASE_URL}/default-og-image.jpg`,
        url = BASE_URL,
        type = 'website',
        siteName = DEFAULT_SEO.siteName,
        locale = DEFAULT_SEO.locale,
        author = DEFAULT_SEO.author,
        publishedTime,
        modifiedTime,
        section,
        tags = [],
        twitterCard = DEFAULT_SEO.twitterCard,
        twitterSite = DEFAULT_SEO.twitterSite,
        twitterCreator,
        canonical,
        robots = DEFAULT_SEO.robots,
        jsonLd,
    } = config;

    const fullTitle = title.includes('SHP-Learner') ? title : `${title} | SHP-Learner`;
    const fullUrl = url.startsWith('http') ? url : `${BASE_URL}${url}`;
    const canonicalUrl = canonical || fullUrl;

    useHead({
        title: fullTitle,
        meta: [
            // Basic meta tags
            { name: 'description', content: description },
            { name: 'keywords', content: keywords.join(', ') },
            { name: 'author', content: author },
            { name: 'robots', content: robots },

            // Open Graph
            { property: 'og:title', content: fullTitle },
            { property: 'og:description', content: description },
            { property: 'og:type', content: type },
            { property: 'og:url', content: fullUrl },
            { property: 'og:image', content: image },
            { property: 'og:site_name', content: siteName },
            { property: 'og:locale', content: locale },

            // Open Graph - Article specific
            ...(publishedTime ? [{ property: 'article:published_time', content: publishedTime }] : []),
            ...(modifiedTime ? [{ property: 'article:modified_time', content: modifiedTime }] : []),
            ...(section ? [{ property: 'article:section', content: section }] : []),
            ...tags.map(tag => ({ property: 'article:tag', content: tag })),

            // Twitter Card
            { name: 'twitter:card', content: twitterCard },
            { name: 'twitter:site', content: twitterSite },
            { name: 'twitter:title', content: fullTitle },
            { name: 'twitter:description', content: description },
            { name: 'twitter:image', content: image },
            ...(twitterCreator ? [{ name: 'twitter:creator', content: twitterCreator }] : []),

            // Mobile
            { name: 'viewport', content: 'width=device-width, initial-scale=1.0' },
            { name: 'theme-color', content: '#2563eb' }, // Blue color
        ],
        link: [
            { rel: 'canonical', href: canonicalUrl },
            { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
        ],
        ...(jsonLd && {
            script: [
                {
                    type: 'application/ld+json',
                    innerHTML: JSON.stringify(jsonLd),
                },
            ],
        }),
    });
}

/**
 * SEO for Home page
 */
export function useHomeSEO() {
    useSEO({
        title: 'SHP-Learner - Online Learning Platform',
        description: 'Explore our wide range of meticulously crafted courses to elevate your skills and career. Learn from expert instructors in various fields.',
        keywords: ['online courses', 'e-learning', 'skill development', 'education', 'training', 'professional development'],
        url: '/',
        type: 'website',
        jsonLd: {
            '@context': 'https://schema.org',
            '@type': 'EducationalOrganization',
            name: 'SHP-Learner',
            description: 'Online learning platform offering diverse courses for skill development',
            url: BASE_URL,
            logo: `${BASE_URL}/logo.png`,
            sameAs: [
                'https://twitter.com/SandeshPat007',
                'https://www.linkedin.com/in/sandesh-patel07',
                'https://www.instagram.com/sandesh_patel007',
            ],
        },
    });
}

/**
 * SEO for Course Detail page
 */
export function useCourseSEO(course: CoursePageSEO) {
    const price = course.isFree ? 'Free' : `₹${course.price}`;

    useSEO({
        title: course.title,
        description: course.description,
        keywords: [course.category, course.level, 'online course', 'e-learning', course.instructor],
        url: `/course/${course.title.toLowerCase().replace(/\s+/g, '-')}`,
        type: 'article',
        section: course.category,
        tags: [course.category, course.level, course.instructor],
        image: course.image,
        jsonLd: {
            '@context': 'https://schema.org',
            '@type': 'Course',
            name: course.title,
            description: course.description,
            provider: {
                '@type': 'Organization',
                name: 'SHP-Learner',
                sameAs: BASE_URL,
            },
            instructor: {
                '@type': 'Person',
                name: course.instructor,
            },
            offers: {
                '@type': 'Offer',
                price: course.isFree ? '0' : course.price.toString(),
                priceCurrency: 'INR',
                availability: 'https://schema.org/InStock',
            },
            aggregateRating: course.reviewCount > 0 ? {
                '@type': 'AggregateRating',
                ratingValue: course.rating.toString(),
                reviewCount: course.reviewCount.toString(),
            } : undefined,
            educationalLevel: course.level,
            courseMode: 'online',
        },
    });
}

/**
 * SEO for Login page
 */
export function useLoginSEO() {
    useSEO({
        title: 'Login',
        description: 'Log in to your SHP-Learner account to access your courses and continue learning.',
        keywords: ['login', 'sign in', 'account access'],
        url: '/login',
        robots: 'noindex, nofollow', // Don't index login pages
    });
}

/**
 * SEO for Register page
 */
export function useRegisterSEO() {
    useSEO({
        title: 'Register',
        description: 'Create your SHP-Learner account and start learning today. Join thousands of students improving their skills.',
        keywords: ['register', 'sign up', 'create account', 'join'],
        url: '/register',
    });
}

/**
 * SEO for Profile page
 */
export function useProfileSEO(profile: ProfilePageSEO) {
    useSEO({
        title: `${profile.username}'s Profile`,
        description: profile.bio || `View ${profile.username}'s profile on SHP-Learner`,
        keywords: ['profile', 'user', profile.role],
        url: `/profile`,
        type: 'profile',
        image: profile.image,
        robots: 'noindex, nofollow', // Don't index user profiles
        jsonLd: {
            '@context': 'https://schema.org',
            '@type': 'Person',
            name: profile.username,
            description: profile.bio,
            image: profile.image,
        },
    });
}

/**
 * SEO for FAQ page
 */
export function useFAQSEO() {
    useSEO({
        title: 'Frequently Asked Questions',
        description: 'Find answers to common questions about SHP-Learner courses, enrollment, and platform features.',
        keywords: ['FAQ', 'help', 'questions', 'support', 'answers'],
        url: '/faq',
        type: 'website',
    });
}

/**
 * SEO for 404 page
 */
export function useNotFoundSEO() {
    useSEO({
        title: 'Page Not Found',
        description: 'The page you are looking for could not be found.',
        url: '/404',
        robots: 'noindex, nofollow',
    });
}

/**
 * Helper to generate course structured data
 */
export function generateCourseStructuredData(course: any) {
    return {
        '@context': 'https://schema.org',
        '@type': 'Course',
        name: course.title,
        description: course.description,
        provider: {
            '@type': 'Organization',
            name: 'SHP-Learner',
            sameAs: BASE_URL,
        },
        instructor: {
            '@type': 'Person',
            name: course.instructor?.username || 'Unknown',
        },
        offers: {
            '@type': 'Offer',
            price: course.is_free ? '0' : course.price?.toString() || '0',
            priceCurrency: 'INR',
            availability: 'https://schema.org/InStock',
        },
        aggregateRating: course.number_of_reviews > 0 ? {
            '@type': 'AggregateRating',
            ratingValue: course.average_rating?.toString() || '0',
            reviewCount: course.number_of_reviews?.toString() || '0',
        } : undefined,
        educationalLevel: course.level,
        courseMode: 'online',
    };
}
