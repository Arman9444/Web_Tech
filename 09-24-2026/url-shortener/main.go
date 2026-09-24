// curl.exe -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{\"url\":\"https://go.dev\"}'
// curl.exe -i http://localhost:8080/code


package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var urls = map[string]string{}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func randomCode() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func shorten(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		http.Error(w, "url must start with http:// or https://", http.StatusBadRequest)
		return
	}

	code := randomCode()
	for urls[code] != "" {
		code = randomCode()
	}
	urls[code] = req.URL

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{ShortURL: "http://localhost:8080/" + code})
}

func redirect(w http.ResponseWriter, r *http.Request) {
	url, ok := urls[r.PathValue("code")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	http.HandleFunc("POST /shorten", shorten)
	http.HandleFunc("GET /{code}", redirect)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}