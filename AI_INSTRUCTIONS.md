# AI Coding Instructions for AdoptionOS

**Role:** Senior Full-Stack Engineer (Go/Vue).
**Goal:** Maintain a clean, modular monolith with strict type safety between Backend (Go) and Frontend (TypeScript).

---

## 1. Project Overview & Architecture

**AdoptionOS** is a shelter management ERP running on Raspberry Pi 5.

- **Repository:** Monorepo (`backend/` and `frontend/`).
- **Backend:** Go 1.24 (Standard Library `net/http`), PostgreSQL 16 (JSONB-heavy schema).
- **Frontend:** Vue 3 (Composition API), TypeScript, Pinia, Vite.
- **Data Flow:** REST API -> Frontend Fetch -> Pinia State Management -> Vue Reactivity.

---

## 2. Getting Started & Development

### 2.1 Initial Setup
To install all dependencies for both frontend and backend, run this from the root directory:
```bash
make setup
```

### 2.2 Running the Development Servers
Run the backend and frontend servers in separate terminals.

- **Run Backend (from root):**
  ```bash
  make run-be
  ```
- **Run Frontend (from root):**
  ```bash
  make run-fe
  ```

---

## 3. Backend Guidelines (Go)

**Location:** `/backend`

### 3.1 Core Principles

- **No Frameworks:** Use standard `net/http` and `http.ServeMux`. Do not introduce Gin, Echo, or Fiber.
- **Database:** Use `database/sql` with `lib/pq`.
  - **JSONB:** Use `JSONB` columns for complex nested data (e.g., `medical`, `physical`).
  - **Coalesce:** Always usage `COALESCE(column, '{}')` in SQL queries to prevent `null` pointer crashes in Go structs.
- **Error Handling:** Return JSON errors. Do not panic. Use the `serverError` or `clientError` helpers in `cmd/api/helpers.go`.

### 3.2 File Structure & Key Files

- `cmd/api/`: Application entry point and routing.
  - `main.go`: Server startup.
  - `routes.go`: All API routes defined here.
  - `handlers.go`: HTTP handlers for the routes.
- `internal/data/`: SQL queries and struct definitions.
- `internal/models/`: **Source of Truth** for data types. Structs here must match DB columns exactly.
- `migrations/`: SQL up/down files. **Never modify the DB without creating a migration file.**

### 3.3 Style & Commands

- **Variable Names:** Short and descriptive (e.g., `mux`, `db`, `pet`).
- **JSON Tags:** All structs exported to API must have `` `json:"camelCase"` `` tags.
- **Handlers:** Naming convention `Func(app *Application) HandlerName(w http.ResponseWriter, r *http.Request)`.
- **Formatting & Dependencies:** Run from `backend/`:
  ```bash
  make tidy
  ```

---

## 4. Frontend Guidelines (Vue 3)

**Location:** `/frontend`

### 4.1 Core Principles

- **Composition API:** Always use `<script setup lang="ts">`. Options API is forbidden.
- **Type Safety:** strict `ts` mode. No `any`.
- **State Management:** Use **Pinia** for all global state. Create new stores in `src/stores/`.
- **Data Fetching:**
  - Use `fetch` API for network requests.
  - API Base URL is relative (e.g., `/pets` or `/api/v1/pets`) as the proxy handles routing.

### 4.2 Component Structure

- **Order:** `<script setup>`, then `<template>`, then `<style scoped>`.
- **Props:** Define using TypeScript interfaces.
  ```ts
  // Example
  import type { IPet } from '@/models/pet'; // Adjust path as needed
  defineProps<{ pet: IPet }>();
  ```

### 4.3 Testing, Linting & Formatting

Run all commands from within the `frontend/` directory.

- **Unit Tests:**
  ```bash
  npm run test:unit
  ```
- **Linting (with auto-fix):**
  ```bash
  npm run lint
  ```
- **Formatting:**
  ```bash
  npm run format
  ```