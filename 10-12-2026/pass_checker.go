package main

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := HealthResponse{
		StatusCode: http.StatusOK,
		Message:    "ok",
	}
	json.NewEncoder(w).Encode(resp)
}

type user struct {
	id       string
	password string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login endpoint reached (not implemented yet)"))
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	w.Write([]byte("login endpoint reached (not implemented yet)"))

}

func main() {
	http.HandleFunc("/api/v1/health", health)
	http.HandleFunc("POST /api/v1/login", login)
	http.ListenAndServe(":8080", nil)
}
