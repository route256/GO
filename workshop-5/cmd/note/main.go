package main

import (
	"context"
	"log"

	"route256/ws5/cmd/note/app"
)

func main() {
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("cannot create app: %s", err.Error())
	}

	if err = a.Run(ctx); err != nil {
		log.Fatalf("cannot run app: %s", err.Error())
	}
}
