---
trigger: model_decision
description: Deployment & Makefile Commands
---

* **Never** suggest running servers manually with `go run` or `npm run dev` for production tasks.
* **ALWAYS** use the Makefile commands:
    * Use `make dev` for local development (runs both FE/BE).
    * Use `make deploy` for production updates (handles zero-downtime swapping).
    * Use `make flash` for the Pico 2 W.