/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./static/index.html"],
  theme: {
    extend: {},
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: ["cmyk", "dracula"]
  }
}

