package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "ok")
		w.WriteHeader(200)
	})
	a := 10
	b := 20
	http.HandleFunc("/api/v1/add", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "%v + %v = %v", a, b, a+b)
		w.WriteHeader(http.StatusOK)
	})
	log.Println("Backend server live on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
