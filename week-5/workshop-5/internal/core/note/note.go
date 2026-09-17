package note

import (
	"context"

	"github.com/pkg/errors"

	"route256/ws5/internal/model"
)

type storage interface {
	Create(ctx context.Context, info model.NoteInfo) (uint64, error)
	List(ctx context.Context, user uint64) ([]model.Note, error)
}

func New(s storage) *Core {
	return &Core{
		storage: s,
	}
}

type Core struct {
	storage storage
}

func (c *Core) Create(ctx context.Context, req model.NoteInfo) (uint64, error) {
	switch {
	case req.UserID == 0:
		return 0, errors.Wrapf(model.ErrInvalidRequest, "cannot be empty, field: [user_id]")
	case req.Title == "":
		return 0, errors.Wrapf(model.ErrInvalidRequest, "cannot be empty, field: [title]")
	case req.Content == "":
		return 0, errors.Wrapf(model.ErrInvalidRequest, "cannot be empty, field: [content]")
	}

	result, err := c.storage.Create(ctx, req)
	if err != nil {
		return 0, errors.Wrap(err, "error while creating note in storage")
	}

	return result, nil
}

func (c *Core) List(ctx context.Context, userID uint64) ([]model.Note, error) {
	notes, err := c.storage.List(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get list of notes")
	}

	return notes, nil
}
