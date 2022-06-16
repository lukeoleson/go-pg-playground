package examples

import (
	"fmt"
	"testing"

	"github.com/lukeoleson/go-pg-playground/data/dal"
	"github.com/stretchr/testify/assert"

	"github.com/go-pg/pg/v10"
	"github.com/lukeoleson/go-pg-playground/data/models"
)

func TestUpdateMultipleRowsByNonPrimaryKey(t *testing.T) {
	db := dal.Connect()
	defer db.Close()

	testRowIds := []int{3, 4, 136, 179}
	// there are 3 eds and 1 jennifer in the test db
	currentFirstNames := []string{"ED", "JENNIFER"}
	updatedFirstNames := []string{"ANNA", "JOE", "PRINCE", "RALF"}
	actors := []models.Actor{
		{
			FirstName: updatedFirstNames[0],
		},
		{
			FirstName: updatedFirstNames[1],
		},
		{
			FirstName: updatedFirstNames[2],
		},
		{
			FirstName: updatedFirstNames[3],
		},
	}

	// It looks like this query will update all the ED and JENNIFER rows to the names in updatedFirstNames.
	_, err := db.Model(&actors).
		Column("first_name").
		WhereIn("a.first_name IN (?)", currentFirstNames).
		Returning("*").
		Update()
	if err != nil {
		fmt.Println(err)
	}

	var selectedActors []models.Actor
	err = db.Model(&selectedActors).WhereIn("id IN (?)", testRowIds).Select()
	assert.NoError(t, err)
	// The correct number of rows were updated...
	assert.Equal(t, len(testRowIds), len(selectedActors))

	// ... but each row was updated with the first name of the first element from `actors`.
	for _, actor := range selectedActors {
		assert.Equal(t, updatedFirstNames[0], actor.FirstName)
		assert.NotEqual(t, actor.FirstName, updatedFirstNames[1])
	}

	// Question: How would the update query here have known which EDs and which Jennifer's should be updated
	// with which names from updatedFirstNames?

	// Note: The SQL generated by this go-pg query is using a static list of strings
	// in the WHERE clause. We need the WHERE clause to specify that it is looking at the
	// first_name field from the _data table to create a mapping between the rows and
	// the update values.

	// Correct Approach: Add the primary keys to the `actors` objects and use `WherePK()`.
}

//func TestUpdateMultipleRowsByPrimaryKey(t *testing.T) {
//	actor := models.Actor{
//		Id:        40,
//		FirstName: "prince",
//	}
//	actor6 := models.Actor{
//		Id:        200,
//		FirstName: "anna",
//	}
//
//	//currentFirstNames := []string{"PENELOPE", "NICK"}
//
//	_, err := db.Model(&actor5, &actor6).
//		Column("first_name").
//		WherePK().
//		Returning("*").
//		Update()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	var selectedActors []models.Actor
//	err = db.Model(&selectedActors).Select()
//	models.PrintActorsOrError(err, selectedActors)
//}

func Update(db *pg.DB) {

	////////////////////
	//// go-pg can use the values in the struct fed into `Model` as the values for the update.
	//actor1 := Actor{
	//	ActorId:   201,
	//	FirstName: "Rachel", // this value is "Luke" in the database for Actor 202
	//	LastName:  "Oleson",
	//}
	//// As per the documentation, go-pg converts struct field names to lowercase snakecase, so here we tell it to do the update using the value held by actor1.FirstName but go-pg knows it as "first_name".
	//_, err := db.Model(&actor1).Set("first_name = ?first_name").WherePK().Update()
	//// Note that the updated at time is the default Go value when we print this struct - this is because the Update query does NOT return the updated values by default.
	//printActorsOrError(err, []Actor{actor1})
	//
	//////////////////
	//actor2 := Actor{
	//	ActorId:   202,
	//	FirstName: "Megan", // this value is "Luke" in the database for Actor 202
	//	LastName:  "Oleson",
	//}
	//// Including a Returning call will return the updated row into our struct
	//_, err = db.Model(&actor2).Set("first_name = ?first_name").WherePK().Returning("*").Update()
	//// Now we see the Update At columns DB value.
	//printActorsOrError(err, []Actor{actor2})
	//
	///////////////////
	//actor3 := Actor{
	//	ActorId:   201,
	//	FirstName: "Leif",
	//	LastName:  "Oleson",
	//}
	//
	//actor4 := Actor{
	//	ActorId:   202,
	//	FirstName: "Luke",
	//	LastName:  "Oleson",
	//}
	//// As before, the struct passed into the Update method will receive the returned values while the struct passed into the Model method remains unchanged.
	//_, err = db.Model(&actor3).Set("first_name = ?first_name").WherePK().Returning("*").Update(&actor4)
	//// Now we see the Update At columns DB value.
	//printActorsOrError(err, []Actor{actor3})
	//printActorsOrError(err, []Actor{actor4})

	//actor5 := Actor{
	//	FirstName: "prince",
	//}
	//actor6 := Actor{
	//	FirstName: "anna",
	//}
	//actors := []Actor{actor5, actor6}
	//
	//currentFirstNames := []string{"joe", "luke"}
	//
	//_, err := db.Model(&actors).
	//	Column("first_name").
	//	Where("a.first_name IN (?)", pg.In(currentFirstNames)).
	//	Returning("*").
	//	Update()
	//
	//printActorsOrError(err, actors)

}