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
	fmt.Println("CHECKING AUTH")
	apiKey := r.Header.Get("Key")
	expectedAPIKey := os.Getenv("LIBERO_API_KEY")
	fmt.Println(apiKey)
	// expectedAPIKey := "hello"
	if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedAPIKey)) != 1 {
		http.Error(w, "Invalid API key", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
