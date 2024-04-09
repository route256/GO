package main

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Foo struct {
	Bid int `db:"bad"`
	T   string
}

func main() {
	var (
		ctx    = context.Background()
		dsn    = "user=mpak dbname=test sslmode=disable"
		driver = "postgres"
	)

	db, err := sqlx.ConnectContext(ctx, driver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	const query = `select bid, t from b limit 5;`

	var tmps []Foo
	// var tmps = make(map[string]interface{})
	rows, err := db.QueryxContext(ctx, query)
	for rows.Next() {
		var tmp Foo
		err = rows.StructScan(&tmp)
		tmps = append(tmps, tmp)
		// err = rows.MapScan(tmps)
		// dest, err := rows.SliceScan()
	}

	for _, v := range tmps {
		log.Println(v)
	}

	// err = db.SelectContext(ctx, &tmps, query)
	// err = db.GetContext(ctx, &tmp, query)
}
