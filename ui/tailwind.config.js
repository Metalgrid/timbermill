module.exports = {
  purge: ["./src/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      colors: require("daisyui/colors"),
    },
  },
  variants: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
