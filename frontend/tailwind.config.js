/** @type {import('tailwindcss').Config} */

import colors from 'tailwindcss/colors';

export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    colors: {
      offwhite: '#E6E2E2',
      cyan: '#39D6E1',
      extend: {},
      black: colors.black,
      'black-rgba': 'rgba(0, 0, 0, 0.8)',
      white: colors.white,
      gray: colors.gray,
    },
  },
  plugins: [],
};
