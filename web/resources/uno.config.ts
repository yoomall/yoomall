import { defineConfig } from 'unocss'

export default defineConfig({
  // ...UnoCSS options
  content:{
    filesystem:[
      'src/**/*.{vue,js,ts,jsx,tsx}',
      "../../templates/**/*.{vue,js,ts,jsx,tsx,html}",
    ]
  }
})