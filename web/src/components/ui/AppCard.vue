<template>
  <section :class="cardClass">
    <header v-if="$slots.header" class="ui-card__header">
      <slot name="header" />
    </header>

    <slot />

    <footer v-if="$slots.footer" class="ui-card__footer">
      <slot name="footer" />
    </footer>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import {
  cardPaddingClassMap,
  cardSizeClassMap,
  type AppCardPadding,
  type AppCardSize,
} from './theme';

const props = withDefaults(defineProps<{
  size?: AppCardSize;
  padding?: AppCardPadding;
  elevated?: boolean;
  centered?: boolean;
}>(), {
  size: 'md',
  padding: 'md',
  elevated: true,
  centered: false,
});

const cardClass = computed(() => [
  'ui-card',
  cardSizeClassMap[props.size],
  cardPaddingClassMap[props.padding],
  {
    'ui-card--elevated': props.elevated,
    'ui-card--centered': props.centered,
  },
]);
</script>
