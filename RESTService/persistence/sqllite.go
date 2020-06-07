package persistence

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ninckblokje/golang-bits-n-bytes/RESTService/beer"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	// sqlite driver will be called using sqlx helper library
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteBeerStore struct {
	db *sqlx.DB
}

const (
	beertable string = `CREATE TABLE IF NOT EXISTS BEERS(
	id TEXT NOT NULL,
	beer TEXT NOT NULL,
	brewery TEXT NOT NULL,
	added_on DATE NOT NULL
)`

	beerinsert string = `INSERT INTO BEERS (id, beer, brewery, added_on) VALUES (?, ?, ?, ?)`

	beerget string = `SELECT * FROM BEERS WHERE id = ?`

	beergetall string = `SELECT * FROM BEERS`

)



// Init initialized db if not existing, does schema creation and returns a pointer to the sqlx.DB instance
func Init() (*SQLiteBeerStore, error) {
	db, err := sqlx.Open("sqlite3", os.Getenv("TEMP")+"/db.sqlite")
	if err != nil {
		return nil, fmt.Errorf("unable to open sqlite database: %+v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to connect to sqlite database")
	}

	// Create table
	// In real life we should do more robuust migrations. ;-)
	rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table' AND name='BEER';`)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to query sql database")
	}

	log.Println("creating database tables if not already existing")
	db.Exec(beertable)

	sqlitebs := SQLiteBeerStore{
		db: db,
	}

	return &sqlitebs, nil
}

func (bs SQLiteBeerStore) PostBeer(nb *beer.NewBeer) (beer.Beer, error) {
	br := beer.Beer{
		Id:      uuid.New().String(),
		Name:    nb.Name,
		Brewery: nb.Brewery,
		Reviews: []beer.Rating{},
		AddedOn: time.Now(),
	}

	v := validator.New()
	err := v.Struct(br)
	if err != nil {
		return br, fmt.Errorf("error validating input: %+v", err)
	}

	_, err = bs.db.Exec(beerinsert, br.Id, br.Name, br.Brewery, br.AddedOn)
	if err != nil {
		return beer.Beer{}, fmt.Errorf("unable to insert br into database: %+v", err)
	}

	return br, nil
}

func (bs SQLiteBeerStore) GetBeer(i string) (beer.Beer, error) {
	b := bs.db.QueryRow(beerget, i)
	br := beer.Beer{}
	b.Scan(&br.Id, &br.Name, &br.Brewery, &br.AddedOn)
	return br, nil
}

func(bs SQLiteBeerStore) GetAllBeers() ([]beer.Beer, error) {
	var beers []beer.Beer

	rows, err := bs.db.Query(beergetall)
	if err != nil {
		return beers, fmt.Errorf("unable to get beers : %+v", err)
	}

	for rows.Next(){
		var br beer.Beer
		err = rows.Scan(&br.Id, &br.Name, &br.Brewery, &br.AddedOn)
		if err != nil {
			return beers, fmt.Errorf("unable to parse beers : %+v", err)
		}

		beers = append(beers, br)
	}


	return beers, nil
}

