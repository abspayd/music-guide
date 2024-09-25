/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./views/**/*.{templ,html,js}"],
	theme: {
	},
	plugins: [
		require('@tailwindcss/forms'),
	],
}
