package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", checkAPIKey)
	http.ListenAndServe(":4444", nil)
}

func checkAPIKey(w http.ResponseWriter, r *http.Request) {
	var authd bool = true
	apiKey := r.Header.Get("Key")
	expectedAPIKey := os.Getenv("LIBERO_API_KEY")
	if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedAPIKey)) != 1 {
		http.Error(w, "Invalid API key", http.StatusUnauthorized)
		authd = false
		return
	}
	fmt.Println("FROM", r.Header.Get("X-Forwarded-For"), "AUTHORIZED", authd)
	w.WriteHeader(http.StatusOK)
}
