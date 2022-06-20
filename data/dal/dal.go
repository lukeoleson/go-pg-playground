package dal

import (
	pgV10 "github.com/go-pg/pg/v10"
	pgV9 "github.com/go-pg/pg/v9"
	"github.com/lukeoleson/go-pg-playground/logging"
)

func ConnectV10() *pgV10.DB {
	db := pgV10.Connect(&pgV10.Options{
		Addr:     ":5432",
		User:     "lukeoleson",
		Password: "pass",
		Database: "pagila",
	})

	db.AddQueryHook(logging.QueryLoggerV10{})

	return db
}

func ConnectV9() *pgV9.DB {
	db := pgV9.Connect(&pgV9.Options{
		Addr:     ":5432",
		User:     "lukeoleson",
		Password: "pass",
		Database: "pagila",
	})

	db.AddQueryHook(logging.QueryLoggerV9{})

	return db
}
