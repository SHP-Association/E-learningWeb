<template>
  <div class="quiz-card" @click="handleClick">
    <div class="quiz-card-header">
      <h3 class="quiz-title">{{ quiz.title }}</h3>
      <span class="quiz-badge">Quiz</span>
    </div>
    
    <p v-if="quiz.description" class="quiz-description">
      {{ quiz.description }}
    </p>
    
    <div class="quiz-meta">
      <div class="meta-item">
        <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <span>Passing Score: {{ quiz.passing_score }}%</span>
      </div>
      
      <div v-if="courseName" class="meta-item">
        <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
        </svg>
        <span>{{ courseName }}</span>
      </div>
    </div>
    
    <div class="quiz-footer">
      <button 
        class="quiz-btn" 
        :class="{ 'enrolled': isEnrolled, 'not-enrolled': !isEnrolled }"
      >
        <span v-if="isEnrolled">Take Quiz</span>
        <span v-else>Enroll to Take Quiz</span>
      </button>
      
      <div v-if="attemptData" class="attempt-status">
        <span class="status-badge" :class="{ 'passed': attemptData.passed }">
          {{ attemptData.passed ? 'Passed' : 'Failed' }} - {{ attemptData.score }}%
        </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';
import type { Quiz, QuizAttempt } from '../types/api.types';

interface Props {
  quiz: Quiz;
  isEnrolled?: boolean;
  courseName?: string;
  attemptData?: QuizAttempt | null;
}

const props = withDefaults(defineProps<Props>(), {
  isEnrolled: false,
  courseName: '',
  attemptData: null
});

const emit = defineEmits<{
  click: [quiz: Quiz, isEnrolled: boolean];
}>();

const handleClick = () => {
  emit('click', props.quiz, props.isEnrolled);
};
</script>

<style scoped>
.quiz-card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition: all 0.3s ease;
  border: 2px solid transparent;
}

.quiz-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
  border-color: #3b82f6;
}

.quiz-card-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 1rem;
}

.quiz-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #1f2937;
  margin: 0;
  flex: 1;
}

.quiz-badge {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
}

.quiz-description {
  color: #6b7280;
  font-size: 0.875rem;
  line-height: 1.5;
  margin-bottom: 1rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.quiz-meta {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #6b7280;
  font-size: 0.875rem;
}

.icon {
  width: 1.25rem;
  height: 1.25rem;
  color: #9ca3af;
}

.quiz-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
  margin-top: 1.5rem;
  padding-top: 1rem;
  border-top: 1px solid #e5e7eb;
}

.quiz-btn {
  flex: 1;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  font-weight: 600;
  font-size: 0.875rem;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
}

.quiz-btn.enrolled {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.quiz-btn.enrolled:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.quiz-btn.not-enrolled {
  background: #f3f4f6;
  color: #6b7280;
  border: 2px solid #d1d5db;
}

.quiz-btn.not-enrolled:hover {
  background: #e5e7eb;
  border-color: #9ca3af;
}

.attempt-status {
  display: flex;
  align-items: center;
}

.status-badge {
  padding: 0.5rem 1rem;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 600;
  background: #fef3c7;
  color: #92400e;
}

.status-badge.passed {
  background: #d1fae5;
  color: #065f46;
}

@media (max-width: 768px) {
  .quiz-card {
    padding: 1rem;
  }
  
  .quiz-title {
    font-size: 1.125rem;
  }
  
  .quiz-footer {
    flex-direction: column;
    align-items: stretch;
  }
  
  .quiz-btn {
    width: 100%;
  }
}
</style>
