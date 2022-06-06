package examples

import (
	"github.com/go-pg/pg/v10"
	"github.com/lukeoleson/go-pg-playground/db/models"
)

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

	actor5 := models.Actor{
		FirstName: "luke",
	}
	actor6 := models.Actor{
		FirstName: "anna",
	}

	currentFirstNames := []string{"joe", "prince"}

	_, err := db.Model(&actor5, &actor6).
		Column("first_name").
		WhereIn("a.first_name IN (?)", currentFirstNames).
		Returning("*").
		Update()

	models.PrintActorsOrError(err, []models.Actor{actor5, actor6})

}
