package dal

import (
	"github.com/go-pg/pg/v10"
	"github.com/lukeoleson/go-pg-playground/logging"
)

func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "lukeoleson",
		Password: "pass",
		Database: "pagila",
	})

	db.AddQueryHook(logging.QueryLogger{})

	return db
}
