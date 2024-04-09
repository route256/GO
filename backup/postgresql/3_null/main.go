package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var (
		ctx    = context.Background()
		dsn    = "user=mpak dbname=test sslmode=disable"
		driver = "postgres"
	)

	db, err := sql.Open(driver, dsn)
	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	var foo sql.NullInt32
	if err = db.QueryRow("select null;").Scan(&foo); err != nil {
		log.Fatal(err)
	}
	log.Print(foo)
}
