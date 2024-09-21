/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./web/**/*.{html,js}"],
	theme: {
	},
	plugins: [
		require('@tailwindcss/forms'),
	],
}
