/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    screens: {
      xs: "400px",
      // 25rem
      sm: "640px",
      // => @media (min-width: 640px) { ... }

      md: "768px",
      // => @media (min-width: 768px) { ... }

      hd: "1000px",

      lg: "1024px",
      // => @media (min-width: 1024px) { ... }

      wide: "1160px",

      xl: "1280px",
      // => @media (min-width: 1280px) { ... }

      "2xl": "1536px",
      // => @media (min-width: 1536px) { ... }
    },
    extend: {
      spacing: { 320: "320px", 280: "280" },
      colors: {
        "bp-header-color": "#394b59",
        "bp-card-color": "#30404d",
        "bp-bg": "#293742",
      },
    },
  },
  plugins: [],
};
