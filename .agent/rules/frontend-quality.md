---
trigger: always_on
---

* **Code Quality & Linting (Strict):**
    * **No 'any' Type:** You MUST NOT use the `any` type in Frontend (TS) or Backend (Go `interface{}`). Use specific interfaces, types, or generics.
    * **No AI Comments:** Do NOT leave explanatory comments (e.g., `// This function does X`) in the code. Code must be self-documenting.
    * **Exception:** You MAY add comments strictly to suppress unavoidable linting warnings (e.g., `// eslint-disable-next-line`).
    * **500-Line Cap:** Files MUST NOT exceed 500 lines. Refactor logic into Composables (`usePetLogic.ts`) or sub-components immediately if you hit this limit.
    * **Zero-Lint Policy:** You are responsible for fixing all ESLint and TypeScript errors. Do not leave unused variables or missing imports.

* **Component Architecture (Smart vs. Dumb):**
    * **UI Components (src/components/ui):** Must be "Dumb." They strictly receive data via **Props** and communicate actions via **Emits**. They MUST NOT make API calls or access Pinia stores directly.
    * **Smart Components (Views/Pages):** Handle the logic, API calls, and state management, passing data down to dumb components.

* **Layout Strategy (Hybrid):**
    * **Public Pages:** Mobile-First. Default CSS is for phones; use media queries for larger screens.
    * **Admin/Dashboard:** Desktop-First. Optimized for high-density data, but must remain functional on mobile.

* **Styling & Theming (Strict variables):**
    * **Palette:**
        * `--color-primary` (#00a5ad) /* Teal */
        * `--color-secondary` (#1976d2) /* Blue */
        * `--color-tertiary` (#6b5b95) /* Purple */
        * `--color-danger` (#c73a67) /* Red/Error */
        * `--color-warning` (#deb024) /* Yellow */
        * `--color-neutral` (#1f2121) /* Dark Grey */
        * `--text-primary` (var(--color-neutral))
        * `--text-secondary` (#374151)
        * `--text-inverse` (#f8f8f3)
    * **Color Logic:** Use Relative Color Syntax for variants (e.g., `hsl(from var(--color-primary) h s 90%)`). NO magic hex codes.
    * **Spacing & Sizing:** Use `rem` units exclusively. Align to a 4px grid (0.25rem, 0.5rem, 1rem). DO NOT use random pixel values (e.g., `13px`).
    * **Typography:** Use project CSS variables for font-sizes (e.g., `var(--text-xl)`), do not hardcode pixels.
    * **Nesting:** Use CSS nesting syntax (e.g., `.card { &:hover { ... } }`).

* **Assets & Icons (No 3rd Party Libs):**
    * **Iconography:** Do NOT use external icon libraries. Use **Raw SVGs** inside a `<template>` or a simple custom Icon component.
    * **SVG Coloring:** All SVGs must use `fill="currentColor"` or `stroke="currentColor"` to inherit the parent text color.
    * **Images:** All images below the fold must use `loading="lazy"`.

* **Vue Best Practices:**
    * **Computed Over Logic:** Move complex logic from `<template>` to `computed` properties.
    * **Naming:** Multi-word component names required (e.g., `PetCard.vue`).

* **Accessibility (A11y) - CRITICAL:**
    * **Inputs:** Must have `<label>` or `aria-label`. Placeholders are NOT labels.
    * **Focus:** Never remove `outline: none` without a custom `:focus-visible` replacement.
    * **Semantics:** Use semantic tags (`<button>`, `<nav>`) over `<div>`.