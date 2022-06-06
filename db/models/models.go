package models

import (
	"fmt"
	"time"
)

type Actor struct {
	tableName  struct{} `pg:"actor, alias:a"`
	ActorId    int      `pg:"type:uuid,pk"`
	FirstName  string
	LastName   string
	LastUpdate time.Time
}

type Film struct {
	tableName          struct{} `pg:"film, alias:f"`
	FilmId             int
	Title              string
	Description        string
	ReleaseYear        int
	LanguageId         int
	OriginalLanguageId int
	RentalDuration     int
	RentalRate         int
	Length             int
	ReplacementCost    int
	Rating             string // This is an mpaa_rating - what does that mean!?
	LastUpdate         time.Time
	SpecialFeatures    []string
	Fulltext           string // This is a tsvector...whatever that is! Also note the name - we don't want full_text (snakecase) here
}

func PrintActorsOrError(err error, actorSlices ...[]Actor) {

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for _, actors := range actorSlices {
		if len(actors) == 0 {
			fmt.Printf("\nRows Returned\n-------------\n")
			fmt.Printf("none\n")
		} else {
			PrintActors(actors)
		}
	}
}

func PrintActors(actors []Actor) {
	fmt.Printf("\nRows Returned\n-------------\n")

	if len(actors) > 6 {
		var firstHalf []Actor
		firstHalf = append(firstHalf, actors[:3]...)

		secondHalf := actors[3:]
		if len(secondHalf) > 3 {
			secondHalf = secondHalf[len(secondHalf)-3:]
		}

		for _, actor := range firstHalf {
			fmt.Printf("%v\n", actor)
		}

		fmt.Println("...")

		for _, actor := range secondHalf {
			fmt.Printf("%v\n", actor)
		}
	} else {
		for _, actor := range actors {
			fmt.Printf("%v\n", actor)
		}
	}
}

func PrintFilms(films []Film) {
	fmt.Printf("\nResults\n")
	fmt.Printf("-------")

	if len(films) > 6 {
		var firstHalf []Film
		firstHalf = append(firstHalf, films[:3]...)

		secondHalf := films[3:]
		if len(secondHalf) > 3 {
			secondHalf = secondHalf[len(secondHalf)-3:]
		}

		for _, actor := range firstHalf {
			fmt.Printf("%v\n", actor)
		}

		fmt.Println("...")

		for _, actor := range secondHalf {
			fmt.Printf("%v\n", actor)
		}
	} else {
		for _, actor := range films {
			fmt.Printf("%v\n", actor)
		}
	}
}
