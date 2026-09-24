package main

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// Each IP can make at most 5 requests per minute
const maxRequests = 5
const window = time.Minute

type client struct {
	count int       // requests made in the current minute
	start time.Time // when the current minute started
}

// Stores IP -> its request count
var clients = map[string]*client{}
var clientsMu sync.Mutex

// allow returns true if this IP is still under the limit
func allow(ip string) bool {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	c, ok := clients[ip]

	// New IP, or its minute is over -> start counting again from 1
	if !ok || time.Since(c.start) > window {
		clients[ip] = &client{count: 1, start: time.Now()}
		return true
	}

	// Already used all requests for this minute
	if c.count >= maxRequests {
		return false
	}

	c.count++
	return true
}

// rateLimit wraps a handler and blocks the request if the IP is over the limit
func rateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		if !allow(ip) {
			http.Error(w, "too many requests, try again later", http.StatusTooManyRequests)
			return
		}

		next(w, r)
	}
}

// GET /hello   a simple endpoint to test the rate limiter
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, your request was allowed\n"))
}

func main() {
	mux := http.NewServeMux()

	// Routes (protected by the rate limiter)
	mux.HandleFunc("GET /hello", rateLimit(hello))

	log.Println("rate limiter running on :8081")

	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
