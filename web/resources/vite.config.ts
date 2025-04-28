import { defineConfig, PluginOption } from "vite";
import babel from "@rollup/plugin-babel";
import UnoCSS from "unocss/vite";
import ViteRestart from 'vite-plugin-restart'

export default defineConfig({
  plugins: [UnoCSS(),ViteRestart({
    restart:[
      "../../templates/**/*",
    ]
  })],
  css: {
    preprocessorOptions: {
      scss: {
        silenceDeprecations: ["legacy-js-api"],
      },
    },
  },
  build: {
    manifest: "manifest.json",
    outDir: "../../public",
    emptyOutDir: true,
    sourcemap: true,
    minify: true,
    assetsDir: "assets",
    target: ["es2020", "chrome100"],
    modulePreload: {
      polyfill: true,
    },
    assetsInlineLimit: 0,
    rollupOptions: {
      input: ["src/main.js", "src/scss/main.scss"],
      plugins: [
        babel({
          babelHelpers: "bundled",
          presets: [
            [
              "@babel/preset-env",
              {
                corejs: 3,
                useBuiltIns: "entry",
                targets: {
                  chrome: "100",
                  edge: "100",
                  firefox: "100",
                  safari: "100",
                  ie: "11",
                },
              },
            ],
          ],
          exclude: "node_modules/**",
        }),
      ],
    },
  },
});
