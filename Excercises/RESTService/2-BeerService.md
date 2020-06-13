# Excercise 2: Implement beer service objects

Now that the base application is running smoothly, it's time to add some serious application logic to it. We want to write a small service used to add beers to a database. Very original, we know.

Let's start by implementing the beer business structs we need to create, store and retrieve 

From the Scaffold directory, create a directory called beer. Add a beer.go file in this directory and add the following content.

```go
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

// PostBeer creates a new beer, and adds it to the database
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

// GetBeer returns a beer from the database, using its uuid
func GetBeer(i string, db Beers) (Beer, error) {
	b, ok := db[i]
	if !ok {
		return b, fmt.Errorf("beer id does not exist in database")
	}

	return b, nil
}

```

It's now time to refactor the handlers.go file. We need to add methods, or better said, http handlers to add and retrieve beers. Handlers are a short hand for functions in Go that implement the handler interface. If you satisfy this interface, you can pass these functions to a http router / server which can then use it to handle web requests.

Add the following struct to the handlers.go file.

```go
// Syntappd hold the components to run this REST service
type Syntappd struct {
       l *log.Logger
       d beer.Beers
}
```

We will use this struct to add our handlers to, and use it as a place to store our logger and beer database.

The New function will be refactored. No longer we will return a router, we will return a SynTappd instance. Replace the function with the following.

```go
// db should be converted to an interface / real database implementation
func New(beers beer.Beers) *Syntappd {
       // Setup Syntappd object
       app := Syntappd{
               l: log.New(os.Stdout, "", 0),
               d: beers,
       }

       return &app
}
```

Change the signature of the Healthcheck and Home functions, so they are no longer functions but methods of the SynTappd struct. The way we do is by adding the struct before the function name.

```go
// Healthcheck returns the health of the overall system
func (app *Syntappd) Healthcheck(w http.ResponseWriter, r *http.Request) {
```

```go
// Home returns the homepage which is not really used
func (app *Syntappd) Home(w http.ResponseWriter, r *http.Request) {
```

Now we will add three more functions to this file, used to save and retrieve one, or all beers from the in memory database.

```go
//PostBeer creates a new beer and stores in the datastore
func (app *Syntappd) PostBeer(w http.ResponseWriter, r *http.Request) {
       var nb beer.NewBeer

       w.Header().Set("Content-Type", "application/json")

       req, err := ioutil.ReadAll(r.Body)
       if err != nil {
               w.WriteHeader(http.StatusInternalServerError)
               w.Write([]byte(`{ "error": "Unable to read body" }`))
               return
       }

       err = json.Unmarshal(req, &nb)
       if err != nil {
               w.WriteHeader(http.StatusBadRequest)
               w.Header().Set("Content-Type", "application/json")
               w.Write([]byte(`{ "error": "invalid request body"`))
               return
       }

       rb := beer.PostBeer(&nb, app.d)

       json.NewEncoder(w).Encode(rb)

}
```

```go
// GetBeer retrieves one individual beer
func (app *Syntappd) GetBeer(w http.ResponseWriter, r *http.Request) {
       vars := mux.Vars(r)

       w.Header().Set("Content-Type", "application/json")

       b, err := beer.GetBeer(vars["id"], app.d)
       if err != nil {
               w.WriteHeader(http.StatusNotFound)
               w.Write([]byte(`{ "error": "beer does not exist" }`))
               return
       }

       err = json.NewEncoder(w).Encode(b)
       if err != nil {
               w.WriteHeader(http.StatusInternalServerError)
               w.Write([]byte(`{ "error": "Unable to retrieve stored beer"`))
               return
       }
}

// GetBeer retrieves all beers in the datastore
func (app *Syntappd) GetAllBeers(w http.ResponseWriter, r *http.Request) {
       var beers []beer.Beer

       w.Header().Set("Content-Type", "application/json")

       for _, value := range app.d {
               beers = append(beers, value)
       }

       w.WriteHeader(http.StatusOK)
       err := json.NewEncoder(w).Encode(beers)
       if err != nil {
               w.WriteHeader(http.StatusInternalServerError)
               w.Write([]byte(`{ "error": "Unable to retrieve stored beers"`))
               return
       }
}
```

Now we need to tie the new code together and add it to the main function. We start with commenting out the sqlite code. Go will not compile if you create a struct or variable you will not use. We will implement sqlite in excercise 4. We will also change how we instantiate our routers and remove the r := handlers.New(db) call.

```go
       // Will be used later in the tutorial
       //db, err := persistence.Init()
       //if err != nil {
       //      fmt.Printf("FATAL: %+v\n", err)
       //      os.Exit(1)
       //}

       app := handlers.New(beer.Beers{})

       // Setup Handlers
       r := mux.NewRouter()
       r.HandleFunc("/", app.Home)
       r.HandleFunc("/health", app.Healthcheck)
       r.HandleFunc("/beer", app.PostBeer).Methods("POST")
       r.HandleFunc("/beer", app.GetAllBeers).Methods("GET")
       r.HandleFunc("/beer/{id}", app.GetBeer).Methods("GET")

       	srv := http.Server{
		Handler: r,
		Addr:    "localhost:8080",
	}

        log.Println("Application started succesfully!")

	log.Fatal(srv.ListenAndServe())
```

## Test the beer logic

You can test if your application works using the following curl logic.

```bash
curl --location --request POST 'http://localhost:8080/beer' \
--header 'Content-Type: application/json' \
--data-raw '{ "brewery": "SynBeer Lab", "name": "Platinum Blond" }'

curl --location --request POST 'http://localhost:8080/beer' \
--header 'Content-Type: application/json' \
--data-raw '{ "brewery": "SynBeer Lab", "name": "Golden Tripel" }'
```

```bash
curl --location --request GET 'localhost:8080/beer'

# Change uuid to one of the uuid from one of the beers that were returned in the previous call
curl --location --request GET 'localhost:8080/beer/92a63846-414a-4022-b15a-8a40ab21a520'
```

If you want to see this in a more formatted view, install jq. That will format json for you.

```bash
bas@DESKTOP-RFVONSL: /mnt/c/data/golang-bits-n-bytes $ curl --location --request GET 'localhost:8080/beer' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   309  100   309    0     0  34333      0 --:--:-- --:--:-- --:--:-- 34333
[
  {
    "id": "d8e9255f-efa3-42f1-a293-9a63f67bf37e",
    "name": "Platinum Blond",
    "brewery": "SynBeer Lab",
    "reviews": [],
    "added_on": "2020-06-13T17:06:32.2683288+02:00"
  },
  {
    "id": "343bedf8-74fc-4a8b-9046-223245fa9227",
    "name": "Golden Tripel",
    "brewery": "SynBeer Lab",
    "reviews": [],
    "added_on": "2020-06-13T17:06:33.0233902+02:00"
  }
]
```

## Additional excercises

After you are done with this tutorial and you feel like you want to do more hands on Go programming, it might be interesting to:
- Implement updating a beer
- Implement deleting a beer