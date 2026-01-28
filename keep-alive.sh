#!/bin/bash

# Configuration
URL="https://idohr.app/healthz"
LOG_FILE="/tmp/idohr-keep-alive.log"

# Perform ping
RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$URL")

# Log result (optional, can be disabled to save disk space)
echo "$(date): Ping $URL - Response: $RESPONSE" >> "$LOG_FILE"

# Keep log size under control (last 100 lines)
tail -n 100 "$LOG_FILE" > "$LOG_FILE.tmp" && mv "$LOG_FILE.tmp" "$LOG_FILE"
