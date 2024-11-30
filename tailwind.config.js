/** @type {import('tailwindcss').Config} */
export default {
	content: ["./views/**/*.{templ,html,js}"],
	theme: {},
	plugins: [
		require('@tailwindcss/forms')
	]
}
