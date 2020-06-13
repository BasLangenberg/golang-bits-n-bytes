# Excercise 5: Add sqlite support and start using it.

At the top of the sqlite.go file, under the import section, add the following code

```go
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
```

The constants are the fixed, prepared statements we will use to manipulate the table in sqlite. Constants are safe, and they will be compiled into the application for fast access.

Add the following code to the Init function, making sure the beer table will be created if it doesn't exist. This should be added on line 42.

```go
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
```

Make sure to change the signature for the init function.

```go
func Init() (*SQLiteBeerStore, error) {
``` 

Now, satisfy the interface by adding the three functions it required.

```go
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
```

In handlers.go, change the GetAllBeers functions to look like the function below.

```go
// GetAllBeers retrieves all beers in the datastore
func (app *Syntappd) GetAllBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	beers, err := app.d.GetAllBeers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Unable to retrieve stored beers from database"`))
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(beers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Unable to retrieve stored beers"`))
		return
	}
}
```

After all functionality is implemented, re-enable the sqlite in main.go, and pass the sqlite object instead of the in memory beer store to the New function.

```go
       sqlite, err := persistence.Init()
       if err != nil {
               fmt.Printf("FATAL: %+v\n", err)
               os.Exit(1)
       }
       app := handlers.New(sqlite)
```

## Additional excercises
- The validation code is reused in both interface implementation methods. Refactor this to make the code more DRY.
- Move the in memory database to the persistence package as well
