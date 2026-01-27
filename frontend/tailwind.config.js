
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          500: '#667eea',
          600: '#5568d3',
          700: '#4c51bf',
        },
        secondary: {
          500: '#764ba2',
        }
      }
    },
  },
  plugins: [],
}
