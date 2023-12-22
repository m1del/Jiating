/** @type {import('tailwindcss').Config} */

import colors from 'tailwindcss/colors';

export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
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
      extend: {},
      black: colors.black,
      'black-rgba': 'rgba(0, 0, 0, 0.8)',
      white: colors.white,
      gray: colors.gray,
      red: colors.red,
      transparent: colors.transparent,
    },
  },
  plugins: [
    function ({ addVariant }) {
      addVariant('child', '& > *');
      addVariant('child-hover', '& > *:hover');
    },
  ],
};
