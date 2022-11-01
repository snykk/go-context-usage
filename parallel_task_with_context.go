package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workerWG(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerWG(ctx, i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers stopped")
}
