package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", checkAPIKey)
	http.ListenAndServe(":4444", nil)
}

func checkAPIKey(w http.ResponseWriter, r *http.Request) {
	var authd bool = true
	// apiKey := r.Header.Get("Key")
	apiKey := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	// fmt.Println(r.Header)
	expectedAPIKey := os.Getenv("LIBERO_API_KEY")
	if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedAPIKey)) != 1 {
		http.Error(w, "Invalid API key. Set the `Key` header to authorize your request.", http.StatusUnauthorized)
		authd = false
		return
	}
	fmt.Println("FROM", r.Header.Get("X-Forwarded-For"), "AUTHORIZED", authd)
	w.WriteHeader(http.StatusOK)
}
