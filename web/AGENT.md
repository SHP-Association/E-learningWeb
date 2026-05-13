# Frontend Agent Guide

## Scope
This `web/` directory is a Vue 3 + Vite SPA using Vue Router, Pinia, and a typed API service layer.

## Quick Start
```bash
cd web
npm install
npm run dev
```

Other commands:
```bash
npm run build
npm run preview
```

## Entry Points
- `src/main.ts`: app bootstrap (Vue app mount + plugin wiring).
- `src/App.vue`: top-level shell.
- `src/router.ts`: route table and navigation guard hook.

## Architecture Map
- `src/pages/`: route-level views (`Home`, `CourseDetail`, `Enroll`, auth/password reset, profile, quiz, FAQ).
- `src/components/`: shared UI blocks and layout.
- `src/components/ui/`: reusable primitive UI components.
- `src/stores/`: Pinia stores for user/course/enrollment state.
- `src/services/api.service.ts`: centralized HTTP client and endpoint methods.
- `src/types/`: API/SEO shared TS types.
- `src/composables/`: reusable logic hooks (ex: SEO).
- `src/utils/`: helpers and response transformers.

## Routing
Defined in `src/router.ts` with lazy-loaded page components.
Important paths include:
- `/`
- `/course/:slug`
- `/enroll/:slug`
- `/quiz/:id`
- auth/password reset flows

Unknown routes redirect to `/`.

## API Integration Rules
- Base URL: `VITE_APP_BACKEND_URL` (fallback `http://localhost:8001`).
- Requests include `credentials: include`.
- Non-GET requests try to send CSRF header from `csrftoken` cookie.
- Keep all new API calls in `src/services/api.service.ts` instead of inline `fetch` in components.

## Styling
- Tailwind is configured via `tailwind.config.js` and PostCSS.
- Global styles live in `src/style.css`.
- Reuse existing UI primitives in `src/components/ui/` before creating new one-off components.

## Safety Rules for Future Agents
- Keep route/component naming aligned with existing conventions.
- Update `src/types/api.types.ts` when backend response contracts change.
- Prefer store-driven state updates over duplicated local fetching logic.
- Build-check with `npm run build` after frontend changes.
