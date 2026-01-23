---
trigger: model_decision
description: Project Hardware & Environment Specs (Pi 5 + Pico 2W)
---

* **Server Hardware:** Raspberry Pi 5.
* **Microcontroller:** Raspberry Pi Pico 2 W (RP2350 chip). 
* **Storage Reality:** The SD card is for the OS only. The NVMe drive is mounted at `/mnt/nvme/adoption-os`.
* **Data Persistence:** ALL database files (SQLite), uploaded assets, and pet photos MUST be stored on the NVMe path (`/mnt/nvme/adoption-os/...`).
* **Network:** The public URL is `idohr.app` (served via Cloudflare Tunnel).
* **Frontend Config:** The Vite server must allow the host `idohr.app`.