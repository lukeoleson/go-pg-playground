package v9examples

import (
	"fmt"
	"testing"

	"github.com/lukeoleson/go-pg-playground/data/dal"
	"github.com/stretchr/testify/assert"

	"github.com/lukeoleson/go-pg-playground/data/models"
)

func TestV9_UpdateMultipleRowsByNonPrimaryKey_Solution(t *testing.T) {
	db := dal.ConnectV9()
	defer db.Close()

	///////////////////////////////////////////////////////////////////
	// 1. Retrieve the rows you want to update from the DB with the PK
	// (or get the PK some other way into the slice)

	testRowIds := []int{22, 23}
	var actors []models.Actor
	err := db.Model(&actors).
		Column("actor_id").
		WhereIn("actor_id IN (?)", testRowIds).
		Select()
	assert.NoError(t, err)

	// Assertions
	// ...the correct number of rows were retrieved...
	assert.Equal(t, len(testRowIds), len(actors))
	// ...and they contain only the Primary Key...
	for _, actor := range actors {
		assert.NotEmpty(t, actor.ActorId)
		assert.Empty(t, actor.FirstName)
		assert.Empty(t, actor.LastUpdate)
	}

	/////////////////////////////////////////////////////////////////////////
	// 2. Add the values you want updated to the slice and perform the update

	newFirstNames := []string{"ANNA", "JOE"}
	for i := range actors {
		actors[i].FirstName = newFirstNames[i]
	}

	// convert to a slice of interfaces so we can use the variadic operator to feed
	// in all our models to the query
	b := make([]interface{}, len(actors))
	for i := range actors {
		b[i] = actors[i]
	}

	// note that we have to hardcode the WHERE clause here since using something like
	// `Where(a.actor_id IN (?), testRowIds`)` will create a query that filters against
	// a list rather than against the data from our VALUE table.
	// e.g `WHERE a.actor_id IN (1,2)` rather than `WHERE a.actor_id IN _data.actor_id)`.
	_, err = db.Model(b...).
		Column("first_name").
		Where("a.actor_id IN (_data.actor_id)").
		Returning("*").
		Update()
	if err != nil {
		fmt.Println(err)
	}

	// Assertions
	// ...get the latest data from the db
	var selectedActors []models.Actor
	err = db.Model(&selectedActors).WhereIn("actor_id IN (?)", testRowIds).Select()
	assert.NoError(t, err)

	// ...check that the correct number of rows were updated...
	assert.Equal(t, len(testRowIds), len(selectedActors))
	// ...and that the correct row now has the (correct) newly updated first_name
	for i, actor := range selectedActors {
		assert.Equal(t, testRowIds[i], actor.ActorId)
		assert.Equal(t, newFirstNames[i], actor.FirstName)
	}
}
