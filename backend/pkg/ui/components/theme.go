package components

import (
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui/icons"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ThemeToggle() Node {
	return Label(
		Class("swap swap-rotate text-secondary-text hover:text-white transition-colors cursor-pointer"),
		Input(
			Type("checkbox"),
			Class("theme-controller"),
			Attr("onchange", "(function(e) { const theme = e.target.checked ? 'light' : 'dark'; document.documentElement.setAttribute('data-theme', theme); localStorage.setItem('theme', theme); })(event)"),
			ID("theme-toggle-checkbox"),
		),
		// Sun icon (for light mode)
		Span(Class("swap-on fill-none stroke-current transition-transform duration-500"), icons.Icon("Sun", "w-4 h-4")),
		// Moon icon (for dark mode)
		Span(Class("swap-off fill-none stroke-current transition-transform duration-500"), icons.Icon("Moon", "w-4 h-4")),
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
		@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700;800;900&display=swap');

		:root {
			/* SlingSight-inspired Premium Palette */
			--color-page-bg: #070c14;
			--color-card-bg: #0d1622;
			--color-card-border: rgba(255, 255, 255, 0.05);
			
			/* Accents: Teal & Cyan */
			--color-accent: #2ec4b6;           /* Primary Teal */
			--color-accent-hover: #26a69a;
			--color-accent-muted: rgba(46, 196, 182, 0.1);
			
			--color-primary-text: #f8fafc;     /* Slate 50 */
			--color-secondary-text: #94a3b8;   /* Slate 400 */
			--color-tertiary-text: #64748b;    /* Slate 500 */
			
			--color-divider: rgba(255, 255, 255, 0.04);
			--color-hover: rgba(255, 255, 255, 0.03);
			
			/* Semantic */
			--color-danger: #ff4d4d;
			--color-success: #2ec4b6;
			
			/* Shell Gradients - Exact matching */
			--ui-shell-start: #0a1622;
			--ui-shell-end: #070c14;

			/* Shadows */
			--shadow-premium: 0 0 50px -12px rgba(0, 0, 0, 0.5);
			--shadow-glow: 0 0 15px -3px rgba(46, 196, 182, 0.4);
			
			--font-main: 'Poppins', sans-serif;
			--font-display: 'Poppins', sans-serif;
		}

		[data-theme=light] {
			--color-page-bg: #f8fafc;
			--color-card-bg: #ffffff;
			--color-card-border: #e2e8f0;
			--color-accent: #0d9488;
			--color-primary-text: #0f172a;
			--color-secondary-text: #475569;
			--color-tertiary-text: #94a3b8;
			--color-divider: rgba(15, 23, 42, 0.08);
			--ui-shell-start: #ffffff;
			--ui-shell-end: #f8fafc;
		}

		body {
			background-color: var(--color-page-bg) !important;
			color: var(--color-primary-text) !important;
			font-family: var(--font-main);
			-webkit-font-smoothing: antialiased;
			overflow-x: hidden;
		}

		/* Premium Glassmorphism */
		.glass-modern {
			background: linear-gradient(180deg, var(--ui-shell-start), var(--ui-shell-end));
			backdrop-filter: blur(20px);
			-webkit-backdrop-filter: blur(20px);
			border-right: 1px solid var(--color-card-border);
			box-shadow: var(--shadow-premium);
		}

		/* Sidebar Link Overhaul - Exact Matching */
		.sidebar-link {
			position: relative;
			transition: all 0.2s ease;
		}

		.sidebar-link.active {
			background: rgba(46, 196, 182, 0.08);
			color: var(--color-accent) !important;
		}
		
		.sidebar-link.active Span {
			color: var(--color-accent) !important;
		}

		.sidebar-link:hover:not(.active) {
			background: var(--color-hover);
			color: var(--color-primary-text);
		}

		/* NavGroup Vertical Line */
		.nav-group-border {
			border-left: 1px solid rgba(255, 255, 255, 0.1);
			margin-left: 20px;
			padding-left: 10px;
		}

		/* Typography Utilities */
		.tracking-ultra {
			letter-spacing: 0.15em;
		}

		/* Ultra-thin Scrollbar */
		.custom-scrollbar::-webkit-scrollbar {
			width: 3px;
		}
		.custom-scrollbar::-webkit-scrollbar-track {
			background: transparent;
		}
		.custom-scrollbar::-webkit-scrollbar-thumb {
			background: var(--color-divider);
			border-radius: 10px;
		}

		/* Interactive Elements */
		.drop-shadow-glow {
			filter: drop-shadow(0 0 5px var(--color-accent));
		}

		/* Logout Special Case */
		.logout-btn {
			color: var(--color-danger);
			opacity: 0.8;
			transition: all 0.3s ease;
		}
		.logout-btn:hover {
			opacity: 1;
			background: rgba(255, 77, 77, 0.05);
			transform: translateX(4px);
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
