/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./web/**/*.html",
        "./web/**/*.templ",
        "./node_modules/flowbite/**/*.js",
    ],
    theme: {
        extend: {},
    },
    plugins: [
        require('flowbite/plugin'),
    ],
}

