package main

import (
	"context"
	"fmt"
	"time"
)

func processBatch(ctx context.Context, data []int) error {
	for _, v := range data {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Printf("Processing item %d\n", v)
			time.Sleep(500 * time.Millisecond)
		}

		// Simulasi error jika data == 5
		if v == 5 {
			return fmt.Errorf("critical error on item %d", v)
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data := []int{1, 2, 3, 4, 5, 6, 7}

	go func() {
		if err := processBatch(ctx, data); err != nil {
			fmt.Println("Batch processing error:", err)
			cancel() // Batalkan proses jika ada error
		}
	}()

	// Simulasi pekerjaan lain
	time.Sleep(5 * time.Second)
	fmt.Println("Main program finished")
}
