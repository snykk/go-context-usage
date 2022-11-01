package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context, name string, duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Printf("Task %s completed\n", name)
	case <-ctx.Done():
		fmt.Printf("Task %s canceled: %v\n", name, ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Task pertama yang bergantung pada task kedua
	go func() {
		task(ctx, "Task 1", 2*time.Second)
		cancel() // Membatalkan semua task lain jika Task 1 selesai
	}()

	// Task kedua
	go task(ctx, "Task 2", 5*time.Second)

	time.Sleep(6 * time.Second)
	fmt.Println("All tasks finished")
}
