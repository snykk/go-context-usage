package main

import (
	"context"
	"fmt"
	"time"
)

func streamData(ctx context.Context, ch chan int) {
	defer close(ch)
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Streaming stopped:", ctx.Err())
			return
		case ch <- i:
			time.Sleep(500 * time.Millisecond) // Simulasi delay streaming
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch := make(chan int)
	go streamData(ctx, ch)

	for v := range ch {
		fmt.Println("Received:", v)
	}
}
