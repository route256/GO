package note

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"route256/ws5/internal/model"
)

type database interface {
	Get(ctx context.Context, dest interface{}, sql string, args ...interface{}) error
	Create(ctx context.Context, sql string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, sql string, args ...interface{}) error
}

func New(db database) *Storage {
	return &Storage{
		db: db,
	}
}

type Storage struct {
	db database
}

func (s *Storage) Create(ctx context.Context, info model.NoteInfo) (uint64, error) {
	var result uint64

	createNoteSQL := "INSERT INTO note( title, content, user_id) VALUES ($1, $2, $3) RETURNING id;"
	if err := s.db.Get(ctx, &result, createNoteSQL,
		info.Title,
		info.Content,
		info.UserID,
	); err != nil {
		return 0, errors.Wrap(err, "error while creating note")
	}

	createLogSQL := "INSERT INTO logs( note_id, content)VALUES ($1, $2)"
	if err := s.db.Create(ctx, createLogSQL,
		result,
		"example information for note creation",
	); err != nil {
		return 0, errors.Wrap(err, "error while creating log")
	}

	return result, nil
}

func (s *Storage) List(ctx context.Context, user uint64) ([]model.Note, error) {
	sql := `SELECT id, user_id, title, content, created_at, updated_at FROM note WHERE user_id=$1 ORDER BY created_at`

	type Row struct {
		ID        uint64    `db:"id"`
		UserID    uint64    `db:"user_id"`
		Title     string    `db:"title"`
		Content   string    `db:"content"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	var rows []Row
	if err := s.db.Select(ctx, &rows, sql, user); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return []model.Note{}, nil
		}
		return nil, errors.Wrap(err, "error while selecting notes")
	}

	result := make([]model.Note, 0, len(rows))
	for _, row := range rows {
		result = append(result, model.Note{
			ID: row.ID,
			Info: model.NoteInfo{
				Title:   row.Title,
				Content: row.Content,
				UserID:  row.UserID,
			},
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}

	return result, nil
}
