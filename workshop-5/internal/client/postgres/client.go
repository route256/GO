package postgres

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"

	"route256/ws5/internal/model"
)

type Config struct {
	DSN                   string
	MaxConnections        int32
	MaxConnectionIdleTime time.Duration
	PingTotal             int
	PingDelay             time.Duration
}

func New(ctx context.Context, config Config) (*Adapter, error) {
	poolConfig, err := pgxpool.ParseConfig(config.DSN)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = config.MaxConnections
	poolConfig.MaxConnIdleTime = config.MaxConnectionIdleTime

	client, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	for i := 0; i < config.PingTotal; i++ {
		if err = client.Ping(ctx); err != nil {
			time.Sleep(config.PingDelay)
			continue
		}
		break
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to database")
	}
	return &Adapter{
		client: client,
	}, nil
}

type Adapter struct {
	client *pgxpool.Pool
}

func (c *Adapter) Close() {
	c.client.Close()
}

func (c *Adapter) Get(ctx context.Context, dest interface{}, sql string, args ...interface{}) error {
	var p provider = c.client
	row := p.QueryRow(ctx, replaceSQLChars(sql), args...)
	if err := row.Scan(dest); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.ErrNotFound
		}
		return errors.Wrap(err, "error while scanning value")
	}

	return nil
}

func (c *Adapter) Select(ctx context.Context, dest interface{}, sql string, args ...interface{}) error {
	var p provider = c.client
	rows, err := p.Query(ctx, replaceSQLChars(sql), args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.ErrNotFound
		}
		return errors.Wrap(err, "error while executing sql")
	}

	return pgxscan.ScanAll(dest, rows)
}

func (c *Adapter) Create(ctx context.Context, sql string, args ...interface{}) error {
	var p provider = c.client
	if _, err := p.Exec(ctx, replaceSQLChars(sql), args...); err != nil {
		return errors.Wrap(err, "error while executing sql")
	}

	return nil
}

func (c *Adapter) Delete(ctx context.Context, sql string, args ...interface{}) error {
	var p provider = c.client
	if _, err := p.Exec(ctx, replaceSQLChars(sql), args...); err != nil {
		return errors.Wrap(err, "error while executing sql")
	}

	return nil
}

type provider interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

var rxSpaces = regexp.MustCompile(`\s+`)

// replaceSQLChars remove all non-printable symbols (as a \t, \n) for more comfortable reading SQL ("humanity").
func replaceSQLChars(sql string) string {
	sql = strings.Replace(sql, "\n", " ", -1)
	sql = strings.Replace(sql, "\t", " ", -1)
	sql = strings.Replace(sql, " ,", ",", -1)
	sql = strings.Replace(sql, ", ", ",", -1)

	return strings.TrimSpace(rxSpaces.ReplaceAllString(sql, " "))
}
