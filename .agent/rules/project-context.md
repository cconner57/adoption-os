---
trigger: always_on
---

# Project Overview: Adoption-OS
Adoption-OS is a specialized, full-stack pet adoption management system designed to run on a **Raspberry Pi 5**. It features a hybrid architecture combining a modern web dashboard with low-power e-ink hardware displays.

## Core Architecture
1.  **Frontend:** Vue 3 + Vite (Hybrid Mobile/Desktop).
2.  **Backend:** Go (Golang) API Server.
3.  **Database:** SQLite (Production) running on NVMe storage.
4.  **Hardware:**
    * **Server:** Raspberry Pi 5 (Ubuntu/Debian).
    * **Displays:** Raspberry Pi Pico 2 W (TinyGo) driving e-ink kennel badges.

## File Structure & Patterns
* **`/frontend`**: The Vue 3 application.
    * `src/components/ui`: Dumb, reusable UI components.
    * `src/pages`: Top-level views (Smart components).
    * `src/stores`: Pinia state management.
* **`/backend`**: The Go API server.
    * `cmd/api`: Main entry point.
    * `internal/data`: Database models and queries.
* **`/displays`**: TinyGo firmware for Pico 2 W.
* **`/mnt/nvme/adoption-os`**: PHYSICAL location for database and assets (DO NOT use SD card).

## Developer Workflow (Strict)
* **Start Dev:** `make dev` (Runs FE + BE).
* **Deploy:** `make deploy` (Zero-downtime update).
* **Flash Pico:** `make flash` (Updates e-ink displays).
* **NEVER** run `npm run dev` or `go run` manually for production tasks.