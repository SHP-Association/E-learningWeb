<template>
  <div :class="['ui-field', toneClassMap[tone]]">
    <label v-if="label" :for="resolvedId" class="ui-field__label">{{ label }}</label>
    <textarea
      v-bind="attrs"
      :id="resolvedId"
      :name="name"
      :value="modelValue"
      :placeholder="placeholder"
      :required="required"
      :disabled="disabled"
      :rows="rows"
      :aria-invalid="Boolean(error) || undefined"
      :aria-describedby="describedBy || undefined"
      :class="controlClass"
      @input="handleInput"
    ></textarea>
    <p v-if="hint" :id="hintId" class="ui-field__hint">{{ hint }}</p>
    <p v-if="error" :id="errorId" class="ui-field__error">{{ error }}</p>
  </div>
</template>

<script setup lang="ts">
import { computed, getCurrentInstance, useAttrs } from 'vue';
import { toneClassMap, type AppTone } from './theme';

defineOptions({
  inheritAttrs: false,
});

const attrs = useAttrs();
const instance = getCurrentInstance();

const props = withDefaults(defineProps<{
  modelValue?: string | number;
  id?: string;
  label?: string;
  name?: string;
  placeholder?: string;
  required?: boolean;
  disabled?: boolean;
  hint?: string;
  error?: string;
  rows?: number;
  tone?: AppTone;
}>(), {
  modelValue: '',
  id: undefined,
  label: '',
  name: undefined,
  placeholder: '',
  required: false,
  disabled: false,
  hint: '',
  error: '',
  rows: 4,
  tone: 'primary',
});

const emit = defineEmits<{
  (event: 'update:modelValue', value: string): void;
}>();

const resolvedId = computed(() => props.id || `app-textarea-${instance?.uid ?? 'field'}`);
const hintId = computed(() => `${resolvedId.value}-hint`);
const errorId = computed(() => `${resolvedId.value}-error`);

const describedBy = computed(() => {
  const ids = [];

  if (props.hint) {
    ids.push(hintId.value);
  }

  if (props.error) {
    ids.push(errorId.value);
  }

  return ids.join(' ');
});

const controlClass = computed(() => [
  'ui-field__control',
  'ui-field__textarea',
  {
    'ui-field__control--error': props.error,
  },
]);

const handleInput = (event: Event) => {
  emit('update:modelValue', (event.target as HTMLTextAreaElement).value);
};
</script>
