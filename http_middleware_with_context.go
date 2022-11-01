package main

import (
	"context"
	"fmt"
	"net/http"
)

type key string

const userKey key = "userID"

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulasi mengambil userID dari token
		userID := "12345"
		ctx := context.WithValue(r.Context(), userKey, userID)

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userKey)
	if userID == nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Hello, User %s!\n", userID)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/user", authMiddleware(http.HandlerFunc(userHandler)))

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", mux)
}
