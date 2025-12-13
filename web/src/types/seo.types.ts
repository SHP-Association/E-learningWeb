// SEO Types and Interfaces

export interface SEOConfig {
    title?: string;
    description?: string;
    keywords?: string[];
    author?: string;
    image?: string;
    url?: string;
    type?: 'website' | 'article' | 'product' | 'profile';
    siteName?: string;
    locale?: string;
    publishedTime?: string;
    modifiedTime?: string;
    section?: string;
    tags?: string[];
    twitterCard?: 'summary' | 'summary_large_image' | 'app' | 'player';
    twitterSite?: string;
    twitterCreator?: string;
    canonical?: string;
    robots?: string;
    jsonLd?: Record<string, any>;
}

export interface CoursePageSEO {
    title: string;
    description: string;
    instructor: string;
    category: string;
    level: string;
    price: number;
    isFree: boolean;
    rating: number;
    reviewCount: number;
    image?: string;
}

export interface ProfilePageSEO {
    username: string;
    bio?: string;
    role: string;
    image?: string;
}
