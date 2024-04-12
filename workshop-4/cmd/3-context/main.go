package main

import (
	"context"
	"fmt"
	"time"
)

type privateStruct struct{}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, privateStruct{}, 5)
	v := ctx.Value(privateStruct{}).(int)
	fmt.Println(v)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	// done channel
	<-ctx.Done()
}
