<template>
  <Teleport to="body">
    <Transition name="fade">
      <div v-if="modelValue" class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6">
        <!-- Backdrop -->
        <div 
          class="fixed inset-0 bg-slate-950/80 backdrop-blur-md" 
          @click="$emit('update:modelValue', false)"
        ></div>

        <!-- Dialog Panel -->
        <div 
          :class="wide ? 'max-w-2xl' : 'max-w-md'"
          class="dialog relative z-10 w-full overflow-hidden border border-white/10 bg-zinc-900 shadow-2xl ring-1 ring-white/5"
          style="border-radius: 1.5rem;"
        >
          <!-- Header -->
          <div class="flex items-center justify-between border-b border-white/5 p-6">
            <h3 class="text-xl font-bold tracking-tight text-white">{{ title }}</h3>
            <button 
              class="rounded-full p-2 text-zinc-400 transition-all hover:bg-white/5 hover:text-white"
              @click="$emit('update:modelValue', false)"
            >
              <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>

          <!-- Content -->
          <div class="p-6 text-zinc-300">
            <slot />
          </div>

          <!-- Optional Footer Slot -->
          <div v-if="$slots.footer" class="border-t border-white/5 p-6">
            <slot name="footer" />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { watch, onDeactivated } from 'vue';
import { onBeforeRouteLeave } from 'vue-router';

interface Props {
  modelValue: boolean;
  title: string;
  wide?: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits(['update:modelValue']);

// Navigation protection: Ensure dialog closes on route change
onBeforeRouteLeave(() => {
  emit('update:modelValue', false);
});

// Deactivation protection
onDeactivated(() => {
  emit('update:modelValue', false);
});

// Body scroll locking logic
watch(
  () => props.modelValue,
  (isOpen) => {
    const html = document.querySelector('html');
    if (!html) return;
    
    if (isOpen) {
      html.classList.add('overflow-hidden');
    } else {
      html.classList.remove('overflow-hidden');
    }
  },
  { immediate: true }
);
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.fade-enter-active .dialog {
  animation: dialog-pop 0.4s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.fade-leave-active .dialog {
  animation: dialog-pop 0.2s cubic-bezier(0.34, 1.56, 0.64, 1) reverse;
}

@keyframes dialog-pop {
  from {
    opacity: 0;
    transform: scale(0.9) translateY(20px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}
</style>
