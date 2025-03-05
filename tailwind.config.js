/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'dark-blue': '#0a192f',
        'light-blue': '#172a45',
        'red': '#e63946',
        'black': '#121212',
        'gray': '#8892b0',
        'light-gray': '#ccd6f6',
      },
    },
  },
  plugins: [],
}
