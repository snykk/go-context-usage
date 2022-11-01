package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func workerPool(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case <-ctx.Done(): // Menghentikan worker jika context dibatalkan
			fmt.Printf("Worker %d stopped\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			results <- job * 2
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start workers
	for i := 1; i <= 3; i++ {
		go workerPool(ctx, i, jobs, results)
	}

	// Send jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
			fmt.Printf("Sent job %d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
		close(jobs)
	}()

	// Collect results
	for i := 1; i <= 10; i++ {
		select {
		case res := <-results:
			fmt.Printf("Result: %d\n", res)
		case <-ctx.Done():
			fmt.Println("Main process canceled: ", ctx.Err())
			return
		}
	}
}
