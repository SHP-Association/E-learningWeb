<template>
  <div class="min-h-screen bg-gray-50 px-4 py-12 sm:px-6 lg:px-8">
    <div class="mx-auto max-w-4xl">
      <h1 class="mb-4 text-center text-5xl font-extrabold text-blue-800">
        Frequently Asked Questions
      </h1>
      <p class="mb-10 text-center text-xl text-gray-600">
        Find answers to the most common questions about our platform and courses.
      </p>

      <AppCard size="lg" padding="md" elevated centered class="mb-10">
        <label for="search" class="sr-only">Search FAQs</label>
        <AppInput
          id="search"
          v-model="searchTerm"
          type="text"
          placeholder="Search questions..."
          tone="primary"
          aria-label="Search FAQs"
        />
      </AppCard>

      <div class="space-y-4">
        <div
          v-for="faq in filteredFaqs"
          :key="faq.id"
          class="overflow-hidden rounded-xl border border-gray-100 bg-white shadow-lg transition-all duration-300"
        >
          <button
            class="flex w-full items-center justify-between rounded-t-xl p-6 text-left transition hover:bg-blue-50 focus:outline-none focus:ring-2 focus:ring-blue-500"
            @click="toggleFaq(faq.id)"
            :id="`faq-question-${faq.id}`"
            :aria-expanded="openFaq === faq.id"
            :aria-controls="`faq-answer-${faq.id}`"
            :aria-label="`Toggle answer for: ${faq.question}`"
          >
            <span class="text-xl font-semibold text-blue-800">{{ faq.question }}</span>
            <svg
              class="h-6 w-6 transform text-blue-500 transition-transform duration-300"
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
            class="border-t border-gray-100 p-6 pt-0 text-lg text-gray-700"
            role="region"
            :aria-labelledby="`faq-question-${faq.id}`"
          >
            <p class="whitespace-pre-line">{{ faq.answer }}</p>
          </div>
        </div>
      </div>

      <AppCard size="lg" padding="lg" elevated centered class="mt-12 ui-card--accent ui-tone-primary">
        <template #header>
          <div class="text-center">
            <h2 class="text-3xl font-extrabold text-blue-800">Can't find your answer?</h2>
            <p class="mt-3 text-lg text-gray-700">Ask us directly and we'll get back to you soon.</p>
          </div>
        </template>

        <form @submit.prevent="handleSubmitQuestion" class="mx-auto max-w-2xl space-y-4">
          <label for="question" class="sr-only">Your Question</label>
          <AppTextarea
            id="question"
            v-model="question"
            :rows="4"
            placeholder="Type your question here..."
            tone="primary"
            required
            aria-label="Your question"
          />

          <AppButton
            type="submit"
            tone="primary"
            :loading="isSubmitting"
            loading-label="Submitting..."
            :disabled="isSubmitting"
            block
          >
            Submit Question
          </AppButton>
        </form>

        <div v-if="successMessage || errorMessage" class="mt-4 space-y-3">
          <AlertMessage
            v-if="successMessage"
            type="success"
            :message="successMessage"
            compact
          />
          <AlertMessage
            v-if="errorMessage"
            type="error"
            :message="errorMessage"
            compact
          />
        </div>
      </AppCard>

      <div class="feedback-btn mt-10 text-center">
        <a
          href="mailto:patelbr5118s@gmail.com?subject=Website Feedback"
          class="inline-block rounded-lg bg-green-500 px-8 py-3 text-lg font-semibold text-white shadow transition duration-200 hover:bg-green-600"
        >
          Send us Feedback
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useFAQSEO } from '../composables/useSEO';
import AlertMessage from '../components/AlertMessage.vue';
import AppButton from '../components/ui/AppButton.vue';
import AppCard from '../components/ui/AppCard.vue';
import AppInput from '../components/ui/AppInput.vue';
import AppTextarea from '../components/ui/AppTextarea.vue';
import { apiService } from '../services/api.service';
import { sanitizeInput } from '../utils/helpers';

useFAQSEO();

interface FaqItem {
  id: string | number;
  question: string;
  answer: string;
}

const staticFaqs: FaqItem[] = [
  { id: 'static-1', question: 'What is this platform about?', answer: 'Our platform provides high-quality online courses across various domains to help you learn and grow your skills.' },
  { id: 'static-2', question: 'How do I buy a course?', answer: 'Choose a course, add it to your cart, and complete the checkout process using UPI, debit/credit cards, or wallets.' },
  { id: 'static-3', question: 'Do you offer certificates?', answer: 'Yes, we provide certificates for most paid courses after successful completion.' },
  { id: 'static-4', question: 'Can I access the course after purchase forever?', answer: 'Yes, once you purchase a course, you get lifetime access unless specified otherwise.' },
  { id: 'static-5', question: 'Is there a mobile app?', answer: 'We are working on our mobile app. Meanwhile, you can access all features through your mobile browser.' },
  { id: 'static-6', question: 'How can I get support?', answer: 'You can email us at patelbr5118s@gmail.com or use the contact form on our website.' },
];

const faqs = ref<FaqItem[]>([]);
const searchTerm = ref('');
const question = ref('');
const openFaq = ref<string | number | null>(null);
const isSubmitting = ref(false);
const successMessage = ref('');
const errorMessage = ref('');

onMounted(async () => {
  try {
    const data = await apiService.get<FaqItem[]>('/api/faqs/');
    faqs.value = data;
  } catch (error) {
    console.error('Failed to load FAQs:', error);
    faqs.value = [];
  }
});

const allFaqs = computed(() => [...staticFaqs, ...faqs.value]);

const filteredFaqs = computed(() => {
  if (!searchTerm.value) return allFaqs.value;
  const term = searchTerm.value.toLowerCase();

  return allFaqs.value.filter((faq) =>
    faq.question.toLowerCase().includes(term) || faq.answer.toLowerCase().includes(term),
  );
});

const toggleFaq = (id: string | number) => {
  openFaq.value = openFaq.value === id ? null : id;
};

const handleSubmitQuestion = async () => {
  if (!question.value.trim()) return;

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
