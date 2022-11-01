package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Parent context with timeout
	parentCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Child context with additional value
	childCtx := context.WithValue(parentCtx, "operationID", "12345")

	go doTask(childCtx)

	// Simulate some work
	time.Sleep(6 * time.Second)
	fmt.Println("Main finished")
}

func doTask(ctx context.Context) {
	operationID := ctx.Value("operationID")
	fmt.Println("Starting task with operationID:", operationID)

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		fmt.Println("Task canceled due to:", ctx.Err())
	}
}
