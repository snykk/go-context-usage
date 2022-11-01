package main

import (
	"context"
	"fmt"
)

func main() {
	type key string
	k := key("userID")

	ctx := context.WithValue(context.Background(), k, "12345")

	process(ctx, k)
}

func process(ctx context.Context, k interface{}) {
	if v := ctx.Value(k); v != nil {
		fmt.Println("User ID:", v)
	} else {
		fmt.Println("No value found")
	}
}
