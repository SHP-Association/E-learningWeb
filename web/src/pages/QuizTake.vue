<template>
  <div class="quiz-take-container">
    <div v-if="loading" class="loading-state">
      <div class="spinner"></div>
      <p>Loading quiz...</p>
    </div>

    <div v-else-if="error" class="error-state">
      <h2>{{ error }}</h2>
      <router-link to="/" class="btn-primary">Go Home</router-link>
    </div>

    <div v-else-if="!isEnrolled" class="enrollment-required">
      <div class="enrollment-card">
        <svg class="lock-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
        </svg>
        <h2>Enrollment Required</h2>
        <p>You need to enroll in the course to take this quiz.</p>
        <button @click="goToCourse" class="btn-primary">View Course & Enroll</button>
      </div>
    </div>

    <div v-else-if="quiz && questions.length > 0" class="quiz-content">
      <div class="quiz-header">
        <h1>{{ quiz.title }}</h1>
        <p v-if="quiz.description" class="quiz-desc">{{ quiz.description }}</p>
        <div class="quiz-info">
          <span class="info-badge">{{ questions.length }} Questions</span>
          <span class="info-badge">Passing Score: {{ quiz.passing_score }}%</span>
        </div>
      </div>

      <div v-if="!submitted" class="questions-container">
        <div 
          v-for="(question, index) in questions" 
          :key="question.id"
          class="question-card"
        >
          <div class="question-header">
            <span class="question-number">Question {{ index + 1 }}</span>
          </div>
          
          <h3 class="question-text">{{ question.question_text }}</h3>
          
          <div v-if="question.question_type === 'mcq'" class="choices">
            <label 
              v-for="choice in question.choices" 
              :key="choice.id"
              class="choice-option"
              :class="{ 'selected': userAnswers[question.id] === choice.id }"
            >
              <input 
                type="radio" 
                :name="`question-${question.id}`"
                :value="choice.id"
                v-model="userAnswers[question.id]"
              />
              <span class="choice-text">{{ choice.choice_text }}</span>
            </label>
          </div>
          
          <div v-else class="text-answer">
            <textarea 
              v-model="userAnswers[question.id]"
              placeholder="Type your answer here..."
              rows="4"
            ></textarea>
          </div>
        </div>

        <div class="quiz-actions">
          <button @click="submitQuiz" class="btn-submit" :disabled="submitting">
            <span v-if="submitting">Submitting...</span>
            <span v-else>Submit Quiz</span>
          </button>
        </div>
      </div>

      <div v-else class="results-container">
        <div class="results-card">
          <div class="results-header" :class="{ 'passed': result.passed, 'failed': !result.passed }">
            <svg v-if="result.passed" class="result-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <svg v-else class="result-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <h2>{{ result.passed ? 'Congratulations!' : 'Keep Trying!' }}</h2>
            <p class="score">Your Score: {{ result.score }}%</p>
            <p class="status">{{ result.passed ? 'You passed the quiz!' : 'You need ' + quiz.passing_score + '% to pass.' }}</p>
          </div>
          
          <div class="results-actions">
            <button @click="retakeQuiz" class="btn-secondary">Retake Quiz</button>
            <router-link :to="`/course/${courseSlug}`" class="btn-primary">Back to Course</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { apiService } from '../services/api.service';
import type { Quiz, Question } from '../types/api.types';

const route = useRoute();
const router = useRouter();

const quizId = computed(() => parseInt(route.params.id as string));
const loading = ref(true);
const error = ref('');
const quiz = ref<Quiz | null>(null);
const questions = ref<Question[]>([]);
const userAnswers = ref<Record<number, any>>({});
const submitted = ref(false);
const submitting = ref(false);
const result = ref<any>(null);
const isEnrolled = ref(false);
const courseSlug = ref('');

onMounted(async () => {
  try {
    // Fetch quiz details
    quiz.value = await apiService.getQuiz(quizId.value);
    
    // Fetch questions
    const questionsData = await apiService.getQuestions(quizId.value);
    questions.value = questionsData.items;
    
    // Check enrollment (simplified - you may need to get course ID from lesson)
    // For now, we'll assume user is enrolled if they can access the quiz
    isEnrolled.value = true; // TODO: Implement proper enrollment check
    
    loading.value = false;
  } catch (err: any) {
    error.value = err.message || 'Failed to load quiz';
    loading.value = false;
  }
});

const submitQuiz = async () => {
  // Validate all questions are answered
  const unanswered = questions.value.filter(q => !userAnswers.value[q.id]);
  if (unanswered.length > 0) {
    alert(`Please answer all questions. ${unanswered.length} question(s) remaining.`);
    return;
  }

  submitting.value = true;
  try {
    // Format answers for submission
    const answers = questions.value.map(q => ({
      question_id: q.id,
      selected_choice_id: q.question_type === 'mcq' ? userAnswers.value[q.id] : undefined,
      text_answer: q.question_type === 'text' ? userAnswers.value[q.id] : undefined
    }));

    // Submit quiz
    const response = await apiService.submitQuiz(quizId.value, { answers });
    
    // Calculate score (simplified)
    let correct = 0;
    questions.value.forEach(q => {
      if (q.question_type === 'mcq') {
        const selectedChoice = q.choices?.find(c => c.id === userAnswers.value[q.id]);
        if (selectedChoice?.is_correct) correct++;
      }
    });
    
    const score = Math.round((correct / questions.value.length) * 100);
    const passed = score >= (quiz.value?.passing_score || 70);
    
    result.value = { score, passed };
    submitted.value = true;
  } catch (err: any) {
    alert(err.message || 'Failed to submit quiz');
  } finally {
    submitting.value = false;
  }
};

const retakeQuiz = () => {
  userAnswers.value = {};
  submitted.value = false;
  result.value = null;
  window.scrollTo({ top: 0, behavior: 'smooth' });
};

const goToCourse = () => {
  if (courseSlug.value) {
    router.push(`/course/${courseSlug.value}`);
  } else {
    router.push('/');
  }
};
</script>

<style scoped>
.quiz-take-container {
  max-width: 900px;
  margin: 2rem auto;
  padding: 0 1rem;
}

.loading-state, .error-state {
  text-align: center;
  padding: 4rem 2rem;
}

.spinner {
  width: 50px;
  height: 50px;
  border: 4px solid #f3f4f6;
  border-top-color: #3b82f6;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 1rem;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.enrollment-required {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
}

.enrollment-card {
  background: white;
  padding: 3rem;
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  text-align: center;
  max-width: 500px;
}

.lock-icon {
  width: 64px;
  height: 64px;
  color: #ef4444;
  margin: 0 auto 1.5rem;
}

.quiz-header {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  margin-bottom: 2rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.quiz-header h1 {
  margin: 0 0 0.5rem;
  color: #1f2937;
}

.quiz-desc {
  color: #6b7280;
  margin-bottom: 1rem;
}

.quiz-info {
  display: flex;
  gap: 1rem;
  flex-wrap: wrap;
}

.info-badge {
  background: #eff6ff;
  color: #1e40af;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 600;
}

.questions-container {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.question-card {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.question-header {
  margin-bottom: 1rem;
}

.question-number {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 600;
}

.question-text {
  color: #1f2937;
  font-size: 1.125rem;
  margin-bottom: 1.5rem;
}

.choices {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.choice-option {
  display: flex;
  align-items: center;
  padding: 1rem;
  border: 2px solid #e5e7eb;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.choice-option:hover {
  border-color: #3b82f6;
  background: #eff6ff;
}

.choice-option.selected {
  border-color: #3b82f6;
  background: #eff6ff;
}

.choice-option input[type="radio"] {
  margin-right: 1rem;
  width: 20px;
  height: 20px;
}

.choice-text {
  flex: 1;
  color: #374151;
}

.text-answer textarea {
  width: 100%;
  padding: 1rem;
  border: 2px solid #e5e7eb;
  border-radius: 8px;
  font-family: inherit;
  font-size: 1rem;
  resize: vertical;
}

.quiz-actions {
  margin-top: 2rem;
  text-align: center;
}

.btn-submit {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 1rem 3rem;
  border-radius: 8px;
  font-weight: 600;
  font-size: 1rem;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-submit:hover:not(:disabled) {
  transform: scale(1.05);
  box-shadow: 0 8px 16px rgba(102, 126, 234, 0.4);
}

.btn-submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.results-card {
  background: white;
  padding: 3rem;
  border-radius: 16px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.results-header {
  padding: 2rem;
  border-radius: 12px;
  margin-bottom: 2rem;
}

.results-header.passed {
  background: linear-gradient(135deg, #d1fae5 0%, #a7f3d0 100%);
}

.results-header.failed {
  background: linear-gradient(135deg, #fee2e2 0%, #fecaca 100%);
}

.result-icon {
  width: 64px;
  height: 64px;
  margin: 0 auto 1rem;
}

.results-header.passed .result-icon {
  color: #065f46;
}

.results-header.failed .result-icon {
  color: #991b1b;
}

.score {
  font-size: 2rem;
  font-weight: 700;
  margin: 1rem 0;
}

.results-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

.btn-primary, .btn-secondary {
  padding: 0.75rem 2rem;
  border-radius: 8px;
  font-weight: 600;
  text-decoration: none;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-secondary {
  background: white;
  color: #667eea;
  border: 2px solid #667eea;
}

.btn-primary:hover, .btn-secondary:hover {
  transform: scale(1.05);
}

@media (max-width: 768px) {
  .quiz-take-container {
    padding: 0 0.5rem;
  }
  
  .quiz-header, .question-card {
    padding: 1.5rem;
  }
  
  .enrollment-card {
    padding: 2rem;
  }
}
</style>
