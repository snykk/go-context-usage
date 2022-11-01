package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set deadline untuk request ini
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	ch := make(chan string, 1)
	go func() {
		// Simulasi pemrosesan data
		time.Sleep(3 * time.Second)
		ch <- "Data processed"
	}()

	select {
	case res := <-ch:
		fmt.Fprintln(w, res)
	case <-ctx.Done():
		http.Error(w, "Request timeout", http.StatusGatewayTimeout)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
