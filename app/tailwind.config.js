/** @type {import('tailwindcss').Config} */
export default {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  theme: {
    extend: {
      colors: {
        white: '#ffffff',
        backg: '#F2F2F2',
        backSoft: '#F6F8FA',
        primary: '#3C7363',
        primaryVariant: '#0F926C',
        darkGray: '#737285',
        softGray: '#CECECE',
        // softprimary: '#d9e8e3',
        softprimary: '#ebf5f1',
        complementary: '#f48c06',
        complementarySoft: '#ffb300',
        orange: '#fca311',
        color1: '#355645',
        black25: 'rgba(0, 0, 0, 0.25)',
        piel: '#D7B49E'
      },
      fontFamily: {
        roboto: ['Roboto', 'sans-serif'],
        opensans: ['Open Sans', 'sans-serif']
      }
    }
  },
  plugins: []
}
