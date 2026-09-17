package model

import (
	"time"
)

type Note struct {
	ID        uint64
	Info      NoteInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NoteInfo struct {
	UserID  uint64
	Title   string
	Content string
}
