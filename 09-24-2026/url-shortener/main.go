package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
)

// Stores short code -> original URL (kept in memory)
var urls = map[string]string{}
var urlsMu sync.Mutex

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Make a random 6 character code like "aZ3kP9"
func generateCode() string {
	code := make([]byte, 6)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

// Shorten Request
type ShortenRequest struct {
	URL string `json:"url"`
}

// Shorten Response
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// POST /shorten   body: {"url": "https://google.com"}
func shorten(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// URL must start with http:// or https://
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		http.Error(w, "url must start with http:// or https://", http.StatusBadRequest)
		return
	}

	urlsMu.Lock()
	code := generateCode()
	// If the code is already used, make a new one
	for urls[code] != "" {
		code = generateCode()
	}
	urls[code] = req.URL
	urlsMu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{
		ShortURL: "http://localhost:8080/" + code,
	})
}

// GET /{code}   sends the user to the original URL
func redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	urlsMu.Lock()
	url, ok := urls[code]
	urlsMu.Unlock()

	if !ok {
		http.Error(w, "short url not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("POST /shorten", shorten)
	mux.HandleFunc("GET /{code}", redirect)

	log.Println("url shortener running on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
