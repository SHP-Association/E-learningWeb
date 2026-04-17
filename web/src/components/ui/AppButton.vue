<template>
  <button
    v-bind="attrs"
    :type="type"
    :disabled="isDisabled"
    :aria-busy="loading || undefined"
    :class="buttonClass"
  >
    <span :class="['ui-button__content', { 'ui-button__content--hidden': loading }]">
      <slot />
    </span>

    <span v-if="loading" class="ui-button__overlay">
      <span class="ui-spinner" aria-hidden="true"></span>
      <span v-if="loadingLabel">{{ loadingLabel }}</span>
      <span v-else class="sr-only">Loading</span>
    </span>
  </button>
</template>

<script setup lang="ts">
import { computed, useAttrs } from 'vue';
import {
  buttonVariantClassMap,
  toneClassMap,
  type AppButtonVariant,
  type AppTone,
} from './theme';

defineOptions({
  inheritAttrs: false,
});

const attrs = useAttrs();

const props = withDefaults(defineProps<{
  variant?: AppButtonVariant;
  tone?: AppTone;
  type?: 'button' | 'submit' | 'reset';
  disabled?: boolean;
  loading?: boolean;
  block?: boolean;
  loadingLabel?: string;
}>(), {
  variant: 'solid',
  tone: 'primary',
  type: 'button',
  disabled: false,
  loading: false,
  block: false,
  loadingLabel: '',
});

const isDisabled = computed(() => props.disabled || props.loading);

const buttonClass = computed(() => [
  'ui-button',
  toneClassMap[props.tone],
  buttonVariantClassMap[props.variant],
  {
    'ui-button--block': props.block,
  },
]);
</script>
