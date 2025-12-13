<template>
  <div class="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-4xl mx-auto">
      <h1 class="text-5xl font-extrabold text-center text-blue-800 mb-4">
        Frequently Asked Questions
      </h1>
      <p class="text-xl text-center text-gray-600 mb-10">
        Find answers to the most common questions about our platform and courses.
      </p>

      <div class="mb-10 p-6 bg-white rounded-xl shadow-lg border border-gray-100">
        <label for="search" class="sr-only">Search FAQs</label>
        <div class="relative">
          <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <input
            id="search"
            type="text"
            v-model="searchTerm"
            placeholder="Search questions..."
            class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-xl text-lg focus:ring-blue-500 focus:border-blue-500 transition"
          />
        </div>
      </div>

      <div class="space-y-4">
        <div v-for="faq in filteredFaqs" :key="faq.id" class="bg-white rounded-xl shadow-lg border border-gray-100 overflow-hidden transition-all duration-300">
          <button
            class="w-full text-left p-6 flex justify-between items-center focus:outline-none hover:bg-blue-50 transition focus:ring-2 focus:ring-blue-500 rounded-t-xl"
            @click="toggleFaq(faq.id)"
            :id="`faq-question-${faq.id}`"
            :aria-expanded="openFaq === faq.id"
            :aria-controls="`faq-answer-${faq.id}`"
            :aria-label="`Toggle answer for: ${faq.question}`"
          >
            <span class="text-xl font-semibold text-blue-800">{{ faq.question }}</span>
            <svg
              class="h-6 w-6 text-blue-500 transform transition-transform duration-300"
              :class="{ 'rotate-180': openFaq === faq.id }"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              aria-hidden="true"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </button>

          <div
            v-show="openFaq === faq.id"
            :id="`faq-answer-${faq.id}`"
            class="p-6 pt-0 text-gray-700 text-lg border-t border-gray-100"
            role="region"
            :aria-labelledby="`faq-question-${faq.id}`"
          >
            <!-- Safe text rendering instead of v-html to prevent XSS -->
            <p class="whitespace-pre-line">{{ faq.answer }}</p>
          </div>
        </div>
      </div>

      <div class="mt-12 p-8 bg-blue-50 rounded-xl shadow-xl border border-blue-200">
        <h2 class="text-3xl font-extrabold text-blue-800 mb-4 text-center">Can't find your answer?</h2>
        <p class="text-center text-lg text-gray-700 mb-6">Ask us directly and we'll get back to you soon.</p>
        
        <form @submit.prevent="handleSubmitQuestion" class="max-w-2xl mx-auto">
          <div class="form-group mb-4">
            <label for="question" class="sr-only">Your Question</label>
            <textarea
              id="question"
              rows="4"
              placeholder="Type your question here..."
              class="w-full resize-y border border-blue-300 rounded-lg p-3 text-base focus:border-blue-500 focus:ring-2 focus:ring-blue-300 transition"
              v-model="question"
              required
            ></textarea>
          </div>
          <button
            type="submit"
            class="bg-blue-900 text-white hover:bg-blue-700 px-8 py-3 rounded-lg font-semibold text-lg shadow transition w-full"
            :disabled="isSubmitting"
          >
            {{ isSubmitting ? 'Submitting...' : 'Submit Question' }}
          </button>
        </form>
        
        <div v-if="successMessage" class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded-lg mt-4 font-medium transition">
          {{ successMessage }}
        </div>
        <div v-if="errorMessage" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg mt-4 font-medium transition">
          {{ errorMessage }}
        </div>
      </div>

      <div class="feedback-btn text-center mt-10">
        <a
          href="mailto:patelbr5118s@gmail.com?subject=Website Feedback"
          class="inline-block bg-green-500 text-white hover:bg-green-600 px-8 py-3 rounded-lg font-semibold text-lg shadow transition duration-200"
        >
          Send us Feedback
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useFAQSEO } from '../composables/useSEO';
import { apiService } from '../services/api.service';
import { sanitizeInput } from '../utils/helpers';

// Set SEO for FAQ page
useFAQSEO();

// Type for a single FAQ item
interface FaqItem {
  id: string | number;
  question: string;
  answer: string;
}

// Static FAQs
const staticFaqs: FaqItem[] = [
  { id: 'static-1', question: 'What is this platform about?', answer: 'Our platform provides high-quality online courses across various domains to help you learn and grow your skills.' },
  { id: 'static-2', question: 'How do I buy a course?', answer: 'Choose a course, add it to your cart, and complete the checkout process using UPI, debit/credit cards, or wallets.' },
  { id: 'static-3', question: 'Do you offer certificates?', answer: 'Yes, we provide certificates for most paid courses after successful completion.' },
  { id: 'static-4', question: 'Can I access the course after purchase forever?', answer: 'Yes, once you purchase a course, you get lifetime access unless specified otherwise.' },
  { id: 'static-5', question: 'Is there a mobile app?', answer: 'We are working on our mobile app. Meanwhile, you can access all features through your mobile browser.' },
  { id: 'static-6', question: 'How can I get support?', answer: 'You can email us at patelbr5118s@gmail.com or use the contact form on our website.' },
];

// State variables
const faqs = ref<FaqItem[]>([]);
const searchTerm = ref('');
const question = ref('');
const openFaq = ref<string | number | null>(null);
const isSubmitting = ref(false);
const successMessage = ref('');
const errorMessage = ref('');

// Fetch FAQs from API on mount
onMounted(async () => {
  try {
    const data = await apiService.get<FaqItem[]>('/api/faqs/');
    faqs.value = data;
  } catch (error) {
    console.error('Failed to load FAQs:', error);
    faqs.value = [];
  }
});

// Combine static and dynamic FAQs
const allFaqs = computed(() => [...staticFaqs, ...faqs.value]);

// Filter FAQs based on search term
const filteredFaqs = computed(() => {
  if (!searchTerm.value) return allFaqs.value;
  const term = searchTerm.value.toLowerCase();
  return allFaqs.value.filter(faq =>
    faq.question.toLowerCase().includes(term) || faq.answer.toLowerCase().includes(term)
  );
});

// Toggle FAQ accordion
const toggleFaq = (id: string | number) => {
  openFaq.value = openFaq.value === id ? null : id;
};

// Handle question submission
const handleSubmitQuestion = async () => {
  if (!question.value.trim()) return;

  // Sanitize user input before sending
  const sanitizedQuestion = sanitizeInput(question.value.trim());

  isSubmitting.value = true;
  successMessage.value = '';
  errorMessage.value = '';

  try {
    await apiService.post('/api/contact/', { question: sanitizedQuestion });
    successMessage.value = 'Thank you for your question! We will get back to you shortly.';
    question.value = '';
  } catch (error: any) {
    errorMessage.value = error.message || 'Failed to submit question. Please try again.';
  } finally {
    isSubmitting.value = false;
  }
};
</script>