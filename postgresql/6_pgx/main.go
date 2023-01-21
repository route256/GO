package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx := context.Background()

	conn, _ := pgx.Connect(ctx, "user=mpak dbname=test sslmode=disable")
	if err := conn.Ping(ctx); err != nil {
		log.Fatal("error pinging db: ", err)
	}

	pool, _ := pgxpool.Connect(ctx, "user=mpak dbname=test sslmode=disable")
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("error pinging db: ", err)
	}

	db, _ := sql.Open("postgres", "user=mpak dbname=test sslmode=disable")
	if err := db.Ping(); err != nil {
		log.Fatal("error pinging db: ", err)
	}

	var (
		tx  pgx.Tx
		err error
	)
	if tx, err = conn.Begin(ctx); err != nil {
		log.Fatal("error beginning tx: ", err)
	}
	defer func() {
		if err = tx.Commit(ctx); err != nil {
			tx.Rollback(ctx)
		}
	}()

}
