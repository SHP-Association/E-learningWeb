// vite-env.d.ts
/// <reference types="vite/client" />

interface ImportMetaEnv {
    readonly VITE_APP_BACKEND_URL: string;
    readonly VITE_APP_BASE_URL: string;
    // Add other env variables here
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}
