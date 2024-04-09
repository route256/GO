package database

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/domain"
)

type RatesDB struct {
	db *sql.DB
}

func NewRateDB(db *sql.DB) *RatesDB {
	return &RatesDB{db}
}

func (db *RatesDB) AddRate(ctx context.Context, date time.Time, rate domain.Rate) error {
	builder := sq.Insert("rates").Columns(
		"created_at",
		"code",
		"nominal",
		"kopecks",
		"original",
		"ts",
	).Values(
		time.Now(),
		rate.Code,
		rate.Nominal,
		rate.Kopecks,
		rate.Original,
		rate.Ts,
	)
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = db.db.ExecContext(ctx, query, args...)

	return err
}

func (db *RatesDB) GetRate(ctx context.Context, code string, date time.Time) (*domain.Rate, error) {

	builder := sq.Select(
		"id",
		"code",
		"nominal",
		"kopecks",
		"original",
		"ts",
		"created_at",
		"updated_at",
		"deleted_at",
	).From("rates").Where(sq.Eq{"code": code})

	if !date.IsZero() {
		builder = builder.Where(sq.Eq{"ts": date})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var rate domain.Rate
	err = db.db.QueryRowContext(ctx, query, args...).Scan(&rate)
	return &rate, err
}
