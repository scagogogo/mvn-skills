// Custom theme configuration for mvn-skills website
// Flat design: no purple, no gradients, small border-radius, Slate neutral palette

const theme = {
  // Primary colors — Blue-600 based, no purple
  colors: {
    primary: '#2563EB',        // Blue-600
    primaryDark: '#1D4ED8',    // Blue-700
    primaryLight: '#EFF6FF',   // Blue-50

    accent: '#EA580C',         // Orange-600 (Maven brand)
    accentDark: '#C2410C',     // Orange-700
    accentLight: '#FFF7ED',    // Orange-50

    // Neutrals — Slate palette
    bg: '#FFFFFF',
    bgSecondary: '#F8FAFC',    // Slate-50
    bgDark: '#0F172A',         // Slate-900
    bgDarkSecondary: '#1E293B', // Slate-800

    text: '#0F172A',           // Slate-900
    textSecondary: '#475569',  // Slate-600
    textMuted: '#94A3B8',      // Slate-400
    textLight: '#F8FAFC',      // Slate-50
    textLightSecondary: '#94A3B8', // Slate-400

    border: '#E2E8F0',         // Slate-200
    borderLight: '#F1F5F9',    // Slate-100
    borderDark: '#334155',     // Slate-700

    // Status
    success: '#16A34A',        // Green-600
    warning: '#CA8A04',        // Yellow-600
    error: '#DC2626',          // Red-600
    info: '#0284C7',           // Sky-600
  },

  // Border radius — flat design, strictly small
  radius: {
    xs: '2px',
    sm: '4px',
    md: '6px',
  },

  // Spacing — 8px grid system
  spacing: {
    xs: '4px',
    sm: '8px',
    md: '16px',
    lg: '24px',
    xl: '32px',
    xxl: '48px',
    section: '80px',
  },

  // Typography
  font: {
    family: `-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif`,
    mono: `"SF Mono", "Fira Code", "Fira Mono", "Roboto Mono", Menlo, Monaco, Consolas, monospace`,
    size: {
      xs: '12px',
      sm: '14px',
      base: '16px',
      lg: '18px',
      xl: '20px',
      xxl: '24px',
      hero: '32px',
      display: '48px',
    },
  },

  // Shadows — flat design: none or minimal
  shadow: {
    none: 'none',
    sm: '0 1px 2px rgba(0, 0, 0, 0.05)',
  },

  // Layout
  layout: {
    maxWidth: '1200px',
    headerHeight: '64px',
  },
};

export default theme;
