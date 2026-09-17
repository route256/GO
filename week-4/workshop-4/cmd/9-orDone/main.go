package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)

	consume(ctx, ch)
}

func orDone[T any](done <-chan struct{}, anyCh <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return

			case v, ok := <-anyCh:
				if !ok {
					return
				}

				select {
				case out <- v:
				case <-done:
				}
			}
		}
	}()

	return out
}

func consume(ctx context.Context, inSink chan int) {
	for v := range orDone(ctx.Done(), inSink) {
		fmt.Println(v)
	}
}
