import { defineConfig, globalIgnores } from 'eslint/config'
import nextVitals from 'eslint-config-next/core-web-vitals'
import nextTs from 'eslint-config-next/typescript'
import prettierPlugin from 'eslint-plugin-prettier'
import prettierConfig from 'eslint-config-prettier'

export default defineConfig([
  ...nextVitals,
  ...nextTs,
  globalIgnores(['.next/**', 'node_modules/**', 'build/**']),

  {
    plugins: {
      prettier: prettierPlugin,
    },
    rules: {
      // Força as regras de estilo da Rocketseat via Prettier
      'prettier/prettier': [
        'error',
        {
          singleQuote: true,
          semi: false,
          trailingComma: 'all',
          arrowParens: 'avoid',
          endOfLine: 'auto',
          printWidth: 80,
        },
      ],
      // Regras de limpeza de código
      'no-unused-vars': 'off',
      '@typescript-eslint/no-unused-vars': ['warn'],
      'react/self-closing-comp': 'error',
      'react/jsx-no-useless-fragment': 'warn',
    },
  },
  // Desativa conflitos (deve ser o último)
  prettierConfig,
])
