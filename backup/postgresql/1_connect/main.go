package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var (
		ctx = context.Background()
		dsn = "user=mpak dbname=test sslmode=disable"
	)

	db, err := sql.Open("postgres", dsn)
	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	conn, err := db.Conn(ctx)
	defer conn.Close()
	if err = conn.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	defer tx.Rollback()
}

func tx(ctx context.Context, db *sql.DB) (err error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	return
}
