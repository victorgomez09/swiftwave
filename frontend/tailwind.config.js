/** @type {import('tailwindcss').Config} */
export default {
  darkMode: false,
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  safelist: ['hidden'],
  daisyui: {
    themes: ['winter', 'night']
  },
  plugins: [require('@tailwindcss/typography'), require("daisyui")]
}
