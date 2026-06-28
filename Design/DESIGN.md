---
name: OddsLab
colors:
  surface: '#fbf9fa'
  surface-dim: '#dbd9db'
  surface-bright: '#fbf9fa'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#f5f3f4'
  surface-container: '#efedef'
  surface-container-high: '#e9e7e9'
  surface-container-highest: '#e4e2e3'
  on-surface: '#1b1c1d'
  on-surface-variant: '#43474c'
  inverse-surface: '#303032'
  inverse-on-surface: '#f2f0f2'
  outline: '#74777d'
  outline-variant: '#c4c6cd'
  surface-tint: '#4f6073'
  primary: '#172839'
  on-primary: '#ffffff'
  primary-container: '#2d3e50'
  on-primary-container: '#97a9be'
  inverse-primary: '#b6c8df'
  secondary: '#5f5e5b'
  on-secondary: '#ffffff'
  secondary-container: '#e1dfdb'
  on-secondary-container: '#63635f'
  tertiary: '#3c2100'
  on-tertiary: '#ffffff'
  tertiary-container: '#56360f'
  on-tertiary-container: '#cd9f6f'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
  primary-fixed: '#d2e4fb'
  primary-fixed-dim: '#b6c8df'
  on-primary-fixed: '#0a1d2d'
  on-primary-fixed-variant: '#37485b'
  secondary-fixed: '#e4e2dd'
  secondary-fixed-dim: '#c8c6c2'
  on-secondary-fixed: '#1b1c19'
  on-secondary-fixed-variant: '#474744'
  tertiary-fixed: '#ffdcbc'
  tertiary-fixed-dim: '#efbd8a'
  on-tertiary-fixed: '#2c1700'
  on-tertiary-fixed-variant: '#614018'
  background: '#fbf9fa'
  on-background: '#1b1c1d'
  surface-variant: '#e4e2e3'
typography:
  headline-lg:
    fontFamily: Inter
    fontSize: 36px
    fontWeight: '700'
    lineHeight: '1.2'
    letterSpacing: -0.02em
  headline-md:
    fontFamily: Inter
    fontSize: 28px
    fontWeight: '600'
    lineHeight: '1.3'
    letterSpacing: -0.01em
  headline-sm:
    fontFamily: Inter
    fontSize: 20px
    fontWeight: '600'
    lineHeight: '1.4'
  body-md:
    fontFamily: Inter
    fontSize: 16px
    fontWeight: '400'
    lineHeight: '1.6'
  body-sm:
    fontFamily: Inter
    fontSize: 14px
    fontWeight: '400'
    lineHeight: '1.5'
  label-xs:
    fontFamily: Inter
    fontSize: 12px
    fontWeight: '500'
    lineHeight: '1.2'
    letterSpacing: 0.02em
  headline-lg-mobile:
    fontFamily: Inter
    fontSize: 28px
    fontWeight: '700'
    lineHeight: '1.2'
rounded:
  sm: 0.125rem
  DEFAULT: 0.25rem
  md: 0.375rem
  lg: 0.5rem
  xl: 0.75rem
  full: 9999px
spacing:
  base: 8px
  xs: 4px
  sm: 12px
  md: 24px
  lg: 40px
  xl: 64px
  gutter: 24px
  margin-mobile: 16px
  margin-desktop: 48px
---

## Brand & Style
The design system is anchored in a philosophy of **Sophisticated Minimalism**, tailored for a prediction market simulator that prioritizes clarity over hype. It rejects the aggressive, neon-soaked aesthetics of typical "fintech" or "crypto" platforms in favor of a warm, academic, and professional environment.

The target audience consists of strategic thinkers and learners who value intellectual rigor. The UI evokes a sense of calm confidence and trustworthiness through generous whitespace, high-quality typography, and a "warm-analog" digital feel. It blends **Corporate Modern** structure with a **Tactile** warmth, moving away from cold interfaces toward something that feels like a premium financial journal or a high-end educational tool.

## Colors
The palette is centered on a high-contrast relationship between **Deep Slate** (authority) and **Warm Cream** (comfort). Unlike pure white-and-gray systems, the use of cream backgrounds reduces eye strain and reinforces the "simulator" and "learning" aspect of the brand.

- **Primary & Secondary:** Used for the core brand identity and structural grounding.
- **Accents:** **Amber Gold** is reserved for highlights and achievement moments. **Emerald Green** and **Rust Coral** are used semantically for market fluctuations (up/down) and success/error states, but are slightly desaturated to maintain the sophisticated tone.
- **Neutrals:** A range of warm grays ensures that even functional elements like borders and backgrounds feel integrated into the "warm" aesthetic.

## Typography
This design system utilizes **Inter** exclusively to maintain a systematic, utilitarian, and highly legible appearance. The hierarchy is strictly enforced to guide the user through complex market data.

- **Headlines:** Use a slightly tighter letter-spacing and heavier weights to provide a sturdy anchor for page content.
- **Body Text:** Set with a generous line-height (1.6) to ensure the "educational" content is easy to digest.
- **Labels:** Small and Tiny variants use a medium weight (500) to remain legible even at high densities, typical for data tables or market tickers.

## Layout & Spacing
The layout follows a **Fluid Grid** model with a max-width container for desktop environments. We use an 8px base unit for all spatial relationships.

- **Grid:** A 12-column grid for desktop, 8-column for tablet, and 4-column for mobile.
- **Rhythm:** Use `md` (24px) for most internal component padding and `lg` (40px) for vertical section spacing.
- **Adaptive Rules:** On mobile, margins shrink to 16px to maximize data visualization space, while component heights remain consistent (min 44px for touch targets).

## Elevation & Depth
Depth is conveyed through **Tonal Layers** and subtle shadows rather than aggressive gradients. 

- **Level 0 (Base):** The `Warm Cream` background.
- **Level 1 (Cards/Containers):** Flat white background with a 1px `Light Gray` border.
- **Level 2 (Interactive/Floating):** Use a soft, diffused shadow: `0 2px 8px rgba(45,62,80,0.08)`. This creates a subtle "lift" for interactive elements like hover-state cards or dropdown menus.
- **Layering:** Avoid stacking multiple shadows. Use the `Very Light Gray` surface for nested containers (like input fields inside a card) to create "wells" rather than "peaks."

## Shapes
The shape language is **Soft and Precise**. It uses a standard 6px corner radius for most functional elements to strike a balance between professional geometry and approachable warmth.

- **Standard (6px):** Buttons and Input fields.
- **Large (8px):** Market cards and main containers.
- **Pill:** Reserved exclusively for status indicators (e.g., "Open", "Closed", "Live") to distinguish them from interactive buttons.

## Components
- **Buttons:** 
    - *Primary:* Deep Slate fill, Warm Cream text. Bold and authoritative.
    - *Success/Danger:* Emerald or Rust fill with Warm Cream text. 6px radius.
    - *Secondary:* Transparent fill, 1px Deep Slate border.
- **Cards:** White background, 8px radius, 1px Light Gray border. Use the subtle shadow only on the primary "Market" cards to emphasize their importance.
- **Inputs:** Very Light Gray background with a 1px Light Gray border. On focus, the border transitions to Deep Slate.
- **Chips/Status:** Smaller text (12px), pill-shaped, using low-opacity versions of the accent colors for the background.
- **Lists:** Clean rows with 1px bottom borders in `Light Gray`. No alternating row colors; use whitespace for separation.
- **Portfolio Tracker:** A custom component using a `Deep Slate` header with `Amber Gold` data points to emphasize value and growth.