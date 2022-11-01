package main

import (
	"context"
	"fmt"
)

type MetadataKey string

func main() {
	// Simulasi RPC dengan metadata
	metadata := map[string]string{
		"Authorization": "Bearer abc123",
		"RequestID":     "req-45678",
	}

	ctx := context.Background()
	ctx = withMetadata(ctx, metadata)

	handleRequest(ctx)
}

func withMetadata(ctx context.Context, metadata map[string]string) context.Context {
	for k, v := range metadata {
		ctx = context.WithValue(ctx, MetadataKey(k), v)
	}
	return ctx
}

func handleRequest(ctx context.Context) {
	auth := ctx.Value(MetadataKey("Authorization"))
	reqID := ctx.Value(MetadataKey("RequestID"))

	if auth == nil || reqID == nil {
		fmt.Println("Missing metadata")
		return
	}

	fmt.Printf("Processing request %s with token %s\n", reqID, auth)
}
