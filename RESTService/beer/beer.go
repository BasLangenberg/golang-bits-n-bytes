package beer

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Beer holds values for a single beer object
type Beer struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Brewery string    `json:"brewery"`
	Reviews []Rating  `json:"reviews"`
	AddedOn time.Time `json:"added_on"`
}

// NewBeer defines the data structure required to create a new Beer
type NewBeer struct {
	Name    string `json:"name"`
	Brewery string `json:"brewery"`
}

// Beers acts as an in memory database for Beer objects
type Beers map[string]Beer

// Rating adds a rating for a beer object
type Rating struct {
	Id      string    `json:"id"`
	BeerId  string    `json:"beer_id"`
	AddedOn time.Time `json:"added_on"`
	Score   float32   `json:"score"`
	Review  string    `json:"review"`
}

func PostBeer(nb *NewBeer, db Beers) Beer {

	beer := Beer{
		Id:      uuid.New().String(),
		Name:    nb.Name,
		Brewery: nb.Brewery,
		Reviews: []Rating{},
		AddedOn: time.Now(),
	}

	db[beer.Id] = beer
	return beer
}

func GetBeer(i string, db Beers) (Beer, error) {
	b, ok := db[i]
	if !ok {
		return b, fmt.Errorf("beer id does not exist in database")
	}

	return b, nil
}
