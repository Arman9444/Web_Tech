package main

import "net/http"

func userHandler() {

}
func main() {
	http.HandleFunc("/arman", userHandler)
	http.ListenAndServe(":8084", nil)
}
