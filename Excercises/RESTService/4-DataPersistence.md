# Excercise 4: Persist data in a sqlite database

This is a pretty lengthy excercise, showing the way Go implement the power of interfaces. We will first refactor the in-memory database in a way it implements the interface we will design at the start. After this is done, we will implement sqlite support in the persistence layer.

## Define and implement the interface

Let's start by defining the interface. Open beers.go and add the following interface definition. An interface definition is a list of methods the type should implement. It does not define variables.

```go
// BeerStore defines the methods for Beer Storage implementations
type BeerStore interface {
       PostBeer(nb *NewBeer) (Beer, error)
       GetBeer(i string) (Beer, error)
       GetAllBeers() ([]Beer, error)
}
```

Now we have an interface, we will rename our in memory map to be called something more descriptive.

change

```go
type Beers map[string]Beer
```

to

```go
type InMemoryBeerStore map[string]Beer
```

Now, we will change the existing PostBeer, GetBeer and GetAllBeers methods to be methods of the InMemoryBeerStore type.

```go
// OLD: func PostBeer(nb *NewBeer, db Beers) (Beer, error) {
func (db InMemoryBeerStore) PostBeer(nb *NewBeer) (Beer, error) {
```

```go
// OLD func GetBeer(i string, db Beers) (Beer, error) {
func (bs InMemoryBeerStore) GetBeer(i string) (Beer, error) {
```

Let's also add the GetAllBeers method

```go
func(bs InMemoryBeerStore) GetAllBeers() ([]Beer, error) {
       var beers []Beer
       for _, value := range bs {
               beers = append(beers, value)
       }

       return beers, nil
}
```

Change the SynTappd type to look like this. It will now accept the interface instead of a map of beers

```go
// Syntappd hold the components to run this REST service
type Syntappd struct {
    l *log.Logger
    d beer.BeerStore
}
```

Change the New function to reflect this.

```go
func New(beerstore beer.BeerStore) *Syntappd {
    // Setup Syntappd object
    app := Syntappd{
        l: log.New(os.Stdout, "", 0),
        d: beerstore,
    }
    return &app
}
```

In the PostBeer handler, change this:

```go
// OLD  rb, err := beer.PostBeer(&nb, app.d)
       rb, err := app.d.PostBeer(&nb)
```

In the GetBeer handler, change this:

```go
// OLD b, err := beer.GetBeer(vars["id"], app.d)
    b, err := app.d.GetBeer(vars["id"])
``` 

Remove the for loop from GetAllBeers, we will not use it anymore due to the new GetAllBeers function within the interface. (Which is a different function!)

```go
// REMOVE THIS
-       for _, value := range app.d {
-               beers = append(beers, value)
-       }
```

Last change is in main.go. Pass an InMemoryBeerStore for now, since we renamed the object.

```go
// OLD app := handlers.New(beer.Beers{})
       app := handlers.New(beer.InMemoryBeerStore{})
```

Now is a good time to start the app, and check if adding and retrieving beers still work. Use the curl statements from previous statements.

## Add sqlite support and start using it.

At the top of the sqlite.go file, add the following code

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

Add the following code to the Init function, making sure the beer table will be created if it doesn't exist.

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

Make sure to change the signature, as well as the first line moving the sqlite database to a temporary directory.

```go
func Init() (*SQLiteBeerStore, error) {
       db, err := sqlx.Open("sqlite3", os.Getenv("TEMP")+"/db.sqlite")
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