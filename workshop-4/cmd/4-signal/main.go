package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

type privateStruct struct{}

func main() {
	// notify := make(chan os.Signal)

	// signal.Notify(notify, syscall.SIGINT)

	// fmt.Println("waiting for signal")
	// <-notify
	// fmt.Println("got signal")

	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT)
	defer cancel()

	fmt.Println("waiting for signal")
	<-ctx.Done()
	fmt.Println("got signal")
}
