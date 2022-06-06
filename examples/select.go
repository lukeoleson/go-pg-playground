package examples

import (
	"github.com/go-pg/pg/v10"
	"github.com/lukeoleson/go-pg-playground/db/models"
)

// DONE
// Title: SELECT...FROM
// Task: Find all the rows in the actor table returning all columns.
// SQL Text: In SQL, to find all the `actors` in the `actor` table we write:
// SQL Code: SELECT * FROM actor;
// SQL Result:
// actor_id | first_name  |  last_name   |      last_update
// ----------+-------------+--------------+------------------------
// 1 | PENELOPE    | GUINESS      | 2020-02-15 01:34:33-08
// 2 | NICK        | WAHLBERG     | 2020-02-15 01:34:33-08
// 3 | ED          | CHASE        | 2020-02-15 01:34:33-08
// ...
// 198 | MARY        | KEITEL       | 2020-02-15 01:34:33-08
// 199 | JULIA       | FAWCETT      | 2020-02-15 01:34:33-08
// 200 | THORA       | TEMPLE       | 2020-02-15 01:34:33-08
func SelectFrom(db *pg.DB) {

	// The `Model` method determines the table in the query's `FROM` clause, while the `Select` method can specify where to store the returned rows.
	var actors1 []models.Actor
	err := db.Model(&models.Actor{}).Select(&actors1)
	models.PrintActorsOrError(err, actors1)

	// The receiving variable can alternatively be specified in the `Model` method where it will serve to determine the table in the query's `FROM` clause as well as receive the returned rows.
	var actors2 []models.Actor
	err = db.Model(&actors2).Select()
	models.PrintActorsOrError(err, actors2)

	// Providing a variable to both the `Model` and `Select` methods will store the returned rows in the variable passed to `Select` leaving the variable passed into `Model` unchanged.
	var actors3 []models.Actor
	var actors4 []models.Actor
	err = db.Model(&actors3).Select(&actors4)
	models.PrintActorsOrError(err, actors3, actors4)
}
