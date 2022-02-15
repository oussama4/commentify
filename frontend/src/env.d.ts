interface ImportMetaEnv {
  readonly VITE_COMMENTIFY_BASE_URL: string
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}