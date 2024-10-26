import { defineConfig } from 'unocss'

export default defineConfig({
  // ...UnoCSS options
  content:{
    filesystem:[
      'src/**/*.{vue,js,ts,jsx,tsx}',
      "../../templates/**/*.{vue,js,ts,jsx,tsx,html}",
    ]
  },
  theme:{
    colors:{
        primary: '#4233cf',
        secondary: '#0F609B',
        accent: '#F59E0B',
        background: '#F9F9F9'
        
    }
  },
  shortcuts:{
    "title-1":"font-bold text-2xl",
    "title-2":"font-bold text-xl",
    "title-3":"font-bold text-lg",
    "btn":"cursor-pointer transition duration-300 ease-in-out hover:opacity-80 px-2 bg-gray-200 text-gray-800 rounded-1 hover:bg-gray-300 hover:text-gray-900"
  }
})