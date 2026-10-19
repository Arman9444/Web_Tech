package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("arman-me")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type User struct {
	Id   string `json:"username"`
	Pass string `json:"password"`
}

func heartbit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintln(w, `{"status code": 200}, {"status": "ok"}`)
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user := User{
		Id: "Arman", Pass: "123456",
	}
	if req.Id != user.Id || req.Pass != user.Pass {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	log.Println("login:", req.Id)

	expiresAt := time.Now().Add(15 * time.Minute)
	claims := Claims{
		Username: req.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "could not create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": signedToken})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/heartbit", heartbit)
	mux.HandleFunc("/api/v1/login", login)
	log.Println("2200 server running")
	http.ListenAndServe(":2200", mux)
}
