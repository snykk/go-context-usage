package main

import (
	"context"
	"fmt"
	"time"
)

func generateNumbers(ctx context.Context, out chan<- int) {
	defer close(out)
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Generator stopped")
			return
		case out <- i:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func squareNumbers(ctx context.Context, in <-chan int, out chan<- int) {
	defer close(out)
	for n := range in {
		select {
		case <-ctx.Done():
			fmt.Println("Square stopped")
			return
		case out <- n * n:
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	numbers := make(chan int)
	squares := make(chan int)

	go generateNumbers(ctx, numbers)
	go squareNumbers(ctx, numbers, squares)

	for n := range squares {
		fmt.Println("Result:", n)
	}
}
