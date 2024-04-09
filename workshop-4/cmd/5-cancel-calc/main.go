package main

import (
	"context"
	"fmt"
)

type privateStruct struct{}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result, err := pricyOperation(ctx)
	fmt.Printf("result: %v\n", result)
	fmt.Printf("err: %v\n", err)
}

func pricyOperation(ctx context.Context) (int, error) {
	out := 0

	for i := 0; i < 1_000_000_000; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			out += i
		}

	}

	return out, nil
}
