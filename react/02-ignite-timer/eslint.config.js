import js from '@eslint/js'
import globals from 'globals'
import react from 'eslint-plugin-react'
import reactHooks from 'eslint-plugin-react-hooks'
import reactRefresh from 'eslint-plugin-react-refresh'
import tseslint from '@typescript-eslint/eslint-plugin'
import tsParser from '@typescript-eslint/parser'
import prettierPlugin from 'eslint-plugin-prettier'
import rocketseatLint from '@rocketseat/eslint-config/react.js';

export default [
    {
        ignores: ['dist', 'node_modules'],
    },
    {
        // define os arquivos a serem verificados
        files: ['**/*.{ts,tsx}'],
        languageOptions: {
            parser: tsParser,
            ecmaVersion: 2021,
            sourceType: 'module',
            globals: globals.browser,
            parserOptions: {
                ecmaFeatures: {
                    jsx: true, // Suporte a JSX
                },
            },
        },
        plugins: {
            react,
            'react-hooks': reactHooks,
            'react-refresh': reactRefresh,
            '@typescript-eslint': tseslint,
            prettier: prettierPlugin,
            rocketseatLint,
        },
        rules: {
            'prettier/prettier': [
                'error',
                {
                    printWidth: 80,
                    tabWidth: 2,
                    singleQuote: true,
                    trailingComma: 'all',
                    arrowParens: 'always',
                    semi: false,
                    endOfLine: 'auto',
                },
            ],
            'react/no-unknown-property': 'error',
            ...reactHooks.configs.recommended.rules,
            ...tseslint.configs.recommended.rules,
            ...js.configs.recommended.rules,
            'react/self-closing-comp': 'error',
            '@typescript-eslint/no-empty-object-type': 'off',
            'react/react-in-jsx-scope': 'off',
            'react/prop-types': 'off',
            'no-unused-vars': 'off',
            '@typescript-eslint/no-unused-vars': ['warn', { 
                vars: 'all', 
                args: 'after-used', 
                ignoreRestSiblings: true 
            }],
            '@typescript-eslint/no-explicit-any': 'warn',
        },
        settings: {
            react: {
                version: 'detect',
            },
        },
    },
]