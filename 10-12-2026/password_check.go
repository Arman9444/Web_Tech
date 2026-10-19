package main

import (
	"fmt"
	"net/http"
)

type user struct {
	id       string
	password string
}

func password_checker(w http.ResponseWriter, r *http.Request) {
	correct_password := "12345abcde"
	pass_given_by_user := r.URL.Query().Get("password")
	if correct_password == pass_given_by_user {
		fmt.Fprintln(w, "You have logged in successfully")
	} else {
		fmt.Fprintln(w, "Sorry,your password is not correct !!! please try again")
	}
}
func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/vi/login", password_checker)
	http.ListenAndServe(":8080", nil)
}
