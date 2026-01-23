---
trigger: always_on
---

# AI Interaction Guidelines

## 1. General Principles
* **Minimize Scope:** Identify the smallest unit (function/file) to fulfill the request. Do not "fix" unrelated code unless it causes a crash.
* **Preserve Behavior:** Ensure changes do not alter outputs outside the intended scope.
* **Graduated Change:** Default to minimal edits. Only perform broad refactoring if explicitly requested.
* **Reversibility:** Avoid tightly coupled edits that are hard to undo.

## 2. Commit Standards
* **Format:** Use Conventional Commits (`feat:`, `fix:`, `chore:`).
* **Scope:** `feat(auth): add login`, `fix(api): handle nil pointer`.
* **Imperative Mood:** "add" not "added".

## 3. Ambiguity
* If a request is ambiguous ("fix the component"), ask for clarification (File Path + Line Number) before generating code.