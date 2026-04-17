export type AppTone = 'primary' | 'secondary' | 'success' | 'danger';
export type AppButtonVariant = 'solid' | 'outline' | 'ghost';
export type AppCardSize = 'sm' | 'md' | 'lg';
export type AppCardPadding = 'sm' | 'md' | 'lg';

export const toneClassMap: Record<AppTone, string> = {
  primary: 'ui-tone-primary',
  secondary: 'ui-tone-secondary',
  success: 'ui-tone-success',
  danger: 'ui-tone-danger',
};

export const buttonVariantClassMap: Record<AppButtonVariant, string> = {
  solid: 'ui-button--solid',
  outline: 'ui-button--outline',
  ghost: 'ui-button--ghost',
};

export const cardSizeClassMap: Record<AppCardSize, string> = {
  sm: 'ui-card--sm',
  md: 'ui-card--md',
  lg: 'ui-card--lg',
};

export const cardPaddingClassMap: Record<AppCardPadding, string> = {
  sm: 'ui-card--padding-sm',
  md: 'ui-card--padding-md',
  lg: 'ui-card--padding-lg',
};
