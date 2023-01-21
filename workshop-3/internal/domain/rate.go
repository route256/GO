package domain

import (
	"time"
)

type Rate struct {
	ID        int
	Code      string
	Nominal   int64
	Kopecks   int64
	Original  string
	Ts        time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
