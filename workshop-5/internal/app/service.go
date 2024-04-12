package app

import (
	"context"

	"route256/ws5/internal/model"
	pb "route256/ws5/pkg"
)

func New(note notesService) *Implementation {
	return &Implementation{
		note: note,
	}
}

type Implementation struct {
	pb.UnimplementedNoteServer

	note notesService
}

type notesService interface {
	Create(ctx context.Context, note model.NoteInfo) (uint64, error)
	List(ctx context.Context, user uint64) ([]model.Note, error)
}
