package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("HelloWorld")

	ctx := context.Background()

	value := 1234
	ctx = context.WithValue(ctx, "current_user", value)
}
