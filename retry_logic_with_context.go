package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func unreliableTask() error {
	if rand.Float32() < 0.7 {
		return errors.New("task failed")
	}
	return nil
}

func retryWithContext(ctx context.Context, maxRetries int, task func() error) error {
	var lastErr error
	for i := 0; i < maxRetries; i++ {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled: %w", ctx.Err())
		default:
			lastErr = task()
			if lastErr == nil {
				return nil
			}
			fmt.Printf("Retry %d failed: %v\n", i+1, lastErr)
			time.Sleep(500 * time.Millisecond)
		}
	}
	return fmt.Errorf("all retries failed: %w", lastErr)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := retryWithContext(ctx, 5, unreliableTask)
	if err != nil {
		fmt.Println("Task failed:", err)
	} else {
		fmt.Println("Task succeeded")
	}
}
