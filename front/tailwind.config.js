/** @type {import('tailwindcss').Config} */
// eslint-disable-next-line @typescript-eslint/no-var-requires
const { theme } = require('./src/tailwind/tailwind.theme');

module.exports = {
    important: '#root',
    prefix: 'tw-',
    content: ['./src/**/*.{js,ts,jsx,tsx}'],
    theme,
    plugins: [],
};
