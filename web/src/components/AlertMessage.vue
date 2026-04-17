<template>
  <div :class="alertClass" role="alert">
    <p v-if="title" class="ui-alert__title">{{ title }}</p>

    <div class="ui-alert__content">
      <slot v-if="$slots.default" />
      <template v-else>
        <p v-if="message">{{ message }}</p>
        <p v-if="details" class="mt-2 text-sm">{{ details }}</p>
        <ul v-if="messages.length" class="list-disc pl-5 space-y-1">
          <li v-for="entry in messages" :key="entry">{{ entry }}</li>
        </ul>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

type AlertType = 'error' | 'success' | 'warning' | 'info';

const props = withDefaults(defineProps<{
  type?: AlertType;
  title?: string;
  message?: string;
  details?: string;
  messages?: string[];
  compact?: boolean;
}>(), {
  type: 'error',
  title: '',
  message: '',
  details: '',
  messages: () => [],
  compact: false,
});

const toneClassMap: Record<AlertType, string> = {
  error: 'ui-tone-danger',
  success: 'ui-tone-success',
  warning: 'ui-tone-warning',
  info: 'ui-tone-primary',
};

const alertClass = computed(() => [
  'ui-alert',
  toneClassMap[props.type],
  props.compact ? 'ui-alert--compact' : 'ui-alert--default',
]);
</script>
