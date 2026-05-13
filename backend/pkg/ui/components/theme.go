package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ThemeToggle() Node {
	return Label(
		Class("swap swap-rotate btn btn-ghost btn-circle btn-md hover:bg-card-bg transition-all duration-300 ease-in-out"),
		Input(
			Type("checkbox"),
			Class("theme-controller"),
			Attr("onchange", "(function(e) { const theme = e.target.checked ? 'light' : 'dark'; document.documentElement.setAttribute('data-theme', theme); localStorage.setItem('theme', theme); })(event)"),
			ID("theme-toggle-checkbox"),
		),
		// Sun icon (for dark mode)
		Span(Class("swap-on fill-none stroke-current transition-transform duration-500 hover:rotate-45"), Icon("Sun", "w-5 h-5")),
		// Moon icon (for light mode)
		Span(Class("swap-off fill-none stroke-current transition-transform duration-500 -rotate-12 hover:rotate-0"), Icon("Moon", "w-5 h-5")),
		Script(Raw(`
			(function() {
				const theme = localStorage.getItem('theme') || 'dark';
				const checkbox = document.getElementById('theme-toggle-checkbox');
				if (checkbox) {
					checkbox.checked = theme === 'light';
				}
			})();
		`)),
	)
}

func ThemeStyles() Node {
	return Raw(`
	<style>
		@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Outfit:wght@400;500;600;700&display=swap');

		:root {
			/* Enhanced Original Dark Theme (Deep Blue / Teal) */
			--color-page-bg: #0a1520;
			--color-card-bg: #112132;
			--color-card-border: #1a2c3f;
			--color-accent: #00c4a0;
			--color-accent-hover: #00e0b8;
			--color-accent-muted: rgba(0, 196, 160, 0.1);
			--color-primary-text: #ffffff;
			--color-secondary-text: #94a3b8;
			--color-tertiary-text: #64748b;
			--color-divider: rgba(255, 255, 255, 0.08);
			--color-zebra: rgba(255, 255, 255, 0.02);
			--color-hover: rgba(255, 255, 255, 0.04);
			--color-warning: #ffb432;
			--color-danger: #ff5a5a;
			--color-success: #00c4a0;
			--color-info: #3b82f6;

			/* Shadows */
			--shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
			--shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.2), 0 2px 4px -1px rgba(0, 0, 0, 0.1);
			--shadow-md: 0 10px 15px -3px rgba(0, 0, 0, 0.3), 0 4px 6px -2px rgba(0, 0, 0, 0.1);
			--shadow-lg: 0 20px 25px -5px rgba(0, 0, 0, 0.4), 0 10px 10px -5px rgba(0, 0, 0, 0.2);
			
			/* Border Radius */
			--radius-sm: 0.375rem;
			--radius: 0.5rem;
			--radius-md: 0.75rem;
			--radius-lg: 1rem;
			--radius-xl: 1.5rem;

			--font-main: 'Inter', sans-serif;
			--font-display: 'Outfit', sans-serif;
		}

		[data-theme=light] {
			/* Modern Slate & Lucid Teal Palette (Optimized for UX) */
			--color-page-bg: #e2e8f0;
			--color-card-bg: #f8fafc;
			--color-card-border: #cbd5e1;
			--color-accent: #14b8a6;
			--color-accent-hover: #0d9488;
			--color-accent-muted: rgba(20, 184, 166, 0.1);
			--color-primary-text: #0f172a;
			--color-secondary-text: #475569;
			--color-tertiary-text: #94a3b8;
			--color-divider: rgba(15, 23, 42, 0.08);
			--color-zebra: rgba(15, 23, 42, 0.03); 
			--color-hover: rgba(15, 23, 42, 0.05);
			--color-warning: #f59e0b;
			--color-error: #ef4444;
			--color-success: #10b981;
			--color-info: #6366f1;

			/* Refined Soft UI Shadows */
			--shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
			--shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05), 0 2px 4px -1px rgba(0, 0, 0, 0.03);
			--shadow-md: 0 10px 15px -3px rgba(0, 0, 0, 0.08), 0 4px 6px -2px rgba(0, 0, 0, 0.04);
			--shadow-lg: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
		}

		body {
			background-color: var(--color-page-bg) !important;
			color: var(--color-primary-text) !important;
			font-family: var(--font-main);
			transition: background-color 0.4s cubic-bezier(0.4, 0, 0.2, 1), color 0.4s cubic-bezier(0.4, 0, 0.2, 1);
			-webkit-font-smoothing: antialiased;
			-moz-osx-font-smoothing: grayscale;
		}

		h1, h2, h3, h4, h5, h6 {
			font-family: var(--font-display);
			font-weight: 700;
			letter-spacing: -0.02em;
		}

		.card {
			background-color: var(--color-card-bg);
			border: 1px solid var(--color-card-border);
			border-radius: var(--radius-lg);
			box-shadow: var(--shadow);
			transition: transform 0.2s ease, box-shadow 0.2s ease, background-color 0.3s ease;
		}

		.btn {
			border-radius: var(--radius);
			font-weight: 600;
			text-transform: none;
			letter-spacing: 0.01em;
			transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
		}

		::-webkit-scrollbar {
			width: 10px;
			height: 10px;
		}
		::-webkit-scrollbar-track {
			background: var(--color-page-bg);
		}
		::-webkit-scrollbar-thumb {
			background: var(--color-divider);
			border-radius: 5px;
			border: 2px solid var(--color-page-bg);
		}
		::-webkit-scrollbar-thumb:hover {
			background: var(--color-tertiary-text);
		}
	</style>
	`)
}

func ThemeInitScript() Node {
	return Script(Raw(`
		(function() {
			const theme = localStorage.getItem('theme') || 'dark';
			document.documentElement.setAttribute('data-theme', theme);
		})();
	`))
}
