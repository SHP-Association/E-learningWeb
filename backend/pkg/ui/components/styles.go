package components

import (
	. "maragu.dev/gomponents"
)

type (
	Color int
	Size  int
)

const (
	ColorNone Color = iota
	ColorNeutral
	ColorPrimary
	ColorSecondary
	ColorAccent
	ColorInfo
	ColorSuccess
	ColorWarning
	ColorError
	ColorLink
)

const (
	SizeExtraSmall Size = iota
	SizeSmall
	SizeMedium
	SizeLarge
	SizeExtraLarge
)

func PremiumStyles() Node {
	return Raw(`
	<style>
		:root {
			--page-bg: var(--color-page-bg);
			--card-bg: var(--color-card-bg);
			--table-bg: var(--color-card-bg);
			
			/* Define premium default (dark mode) zebra & hover */
			--color-zebra: rgba(255, 255, 255, 0.015);
			--color-hover: rgba(255, 255, 255, 0.04);
			
			--table-zebra: var(--color-zebra);
			--table-hover: var(--color-hover);
			--accent: var(--color-accent);
			--accent-soft: color-mix(in oklab, var(--color-accent) 15%, transparent);
			--accent-glow: color-mix(in oklab, var(--color-accent) 35%, transparent);
			--accent-muted: color-mix(in oklab, var(--color-accent) 25%, transparent);
			--primary-text: var(--color-primary-text);
			--secondary-text: var(--color-secondary-text);
			--divider: var(--color-divider);
			--danger: var(--color-error);
			
			/* Premium Elevation Tokens */
			--shadow-lume: 0 0 20px -5px var(--accent-glow);
			--glass-border: color-mix(in oklab, var(--color-divider) 50%, transparent);
			
			/* Theme-Responsive Form Inputs */
			--input-bg: rgba(0, 0, 0, 0.25);
			--input-border: rgba(255, 255, 255, 0.08);
			--input-border-hover: rgba(255, 255, 255, 0.16);
		}
		
		[data-theme=light] {
			--color-zebra: rgba(15, 23, 42, 0.02) !important;
			--color-hover: rgba(15, 23, 42, 0.04) !important;
			--input-bg: rgba(15, 23, 42, 0.035) !important;
			--input-border: rgba(15, 23, 42, 0.1) !important;
			--input-border-hover: rgba(15, 23, 42, 0.2) !important;
		}
		
		.glass-frost {
			background: color-mix(in oklab, var(--color-card-bg) 70%, transparent) !important;
			backdrop-filter: blur(16px) saturate(180%) !important;
			-webkit-backdrop-filter: blur(16px) saturate(180%) !important;
			border: 1px solid var(--glass-border) !important;
		}

		.glass-modern {
			background: color-mix(in oklab, var(--color-card-bg) 85%, transparent) !important;
			backdrop-filter: blur(12px) !important;
			-webkit-backdrop-filter: blur(12px) !important;
			border: 1px solid var(--color-divider) !important;
			box-shadow: var(--shadow-md) !important;
		}
		[data-theme=light] .glass-modern, [data-theme=light] .glass-frost {
			background: var(--color-card-bg) !important;
			backdrop-filter: none !important;
			-webkit-backdrop-filter: none !important;
			border: 1.5px solid var(--color-card-border) !important;
		}

		.teal-lume {
			box-shadow: var(--shadow-lume) !important;
			border-color: var(--color-accent) !important;
		}
		.yellow-lume {
			box-shadow: 0 0 20px -5px rgba(245, 158, 11, 0.4) !important;
			border-color: var(--color-brand-yellow) !important;
		}

		.btn-brand {
			background: var(--color-accent) !important;
			color: #ffffff !important;
			font-family: var(--font-main) !important;
			font-weight: 700 !important;
			border: none !important;
			border-radius: 14px !important;
			transition: all 0.4s cubic-bezier(0.19, 1, 0.22, 1) !important;
			box-shadow: 0 4px 12px var(--color-accent-muted) !important;
		}
		.btn-brand:hover {
			filter: brightness(1.1) !important;
			transform: translateY(-2px) !important;
			box-shadow: 0 8px 20px -5px var(--color-accent-muted) !important;
		}
		.btn-brand:active { transform: translateY(0) scale(0.96) !important; }

		.admin-form-input {
			width: 100% !important;
			background: var(--input-bg) !important;
			border: 1.5px solid var(--input-border) !important;
			padding: 12px 16px !important;
			border-radius: 14px !important;
			color: var(--color-primary-text) !important;
			font-size: 13px !important;
			transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1) !important;
			outline: none !important;
			box-shadow: inset 0 2px 4px 0 rgba(0, 0, 0, 0.05) !important;
		}
		.admin-form-input:hover {
			border-color: var(--input-border-hover) !important;
			background: var(--input-bg) !important;
		}
		.admin-form-input:focus {
			border-color: var(--color-accent) !important;
			background: var(--color-card-bg) !important;
			box-shadow: 0 0 0 4px var(--accent-soft) !important;
		}

		.flex-col.gap-1\.5 label {
			font-family: 'Poppins', sans-serif !important;
			font-size: 13px !important;
			font-weight: 600 !important;
			color: var(--color-secondary-text) !important;
			letter-spacing: 0.015em !important;
			transition: color 0.2s ease !important;
		}
		.flex-col.gap-1\.5:focus-within label {
			color: var(--color-accent) !important;
		}

		.btn-teal {
			background: var(--color-accent) !important;
			color: #ffffff !important;
			font-family: var(--font-main) !important;
			font-weight: 700 !important;
			text-transform: uppercase !important;
			letter-spacing: 0.1em !important;
			border: none !important;
			border-radius: 14px !important;
			transition: all 0.4s cubic-bezier(0.19, 1, 0.22, 1) !important;
			box-shadow: 0 4px 12px var(--color-accent-muted) !important;
		}
		.btn-teal:hover {
			filter: brightness(1.1) !important;
			transform: translateY(-2px) !important;
			box-shadow: 0 12px 24px -8px var(--color-accent) !important;
		}
		.btn-teal:active { transform: translateY(0) scale(0.96) !important; }

		.btn-danger {
			background: var(--danger) !important;
			color: #ffffff !important;
			font-family: var(--font-main) !important;
			font-weight: 700 !important;
			text-transform: uppercase !important;
			letter-spacing: 0.1em !important;
			border: none !important;
			border-radius: 14px !important;
			transition: all 0.4s cubic-bezier(0.19, 1, 0.22, 1) !important;
		}
		.btn-danger:hover {
			filter: brightness(1.15) !important;
			transform: translateY(-2px) !important;
			box-shadow: 0 12px 24px -8px var(--danger) !important;
		}

		.admin-card {
			background: var(--color-card-bg) !important;
			border: 1px solid var(--color-divider) !important;
			border-radius: 24px !important;
			transition: all 0.4s cubic-bezier(0.23, 1, 0.32, 1) !important;
			box-shadow: var(--shadow-sm) !important;
			overflow: hidden !important;
		}
		.admin-card:hover {
			border-color: var(--color-accent) !important;
			box-shadow: 12px 14px 40px -10px rgba(0,0,0,0.4) !important;
			transform: translateY(-5px) !important;
		}
		[data-theme=dark] .admin-card {
			background: color-mix(in oklab, var(--color-card-bg) 95%, transparent) !important;
			backdrop-filter: blur(8px);
		}

		.sticky-glass {
			position: sticky !important;
			top: 0 !important;
			z-index: 40 !important;
			background: color-mix(in oklab, var(--color-page-bg) 80%, transparent) !important;
			backdrop-filter: blur(12px) !important;
			border-bottom: 1px solid var(--color-divider) !important;
		}

		.custom-scrollbar::-webkit-scrollbar { width: 6px; height: 6px; }
		.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
		.custom-scrollbar::-webkit-scrollbar-thumb { 
			background: color-mix(in oklab, var(--color-divider) 50%, transparent); 
			border-radius: 10px; 
		}
		.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: var(--color-accent); }

		/* Improved Table Sticky Logic - GUARANTEED OPAQUE */
		.sticky-column-left {
			position: sticky !important;
			left: 0 !important;
			z-index: 20 !important;
			background-color: var(--color-card-bg) !important;
			box-shadow: 12px 0 16px -8px rgba(0,0,0,0.3) !important;
			transition: all 0.3s ease !important;
			opacity: 1 !important;
			backdrop-filter: none !important;
		}
		.sticky-column-left::before {
			content: "" !important;
			position: absolute !important;
			inset: 0 !important;
			z-index: -1 !important;
			background-color: var(--color-card-bg) !important;
		}
		.admin-table tr:nth-child(even) .sticky-column-left { 
			background-color: var(--color-card-bg) !important;
			background-image: linear-gradient(var(--color-zebra), var(--color-zebra)) !important;
		}
		.admin-table tr:hover .sticky-column-left { 
			background-color: var(--color-card-bg) !important;
			background-image: linear-gradient(var(--color-hover), var(--color-hover)) !important;
		}
		thead th.sticky-column-left { 
			background-color: var(--color-card-bg) !important;
			z-index: 30 !important; 
			opacity: 1 !important;
		}

		.sticky-column-right {
			position: sticky !important;
			right: 0 !important;
			z-index: 20 !important;
			background-color: var(--color-card-bg) !important;
			box-shadow: -12px 0 16px -8px rgba(0,0,0,0.3) !important;
			transition: all 0.3s ease !important;
			opacity: 1 !important;
			backdrop-filter: none !important;
		}
		.sticky-column-right::before {
			content: "" !important;
			position: absolute !important;
			inset: 0 !important;
			z-index: -1 !important;
			background-color: var(--color-card-bg) !important;
		}
		.admin-table tr:nth-child(even) .sticky-column-right { 
			background-color: var(--color-card-bg) !important;
			background-image: linear-gradient(var(--color-zebra), var(--color-zebra)) !important;
		}
		.admin-table tr:hover .sticky-column-right { 
			background-color: var(--color-card-bg) !important;
			background-image: linear-gradient(var(--color-hover), var(--color-hover)) !important;
		}
		thead th.sticky-column-right { 
			background-color: var(--color-card-bg) !important;
			z-index: 30 !important; 
			opacity: 1 !important;
		}
		
		.admin-table {
			width: 100%;
			border-collapse: separate !important;
			border-spacing: 0 !important;
		}
		.admin-table th {
			background-color: var(--color-card-bg) !important;
			color: var(--color-secondary-text) !important;
			font-weight: 700 !important;
			font-size: 11px !important;
			text-transform: uppercase !important;
			letter-spacing: 0.1em !important;
			padding: 16px 24px !important;
			border-bottom: 1.5px solid var(--color-divider) !important;
			white-space: nowrap !important;
		}
		.admin-table td {
			padding: 16px 24px !important;
			color: var(--color-primary-text) !important;
			font-size: 13px !important;
			border-bottom: 1px solid var(--color-divider) !important;
			white-space: nowrap !important;
		}
		
		.admin-table-container {
			border-radius: 20px !important;
			border: 1px solid var(--color-divider) !important;
			overflow-x: auto !important;
			overflow-y: hidden !important;
			scroll-behavior: smooth;
		}
		.admin-table-container::-webkit-scrollbar { height: 6px; }
		.admin-table-container::-webkit-scrollbar-track { background: transparent; }
		.admin-table-container::-webkit-scrollbar-thumb {
			background: color-mix(in oklab, var(--color-accent) 40%, transparent);
			border-radius: 10px;
		}
		.admin-table-container::-webkit-scrollbar-thumb:hover { background: var(--color-accent); }
		
		.table-scroll-btn {
			position: absolute;
			top: 50%;
			transform: translateY(-50%);
			z-index: 40;
			width: 32px;
			height: 32px;
			display: flex;
			align-items: center;
			justify-content: center;
			border-radius: 50%;
			border: 1px solid var(--color-divider);
			background-color: var(--color-card-bg);
			color: var(--color-accent);
			cursor: pointer;
			transition: all 0.2s ease;
			box-shadow: 0 2px 12px rgba(0,0,0,0.4);
			opacity: 0;
			pointer-events: none;
		}
		.table-scroll-btn.visible { opacity: 1; pointer-events: auto; }
		.table-scroll-btn:hover {
			background-color: var(--color-accent);
			color: var(--color-page-bg);
			transform: translateY(-50%) scale(1.1);
		}
		.table-scroll-btn-left  { left:  -16px; }
		.table-scroll-btn-right { right: -16px; }

		/* Overridden by the main definition above */
		
		.admin-form-section-header {
			display: flex;
			align-items: center;
			gap: 16px;
			margin-bottom: 24px;
		}
		.admin-form-section-title {
			font-family: var(--font-main);
			font-size: 11px;
			font-weight: 800;
			text-transform: uppercase;
			letter-spacing: 0.25em;
			color: var(--color-accent);
			white-space: nowrap;
		}
		.admin-form-section-divider {
			height: 1px;
			background: var(--color-divider);
			flex: 1;
			opacity: 0.3;
		}

		@keyframes shake {
			0%, 100% { transform: translateX(0); }
			25% { transform: translateX(-5px); }
			75% { transform: translateX(5px); }
		}
		.animate-shake { animation: shake 0.4s cubic-bezier(.36,.07,.19,.97) both; }
		
		@keyframes pulse-subtle {
			0%, 100% { opacity: 1; }
			50% { opacity: 0.7; }
		}
		.animate-pulse-subtle { animation: pulse-subtle 2s ease-in-out infinite; }

		.drop-shadow-glow {
			filter: drop-shadow(0 0 8px var(--accent-glow));
		}
	</style>
	`)
}
