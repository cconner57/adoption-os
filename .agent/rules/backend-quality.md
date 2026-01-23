---
trigger: model_decision
description: Go Error Handling & Logging Standards
---

* **No Panics:** NEVER use `panic()` in the backend code. All errors must be handled gracefully and returned to the caller.
* **Explicit Handling:** Do not use `_` to ignore errors. Always check `if err != nil`.
* **Structured Logging:**
    * Use the standard `log` package or `slog`.
    * Log errors with context (e.g., `log.Printf("Failed to fetch pet %s: %v", id, err)`), not just `log.Println(err)`.
* **HTTP Errors:** When an API handler fails, return a JSON error response (e.g., `{"error": "message"}`) with the correct HTTP status code (400 vs 500).