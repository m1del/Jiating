/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        offwhite: '#E6E2E2',
        cyan: {
          DEFAULT: '#39D6E1', // Base color
          '50': '#E6F8FB',
          '100': '#CDF3F7',
          '200': '#9DE9F0',
          '300': '#6DDFE9',
          '400': '#3DD6E1',   // slightly lighter
          '500': '#39D6E1',   // Base color
          '600': '#35C2C9',   // slightly darker
          '700': '#2EA5B2',
          '800': '#27879A',
          '900': '#206A83',
        },

        'black-rgba': 'rgba(0, 0, 0, 0.8)',
      },
    },
  },
  plugins: [
    function ({ addVariant }) {
      addVariant('child', '& > *');
      addVariant('child-hover', '& > *:hover');
    },
  ],
};
