package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upInitRate, downInitRate)
}

func upInitRate(tx *sql.Tx) error {
	const query = `
	create table rates
	(
		id integer generated always as identity,
		code       text,
		nominal    bigint,
		kopecks    bigint,
		original   text,
		ts         timestamp,
		created_at timestamp,
		updated_at timestamp
	);
	
	create index rates_code_ts_idx on rates(code, ts);
	`

	_, err := tx.Exec(query)
	return err
}

func downInitRate(tx *sql.Tx) error {
	const query = `
	drop index rates_code_ts_idx;
	drop table rates;
	`

	_, err := tx.Exec(query)
	return err
}
