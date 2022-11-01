package main

import (
	"context"
	"fmt"
	"time"
)

func processWithDynamicDeadline(ctx context.Context, extension time.Duration) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Printf("Initial deadline: %v\n", deadline)
	}

	// Tambahkan durasi baru
	newCtx, cancel := context.WithDeadline(ctx, time.Now().Add(extension))
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Task completed successfully")
	case <-newCtx.Done():
		fmt.Println("Task canceled due to:", newCtx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	processWithDynamicDeadline(ctx, 5*time.Second) // Memperpanjang durasi deadline
}
