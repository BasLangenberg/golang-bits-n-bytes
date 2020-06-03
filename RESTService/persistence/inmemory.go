package persistence

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ninckblokje/golang-bits-n-bytes/RESTService/beer"
	"time"
)

// InMemoryBeerStore acts as an in memory database for Beer objects
type InMemoryBeerStore map[string]beer.Beer

func (bs InMemoryBeerStore) PostBeer(nb *beer.NewBeer) (beer.Beer, error) {
	beer := beer.Beer{
		Id:      uuid.New().String(),
		Name:    nb.Name,
		Brewery: nb.Brewery,
		Reviews: []beer.Rating{},
		AddedOn: time.Now(),
	}

	v := validator.New()
	err := v.Struct(beer)
	if err != nil {
		return beer, fmt.Errorf("error validating input: %+v", err)
	}

	bs[beer.Id] = beer
	return beer, nil
}

func (bs InMemoryBeerStore) GetBeer(i string) (beer.Beer, error) {
	b, ok := bs[i]
	if !ok {
		return b, fmt.Errorf("beer id does not exist in database")
	}

	return b, nil
}

func(bs InMemoryBeerStore) GetAllBeers() ([]beer.Beer, error) {
	var beers []beer.Beer
	for _, value := range bs {
		beers = append(beers, value)
	}

	return beers, nil
}
