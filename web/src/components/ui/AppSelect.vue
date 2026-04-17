<template>
  <div :class="['ui-field', toneClassMap[tone]]">
    <label v-if="label" :for="resolvedId" class="ui-field__label">{{ label }}</label>
    <div class="ui-select">
      <select
        v-bind="attrs"
        :id="resolvedId"
        :name="name"
        :value="selectedValue"
        :required="required"
        :disabled="disabled"
        :aria-invalid="Boolean(error) || undefined"
        :aria-describedby="describedBy || undefined"
        :class="controlClass"
        @change="handleChange"
      >
        <option v-if="placeholder" value="" disabled>{{ placeholder }}</option>
        <option
          v-for="option in normalizedOptions"
          :key="String(option.value)"
          :value="String(option.value)"
          :disabled="option.disabled"
        >
          {{ option.label }}
        </option>
      </select>
    </div>
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

type SelectOption =
  | string
  | number
  | {
      label: string;
      value: string | number;
      disabled?: boolean;
    };

const attrs = useAttrs();
const instance = getCurrentInstance();

const props = withDefaults(defineProps<{
  modelValue?: string | number;
  id?: string;
  label?: string;
  name?: string;
  options: SelectOption[];
  placeholder?: string;
  required?: boolean;
  disabled?: boolean;
  hint?: string;
  error?: string;
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
  tone: 'primary',
});

const emit = defineEmits<{
  (event: 'update:modelValue', value: string | number): void;
}>();

const resolvedId = computed(() => props.id || `app-select-${instance?.uid ?? 'field'}`);
const hintId = computed(() => `${resolvedId.value}-hint`);
const errorId = computed(() => `${resolvedId.value}-error`);

const normalizedOptions = computed(() =>
  props.options.map((option) => {
    if (typeof option === 'string' || typeof option === 'number') {
      return {
        label: String(option),
        value: option,
        disabled: false,
      };
    }

    return {
      label: option.label,
      value: option.value,
      disabled: option.disabled ?? false,
    };
  }),
);

const selectedValue = computed(() => String(props.modelValue ?? ''));

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
  'ui-field__select',
  {
    'ui-field__control--error': props.error,
  },
]);

const handleChange = (event: Event) => {
  const nextValue = (event.target as HTMLSelectElement).value;
  const matchedOption = normalizedOptions.value.find((option) => String(option.value) === nextValue);

  emit('update:modelValue', matchedOption?.value ?? nextValue);
};
</script>
