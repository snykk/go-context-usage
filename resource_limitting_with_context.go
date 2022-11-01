package main

import (
	"context"
	"fmt"
	"time"
)

func limitedResource(ctx context.Context, limit chan struct{}) {
	select {
	case limit <- struct{}{}:
		// Simulasi penggunaan resource
		fmt.Println("Resource acquired")
		time.Sleep(1 * time.Second)
		<-limit // Bebaskan resource
		fmt.Println("Resource released")
	case <-ctx.Done():
		fmt.Println("Failed to acquire resource:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	limit := make(chan struct{}, 2) // Maksimum 2 goroutine yang bisa mengakses resource
	for i := 0; i < 5; i++ {
		go limitedResource(ctx, limit)
	}

	time.Sleep(6 * time.Second)
}
