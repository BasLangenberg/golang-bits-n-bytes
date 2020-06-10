package beer

import (
	"time"
)

// BeerStore defines the methods for Beer Storage implementations
type BeerStore interface {
	PostBeer(nb *NewBeer) (Beer, error)
	GetBeer(i string) (Beer, error)
	GetAllBeers() ([]Beer, error)
}

// Beer holds values for a single beer object
type Beer struct {
	Id      string    `json:"id"`
	Name    string    `json:"name" validate:"required"`
	Brewery string    `json:"brewery" validate:"required"`
	Reviews []Rating  `json:"reviews"`
	AddedOn time.Time `json:"added_on"`
}

// NewBeer defines the data structure required to create a new Beer
type NewBeer struct {
	Name    string `json:"name"`
	Brewery string `json:"brewery"`
}

// Rating adds a rating for a beer object
type Rating struct {
	Id      string    `json:"id"`
	BeerId  string    `json:"beer_id"`
	AddedOn time.Time `json:"added_on"`
	Score   float32   `json:"score"`
	Review  string    `json:"review"`
}
