/** @type {import('tailwindcss').Config} */
export default {
	content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
	theme: {
		extend: {
			colors: {
				primary: "#212936",
				textColor: "#c9d1d9",
				accentOne: "#0366d6",
				accentTwo: "#28a745",
				hoverColor: "#0056b3",
				errorColor: "#d73a49",
				sideBg: "#161b22",
				sideText: "#8b949e",
				cardBg: "#161b22",
				btnColor: "#ffffff",
			},
		},
	},
	plugins: [],
};
