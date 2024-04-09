package main

import (
	"context"
	"database/sql"
	"fmt"
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

	const query = `
		select abalance
		from pgbench_accounts
		where bid = %s
	`
	// arg := "58 or 1 = 1"
	arg := "'58';DROP schema public--"

	var foo int
	if err = db.QueryRow(fmt.Sprintf(query, arg)).Scan(&foo); err != nil {
		log.Fatal(err)
	}
	log.Print(foo)

}
