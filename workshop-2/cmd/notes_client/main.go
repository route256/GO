package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/client/notes"
	"gitlab.ozon.dev/go/classroom-8/students/workshop-2/internal/model"
	desc "gitlab.ozon.dev/go/classroom-8/students/workshop-2/pkg/api/notes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(os.Args[1], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	grpcNotesClient := desc.NewNotesClient(conn)

	adapter := notes.NewClient(grpcNotesClient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := adapter.SaveNote(ctx, &model.Note{
		Title:   gofakeit.CarModel(),
		Content: gofakeit.CarMaker(),
	})

	if err != nil {
		log.Fatalf("failed to send create note request: %v", err)
	}
	log.Printf("Added note ID: %d\n", id)

	rs, err := adapter.ListNotes(ctx)

	for _, n := range rs {
		log.Printf("ID: %v\n", n.Id)
		log.Printf("Title: %v\n", n.Title)
		log.Printf("Content: %v\n", n.Content)
	}
}
