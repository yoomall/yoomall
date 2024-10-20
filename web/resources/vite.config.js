import { defineConfig } from 'vite';
import laravel from 'laravel-vite-plugin';
import babel from '@rollup/plugin-babel'

export default defineConfig({
  plugins: [
    laravel({
      input: ['src/main.js'],
      refresh: true,
      publicDirectory: '../../public',
      buildDirectory: '.',
    }),
  ],
  esbuild: {
    target: 'es2020',
  },
  build: {
    target: [
      'es2020',
      'chrome100'
    ],
    modulePreload: {
      polyfill: true
    },
    rollupOptions: {
      plugins: [
        babel({
          babelHelpers: 'bundled',
          presets: [['@babel/preset-env', {
            "corejs": 3,
            "useBuiltIns": "entry",
            targets: {
              chrome: "100",
              edge: "100",
              firefox: "100",
              safari: "100",
              ie: "11"
            }
          }]],
          exclude: 'node_modules/**',
        })
      ]
    }
  },

});
