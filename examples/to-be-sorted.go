package examples

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

//func randomized() {
//	var actors1 []models.Actor
//	_ = db.Model(&models.Actor{}).Where("first_name LIKE ?", "PEN%").Limit(10).OrderExpr("random()").Select(&actors1)
//	for _, actor := range actors1 {
//		fmt.Println(actor)
//	}
//}

// Title: Complex Queries
// TODO: Write something that uses CTEs and Joins and Aggregate functions.

//func selectColumnsFrom() {
//	// Note that for retrieving specific columns we must use the Select() syntax.
//	// Specifying the variables in the model does not tell go-pg what table we want to query.
//	// So this query will throw an error.
//	var firstName, lastName string
//	err = db.Model(&firstName, &lastName).
//		Column("first_name", "last_name").
//		Limit(1).
//		Select()
//
//	fmt.Printf("\nResults\nfirst_name: %v, last_name: %v\n", firstName, lastName)
//
//	// ... while this one will work
//	err = db.Model(&Actor{}).
//		Column("first_name", "last_name").
//		Limit(1).
//		Select(&firstName, &lastName)
//
//	fmt.Printf("\nResults\nfirst_name: %v, last_name: %v\n", firstName, lastName)
//
//	// Specifying which columns to retrieve with the Column() method will return ONLY those
//	// columns, and as such will only write those values to our receiver struct, leaving all the
//	// other values in place.
//	actor3 := Actor{
//		ActorId:   999,
//		FirstName: "LUKE",
//	}
//	err = db.Model(&actor3).
//		Where("actor_id = ?", 1).
//		Select()
//	printActors([]Actor{actor3})
//
//	// The order in which we specify the query Methods between the Model() method and the
//	// Select() method only matters in some cases - e.g. the Where()'s must be ordered correctly
//	// to create the correct AND logical condition, but the Column(), Group(), Order(), Limit(), etc. need
//	// not be.
//	// The query ...
//	var films []Film
//	err = db.Model(&films).
//		ColumnExpr("length").
//		Where("length > ?", 80).
//		Where("length < ?", 120).
//		Group("length").
//		Order("length ASC").
//		Limit(100).
//		Select()
//	printFilms(films)
//
//	// ...is equivalent to...
//	err = db.Model(&films).
//		Limit(100).
//		Order("length ASC").
//		Group("length").
//		Where("length > ?", 80).
//		Where("length < ?", 120).
//		ColumnExpr("length").
//		Select()
//	printFilms(films)
//
//	return err
//}

//
// CTE
//
//func ctes() {
//	var actors []models.Actor
//	pActors := db.Model(&models.Actor{}).Where("first_name LIKE ?", "P%")
//
//	err := db.Model().
//		With("p_actors", pActors).
//		Table("p_actors").
//		ColumnExpr("first_name, last_name").
//		Limit(3).
//		Select()
//
//	checkErr(err)
//
//	printActors(actors)
//}

/*
Title: SELECT col1, col2, ...
*/
//func getActorFirstAndLastNameById(actorId int) (Actor, error) {
//	// The key takeaway here is that ONLY the FirstName and LastName are populated
//	// in the .Model(Actor).
//	var actor Actor
//	// Note that using column here does not add the alias to the column names
//	err := db.Model(&actor).
//		Column("first_name", "last_name").
//		Where("actor_id = ?", actorId).
//		Select()
//
//	printActors([]Actor{actor})
//	// Note that the order of the methods in this method-chain between Model() and Select() is not important.
//	// So, the first query and the second are equivalent
//	err = db.Model(&actor).
//		Where("actor_id = ?", actorId).
//		Column("first_name", "last_name").
//		Select()
//
//	printActors([]Actor{actor})
//
//	// The important note here is that nothing gets written to the Model(&Actor)
//	var firstName, lastName string
//	// Note an alternate syntax for specifying columns with ColumnExpr() where the columns list is expressed as a single string rather than a comma delineated list.
//	err := db.Model(&Actor{}).
//		ColumnExpr("first_name, last_name").
//		Where("actor_id = ?", actorId).
//		Select(&firstName, &lastName)
//	checkErr(err)
//	return firstName, lastName
//
//	return actor, err
//}

/*
Title: WHERE ... LIKE
Problem: Find all actors whose first_name begins with a given letter
SQL Description: The `LIKE` keyword allows us to apply a partial-match filter to our WHERE clause.
SQL Example: To find the actors whose first names start with the letter 'P' we can use the `LIKE` keyword with the `%` placeholder symbol.
SQL Code: SELECT * FROM actor WHERE first_name LIKE 'P%';
actor_id | first_name | last_name |      last_update
----------+------------+-----------+------------------------
        1 | PENELOPE   | GUINESS   | 2020-02-15 01:34:33-08
       46 | PARKER     | GOLDBERG  | 2020-02-15 01:34:33-08
       54 | PENELOPE   | PINKETT   | 2020-02-15 01:34:33-08
      104 | PENELOPE   | CRONYN    | 2020-02-15 01:34:33-08
      120 | PENELOPE   | MONROE    | 2020-02-15 01:34:33-08
Go Code Output:
Generated Query
SELECT a.actor_id, a.first_name, a.last_name, a.last_update
FROM actor AS a
WHERE (first_name LIKE 'P%')

Results
{{} 1 PENELOPE GUINESS 2020-02-15 01:34:33 -0800 PST}
{{} 46 PARKER GOLDBERG 2020-02-15 01:34:33 -0800 PST}
{{} 54 PENELOPE PINKETT 2020-02-15 01:34:33 -0800 PST}
{{} 104 PENELOPE CRONYN 2020-02-15 01:34:33 -0800 PST}
{{} 120 PENELOPE MONROE 2020-02-15 01:34:33 -0800 PST}
*/
//func getActorsByFirstLetterOfFirstName(letter string) error {
//	var actors []models.Actor
//
//	err := db.Model(&models.Actor{}).
//		// Note that the % wildcard is concatenated with the placeholder value rather than placed in the WHERE clauses’ format string.
//		Where("first_name LIKE ?", letter+"%").
//		Select(&actors)
//
//	printActors(actors)
//
//	return err
//}

/*
Title: DISTINCT
Problem: Find a deduplicated list of actor first names which all begin with the same letter.
SQL Description: The DISTINCT keyword allows us to de-duplicate a set of rows based on the values of the specified columns.
SQL Example: To remove all the "PENELOPE"s except 1 from our list of actors whose name begins with 'P', we could use:
SQL Code:
pagila=# SELECT DISTINCT first_name FROM actor WHERE first_name LIKE 'P%';
first_name
------------
PARKER
PENELOPE
Go Code Output:
Generated Query
SELECT DISTINCT first_name
FROM actor AS a
WHERE (first_name LIKE 'P%')

Results
0: {{} 0 PARKER  0001-01-01 00:00:00 +0000 UTC}
1: {{} 0 PENELOPE  0001-01-01 00:00:00 +0000 UTC}
*/
//func getUniqueListOfActorsByFirstLetterOfFirstName(firstName string) error {
//	var actors []models.Actor
//
//	err := db.Model(&actors).
//		Distinct().
//		Column("first_name").
//		Where("first_name LIKE ?", "P%").
//		Select()
//
//	printActors(actors)
//
//	return err
//}

/*
Title:
Problem:
SQL Description:
SQL Example:
SQL Code:
Go Code Output:
*/

/*
The order of the method calls between Model() and Select() is not important outside of readability.
*/

// Model() == FROM
// Column() == SELECT colName1, colName2 (omitting Column == SELECT *)
// Select(&v) == Where to store the retrieved columns (can be a struct or variables);
// 	- specifying nothing in Column will return everything (overwriting everything in your struct)
// 	- specifying something in Column will only write the specified columns into the reciever.
//	- specifying nothing in the Select() will store the returned data into the struct put into Model()
func selectExamples(db *pg.DB) {
	//SELECT DISTINCT column, AGG_FUNC(column_or_expression), …
	//FROM mytable
	//JOIN another_table
	//  ON mytable.column = another_table.column
	//WHERE constraint_expression
	//GROUP BY column
	//HAVING constraint_expression
	//ORDER BY column ASC/DESC
	//LIMIT count OFFSET COUNT;

	//
	// SELECTING DATA
	//

	// Title: SELECT * FROM ...
	// Title: WHERE
	// Title: WHERE ... LIKE
	// Title: SELECT col1, col2, ...
	// Title: DISTINCT
	// Title: LIMIT

	// #4
	// Given a LIMIT, return the corresponding number of actors
	// SELECT * FROM actor LIMIT 10;
	// func getActors(limit int) []Actor
	//
	// #5
	// Find all actors with the given first_name
	// SELECT * FROM actor WHERE first_name = 'Penelope';
	// func getActorsByFirstName(firstName string) []Actor
	//
	// #6
	// Find all actors whose first_name begins with a given letter
	// SELECT * FROM actor WHERE first_name LIKE 'P%';
	// func getActorsByFirstLetterOfFirstName(firstName string) []Actor
	//
	// #7
	// Find a deduplicated list of actor first names where the first name begins with a given letter
	// SELECT DISTINCT first_name FROM actor WHERE first_name LIKE 'P%';
	// func getActorsByFirstLetterOfFirstName(firstName string) []Actor
	//
	// #8
	// Find all actors whose first_name and last_name begin with the given letters
	// SELECT * FROM actor WHERE first_name LIKE 'P%' AND last_name LIKE 'G%'
	// func getActorsByFirstLetterOfFirstName(firstName, lastName string) []Actor
	//
	// #9
	// Find all actors whose first_name OR last_name begins with the given letter
	// SELECT * FROM actor WHERE first_name LIKE 'P%' OR last_name LIKE 'P%'
	// func getActorsByFirstLetterOfFirstOrLastName(firstName, lastName string) []Actor
	//
	// #10
	// Given a list of actor IDs, return the corresponding rows by actor_id from the actor table
	// SELECT * FROM film WHERE film_id IN ('1', '2', '3')
	// func getMoviesByID(filmIDs []string) []Actor

	//
	// AGGREGATE FUNCTIONS
	//

	//
	// JOINS
	//

	//
	// DISTINCT
	//

	//
	// GROUPING DATA
	//
	// Group the movies in the film table by rating
	// SELECT COUNT(id), rating FROM film GROUP BY rating
	// func getMovieCountByRating() []string
	//
	// Find how many movies have what rating from the film table
	// SELECT COUNT(id), rating FROM film GROUP BY rating
	// func getMovieCountByRating() map[int]string

	//
	// TRANSACTIONS
	//

	//
	// LOCKING ROWS
	//

	////
	//// Select all the columns and scan the values into a struct leaving the .Model() untouched
	////
	//actor3 := Actor{
	//	ActorId:   999,
	//	FirstName: "Luke",
	//	LastName:  "Oleson",
	//}
	//
	//printPreQuery(actor3)
	//
	//var actor4 Actor
	//
	//// Instead of scanning the values into actor3, this will store all the values returned in
	//// actor4 leaving actor3 untouched
	//err = db.Model(&actor3).Limit(1).Select(&actor4)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//}
	//
	//printPostQuery(actor3)
	//printPostQuery(actor4)
	//
	////
	//// Select specific columns into the specified struct
	////
	//actor5 := Actor{
	//	ActorId:   999,
	//	FirstName: "Luke",
	//	LastName:  "Oleson",
	//}
	//
	//printPreQuery(actor5)
	//
	//err = db.Model(&actor5).Column("first_name").Limit(1).Select(&actor5)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//}
	//
	//printPostQuery(actor5)

}

// Singles
// Lists
// SQL keywords
// Quoting
// 	single quotes
// 	double quotes
// Unquoting
//func placeholders(db *pg.DB) {
//	var actor1 Actor
//	printPreQuery(actor1)
//
//	err := db.Model(&actor1).Where("first_name = ?", "PENELOPE").Select()
//	if err != nil {
//		fmt.Println("err: ", err)
//	}
//
//	printPostQuery(actor1)
//}
//
////
//// Aggregate Queries
////
//func columnExpressions(db *pg.DB) {
//	var actor1 Actor
//	printPreQuery(actor1)
//
//	var numOfActors int
//	fmt.Printf("\ncount before: %v \n\n", numOfActors)
//
//	err := db.Model(&actor1).ColumnExpr("sum(?)", pg.Ident("actor_id")).Select(&numOfActors)
//	if err != nil {
//		fmt.Println("err: ", err)
//	}
//
//	fmt.Println("numOfActors: ", numOfActors)
//	printPostQuery(actor1)
//}

//
// Aliases
//

// If you don't add an alias to the object def. it will prefix all columns with the table name
// adding `pg:"alias: a"` creates the alias "a" for the table
// adding `pg:"alias:"` will define NO alias (the column names will not be prefixed.

//
// Constraints
//
// relationalConstraint illustrates that you cannot use a foreign key
// in a table if that foreign key does not exist as a primary key in
// the referenced table.
//func relationalConstraint(db *pg.DB) {
//	contact := Contact{
//		Id:        uuid.NewV4().String(),
//		FirstName: "Justin",
//		LastName:  "Weissig",
//		Account:   "20544fd3-29a1-4780-9008-2b544f25925f",
//		CreatedAt: time.Time{},
//	}
//	_, err := db.Model(&contact).Insert()
//	if err != nil {
//		fmt.Println("error inserting contact. error: ", err)
//	}
//}

//
// Bugs
//

//// returningBug illustrates how the returning clause in the update
//// query writes the same row (the last one it updates) to every record
//// in the accounts array
//func returningBug(db *pg.DB) {
//
//	var accounts []Account
//	ctx := context.Background()
//	err := db.RunInTransaction(ctx, func(tx *pg.Tx) error {
//		err := tx.Model(&accounts).Select()
//		if err != nil {
//			return err
//		}
//
//		var accountIDs []string
//		for _, account := range accounts {
//			accountIDs = append(accountIDs, account.Id)
//		}
//
//		_, err = tx.Model(&accounts).
//			WhereIn("account.id IN (?)", accountIDs).
//			Set("email = ?", "testing6!").
//			Returning("*").
//			Update()
//		if err != nil {
//			return err
//		}
//
//		return nil
//	})
//	if err != nil {
//		fmt.Println("error: ", err)
//	}
//
//	printAccounts(accounts)
//}
//
//func printAccounts(accounts []Account) {
//	fmt.Println("# accounts: ", len(accounts))
//	fmt.Println("Accounts: ")
//	for _, account := range accounts {
//		fmt.Println(account)
//	}
//}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
