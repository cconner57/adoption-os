package main

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func (app *application) rateLimit(next http.Handler) http.Handler {
	// Define a client struct to hold the rate limiter and last seen time
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	// Launch a background goroutine to remove old entries from the clients map.
	go func() {
		for {
			time.Sleep(10 * time.Minute)

			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only rate limit if enabled in config?
		// For now, always on as per request.

		// Extract the IP address from the request.
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		// Lock the mutex to protect this block from race conditions.
		mu.Lock()

		if _, found := clients[ip]; !found {
			// Create a new limiter: 1 request per second, with burst of 5.
			clients[ip] = &client{
				limiter: rate.NewLimiter(1, 5),
			}
		}

		// Update the last seen time for the client.
		clients[ip].lastSeen = time.Now()

		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			app.errorResponse(w, r, http.StatusTooManyRequests, "rate limit exceeded")
			return
		}

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
