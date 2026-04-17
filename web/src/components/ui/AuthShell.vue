<template>
  <div :class="['ui-auth-shell', toneClassMap[tone]]">
    <AppCard size="md" padding="lg" elevated centered class="ui-auth-shell__card">
      <template #header>
        <div class="ui-auth-shell__header">
          <h1 class="ui-auth-shell__title animate-fade-in-down">{{ title }}</h1>
          <p v-if="subtitle" class="ui-auth-shell__subtitle animate-fade-in anim-delay-200">
            {{ subtitle }}
          </p>
          <div v-if="$slots.status" class="animate-fade-in anim-delay-300">
            <slot name="status" />
          </div>
        </div>
      </template>

      <div class="animate-fade-in anim-delay-300">
        <slot />
      </div>

      <template v-if="$slots.footer" #footer>
        <div class="animate-fade-in anim-delay-500">
          <slot name="footer" />
        </div>
      </template>
    </AppCard>
  </div>
</template>

<script setup lang="ts">
import AppCard from './AppCard.vue';
import { toneClassMap, type AppTone } from './theme';

withDefaults(defineProps<{
  title: string;
  subtitle?: string;
  tone?: AppTone;
}>(), {
  subtitle: '',
  tone: 'primary',
});
</script>
