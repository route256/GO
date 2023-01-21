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

	var res sql.Result
	if res, err = db.Exec("delete from a where aid = $1", 1); err != nil {
		log.Fatal(err)
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		log.Fatal("Not Found")
	}
	// res.LastInsertId()

	var ID int
	row := db.QueryRow("insert into a (aid) values $1 returning aid", 1)
	err = row.Scan(&ID)
	// err = row.Err()
	if err == sql.ErrNoRows {
		log.Fatal("Not Found")
	}
	// err = db.QueryRow("insert into a (aid) values $1 returning aid", 1).Scan(&ID)

	var IDs []int
	rows, err := db.Query("select aid from a")
	defer rows.Close()
	for rows.Next() {
		var ID int
		err = rows.Scan(&ID, &ID)
		IDs = append(IDs, ID)
	}
	err = rows.Err()
}
