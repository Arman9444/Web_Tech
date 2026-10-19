// curl.exe -i http://localhost:8081/hello

package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

const limit = 5

var counts = map[string]int{}
var start = time.Now()

func hello(w http.ResponseWriter, r *http.Request) {
	if time.Since(start) > time.Minute {
		counts = map[string]int{}
		start = time.Now()
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	counts[ip]++
	if counts[ip] > limit {
		http.Error(w, "too many requests", http.StatusTooManyRequests)
		return
	}

	w.Write([]byte("hello\n"))
}

func main() {
	http.HandleFunc("/hello", hello)

	log.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
