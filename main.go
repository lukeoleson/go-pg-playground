package main

import (
	"github.com/go-pg/pg/v10"
)

//
// main
//
var db *pg.DB

func main() {
	//db = pg.Connect(&pg.Options{
	//	Addr:     ":5432",
	//	User:     "lukeoleson",
	//	Password: "pass",
	//	Database: "pagila",
	//})
	//defer db.Close()
	//
	//db.AddQueryHook(logging.QueryLogger{})

	//randomized()
	//examples.Update(db)
}
