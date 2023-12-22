/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["internal/template/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms")],
};
